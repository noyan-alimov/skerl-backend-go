package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/noyan-alimov/skerl-backend/model"
	"gorm.io/gorm"
)

type completeQuizReqBody struct {
	StudentId    int   `json:"studentId"`
	QuizId       int   `json:"quizId"`
	QuestionsIds []int `json:"questionsIds"`
	AnswersIds   []int `json:"answersIds"`
}

type completeQuizResBody struct {
	CorrectlyAnsweredQuestions   []model.AnswerStudent `json:"correctlyAnsweredQuestions"`
	InCorrectlyAnsweredQuestions []model.AnswerStudent `json:"inCorrectlyAnsweredQuestions"`
}

func CompleteQuiz(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	reqBody := completeQuizReqBody{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&reqBody); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	quizStudent := model.QuizStudent{}
	quizStudent.QuizId = uint(reqBody.QuizId)
	quizStudent.StudentId = uint(reqBody.StudentId)

	if err := db.Save(&quizStudent).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	for _, questionId := range reqBody.QuestionsIds {
		questionStudent := model.QuestionStudent{}
		questionStudent.QuestionId = uint(questionId)
		questionStudent.StudentId = uint(reqBody.StudentId)
		questionStudent.QuizId = uint(reqBody.QuizId)

		if err := db.Save(&questionStudent).Error; err != nil {
			respondError(w, http.StatusInternalServerError, err.Error())
			return
		}

		for _, answerId := range reqBody.AnswersIds {
			answerStudent := model.AnswerStudent{}
			answerStudent.AnswerId = uint(answerId)
			answerStudent.StudentId = uint(reqBody.StudentId)
			answerStudent.QuestionId = uint(questionId)
			answerStudent.QuizId = uint(reqBody.QuizId)

			answer := getAnswerOr404(db, strconv.Itoa(answerId), w, r)
			if answer.IsCorrect {
				answerStudent.IsAnswerCorrect = true
			} else {
				answerStudent.IsAnswerCorrect = false
			}

			if err := db.Save(&answerStudent).Error; err != nil {
				respondError(w, http.StatusInternalServerError, err.Error())
				return
			}
		}
	}

	answersStudents := []model.AnswerStudent{}

	if err := db.Find(&answersStudents, model.AnswerStudent{StudentId: uint(reqBody.StudentId), QuizId: uint(reqBody.QuizId)}).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	correctlyAnsweredQuestions := make([]model.AnswerStudent, 0)
	inCorrectlyAnsweredQuestions := make([]model.AnswerStudent, 0)

	for _, answerStudent := range answersStudents {
		if answerStudent.IsAnswerCorrect {
			correctlyAnsweredQuestions = append(correctlyAnsweredQuestions, answerStudent)
		} else {
			inCorrectlyAnsweredQuestions = append(inCorrectlyAnsweredQuestions, answerStudent)
		}
	}

	response := completeQuizResBody{correctlyAnsweredQuestions, inCorrectlyAnsweredQuestions}

	respondJSON(w, http.StatusOK, response)
}
