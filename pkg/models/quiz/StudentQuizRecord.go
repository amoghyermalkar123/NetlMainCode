package quiz

type StudentQuizRecord struct {
	UserId     int    `json:"userid" bson:"userid"`
	QuizId     int    `json:"quizid" bson:"quizid"`
	DateOfQuiz string `json:"date" bson:"dateofquiz"`
	FinishTime string `json:"finishedAt" bson:"finishtime"`
	Score      int    `json:"score" bson:"score"` // client rendered
}
