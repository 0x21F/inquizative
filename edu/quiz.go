package edu

import (
	"time"
)

type Quiz struct {
	Questions    []Question
	CreationDate time.Time
	UserID       []byte
	QuizID       []byte
}

type Question interface {
	score() float32
}

// LLM will grade this
type FRQuestion struct {
	QuestionID []byte
	Index      int
	Question   string
	Content    string
	Score      float32
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
	Correct     bool   `json:"correct"`
	Explanation string `json:"explanation"`
}
