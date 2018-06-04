package server

import (
	"time"
	"net/http"

	"github.com/Gin/interate-with-vue/collector"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
)

type Server struct {
	r         *gin.Engine
	collector *collector.Collector
}

func NewServer() *Server {
	r := gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.MaxAge = 5 * time.Minute 

	r.Use(cors.New(corsConfig))
	return &Server{
		r:         r,
		collector: collector.NewCollector(),
	}
}

func (self *Server) GetData(c *gin.Context) {
	data, err := self.collector.GetData()
	if err != nil {
		c.JSON(
			http.StatusOK,
			gin.H{
				"success": false,
				"reason": "error query data",
			},
		)
		return
	}
	c.JSON(
		http.StatusOK,
		gin.H{
			"success": true,
			"data": data,
		},
	)
}

func (self *Server) Run() {
	self.r.Use(static.Serve("/", static.LocalFile("/home/nguyen.thi.tu.anh/go/src/github.com/Gin/interate-with-vue/app", true)))

	api := self.r.Group("/api")
	api.GET("/data", self.GetData)
	
	self.r.Run(":8001")
}
