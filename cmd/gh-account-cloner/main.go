package main

import (
    "fmt"
    "log"
    "os"
    "strings"

    "gh-account-cloner/internal/github"
)

func main() {
    if len(os.Args) < 2 {
        log.Fatal("Usage: gh-account-cloner <GitHub User URL>")
    }

    url := os.Args[1]
    parts := strings.Split(url, "/")
    username := parts[len(parts)-1]

    if username == "" {
        log.Fatal("Invalid GitHub URL")
    }

    repos, err := github.FetchRepos(username)
    if err != nil {
        log.Fatalf("Failed to fetch repos: %v", err)
    }

    fmt.Printf("Found %d repos for %s\n", len(repos), username)

    for _, repo := range repos {
        fmt.Println("Cloning:", repo.CloneURL)
        err := github.CloneRepo(repo.CloneURL, username)
        if err != nil {
            log.Printf("Failed to clone %s: %v", repo.Name, err)
        }
    }
}
