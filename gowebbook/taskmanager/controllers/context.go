package controllers

import(
	"gopkg.in/mgo.v2"
	"github.com/ssjsk/gowebbook/taskmanager/common"
)

//struct used for maintaining http request context
type Context struct{
	MongoSession *mgo.Session
}

//close mgo.session
func(c *Context) Close(){
	c.MongoSession.Close()
}

//Return mgo.Collection for the given name
func(c *Context) DbCollection(name string) *mgo.Collection{
	return c.MongoSession.DB(common.AppConfig.Database).C(name)
}

//Create a new context object for each http request
func NewContext() *Context{
	session := common.GetSession().Copy()
	context := &Context{
		MongoSession: session,
	}
	return context
}