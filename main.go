package main

import (
	"fmt"
	"os"
	libReadWord "sambav_law/lib"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error while loading env File ")
	}

	input := os.Getenv("INPUTFOLDER")
	chatId := os.Getenv("CHATID")
	botID := os.Getenv("BOTTOKEN")
	//Fetching list of doc file in the input folder
	fileNames, err := libReadWord.FetchDocFilePath(input)
	if err != nil || len(fileNames) <= 0 {
		fmt.Println("error While reading filenames in input folder")
		os.Exit(1)
	}

	for _, fileName := range fileNames {
		filepath := input + "/" + fileName
		Quiz, err := ProcessFile(filepath)
		if err != nil {
			fmt.Println("Error while parsing File ")
			continue
		}
		SendTelegramQuiz(Quiz, botID, chatId)
	}

	fmt.Println("Telgram Poster Program Completed")
}

func ProcessFile(filePath string) (Quiz libReadWord.ParaQuiz, err error) {
	return Quiz, err
}

func SendTelegramQuiz(Quiz libReadWord.ParaQuiz, botToken string, chatId string) (response string, err error) {
	return response, err
}
