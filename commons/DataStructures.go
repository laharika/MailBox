package commons

type Mail struct {
    Id  int         `json:"id"`
    Subject string  `json:"subject"`
    To string       `json:"to"`
    From string     `json:"from"`
    Mail_type string     `json:"mail_type"`
    Body string     `json:"body"`
    Reply_id int    `json:"reply_id"`
}

type Mails struct {
    Mail_details []Mail `json:"mail_details"`
}

