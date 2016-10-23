package controllers

import(
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ssjsk/gowebbook/taskmanager/data"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//handler for http posts - "/tasks"
//insert a new task document
func CreateTask(w http.ResponseWriter, r *http.Request){
	var dataResource TaskResource
	//Decode incoming Task json
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid task data",
			500,
		)
		return
	}
	task := &dataResource.data
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("tasks")
	repo := &data.TaskRepository{c}
	//Insert a task document
	repo.Create(task)
	if j, err := json.Marshal(TaskResource{Data: *task}); err != nil{
		common.DisplayAppError(
			w,
			err,
			"An unexpected error has occurred",
			500,
		)
		return
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(j)
	}
}
//Handler for http Get -"/tasks"
//Returns all task documents
func GetTasks(w http.ResponseWriter, r *http.Request){
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("task")
	repo := &data.TaskRepository{c}
	tasks := repo.GetAll()
	j, err := json.Marshal(TaskResource{Data: tasks})
	if err != nil{
		common.DisplayAppError(
			w,
			err,
			"An unexpected error has occurred",
			500,
		)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type","application/json")
	w.Write(j)
}

//Handler for HTTP Get - "/task/{id}"
//Returns a single task document by id
func GetTaskById(w http.ResponseWriter, r *http.Request){
	//Get id from incoming url
	vars := mux.Vars(r)
	id := vars["id"]
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("tasks")
	repo := &data.TaskRepository{c}
	task, err := repo.GetById(id)
	if err != nil {
		if err == mgo.ErrNotFound{
			w.WriteHeader(http.StatusNoContent)
			return
		} else {
			common.DisplayAppError(
				w,
				err,
				"An unexpected error has occurred",
				500,
			)
			return
		}
	}
	if j, err := json.Marshal(task); err != nil{
		common.DisplayAppError(
			w,
			err,
			"An unexpected error has occurred",
			500,
		)
		return
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(j)
	}
}

//handler for http get - "/tasks/users/{id}"
//returns all tasks created by a user
func GetTasksByUser(w http.ResponseWriter, r *http.Request){
	//Get id from incoming url
	vars := mux.Vars(r)
	user := vars["id"]
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("task")
	repo := &data.TaskRepository{c}
	tasks := repo.GetByUser(user)
	j, err := json.Marshal(TaskResource{Data: tasks})
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"An unexpected error has occurred",
			500,
		)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

//Handler for HTTP Put - "/tasks/{id}"
//Update an existing task document
func UpdateTask(w http.ResponseWriter, r *http.Request){
	//Get id from incoming url
	vars := mux.Vars(r)
	id := bson.ObjectIdHex(vars["id"])
	var dataResource TaskResource

	//decode incoming task json
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid Task Data",
			500,
		)
		return
	}
	task := &dataResource.Data
	task.Id = id
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("tasks")
	repo := &data.TaskRepository{c}
	//Update an existing task document
	if err := repo.Update(task); err != nil{
		common.DisplayAppError(
			w,
			err,
			"An unexpected error has occurred",
			500,
		)
		return
	} else {
		w.WriteHeader(http.StatusNoContent)
	}
}

//Handler for HTTP Delete - "/tasks/{id}"
//Delete an existing task document
func DeleteTask(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id := vars["id"]
	context := NewContext()
	defer context.Close()
	c  := context.DbCollection("tasks")
	repo := &data.TaskRepository{c}
	//Delete an existing Task document
	err := repo.Delete(id)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"An unexpected error has occurred",
			500,
		)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}