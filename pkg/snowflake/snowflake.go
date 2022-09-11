package snowflake

import (
	"time"

	"github.com/bwmarrin/snowflake"
	"github.com/sjxiang/bluebell/settings"
)


var node *snowflake.Node


// func Init(startTime string, machineID int64) (err error) {}

func Init(cfg *settings.AppConfig) (err error) {
	
	var st time.Time
	st, err = time.Parse("2006-01-02", cfg.StartTime)
	if err != nil {
		return
	}

	snowflake.Epoch = st.UnixNano() / 1000000
	node, err = snowflake.NewNode(cfg.MachineID)
	
	return
}


func GetID() int64 {
	return node.Generate().Int64()  // .String() 可以换
}