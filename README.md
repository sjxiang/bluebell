```
bluebell 风铃草 nga



controller 获取参数 & 参数校验
logic 业务拼装
dao 

用户表结构设计

    user_id 

    1. 自增 id

        敏感信息 用户数泄漏
        分库分表 

    2. uuid 
        
        排序、检索不方便


    3. snowflake 雪花算法

        64 bit 组成

        时间戳 
            指定 startTime， 可容纳 69 年
        
        工作机器 id
            10 bit，可设置 0 ~ 1024 个节点

        序列号
            12 bit，1 ms，产生 4096 个 id
        


ctx.ShouldBindJSON(obj)

    JSON 序列化出现错误，panic
    "json: unsupported type: func() string"
    
    Tips: 客户端最少发送 `{}`


错误处理    
    日志记录要详细
        zap.L().Error("现象_PKG_Function", zap.Error(err)))  // "record not found"

    给客户端返回（不透露具体信息）
        errors.new("用户不存在")  


GORM

    Find vs. First，结果出入有点大

        SELECT * FROM `user` WHERE username = 'sjxia1g' ORDER BY `user`.`id` LIMIT 1  // 查得出
        SELECT * FROM `user` WHERE username = 'xjxiang'  // 查不出



    不要考虑数据库操作失败（大量下面这种操作，真的下头）

        if result.Error != nil {

        }

        if user.ID > 0 {

        }



加密
    明文密码和哈希值，不能颠倒位置
    {"error": "crypto/bcrypt: hashedPassword is not the hash of the given password"}



govalidator
    对应 Param，只能是 string

    还要转 strconv.ParseInt(str, 10, 64)


```