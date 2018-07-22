package controller

import (
    "fmt"
    "mailbox_server/commons"
    "encoding/json"
    "io/ioutil"
    "os"
    "strconv"

)

func GetAllMails() ([]commons.Mail, error) {

    arr := make([]commons.Mail,0)
    jsonFile, err := os.Open("../files/data.json")
    if err != nil {
        fmt.Println(err)
        return nil,err
    }
    fmt.Println("Opened Successfully !!")

    defer jsonFile.Close();

    byteValue, _ := ioutil.ReadAll(jsonFile)
    json.Unmarshal(byteValue, &arr)
    fmt.Println(arr)
    for i := 0; i < len(arr); i++ {

        fmt.Println("ID: " + strconv.Itoa(arr[i].Id))
        fmt.Println("from:" + arr[i].From)
        fmt.Println("subject:" + arr[i].Subject)
    }

    return arr,nil
}


func ChangeCategory(Id int, Mail_type string) error {
    mails := make([]commons.Mail,0)
    fmt.Println("Id: ", Id)
    fmt.Println("Mail_type: ", Mail_type)
    jsonFile, err := os.Open("../files/data.json")
    if err != nil {
        fmt.Println(err)
        return err
    }
    fmt.Println("Opened Successfully !!")
    byteValue, _ := ioutil.ReadAll(jsonFile)
    json.Unmarshal(byteValue, &mails)
    jsonFile.Close();

    for i := 0; i < len(mails); i++ {
        if mails[i].Id == Id {
            mails[i].Mail_type = Mail_type
        }
    }

    output, err := json.MarshalIndent(&mails, "", "\t\t")
    if err != nil {
        fmt.Println("Error Marshaling to JSON: ", err)
        return err
    }

    err = ioutil.WriteFile("../files/data.json", output, 06444)
    if(err != nil) {
        fmt.Println("Error writing to file:", err)
        return err
    }
    return nil
}

func AddReply(Id int, Body string) error{
    var reply_mail commons.Mail
    mails := make([]commons.Mail,0)
    jsonFile, err := os.Open("../files/data.json")
    if err != nil {
        fmt.Println(err)
        return err
    }
    fmt.Println("Opened Successfully !!")
    byteValue, _ := ioutil.ReadAll(jsonFile)
    json.Unmarshal(byteValue, &mails)
    jsonFile.Close();

    for i:=0;i<len(mails);i++ {
        if mails[i].Id == Id {
            reply_mail.Id = len(mails) + 1
            reply_mail.To = mails[i].From
            reply_mail.From = mails[i].To
            reply_mail.Body = Body
            reply_mail.Subject = mails[i].Subject
            mails[i].Reply_id = reply_mail.Id
            reply_mail.Mail_type = mails[i].Mail_type
        }
    }
    mails = append(mails, reply_mail)

    output, err := json.MarshalIndent(&mails, "", "\t\t")
    if err != nil {
        fmt.Println("Error Marshaling to JSON: ", err)
        return err
    }

    err = ioutil.WriteFile("../files/data.json", output, 06444)
    if(err != nil) {
        fmt.Println("Error writing to file:", err)
        return err
    }
    return nil
}

/*
func GetMailById(Id int) (common.Mail) {
    mails := make([]commons.Mail,0)
    jsonFile, err := os.Open("../files/data.json")
    if err != nil {
        fmt.Println(err)
        return err
    }
    fmt.Println("Opened Successfully !!")
    byteValue, _ := ioutil.ReadAll(jsonFile)
    json.Unmarshal(byteValue, &mails)
    jsonFile.Close();

    new_body := ""

    for i := 0;i < len(mails); i++ {
        if mails[i].Id == Id {
            new_body = new_body + mails[i].Body
            Id = mails[i].Id
        }

    }
}
*/
