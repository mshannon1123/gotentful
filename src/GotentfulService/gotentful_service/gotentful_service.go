package gotentful_service

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"
	"reflect"
	"log"
)


// Standard Gotentful Service
type GotentfulService struct {
	BasePath		string
	Session                 *mgo.Session
	DatabaseName		string
	CollectionName		string
	JsonType 		reflect.Type
	SingleLookupFunc	func(ps httprouter.Params, id string, action string) interface{}
	MultiLookupFunc		func(ps httprouter.Params) interface{}
	ValidateFunc		func(ps httprouter.Params, action string, record *interface{}, session *mgo.Session) bool
}

func (mongo *GotentfulService) RegisterStandardRoutes(router *httprouter.Router) {
	router.GET(mongo.BasePath, mongo.List)
	router.POST(mongo.BasePath, mongo.Create)
	router.GET(mongo.BasePath+"/:id", mongo.Index)
	router.PUT(mongo.BasePath+"/:id", mongo.Upsert)
	router.DELETE(mongo.BasePath+"/:id", mongo.Remove)
}

// Generate Standard Structures and Indexes
func (mongo *GotentfulService) GenerateIndexes(indexes []mgo.Index) {
	for _, index := range indexes {
		mongo.Session.DB(mongo.DatabaseName).C(mongo.CollectionName).EnsureIndex(index)
	}
}

func (mongo *GotentfulService) List(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	slice := reflect.MakeSlice(reflect.SliceOf(mongo.JsonType), 0, 0)

	// Create a pointer to a slice value and set it to the slice
	records := reflect.New(slice.Type())
	records.Elem().Set(slice)

	dbs := mongo.Session.Copy()
	defer dbs.Close()
	c := dbs.DB(mongo.DatabaseName).C(mongo.CollectionName)

	err := c.Find(mongo.MultiLookupFunc(ps)).All(records.Interface())
	if err != nil {
		log.Fatalf("Unable to retrieve %ss: %s\n", mongo.CollectionName, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(records.Elem().Interface())
}

func (mongo *GotentfulService) Index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	id := ps.ByName(":id")
	if (id == "") {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	dbs := mongo.Session.Copy()
	defer dbs.Close()
	c := dbs.DB(mongo.DatabaseName).C(mongo.CollectionName)

	record := reflect.New(mongo.JsonType).Elem().Interface()
	err := c.Find(mongo.SingleLookupFunc(ps, id, "index")).One(&record)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&record)
}

func (mongo *GotentfulService) Create(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	decoder := json.NewDecoder(r.Body)

	record := reflect.New(mongo.JsonType).Elem().Interface()
	err := decoder.Decode(&record)
	if err != nil {
		log.Fatalf("Unable to parse %s (Create): %s\n", mongo.CollectionName, err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	dbs := mongo.Session.Copy()
	defer dbs.Close()
	c := dbs.DB(mongo.DatabaseName).C(mongo.CollectionName)

	newId := "TEST" //bson.NewObjectId()

	if !mongo.ValidateFunc(ps, "create", &record, dbs) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	change := mgo.Change{
		Update: record,
		Upsert: true,
		Remove: false,
		ReturnNew: true,
	}

	_, err2 := c.Find(mongo.SingleLookupFunc(ps, newId, "create")).Apply(change, &record)
	if err2 != nil {
		log.Fatalf("Unable to create %s: %s\n", mongo.CollectionName, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(record)
}

func (mongo *GotentfulService) Upsert(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	id := ps.ByName(":id")
	if (id == "") {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	decoder := json.NewDecoder(r.Body)

	record := reflect.New(mongo.JsonType).Elem().Interface()
	err := decoder.Decode(&record)
	if err != nil {
		log.Fatalf("Unable to parse %s (Upsert): %s\n", mongo.CollectionName, err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	dbs := mongo.Session.Copy()
	defer dbs.Close()
	c := dbs.DB(mongo.DatabaseName).C(mongo.CollectionName)

	if !mongo.ValidateFunc(ps, "upsert", &record, dbs) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	change := mgo.Change{
		Update: record,
		Upsert: true,
		Remove: false,
		ReturnNew: true,
	}

	info, err2 := c.Find(mongo.SingleLookupFunc(ps, id, "upsert")).Apply(change, &record)
	if err2 != nil {
		log.Fatalf("Unable to upsert %s (%s): %s\n", mongo.CollectionName, id, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if info.Matched > 0 {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusCreated)
	}

	json.NewEncoder(w).Encode(record)
}

func (mongo *GotentfulService) Remove(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	record := reflect.New(mongo.JsonType).Elem().Interface()
	id := ps.ByName(":id")
	if (id == "") {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	dbs := mongo.Session.Copy()
	defer dbs.Close()
	c := dbs.DB(mongo.DatabaseName).C(mongo.CollectionName)

	if !mongo.ValidateFunc(ps, "remove", &record, dbs) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	change := mgo.Change{
		Upsert: false,
		Remove: true,
		ReturnNew: false,
	}

	info, err := c.Find(mongo.SingleLookupFunc(ps, id, "remove")).Apply(change, &record)
	if err != nil {
		log.Fatalf("Unable to remove %s (%s): %s\n", mongo.CollectionName, id, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if info.Removed == 0 {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusOK)
	}

	json.NewEncoder(w).Encode(record)
}
