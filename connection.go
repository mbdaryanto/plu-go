package main

import (
	"fmt"
	"os"

	"github.com/go-sql-driver/mysql"
)

func GetDsn() string {
	config := mysql.NewConfig()
	config.User = os.Getenv("DB_USER")
	config.Passwd = os.Getenv("DB_PASS")
	config.Net = "tcp"
	config.Addr = fmt.Sprintf("%v:%v", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"))
	config.DBName = os.Getenv("DB_DATABASE")
	config.AllowNativePasswords = true
	return config.FormatDSN()
}
