package handler

import (
	"net/http"
	"strconv"

	"github.com/noyan-alimov/skerl-backend/model"
	"gorm.io/gorm"
)

// This handler helps to find out what quizzes completed a specific student
// Or find out which students completed a specific quiz
func GetQuizzesAndStudents(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	quizzesStudents := []model.QuizStudent{}

	quizIdString := r.FormValue("quizId")
	studentIdString := r.FormValue("studentId")

	if len(quizIdString) > 0 && len(studentIdString) > 0 {
		respondError(w, http.StatusBadRequest, "Please provide only one query parameter")
		return
	}

	if len(quizIdString) == 0 && len(studentIdString) == 0 {
		respondError(w, http.StatusBadRequest, "Please provide a query parameter")
		return
	}

	if len(quizIdString) > 0 {
		quizIdInt, err := strconv.Atoi(quizIdString)
		if err != nil {
			respondError(w, http.StatusBadRequest, err.Error())
			return
		}

		if err := db.Find(&quizzesStudents, model.QuizStudent{QuizId: uint(quizIdInt)}).Error; err != nil {
			respondError(w, http.StatusInternalServerError, err.Error())
			return
		}

		respondJSON(w, http.StatusOK, quizzesStudents)
	}

	if len(studentIdString) > 0 {
		studentIdInt, err := strconv.Atoi(studentIdString)
		if err != nil {
			respondError(w, http.StatusBadRequest, err.Error())
			return
		}

		if err := db.Find(&quizzesStudents, model.QuizStudent{StudentId: uint(studentIdInt)}).Error; err != nil {
			respondError(w, http.StatusInternalServerError, err.Error())
			return
		}

		respondJSON(w, http.StatusOK, quizzesStudents)
	}
}
