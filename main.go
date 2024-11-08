package commu_module

type EnrollAccountEntity struct {
	Type      string `json:"type"`
	Address   string `json:"address"`
	Pubkey    string `json:"pubkey"`
	Signature string `json:"signature(sK(Address))"`
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

var interfaceList = GetAddressList("interface-list.txt")
var redisList = GetAddressList("redis-list.txt")
var serverSelectionNum int32 = 0

// API로 사용되는 메인 함수들

// JoinNetwork하면 해당 struct정보를
func JoinNetwork(nodeData EnrollAccountEntity) {
	// 들어온 데이터를 protobuf써서 interface server에 저장요청 보내기
	go SubscriptionCommitteeListChannel()
}

func ReqeustSetupCommittee(requestData RequestSetupcommittee) {
	RequestSetupCommitteeToInterface(requestData.Round)
}

func LeaveNetwork(nodeData LeaveAccountEntity) {
	// redis등 구독정보를 해제하도록 설정
}
