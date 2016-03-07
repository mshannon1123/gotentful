package sync

import (
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

const basePath string = "/sync"

func RegisterRoutes(router *httprouter.Router, session *mgo.Session, databaseName string) {

}
