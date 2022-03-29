/*
 * @Author: eeeecj
 * @Date: 2022-03-25 14:18:17
 * @LastEditors: eeeecj
 * @LastEditTime: 2022-03-28 16:38:34
 * @Description:
 */

package global

import (
	"project/vue_admin/gin-vue-admin-my/config"

	"go.uber.org/zap"

	"gorm.io/gorm"

	"github.com/spf13/viper"
)

var (
	GAB_CONFIG *config.Server
	GAB_VP     *viper.Viper
	GAB_DB     *gorm.DB
	GAB_LOG    *zap.Logger
)
