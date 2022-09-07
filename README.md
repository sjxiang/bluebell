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
    


```