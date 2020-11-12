package quiz

type QnA struct {
	QnA_Id      int    `bson:"qna_id"`
	ChapterName string `bson:"chapter_name"`
	SubjectName string `bson:"subject_name"`

	Question string `bson:"question"`

	Option1 string `bson:"option_a"`
	Option2 string `bson:"option_b"`
	Option3 string `bson:"option_c"`
	Option4 string `bson:"option_d"`

	Answer string `bson:"answer"`
}
