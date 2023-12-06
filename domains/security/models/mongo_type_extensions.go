package models

type UserCredentialsMongoDB struct {
	Username string `bson:"username"`
	Password string `bson:"password"`
}

func (protoUserCredentials *UserCredentials) ConvertToMongoReadable() (*UserCredentialsMongoDB, error) {
	return &UserCredentialsMongoDB{
		Username: protoUserCredentials.Username,
		Password: protoUserCredentials.Password,
	}, nil
}

func (mongoUserCredentials *UserCredentialsMongoDB) ConvertToProtoMessage() (*UserCredentials, error) {
	return &UserCredentials{
		Username: mongoUserCredentials.Username,
		Password: mongoUserCredentials.Password,
	}, nil
}
