package mongo

import (
	mgo "gopkg.in/mgo.v2"
	"github.com/astaxie/beego"
)

var session *mgo.Session

func Session() *mgo.Session {
	if session == nil {
		var err error
		session, err = mgo.Dial(beego.AppConfig.String("dbLink"))
		if err != nil {
			panic(err)
		}
	}
	return session
}

func Collection(collection string) *mgo.Collection {
	return Session().DB(beego.AppConfig.String("dbName")).C(collection)
}

func Close() {
	if session != nil {
		session.Close()
		if err := recover(); err != nil {
			panic(err)
		}
	}
}
