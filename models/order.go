package models

type NodePosition struct {
	X                     float64 `json:"x"`
	Y                     float64 `json:"y"`
	Theta                 float64 `json:"theta,omitempty"`
	AllowedDeviationXY    float64 `json:"allowedDeviationXY,omitempty"`
	AllowedDeviationTheta float64 `json:"allowedDeviationTheta,omitempty"`
	MapID                 string  `json:"mapId"`
	MapDescription        string  `json:"mapDescription,omitempty"`
}

type Node struct {
	NodeID          string       `json:"nodeId"`
	SequenceID      int          `json:"sequenceId"`
	NodeDescription string       `json:"nodeDescription,omitempty"`
	Released        bool         `json:"released"`
	NodePosition    NodePosition `json:"nodePosition,omitempty"`
	Actions         []Action     `json:"actions"`
}

type ControlPoint struct {
	X      float64 `json:"x"`
	Y      float64 `json:"y"`
	Weight float64 `json:"weight,omitempty"`
}

type Trajectory struct {
	Degree        int            `json:"degree"`
	KnotVector    []float64      `json:"knotVector"`
	ControlPoints []ControlPoint `json:"controlPoints"`
}

type Corridor struct {
	LeftWidth        float64 `json:"leftWidth"`
	RightWidth       float64 `json:"rightWidth"`
	CorridorRefPoint string  `json:"corridorRefPoint,omitempty"`
}

type Edge struct {
	EdgeID           string      `json:"edgeId"`
	SequenceID       int         `json:"sequenceId"`
	EdgeDescription  string      `json:"edgeDescription,omitempty"`
	Released         bool        `json:"released"`
	StartNodeID      string      `json:"startNodeId"`
	EndNodeID        string      `json:"endNodeId"`
	MaxSpeed         float64     `json:"maxSpeed,omitempty"`
	MaxHeight        float64     `json:"maxHeight,omitempty"`
	MinHeight        float64     `json:"minHeight,omitempty"`
	Orientation      float64     `json:"orientation,omitempty"`
	OrientationType  string      `json:"orientationType,omitempty"`
	Direction        string      `json:"direction,omitempty"`
	RotationAllowed  *bool       `json:"rotationAllowed,omitempty"`
	MaxRotationSpeed float64     `json:"maxRotationSpeed,omitempty"`
	Length           float64     `json:"length,omitempty"`
	Trajectory       *Trajectory `json:"trajectory,omitempty"`
	Corridor         *Corridor   `json:"corridor,omitempty"`
	Actions          []Action    `json:"actions"`
}

type OrderMessage struct {
	HeaderID      int    `json:"headerId"`
	Timestamp     string `json:"timestamp"`
	Version       string `json:"version"`
	Manufacturer  string `json:"manufacturer"`
	SerialNumber  string `json:"serialNumber"`
	OrderID       string `json:"orderId"`
	OrderUpdateID int    `json:"orderUpdateId"`
	ZoneSetID     string `json:"zoneSetId,omitempty"`
	Nodes         []Node `json:"nodes"`
	Edges         []Edge `json:"edges"`
}
