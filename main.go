package main

import (
	"bufio"
	"fmt"
	"github.com/c-bata/go-prompt"
	"os"
	"os/exec"
)

func Completer(d prompt.Document) []prompt.Suggest {
	var s []prompt.Suggest
	var CommitTypes = map[string]string{
		"feat":     "기능 추가",
		"fix":      "버그 수정",
		"refactor": "리팩터링",
		"docs":     "문서화",
		"ci":       "CI/CD",
		"perf":     "성능 개선",
		"chore":    "잡무",
		"test":     "테스트",
		"style":    "css 등 스타일 작업",
	}

	for k, v := range CommitTypes {
		s = append(s, prompt.Suggest{Text: k, Description: v})
	}

	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func main() {
	r := bufio.NewReader(os.Stdin)
	fmt.Printf("🛠 Gitty a handy git commit formatter 🛠\n")
	fmt.Printf("Select Commit Type: \n\n\n\n\n\n")
	t := prompt.Input(
		">>> ",
		Completer,
		prompt.OptionCompletionOnDown(),
		prompt.OptionDescriptionBGColor(prompt.LightGray),
		prompt.OptionShowCompletionAtStart(),
		prompt.OptionMaxSuggestion(5),
	)
	fmt.Println("selected: " + t)
	fmt.Printf("Enter commit message: \n%s: ", t)
	originalMsg, _ := r.ReadString('\n')

	msg := fmt.Sprintf("%s: %s", t, originalMsg)
	cmd := exec.Command("git", "commit", "-m", msg)

	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Errorf("error occured: %s", err.Error()))
	}
}
