<p>本项目是用go语言写的把mysql数据迁移到分库分表。</p>

## 特点

1. 使用go协程、连接池特性
1. 协程执行错误时记录日志
1. 程序执行时进度显示
1. 不同机器可以通过配置分批查询参数和批量插入参数达到最佳性能

## 使用方式：

```go
//创建数据表
go run main.go -mode=tables -ac=create

//删除数据表
go run main.go -mode=tables -ac=drop

//迁移数据
go run main.go -mode=migrate
```
