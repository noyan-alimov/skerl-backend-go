package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/noyan-alimov/skerl-backend/model"
	"gorm.io/gorm"
)

func CreateTeacher(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	teacher := model.Teacher{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&teacher); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&teacher).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(w, http.StatusCreated, teacher)
}

func GetTeacher(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	teacher := getTeacherOr404(db, id, w, r)
	if teacher == nil {
		return
	}

	quizzes := []model.Quiz{}

	if err := db.Find(&quizzes, model.Quiz{TeacherId: teacher.ID}).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	teacher.Quizzes = quizzes

	respondJSON(w, http.StatusOK, teacher)
}

func getTeacherOr404(db *gorm.DB, id string, w http.ResponseWriter, r *http.Request) *model.Teacher {
	teacher := model.Teacher{}

	if err := db.First(&teacher, id).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}

	return &teacher
}
