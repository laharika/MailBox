# MailBox

PLEASE FOLLOW THE STEPS BELOW TO SET UP THE SERVER

1. Create 'MailBox' folder and 'src' and 'pkg' subfolders under it. Paste all the folders in this repository within the "src" folder.
2. Set the 'GOPATH' environment variable to 'MailBox' location.
3.Open command prompt on the 'Mailbox' location and please type the following command to download all dependencies go get ./...
Alternatively, you could manually download the dependencies by firing the following commands
syntax: go get 
    1. github.com/julienschmidt/httprouter
    2. github.com/rs/cors
3. Run the program by navigating to 'main' folder and executing 'go run app.go' command.
