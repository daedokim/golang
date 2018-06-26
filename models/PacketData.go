package models

// PacketData 는 패킷데이터의 구조체이다
type PacketData struct {
	PacketNum  int         `json:"packetNum"`
	PacketData interface{} `json:"packetData"`
	Error      Error       `json:"error"`
}

// Error 오류
type Error struct {
	ErrorCode int    `json:"errorcode"`
	Message   string `json:"message"`
}
