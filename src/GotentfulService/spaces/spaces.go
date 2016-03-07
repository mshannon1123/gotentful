package spaces

import (
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"GotentfulService/gotentful_service"
	"GotentfulService/models"
	"reflect"
	"gopkg.in/mgo.v2/bson"
)

const collectionName string = "spaces"
const basePath string = "/spaces"

func RegisterRoutes(router *httprouter.Router, session *mgo.Session, databaseName string) {

	asset := gotentful_service.GotentfulService{
		BasePath: basePath,
		Session: session,
		DatabaseName: databaseName,
		CollectionName: collectionName,
		JsonType: reflect.TypeOf(models.Space{}),
	}


	//SingleLookupFunc
	asset.SingleLookupFunc = func(ps httprouter.Params, id string, action string) interface{} {
		return bson.M{"_id": id}
	}

	//MultiLookupFunc
	asset.MultiLookupFunc = func(ps httprouter.Params) interface{} {
		return nil;
	}

	//ValidateFunc
	asset.ValidateFunc = func(ps httprouter.Params, action string, record *interface{}, session *mgo.Session) bool {
		record_value := *record
		_, ok := record_value.(models.Space)
		if ok {
			return true
		}
		return false
	}

	//asset.GenerateIndexes()
	asset.RegisterStandardRoutes(router)
}