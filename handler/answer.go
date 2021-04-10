package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/noyan-alimov/skerl-backend/model"
	"gorm.io/gorm"
)

func CreateAnswer(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	answer := model.Answer{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&answer); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&answer).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(w, http.StatusCreated, answer)
}

func GetAnswersByQuestionId(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	answers := []model.Answer{}

	questionIdString := r.FormValue("questionId")

	questionIdInt, err := strconv.Atoi(questionIdString)
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := db.Find(&answers, model.Answer{QuestionId: uint(questionIdInt)}).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(w, http.StatusOK, answers)
}

func UpdateAnswer(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	answer := getAnswerOr404(db, id, w, r)
	if answer == nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&answer); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&answer).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(w, http.StatusOK, answer)
}

func DeleteAnswer(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	answer := getAnswerOr404(db, id, w, r)
	if answer == nil {
		return
	}

	if err := db.Delete(&answer).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(w, http.StatusOK, nil)
}

func getAnswerOr404(db *gorm.DB, id string, w http.ResponseWriter, r *http.Request) *model.Answer {
	answer := model.Answer{}

	if err := db.First(&answer, id).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}

	return &answer
}
