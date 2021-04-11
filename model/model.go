package model

import "gorm.io/gorm"

type Teacher struct {
	gorm.Model
	Name    string `json:"name"`
	Quizzes []Quiz `json:"quizzes"`
}

type Quiz struct {
	gorm.Model
	Name      string     `json:"name"`
	TeacherId uint       `json:"teacherId"`
	Questions []Question `json:"questions"`
}

type Question struct {
	gorm.Model
	Question string   `json:"question"`
	QuizId   uint     `json:"quizId"`
	Answers  []Answer `json:"answers"`
}

type Answer struct {
	gorm.Model
	Answer     string `json:"answer"`
	IsCorrect  bool   `json:"isCorrect"`
	QuestionId uint   `json:"questionId"`
}

type Student struct {
	gorm.Model
	Name string `json:"name"`
}

type QuizStudent struct {
	gorm.Model
	QuizId    uint
	StudentId uint
}

type QuestionStudent struct {
	gorm.Model
	QuestionId uint
	StudentId  uint
	QuizId     uint
}

type AnswerStudent struct {
	gorm.Model
	IsAnswerCorrect bool `json:"isAnswerCorrect"`
	AnswerId        uint `json:"answerId"`
	StudentId       uint `json:"studentId"`
	QuestionId      uint `json:"questionId"`
	QuizId          uint `json:"quizId"`
}

func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Teacher{}, &Quiz{}, &Question{}, &Answer{}, &Student{}, &QuizStudent{}, &QuestionStudent{}, &AnswerStudent{})
	return db
}
