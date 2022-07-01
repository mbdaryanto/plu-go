package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Env struct {
	db *gorm.DB
}

type Item struct {
	IDItem      int `gorm:"primaryKey"`
	Kode        string
	Barcode     string
	Nama        string
	KodePabrik  string
	HargaNormal float64
	HargaJual   float64
}

func main() {
	if len(os.Args) > 1 {
		command := os.Args[1]
		// if using create_config on command line
		if command == "create_config" {
			CreateConfig()
			return
		}
	}

	var dsn string
	if setting, err := GetSetting(); err != nil {
		log.Fatal(err)
	} else {
		dsn = setting.GetDsn()
	}
	// dsn := GetDsn()
	// fmt.Println(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	env := &Env{db: db}
	router := gin.Default()
	// router.Static("/", "./public/index.html")
	router.GET("/item", env.getItem)
	router.GET("/", func(c *gin.Context) {
		c.File("./simple.html")
	})

	router.Run("localhost:8080")
}

func (e *Env) getItem(c *gin.Context) {
	code := c.Query("code")

	// sql := e.db.ToSQL(func(tx *gorm.DB) *gorm.DB {
	// 	return tx.Table("mitem").Where(
	// 		"Aktif", "Ya",
	// 	).Where(
	// 		e.db.Where("Kode like ?", code).Or("Barcode like ?", code),
	// 	).Clauses(clause.OrderBy{
	// 		Expression: clause.Expr{SQL: "Barcode like ? desc", Vars: []interface{}{code}, WithoutParentheses: true},
	// 	}).Limit(1).Find(&Item{})
	// })
	// fmt.Println(sql)
	// SELECT * FROM `mitem` WHERE `Aktif` = 'Ya' AND (Kode like '00000180' OR Barcode like '00000180') ORDER BY Barcode like '00000180' desc LIMIT 1

	var item Item
	e.db.Table(
		"mitem",
	).Select(
		"IDItem, Kode, Barcode, Nama, KodePabrik, HargaNormal, HargaJual",
	).Where(
		"Aktif", "Ya",
	).Where(
		e.db.Where(
			"Kode like ?", code,
		).Or(
			"Barcode like ?", code,
		),
	).Clauses(clause.OrderBy{
		Expression: clause.Expr{
			SQL:                "Barcode like ? desc",
			Vars:               []interface{}{code},
			WithoutParentheses: true,
		},
	}).Limit(1).Find(&item)

	c.IndentedJSON(200, item)
}
