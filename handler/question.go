package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/noyan-alimov/skerl-backend/model"
	"gorm.io/gorm"
)

func CreateQuestion(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	question := model.Question{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&question); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&question).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(w, http.StatusCreated, question)
}

func GetQuestionsByQuizId(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	questions := []model.Question{}

	quizIdString := r.FormValue("quizId")

	quizIdInt, err := strconv.Atoi(quizIdString)
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := db.Find(&questions, model.Question{QuizId: uint(quizIdInt)}).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(w, http.StatusOK, questions)
}

func GetQuestion(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	question := getQuestionOr404(db, id, w, r)
	if question == nil {
		return
	}

	answers := []model.Answer{}

	// query answers for the question
	if err := db.Find(&answers, model.Answer{QuestionId: question.ID}).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	question.Answers = answers

	respondJSON(w, http.StatusOK, question)
}

func UpdateQuestion(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	question := getQuestionOr404(db, id, w, r)
	if question == nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&question); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&question).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(w, http.StatusOK, question)
}

func DeleteQuestion(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	question := getQuestionOr404(db, id, w, r)
	if question == nil {
		return
	}

	if err := db.Delete(&question).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(w, http.StatusOK, nil)
}

func getQuestionOr404(db *gorm.DB, id string, w http.ResponseWriter, r *http.Request) *model.Question {
	question := model.Question{}

	if err := db.First(&question, id).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}

	return &question
}
