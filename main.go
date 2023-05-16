package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

type Intent struct {
	Name     string   `json:"name"`
	Keywords []string `json:"keywords"`
	Response string   `json:"response"`
}

type IntentData struct {
	Intents []Intent `json:"intents"`
}

var intents []Intent

func main() {
	loadIntentsFromFile("intents.json")

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Chatbot is ready. Start typing your messages.")

	for {
		fmt.Print("> ")
		message, _ := reader.ReadString('\n')
		message = strings.TrimSpace(message)

		if message == "exit" {
			fmt.Println("Exiting...")
			break
		}

		intent := classifyIntent(message)
		response := handleIntent(intent)

		fmt.Println("Chatbot:", response)
	}
}

func loadIntentsFromFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Failed to open intents file:", err)
	}
	defer file.Close()

	var intentData IntentData
	err = json.NewDecoder(file).Decode(&intentData)
	if err != nil {
		log.Fatal("Failed to parse intents file:", err)
	}

	intents = intentData.Intents
}

func classifyIntent(message string) string {
	for _, intent := range intents {
		for _, keyword := range intent.Keywords {
			if strings.Contains(strings.ToLower(message), strings.ToLower(keyword)) {
				return intent.Name
			}
		}
	}

	return "unknown"
}

func handleIntent(intentName string) string {
	for _, intent := range intents {
		if intent.Name == intentName {
			return intent.Response
		}
	}

	return "Sorry, I couldn't understand your request."
}
