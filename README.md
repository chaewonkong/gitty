# ✨ Gitty
> a handy git commit formatter

Gitty is a CLI tool built with Go, helps you to format commit message according to [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/)

![ezgif com-video-to-gif-converter (1)](https://github.com/chaewonkong/gitty/assets/33085082/f226403a-4753-423a-8240-b42b66911230)


```text
// Message format
<type>[optional scope]: <description>
```


### Select Options
```text
"feat":     "Adding new features"
"fix":      "Fixing the bugs"
"refactor": "Refactoring codes"
"docs":     "Documentation"
"ci":       "CI/CD improvements"
"perf":     "Performance Enhancement"
"chore":    "Tedious task done"
"test":     "Adding/Editing tests"
"style":    "Style changes (like css)"
```

## Installation
If you are using MacOS, you can install it with brew.

```shell
brew tap chaewonkong/homebrew-gitty
brew install gitty
```

Or if you are familiar to Go, just run the following
```shell
 go install github.com/chaewonkong/gitty@v1.0.5
```


 ## Run
All you need to do is just run `gitty` in your terminal after git add.
```shell
gitty
```
