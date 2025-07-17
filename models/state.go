package models

type AGVPosition struct {
	X                   Float64 `json:"x"`
	Y                   Float64 `json:"y"`
	Theta               Float64 `json:"theta"`
	MapID               string  `json:"mapId"`
	MapDescription      string  `json:"mapDescription,omitempty"`
	PositionInitialized bool    `json:"positionInitialized"`
	LocalizationScore   Float64 `json:"localizationScore"`
	DeviationRange      Float64 `json:"deviationRange,omitempty"`
}

type Velocity struct {
	Vx    Float64 `json:"vx"`
	Vy    Float64 `json:"vy"`
	Omega Float64 `json:"omega"`
}

type NodeState struct {
	NodeID          string        `json:"nodeId"`
	SequenceID      int           `json:"sequenceId"`
	NodeDescription string        `json:"nodeDescription,omitempty"`
	Released        bool          `json:"released"`
	NodePosition    *NodePosition `json:"nodePosition,omitempty"`
}

type EdgeState struct {
	EdgeID          string      `json:"edgeId"`
	SequenceID      int         `json:"sequenceId"`
	EdgeDescription string      `json:"edgeDescription,omitempty"`
	Released        bool        `json:"released"`
	Trajectory      *Trajectory `json:"trajectory,omitempty"`
}

type ActionState struct {
	ActionID          string `json:"actionId"`
	ActionType        string `json:"actionType,omitempty"`
	ActionDescription string `json:"actionDescription,omitempty"`
	ActionStatus      string `json:"actionStatus"`
	ResultDescription string `json:"resultDescription,omitempty"`
}

type BatteryState struct {
	BatteryCharge  Float64 `json:"batteryCharge"`
	BatteryVoltage Float64 `json:"batteryVoltage,omitempty"`
	BatteryHealth  Float64 `json:"batteryHealth,omitempty"`
	Charging       bool    `json:"charging"`
	Reach          Float64 `json:"reach,omitempty"`
}

type ErrorReference struct {
	ReferenceKey   string `json:"referenceKey"`
	ReferenceValue string `json:"referenceValue"`
}

type Error struct {
	ErrorType        string           `json:"errorType"`
	ErrorReferences  []ErrorReference `json:"errorReferences,omitempty"`
	ErrorDescription string           `json:"errorDescription,omitempty"`
	ErrorHint        string           `json:"errorHint,omitempty"`
	ErrorLevel       string           `json:"errorLevel"`
}

type SafetyState struct {
	EStop          string `json:"eStop"`
	FieldViolation bool   `json:"fieldViolation"`
}

type StateMessage struct {
	HeaderID              uint64        `json:"headerId"`
	Timestamp             string        `json:"timestamp"`
	Version               string        `json:"version"`
	Manufacturer          string        `json:"manufacturer"`
	SerialNumber          string        `json:"serialNumber"`
	OrderID               string        `json:"orderId"`
	OrderUpdateID         int           `json:"orderUpdateId"`
	LastNodeID            string        `json:"lastNodeId"`
	LastNodeSequenceID    int           `json:"lastNodeSequenceId"`
	NodeStates            []NodeState   `json:"nodeStates"`
	EdgeStates            []EdgeState   `json:"edgeStates"`
	AGVPosition           *AGVPosition  `json:"agvPosition"`
	Velocity              *Velocity     `json:"velocity"`
	Driving               bool          `json:"driving"`
	Paused                bool          `json:"paused"`
	NewBaseRequest        bool          `json:"newBaseRequest"`
	DistanceSinceLastNode Float64       `json:"distanceSinceLastNode"`
	OperatingMode         string        `json:"operatingMode"`
	ActionStates          []ActionState `json:"actionStates"`
	BatteryState          BatteryState  `json:"batteryState"`
	Errors                []Error       `json:"errors"`
	Information           []interface{} `json:"information"`
	SafetyState           SafetyState   `json:"safetyState"`
}
