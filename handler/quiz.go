package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/noyan-alimov/skerl-backend/model"
	"gorm.io/gorm"
)

func CreateQuiz(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	quiz := model.Quiz{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&quiz); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&quiz).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(w, http.StatusCreated, quiz)
}

// func GetAllQuizzes(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
// 	quiz := model.Quiz{}

// }

func GetQuiz(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	quiz := getQuizOr404(db, id, w, r)
	if quiz == nil {
		return
	}

	respondJSON(w, http.StatusOK, quiz)
}

func getQuizOr404(db *gorm.DB, id string, w http.ResponseWriter, r *http.Request) *model.Quiz {
	quiz := model.Quiz{}

	if err := db.First(&quiz, id).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}

	return &quiz
}
