package models

// ConnectionMessage는 연결 상태를 알리기 위한 메시지입니다.
type ConnectionMessage struct {
	HeaderID        uint64 `json:"headerId"`
	Timestamp       string `json:"timestamp"`
	Version         string `json:"version"`
	Manufacturer    string `json:"manufacturer"`
	SerialNumber    string `json:"serialNumber"`
	ConnectionState string `json:"connectionState"`
}
