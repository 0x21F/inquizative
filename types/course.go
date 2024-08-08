package edu

import (
	"context"
	"encoding/json"
	"io"

	"code.sajari.com/docconv"
	"github.com/ollama/ollama/api"
)

type Course struct {
	CourseID   []byte
	CourseName string
	Topics     []*Topic
}

type Topic struct {
	Name      string
	Summaries []*Summary
}

type Summary struct {
	Content  string `json:"content"`
	Filename string
}

func GenerateSummary(client api.Client, filename, courseName string, file io.Reader) (*Summary, error) {
	fileContents, _, err := docconv.ConvertPDF(file)
	if err != nil {
		return nil, err
	}

	req := &api.GenerateRequest{
		Model:  "llama3.1",
		System: "You are an instructor for the course " + courseName + ". Read the following text, summarize it, and respond in Json. The summary should be sent with the name \"content\"",
		Prompt: fileContents,
		Stream: new(bool),
	}

	res := &Summary{
		Filename: filename,
	}

	err = client.Generate(context.TODO(), req, func(resp api.GenerateResponse) error {
		err := json.Unmarshal([]byte(resp.Response), res)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}
