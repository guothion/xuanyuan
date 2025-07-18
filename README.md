# xuanyuan
the service of xuanyuan create by gin

# 框架搭建

## 配置解决方案
### 插件
- [github.com/spf13/viper](https://github.com/spf13/viper)
### 为什么使用 viper?
- 支持 JSON/TOML/YAML/HCL/envfile/Java properties 等多种格式的配置文件；
- 可以设置监听配置文件的修改，修改时自动加载新的配置；
- 从环境变量、命令行选项和io.Reader中读取配置；
- 从远程配置系统中读取和监听修改，如 etcd/Consul；
- 代码逻辑中显示设置键值。

## 日志解决方案
### 插件
- [gopkg.in/natefinch/lumberjack.v2](https://github.com/natefinch/lumberjack)
- [github.com/sirupsen/logrus](https://github.com/sirupsen/logrus)

## 数据库解决方案
这里我们使用 [gorm](https://gorm.io/)，这个算是Gin框架官方推荐的 ORM 库了。
### mysql 实现方案
这里我们使用的是 mysql 的库 `gorm.io/driver/mysql`，官方支持。
命令：
```bash
go get gorm.io/driver/mysql
go get gorm.io/gorm
```

### mongo 实现方案
