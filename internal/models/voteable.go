package models

// Voteable ...
type Voteable struct {
	UUID     string `dynamo:"ID,hash"`
	Answers  map[string]string
	Question string
}
