package base_service

import (
	"os"

	"github.com/bwmarrin/snowflake"
)

var snowflakeNode *snowflake.Node

func getSnowflakeNode() *snowflake.Node {
	if snowflakeNode == nil {
		n, err := snowflake.NewNode(1)
		if err != nil {
			println(err)
			os.Exit(1)
		}
		snowflakeNode = n
	}
	return snowflakeNode
}

func GenID() int64 {
	return getSnowflakeNode().Generate().Int64()
}
