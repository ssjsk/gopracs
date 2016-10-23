package main
import (
	"log"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func main(){
	session, err := mgo.Dial("localhost")
	if err != nil{
		panic(err)
	}

	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	//get collection
	collection := session.DB("taskdb").C("categories")

	docM := map[string]string{
		"name" : "Open Source",
		"description" : "Tasks for open-source projects",
	}

	err = collection.Insert(docM)
	if err != nil{
		log.Fatal(err)
	}

	docD := bson.D{
		{"name", "Project"},
		{"description", "Project Tasks"},
	}

	//insert a doc slice
	err = collection.Insert(docD)
	if err != nil {
		log.Fatal(err)
	}
}