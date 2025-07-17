package models

type TypeSpecification struct {
	SeriesName        string   `json:"seriesName"`
	SeriesDescription string   `json:"seriesDescription,omitempty"`
	AGVKinematic      string   `json:"agvKinematic"`
	AGVClass          string   `json:"agvClass"`
	MaxLoadMass       Float64  `json:"maxLoadMass"`
	LocalizationTypes []string `json:"localizationTypes"`
	NavigationTypes   []string `json:"navigationTypes"`
}

type PhysicalParameters struct {
	SpeedMin        Float64 `json:"speedMin"`
	SpeedMax        Float64 `json:"speedMax"`
	AccelerationMax Float64 `json:"accelerationMax"`
	DecelerationMax Float64 `json:"decelerationMax"`
	HeightMin       Float64 `json:"heightMin,omitempty"`
	HeightMax       Float64 `json:"heightMax"`
	Width           Float64 `json:"width"`
	Length          Float64 `json:"length"`
}

type MaxStringLens struct {
	MsgLen          int  `json:"msgLen,omitempty"`
	TopicSerialLen  int  `json:"topicSerialLen,omitempty"`
	TopicElemLen    int  `json:"topicElemLen,omitempty"`
	IDLen           int  `json:"idLen,omitempty"`
	IDNumericalOnly bool `json:"idNumericalOnly,omitempty"`
	EnumLen         int  `json:"enumLen,omitempty"`
	LoadIDLen       int  `json:"loadIdLen,omitempty"`
}

type MaxArrayLens struct {
	OrderNodes              int `json:"order.nodes,omitempty"`
	OrderEdges              int `json:"order.edges,omitempty"`
	NodeActions             int `json:"node.actions,omitempty"`
	EdgeActions             int `json:"edge.actions,omitempty"`
	ActionsParameters       int `json:"actions.actionsParameters,omitempty"`
	InstantActions          int `json:"instantActions,omitempty"`
	TrajectoryKnotVector    int `json:"trajectory.knotVector,omitempty"`
	TrajectoryControlPoints int `json:"trajectory.controlPoints,omitempty"`
	StateNodeStates         int `json:"state.nodeStates,omitempty"`
	StateEdgeStates         int `json:"state.edgeStates,omitempty"`
	StateLoads              int `json:"state.loads,omitempty"`
	StateActionStates       int `json:"state.actionStates,omitempty"`
	StateErrors             int `json:"state.errors,omitempty"`
	StateInformation        int `json:"state.information,omitempty"`
	ErrorErrorReferences    int `json:"error.errorReferences,omitempty"`
	InfoInfoReferences      int `json:"information.infoReferences,omitempty"`
}

type Timing struct {
	MinOrderInterval      Float64 `json:"minOrderInterval"`
	MinStateInterval      Float64 `json:"minStateInterval"`
	DefaultStateInterval  Float64 `json:"defaultStateInterval,omitempty"`
	VisualizationInterval Float64 `json:"visualizationInterval,omitempty"`
}

type ProtocolLimits struct {
	MaxStringLens MaxStringLens `json:"maxStringLens"`
	MaxArrayLens  MaxArrayLens  `json:"maxArrayLens"`
	Timing        Timing        `json:"timing"`
}

type OptionalParameter struct {
	Parameter   string `json:"parameter"`
	Support     string `json:"support"`
	Description string `json:"description,omitempty"`
}

type ActionParameterFactsheet struct {
	Key           string `json:"key"`
	ValueDataType string `json:"valueDataType"`
	Description   string `json:"description,omitempty"`
	IsOptional    bool   `json:"isOptional,omitempty"`
}

type AGVAction struct {
	ActionType        string                     `json:"actionType"`
	ActionDescription string                     `json:"actionDescription,omitempty"`
	ActionScopes      []string                   `json:"actionScopes"`
	ActionParameters  []ActionParameterFactsheet `json:"actionParameters,omitempty"`
	ResultDescription string                     `json:"resultDescription,omitempty"`
	BlockingTypes     []string                   `json:"blockingTypes,omitempty"`
}

type ProtocolFeatures struct {
	OptionalParameters []OptionalParameter `json:"optionalParameters"`
	AGVActions         []AGVAction         `json:"agvActions"`
}

type PolygonPoint struct {
	X Float64 `json:"x"`
	Y Float64 `json:"y"`
}

type Envelope2D struct {
	Set           string         `json:"set"`
	PolygonPoints []PolygonPoint `json:"polygonPoints"`
	Description   string         `json:"description,omitempty"`
}

type AGVGeometry struct {
	Envelopes2D []Envelope2D `json:"envelopes2d,omitempty"`
}

type LoadSpecification struct {
	LoadPositions []string `json:"loadPositions,omitempty"`
}

type Factsheet struct {
	HeaderID           uint64             `json:"headerId"`
	Timestamp          string             `json:"timestamp"`
	Version            string             `json:"version"`
	Manufacturer       string             `json:"manufacturer"`
	SerialNumber       string             `json:"serialNumber"`
	TypeSpecification  TypeSpecification  `json:"typeSpecification"`
	PhysicalParameters PhysicalParameters `json:"physicalParameters"`
	ProtocolLimits     ProtocolLimits     `json:"protocolLimits"`
	ProtocolFeatures   ProtocolFeatures   `json:"protocolFeatures"`
	AGVGeometry        AGVGeometry        `json:"agvGeometry"`
	LoadSpecification  LoadSpecification  `json:"loadSpecification"`
}
