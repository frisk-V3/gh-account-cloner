package github

import (
    "os"
    "os/exec"
)

func CloneRepo(cloneURL, baseDir string) error {
    cmd := exec.Command("git", "clone", cloneURL)
    cmd.Dir = baseDir
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    return cmd.Run()
}
