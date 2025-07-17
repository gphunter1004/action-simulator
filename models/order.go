package models

import (
	"encoding/json"
	"strconv"
)

// Float64 is a custom type that ensures JSON marshaling with decimal point
type Float64 float64

func (f Float64) MarshalJSON() ([]byte, error) {
	return json.Marshal(strconv.FormatFloat(float64(f), 'f', 1, 64))
}

func (f *Float64) UnmarshalJSON(data []byte) error {
	var v float64
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	*f = Float64(v)
	return nil
}

type NodePosition struct {
	X                     Float64 `json:"x"`
	Y                     Float64 `json:"y"`
	Theta                 Float64 `json:"theta,omitempty"`
	AllowedDeviationXY    Float64 `json:"allowedDeviationXY,omitempty"`
	AllowedDeviationTheta Float64 `json:"allowedDeviationTheta,omitempty"`
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
	X      Float64 `json:"x"`
	Y      Float64 `json:"y"`
	Weight Float64 `json:"weight,omitempty"`
}

type Trajectory struct {
	Degree        int            `json:"degree"`
	KnotVector    []Float64      `json:"knotVector"`
	ControlPoints []ControlPoint `json:"controlPoints"`
}

type Corridor struct {
	LeftWidth        Float64 `json:"leftWidth"`
	RightWidth       Float64 `json:"rightWidth"`
	CorridorRefPoint string  `json:"corridorRefPoint,omitempty"`
}

type Edge struct {
	EdgeID           string      `json:"edgeId"`
	SequenceID       int         `json:"sequenceId"`
	EdgeDescription  string      `json:"edgeDescription,omitempty"`
	Released         bool        `json:"released"`
	StartNodeID      string      `json:"startNodeId"`
	EndNodeID        string      `json:"endNodeId"`
	MaxSpeed         Float64     `json:"maxSpeed,omitempty"`
	MaxHeight        Float64     `json:"maxHeight,omitempty"`
	MinHeight        Float64     `json:"minHeight,omitempty"`
	Orientation      Float64     `json:"orientation,omitempty"`
	OrientationType  string      `json:"orientationType,omitempty"`
	Direction        string      `json:"direction,omitempty"`
	RotationAllowed  *bool       `json:"rotationAllowed,omitempty"`
	MaxRotationSpeed Float64     `json:"maxRotationSpeed,omitempty"`
	Length           Float64     `json:"length,omitempty"`
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
