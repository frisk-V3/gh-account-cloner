package main

import (
    "fmt"
    "strings"

    "gh-account-cloner/internal/github"

    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/dialog"
    "fyne.io/fyne/v2/widget"
)

func main() {
    a := app.New()
    w := a.NewWindow("GitHub Account Cloner (Linux)")

    urlEntry := widget.NewEntry()
    urlEntry.SetPlaceHolder("https://github.com/torvalds")

    selectedDir := ""

    dirLabel := widget.NewLabel("保存先フォルダ: 未選択")

    selectDirBtn := widget.NewButton("フォルダを選択", func() {
        dialog.ShowFolderOpen(func(uri fyne.ListableURI, err error) {
            if uri != nil {
                selectedDir = uri.Path()
                dirLabel.SetText("保存先フォルダ: " + selectedDir)
            }
        }, w)
    })

    logArea := widget.NewMultiLineEntry()
    logArea.SetPlaceHolder("ログがここに表示されます")

    cloneBtn := widget.NewButton("クローン開始", func() {
        if selectedDir == "" {
            logArea.SetText("エラー: 保存先フォルダを選択してください")
            return
        }

        url := urlEntry.Text
        parts := strings.Split(url, "/")
        username := parts[len(parts)-1]

        repos, err := github.FetchRepos(username)
        if err != nil {
            logArea.SetText(fmt.Sprintf("リポジトリ取得失敗: %v", err))
            return
        }

        logArea.SetText(fmt.Sprintf("%d 個のリポジトリを発見\n", len(repos)))

        for _, repo := range repos {
            logArea.SetText(logArea.Text + "Cloning: " + repo.CloneURL + "\n")
            err := github.CloneRepo(repo.CloneURL, selectedDir)
            if err != nil {
                logArea.SetText(logArea.Text + fmt.Sprintf("失敗: %v\n", err))
            }
        }

        logArea.SetText(logArea.Text + "\n完了！")
    })

    w.SetContent(container.NewVBox(
        widget.NewLabel("GitHub ユーザー URL を入力"),
        urlEntry,
        selectDirBtn,
        dirLabel,
        cloneBtn,
        logArea,
    ))

    w.Resize(fyne.NewSize(500, 500))
    w.ShowAndRun()
}
