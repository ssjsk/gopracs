package data

import(
	"github.com/ssjsk/gowebbook/taskmanager/models"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserRepository struct{
	C *mgo.Collection
}

func (r *UserRepository) CreateUser(user *models.User) error{
	obj_id := bson.NewObjectId()
	user.Id = obj_id
	hpass, err := bcrypt.GenerateFromPassword([](user.Passord), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	user.HashPassword = hpass
	//clear incoming text password
	user.Passord = ""
	err = r.C.Insert(&user)
	return err
}

func (r *UserRepository) Login(user models.User) (u models.User, err error){
	err = r.C.Find(bson.M{"email": user.Email}).One(*u)
	if err != nil {
		return
	}
	//validate password
	err = bcrypt.CompareHashAndPassword(u.HashPassword, []byte(user.Passord))
	if err != nil{
		u = models.User{}
	}
	return
}