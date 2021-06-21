package main

import (
	"log"
	"math/rand"
	"time"
)

func botJoinGame(gameId, botName string) *PlayerObject {
	player := JoinGame(gameId, botName).Player
	if player.Game_id == "" {
		log.Fatal(botName + " not created succesfully")
	}
	return player
}

func botMakeMove(board [ROW_SIZE][COLUMN_SIZE]int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(7)
}