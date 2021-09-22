package api

import (
	"encoding/json"
	"github.com/causality/penbox/pkg/payloads/flask"
	"net/http"
)

type Session struct {
	Secret  string            `json:"secret"`
	Session string            `json:"session"`
	Data    map[string]string `json:"data"`
}

// Analysis of session or forging a session according to Secret
// Already mapped to /payloads/flask/session
func flaskSession(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/json")

	if r.Method == "GET" {
		// return help message
		type Message struct {
			Info string                 `json:"info"`
			Help map[string]interface{} `json:"help"`
		}
		m, _ := json.Marshal(Message{
			Info: "Error",
			Help: helpInfo(),
		})
		w.Write(m)
		return
	}

	r.ParseForm()
	action := r.FormValue("action")

	if action == "parse" {
		// get for data
		secret := r.FormValue("secret")
		session := r.FormValue("session")

		// parse session
		data, _ := flask.ParseSession(secret, session)
		jsonResult, _ := json.Marshal(Session{
			secret,
			session,
			data,
		})

		// return
		w.Write(jsonResult)
		return
	}

	if action == "forge" {
		// get form data
		secret := r.FormValue("secret")
		jsonData := r.FormValue("data")
		var data map[string]string
		json.Unmarshal([]byte(jsonData), &data)

		// forge session
		session, _ := flask.ForgeSession(secret, data)
		jsonResult, _ := json.Marshal(Session{
			secret,
			session,
			data,
		})

		// return
		w.Write(jsonResult)
		return
	}

	errInfo, _ := json.Marshal(map[string]string{"info": "make sure post correct action, get this page for more help"})
	w.Write(errInfo)
	return
}

func helpInfo() map[string]interface{} {
	var help = make(map[string]interface{})
	help["parse session"] = map[string]interface{}{
		"method": "post",
		"data": map[string]string{
			"action":  "parse",
			"secret":  "secret_key",
			"session": "session.to.parse",
		},
	}
	help["forge session"] = map[string]interface{}{
		"method": "post",
		"data": map[string]string{
			"action": "forge",
			"secret": "secret_key",
			"data":   "{key1: value1, key2: value2}",
		},
	}
	return help
}
