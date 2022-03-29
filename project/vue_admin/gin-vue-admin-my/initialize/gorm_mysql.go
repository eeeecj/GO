/*
 * @Author: eeeecj
 * @Date: 2022-03-25 15:05:26
 * @LastEditors: eeeecj
 * @LastEditTime: 2022-03-27 12:06:47
 * @Description:
 */

package initialize

import (
	"errors"
	"project/vue_admin/gin-vue-admin-my/global"

	"gorm.io/gorm/schema"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GormMysql() (*gorm.DB, error) {
	m := global.GAB_CONFIG.Mysql

	if m.DbName == "" {
		return nil, errors.New("database name is empty")
	}
	mysqlConfig := mysql.Config{
		DSN:                       m.Dsn(),
		DefaultStringSize:         191,
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}
	db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{
		SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "t_",  // table name prefix, table for `User` would be `t_users`
			SingularTable: false, // use singular table name, table for `User` would be `user` with this option enabled
		},
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	if err != nil {
		return nil, err
	}
	DB, err := db.DB()
	if err != nil {
		return nil, err
	}
	DB.SetMaxIdleConns(m.MaxIdleConn)
	DB.SetMaxOpenConns(m.MaxOpenConn)
	return db, nil
}
