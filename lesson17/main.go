package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

func main() {
	// Permissions -- Ruxsatlar (Read, Write, ReadAndWrite, Delete, Execute)
	// Roles -- Rollar (Admin, User, Developer, Moderator)

	// os - paket // Create, Read, Write, Remove, Execute, Stat, Open, Close
	// io - paket // Read, Write, Copy, ReadAll

	// file, err := os.Create("text.txt")
	// if err != nil {
	// 	log.Fatalf("Error creating file: %v", err)
	// }
	// | -> OR
	// Perimissions: 0644 -> 6 = 4 + 2 = Read + Write, 4 = Read, 4 = Read
	// 4 - Read
	// 2 - Write
	// 1 - Execute
	// 0 - No permission
	file, err := os.OpenFile("text.txt", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		WriteErrorText(err.Error())
		return
	}
	defer file.Close()
	// file2, err := os.Open("text2.txt")
	// if err != nil {
	// 	WriteErrorText(err.Error())
	// 	return
	// }

	_, err = file.WriteString("Salom Dunyo\n")
	if err != nil {
		WriteErrorText(err.Error())
		return
	}

	info, err := os.ReadFile("text.txt")
	if err != nil {
		WriteErrorText(err.Error())
		return
	}
	fmt.Println("-->> ", string(info))

	// check file exists or not
	// exist, err := os.Stat("text.txt")
	// if err != nil {
	// 	WriteErrorText(err.Error())
	// 	return
	// }

	text, err := io.ReadAll(file)
	if err != nil {
		WriteErrorText(err.Error())
		return
	}
	// io.Copy(file, file2)
	fmt.Println("-->> ", string(text))

	files, err := os.ReadDir("images")
	if err != nil {
		WriteErrorText(err.Error())
		return
	}
	fmt.Println("Files in images directory:")
	for _, f := range files {
		fmt.Println(f.Name())
	}

	filepath.Walk()

	filepath.IsAbs("text.txt")

	
}

func WriteErrorText(txt string) {
	file, err := os.OpenFile("errors.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		log.Fatalf("Error creating error log file: %v", err)
	}
	defer file.Close()
	_, err = file.WriteString(txt + "\n")
	if err != nil {
		log.Fatalf("Error writing to error log file: %v", err)
	}
}
