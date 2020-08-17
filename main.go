package main

import (
	
	"encoding/json"
	"github.com/gorilla/mux"
	
	"io/ioutil"
	"log"
	"net/http"	
	"strconv"	
)
type IndexResponseType2 struct{
	Body string `json:"message"`
	Label string `json:"token"`
}
type IndexResponse struct{
	Body string `json:"message"`
}
type PostWriteSchema struct {
	Body string `json:"username"`
	pass string `json:"password"`
}

type BuyPostSchema struct{
	Body int `json:"product_id"`
}

func main(){
r := mux.NewRouter()
r.HandleFunc("/helloworld",func(writer http.ResponseWriter, request *http.Request){
	resp := IndexResponse{Body:"Hello world"}
	jsonResp,err:= json.Marshal(resp)
	if err!=nil{
	writer.WriteHeader(200)
	}
se, err  := writer.Write(jsonResp)
if err!=nil{
	log.Print(se)
}

})
r.HandleFunc("/goodbyeworld",func(writer http.ResponseWriter, request *http.Request){
	writer.WriteHeader(404)
})
r.HandleFunc("/signup",func(writer http.ResponseWriter, request *http.Request){
	doc := PostWriteSchema{}
		requestBody, err := ioutil.ReadAll(request.Body)
		err = json.Unmarshal(requestBody,&doc)
		if err != nil {
			return
		}
		
		if err != nil {
			return
		}
		username := doc.Body
		response:="User "
		response= response+username+" signed up successfully."
		res  :=IndexResponseType2{Body:response , Label:"x1xPdDnO69cs8MiyWhaHMxqf8lnIun7J"}
		jsonRes , err := json.Marshal(res)
		if err!=nil{

		}
		writer.WriteHeader(200)
		writer.Write(jsonRes)
	})

r.HandleFunc("/buy_product",func(writer http.ResponseWriter, request *http.Request){
	docc := BuyPostSchema{}
	requestBody, err := ioutil.ReadAll(request.Body)
	err = json.Unmarshal(requestBody,&docc)
	if err != nil {
		return
	}
	
	if err != nil {
		return
	}
	product_id := docc.Body
	result := strconv.Itoa(product_id)
	response:="product with id "
	response= response+result+" purchased successfully."
	log.Print(response)
	res  := IndexResponse{Body : response}
	jsonRes , err := json.Marshal(res)
	if err!=nil{
		return
	}
	writer.WriteHeader(200)
	writer.Write(jsonRes)
})
http.Handle("/",r)
err := http.ListenAndServe(":8080", r) 
if err != nil {
	log.Fatalf("failed to listen")
}
}
