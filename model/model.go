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

func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Teacher{}, &Quiz{}, &Question{}, &Answer{})
	return db
}
