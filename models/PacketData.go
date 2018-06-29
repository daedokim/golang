package models

// PacketData 는 패킷데이터의 구조체이다
type PacketData struct {
	PacketNum int         `json:"packetNum"`
	Data      interface{} `json:"data"`
	Error     error       `json:"error"`
}
