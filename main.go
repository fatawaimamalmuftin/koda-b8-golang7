package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type Char struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Status  string `json:"status"`
	Species string `json:"species"`
	Gender  string `json:"gender"`
}

type Resp struct {
	Results []Char `json:"results"`
}

func srcTalent(data []Char, keyword string) []Char {
	var result []Char
	for _, item := range data {
		if strings.Contains(strings.ToLower(item.Name), strings.ToLower(keyword)) {
			result = append(result, item)
		}
	}
	return result
}

func fetch(link string) Resp {
	var data Resp
	res, err := http.Get(link)
	if err != nil {
		fmt.Println(err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println(err)
	}
	return data
}

func main() {

	fmt.Println()
	data := fetch("https://rickandmortyapi.com/api/character")
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Masukan nama : ")
	keyword, _ := reader.ReadString('\n')
	if strings.TrimSpace(keyword) == "0" {
		os.Exit(0)
	}
	keyword = strings.TrimSpace((keyword))

	result := srcTalent(data.Results, keyword)

	if len(result) == 0 {
		fmt.Println("Nama tidak di temukan . . .")
		main()
	}

	fmt.Println("Hasil pencarian : ")
	for _, item := range result {
		fmt.Println("- ", item.Name)
		fmt.Println()
		fmt.Println("0 > untuk exit")
		main()
	}
}
