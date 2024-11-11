package commu_module

import (
	"crypto/ed25519"
)

// 최초 등록시 필요한 struct
type nodeAccountEntity struct {
	Address   string `json:"address"`
	Pubkey    string `json:"pubkey"`
	Signature string `json:"signature(sK(Address))"`
}

type ServerList struct {
	interfaceList []string
	redisList     []string
}

type LeaveAccountEntity struct {
	Address   string
	Pubkey    string
	Signature string
}

/* type vrfValueEntity struct {
	Val    string
	Proof  []byte
	PubKey []byte
} */

// committee request할 때 보내는 데이터
// 해당 데이터는 grpc에 정의
/* type CommitteeNodeInfo struct {
	Round     int32          `json:"round"`
	Address   string         `json:"address"`
	VrfPubKey []byte         `json:"vrfpubkey"`
	VrfResult vrfValueEntity `json:"vrfresult"`
} */

// committee request에 대한 response로 받는 데이터
type CommitteeInfo struct {
	AggregateCommit []byte
	AggregatePubKey []byte
	//CommitteeList   []CommitteeNodeInfo
	PrimaryNodeInfo string
	isLeader        bool
}

type KeyPair struct {
	PublicKey ed25519.PublicKey
	SecretKey ed25519.PrivateKey
}

var globalKeyPair KeyPair
var serverList ServerList
var nodeAccount nodeAccountEntity
var serverSelectionNum int32 = 0
var storedNodeIP string

// API로 사용되는 메인 함수들

// nodeIP주소는 블록체인 노드 자기자신의 주소
// servers는 interface목록 및 redis 서버목록
// 일단 interface server 리더 선정방식에 대해 정해진 바가 없으므로 단일 서버로 테스트 진행
func JoinNetwork(nodeIP string, servers ServerList) {
	// 들어온 데이터를 grpc써서 interface server에 저장요청 보내기
	serverList = servers
	storedNodeIP = nodeIP

	generateGlobalKeyPair()
	go subscriptionCommitteeListChannel()
	requestEnrollNodeDataToInterface(storedNodeIP)
}

// vrf 실행하고 결과값 포함해서 보내기
func ReqeustSetupCommittee(round int32) {
	requestSetupCommitteeToInterface(round)
}

func LeaveNetwork(nodeData LeaveAccountEntity) {

	// redis등 구독정보를 해제하도록 설정
}
