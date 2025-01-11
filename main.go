package main

import (
        "fmt"
        "net/http"
        "os"
        "encoding/json"
)

type ApiResponse struct {
        Type string `json:"type"`
        CreatedAt string `json:"created_at"`
        Repo struct {
                Name string `json:"name"`
        } `json:"repo"`
}

func main() {

        var username string

        fmt.Println("The GitHub activity tracker program")
        fmt.Print("Please enter your GitHub username: ")
        fmt.Scan(&username)

        url := fmt.Sprintf("https://api.github.com/users/%s/events", username)

        response, err := http.Get(url)
        if err != nil {
                fmt.Println(err)
                os.Exit(1)
        }

        defer response.Body.Close()

        var results []ApiResponse

        err = json.NewDecoder(response.Body).Decode(&results)

        if err != nil {
                fmt.Println(err)
                os.Exit(1)
        }

        for _, result := range results {
                fmt.Printf(
                        "Type: %s, Created at: %s, Repo: %s\n",
                        result.Type, result.CreatedAt, result.Repo,
                )
        }
}