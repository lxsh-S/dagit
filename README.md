<img width="1920" height="1200" alt="2026_08_26_14_42_15_screenshot" src="https://github.com/user-attachments/assets/24240c79-9196-42dd-92e3-501abefbb3bf" />

# dagit

Visualize your Git repository in your terminal!

`NOTE:- Make sure to have aleast 120x35 cells on the terminal so that dagit renders properly`(You can increase it to as big as you want but I prefer not smaller than the min.(120x35)

# Design choices and goals

dagit was made because of my interest in CLI tools, although it's a TUI :P.

# Overview

- All the data displayed by `dagit` is fetched using git commands on your repository!
- Currently dagit has 4 sections/panels:

- Left panel(Yes i call it this in the code too T_T)
- Visualizer
- Log
- Status

- After fetching the data, we use charm's lipgloss and bubbletea libraries to make this awesome TUI look more beautiful

- We can toggle between the `status and Visualizer` panel using the `v` key on the keyboard

- And the deafault tickrate of updating the data on TUI is 1 second. The user can choose between the following using the `+/-` keys:-
  - 500 milliseconds
  - 1 second
  - 3 seconds
  - 5 seconds
  - 10 seconds

### Left panel

We fetch and display the following data here:-

```
* Repo name 
* The repo's owner name 
* Current branch 
* The url of the remote
```

All of this can be fetched without the internet as we are just running commands and removing thing that we dont need from the output

### Visualizer

We display a git graph of the repo here, showing different branches. Here we fetch the following data

```
* Hash of the commit
* Committer's username
* How long has it been from the commit time 
```

Then we beautify the graph and display it in this section

### Log

We try to give detail about the commits done in this section of datgit. We currently fetch:-

```
* Hash of the commit
* Commit message 
* committer's username 
```

### Status

Here, like `lazygit` we show the staged/modified files on the work tree

# We need a setup?

No, you dont need to setup anything in prior of downloading `dagit`.
 I'm making sure it stays simple. Just make sure you launch it inside a git repo and enjoy!

# Installation

- For arch users/larp maxxers(myslef too) I've pushed it to the AUR:

```
yay -S dagit-bin
```

- using go:

```
go install github.com/lxsh-S/dagit@latest
```

- of course you can download it from the `releases page`

# Contribute

```
git clone https://github.com/lxsh-S/dagit.git
cd dagit
go run main.go 
```
