package helper

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// 專案(log)輸出

//-------------------------------------------------------------------------------------------------[Variable]

const timeFormat = "2006-01-02T15:04:05"

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

	l, e := c.Build(zap.AddCallerSkip(1)) // fix caller skip
	if e != nil {
		return nil, e
	}
	return l, nil
}

// NewFileLogger ...
// @see https://github.com/uber-go/zap/blob/master/FAQ.md#does-zap-support-log-rotation
func NewFileLogger(path string) *zap.Logger {
	h := lumberjack.Logger{
		Filename:   fmt.Sprintf("./dist/logs/%s/.log", path), // 文件輸出位置
		MaxSize:    10,                                       // 文件大小 MB
		LocalTime:  true,                                     // 是否使用本地時間
		Compress:   false,                                    // 是否壓縮檔案 ( 大小差滿多的，但不確定效能會不會影響很大 )
		MaxAge:     30,                                       // 預設值是不刪除舊檔(單位天), 修改為 30 天
		MaxBackups: 50,                                       // 保留多少個備份檔( 受限 MaxAge )，預設全保留 ( 500MB )
	}
	ws := zapcore.AddSync(&h)

	ec := zap.NewProductionEncoderConfig()
	// _encoderConfig.TimeKey = "" // 空字串 不顯示欄位
	// _encoderConfig.MessageKey = ""
	ec.LevelKey = ""
	ec.EncodeTime = timeEncoder

	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(ec), // zapcore.NewJSONEncoder(encoderConfig),
		ws,                            // ...
		zapcore.InfoLevel,             //
	)

	l := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))

	return l
}

// NewEmptyLogger ...
func NewEmptyLogger() *zap.Logger {
	return zap.NewNop()
}

//-------------------------------------------------------------------------------------------------[Color]

// ConsoleColor 背景色
//@see https://stackoverflow.com/questions/5762491/how-to-print-color-in-console-using-system-out-println
type ConsoleColor string

// ToString ...
func (v ConsoleColor) ToString() string {
	return string(v)
}

// ANSIBlackBG ...
const ANSIBlackBG ConsoleColor = "\u001B[40m%s\u001B[0m"

// ANSIRedBG ...
const ANSIRedBG ConsoleColor = "\u001B[41m%s\u001B[0m"

// ANSIGreenBG ...
const ANSIGreenBG ConsoleColor = "\u001B[42m%s\u001B[0m"

// ANSIYellowBG ...
const ANSIYellowBG ConsoleColor = "\u001B[43m%s\u001B[0m"

// ANSIBlueBG ...
const ANSIBlueBG ConsoleColor = "\u001B[44m%s\u001B[0m"

// ANSIPurpleBG ...
const ANSIPurpleBG ConsoleColor = "\u001B[45m%s\u001B[0m"

// ANSICyanBG ...
const ANSICyanBG ConsoleColor = "\u001B[46m%s\u001B[0m"

// ANSIWhiteBG ...
const ANSIWhiteBG ConsoleColor = "\u001B[47m%s\u001B[0m"

// ANSIDefaultBG ...
const ANSIDefaultBG ConsoleColor = "%s"

//-------------------------------------------------------------------------------------------------[Custom]

// KeyValuePair ...
type KeyValuePair map[string]interface{}

// Add ...
func (v KeyValuePair) Add(k string, p interface{}) KeyValuePair {
	v[k] = p
	return v
}

// Logger ...
type Logger struct {
	ConsoleColor
	mFile    *zap.Logger
	mConsole *zap.Logger
}

// Debug ...
// output value as json
// or use KeyValuePair
//
// var pair = KeyValuePair{}
// pair.Add("a", "b").Add("c", "d")
func (v Logger) Debug(output interface{}) {
	if v.mConsole == nil {
		return
	}

	j, e := json.Marshal(output)

	if e != nil {
		fmt.Println(e)
		return
	}

	c := fmt.Sprintf("\n"+v.ConsoleColor.ToString(), string(j))
	v.mConsole.Debug(c)
}

// Info @see Debug
func (v Logger) Info(output interface{}) {
	j, e := json.Marshal(output)

	if e != nil {
		fmt.Println(e)
		return
	}

	c := string(j)
	v.mFile.Info(c)

	if v.mConsole == nil {
		return
	}

	c = fmt.Sprintf("\n"+v.ConsoleColor.ToString(), c)
	v.mConsole.Debug(c)

}

// Fatal @see Debug
func (v Logger) Fatal(output interface{}) {
	v.Info(output)
	os.Exit(0)
}

// NewLogger ...
func NewLogger(fullFilePath string, useConsole bool, cc ConsoleColor) *Logger {
	l := &Logger{
		ConsoleColor: cc,
	}
	l.mFile = NewFileLogger(fullFilePath)

	if useConsole {
		var e error
		l.mConsole, e = NewConsoleLogger()
		if e != nil {
			panic(e)
		}
	}
	l.Info(KeyValuePair{"start": time.Now().Format(timeFormat)})

	return l
}
