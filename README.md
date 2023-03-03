# Install
Clone this repo first, obviously. Then...
1. [Install go](https://go.dev/dl/)
2. `brew install sox` <- download SoX
3. `go build -o satansays satansays.go && sudo mv satansays /usr/local/bin/` <- make binary, move to PATH

# Run
#### Slightly Cool Option
4. Enter `satansays` in your terminal and tell satan what to say

<!-- TODO add options for someone using zshell but no .zshrc file yet -->
#### Way Cooler Option (if you run zsh) ⫛
4. Run this command to make a bash function in your .zshrc file and make it ready to use
```
echo -e "\nhail() {\nif [ "$1" = "satan" ]; then\nsatansays\nelse\necho "Invalid command"\nfi\n}\n" >> ~/.zshrc && exec zsh
```

5. Enter `hail satan` into the terminal to speak through satan himself 𓄋
6. Exit with `jesus is lord`


# Uninstall
8. `sudo rm /usr/local/bin/satansays` <- remove binary file. 
9. `brew uninstall sox` <- uninstall SoX
10. `rm -rfi *` <- run in cloned repo to burn it