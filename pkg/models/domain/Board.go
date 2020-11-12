package domain

import (
	model "netl/pkg/models/quiz"
)

type Board struct {
	UserId   int                       `json:"userid"`
	Subjects []SubjectWiseProgress     `json:"subjects"`
	Quizes   []model.StudentQuizRecord `json:"quizes"`
}
