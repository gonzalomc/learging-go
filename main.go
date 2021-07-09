package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"elun.cl/go-tutorial/survey"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		allSurveys, err := survey.GetAllSurvey(context.Background())
		if err != nil {
			log.Fatal("Error al parsear lista de encuestas")
		}
		sv, err := json.Marshal(allSurveys)
		if err != nil {
			log.Fatal("Error .....")
		}
		return c.Type("json", "utf-8").Send(sv)
	})

	app.Get("/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			log.Fatal("Error al convertir parametro")
		}
		surveyDetail, err := survey.GetSurveyByID(context.Background(), id)
		if err != nil {
			log.Fatal("Error")
		}
		sv, err := json.Marshal(surveyDetail)
		if err != nil {
			log.Fatal("Eerrorcico")
		}
		return c.Type("json", "utf-8").Send(sv)
	})

	// Crear una nueva pregunta asociada a la encuesta
	app.Post("/question", func(c *fiber.Ctx) error {
		question := new(survey.Question)
		err := c.BodyParser(question)
		if err != nil {
			fmt.Println("Llegando aca")
		}
		newQ, err := survey.CreateQuestion(context.Background(), *question)
		if err != nil {
			fmt.Println("Error desde servicio")
		}
		newQJson, err := json.Marshal(newQ)
		if err != nil {
			fmt.Println("Error al parsear una encuesta")
		}
		return c.Type("json", "utf-8").Send(newQJson)
	})
	app.Listen(":3000")
}
