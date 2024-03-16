// Package hello handles requestss
package hello

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
	"io/ioutil"
	"github.com/google/uuid"
	"golangbook/structs"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func init() {
	rand.Seed(time.Now().UnixNano())
}

func getip2() string {
    req, err := http.Get("http://ip-api.com/json/")
    if err != nil {
        return err.Error()
    }
    defer req.Body.Close()

    body, err := ioutil.ReadAll(req.Body)
    if err != nil {
        return err.Error()
    }

    var ip structs.IP
    json.Unmarshal(body, &ip)

    return ip.Query
}

func HelloHandler(w http.ResponseWriter, r *http.Request, myUUID uuid.UUID) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Status", "200")
	hTTPResponse := structs.HTTPResponse{
		Status:      200,
		Application: "hello-weather",
		IP: getip2(),
		UUID:        myUUID,
		Data:        randSeq(200000),
	}

	err := json.NewEncoder(w).Encode(hTTPResponse)
	if err != nil {
		fmt.Fprintf(w, "%+v", hTTPResponse)
	}
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Service alive and reachable")
}

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
