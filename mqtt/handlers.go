package mqtt

import (
	"encoding/json"
	"fmt"
	"log"
	"mqtt_agv_simulator/models"
	"mqtt_agv_simulator/services"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// 핸들러 함수들
var orderMessageHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	log.Printf("Order 토픽 수신: %s\n", msg.Topic())
	var order models.OrderMessage
	if err := json.Unmarshal(msg.Payload(), &order); err != nil {
		log.Printf("오류: Order 메시지 파싱 실패 - %v", err)
		return
	}
	services.OrderChan <- order // services 패키지의 채널 사용
}

var instantActionsHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	log.Printf("InstantActions 토픽 수신: %s\n", msg.Topic())
	var instantActions models.InstantActionsMessage
	if err := json.Unmarshal(msg.Payload(), &instantActions); err != nil {
		log.Printf("오류: InstantActions 메시지 파싱 실패 - %v", err)
		return
	}
	services.InstantActionChan <- instantActions // services 패키지의 채널 사용
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	log.Println("✅ MQTT 브로커에 연결되었습니다.")
	// Order 토픽 구독
	orderTopic := fmt.Sprintf("meili/v2/Roboligent/%s/order", "+")
	token := client.Subscribe(orderTopic, 1, orderMessageHandler)
	token.Wait()
	log.Printf("구독 시작: %s\n", orderTopic)

	// InstantActions 토픽 구독
	actionsTopic := fmt.Sprintf("meili/v2/Roboligent/%s/instantActions", "+")
	token = client.Subscribe(actionsTopic, 1, instantActionsHandler)
	token.Wait()
	log.Printf("구독 시작: %s\n", actionsTopic)

	// 연결 성공 시 "ONLINE" 상태 발행 (Retain 플래그 true)
	services.PublishConnectionState(client, "ONLINE", true)
}

var connectionLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	log.Printf("🔌 MQTT 연결이 끊겼습니다: %v", err)
}
