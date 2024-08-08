

# im

>

用户注册
用户登录
用户详情
用户搜索

curl --location --request POST 'localhost:8888/v1/user/register' \
--header 'Content-Type: application/json' \
--data-raw '{
    "mobile":   "12345678910",
    "nickname": "sjxiang",
    "password": "123456qwe",
    "sex":      1,
    "avatar": "sjxiang.jpeg"
}'


curl --location --request POST 'localhost:8888/v1/user/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "mobile": "12344449999",
    "password": "123456"
}'



curl --location --request GET 'localhost:8888/v1/user/detail' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzE3ODQ5OTQsImlhdCI6MTcyMzE0NDk5NCwiaWRlbnRpdHkiOiI2YzczNDFlMS1lMTRlLTRjYWYtODlmNC1hMWNhNGEyZjlkZGMifQ.bN4EKRXroodqWhSzhcsHaVr1K11xWogTWTbS7VEbQEk'
