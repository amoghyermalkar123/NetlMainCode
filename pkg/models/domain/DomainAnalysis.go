package domain

type DomainAnalysis struct {
	UserId   int   `bson:"userid"`
	DomainID int   `bson:"domainid"`
	Subjects []int `bson:"subjectanalysisIDS"`
}
