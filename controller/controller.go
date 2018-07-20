package controller

import (
    "fmt"
    "mailbox_server/commons"
    "encoding/json"
    "io/ioutil"
    "os"
    "strconv"

)

func GetAllMails() {
    fmt.Println("helooooo!!!")

    jsonFile, err := os.Open("../files/data.json")
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println("Opened Successfully !!")

    defer jsonFile.Close();

    byteValue, _ := ioutil.ReadAll(jsonFile)
    var mails commons.Mails
    json.Unmarshal(byteValue, &mails)

    for i := 0; i < len(mails.Mail_details); i++ {

        fmt.Println("ID: " + strconv.Itoa(mails.Mail_details[i].Id))
        fmt.Println("from:" + mails.Mail_details[i].From)
        fmt.Println("subject:" + mails.Mail_details[i].Subject)
    }
}