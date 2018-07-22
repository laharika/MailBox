package handlers

import (
    "encoding/json"
    "fmt"
    "io"
    "net/http"

    "mailbox_server/commons"

    "github.com/julienschmidt/httprouter"
)

//ConfigureRoutes Configures All the routes with the httpRouter
func ConfigureRoutes(router *httprouter.Router) error {
    router.POST("/test", TestPost)
    router.POST("/getallmails",GetAllMailsHandler)
    router.POST("/changecategory", ChangeCategoryHandler)
    router.POST("/addreply", AddReplyHandler)
    return nil
}

func ReturnResponse(w *http.ResponseWriter, response interface{}, e error) {
	if e != nil {
		return
	}
	js, err := json.Marshal(response)
	if err != nil {
		http.Error(*w, err.Error(), http.StatusInternalServerError)
		return
	}

	(*w).Header().Set("Content-Type", "application/json")

	(*w).Write(js)
}

func ParseRequest(w *http.ResponseWriter, body io.ReadCloser,
	inputData interface{}) {

	err := json.NewDecoder(body).Decode(inputData)

	fmt.Println(json.NewDecoder(body).Decode(inputData))

	if err != nil {
		fmt.Fprintf(*w, "ErrBadRequest - "+err.Error())
		return
	}
}
func CreateResponse(status int32, message string, data interface{}) *commons.Response {
	response := commons.Response{}

	response.Status = status
	response.Message = message
	response.Data = data

	return &response
}

func TestPost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	response := "Test Successful"
	fmt.Println("working")

	ReturnResponse(&w, response, nil)

}
