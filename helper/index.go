package helper

import (
	"fmt"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// 專案(log)輸出

//-------------------------------------------------------------------------------------------------[Variable]

const (
	timeFormat = "2006-01-02T15:04:05"
	skip       = 1
)

//-------------------------------------------------------------------------------------------------[Func]

// timeEncoder 排除(+8)時區
func timeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("01-02 15:04:05.000"))
}

// NewConsoleLogger ...
func NewConsoleLogger() (*zap.Logger, error) {
	c := zap.NewDevelopmentConfig()
	c.EncoderConfig.EncodeTime = timeEncoder
	c.EncoderConfig.LevelKey = ""

	l, e := c.Build(zap.AddCallerSkip(skip))
	//l, e := c.Build()
	if e != nil {
		return nil, e
	}
	return l, nil
}

/*
NewFileLogger ...
	@see https://github.com/uber-go/zap/blob/master/FAQ.md#does-zap-support-log-rotation
*/
func NewFileLogger(path, name string) *zap.Logger {
	h := lumberjack.Logger{
		Filename:   fmt.Sprintf("./%s/%s/.log", path, name), // 文件輸出位置
		MaxSize:    10,                                      // 文件大小 MB
		LocalTime:  true,                                    // 是否使用本地時間
		Compress:   false,                                   // 是否壓縮檔案 ( 大小差滿多的，但不確定效能會不會影響很大 )
		MaxAge:     30,                                      // 預設值是不刪除舊檔(單位天), 修改為 30 天
		MaxBackups: 50,                                      // 保留多少個備份檔( 受限 MaxAge )，預設全保留 ( 500MB )
	}
	ws := zapcore.AddSync(&h)

	ec := zap.NewProductionEncoderConfig()
	// ec.TimeKey = "" // 空字串 不顯示欄位
	// ec.MessageKey = ""
	ec.LevelKey = ""
	ec.EncodeTime = timeEncoder

	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(ec),
		// zapcore.NewJSONEncoder(ec),
		ws,                // ...
		zapcore.InfoLevel, //
	)

	l := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(skip))
	// l := zap.New(core, zap.AddCaller())

	return l
}

// NewEmptyLogger ...
func NewEmptyLogger() *zap.Logger {
	return zap.NewNop()
}

//-------------------------------------------------------------------------------------------------

// Logger ...
type Logger struct {
	File    *zap.Logger
	Console *zap.Logger
}

//-------------------------------------------------------------------------------------------------

/*
NewLogger
	path 路徑
	name (資料夾/檔案)名稱
	useConsole 是否啟用 console 功能
*/
func NewLogger(path, name string, useConsole bool) *Logger {
	log := &Logger{
		File: NewFileLogger(path, name),
	}

	if useConsole {
		var e error
		log.Console, e = NewConsoleLogger()
		if e != nil {
			panic(e)
		}
	} else {
		log.Console = zap.NewNop()
	}

	log.File.Info("start", zap.String("time", time.Now().Format(timeFormat)))
	return log
}
