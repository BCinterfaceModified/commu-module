package commu_module

import (
	"crypto/ed25519"
	"fmt"
)

type EnrollAccountEntity struct {
	Type      string `json:"type"`
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

type RequestSetupcommittee struct {
	Round int32
}

type CommitteeInfo struct {
	AggregatePubKey []byte
	CommitteeList   []EnrollAccountEntity
	PrimaryNodeInfo string
	isLeader        bool
}

type KeyPair struct {
	PublicKey ed25519.PublicKey
	SecretKey ed25519.PrivateKey
}

var globalKeyPair KeyPair
var serverList ServerList
var serverSelectionNum int32 = 0

// API로 사용되는 메인 함수들

// JoinNetwork하면 해당 struct정보를
func JoinNetwork(nodeData EnrollAccountEntity, servers ServerList) {
	// 들어온 데이터를 protobuf써서 interface server에 저장요청 보내기
	serverList = servers
	generateGlobalKeyPair()
	go subscriptionCommitteeListChannel()
	requestEnrollNodeDataToInterface(nodeData)
}

// vrf 실행하고 결과값 포함해서 보내기
func ReqeustSetupCommittee(requestData RequestSetupcommittee) {

	requestSetupCommitteeToInterface(requestData.Round)
}

func LeaveNetwork(nodeData LeaveAccountEntity) {
	fmt.Println("module test: Leave Network")
	generateGlobalKeyPair()
	_, _, ratio := generateVrfOutput("LEAVE")

	fmt.Println("VRF Ratio : ", ratio)
	// redis등 구독정보를 해제하도록 설정
}