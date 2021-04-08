# 底層方法

## 函式定義規則
    expect(Any).To(Be)

## 結構修改
    功能
        builder
    不在使用 base
        Builder 後反傳 interface
    改為 inject 功能至指定對像
        例 : game, server
    log 改為一層傳遞一層，建構時需設置
    server 加入 main
    ↑  同一件事，應修改為

    "github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
        修改至新物件擴充類別
    "go get github.com/google/uuid"
        待補



## ( GO )輔助工具
    go vet -json main.go
    delve debug main.go