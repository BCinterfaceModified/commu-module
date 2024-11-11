package commu_module

import (
	"context"
	"crypto/ed25519"
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
// interface는 해당 요청 받을 시 mongodb에 해당 데이터 저장하도록 수행.
func requestEnrollNodeDataToInterface(nodeIP string) {
	client := dialGrpcConnection()

	//등록할 signature 생성(globalkeypair 기반)
	signature := ed25519.Sign(globalKeyPair.SecretKey, []byte(nodeIP))

	for {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		r, err := client.EnrollAccount(ctx, &pb.EnrollAccountRequest{
			Address:   nodeIP,
			Pubkey:    globalKeyPair.PublicKey,
			Signature: signature,
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
	seed := "test"
	vrfProof, _, vrfRatio := generateVrfOutput(seed)

	if !sortition(vrfRatio) {
		log.Println("VRF ratio can't meet threshold")
		return
	}

	client := dialGrpcConnection()

	for {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		r, err := client.SetupCommittee(ctx, &pb.SetupCommitteeRequest{
			Round:     round,
			Nodeip:    storedNodeIP,
			Vrfpubkey: globalKeyPair.PublicKey,
			VrfResult: &pb.VrfValue{
				Val:    seed,
				Proof:  vrfProof,
				Pubkey: globalKeyPair.PublicKey,
			},
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
