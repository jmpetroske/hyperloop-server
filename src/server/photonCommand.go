package main

import (
	"encoding/binary"
	"math"
)

type TestingCommandEnum uint32

const (
	SetVariablesCommandNum     uint32 = 0
	GoToIdleCommandNum         uint32 = 1
	GoToArmedCommandNum        uint32 = 2
	GoToAcceleratingCommandNum uint32 = 3
	TestingCommandNum          uint32 = 4
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
	ExpectedNextMode() uint32 // For testing fake arduino mode
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

func (*MissionParamsCommand) ExpectedNextMode() uint32 {
	return 1
}

type GoToIdleCommand struct {
}

func (*GoToIdleCommand) WriteCommand() []byte {
	contents := make([]byte, 4)
	binary.LittleEndian.PutUint32(contents[0:4], GoToIdleCommandNum)
	return contents
}

func (*GoToIdleCommand) ExpectedNextMode() uint32 {
	return 0
}

type ArmCommand struct {
}

func (*ArmCommand) WriteCommand() []byte {
	contents := make([]byte, 4)
	binary.LittleEndian.PutUint32(contents[0:4], GoToArmedCommandNum)
	return contents
}

func (*ArmCommand) ExpectedNextMode() uint32 {
	return 2
}

type StartCommand struct {
}

func (*StartCommand) WriteCommand() []byte {
	contents := make([]byte, 4)
	binary.LittleEndian.PutUint32(contents[0:4], GoToAcceleratingCommandNum)
	return contents
}

func (*StartCommand) ExpectedNextMode() uint32 {
	return 3
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

func (*TestingCommand) ExpectedNextMode() uint32 {
	return 0
}

type AbortCommand struct {
}

func (*AbortCommand) WriteCommand() []byte {
	return make([]byte, 1)
}

func (*AbortCommand) ExpectedNextMode() uint32 {
	return 0
}
