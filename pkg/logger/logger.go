package logger

import (
	"github.com/jshzhj/ginframe/pkg/setting"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"time"
)

var logger *zap.Logger

func Setup() {
	var (
		core   zapcore.Core
		caller zap.Option
	)
	hook := lumberjack.Logger{
		Filename:   setting.LogSetting.LogSavePath, // 日志文件路径
		MaxSize:    setting.LogSetting.MaxSize,     // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: setting.LogSetting.MaxBackups,  // 日志文件最多保存多少个备份
		// (如果MaxBackups和maxage都为0,则不删除任何旧的日志文件)
		MaxAge:   setting.LogSetting.MaxAge,   // 文件最多保存多少天
		Compress: setting.LogSetting.Compress, // 是否压缩
	}

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 小写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,     // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder, //
		EncodeCaller:   zapcore.ShortCallerEncoder,     // 短路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}

	// 设置日志级别
	atomicLevel := zap.NewAtomicLevel()
	if setting.LogSetting.LogLevel == 0 {
		//开发环境
		atomicLevel.SetLevel(zap.DebugLevel)
	} else {
		//生产环境
		atomicLevel.SetLevel(zap.InfoLevel)
	}
	//开启开发模式，堆栈跟踪
	caller = zap.AddCaller()
	if setting.LogSetting.IsConsole == 0 {
		core = zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderConfig),               // 编码器配置
			zapcore.NewMultiWriteSyncer(zapcore.AddSync(&hook)), // 打印到控制台和文件
			atomicLevel,                                         // 日志级别
		)
	} else {
		core = zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderConfig),                                           // 编码器配置
			zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)), // 打印到控制台和文件
			atomicLevel,                                                                     // 日志级别
		)
	}

	// 开启文件及行号
	development := zap.Development()

	// 设置初始化字段
	filed := zap.Fields()

	// 构造日志
	logger = zap.New(core, caller, development, filed)
	//logger.Info("12", zap.Int("hha", 1))
	//Info("123", zap.Int("hha", 1))
}

//http请求日志,需要记录msg,接口调用耗费时间,请求的url,http状态码,用户的id,链路trace_id
func Info(msg, url string, status int, spent_time time.Duration, user_id, trace_id int) {
	logger.Info(msg,
		zap.String("url", url),
		zap.Int("http_status_code", status),
		zap.Duration("spend_time", spent_time),
		zap.Int("user_id", user_id),
		zap.Int("trace_id", trace_id),
	)
}
func Warn(msg, url string, status int, spent_time time.Duration, user_id, trace_id int) {
	logger.Warn(msg,
		zap.String("url", url),
		zap.Int("http_status_code", status),
		zap.Duration("spend_time", spent_time),
		zap.Int("user_id", user_id),
		zap.Int("trace_id", trace_id),
	)
}
func Fatal(msg, url string, status int, spent_time time.Duration, user_id, trace_id int) {
	logger.Fatal(msg,
		zap.String("url", url),
		zap.Int("http_status_code", status),
		zap.Duration("spend_time", spent_time),
		zap.Int("user_id", user_id),
		zap.Int("trace_id", trace_id),
	)
}
func Panic(msg, url string, status int, spent_time time.Duration, user_id, trace_id int) {
	logger.Panic(msg,
		zap.String("url", url),
		zap.Int("http_status_code", status),
		zap.Duration("spend_time", spent_time),
		zap.Int("user_id", user_id),
		zap.Int("trace_id", trace_id),
	)
}
func Debug(msg, url string, status int, spent_time time.Duration, user_id, trace_id int) {
	logger.Panic(msg,
		zap.String("url", url),
		zap.Int("http_status_code", status),
		zap.Duration("spend_time", spent_time),
		zap.Int("user_id", user_id),
		zap.Int("trace_id", trace_id),
	)
}
func Error(msg, url string, status int, spent_time time.Duration, user_id, trace_id int) {
	logger.Panic(msg,
		zap.String("url", url),
		zap.Int("http_status_code", status),
		zap.Duration("spend_time", spent_time),
		zap.Int("user_id", user_id),
		zap.Int("trace_id", trace_id),
	)
}
