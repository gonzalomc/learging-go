package survey

import (
	"context"
	"fmt"
	"log"

	"elun.cl/go-tutorial/database"
)

// Obtenemos todas las encuestas disponibles
func GetAllSurvey(ctx context.Context) ([]Survey, error) {
	db, err := database.GetConnection()
	if err != nil {
		log.Fatal("Error al conectarse  a la DB")
	}
	surveys, err := db.Query("SELECT * FROM surveys")
	if err != nil {
		log.Fatal("Error en obetener la información")
	}
	defer surveys.Close()

	var svs []Survey
	for surveys.Next() {
		var s Survey
		surveys.Scan(&s.ID, &s.Name, &s.Status)
		svs = append(svs, s)
	}
	return svs, err
}

// Obtenemos una encuesta filtrada por ID
func GetSurveyByID(ctx context.Context, id int) (Survey, error) {
	db, err := database.GetConnection()
	if err != nil {
		log.Fatal("Error")
	}
	query := `SELECT id, name, status FROM surveys where id = $1`
	var sv Survey
	err = db.QueryRow(query, id).Scan(&sv.ID, &sv.Name, &sv.Status)
	if err != nil {
		fmt.Println(err)
		log.Fatal("Erro al obtener la encuesta")
	}

	return sv, nil
}

// Crear una nueva pregunta
func CreateQuestion(ctx context.Context, question Question) (Question, error) {
	db, err := database.GetConnection()
	if err != nil {
		log.Fatal("Error de con.")
	}
	query := `
		INSERT INTO question (title, survey_id) values ($1, $2)
	`
	_, err = db.Exec(query, question.Title, question.SurveyId)
	if err != nil {
		fmt.Println(err)
	}
	return Question{}, nil
}
