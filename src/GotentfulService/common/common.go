package common

import (
	"gopkg.in/mgo.v2"
	"time"
	"log"
)

type Configuration struct {
	Server, Port            string
	MongoDBHosts            []string
	DBUser, DBPwd, Database string
	Session                 *mgo.Session
}

func StartUp() *Configuration {
	config := initConfig()
	config.Session = createDbSession(config.MongoDBHosts, config.Database, config.DBUser, config.DBPwd)
	return config
}

func initConfig() *Configuration {
	var config Configuration
	config.loadAppConfig()
	return &config
}

func (config *Configuration) loadAppConfig() {
	config.Server = ""
	config.Port = "8080"

	config.MongoDBHosts = []string{"127.0.0.1"}
	config.Database = "Gotentful"
	config.DBUser = ""
	config.DBPwd = ""
}

func createDbSession(MongoDBHosts []string, Database string, UserName string, Password string) *mgo.Session {

	// We need this object to establish a session to our MongoDB.
	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:    MongoDBHosts,
		Timeout:  60 * time.Second,
		Database: Database,
		Username: UserName,
		Password: Password,
	}

	// Create a session which maintains a pool of socket connections
	// to our MongoDB.
	mongoSession, err := mgo.DialWithInfo(mongoDBDialInfo)
	if err != nil {
		log.Fatalf("Mongo CreateSession: %s\n", err)
	}

	// Reads may not be entirely up-to-date, but they will always see the
	// history of changes moving forward, the data read will be consistent
	// across sequential queries in the same session, and modifications made
	// within the session will be observed in following queries (read-your-writes).
	// http://godoc.org/labix.org/v2/mgo#Session.SetMode
	mongoSession.SetMode(mgo.Monotonic, true)
	return mongoSession
}
