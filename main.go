package main

import (
	"github.com/double1996/joplin-web-blog/config"
	"github.com/double1996/joplin-web-blog/helpers"
	"github.com/double1996/joplin-web-blog/models"
	"github.com/double1996/joplin-web-blog/pkg/joplin"
	"github.com/double1996/joplin-web-blog/pkg/logger"
	"github.com/double1996/joplin-web-blog/routers"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gopkg.in/robfig/cron.v2"
	"html/template"
)

func init() {
	config.InitConfig()
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()
	engine.Use(logger.LogHandler())

	routers.InitRouter(engine)

	setTemplate(engine)
	setSessions(engine)

	db, err := models.InitDB(config.Conf)
	if err != nil {
		logger.Fatal("Fatal open database", zap.String("database", err.Error()))
	}
	defer db.Close()

	c := cron.New()
	c.AddFunc(config.Conf.Joplin.Cron, joplin.SyncJoplinBlog)
	c.Start()

	engine.Run(":" + config.Conf.Server.Port)
}

func setTemplate(engine *gin.Engine) {
	funcMap := template.FuncMap{
		"dateFormat": helpers.DateFormat,
		"substring":  helpers.Substring,
		"truncate":   helpers.Truncate,
		"isOdd":      helpers.IsOdd,
		"isEven":     helpers.IsEven,
	}
	engine.SetFuncMap(funcMap)
	engine.LoadHTMLGlob("./templates/**/*")
}

func setSessions(engine *gin.Engine) {

}
