

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
    "mobile":   "12345678912",
    "nickname": "sjxiang",
    "password": "123456qwe",
    "sex":      1,
    "avatar": "sjxiang.jpeg"
}'


curl --location --request POST 'localhost:8888/v1/user/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "mobile": "18851762282",
    "password": "123456"
}'


curl --location --request GET 'localhost:8888/v1/user/detail' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzE4MzAzOTAsImlhdCI6MTcyMzE5MDM5MCwiaWRlbnRpdHkiOiJiYzUxZmY3My1kZmMwLTQ0Y2QtODM5OC01NDgzYmU1YjAyZjUifQ.PjLMi8xLej7E49URX8FHIdq3J7Y56xjP1YIi_zQpt04'

```