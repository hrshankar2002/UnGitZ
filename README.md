
# UnGitZ

![made-with-python](https://img.shields.io/badge/Made%20With-Go-blue)
[![Go Reference](https://pkg.go.dev/badge/github.com/celestiaorg/celestia-node.svg)](https://pkg.go.dev/github.com/celestiaorg/celestia-node)
[![GitHub release (latest by date including pre-releases)](https://img.shields.io/github/v/release/hrshankar2002/ungitz)](https://github.com/hrshankar2002/ungitz/releases/tag/v1.2)

Ungitz is a CLI application written in Golang, primarily for simplifying git workflow for developers by automating the tasks of downloading, unzipping and opening a github repository within the user preferred IDE.

## :wrench: Installation-Brew

```
$ brew tap hrshankar2002/hrshankar2002
```

```
$brew install hrshankar2002/hrshankar2002/ungitz
```

## 📖 Demo

## To unzip and open in VSCode/Intellij Idea/Pycharm

```
$ungitz <code>/<idea>/<charm> -f <filename>,<repo name>,<branch name>
```

#### Examples

- `ungitz code -f demo.zip,testname1,main`

- `ungitz idea -f SpringApplication.zip,testname2,master` 

- `ungitz charm -f demo.zip,testname3,master`

## To download, unzip and open in VSCode/Intellij Idea/Pycharm

```
$ungitz <code>/<idea>/<charm> -l <URL>
```

#### Examples

- `ungitz code -l https://github.com/tensorflow/tensorflow/archive/refs/heads/master.zip`

- `ungitz idea -l https://github.com/tensorflow/tensorflow/archive/refs/heads/master.zip`

- `ungitz charm -l https://github.com/tensorflow/tensorflow/archive/refs/heads/master.zip`

## 🚀 Technologies

  - `Go v1.22.6`
  - `Cobra v1.7.0`

## 🤝 Contribute

To contribute, fork the repository and push the changes to the **master** branch. Then submit a pull request for merging with the source. If your code passes the review and checks it will be merged into the master branch.

## 💬 Feedback

Feel free to send us feedback  or file an issue. Feature requests are always welcome.

  
