package logger

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var Logger *zap.Logger
var SugarLogger *zap.SugaredLogger

func LoggerInit() {
	core := zapcore.NewCore(
		getEncoder(),
		zapcore.NewMultiWriteSyncer(
			getWriterSyncer(),
			zapcore.AddSync(os.Stdout)),
		getLevelEnable(),
	)

	// addcaller是为了显示调用者，该参数可省略掉
	Logger = zap.New(core, zap.AddCaller())
	SugarLogger = Logger.Sugar()
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()            // 获取配置实例
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder        // 规定时间格式
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder // 全部大写
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getWriterSyncer() zapcore.WriteSyncer {
	file, _ := os.Create(viper.GetString("logs-path"))
	return zapcore.AddSync(file)
}

func getLevelEnable() zapcore.Level {
	return zapcore.DebugLevel
}
