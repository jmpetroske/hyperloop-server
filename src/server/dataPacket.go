package main

type DataPacket struct {
	Timestamp          uint32  `json:"timestamp"`
	Mode               uint32  `json:"mode"`
	PressureTank       float32 `json:"pressureTank"`
	PressureBraking    float32 `json:"pressureBraking"`
	PressureActuate    float32 `json:"pressureActuate"`
	PressureRetract    float32 `json:"pressureRetract"`
	PressureRelief     float32 `json:"pressureRelief"`
	TemperatureTank    float32 `json:"temperatureTank"`
	TemperaturePiping  float32 `json:"temperaturePiping"`
	TemperatureBraking float32 `json:"temperatureBraking"`
	TemperatureBattery float32 `json:"temperatureBattery"`
	Speed              float32 `json:"speed"`
	Distance           float32 `json:"distance"`
	Vibration1         float32 `json:"vibration1"`
	Vibration2         float32 `json:"vibration2"`
	CurrentBallValve   float32 `json:"currentBallValve"`
	CurrentRelief      float32 `json:"currentRelief"`
	CurrentBraking     float32 `json:"currentBraking"`
	CurrentBattery     float32 `json:"currentBattery"`
	VoltageBallValves  float32 `json:"voltageBallValves"`
	VoltageRelief      float32 `json:"voltageRelief"`
	VoltageBraking     float32 `json:"voltageBraking"`
	VoltageBattery     float32 `json:"voltageBattery"`
	Voltage5vRef       float32 `json:"voltage5vRef"`
	OrientationX       float32 `json:"orientationX"`
	OrientationY       float32 `json:"orientationY"`
	OrientationZ       float32 `json:"orientationZ"`
	AccelerationX      float32 `json:"accelerationX"`
	AccelerationY      float32 `json:"accelerationY"`
	AccelerationZ      float32 `json:"accelerationZ"`
	ImuStatus          uint32  `json:"imuStatus"`
	BrakeStatus        uint32  `json:"brakeStatus"`
	PropulsionStatus   uint32  `json:"propulsionStatus"`
	ActuationThruster  bool    `json:"actuationThruster"`
	ActuationRelief    bool    `json:"actuationRelief"`
	ActuationSafety    bool    `json:"actuationSafety"`
	ActuationBraking   bool    `json:"actuationBraking"`
	EmergencyBrake     bool    `json:"emergencyBrake"`
}
