package state

import (
	"context"
	"mqtt_agv_simulator/models"
	"sync"
)

// AGV의 핵심 고유 정보
const AgvSerialNumber = "DEX0002"

// AGVState는 AGV의 모든 상태 정보를 관리하는 중앙 구조체입니다.
type AGVState struct {
	mu                 sync.Mutex
	CurrentOrder       *models.OrderMessage
	ActionStatus       string
	Position           *models.AGVPosition
	CancelOrderCycle   context.CancelFunc
	LastNodeId         string
	LastNodeSequenceId int
}

// Lock은 AGVState의 뮤텍스를 잠그는 공개 메서드입니다.
func (s *AGVState) Lock() {
	s.mu.Lock()
}

// Unlock은 AGVState의 뮤텍스를 해제하는 공개 메서드입니다.
func (s *AGVState) Unlock() {
	s.mu.Unlock()
}

// 전역 변수들
var (
	HeaderID           uint64
	ConnectionHeaderID uint64
	FactsheetHeaderID  uint64
	AgvState           = &AGVState{ // AGV 전역 상태
		ActionStatus: "FINISHED",
		Position: &models.AGVPosition{
			PositionInitialized: false,
			LocalizationScore:   0.0,
			X:                   0.0,
			Y:                   0.0,
			Theta:               0.0,
			MapID:               "",
		},
	}
)
