package services

import (
	"mqtt_agv_simulator/models"

	paho "github.com/eclipse/paho.mqtt.golang"
)

// 수신된 메시지를 처리하기 위한 채널 (controller가 소유)
var (
	OrderChan         = make(chan models.OrderMessage)
	InstantActionChan = make(chan models.InstantActionsMessage)
)

// AgvLogicController는 주문과 즉시액션을 받아 처리 흐름을 제어합니다.
func AgvLogicController(client paho.Client) {
	for {
		select {
		case newOrder := <-OrderChan:
			HandleNewOrder(client, &newOrder)
		case instantActions := <-InstantActionChan:
			HandleInstantActions(client, &instantActions)
		}
	}
}
