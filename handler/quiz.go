package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

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

func GetAllQuizzes(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	quizzes := []model.Quiz{}

	teacherIdString := r.FormValue("teacherId")
	if len(teacherIdString) == 0 {
		if err := db.Find(&quizzes).Error; err != nil {
			respondError(w, http.StatusInternalServerError, err.Error())
			return
		}

		respondJSON(w, http.StatusOK, quizzes)
		return
	}

	teacherIdInt, err := strconv.Atoi(teacherIdString)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := db.Find(&quizzes, model.Quiz{TeacherId: uint(teacherIdInt)}).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(w, http.StatusOK, quizzes)
}

func GetQuiz(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	quiz := getQuizOr404(db, id, w, r)
	if quiz == nil {
		return
	}

	questions := []model.Question{}

	// query questions for the quiz
	if err := db.Find(&questions, model.Question{QuizId: quiz.ID}).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	quiz.Questions = questions

	respondJSON(w, http.StatusOK, quiz)
}

func UpdateQuiz(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	quiz := getQuizOr404(db, id, w, r)
	if quiz == nil {
		return
	}

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

	respondJSON(w, http.StatusOK, quiz)
}

func DeleteQuiz(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	quiz := getQuizOr404(db, id, w, r)
	if quiz == nil {
		return
	}

	if err := db.Delete(&quiz).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(w, http.StatusOK, nil)
}

func getQuizOr404(db *gorm.DB, id string, w http.ResponseWriter, r *http.Request) *model.Quiz {
	quiz := model.Quiz{}

	if err := db.First(&quiz, id).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}

	return &quiz
}
