package initial

import (
	"path/filepath"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)
var Logger *zap.Logger
func InitialLogger() {
	encoder_config := zap.NewProductionEncoderConfig()
	// 设置日志记录中时间的格式
	encoder_config.EncodeTime = zapcore.ISO8601TimeEncoder
	// 日志Encoder 还是JSONEncoder，把日志行格式化成JSON格式的
	encoder := zapcore.NewJSONEncoder(encoder_config)
	// 设置日志路径
	fileWriteSyncer := logWritter()
	// 最终实现写日志
	core := zapcore.NewTee(
	// 同时向控制台和文件写日志， 生产环境记得把控制台写入去掉，日志记录的基本是Debug 及以上，生产环境记得改成Info
		// zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zapcore.DebugLevel),
		zapcore.NewCore(encoder, fileWriteSyncer, zapcore.DebugLevel),
		
	)
	Logger = zap.New(core,zap.AddCaller(),zap.AddCallerSkip(1))
   }
func logWritter() zapcore.WriteSyncer{	
	log_path := filepath.Join(GetValue("base_dir").(string),"proxy.log")
	lumberWriteSyncer := &lumberjack.Logger{
	Filename:   log_path,
	MaxSize:    10, // megabytes
	MaxBackups: 100,
	MaxAge:     28,    // days
	Compress:   false, //Compress确定是否应该使用gzip压缩已旋转的日志文件。默认值是不执行压缩。
}
return zapcore.AddSync(lumberWriteSyncer)
}