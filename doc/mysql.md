


# 使用 go-zero 过程中，遇到的 mysql 疑难杂症


schema.sql 

建表语句，表名为 `users`，则 entity 为 model.Users，命名差点意思。

建表语句，字段 id 设置 AUTO_INCREMENT；那么，创建记录，则没有 id 字段；如果使用 uuid，慎重考虑。



