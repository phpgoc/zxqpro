# 介绍
- 想要做一个极简的项目管理工具，适用于5-30人的小公司，学习成本极低
- 开发中
- 从没写过前端代码，前端基本都是在面向AI编程

# 目录结构
- font 前端使用react， ant design
- server 后端使用gin， go-swagger 默认  sqlite go-cache

# 开发依赖

## go npm node

## go 格式化工具
```shell
go install mvdan.cc/gofumpt
```

##  gin swagger 生成文档
```shell
go install github.com/swaggo/swag/cmd/swag@latest
```

## gin swagger 生成文档命令
```shell
swag init -g cmd/server/main.go -d  server -o server/docs 
```

## 提交自动格式化
```shell
cp scripts/pre-commit .git/hooks/
chmod +x .git/hooks/pre-commit
```

# 运行
- server
```shell
cd server
go  mod tidy
go run cmd/server/main.go
```
- front
```shell
cd font
npm install
npm run start
```