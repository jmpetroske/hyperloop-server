package main

type TestingCommandEnum int

const (
	EngageBreaks         TestingCommandEnum = 1 << iota
	DisengageBreaks      TestingCommandEnum = 1 << iota
	EngageSolenoids      TestingCommandEnum = 1 << iota
	DisengageSolenoids   TestingCommandEnum = 1 << iota
	EngageBallValves     TestingCommandEnum = 1 << iota
	DisengageBallValves TestingCommandEnum = 1 << iota
)

type PhotonCommand interface {
	WriteCommand() []byte
}

type MissionParamsCommand struct {
	Distance float32
	Pressure float32
	TopSpeed float32
}

func (*MissionParamsCommand) WriteCommand() []byte {
	return make([]byte, 0)
}

type ArmCommand struct {
}

func (*ArmCommand) WriteCommand() []byte {
	return make([]byte, 0)
}

type StartCommand struct {
}

func (*StartCommand) WriteCommand() []byte {
	return make([]byte, 0)
}

type TestingCommand struct {
	TestCommand TestingCommandEnum
}

func (*TestingCommand) WriteCommand() []byte {
	return make([]byte, 0)
}

type AbortCommand struct {
}

func (*AbortCommand) WriteCommand() []byte {
	return make([]byte, 0)
}
