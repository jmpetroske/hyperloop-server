package main

import (
	"encoding/binary"
)

type TestingCommandEnum uint32

const (
	SetVariablesCommandNum     uint32 = 0
	GoToTestingCommandNum      uint32 = 1
	GoToStandbyCommandNum      uint32 = 2
	TestingCommandNum          uint32 = 3
	GoToIdleCommandNum         uint32 = 4
	GoToAcceleratingCommandNum uint32 = 5
)

const (
	EngageBreaks        TestingCommandEnum = 0
	DisengageBreaks     TestingCommandEnum = 1
	EngageSolenoids     TestingCommandEnum = 2
	DisengageSolenoids  TestingCommandEnum = 3
	EngageBallValves    TestingCommandEnum = 4
	DisengageBallValves TestingCommandEnum = 5
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

func (t *TestingCommand) WriteCommand() []byte {
	contents := make([]byte, 8)
	binary.LittleEndian.PutUint32(contents[0:4], TestingCommandNum)
	binary.LittleEndian.PutUint32(contents[5:8], uint32(t.TestCommand))
	return contents
}

type AbortCommand struct {
}

func (*AbortCommand) WriteCommand() []byte {
	return make([]byte, 0)
}
