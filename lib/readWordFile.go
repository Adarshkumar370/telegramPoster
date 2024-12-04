package libReadWord

import (
	"fmt"
	"os"
	"path/filepath"
)

func FetchDocFilePath(folder string) (filePaths []string, err error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
			filePaths = nil
		}
	}()
	// fmt.Println("Reading Files in Folder:  " + folder)

	//Reading Entries in Input Folder
	dirEntries, err := os.ReadDir(folder)
	if err != nil {
		fmt.Printf("Error reading directory: %v\n", err)
		return filePaths, err
	}

	//Fetching all Word document from the file list
	for _, file := range dirEntries {
		ext := filepath.Ext(file.Name())
		if ext == ".doc" || ext == ".docx" {
			filePaths = append(filePaths, file.Name())
		}
	}
	// fmt.Println("File Names:", filePaths)
	return filePaths, err
}

func ParseDocument(filepath string) (docText string, err error) {
	return docText, err
}

func ParseQuiz(input string) (ques ParaQuiz, err error) {
	return ques, err
}
