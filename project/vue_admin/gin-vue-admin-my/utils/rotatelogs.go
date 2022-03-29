/*
 * @Author: eeeecj
 * @Date: 2022-03-28 14:56:54
 * @LastEditors: eeeecj
 * @LastEditTime: 2022-03-28 15:32:13
 * @Description:
 */

package utils

import (
	"os"
	"project/vue_admin/gin-vue-admin-my/global"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap/zapcore"
)

/**
 * @Description: 获取同步写入入口
 * @param {string} filename
 * @return {*}
 */
func GetWriterSyncer(filename string) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    10, // megabytes
		MaxBackups: 200,
		MaxAge:     30,   //days
		Compress:   true, // disabled by default
	}
	if global.GAB_CONFIG.Zap.LoginConsole {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(lumberJackLogger))
	}
	return zapcore.AddSync(lumberJackLogger)
}
