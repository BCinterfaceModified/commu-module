package commu_module

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gomodule/redigo/redis"
)

func GetAddressList(fileName string) []string {
	file, err := os.Open("/usr/local/bin/" + fileName)
	if err != nil {
		log.Fatalf("can't open file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// 각 줄을 저장할 슬라이스 생성
	var lines []string

	// 파일을 줄 단위로 읽어서 슬라이스에 저장
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// 스캐너 에러 확인
	if err := scanner.Err(); err != nil {
		log.Fatalf("failed to read file: %v", err)
	}

	return lines
}

func PublishMessageToRedis(channelName string, message []byte) {
	redisHost, redisPort := ParseAddress(redisList[0])
	c, _ := redis.DialURL("redis://" + redisHost + redisPort)
	if c == nil {
		fmt.Println("Error Occured: PublishMessageToRedis")
	}

	c.Do("PUBLISH", channelName, message)
}

func ParseAddress(fullAddress string) (string, string) {
	parts := strings.SplitN(fullAddress, ":", 2)

	//Error check
	if len(parts) != 2 {
		fmt.Println(`Invalid Address Format!
		Address must be in the following format: address:port`)
	}

	//parts[0]: Host, parts[1]: Port
	return parts[0], parts[1]
}
