// Contains the handlers for the receiver entity
package handlers

import "net/http"

func GetReceiver(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get Receiver"))
}

func GetAllReceivers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get Receivers"))
}

func CreateReceiver(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create Receiver"))
}

func UpdateReceiver(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update Receiver"))
}

func DeleteReceiver(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete Receiver"))
}
