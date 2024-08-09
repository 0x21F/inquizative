package types

import (
	"time"

	"github.com/ollama/ollama/api"
)

type Quiz struct {
	Questions    []Question `json:"quiz"`
	CreationDate time.Time
	UserID       []byte
	QuizID       []byte
}

type Question interface {
	score(api.Client, string) float32
}

// LLM will grade this
type FRQuestion struct {
	QuestionID []byte
	Index      int
	Question   string `json:"question"`
	Content    string `json:"content"`
	Score      float32
	Fb         Feedback `json:"feedback"`
}

// only one correct answer for this one
// if you're looking for MCQ with multiple
// answers see SelectionQuestion
type MCQuestion struct {
	QuestionID []byte
	Index      int
	Question   string     `json:"question"`
	Options    []string   `json:"options"`
	Fb         []Feedback `json:"feedback"`
	Score      float32
}

type Feedback struct {
	Correct bool `json:"correct"`

	Explanation string `json:"explanation"`
}

func (quiz Quiz) Grade(llm api.Client, content string) float32 {
	var sum float32 = 0.0

	for _, question := range quiz.Questions {
		sum += question.score(llm, content)
	}

	return sum / float32(len(quiz.Questions))
}

func (q FRQuestion) score() float32 {
	var res float32 = 0.0
	courseName := "Calculus"
	req := &api.GenerateRequest{
		Model:  "llama3.1",
		System: "You are a teaching assistant for " + courseName + " grading a free response question. The supplied text will be graded on a scale of 0.0 to 10.0 and given feedback ",
		Prompt: "",
		Stream: new(bool),
	}

	return res
}
