package services

import (
	"context"
	"log"
	"mqtt_agv_simulator/models"
	"mqtt_agv_simulator/state"
	"time"

	paho "github.com/eclipse/paho.mqtt.golang"
)

const (
	waitCount         = 3
	initializingCount = 5
	runningCount      = 10
)

func HandleNewOrder(client paho.Client, newOrder *models.OrderMessage) {
	state.AgvState.Lock()
	isBusy := state.AgvState.ActionStatus != "FINISHED" && state.AgvState.ActionStatus != "FAILED"
	currentStatus := state.AgvState.ActionStatus
	state.AgvState.Unlock()

	if isBusy {
		log.Printf("AGV가 바쁨 (상태: %s). 새로운 Order 거절: OrderID=%s", currentStatus, newOrder.OrderID)
		go PublishRejectedState(client, newOrder)
	} else {
		log.Printf("새로운 Order 수신 및 처리 시작: OrderID = %s", newOrder.OrderID)
		state.AgvState.Lock()
		if state.AgvState.CancelOrderCycle != nil {
			state.AgvState.CancelOrderCycle()
		}
		ctx, cancel := context.WithCancel(context.Background())
		state.AgvState.CurrentOrder = newOrder
		state.AgvState.ActionStatus = "WAITING"
		state.AgvState.CancelOrderCycle = cancel
		if len(newOrder.Nodes) > 0 {
			state.AgvState.LastNodeId = newOrder.Nodes[0].NodeID
			state.AgvState.LastNodeSequenceId = newOrder.Nodes[0].SequenceID
		}
		state.AgvState.Unlock()

		go runOrderCycle(ctx)
	}
}

// runOrderCycle은 주문 상태를 시간에 따라 업데이트합니다.
func runOrderCycle(ctx context.Context) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	counter := 0
	for {
		select {
		case <-ticker.C:
			state.AgvState.Lock()
			currentStatus := state.AgvState.ActionStatus

			switch currentStatus {
			case "WAITING":
				if counter >= waitCount-1 {
					state.AgvState.ActionStatus = "INITIALIZING"
					counter = 0
				} else {
					counter++
				}
			case "INITIALIZING":
				if counter >= initializingCount-1 {
					state.AgvState.ActionStatus = "RUNNING"
					counter = 0
				} else {
					counter++
				}
			case "RUNNING":
				if counter >= runningCount-1 {
					state.AgvState.ActionStatus = "FINISHED"
					counter = 0
				} else {
					counter++
				}
			}
			state.AgvState.Unlock()
		case <-ctx.Done():
			log.Printf("Order 주기 중단됨 (Context Done)")
			return
		}
	}
}
