package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Db *gorm.DB

type Foo struct {
	Timezone    string
	TheDatetime time.Time
}

func main() {
	// setup database
	TZ := os.Getenv("TZ")
	if TZ == "" {
		TZ = "Local"
	}
	var connectionString = "merico:merico@tcp(mysql:3306)/lake?charset=utf8mb4&parseTime=True&loc=" + url.QueryEscape(TZ)

	var err error

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,  // Slow SQL threshold
			LogLevel:                  logger.Error, // Log level
			IgnoreRecordNotFoundError: true,         // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,         // Disable color
		},
	)

	Db, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		fmt.Println("ERROR: >>> Mysql failed to connect")
		panic(err)
	}
	err = Db.AutoMigrate(&Foo{})
	if err != nil {
		panic(err)
	}

	// construct time.Time from different timezone
	now := time.Now()
	tz0, err := time.Parse("2006-01-02T15:04:05-0700", "2021-01-19T00:00:00+0000")
	if err != nil {
		panic(err)
	}
	tz1, err := time.Parse("2006-01-02T15:04:05-0700", "2021-01-19T01:00:00+0100")
	if err != nil {
		panic(err)
	}

	// print origin value
	fmt.Println("origin values")
	fmt.Printf(" local now: %v\n", now)
	fmt.Printf(" tz0: %v\n", tz0)
	fmt.Printf(" tz1: %v\n", tz1)

	// save to database
	fmt.Println()
	fmt.Println("save to database")
	Db.Exec("truncate table foos")
	Db.Create(&Foo{Timezone: "local now", TheDatetime: now})
	Db.Create(&Foo{Timezone: "tz0", TheDatetime: tz0})
	Db.Create(&Foo{Timezone: "tz1", TheDatetime: tz1})

	// load from database
	foos := make([]*Foo, 0)
	Db.Find(&foos)
	fmt.Println()
	fmt.Println("load from database")
	for _, foo := range foos {
		fmt.Printf(" %v: %v\n", foo.Timezone, foo.TheDatetime)
	}
}
