package student

type LoginPayload struct {
	Email    string `bson:"email"`
	Password string `bson:"password"`
}
