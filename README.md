# zhy-log

使日志打印更高效、支持不同级别写入不同文件、输出方式等功能

## Quick Start

使用时请查看`options.go`文件中五种选择器的使用方式，下面为基本使用方式

```go
logrus.InitLogrus(WithLevel("info"))
logger.Info("info")
```
当你输入级别选择器时，会默认打印终端，如果你想指定输出，下面为使用方式

```go
logrus.InitLogrus(WithLevel("info"), WithOutPut("all"))
logger.Info("info")
```
`WithOutPut`可填入`all` 、`console`、 `file`三种，分别为输出终端和记入文件、终端、文件，当然当你选择`all`、`file`
两种输出形式时，你需要填写输出文件的属性，包括路径、文件最多数、清理时间等，当然如果你不填写的话，会默认给你生成文件配置，下面是如何自定义文件配置

```go
logrus.InitLogrus(WithLevel("info"), WithOutPut("all"), WithFileConfig(LogWriter{
        Filename:   "./logs.txt",
        MaxSize:    10,    //最大M数，超过则切割
        MaxBackups: 5,     //最大文件保留数，超过就删除最老的日志文件
        MaxAge:     30,    //保存30天
        Compress:   false, //是否压缩
		Used:       true,  //自定义时需要填写true，不然会默认使用defaultWriter
	}))

```
有时我们需要直将`error`错误捕获并写入文件，以便可以查看错误，但不想让低于`error`
级别的日志写入，我们可以这样做

```go
logrus.InitLogrus(WithLevel("info"), WithFileConfig(LogWriter{
        FileName:   "./logs.txt",
        MaxSize:    10,
        MaxBackups: 5,
        MaxAge:     30,
        Compress:   false,
        Used:       true,
}), WithOutPut("all"), WithFileLevel("warn", LogWriter{
        FileName:   "./warn.txt",
        MaxSize:    10,
        MaxBackups: 5,
        MaxAge:     30,
        Compress:   false,
        Used:       true,
}))
logger.Info("info")
logger.Warn("test")
```
这样我们就既得到了一个拥有全部日志的`logs.txt`文件和一个只有`warn`级别的日志了，同时还会将日志输出在终端

### 日志中包含traceId的使用方法
```go
c := logger.NewContext(context.Background(), "1234567")

logger.WithContext(c).Debug("debug")

```

#### Author
hengyuan zhang
