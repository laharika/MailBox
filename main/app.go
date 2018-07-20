package main

import(
    "fmt"
    "mailbox_server/controller"
)
func main() {
	fmt.Printf("Hello, world.\n")
    controller.GetAllMails()
}