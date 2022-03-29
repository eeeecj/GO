/*
 * @Author: eeeecj
 * @Date: 2022-03-25 14:52:26
 * @LastEditors: eeeecj
 * @LastEditTime: 2022-03-28 16:31:43
 * @Description:
 */

package core

import (
	"fmt"
	"os"
	"project/vue_admin/gin-vue-admin-my/global"
	"project/vue_admin/gin-vue-admin-my/utils"
	"time"

	"go.uber.org/zap/zapcore"

	"go.uber.org/zap"
)

func Zap() *zap.Logger {
	if ok, _ := utils.PathExists(global.GAB_CONFIG.Zap.Directory); !ok {
		fmt.Printf("Create file %v", global.GAB_CONFIG.Zap.Directory)
		_ = os.Mkdir(global.GAB_CONFIG.Zap.Directory, os.ModePerm)
	}

	debuglevel := zap.LevelEnablerFunc(func(lv zapcore.Level) bool {
		return lv == zap.DebugLevel
	})

	inforlevel := zap.LevelEnablerFunc(func(lv zapcore.Level) bool {
		return lv == zap.InfoLevel
	})

	warnlevel := zap.LevelEnablerFunc(func(lv zapcore.Level) bool {
		return lv == zap.WarnLevel
	})

	errorlevel := zap.LevelEnablerFunc(func(lv zapcore.Level) bool {
		return lv >= zap.ErrorLevel
	})

	levels := [...]zapcore.Core{
		getCore(fmt.Sprintf("./%s/server_debug.log", global.GAB_CONFIG.Zap.Directory), debuglevel),
		getCore(fmt.Sprintf("./%s/server_info.log", global.GAB_CONFIG.Zap.Directory), inforlevel),
		getCore(fmt.Sprintf("./%s/server_warn.log", global.GAB_CONFIG.Zap.Directory), warnlevel),
		getCore(fmt.Sprintf("./%s/server_error.log", global.GAB_CONFIG.Zap.Directory), errorlevel),
	}
	logger := zap.New(zapcore.NewTee(levels[:]...), zap.AddCaller())
	if global.GAB_CONFIG.Zap.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}

/**
 * @Description:获取日志处理单元
 * @param {string} filename
 * @param {zapcore.LevelEnabler} lv
 * @return {*}
 */
func getCore(filename string, lv zapcore.LevelEnabler) zapcore.Core {
	writer := utils.GetWriterSyncer(filename)
	return zapcore.NewCore(getEncoder(), writer, lv)
}

/**
 * @Description: 返回编码器
 * @param {*}
 * @return {*}
 */
func getEncoder() zapcore.Encoder {
	if global.GAB_CONFIG.Zap.Format == "json" {
		return zapcore.NewJSONEncoder(getEncoderConfig())
	}
	return zapcore.NewConsoleEncoder(getEncoderConfig())
}

/**
 * @Description: 自定义编码器配置
 * @param {*}
 * @return {*}
 */
func getEncoderConfig() zapcore.EncoderConfig {
	config := zapcore.EncoderConfig{
		MessageKey:    "message",
		LevelKey:      "level",
		TimeKey:       "time",
		NameKey:       "name",
		CallerKey:     "caller",
		StacktraceKey: global.GAB_CONFIG.Zap.StacktraceKey,
		LineEnding:    zapcore.DefaultLineEnding,

		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     CustomTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}
	switch global.GAB_CONFIG.Zap.Format {
	case "LowercaseLevelEncoder":
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	case "LowercaseColorLevelEncoder":
		config.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	case "CapitalLevelEncoder":
		config.EncodeLevel = zapcore.CapitalLevelEncoder
	case "CapitalColorLevelEncoder":
		config.EncodeLevel = zapcore.CapitalColorLevelEncoder
	default:
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	}
	return config
}

/**
 * @Description: 时间处理函数
 * @param {time.Time} t
 * @param {zapcore.PrimitiveArrayEncoder} p
 * @return {*}
 */
func CustomTimeEncoder(t time.Time, p zapcore.PrimitiveArrayEncoder) {
	p.AppendString(t.Format(global.GAB_CONFIG.Zap.Prefix + "2006/01/02 - 15:04:05.000"))
}
