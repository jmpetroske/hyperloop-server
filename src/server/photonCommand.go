package main

type TestingCommandEnum int

const (
	EngageBreaks        TestingCommandEnum = iota
	DisengageBreaks     TestingCommandEnum = iota
	EngageSolenoids     TestingCommandEnum = iota
	DisengageSolenoids  TestingCommandEnum = iota
	EngageBallValves    TestingCommandEnum = iota
	DisengageBallValves TestingCommandEnum = iota
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
