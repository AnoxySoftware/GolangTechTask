package models

// Voteable holds questions, answers, and the votes given
type Voteable struct {
	Votes    map[int64]int64
	UUID     string `dynamo:"ID,hash"`
	Question string
	Answers  []string
}
