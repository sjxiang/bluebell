package logger

import (
	"time"
	"strings"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)


func Init() (err error) {

	writeSyncer := getLogWriter(
		viper.GetString("log.filename"),
		viper.GetInt("log.max_size"),
		viper.GetInt("log.max_age"),
		viper.GetInt("log.max_backups"),
	)

	encoder := getEncoder()

	// 设置日志等级
	logLevel := new(zapcore.Level)
	if err := logLevel.UnmarshalText([]byte(viper.GetString("log.level"))); err != nil {
		return err
	}

	// 初始化 core
	// zapcore.DebugLevel
	core := zapcore.NewCore(encoder, writeSyncer, logLevel)
	
	// 初始化 Logger
	lg := zap.New(
		core, 
		zap.AddCaller(),                    // 调用文件和行号，内部使用 runtime.Caller
		zap.AddCallerSkip(1),               // 封装了一层，调用文件去除一层（runtime.Caller(1)） 
		zap.AddStacktrace(zap.ErrorLevel),  // Error 才会显示 stacktrace
	)

	// 替换 zap 库中的全局 logger
	zap.ReplaceGlobals(lg)

	return
}


// getLogWriter 多个输出流 
// 1. 日志文件 .log 
// 2. 控制台 os.Stdout  
func getLogWriter(filename string,  maxSize, maxAge, maxBackup int) zapcore.WriteSyncer {

	// // 按照日期记录 log
	datename := time.Now().Format("2006-01-02")
	filename = strings.ReplaceAll(filename, "test", datename)
	
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxBackups: maxBackup,
		MaxAge:     maxAge,
	}

	return zapcore.AddSync(lumberJackLogger)
}



// getEncoder 设置 Log Entry 格式
func getEncoder() zapcore.Encoder {
	
	encoderConfig := zap.NewProductionEncoderConfig()  // 很多都是默认设置，不用再写一遍
	encoderConfig.EncodeTime = customTimeEncoder  // 时间戳格式
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	
	return zapcore.NewJSONEncoder(encoderConfig)  // JSON 格式
}

// customTimeEncoder 自定义友好的时间戳格式
func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}


