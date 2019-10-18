package gorm

import (
	"fmt"

	_ "github.com/lib/pq"
	"github.com/jinzhu/gorm"

	"github.com/mergeforces/mergeforces-service/config"
)

func New(conf *config.Conf) (*gorm.DB, error) {
	dbConn := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		conf.Db.Host, 5432, conf.Db.Username, conf.Db.Password, conf.Db.DbName)

	return gorm.Open("postgres", dbConn)
}