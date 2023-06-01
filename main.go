package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Response struct {
	Choices []Choice `json:"choices"`
}

type Choice struct {
	Text     string  `json:"text"`
	Logprobs Logprobs `json:"logprobs"`
}

type Logprobs struct {
	Tokens []string  `json:"tokens"`
	Top    []float64 `json:"top_logprobs"`
}

func main() {
	apiKey := "<your-api-key>"
	model := "text-davinci-002"
	prompt := "Write a program to collect data using OPENAI GPT API."
	parameters := map[string]interface{}{
		"prompt":      prompt,
		"temperature": 0.5,
		"max_tokens":  100,
		"n":           1,
		"stop":        "\n",
	}

	jsonValue, _ := json.Marshal(parameters)
	request, err := http.NewRequest("POST", fmt.Sprintf("https://api.openai.com/v1/models/%s/completions", model), bytes.NewBuffer(jsonValue))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var data Response
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var output []string
	for _, choice := range data.Choices {
		output = append(output, choice.Text)
	}

	file, _ := json.MarshalIndent(output, "", " ")
	_ = ioutil.WriteFile("output.json", file, os.ModePerm)
}
