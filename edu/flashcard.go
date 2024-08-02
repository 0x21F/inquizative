package edu

import "time"

type FlashCard struct {
	Front           string `json:"Front"`
	Content         string `json:"content"`
	AdditionalNotes string `json:"additional_notes"`
}

type FlashCardDeck struct {
	Deck         []FlashCard `json:"deck"`
	DeckID       []byte
	UserId       []byte
	CreationDate time.Time
}
