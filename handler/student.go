package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/noyan-alimov/skerl-backend/model"
	"gorm.io/gorm"
)

func CreateStudent(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	student := model.Student{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&student); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&student).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(w, http.StatusCreated, student)
}

func GetStudent(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	student := getStudentOr404(db, id, w, r)
	if student == nil {
		return
	}

	respondJSON(w, http.StatusOK, student)
}

func getStudentOr404(db *gorm.DB, id string, w http.ResponseWriter, r *http.Request) *model.Student {
	student := model.Student{}

	if err := db.First(&student, id).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}

	return &student
}
