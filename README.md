# commu_module

Interface서버와 통신해 커미티 요청등을 할 수 있게 API를 지원하는 module입니다.

---

## Functions(APIs)

### `func JoinNetwork`

**Syntax**:

```go
func JoinNetwork(nodeIP string, servers ServerList) bool
```

```go
type ServerList struct {
	InterfaceList []string
	RedisList     []string
}
```

Network에 참여할 때 해당 함수를 호출합니다. 호출 할때, parameter로 노드 자신의 IP와 통신할 interface server 및 redis의 주소의 리스트를 보냅니다. SeverList구조체는 위에 보이는 바와 같습니다.

### `func RequestSetupCommittee`

**Syntax**:

```go
func ReqeustSetupCommittee(round int32) CommitteeInfo
```

```go
type vrfValue struct {
	Val    string `json:"val"`
	Proof  []byte `json:"pi"`
	PubKey []byte `json:"pk"`
}

type CommitteeNodeInfo struct {
	Round     int32    `json:"round"`
	Address   string   `json:"address"`
	VrfPubKey []byte   `json:"vrfpubkey"`
	VrfResult vrfValue `json:"vrfresult"`
}

type CommitteeInfo struct {
	AggregateCommit []byte              `json:"aggcommit"`
	AggregatePubKey []byte              `json:"aggpubkey"`
	CommitteeList   []CommitteeNodeInfo `json:"committeelist"`
	PrimaryNodeInfo string
}
```

커미티를 요청하는데 사용되는 함수입니다. 요청 시 현재 round값을 보내면 됩니다. 현재는 vrf값을 생성해 threshold값 이내라면 선택되어 committee에 참여시키는 것으로 구현되어 있습니다.

완성된 커미티 결과값은 위에 보이는 `CommitteeInfo` 구조체에 담겨 반환됩니다.

### `func LeaveNetwork`

**Syntax**:

```go
func LeaveNetwork()
```

노드가 Network에서 이탈할 때 호출할 함수입니다. 해당 함수는 반환값이 없으며, 이 함수를 호출하는 것으로 interface server의 DB에 저장되어 있던 노드 정보를 삭제합니다.