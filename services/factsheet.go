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

var factsheetData models.Factsheet

// InitFactsheet는 프로그램 시작 시 AGV의 고정 사양 정보를 설정합니다.
func InitFactsheet() {
	factsheetData = models.Factsheet{
		Version:      "2.0.0",
		Manufacturer: "Roboligent",
		SerialNumber: state.AgvSerialNumber,
		TypeSpecification: models.TypeSpecification{
			SeriesName:        "DEX-Series",
			SeriesDescription: "Differential Drive AGV for general purpose",
			AGVKinematic:      "DIFF",
			AGVClass:          "CARRIER",
			MaxLoadMass:       models.Float64(200.0),
			LocalizationTypes: []string{"NATURAL", "GRID"},
			NavigationTypes:   []string{"AUTONOMOUS"},
		},
		PhysicalParameters: models.PhysicalParameters{
			SpeedMin:        models.Float64(0.0),
			SpeedMax:        models.Float64(2.0),
			AccelerationMax: models.Float64(1.0),
			DecelerationMax: models.Float64(1.5),
			HeightMax:       models.Float64(0.5),
			Width:           models.Float64(0.6),
			Length:          models.Float64(0.8),
		},
		ProtocolLimits: models.ProtocolLimits{
			MaxStringLens: models.MaxStringLens{MsgLen: 65535},
			MaxArrayLens:  models.MaxArrayLens{OrderNodes: 100, OrderEdges: 100},
			Timing: models.Timing{
				MinOrderInterval: models.Float64(0.2),
				MinStateInterval: models.Float64(0.2),
			},
		},
		ProtocolFeatures: models.ProtocolFeatures{
			OptionalParameters: []models.OptionalParameter{
				{Parameter: "state.agvPosition", Support: "SUPPORTED"},
			},
			AGVActions: []models.AGVAction{
				{ActionType: "cancelOrder", ActionScopes: []string{"INSTANT"}},
				{ActionType: "initPosition", ActionScopes: []string{"INSTANT"}},
				{ActionType: "factsheetRequest", ActionScopes: []string{"INSTANT"}},
			},
		},
		AGVGeometry: models.AGVGeometry{
			Envelopes2D: []models.Envelope2D{
				{
					Set: "default",
					PolygonPoints: []models.PolygonPoint{
						{X: models.Float64(0.4), Y: models.Float64(0.3)},
						{X: models.Float64(-0.4), Y: models.Float64(0.3)},
						{X: models.Float64(-0.4), Y: models.Float64(-0.3)},
						{X: models.Float64(0.4), Y: models.Float64(-0.3)},
					},
				},
			},
		},
		LoadSpecification: models.LoadSpecification{
			LoadPositions: []string{"center"},
		},
	}
}

// PublishFactsheet는 사전에 정의된 AGV 사양 정보를 발행합니다.
func PublishFactsheet(client paho.Client) {
	factsheetToSend := factsheetData

	atomic.AddUint64(&state.FactsheetHeaderID, 1)
	factsheetToSend.HeaderID = state.FactsheetHeaderID
	factsheetToSend.Timestamp = time.Now().UTC().Format(time.RFC3339)

	payload, err := json.Marshal(factsheetToSend)
	if err != nil {
		log.Printf("오류: Factsheet 메시지 마샬링 실패 - %v", err)
		return
	}

	topic := fmt.Sprintf("meili/v2/%s/%s/factsheet", factsheetToSend.Manufacturer, factsheetToSend.SerialNumber)
	token := client.Publish(topic, 1, false, payload)
	token.Wait()
	log.Printf("Factsheet 발행 완료: Topic=%s", topic)
}
