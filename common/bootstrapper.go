package common

func StartUp() {
	//Initialize AppConfig variable
	initConfig()

	//Initialize private/public keys for JWT authentication
	initKeys()

	//Start mongodb session
	createDbSession()

	//add indexes into mongodb
	addIndexes()
}