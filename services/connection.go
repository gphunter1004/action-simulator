package services

import (
	"encoding/json"
	"fmt"
	"log"
	"mqtt_agv_simulator/models"
	"mqtt_agv_simulator/state"
	"sync/atomic"
	"time"

	paho "github.com/eclipse/paho.mqtt.golang"
)

// PublishConnectionState는 ONLINE, OFFLINE 상태를 발행합니다.
func PublishConnectionState(client paho.Client, connState string, retained bool) {
	atomic.AddUint64(&state.ConnectionHeaderID, 1)
	message := models.ConnectionMessage{
		HeaderID:        state.ConnectionHeaderID,
		Timestamp:       time.Now().UTC().Format(time.RFC3339),
		Version:         "2.0",
		Manufacturer:    "Roboligent",
		SerialNumber:    state.AgvSerialNumber,
		ConnectionState: connState,
	}
	payload, err := json.Marshal(message)
	if err != nil {
		log.Printf("오류: Connection 메시지 마샬링 실패 - %v", err)
		return
	}
	topic := fmt.Sprintf("meili/v2/Roboligent/%s/connection", state.AgvSerialNumber)
	token := client.Publish(topic, 1, retained, payload)
	token.Wait()
	log.Printf("Connection 상태 발행: %s", connState)
}

// PublishRejectedState는 거절된 주문에 대한 상태를 한 번 발행합니다.
func PublishRejectedState(client paho.Client, rejectedOrder *models.OrderMessage) {
	atomic.AddUint64(&state.HeaderID, 1)

	state.AgvState.Lock()
	busyOrderID := ""
	if state.AgvState.CurrentOrder != nil {
		busyOrderID = state.AgvState.CurrentOrder.OrderID
	}
	agvPosition := state.AgvState.Position
	state.AgvState.Unlock()

	rejectionError := models.Error{
		ErrorType:        "OrderRejected",
		ErrorDescription: fmt.Sprintf("AGV is busy with a previous order (ID: %s)", busyOrderID),
		ErrorLevel:       "WARNING",
		ErrorReferences:  []models.ErrorReference{{ReferenceKey: "orderId", ReferenceValue: rejectedOrder.OrderID}},
	}
	actionStates := []models.ActionState{}
	for _, node := range rejectedOrder.Nodes {
		for _, action := range node.Actions {
			actionStates = append(actionStates, models.ActionState{ActionID: action.ActionID, ActionType: action.ActionType, ActionStatus: "REJECTED"})
		}
	}
	for _, edge := range rejectedOrder.Edges {
		for _, action := range edge.Actions {
			actionStates = append(actionStates, models.ActionState{ActionID: action.ActionID, ActionType: action.ActionType, ActionStatus: "REJECTED"})
		}
	}

	stateMsg := models.StateMessage{
		HeaderID:           state.HeaderID,
		Timestamp:          time.Now().UTC().Format(time.RFC3339),
		Version:            "2.0.0",
		Manufacturer:       "Roboligent",
		SerialNumber:       state.AgvSerialNumber,
		OrderID:            rejectedOrder.OrderID,
		OrderUpdateID:      rejectedOrder.OrderUpdateID,
		AGVPosition:        agvPosition,
		ActionStates:       actionStates,
		Errors:             []models.Error{rejectionError},
		LastNodeID:         "",
		LastNodeSequenceID: 0,
		NodeStates:         []models.NodeState{},
		EdgeStates:         []models.EdgeState{},
		Velocity:           &models.Velocity{Vx: 0.0, Vy: 0.0, Omega: 0.0},
		Driving:            false,
		Paused:             false,
		OperatingMode:      "AUTOMATIC",
		BatteryState:       models.BatteryState{BatteryCharge: 60.0, Charging: false},
		Information:        []interface{}{},
		SafetyState:        models.SafetyState{EStop: "NONE", FieldViolation: false},
	}

	payload, err := json.Marshal(stateMsg)
	if err != nil {
		log.Printf("오류: 거절 메시지 마샬링 실패 - %v", err)
		return
	}

	topic := fmt.Sprintf("meili/v2/Roboligent/%s/state", state.AgvSerialNumber)
	_ = client.Publish(topic, 1, false, payload)
	log.Printf("거절 상태 발행 완료: OrderID=%s", rejectedOrder.OrderID)
}
