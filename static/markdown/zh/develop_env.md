## server开发环境
- 安装go
- 建议设置中国区源
```shell
go env -w GOPROXY=https://goproxy.cn,direct
```
- 下载go依赖
```shell
cd ../../../server
go mod tidy 
```
- 启动服务
```shell
cd ../../../server
go run cmd/server/main.go 
```

## font 开发环境
- 安装nodejs
- 安装bun(建议，用npm也行)
```shell
cd ../../../font 
npm install -g bun
bun install
bun run dev
```

## tauri 开发环境
- 安装rust
- 设置cargo源(建议),参照 [https://rsproxy.cn/](https://rsproxy.cn/) ，[清华](https://mirrors.tuna.tsinghua.edu.cn/help/crates.io-index/)
- 运行(第一次会非常慢，10-30分钟都正常)
```shell
cd ../../../font 
npx tauri dev
```
