package communication_module

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

var interfaceList = GetAddressList("interface-list.txt")
var redisList = GetAddressList("redis-list.txt")

// API로 사용되는 메인 함수들

// JoinNetwork하면 해당 struct정보를
func JoinNetwork(nodeData EnrollAccountEntity) {
	// 들어온 데이터를 protobuf써서 interface server에 저장요청 보내기
	// 데이터 등록하면서 redis 구독작업 미리 하기
}

func ReqeustSetupCommittee(requestData RequestSetupcommittee) {
	//
}

func LeaveNetwork(nodeData LeaveAccountEntity) {
	// redis등 구독정보를 해제하도록 설정
}
