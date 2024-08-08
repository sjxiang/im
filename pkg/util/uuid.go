package util

import (
	"time"
	
	"github.com/bwmarrin/snowflake"
)

type idGenerator struct {
	node *snowflake.Node
}

func NewIDGenerator(now time.Time, machineId int64) (*idGenerator, error) {

	snowflake.Epoch = now.Unix()
	node, err := snowflake.NewNode(machineId)
	if err != nil {
		return nil, err
	}

	return &idGenerator{node: node}, nil
}

func (i *idGenerator) GenerateUUID() int64 {
	return i.node.Generate().Int64()
}
