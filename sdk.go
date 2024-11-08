package commu_module

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	pb "github.com/BCinterfaceModified/commu-module/bcinterface"
	"google.golang.org/grpc"

	"github.com/gomodule/redigo/redis"
)

var recvCommitteeInfo CommitteeInfo
var grpcResult chan int32 = make(chan int32)

// 채널명: CommitteeList -> interface server로부터 redis 통해 committee list 수신함
func SubscriptionCommitteeListChannel() {
	redisConnect, err := redis.Dial("tcp", redisList[serverSelectionNum])
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

func RequestSetupCommitteeToInterface(round int32) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)

	conn, err := grpc.DialContext(ctx, interfaceList[serverSelectionNum], grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Println("did not connect to client server(RequestCommitteeSetupNormal):", err)

		cancel()
		return
	} else {
		c := pb.NewBlockchainInterfaceClient(conn)

		for {
			ctx, cancel = context.WithTimeout(context.Background(), time.Second)
			r, err := c.SetupCommittee(ctx, &pb.SetupCommitteeRequest{
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

}
