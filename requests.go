package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"bytes"
)

func GetCurrentGame(id string) *Response {
	resp, err := http.Get(URL + "/game?id=" + id)
	if err != nil {
		log.Fatalln(err)
		return nil
	}
	defer resp.Body.Close()
	return parseJson(*resp)
}

func CreateGame() *Response {
	reqBody, err := json.Marshal(map[string]string{
		"gameName": "connect4",
	})
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return postToServer(reqBody, "/create")
}

func JoinGame(gameId, playerName string) *Response {
	reqBody, err := json.Marshal(map[string]string{
		"gameId": gameId,
		"playerName": playerName,
	})
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return postToServer(reqBody, "/join")
}

func UpdateGame(gameId string, token string, move int) *Response {
	reqBody, err := json.Marshal(map[string]interface{}{
		"gameId": gameId,
		"token": token,
		"move": move,
	})
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return postToServer(reqBody, "/game")
}

func postToServer(body []byte, endpoint string) *Response {
	resp, err := http.Post(URL + endpoint, "application/json", bytes.NewBuffer(body))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	return parseJson(*resp)
}

func parseJson(resp http.Response) *Response {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
		return nil
	}

	var response Response 
	json.Unmarshal(body, &response)
	return &response
}