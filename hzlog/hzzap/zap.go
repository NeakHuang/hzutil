// Date: 2023/5/25
// Author:
// Description：

package hzzap

import (
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/zero-contrib/logx/zapx"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

func NewDevelopmentLogger(outputPath string, opts ...zap.Option) (*zap.Logger, error) {
	// 创建Lumberjack实例来配置日志文件切割
	lumberjackLogger := &lumberjack.Logger{
		Filename:   outputPath,
		MaxSize:    10,   // 单个日志文件的最大大小（以MB为单位）
		MaxAge:     1,    // 单个日志文件的天数
		MaxBackups: 10,   // 最多保留的旧日志文件数量
		LocalTime:  true, // 使用本地时间
	}
	// 创建WriteSyncer，将日志输出到lumberjackLogger和控制台
	writeSyncer := zapcore.NewMultiWriteSyncer(zapcore.AddSync(lumberjackLogger), zapcore.AddSync(os.Stdout))
	// 配置zap日志核心
	encoderConfig := zap.NewDevelopmentEncoderConfig()
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig),
		writeSyncer,
		zap.NewAtomicLevelAt(zapcore.DebugLevel),
	)
	// 创建Logger实例
	logger := zap.New(core)
	return logger, nil
}

func NewProductionLogger(outputPath string, opts ...zap.Option) (*zap.Logger, error) {
	// 创建Lumberjack实例来配置日志文件切割
	lumberjackLogger := &lumberjack.Logger{
		Filename:   outputPath,
		MaxSize:    10,   // 单个日志文件的最大大小（以MB为单位）
		MaxAge:     1,    // 单个日志文件的天数
		MaxBackups: 10,   // 最多保留的旧日志文件数量
		LocalTime:  true, // 使用本地时间
	}
	// 创建WriteSyncer，将日志输出到lumberjackLogger
	writeSyncer := zapcore.AddSync(lumberjackLogger)
	// 配置zap日志核心
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder // 使用ISO 8601格式的时间编码器
	core := zapcore.NewCore(
		// zapcore.NewJSONEncoder(encoderConfig), // 不再用json格式写入了
		zapcore.NewConsoleEncoder(encoderConfig),
		writeSyncer,
		zap.NewAtomicLevelAt(zapcore.InfoLevel),
	)
	// 创建Logger实例
	logger := zap.New(core)
	return logger, nil
}

func NewZap() {
	writer, err := zapx.NewZapWriter()
	logx.Must(err)
	logx.SetWriter(writer)
}
