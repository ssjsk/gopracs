package main

import(
	"log"
	"time"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"fmt"
)

type Task struct{
	Description string
	Due time.Time
}

type Category struct {
	Id bson.ObjectId `bson:"_id,omitempty"`
	Name string
	Description string
	Tasks	[]Task
}

func main(){
	session, err := mgo.Dial("localhost")
	if err != nil{
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	collection := session.DB("taskdb").C("categories")

	collection.RemoveAll(nil)
	//embedde child collection
	doc := Category{
		bson.NewObjectId(),
		"Open-Source",
		"Tasks for open-source projects",
		[]Task{
			Task{"Create project in mgo", time.Date(2015, time.August, 10, 0, 0, 0, 0,
                  time.UTC)},
            Task{"Create REST API", time.Date(2015, time.August, 20, 0, 0, 0, 0, time.UTC)},
            },
	}
	//Insert a category
	err = collection.Insert(&doc)
	if err != nil{
		log.Fatal(err)
	}

	err = collection.Insert(&Category{
		bson.NewObjectId(),
		"C# ASP.NET MVC Project",
		"Tasks for MVC ASP.NET projects",
		[]Task{
			Task{"Create project in MVC 5", time.Date(2015, time.August, 10, 0, 0, 0, 0,
                  time.UTC)},
            Task{"Create REST API", time.Date(2015, time.August, 20, 0, 0, 0, 0, time.UTC)},
            Task{"Create  Database", time.Date(2015, time.August, 22, 0, 0, 0, 0, time.UTC)},
            },})
	fmt.Println("lets read documents now")

	iter := collection.Find(nil).Iter()
	result := Category{}
	for iter.Next(&result){
		fmt.Printf("Category:%s, Description:%s\n", result.Name, result.Description)
		tasks := result.Tasks

		for _, v := range tasks{
			fmt.Printf("Task: %s Due %v\n", v.Description, v.Due)
		}
	}
	if err = iter.Close(); err != nil{
		log.Fatal(err)
	}

	fmt.Println("Lets sort records")

	iter = collection.Find(nil).Sort("name").Iter()
	result = Category{}
	for iter.Next(&result){
		fmt.Printf("Category:%s, Description:%s\n", result.Name, result.Description)
		tasks := result.Tasks

		for _, v := range tasks{
			fmt.Printf("Task: %s Due %v\n", v.Description, v.Due)
		}
	}
	if err = iter.Close(); err != nil{
		log.Fatal(err)
	}

	fmt.Println("Find only 1 document")

	result = Category{}
	err = collection.Find(bson.M{"name":"Open-Source"}).One(&result)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Printf("Category:%s, Description:%s\n", result.Name, result.Description)
	tasks := result.Tasks

	for _, v := range tasks{
		fmt.Printf("Task: %s Due %v\n", v.Description, v.Due)
	}

}