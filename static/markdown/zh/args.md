[返回](../../../README.md)
# 参数说明
- help 或者 -h 显示帮助信息
- version 显示版本信息
- -d 或者 --debug 开启gin调试模式,默认是生产模式
- -g 或者 --git_pull_interval 设置git拉去的间隔，单位是分钟，默认是10分钟，是cron的*/n,里的n值，31-59都是每小时执行一次，在31分钟执行-59分钟执行
- -p 或者 --port 设置服务端口，默认是8080
- -r 或者 --redis_addr 设置redis地址，默认是使用go-cache，开发时建议使用redis，比如本机有一个redis，就写localhost:6379,如果使用go-cache，每次重启server服务都会清空数据，登录的cookie也会清空
- --gin_log 设置gin日志文件，默认是stdout
- --gorm_log 设置gorm日志文件，默认是stdout
- --self_log 设置自定义日志文件，默认是stdout
- -l 或者 --gorm_log_level 设置gorm日志级别，默认是允许使用i,w,e,s对应info，warn,error,slient， 默认是warn

## 参数建议
- 默认环境就是生成环境合适的配置 如果需要仅仅需要选择日记的文件
- 开发环境建议 -d -r localhost:6379 -l i 