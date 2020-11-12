package domain

type SubjectWiseProgress struct {
	SubjectDocID int    `json:"docid" bson:"subjectDocid"`
	DomainID     int    `json:"domainid" bson:"domainid"`
	ProgressRate int    `json:"progress" bson:"progressRate_per"`
	SubjectName  string `json:"subjectname" bson:"subjectname"`
}
