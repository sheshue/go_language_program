package ipc

import (
	"encoding/json"
)

type IpcClient struct {
	conn chan string
}

func NewIpcClient(server *IpcServer) *IpcClient {
	c := server.Connect()
	return &IpcClient{c}
}

func (client *IpcClient) Call(method, params string) (resp *Response, err error) {
	req := &Request{method, params}
	var b []byte
	var resp1 Response
	var resp2 Response
	b, err = json.Marshal(req)
	if err != nil {
		resp = &resp1
		return
	}
	client.conn <- string(b)
	str := <-client.conn //等待返回值
	err = json.Unmarshal([]byte(str), &resp2)
	resp = &resp2
	return
}

func (client *IpcClient) Close() {
	client.conn <- "CLOSE"
}
