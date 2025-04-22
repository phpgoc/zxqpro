[返回](../../../README.md)


## 启动server
```shell
cd ../../../server
go run cmd/server/main.go 
```


## 启动font
```shell
cd ../../../font
bun run dev
```

## swagger生成文档(修改接口请求和返回值后需要)
```shell
cd ../../..
swag init -g cmd/server/main.go -d  server -o server/docs 
```

## tauri 运行桌面应用
```shell
cd ../../../font
npx tauri dev
```
