package main

import (
	"bufio"
	"fmt"
	"github.com/manifoldco/promptui"
	"os"
	"os/exec"
)

var CommitTypesMap = map[string]string{
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

type CommitTypes struct {
	Type        string
	Description string
}

func main() {
	r := bufio.NewReader(os.Stdin)
	fmt.Printf("✨  Gitty a handy git commit formatter 🛠\n\n")

	var commitTypes []CommitTypes
	for k, v := range CommitTypesMap {
		ct := CommitTypes{
			Type:        k,
			Description: v,
		}

		commitTypes = append(commitTypes, ct)
	}

	tmpl := &promptui.SelectTemplates{
		Label:    "Select Commit Type:",
		Active:   " {{ .Type | cyan }}: ({{ .Description | red }}) \U0001F336 ",
		Inactive: " {{ .Type | cyan }}: ({{ .Description | red }})",
		Selected: "\U0001F336 {{ .Type | red | cyan }}",
	}

	prompt := promptui.Select{
		Label:     "Select Commit Type",
		Items:     commitTypes,
		Templates: tmpl,
	}

	i, _, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
	}

	t := commitTypes[i].Type

	fmt.Printf("Enter commit message >>> \n%s: ", t)
	originalMsg, _ := r.ReadString('\n')

	msg := fmt.Sprintf("%s: %s", t, originalMsg)
	cmd := exec.Command("git", "commit", "-m", msg)

	err = cmd.Run()
	if err != nil {
		fmt.Println(fmt.Errorf("error occured: %s", err.Error()))
	}

	fmt.Println("✨ Committed successfully ✨")
}
