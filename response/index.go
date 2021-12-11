package response

import "encoding/json"

type N int64

const (
	OK N = 0
)

/*
Info
*/
type Info struct {
	Code    N      // 代碼 {成功:0,失敗代碼:非 0 值}
	Message []byte // 備用資訊( 例 : API 回應錯誤代碼 )
	Content []byte // 例: event + grpc
}

/*
Marshal

	轉換過程不應出現錯誤、強制關閉

*/
func (v Info) Marshal() []byte {
	r, e := json.Marshal(v)
	if e != nil {
		panic(e)
	}
	return r
}

/*
Unmarshal
*/
func Unmarshal(data []byte) (Info, error) {
	var r Info
	err := json.Unmarshal(data, &r)
	return r, err
}
