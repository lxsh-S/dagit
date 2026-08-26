<img width="1920" height="1200" alt="2026_08_26_14_42_15_screenshot" src="https://github.com/user-attachments/assets/24240c79-9196-42dd-92e3-501abefbb3bf" />

# dagit
Visualize your Git repository in your terminal!

`NOTE:- Make sure to have aleast 100x 35 cells on the terminal so that dagit renders properly`(You can increase it to as big as you want but I prefer not smaller than the min.(100 x 5)

# Design choices and goals

dagit was made because of my interest in CLI tools, although it's a TUI :P.

# We need a setup?
No, you dont need to setup anything in prior of downloading `dagit`.
 I'm making sure it stays simple. Just make sure you launch it inside a git repo and enjoy!

# Installation
* For arch users/larp maxxers(myslef too) I've pushed it to the AUR:
```
yay -S dagit-bin
```
* using go:
```
go install github.com/lxsh-S/dagit@latest
```
* of course you can download it from the `releases page`

# Contribute
```
git clone https://github.com/lxsh-S/dagit.git
cd dagit
go run main.go 
```
