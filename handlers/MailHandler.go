package handlers

import (
	"fmt"
	"net/http"


	"mailbox_server/commons"
	"mailbox_server/controller"
	"github.com/julienschmidt/httprouter"
)

func GetAllMailsHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

    mails, e := controller.GetAllMails()
    if e != nil {
        //Return Errors
        response := CreateResponse(commons.FAIL, e.Error(), nil)
        ReturnResponse(&w, *response, nil)
        return
    }
    response := CreateResponse(commons.SUCCESS, "", mails)
    ReturnResponse(&w, *response, nil)
    return
}

func ChangeCategoryHandler (w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    fmt.Println("Change Category API request")
    request := commons.ChangeCategoryRequest{}
    ParseRequest(&w, r.Body, &request)
    fmt.Println("change request: " , request )
    err := controller.ChangeCategory(request.Id, request.Mail_type)
    if err != nil {
        response := CreateResponse(commons.FAIL, err.Error(), nil)
        ReturnResponse(&w, *response, nil)
        return
    }
    response := CreateResponse(commons.SUCCESS, "", nil)
    ReturnResponse(&w, *response, nil)
    return
}

func AddReplyHandler (w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    fmt.Println("Reply API request")
    request := commons.ReplyRequest{}
    ParseRequest(&w, r.Body, &request)
    err := controller.AddReply(request.Id, request.Body)
    if err != nil {
     response := CreateResponse(commons.FAIL, err.Error(), nil)
     ReturnResponse(&w, *response, nil)
     return
    }
    response := CreateResponse(commons.SUCCESS, "", nil)
    ReturnResponse(&w, *response, nil)
    return
}