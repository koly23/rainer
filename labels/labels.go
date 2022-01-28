package labels

import (
	"github.com/gin-gonic/gin"
	"github.com/koly23/rainer/db"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

func Create(c *gin.Context) {
	oneDb := db.New()
	data, err := c.GetRawData()
	if err != nil {
		log.Println(err)
	}
	var body interface{}
	err = bson.UnmarshalExtJSON(data, true, &body)
	if err != nil {
		log.Println("parse body failed", err)
	}
	id := oneDb.Create(db.Labels, body)
	c.JSON(200, gin.H{
		"_id": id,
	})
}

func All(c *gin.Context) {
	oneDb := db.New()
	all := oneDb.All(db.Labels, 0, 0)
	c.JSON(200, &all)
}
