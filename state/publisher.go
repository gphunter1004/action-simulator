package state

import (
	"encoding/json"
	"fmt"
	"log"
	"mqtt_agv_simulator/models"
	"sync/atomic"
	"time"

	paho "github.com/eclipse/paho.mqtt.golang"
)

// PublishingLoop는 1초마다 현재 AGV 상태를 MQTT로 발행합니다.
func PublishingLoop(client paho.Client) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		PublishCurrentState(client)
	}
}

// PublishCurrentState는 agvState의 현재 정보를 바탕으로 State 메시지를 만듭니다.
func PublishCurrentState(client paho.Client) {
	AgvState.mu.Lock()
	defer AgvState.mu.Unlock()

	atomic.AddUint64(&HeaderID, 1)

	orderID := ""
	orderUpdateID := 0
	nodeStates := []models.NodeState{}
	edgeStates := []models.EdgeState{}
	actionStates := []models.ActionState{}

	if AgvState.CurrentOrder != nil {
		orderID = AgvState.CurrentOrder.OrderID
		orderUpdateID = AgvState.CurrentOrder.OrderUpdateID

		for _, node := range AgvState.CurrentOrder.Nodes {
			nodeStates = append(nodeStates, models.NodeState{NodeID: node.NodeID, SequenceID: node.SequenceID, Released: node.Released})
			for _, action := range node.Actions {
				actionStates = append(actionStates, models.ActionState{ActionID: action.ActionID, ActionType: action.ActionType, ActionStatus: AgvState.ActionStatus})
			}
		}
		for _, edge := range AgvState.CurrentOrder.Edges {
			edgeStates = append(edgeStates, models.EdgeState{EdgeID: edge.EdgeID, SequenceID: edge.SequenceID, Released: edge.Released})
			for _, action := range edge.Actions {
				actionStates = append(actionStates, models.ActionState{ActionID: action.ActionID, ActionType: action.ActionType, ActionStatus: AgvState.ActionStatus})
			}
		}
	}

	stateMsg := models.StateMessage{
		HeaderID:           HeaderID,
		Timestamp:          time.Now().UTC().Format(time.RFC3339),
		Version:            "2.0.0",
		Manufacturer:       "Roboligent",
		SerialNumber:       AgvSerialNumber, // 같은 패키지 내의 상수 사용
		OrderID:            orderID,
		OrderUpdateID:      orderUpdateID,
		LastNodeID:         AgvState.LastNodeId,
		LastNodeSequenceID: AgvState.LastNodeSequenceId,
		NodeStates:         nodeStates,
		EdgeStates:         edgeStates,
		AGVPosition:        AgvState.Position,
		Velocity:           &models.Velocity{Vx: 0.0, Vy: 0.0, Omega: 0.0},
		Driving:            AgvState.ActionStatus == "RUNNING",
		Paused:             false,
		OperatingMode:      "AUTOMATIC",
		ActionStates:       actionStates,
		BatteryState:       models.BatteryState{BatteryCharge: 60.0, BatteryVoltage: 40.0, Charging: false},
		Errors:             []models.Error{},
		Information:        []interface{}{},
		SafetyState:        models.SafetyState{EStop: "NONE", FieldViolation: false},
	}

	payload, err := json.Marshal(stateMsg)
	if err != nil {
		log.Printf("오류: state 메시지 마샬링 실패 - %v", err)
		return
	}

	topic := fmt.Sprintf("meili/v2/Roboligent/%s/state", AgvSerialNumber) // 같은 패키지 내의 상수 사용
	_ = client.Publish(topic, 0, false, payload)
	log.Printf("State 발행: OrderID=%s, ActionStatus=%s, PosInit=%v", orderID, AgvState.ActionStatus, AgvState.Position.PositionInitialized)
}
