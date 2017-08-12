package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	con "github.com/hiromaily/go-questionnaires/api/controllers"
	"github.com/hiromaily/go-questionnaires/api/libs/fcgi"
	"github.com/hiromaily/golibs/db/mysql"
	lg "github.com/hiromaily/golibs/log"
)

var (
	serverPort        = 8084
	dbHost            = "localhost"
	dbPort     uint16 = 13306
	dbName            = "questionnaire"
	dbUser            = "hiromaily"
	dbPass            = "12345678"
	logPath           = "/var/log/questionnaire.log"
)

var (
	docker = flag.Int("docker", 0, "0:local, 1:Docker")
)

// router
func setURL(r *gin.Engine) {
	//Get questionnaire list
	r.GET("/api/ques", con.GetQuesAction)
	//Register new questionnaire
	r.POST("/api/ques", con.PostRegQuesAction)
	//Delete questionnaire by ID
	r.DELETE("/api/ques/:id", con.DeleteQuesAction)

	//Get answer by ID
	r.GET("/api/answer/:id", con.GetAnswerAction)
	//Register answer by ID
	r.POST("/api/answer/:id", con.PostAnswerAction)
}

// settings for mysql
func setupDB() {
	mysql.New(dbHost, dbName, dbUser, dbPass, dbPort)
}

func init() {
	flag.Parse()

	//log
	lg.InitializeLog(lg.DebugStatus, lg.LogOff, 99,
		"[Questionnaire]", logPath)

	//Docker Settings
	if *docker == 1 {
		lg.Info("docker mode")
		serverPort = 8083
		dbHost = "mysql-server"
		dbPort = 3306
		//dbName     = "questionnaire"
		//dbUser     = "hiromaily"
		//dbPass     = "12345678"
		logPath = "/var/log/questionnaire/api.log"
	}

	setupDB()
}

func main() {
	//settings
	router := gin.Default()

	//URL for JSON API
	setURL(router)

	lg.Debugf("docker flg is %d", *docker)

	//Run
	if *docker == 1 {
		lg.Info("running on fcgi mode.")
		fcgi.Run(router, fmt.Sprintf(":%d", serverPort))
	} else {
		router.Run(fmt.Sprintf(":%d", serverPort))
	}
}
