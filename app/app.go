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
	a.Post("/teachers", a.CreateTeacher)
	a.Get("/teachers/{id}", a.GetTeacher)
	a.Post("/quizzes", a.CreateQuiz)
	a.Get("/quizzes/{id}", a.GetQuiz)
	a.Get("/quizzes", a.GetAllQuizzes)
	a.Put("/quizzes/{id}", a.UpdateQuiz)
	a.Delete("/quizzes/{id}", a.DeleteQuiz)
	a.Post("/questions", a.CreateQuestion)
	a.Get("/questions", a.GetQuestionsByQuizId)
	a.Get("/questions/{id}", a.GetQuestion)
	a.Put("/questions/{id}", a.UpdateQuestion)
	a.Delete("/questions/{id}", a.DeleteQuestion)
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

func (a *App) CreateTeacher(w http.ResponseWriter, r *http.Request) {
	handler.CreateTeacher(a.DB, w, r)
}

func (a *App) GetTeacher(w http.ResponseWriter, r *http.Request) {
	handler.GetTeacher(a.DB, w, r)
}

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
