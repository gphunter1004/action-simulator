package services

import (
	"log"
	"mqtt_agv_simulator/models"
	"mqtt_agv_simulator/state"

	paho "github.com/eclipse/paho.mqtt.golang"
)

func HandleInstantActions(client paho.Client, msg *models.InstantActionsMessage) {
	for _, action := range msg.Actions {
		switch action.ActionType {
		case "cancelOrder":
			log.Println("CancelOrder 액션 수신. 현재 주문을 취소합니다.")
			state.AgvState.Lock()
			if state.AgvState.CancelOrderCycle != nil {
				state.AgvState.CancelOrderCycle()
			}
			state.AgvState.CurrentOrder = nil
			state.AgvState.ActionStatus = "FAILED"
			state.AgvState.CancelOrderCycle = nil
			state.AgvState.LastNodeId = ""
			state.AgvState.LastNodeSequenceId = 0
			state.AgvState.Unlock()

		case "initPosition":
			log.Println("InitPosition 액션 수신. AGV 위치를 초기화합니다.")
			for _, param := range action.ActionParameters {
				if param.Key == "pose" {
					if pose, ok := param.Value.(map[string]interface{}); ok {
						state.AgvState.Lock()
						state.AgvState.Position.X, _ = pose["x"].(float64)
						state.AgvState.Position.Y, _ = pose["y"].(float64)
						state.AgvState.Position.Theta, _ = pose["theta"].(float64)
						state.AgvState.Position.MapID, _ = pose["mapId"].(string)
						state.AgvState.Position.PositionInitialized = true
						state.AgvState.Position.LocalizationScore = 1.0
						state.AgvState.Unlock()
						log.Printf("위치 초기화 완료: X=%.2f, Y=%.2f, Theta=%.2f",
							state.AgvState.Position.X, state.AgvState.Position.Y, state.AgvState.Position.Theta)
					}
				}
			}
		case "factsheetRequest":
			log.Println("FactsheetRequest 액션 수신. Factsheet을 발행합니다.")
			go PublishFactsheet(client)

		default:
			log.Printf("알 수 없는 InstantAction 수신: %s", action.ActionType)
		}
	}
}
