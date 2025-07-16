package mqtt

import (
	"encoding/json"
	"fmt"
	"log"
	"mqtt_agv_simulator/config"
	"mqtt_agv_simulator/models"
	"mqtt_agv_simulator/state" // state 패키지 import 추가
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// NewClient는 MQTT 클라이언트를 생성하고 LWT를 포함한 옵션을 설정합니다.
func NewClient(cfg *config.Config) mqtt.Client {
	opts := mqtt.NewClientOptions()
	brokerURL := fmt.Sprintf("tcp://%s:%s", cfg.Broker, cfg.Port)
	opts.AddBroker(brokerURL)
	opts.SetClientID(cfg.ClientID)
	opts.SetUsername(cfg.Username)
	opts.SetPassword(cfg.Password)

	// Last Will (LWT) 설정
	connectionTopic := fmt.Sprintf("meili/v2/Roboligent/%s/connection", state.AgvSerialNumber) // state 패키지에서 상수 사용
	willMessage := models.ConnectionMessage{
		HeaderID:        0,
		Timestamp:       time.Now().UTC().Format(time.RFC3339),
		Version:         "2.0",
		Manufacturer:    "Roboligent",
		SerialNumber:    state.AgvSerialNumber, // state 패키지에서 상수 사용
		ConnectionState: "CONNECTIONBROKEN",
	}
	willPayload, err := json.Marshal(willMessage)
	if err != nil {
		log.Fatalf("LWT 메시지 생성 실패: %v", err)
	}
	opts.SetWill(connectionTopic, string(willPayload), 1, true) // QoS 1, Retain True
	log.Println("Last Will (CONNECTIONBROKEN) 메시지가 설정되었습니다.")

	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectionLostHandler

	client := mqtt.NewClient(opts)
	return client
}
