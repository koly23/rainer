package labels

import (
	"github.com/gin-gonic/gin"
	"github.com/koly23/rainer/db"
	"github.com/koly23/rainer/logger"
	"go.mongodb.org/mongo-driver/bson"
)

func Create(c *gin.Context) {
	oneDb := db.NewDb()
	data, err := c.GetRawData()
	if err != nil {
		logger.InfoE("db create error", err)
	}
	var body interface{}
	err = bson.UnmarshalExtJSON(data, true, &body)
	logger.InfoA("create label with data", body)
	if err != nil {
		logger.InfoE("parse body failed", err)
	}
	id := oneDb.Create(db.Labels, body)
	c.JSON(200, gin.H{
		"_id": id,
	})
}

func All(c *gin.Context) {
	logger.Info("show all labels")
	oneDb := db.NewDb()
	all := oneDb.All(db.Labels, 0, 0)
	c.JSON(200, &all)
}
