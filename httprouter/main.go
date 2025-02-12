package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.GET("/api/v1/go-version", goVersion)
	router.GET("/api/v1/show-file/:name", getFileContent)
	log.Fatal(http.ListenAndServe(":8080", router))
}

// Implement the handler
func getCommandOutput(command string, arguments ...string) string {
	out, _ := exec.Command(command, arguments...).Output()
	return string(out)
}

func goVersion(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	response := getCommandOutput("go", "version")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}

func getFileContent(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	response := getCommandOutput("more", param.ByName("name"))
	log.Println(response)
	fmt.Fprint(w, response)
}
