# casbin

[参考地址](https://github.com/zupzup/casbin-http-role-example/blob/main/README.md)

```
GET http://localhost:9000/api/v1/hello
很遗憾,权限验证没有通过

POST http://localhost:9000/api/v1/add
增加Policy
增加成功

GET http://localhost:9000/api/v1/hello
恭喜您,权限验证通过
Hello 接收到GET请求..
```

