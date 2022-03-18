# goweb
 gin+docker简单教程

## 使用docker-compose.yml
1. 构建整个容器系统
```shell
docker-compose up -d
```
2. curl命令访问
```shell
curl  -X GET 127.0.0.1:8082/redis
```