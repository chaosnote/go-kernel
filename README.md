# 工具說明

## 撰寫規則
    expect(Any).To(Be)

## 待修改

功能
1. builder

不在使用 base
1. Builder 後反傳 interface

改為 inject 功能至指定對像
1. 例 : game, server    

修改至新物件擴充類別
1. "github.com/golang/protobuf/jsonpb"
1. "github.com/golang/protobuf/proto"

待補
1. "go get github.com/google/uuid"

## ( GO )輔助指命
    go vet -json main.go
    delve debug main.go