package main

type GameObject struct {
	Game_id string
	Object_type string
	Game_name string
	Min_players int8
	Max_players int8
	Num_players int8
	Player_turn string
	Players []string
	Winner string
	State string
	Board [ROW_SIZE][COLUMN_SIZE]int
}

type PlayerObject struct {
	Game_id string
	Object_type string
	User_name string
	Game_token int8
	Token string
}

type Response struct {
	Status int16
	Game GameObject
	Player *PlayerObject
	Message string
}