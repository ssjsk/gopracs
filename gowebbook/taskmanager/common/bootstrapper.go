package common

func Startup(){
	//Initialize AppConfig variable
	initConfig()

	//Initialize private/public keys for JWT authentication
	initKeys()

	//Start MongoDB session
	createDbSession()

	//Add indexes to MongoDB
	addIndexes()
}