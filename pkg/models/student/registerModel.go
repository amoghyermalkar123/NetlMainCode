package student

type RegisterModel struct {
	UserId   int    `bson:"user_id"`
	Email    string `bson:"email"`
	Password string `bson:"password"`
	Year     int    `bson:"year"`
	Branch   string `bson:"branch"`
}
