package controllers

import (
	"github.com/bockbone/taskmanager/common"
	"gopkg.in/mgo.v2"
)

//Struct used for maintaining Http request context
type Context struct {
	MongoSession *mgo.Session
}

//Close mgo.Session
func (c *Context) Close() {
	c.MongoSession.Close()
}

//Returns mgo.collection for the given name
func (c *Context) DbCollection(name string) *mgo.Collection {
	return c.MongoSession.DB(common.AppConfig.Database).C(name)
}


//Create new context object for each http request
func NewContext() *Context {
	session := common.GetSession().Copy()
	context := &Context{
		MongoSession: session,
	}
	return context
}