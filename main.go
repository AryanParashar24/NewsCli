package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

const apiKey = "mFt2RvihhKA5jJXVdaoI2c53k0YogiUH"

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter query (e.g. arts, politics, science, etc.): ")
	query, _ := reader.ReadString('\n')
	query = strings.TrimSpace(query)

	url := fmt.Sprintf("https://api.nytimes.com/svc/topstories/v2/%s.json?api-key=%s", query, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	defer resp.Body.Close()

	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		fmt.Println("Error decoding response:", err)
		os.Exit(1)
	}

	articles := result["results"].([]interface{})

	for i, article := range articles {
		articleData := article.(map[string]interface{})
		title := articleData["title"].(string)
		abstract := articleData["abstract"].(string)

		fmt.Printf("%d. Title: %s\n", i+1, title)
		fmt.Println("Abstract:", abstract)
		fmt.Println("-----------------------------")
	}

	fmt.Print("Enter the number of the article you want to read more about (or 0 to exit): ")
	selectedArticle, _ := reader.ReadString('\n')
	selectedArticle = strings.TrimSpace(selectedArticle)

	if selectedArticle == "0" {
		os.Exit(0)
	}

	index := atoi(selectedArticle) - 1
	if index >= 0 && index < len(articles) {
		articleData := articles[index].(map[string]interface{})
		title := articleData["title"].(string)
		abstract := articleData["abstract"].(string)
		url := articleData["url"].(string)

		fmt.Printf("Title: %s\n", title)
		fmt.Println("Abstract:", abstract)
		fmt.Println("URL:", url)
	} else {
		fmt.Println("Invalid selection")
	}
}

func atoi(s string) int {
	var n int
	for _, ch := range s {
		n = n*10 + int(ch-'0')
	}
	return n
}
