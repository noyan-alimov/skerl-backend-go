package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/noyan-alimov/skerl-backend/config"
	"github.com/noyan-alimov/skerl-backend/handler"
	"github.com/noyan-alimov/skerl-backend/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

func (a *App) Initialize(config *config.Config) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		config.DB.Host,
		config.DB.Username,
		config.DB.Password,
		config.DB.Name,
		config.DB.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Could not connect to database")
	}

	a.DB = model.DBMigrate(db)
	a.Router = mux.NewRouter()
	a.setRouters()
}

func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}

func (a *App) setRouters() {
	// Teacher
	a.Post("/teachers", a.CreateTeacher)
	a.Get("/teachers/{id}", a.GetTeacher)

	// Quiz
	a.Post("/quizzes", a.CreateQuiz)
	a.Get("/quizzes/{id}", a.GetQuiz)
	a.Get("/quizzes", a.GetAllQuizzes)
	a.Put("/quizzes/{id}", a.UpdateQuiz)
	a.Delete("/quizzes/{id}", a.DeleteQuiz)

	// Question
	a.Post("/questions", a.CreateQuestion)
	a.Get("/questions", a.GetQuestionsByQuizId)
	a.Get("/questions/{id}", a.GetQuestion)
	a.Put("/questions/{id}", a.UpdateQuestion)
	a.Delete("/questions/{id}", a.DeleteQuestion)

	// Answer
	a.Post("/answers", a.CreateAnswer)
	a.Get("/answers", a.GetAnswersByQuestionId)
	a.Put("/answers/{id}", a.UpdateAnswer)
	a.Delete("/answers/{id}", a.DeleteAnswer)

	// Student
	a.Post("/students", a.CreateStudent)
	a.Get("/students/{id}", a.GetStudent)

	// Complete Quiz
	a.Put("/completeQuiz", a.CompleteQuiz)

	// Quizzes Students
	a.Get("/quizzesStudents", a.GetQuizzesAndStudents)
}

// Wrap the router for GET method
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

// Wrap the router for POST method
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

// Wrap the router for PUT method
func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("PUT")
}

// Wrap the router for DELETE method
func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("DELETE")
}

// Teacher
func (a *App) CreateTeacher(w http.ResponseWriter, r *http.Request) {
	handler.CreateTeacher(a.DB, w, r)
}

func (a *App) GetTeacher(w http.ResponseWriter, r *http.Request) {
	handler.GetTeacher(a.DB, w, r)
}

// Quiz
func (a *App) CreateQuiz(w http.ResponseWriter, r *http.Request) {
	handler.CreateQuiz(a.DB, w, r)
}

func (a *App) GetQuiz(w http.ResponseWriter, r *http.Request) {
	handler.GetQuiz(a.DB, w, r)
}

func (a *App) GetAllQuizzes(w http.ResponseWriter, r *http.Request) {
	handler.GetAllQuizzes(a.DB, w, r)
}

func (a *App) UpdateQuiz(w http.ResponseWriter, r *http.Request) {
	handler.UpdateQuiz(a.DB, w, r)
}

func (a *App) DeleteQuiz(w http.ResponseWriter, r *http.Request) {
	handler.DeleteQuiz(a.DB, w, r)
}

// Question
func (a *App) CreateQuestion(w http.ResponseWriter, r *http.Request) {
	handler.CreateQuestion(a.DB, w, r)
}

func (a *App) GetQuestionsByQuizId(w http.ResponseWriter, r *http.Request) {
	handler.GetQuestionsByQuizId(a.DB, w, r)
}

func (a *App) GetQuestion(w http.ResponseWriter, r *http.Request) {
	handler.GetQuestion(a.DB, w, r)
}

func (a *App) UpdateQuestion(w http.ResponseWriter, r *http.Request) {
	handler.UpdateQuestion(a.DB, w, r)
}

func (a *App) DeleteQuestion(w http.ResponseWriter, r *http.Request) {
	handler.DeleteQuestion(a.DB, w, r)
}

// Answer
func (a *App) CreateAnswer(w http.ResponseWriter, r *http.Request) {
	handler.CreateAnswer(a.DB, w, r)
}

func (a *App) GetAnswersByQuestionId(w http.ResponseWriter, r *http.Request) {
	handler.GetAnswersByQuestionId(a.DB, w, r)
}

func (a *App) UpdateAnswer(w http.ResponseWriter, r *http.Request) {
	handler.UpdateAnswer(a.DB, w, r)
}

func (a *App) DeleteAnswer(w http.ResponseWriter, r *http.Request) {
	handler.DeleteAnswer(a.DB, w, r)
}

// Student
func (a *App) CreateStudent(w http.ResponseWriter, r *http.Request) {
	handler.CreateStudent(a.DB, w, r)
}

func (a *App) GetStudent(w http.ResponseWriter, r *http.Request) {
	handler.GetStudent(a.DB, w, r)
}

// Complete Quiz
func (a *App) CompleteQuiz(w http.ResponseWriter, r *http.Request) {
	handler.CompleteQuiz(a.DB, w, r)
}

// Quizzes Students
func (a *App) GetQuizzesAndStudents(w http.ResponseWriter, r *http.Request) {
	handler.GetQuizzesAndStudents(a.DB, w, r)
}
