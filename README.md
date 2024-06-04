# Woox — A CLI tool for building go applications.


Woox是一个基于Golang的应用脚手架，它是由Golang生态中各种非常流行的库整合而成的，它们的组合可以帮助你快速构建一个高效、可靠的应用程序。


## 功能
- **Hertz**: https://github.com/cloudwego/hertz
- **Gorm**: https://github.com/go-gorm/gorm
- **Wire**: https://github.com/google/wire
- **Viper**: https://github.com/spf13/viper
- **Zap**: https://github.com/uber-go/zap
- **Golang-jwt**: https://github.com/golang-jwt/jwt
- **Go-redis**: https://github.com/go-redis/redis
- **Testify**: https://github.com/stretchr/testify
- **Sonyflake**: https://github.com/sony/sonyflake
- **Gocron**:  https://github.com/go-co-op/gocron
- **Go-sqlmock**:  https://github.com/DATA-DOG/go-sqlmock
- **Gomock**:  https://github.com/golang/mock
- **Swaggo**:  https://github.com/swaggo/swag
- **Pitaya**:  https://github.com/topfreegames/pitaya
- More...

### 安装

您可以通过以下命令安装Woox：

```bash
go install github.com/go-woox/woox@latest
```

国内用户可以使用`GOPROXY`加速`go install`

```
$ go env -w GO111MODULE=on
$ go env -w GOPROXY=https://goproxy.cn,direct
```

> tips: 如果`go install`成功，却提示找不到woox命令，这是因为环境变量没有配置，可以把 GOBIN 目录配置到环境变量中即可


### 创建新项目

您可以使用以下命令创建一个新的Golang项目：

```bash
// 推荐新用户选择Advanced Layout
woox new projectName
```

