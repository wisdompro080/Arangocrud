package models

import (
	"github.com/arangodb/go-driver"
	log "github.com/sirupsen/logrus"
)

type Config struct {
	Port   string `json:"port"`
	Arango struct {
		Database    string `json:"database"`
		Server      string `json:"server"`
		UserName    string `json:"username"`
		Password    string `json:"password"`
		Collections struct {
			User string `json:"user"`
		} `json:"collections"`
	} `json:"arango"`
	LogLevel log.Level `json:"logLevel"`
}
type Document struct {
	Id   string `json:"id,omitempty"`//
	Name string `json:"name,omitempty"`//required
//	rank
//	recentresult{
//		phy int
//		chem int
//}
}
type DbConnection struct {
	Db driver.Database
	Col driver.Collection
}
