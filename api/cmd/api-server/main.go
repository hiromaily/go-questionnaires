package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	con "github.com/hiromaily/go-questionnaires/api/controllers"
	"github.com/hiromaily/go-questionnaires/api/libs/fcgi"
	"github.com/hiromaily/golibs/db/mysql"
	lg "github.com/hiromaily/golibs/log"
	u "github.com/hiromaily/golibs/utils"
	"os"
	"time"
)

var (
	serverPort        = 8083
	dbHost            = "mysql-server"
	dbPort     uint16 = 3306
	dbName            = "questionnaire"
	dbUser            = "hiromaily"
	dbPass            = "12345678"
	logPath           = "/var/log/questionnaire/api.log"
)

var (
	portNum    = flag.Int("port", 8083, "Port of server")
	heroku     = flag.Int("heroku", 0, "0:no Heroku, 1:Heroku")
	retryCount = flag.Int("rc", 5, "retry count before starting")
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

func setStatic(r *gin.Engine) {
	rootPath := "/go/src/github.com/hiromaily/go-questionnaires/api"

	r.LoadHTMLGlob(rootPath + "/public/*.html")
	r.GET("/", con.GetIndexAction)
	r.GET("/index.html", con.GetIndexAction)

	//r.Static("/favicon.ico", rootPath+"/statics/favicon.ico")
	//r.Static("/", rootPath+"/public/")
	//r.Static("/", rootPath+"/public/index.html")
	r.Static("/admin", rootPath+"/public/admin")
	r.Static("/js", rootPath+"/public/js")
	r.Static("/css", rootPath+"/public/css")
}

// settings for mysql
func setupDB() {
	if *retryCount == 0 {
		*retryCount = 1
	}

	if *heroku == 1 {
		dbHost = "localhost"
	}

	var err error
	for i := 0; i < *retryCount; i++ {
		lg.Info("connecting to db server ...")
		err := mysql.New(dbHost, dbName, dbUser, dbPass, dbPort)
		if err != nil {
			lg.Errorf("db connection failed. %v", err)
			time.Sleep(3 * time.Second)
			continue
		}
		break
	}
	if err != nil {
		panic("database can not be connected.")
	}
}

func init() {
	flag.Parse()

	//log
	lg.InitializeLog(lg.DebugStatus, lg.LogOff, 99,
		"[Questionnaire]", logPath)

	//port
	if os.Getenv("PORT") != "" {
		*portNum = u.Atoi(os.Getenv("PORT"))
	}
	serverPort = *portNum
	lg.Infof("exported Port is %d", serverPort)

	//Database
	setupDB()
}

func main() {
	//settings
	router := gin.Default()

	//URL for JSON API
	setURL(router)

	//Run
	if *heroku == 1 {
		lg.Info("heroku mode")
		//static
		setStatic(router)

		//for localhost running
		router.Run(fmt.Sprintf(":%d", serverPort))
	} else {
		lg.Info("running on fcgi mode.")
		fcgi.Run(router, fmt.Sprintf(":%d", serverPort))
	}
}
