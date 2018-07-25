package commons

const SUCCESS = 0
const FAIL = 1

type Response struct {
	Status  int32       `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Mail struct {
    Id  int         `json:"id"`
    Subject string  `json:"subject"`
    To string       `json:"to"`
    From string     `json:"from"`
    Mail_type string `json:"type"`
    Body string     `json:"body"`
    Reply_id int    `json:"reply_id"`
}

type ChangeCategoryRequest struct {
    Id int              `json:"id"`
    Mail_type string    `json:"mail_type"`
}

type ReplyRequest struct {
    Id int      `json:"id"`
    Body string `json:"Body"`
}

type SearchRequest struct {
    Mail_type string      `json:"mail_type"`
    Search_text string    `json:"search_text"`
}

type GetMailByIDRequest struct {
    Id int `json:"id"`
}