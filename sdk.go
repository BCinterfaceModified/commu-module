package commu_module

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	pb "github.com/BCinterfaceModified/commu-module/bcinterface"

	"github.com/gomodule/redigo/redis"
)

var recvCommitteeInfo CommitteeInfo
var grpcResult chan int32 = make(chan int32)

// 채널명: CommitteeList -> interface server로부터 redis 통해 committee list 수신함
func subscriptionCommitteeListChannel() {
	redisConnect, err := redis.Dial("tcp", serverList.redisList[serverSelectionNum])
	if err != nil {
		log.Println("Error occured when subscription interface: ", err)
	} else {
		log.Println("Success Subscription interface server")
	}
	psc := redis.PubSubConn{Conn: redisConnect}
	psc.Subscribe("CommitteeList")
	for {
		switch v := psc.Receive().(type) {
		case redis.Message:
			if string(v.Data) == "CHANGE_REDIS_ADDRESS" {
				log.Println("receive chage the redis address, break!")
				redisConnect.Close()
				return
			}

			recvCommitteeInfo = CommitteeInfo{}

			json.Unmarshal(v.Data, &recvCommitteeInfo)
		}
	}
}

// proto파일의 EnrollAccount 함수를 이용해서 interface server에 Join한 노드정보 저장
func requestEnrollNodeDataToInterface(enrollAccountEntity EnrollAccountEntity) {
	client := dialGrpcConnection()

	for {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		r, err := client.EnrollAccount(ctx, &pb.EnrollAccountRequest{
			Type:      enrollAccountEntity.Type,
			Address:   enrollAccountEntity.Address,
			Pubkey:    []byte(enrollAccountEntity.Pubkey),
			Signature: []byte(enrollAccountEntity.Signature),
		})
		if err != nil {
			log.Println("ERROR :", err)
			cancel()
		} else {
			fmt.Println(r.GetCode())
			grpcResult <- r.GetCode()
			cancel()
			defer close(grpcResult)
			break
		}
	}
}

func requestSetupCommitteeToInterface(round int32) {
	client := dialGrpcConnection()

	for {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		r, err := client.SetupCommittee(ctx, &pb.SetupCommitteeRequest{
			Round: round,
		})
		if err != nil {
			log.Println("ERROR :", err)
			cancel()
		} else {
			fmt.Println(r.GetCode())
			grpcResult <- r.GetCode()
			cancel()
			defer close(grpcResult)
			break
		}
	}
}
