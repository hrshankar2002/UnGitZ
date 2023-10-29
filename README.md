
# UnGitZ

![made-with-python](https://img.shields.io/badge/Made%20With-Go-blue)

Ungitz is a CLI application written in Golang, primarily for simplifying git workflow for developers by automating the tasks of downloading, unzipping and opening a github repository within the user preferred IDE.

## :wrench: Installation-Brew

`brew tap hrshankar2002/hrshankar2002`

`brew install hrshankar2002/hrshankar2002/ungitz`

## 📖 Demo

## To unzip and open in VSCode/Intellij Idea/Pycharm

`ungitz <code>/<idea>/<charm> -f <filename>,<git branch name>`

#### Examples

- `ungitz code -f demo.zip,main`

- `ungitz idea -f SpringApplication.zip,master` 

- `ungitz charm -f demo.zip,main`

## To download, unzip and open in VSCode/Intellij Idea/Pycharm

`ungitz code/idea/charm -l <URL>,<filename>,<git branch name>`

#### Examples

- `ungitz code -l http://www.example.com/index.html,demo.zip,main`

- `ungitz idea -l http://www.example.com/index.html,SpringApplication.zip,master`

- `ungitz charm -l http://www.example.com/index.html,demo.zip,main`

## 🚀 Technologies

  - `Go v1.22.6`
  - `Cobra v1.7.0`

## 🤝 Contribute

To contribute, fork the repository and push the changes to the **master** branch. Then submit a pull request for merging with the source. If your code passes the review and checks it will be merged into the master branch.

## 💬 Feedback

Feel free to send us feedback  or file an issue. Feature requests are always welcome.

  
