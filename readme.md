

# im

>

用户注册
用户登录
用户详情
用户搜索

```bash

curl --location --request POST 'localhost:8888/v1/user/register' \
--header 'Content-Type: application/json' \
--data-raw '{
    "phone":   "18851762282",
    "nickname": "sjxiang",
    "password": "123456",
    "sex":      1,
    "avatar": "sjxiang.jpeg"
}'


curl --location --request POST 'localhost:8888/v1/user/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "phone": "18851762282",
    "password": "123456"
}'


curl --location --request GET 'localhost:8888/v1/user/detail' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzE5NTIzNDQsImlhdCI6MTcyMzMxMjM0NCwiaWRlbnRpdHkiOiI0NDE5ZjY5MC0xOWQ5LTQ5NDktOGQ2MS05YWUwY2FiYTY0YzEifQ.Ya_KxCI5aZMHiDEkmONi_FEfc9-P72-13J5yqCNAya4'

```


