# ✨ Gitty
> a handy git commit formatter

Helps you to format commit message according to [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/)
Conventional Commit 규칙에 맞도록 커밋 메시지를 작성할 수 있도록 돕는 CLI 툴입니다.

```text
// Message format
<type>[optional scope]: <description>
```

### Select Options
```text
"feat":     "기능 추가"
"fix":      "버그 수정"
"refactor": "리팩터링"
"docs":     "문서화"
"ci":       "CI/CD"
"perf":     "성능 개선"
"chore":    "잡무"
"test":     "테스트"
"style":    "css 등 스타일 작업"
```

## Installation
```shell
 go install github.com/chaewonkong/gitty@latest
```


 ## Run
All you need to do is just run `gitty` in your terminal after git add.
```shell
gitty
```