package service

import (
	"encoding/json"
	"fmt"
	"github.com/TonyXMH/goblog/accountservice/dbclient"
	"net"

	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

var DBClient dbclient.IBoltClient

func GetAccount(w http.ResponseWriter, r *http.Request) {
	accountID := mux.Vars(r)["accountID"]
	account, err := DBClient.QueryAccount(accountID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	account.ServedBy = getIP()
	data, _ := json.Marshal(&account)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
func getIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "error"
	}
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	panic("Unable to determine local IP address(non loopback).Exiting")
}

type healthCheckResp struct {
	Status string `json:"status"`
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {

	dbUp := DBClient.Check()
	if dbUp && isHealthy {
		data, _ := json.Marshal(&healthCheckResp{
			Status: "UP",
		})
		writeJsonResp(w, http.StatusOK, data)
	} else {
		data, _ := json.Marshal(&healthCheckResp{
			Status: "DB unaccessible",
		})
		writeJsonResp(w, http.StatusServiceUnavailable, data)
	}
}

func writeJsonResp(w http.ResponseWriter, status int, data []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	w.WriteHeader(status)
	w.Write(data)
}

var isHealthy = true

func SetHealthyState(w http.ResponseWriter, r *http.Request) {
	state, err := strconv.ParseBool(mux.Vars(r)["state"])
	if err != nil {
		fmt.Println("Invalid request to SetHealthyState allowed values are true or false")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	isHealthy = state
	w.WriteHeader(http.StatusOK)
}
