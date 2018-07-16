package main

import (
	"encoding/binary"
	"math"
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
	EngageBreaks      TestingCommandEnum = 0
	DisengageBreaks   TestingCommandEnum = 1
	OpenThruster      TestingCommandEnum = 2
	CloseThruster     TestingCommandEnum = 3
	OpenBallValves    TestingCommandEnum = 4
	CloseBallValves   TestingCommandEnum = 5
	OpenReleaseValve  TestingCommandEnum = 6
	CloseReleaseValve TestingCommandEnum = 7
)

type PhotonCommand interface {
	WriteCommand() []byte
}

type MissionParamsCommand struct {
	Distance float32
	Pressure float32
	TopSpeed float32
}

func (c *MissionParamsCommand) WriteCommand() []byte {
	contents := make([]byte, 16)
	binary.LittleEndian.PutUint32(contents[0:4], SetVariablesCommandNum)
	binary.LittleEndian.PutUint32(contents[4:8], math.Float32bits(c.Distance))
	binary.LittleEndian.PutUint32(contents[8:12], math.Float32bits(c.Pressure))
	binary.LittleEndian.PutUint32(contents[12:16], math.Float32bits(c.TopSpeed))
	return contents
}

type ArmCommand struct {
}

func (*ArmCommand) WriteCommand() []byte {
	contents := make([]byte, 4)
	binary.LittleEndian.PutUint32(contents[0:4], GoToIdleCommandNum)
	return contents
}

type StartCommand struct {
}

func (*StartCommand) WriteCommand() []byte {
	contents := make([]byte, 4)
	binary.LittleEndian.PutUint32(contents[0:4], GoToAcceleratingCommandNum)
	return contents
}

type TestingCommand struct {
	TestCommand TestingCommandEnum
}

func (t *TestingCommand) WriteCommand() []byte {
	contents := make([]byte, 8)
	binary.LittleEndian.PutUint32(contents[0:], TestingCommandNum)
	binary.LittleEndian.PutUint32(contents[4:], uint32(t.TestCommand))
	return contents
}

type StopTestingCommand struct {
}

func (t *StopTestingCommand) WriteCommand() []byte {
	contents := make([]byte, 4)
	binary.LittleEndian.PutUint32(contents[0:4], GoToStandbyCommandNum)
	return contents
}

type AbortCommand struct {
}

func (*AbortCommand) WriteCommand() []byte {
	return make([]byte, 0)
}
