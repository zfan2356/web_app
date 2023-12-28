package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var db *sqlx.DB

func InitMySQL() (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		viper.GetString("mysql.user"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.host"),
		viper.GetInt("mysql.port"),
		viper.GetString("mysql.dbname"),
	)

	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("connect db failed, err = %v\n", err)
		return err
	}
	db.SetMaxIdleConns(viper.GetInt("mysql.max_idle_conns"))
	db.SetMaxOpenConns(viper.GetInt("mysql.max_open_conns"))

	return err
}

func Close() {
	err := db.Close()
	if err != nil {
		zap.L().Fatal("close mysql failed: ", zap.Error(err))
	}
}
