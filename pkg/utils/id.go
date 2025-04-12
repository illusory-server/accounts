package utils

import "github.com/google/uuid"

var (
	_ IDGenerator = (*UUIDGenerate)(nil)
	_ IDGenerator = (*IDGenerateTest)(nil)
)

type IDGenerator interface {
	GenerateID() string
}

type UUIDGenerate struct{}

func NewUUIDGenerator() UUIDGenerate {
	return UUIDGenerate{}
}

func (u UUIDGenerate) GenerateID() string {
	return uuid.New().String()
}

type IDGenerateTest struct {
	ID string
}

func NewIDGeneratorTest(id string) IDGenerateTest {
	return IDGenerateTest{
		ID: id,
	}
}

func (i IDGenerateTest) GenerateID() string {
	return i.ID
}
