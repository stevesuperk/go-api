package config

const (
	AppMode = "release" //debug or release
	AppPort = ":9999"
	AppName = "go-api"
	MakeMigration = true

	// 签名超时时间
	AppSignExpiry = "120"

	// RSA Private File
	AppRsaPrivateFile = "rsa/private.pem"

	// 超时时间
	AppReadTimeout  = 120
	AppWriteTimeout = 120

	// 日志文件
	AppAccessLogName = "logs/" + AppName + "-access.log"
	AppErrorLogName  = "logs/" + AppName + "-error.log"
	AppGrpcLogName   = "logs/" + AppName + "-grpc.log"

	// 系统告警邮箱信息
	SystemEmailUser = "stevesuperk@gmail.com"
	SystemEmailPass = "" //密码或授权码
	SystemEmailHost = "smtp.gmail.com"
	SystemEmailPort = 465

	// 告警接收人
	ErrorNotifyUser = "stevesuperk@gmail.com"

	// 告警开关 1=开通 -1=关闭
	ErrorNotifyOpen = 1

	// Jaeger 配置信息
	JaegerHostPort = "127.0.0.1:6831"

	// Jaeger 配置开关 1=开通 -1=关闭
	JaegerOpen = 1
)
