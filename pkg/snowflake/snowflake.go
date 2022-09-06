package snowflake

import (
	"fmt"
	"time"

	"github.com/sony/sonyflake"
)


var (
	sonyFlake *sonyflake.Sonyflake  // 实例
	sonyMachineID uint16  // 机器 id
)


func getMachineID() (uint16, error) {
	return sonyMachineID, nil
}


// 需传入当前的机器 id
func Init(machineId uint16) (err error) {
	
	sonyMachineID = machineId

	st, err := time.Parse("2006-01-02", "2020-01-01")
	if err != nil {
		return err
	}

	settings := sonyflake.Settings{
		StartTime: st,
		MachineID: getMachineID,
	}
	
	sonyFlake = sonyflake.NewSonyflake(settings)
	return
}


// 返回生成 id
func GetID() (id uint64, err error) {
	if sonyFlake == nil {
		err = fmt.Errorf("sony flake not inited")
		return
	}

	id, err = sonyFlake.NextID()
	return
}





