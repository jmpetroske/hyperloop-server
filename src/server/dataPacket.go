package main

type DataPacket struct {
	Timestamp                    uint32  `json:"timestamp"`
	Mode                         uint32  `json:"mode"`
	PressureThrusterTank         bool    `json:"pressureThrusterTank"`
	PressureBrakingTank          bool    `json:"pressureBrakingTank"`
	PressureBrakingExtend        bool    `json:"pressureBrakingExtend"`
	PressureBrakingCompress      bool    `json:"pressureBrakingCompress"`
	TemperatureThrusterLeftTank  float32 `json:"temperatureThrusterLeftTank"`
	TemperatureThrusterRightTank float32 `json:"temperatureThrusterRightTank"`
	TemperatureBraking           float32 `json:"temperatureBraking"`
	TemperatureBattery           float32 `json:"temperatureBattery"`
	Speed                        float32 `json:"speed"`
	Distance                     float32 `json:"distance"`
	Vibration1                   float32 `json:"vibration1"`
	Vibration2                   float32 `json:"vibration2"`
	Vibration3                   float32 `json:"vibration3"`
	Vibration4                   float32 `json:"vibration4"`
	ActuationThruster            bool    `json:"actuationThruster"`
	ActuationRelieveValve        bool    `json:"actuationRelieveValve"`
	ActuationSafetyValve         bool    `json:"actuationSafetyValve"`
	ActuationBrakingValve        bool    `json:"actuationBrakingValve"`
	CurrentThrusterLeftValve     float32 `json:"currentThrusterLeftValve"`
	CurrentThrusterRightValve    float32 `json:"currentThrusterRightValve"`
	CurrentBrakingValve          float32 `json:"currentBrakingValve"`
	CurrentBattery               float32 `json:"currentBattery"`
	VoltageThrusterLeftValve     float32 `json:"voltageThrusterLeftValve"`
	VoltageThrusterRightValve    float32 `json:"voltageThrusterRightValve"`
	VoltageBrakingValve          float32 `json:"voltageBrakingValve"`
	VoltageBattery               float32 `json:"voltageBattery"`
	Orientationx                 float32 `json:"orientationx"`
	Orientationy                 float32 `json:"orientationy"`
	Orientationz                 float32 `json:"orientationz"`
	Accelerationx                float32 `json:"accelerationx"`
	Accelerationy                float32 `json:"accelerationy"`
	Accelerationz                float32 `json:"accelerationz"`
	ImuStatus                    uint32  `json:"imuStatus"`
	BrakeStatus                  uint32  `json:"brakeStatus"`
	PropulsionStatus             uint32  `json:"propulsionStatus"`
	EmergencyBrake               bool    `json:"emergencyBrake"`
}
