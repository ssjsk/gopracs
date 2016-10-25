package controllers

import (
	"github.com/ssjsk/gowebbook/taskmanager/models"
)

type(
	/*appError struct{
		Error string `json:"error"`
		Message string `json:"message"`
		HttpStatus int `json:"status"`
	}
	errorResource struct{
		Data appError `json:"data"`
	}
	*/
	//For post - /user/register
	UserResource struct{
		Data models.User `json:"data"`
	}

	//for post - /user/login
	LoginResource struct{
		Data LoginModel `json:"data"`
	}

	//Response for authorized user Post - /user/login
	AuthUserResource struct{
		Data AuthUserModel `json:"data"`
	}

	//Model for authentication
	LoginModel struct{
		Email string `json:"email"`
		Password string `json:"password"`
	}

	//Model for authorized user with access token
	AuthUserModel struct{
		User models.User `json:"user"`
		Token string `json:"token"`
	}

	//For POST/PUT - /tasks
	//For Get -/tasks/id
	TaskResource struct {
		Data models.Task `json:"data"`
	}
	//For Get - /tasks
	TasksResource struct{
		Data []models.Task `json:"data"`
	}

	//for POST/PUT - /notes
	NoteResource struct {
		Data NoteModel `json:"data"`
	}

	//For Get /notes
	//For /notes/tasks/id
	NotesResource struct {
		Data []models.TaskNote `json:"data"`
	}

	//Model for TaskNote
	NoteModel struct{
		TaskId string `json:"taskid"`
		Description string `json:"description"`
	}
)

/*func DisplayAppError(w http.ResponseWriter, handlerError error, message string, code int){
	errObj := appError{
		Error: handlerError.Error(),
		Message: message,
		HttpStatus: code,
	}
	log.Printf("[AppError]: %s\n", handlerError)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriterHeader(code)
	if j, err := json.Marshal(errorResource{Data: errObj}); err == nil{
		w.Write(j)
	}
}*/