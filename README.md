```
bluebell 风铃草



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
    


返回 response 处理

    1. 不到最顶层 controller 处理，划分更细致（烦）
    2. err 分类，漏掉信息；error.Is 判断分类
    
    
        errors.new("用户不存在")  
        "record not found"
        errors.new("密码错误") 
        "crypto/bcrypt: hashedPassword is not the hash of the given password"



GORM

    Find vs. First，结果出入有点大

        SELECT * FROM `user` WHERE username = 'sjxia1g' ORDER BY `user`.`id` LIMIT 1  // 查得出
        SELECT * FROM `user` WHERE username = 'xjxiang'  // 查不出



两者 dao 层方法

差异一
    func Create(username, password string) (user *models.User, err error) 
    func Create(*user *models.User) error


差异二
    func CheckUserExist() error 
    func IsUserExist() bool 
```