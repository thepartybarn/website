package main

import (
	"github.com/sirupsen/logrus"
	"net/http"
)

var (
	_buildDate    string
	_buildVersion string
	log           = logrus.New()
)

func main() {
	var err error
	log.SetLevel(logrus.TraceLevel)
	log.Printf("---------- Program Started %v (%v) ----------", _buildVersion, _buildDate)

	myMux := &CustomMux{DefaultRoute: homeHandler}

	myMux.Handle("/", homeHandler, "NONE")

	myMux.Handle("/api/login", loginHandler, "NONE")

	//http.HandleFunc("/api/addUser", addUserHandler)
	//http.HandleFunc("/api/addFriend", addFriendHandler)

	log.Trace("Opening HTTP Server")
	err = http.ListenAndServe(":80", myMux)
	if err != nil {
		log.Panic(err)
	}
}
func homeHandler(context *Context) {

}
func loginHandler(context *Context) {

}
