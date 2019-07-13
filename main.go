package main

import (
	"github.com/double1996/smart-evernote-blog/config"
	"github.com/double1996/smart-evernote-blog/helpers"
	"github.com/double1996/smart-evernote-blog/models"
	"github.com/double1996/smart-evernote-blog/pkg/logger"
	"github.com/double1996/smart-evernote-blog/routers"
	"go.uber.org/zap"
	"html/template"

	//"github.com/double1996/smart-evernote-blog/pkg/evernote"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()
	engine.Use(logger.LogHandler())

	routers.InitRouter(engine)

	setTemplate(engine)
	setSessions(engine)

	config := config.InitConfig()
	//evernote.SyncEverNoteClient()

	db, err := models.InitDB(config)
	if err != nil {
		logger.Fatal("Fatal open database", zap.String("database", err.Error()))
	}
	defer db.Close()

	engine.Run(":" + config.Server.Port)
}

func setTemplate(engine *gin.Engine) {
	funcMap := template.FuncMap{
		"dataFormat": helpers.DateFormat,
		"substring":  helpers.Substring,
	}
	engine.SetFuncMap(funcMap)
	engine.LoadHTMLGlob("templates/*")
}

func setSessions(engine *gin.Engine) {

}
