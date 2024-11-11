package commu_module

import (
	"context"
	"crypto/ed25519"
	"crypto/sha256"
	"fmt"
	"log"
	"math/big"
	"strings"
	"time"

	pb "github.com/BCinterfaceModified/commu-module/bcinterface"
	"github.com/BCinterfaceModified/commu-module/vrfs"
	"github.com/gomodule/redigo/redis"
	"google.golang.org/grpc"
)

func publishMessageToRedis(channelName string, message []byte) {
	redisHost, redisPort := parseAddress(serverList.redisList[serverSelectionNum])
	c, _ := redis.DialURL("redis://" + redisHost + redisPort)
	if c == nil {
		fmt.Println("Error Occured: PublishMessageToRedis")
	}

	c.Do("PUBLISH", channelName, message)
}

func parseAddress(fullAddress string) (string, string) {
	parts := strings.SplitN(fullAddress, ":", 2)

	//Error check
	if len(parts) != 2 {
		fmt.Println(`Invalid Address Format!
		Address must be in the following format: address:port`)
	}

	//parts[0]: Host, parts[1]: Port
	return parts[0], parts[1]
}

func generateGlobalKeyPair() {
	pk, sk, err := ed25519.GenerateKey(nil)
	if err != nil {
		log.Println("Error in generate PkSk: ", err)
	}

	globalKeyPair.SecretKey = sk
	globalKeyPair.PublicKey = pk
}

func hashRatio(vrfOutput []byte) float64 {
	t := &big.Int{}
	t.SetBytes(vrfOutput[:])

	precision := uint(8 * (len(vrfOutput) + 1))
	max, b, err := big.ParseFloat("0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff", 0, precision, big.ToNearestEven)
	if b != 16 || err != nil {
		log.Fatal("failed to parse big float constant for sortition")
	}

	//hash value as int expression.
	//hval, _ := h.Float64() to get the value
	h := big.Float{}
	h.SetPrec(precision)
	h.SetInt(t)
	//https://stackoverflow.com/questions/13582519/how-to-generate-hash-number-of-a-string-in-go
	ratio := big.Float{}
	cratio, _ := ratio.Quo(&h, max).Float64()

	return cratio
}

// seed must never be exposed to public.
// 해당함수가 호출되기 전에는 반드시 generateGlobalKeyPair에 의해 pk와 sk가 생성되어 있어야한다.
func generateVrfOutput(seed string) ([]byte, []byte, float64) {
	fmt.Println("VRF output function input seed: ", seed)

	hashedSeed := sha256.Sum256([]byte(seed))

	vrfProof, vrfOutput, err := vrfs.Prove(globalKeyPair.PublicKey, globalKeyPair.SecretKey, hashedSeed[:])
	if err != nil {
		log.Fatal(err)
	}
	vrfResult, err := vrfs.Verify(globalKeyPair.PublicKey, vrfProof, hashedSeed[:])
	if err != nil {
		log.Fatal(err)
	}
	if !vrfResult {
		fmt.Println("Error in VRF Result")
	}

	vrfRatio := hashRatio(vrfOutput)

	return vrfProof, vrfOutput, vrfRatio
}

func dialGrpcConnection() pb.BlockchainInterfaceClient {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	conn, err := grpc.DialContext(ctx, serverList.interfaceList[serverSelectionNum], grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Println("did not connect to server:", err)
		return nil
	}

	client := pb.NewBlockchainInterfaceClient(conn)
	return client
}

func sortition(ratio float64) bool {
	sortitionThreshold := 1.0

	fmt.Println(ratio)
	if ratio > sortitionThreshold {
		return false
	}
	return true
}
