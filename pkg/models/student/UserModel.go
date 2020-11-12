package student

type User struct {
	UserId int    `bson:"user_id"`
	Email  string `bson:"email"`
	Year   int    `bson:"year"`
	Branch string `bson:"branch"`
}
