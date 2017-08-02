package performance

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
)

var cache map[string]interface{}

//Scene 场景
type Scene struct {
	Name       string
	Interfaces []Interface
}

//Interface 场景中的接口
type Interface struct {
	Name    string
	URL     string
	Headers map[string]string
	Method  string
	Store   []Store
	Body
}

// Store save some response data to global
type Store struct {
	Type     int8
	FromKey  string
	ToKey    string
	JSONPath string
	Value    string
}

// Body 请求体
type Body struct {
	Content string
}

func (i *Interface) getContentByType() (io.Reader, error) {
	contentType, ok := i.Headers["Content-Type"]
	if !ok {
		return nil, fmt.Errorf("headers has not Content-Type")
	}
	switch contentType {
	case "application/json":
		return bytes.NewBuffer([]byte(i.Body.Content)), nil
	default:
		return nil, fmt.Errorf("body type %s did not surport", contentType)
	}
}

//JSONToStruct json to struct
func JSONToStruct(str string) (*[]Scene, error) {
	var scenes []Scene
	jsonBlob := []byte(str)
	err := json.Unmarshal(jsonBlob, &scenes)
	if err != nil {
		return nil, err
	}
	return &scenes, nil
}

// func (s *Scene) Run() error {
// 	for _, inter := range s.Interfaces {
// 		res, err := inter.Request()
// 		if err != nil {
// 			return err
// 		}
// 	}
// 	res
// 	return nil
// }
