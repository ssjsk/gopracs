package controllers

import(
	"encoding/json"
	"net/http"
	"github.com/ssjsk/gowebbook/taskmanager/common"
	"github.com/ssjsk/gowebbook/taskmanager/data"
	"github.com/ssjsk/gowebbook/taskmanager/models"
)

//Handler for HTTP Post "/users/register"
//Add a new User document

func Register(w http.ResponseWriter, r *http.Request){
	var dataResource UserResource
	//decode incoming user json
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(
				w,
				err,
				"Invalid User Data",
				500,
			)
		return
	}
	user := &dataResource.Data
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("users")
	repo := &data.UserRepository{c}
	//Insert user document
	repo.CreateUser(user)
	//clean up hashpassword to eliminate from response
	user.HashPassword = nil
	if j, err := json.Marshal(UserResource{Data: *user}); err != nil {
		common.DisplayAppError(
			w,
			err,
			"An unexpected error has occurred",
			500,
		)
		return
	} else {
		w.Header().Set("Content-Type","application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(j)
	}
}

//Handler for HTTP post "/users/login"
//Authenticate with username and password
func Login(w http.ResponseWriter, r *http.Request){
	var dataResource LoginResource
	var token string
	//Decode incoming Login json
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil{
		common.DisplayAppError(
			w,
			err,
			"Invalid Login data",
			500,
		)
		return
	}
	loginModel := dataResource.Data
	loginUser := models.User{
		Email: loginModel.Email,
		Password: loginModel.Password,
	}
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("users")
	repo := &data.UserRepository{C: c}
	//Authenticate logged in user
	user, err := repo.Login(loginUser)
	if err != nil{
		common.DisplayAppError(w, err, "Invalid login credentials", 401,)
		return
	} 
	//Generate JWT token
	token, err = common.GenerateJWT(user.Email, "member")
	if err != nil {
		common.DisplayAppError(w, err, "Error while generating access token", 500,)
		return
	}
	
	w.Header().Set("Content-Type","application/json")
	user.HashPassword = nil
	authUser := AuthUserModel{
		User: user,
		Token: token,
	}
	j, err := json.Marshal(AuthUserResource{Data: authUser})
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
	w.Write(j)
}
