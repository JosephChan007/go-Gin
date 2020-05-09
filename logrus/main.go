package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"os"
)

var log = logrus.New()

func initLogger() {
	logfile, _ := os.OpenFile("./gin.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	log.Out = logfile
	log.Formatter = &logrus.JSONFormatter{}
	log.Level = logrus.InfoLevel

	gin.DefaultWriter = log.Out
	gin.SetMode(gin.ReleaseMode)
}

func main() {
	initLogger()

	log.WithFields(logrus.Fields{
		"name": "xiaowang",
		"sex":  "male",
	}).Info("print log infomation...")

}
