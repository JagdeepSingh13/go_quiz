package controller

import (
	"github.com/JagdeepSingh13/go_quiz/internal/service"
	"github.com/gofiber/fiber/v2"
)

type QuizController struct {
	quizService *service.QuizService
}

func Quiz(quizService *service.QuizService) QuizController {
	return QuizController{
		quizService: quizService,
	}
}

// controller calls service service calls collection

func (c QuizController) GetQuizzes(ctx *fiber.Ctx) error {
	quizzes, err := c.quizService.GetQuizzes()
	if err != nil {
		return err
	}

	return ctx.JSON(quizzes)
}
