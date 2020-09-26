package service

import (
	"context"
	"encoding/json"
	"github.com/arangodb/go-driver"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	log "github.com/sirupsen/logrus"
	"test/models"

	//log "github.com/sirupsen/logrus"
	"test/config"
	)
var DbConnect models.DbConnection
func C()  {
	DbConnect.Col,DbConnect.Db =DbConnection()
}
func Create1(c *gin.Context) {
	C()
	db:=config.DbConnect.Db
	collectionName := config.Config.Arango.Collections.User

	ctx := context.Background()
	var details models.Document
	err := c.BindJSON(&details)
	if err != nil {
		log.Fatal(err)
	}

	query := "INSERT {name:" + "'" + details.Name + "'" + ",id:'" + details.Id + "'} IN " + collectionName
	_, _ = db.Query(ctx, query, nil)
}
func Read1(c *gin.Context) {
	C()
    db:=config.DbConnect.Db
	ctx := context.Background()
	query := "FOR d IN Documents RETURN d"
	cursor, err := db.Query(ctx, query, nil)

	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close()
	for {
		var doc models.Document
		_, err := cursor.ReadDocument(ctx, &doc)
		if driver.IsNoMoreDocuments(err) {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		jdoc, err := json.Marshal(doc)
		if err != nil {
			log.Fatal("marshal error ", err)
		}
		c.String(200, string(jdoc))
	}
}
func Remove1(c *gin.Context) {
	C()
	db:=config.DbConnect.Db
	collectionName := config.Config.Arango.Collections.User
	ctx := context.Background()
	key := c.Param("id")
	query := "REMOVE'" + key + "'IN " + collectionName
	_, err := db.Query(ctx, query, nil)
	if err != nil {
		log.Fatal("error in removal", err)
	}
	log.Info("Removed Succesfully")
}
func Update1(c *gin.Context) {
	C()
	db:=config.DbConnect.Db
	//collectionName := config.Config.Arango.Collections.User
	ctx := context.Background()
	var doc models.Document
	_ = c.ShouldBindWith(&doc, binding.JSON)
	//key := c.Param("id")

	//query := "UPDATE'" + key + "'WITH{id:'" + doc.Id + "',name:'" + doc.Name + "'} IN " + collectionName
	query:="for u in Documents filter u._key=='62004' update u with {value:'hell',id:'ya'} in Documents"
	//time.r339
	_, err := db.Query(ctx, query, nil)
	if err != nil {
		log.Fatal("error in removal", err)
	}
	log.Info("updated successfully", doc)
}
