package github

import (
    "encoding/json"
    "fmt"
    "net/http"
)

type Repo struct {
    Name     string `json:"name"`
    CloneURL string `json:"clone_url"`
}

func FetchRepos(username string) ([]Repo, error) {
    url := fmt.Sprintf("https://api.github.com/users/%s/repos", username)

    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    var repos []Repo
    err = json.NewDecoder(resp.Body).Decode(&repos)
    return repos, err
}
