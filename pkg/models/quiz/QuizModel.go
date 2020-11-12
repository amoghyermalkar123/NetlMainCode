package quiz

type QuizModel struct {
	QuizId      int    `bson:"quiz_id"`
	QnA_Ids     []int  `bson:"qna_ids"`
	ChapterName string `bson:"chapter_name"`
	SubjectName string `bson:"subject_name"`
}
