package main

import (
	_ "go_worlder_system/docs"
	"go_worlder_system/infrastructure/router"
	"time"

	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
)

// @title Worlder Prototype API
// @version 1.0
// @description This is a prototype worlder server
// @license.name Worlder Team
// @host 118.27.23.183:8080
// @BasePath /api

// @tag.name UserAuth
// @tag.description
// @tag.name Profile
// @tag.description
// @tag.name Order
// @tag.description
// @tag.name Brand
// @tag.description
// @tag.name Product
// @tag.description
// @tag.name Project
// @tag.description
// @tag.name Search
// @tag.description
// @tag.name Inventory
// @tag.description
func main() {
	// log setting
	log.SetFormatter(&log.JSONFormatter{
		PrettyPrint: true,
	})
	log.SetReportCaller(true)

	// time zone
	const location = "Asia/Tokyo"
	loc, err := time.LoadLocation(location)
	if err != nil {
		loc = time.FixedZone(location, 9*60*60)
	}
	time.Local = loc

	router.Start()
}
