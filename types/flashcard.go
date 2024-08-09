package types

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"time"

	"github.com/ollama/ollama/api"
)

type FlashCard struct {
	FlashCardID     []byte
	Front           string `json:"front"`
	Content         string `json:"content"`
	AdditionalNotes string `json:"additional_notes,omitempty"`
}

type FlashCardDeck struct {
	Deck         []FlashCard `json:"deck"`
	DeckID       []byte
	UserId       []byte
	CreationDate time.Time
}

func GenerateFlashCardDeck(client api.Client, topic Topic) (*FlashCardDeck, error) {
	res := &FlashCardDeck{}

	// this could hypothetically get really slow cause of string concatenation, ahthough
	// I'm not sure how much of an issue this is in go vs with java.
	// fixed.
	var b bytes.Buffer
	for _, summary := range topic.Summaries {
		b.WriteString(summary.Content)
	}

	topic_contents := b.String()

	req := &api.GenerateRequest{
		Model: "llama3.1",
		System: "You are a student studying " + topic.Name + ". Read the following summaries, and" +
			" create a deck of flashcards. These flashcards have three items, a front which has the " +
			"topic to be discussed, content which is the information about the topic, and an optional" +
			" additional notes are for any important context or conditions of the content. Respond in" +
			" json. In the json response, front will have the name \"front\", content will have the " +
			"name \"content\", and the additional notes will be named \"additional_notes\". The array" +
			" of flashcards will be named \"deck\" ",
		Prompt: topic_contents,
		Stream: new(bool),
	}

	err := client.Generate(context.TODO(), req, func(resp api.GenerateResponse) error {
		err := json.Unmarshal([]byte(resp.Response), res)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	res.CreationDate = time.Now()

	return res, nil
}

func (deck FlashCardDeck) Store(db *sql.DB) error {
	db.Exec("INSERT")

	return nil
}
