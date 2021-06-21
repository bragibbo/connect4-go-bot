package main

import (
	"fmt"
	"time"
	"os"
)

const URL string = ""
const GAME_NAME = "connect4"
const COLUMN_SIZE = 7
const ROW_SIZE = 6


func main() {
	args := os.Args
	gameId := ""
	botMode := false
	if len(args) != 3 {
		gameId = CreateGame().Game.Game_id
	} 
	
	if len(args) == 2 { 
		botMode = args[1] != ""
	} else if len(args) == 3 {
		gameId = args[2]
	}
	fmt.Println("Game id:\t" + gameId)


	bot := botJoinGame(gameId, "bragibbo")
	var player2 *PlayerObject
	if !botMode {
		player2 = botJoinGame(gameId, "player2")
	}

	curGame := GetCurrentGame(gameId)
	for curGame.Game.State != "Finish" && curGame.Game.State != "Tie" { 
		if curGame.Game.State != "Waiting" {
			ClearConsole()
			PrettyPrintArrays(curGame.Game.Board)
		}

		if curGame.Game.Player_turn == bot.User_name {
			handleBotMove(curGame.Game.Game_id, bot.Token, curGame.Game.Board, curGame.Game.Player_turn)
		} else if !botMode && curGame.Game.Player_turn == player2.User_name {
			handleBotMove(curGame.Game.Game_id, player2.Token, curGame.Game.Board, curGame.Game.Player_turn)
		}

		time.Sleep(250 * time.Millisecond)
		curGame = GetCurrentGame(gameId) 
	}

	fmt.Println("\n" + curGame.Game.Winner + " Won!")
	fmt.Println("Final Board:")
	PrettyPrintArrays(curGame.Game.Board)
}

func handleBotMove(gameId string, token string, board [ROW_SIZE][COLUMN_SIZE]int, playerName string) {
	move := botMakeMove(board)
	fmt.Println("\t" + playerName + " turn")
	fmt.Printf("\tMove: %d\n", move)
	UpdateGame(gameId, token, move)
}