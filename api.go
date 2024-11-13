package commu_module

import (
	"context"
	"crypto/ed25519"
	"log"
)

// 최초 등록시 필요한 struct
/* type nodeAccountEntity struct {
	Address   string `json:"address"`
	Pubkey    string `json:"pubkey"`
	Signature string `json:"signature(sK(Address))"`
} */

type ServerList struct {
	InterfaceList []string
	RedisList     []string
}

/* type LeaveAccountEntity struct {
	Address   string
	Pubkey    string
	Signature string
} */

type KeyPair struct {
	PublicKey ed25519.PublicKey
	SecretKey ed25519.PrivateKey
}

var (
	globalCtx    context.Context
	globalCancel context.CancelFunc
)

var globalKeyPair KeyPair
var serverList ServerList

// var nodeAccount nodeAccountEntity
var serverSelectionNum int32 = 0
var storedNodeIP string

// API로 사용되는 메인 함수들

// nodeIP주소는 블록체인 노드 자기자신의 주소
// servers는 interface목록 및 redis 서버목록
// 일단 interface server 리더 선정방식에 대해 정해진 바가 없으므로 단일 서버로 테스트 진행

// nodeIP는 노드의 ip주소를 보내주시고, servers에는
// 패키지 내 ServerList 구조체에 값을 채워서 보내주시면 됩니다.
func JoinNetwork(nodeIP string, servers ServerList) bool {
	// 들어온 데이터를 grpc써서 interface server에 저장요청 보내기
	serverList = servers
	storedNodeIP = nodeIP

	generateGlobalKeyPair()

	globalCtx, globalCancel = context.WithCancel(context.Background())
	go subscriptionCommitteeListChannel()
	requestEnrollNodeDataToInterface(storedNodeIP)

	return true
}

// 현재 Round값만 채워서 보내주시면 됩니다.
func ReqeustSetupCommittee(round int32) (recvCommittee CommitteeInfo) {
	requestSetupCommitteeToInterface(round)

	return recvCommitteeInfo
}

// node IP가 LEAVE까지 변경되지 않았을 것을 가정합니다.
func LeaveNetwork() {
	// redis 구독정보 해제
	if globalCancel != nil {
		globalCancel()
	} else {
		log.Println("No active subscription to cancel")
	}

	requestLeaveNodeToInterface()
}
