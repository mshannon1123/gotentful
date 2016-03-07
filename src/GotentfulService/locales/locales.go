package locales

import (
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"GotentfulService/gotentful_service"
	"GotentfulService/models"
	"reflect"
	"gopkg.in/mgo.v2/bson"
)

const collectionName string = "locales"
const basePath string = "/locales"

func RegisterRoutes(router *httprouter.Router, session *mgo.Session, databaseName string) {

	asset := gotentful_service.GotentfulService{
		BasePath: basePath,
		Session: session,
		DatabaseName: databaseName,
		CollectionName: collectionName,
		JsonType: reflect.TypeOf(models.Locale{}),
	}

	//SingleLookupFunc
	asset.SingleLookupFunc = func(ps httprouter.Params, id string, action string) interface{} {
		space_id := ps.ByName(":space_id")
		return bson.M{"space_id": space_id,  "_id": id}
	}

	//MultiLookupFunc
	asset.MultiLookupFunc = func(ps httprouter.Params) interface{} {
		space_id := ps.ByName(":space_id")
		return bson.M{"space_id": space_id}
	}

	//ValidateFunc
	asset.ValidateFunc = func(ps httprouter.Params, action string, record *interface{}, session *mgo.Session) bool {
		return true
	}

	//asset.GenerateIndexes()
	asset.RegisterStandardRoutes(router)
}