package day_38

import "net/http"

const USERNAME = "joker"
const PASSWORD = "riddle"

func Auth(w http.ResponseWriter, r *http.Request) bool {
    username, password, ok := r.BasicAuth()
    if !ok {
        w.Write([]byte(`Error`))
        return false
    }

    isValid := (username == USERNAME) && (password == PASSWORD)
    if !isValid {
        w.Write([]byte(`Wrong username or password`))
        return false
    }

    return true
}