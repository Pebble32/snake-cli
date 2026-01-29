package game

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Score struct{
	Name string		`json:"name"`
	Score int		`json:"score"`
	Date time.Time	`json:"date"`
}

func NewScore() *Score{
	return &Score{Name: "NoName", Score: 0}	
}

func (sc *Score) SaveScore(name string, score int) {
	err := os.MkdirAll("highscore", 0755)
	if err != nil {
		fmt.Println("Error while creating directory", err)
		return
	}
	allScores := sc.LoadScore()

	newScore := Score{
		Name: name,
		Score: score,
		Date: time.Now(),
	}
	allScores = append(allScores, newScore)

	updatedData, err := json.MarshalIndent(allScores, "", " ")
	if err != nil {
		fmt.Println("Error while saving file", err)
		return
	}

	err = os.WriteFile("highscore/highscore.json", updatedData, 0644)
	if err != nil {
		fmt.Println("Error while saving file", err)
		return
	}
}

func (sc *Score) LoadScore() []Score{
	data, err := os.ReadFile("highscore/highscore.json")
	if err != nil {
		return []Score{}
	}

	var allScores []Score
	err = json.Unmarshal(data, &allScores)
	if err != nil {
		fmt.Println("Error while unmarshiling the data", err)
		return []Score{}
	}
	return allScores
}
