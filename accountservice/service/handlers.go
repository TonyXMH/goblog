package service

import (
	"encoding/json"
	"github.com/TonyXMH/goblog/dbclient"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

var DBClient dbclient.IBoltClient

func GetAccount(w http.ResponseWriter, r *http.Request)  {
	accountID:=mux.Vars(r)["accountID"]
	account,err:=DBClient.QueryAccount(accountID)
	if err!=nil{
		w.WriteHeader(http.StatusNotFound)
		return
	}
	data,_:=json.Marshal(&account)
	w.Header().Set("Content-Type","application/json")
	w.Header().Set("Content-Length",strconv.Itoa(len(data)))
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}