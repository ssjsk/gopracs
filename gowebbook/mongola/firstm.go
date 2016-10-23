package main

import(
	"fmt"
	"log"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Category struct{
	Id bson.ObjectId `bson:"_id,omitempty"`
	Name string
	Description string
}

func main(){
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()


	//set session mode to monotonic
	session.SetMode(mgo.Monotonic, true)

	collection := session.DB("taskdb").C("categories")

	doc := Category{
		bson.NewObjectId(),
		"Open Source",
		"Tasks for open-source projects",
	}

	//insert category object
	err = collection.Insert(&doc)
	if err != nil{
		log.Fatal(err)
	}

	err = collection.Insert(&Category{bson.NewObjectId(), "R & D", "R & D Tasks"},
			&Category{bson.NewObjectId(), "Project Management", "Project Management Tasks"})

	var count int
	count, err = collection.Count()
	if err != nil {
		log.Fatal(err)
	}else {
		fmt.Printf("%D records inserted", count)
	}

	
}