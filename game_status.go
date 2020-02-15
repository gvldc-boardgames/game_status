package main

import (
	"fmt"
	"github.com/neo4j/neo4j-go-driver/neo4j"
	"time"
)

type sevenWondersGame struct {
	gameId, gameType, state, creator, name string
	round, age, maxPlayers                 int64
	createdOn                              time.Time
}

func (g sevenWondersGame) String() string {
	return fmt.Sprintf("%v (%v) created by %v on %v for %v players is currently in state %v, age: %v, round: %v", g.name, g.gameType, g.creator, g.createdOn, g.maxPlayers, g.state, g.age, g.round)
}

func convertToGame(record map[string]interface{}) sevenWondersGame {
	return sevenWondersGame{record["gameId"].(string), record["gameType"].(string), record["state"].(string), record["creator"].(string), record["name"].(string), record["round"].(int64), record["age"].(int64), record["maxPlayers"].(int64), record["createdOn"].(time.Time)}
}

func main() {
	var (
		driver  neo4j.Driver
		session neo4j.Session
		result  neo4j.Result
		err     error
	)
	if driver, err = neo4j.NewDriver("bolt://localhost", neo4j.BasicAuth("neo4j", "BoardGames", "")); err != nil {
		fmt.Printf("Error!", err)
	}
	defer driver.Close()
	if session, err = driver.Session(neo4j.AccessModeRead); err != nil {
		fmt.Printf("Error!", err)
	}
	defer session.Close()
	result, err = session.Run("MATCH (g:Game) RETURN g {.*} LIMIT 1", map[string]interface{}{})
	for result.Next() {
		fmt.Println(convertToGame(result.Record().GetByIndex(0).(map[string]interface{})))
	}
}
