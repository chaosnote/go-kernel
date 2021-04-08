package content

// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    simple, err := UnmarshalSimple(bytes)
//    bytes, err = simple.Marshal()

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// UnmarshalModel ...
func UnmarshalModel(data []byte) (Model, error) {
	var r Model
	err := json.Unmarshal(data, &r)
	return r, err
}

// Marshal ...
func (v *Model) Marshal() ([]byte, error) {
	return json.Marshal(v)
}

// Model ...
type Model struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

//-----------------------------------------------------------------------------

// Status ...
type Status string

// ToString ...
func (v Status) ToString() string {
	return string(v)
}

// OK ...
const OK Status = "OK"

//-----------------------------------------------------------------------------

// Write ...
func (v *Model) Write(w http.ResponseWriter, r *http.Request) {
	json, err := v.Marshal()
	if err != nil {
		// 記錄底層錯誤
		http.NotFound(w, r)
		return
	}
	fmt.Fprint(w, string(json))
}

// New ...
func New(c Status) *Model {
	return &Model{
		Code: c.ToString(),
	}
}

// NewMessage ...
func NewMessage(c Status, m string) *Model {
	return &Model{
		Code:    c.ToString(),
		Message: m,
	}
}

// NewData ...
func NewData(c Status, d interface{}) *Model {
	return &Model{
		Code: c.ToString(),
		Data: d,
	}
}
