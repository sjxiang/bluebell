package logic

import (
	"github.com/sjxiang/bluebell/dao/mysql"
	"github.com/sjxiang/bluebell/pkg/snowflake"
)

// 存放业务逻辑的代码，拼装


func Signup() {
	// 判断用户存不存在
	mysql.QueryUserByUsername()

	// 生成 id
	snowflake.GetID()

	// 密码加密

	// 保存进数据库
	mysql.InsertUser()

}