package main

import (
  "errors"
  "github.com/brickshot/termcaptcha/internal"
  "github.com/google/uuid"
  "github.com/gorilla/mux"
  "log"
  "net/http"
  "strings"
)

// map UUIDs to words
var uuidWord = make(map[uuid.UUID]string)
// uuidCheckCount tracks number of checks per uuid. We allow 2. One for the client, one for the service
// they are attempting to access. If the first check is no then this is set to 255.
var uuidCheckCount = make(map[uuid.UUID]uint8)

func GetHandler(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  uuid, err := uuid.Parse(vars["uuid"])
  if err != nil {
    http.Error(w, err.Error(), http.StatusBadRequest)
    return
  }
  if uuidWord[uuid] != "" {
    http.Error(w, errors.New("UUID already used").Error(), http.StatusBadRequest)
    return
  }
  word, captcha := internal.OneLine()
  uuidWord[uuid] = word
  w.Write([]byte(captcha))
}

func CheckHandler(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  uuid, err := uuid.Parse(vars["uuid"])
  if err != nil {
    http.Error(w, err.Error(), http.StatusBadRequest)
    return
  }
  word := strings.ToUpper(vars["word"])

  // if this uuid has been checked more than once we return an error
  if uuidCheckCount[uuid] > 0 {
    http.Error(w, errors.New("Too many checks").Error(), http.StatusBadRequest)
    return
  }

  uuidCheckCount[uuid]++

  if uuidWord[uuid] != word {
    w.Write([]byte("NO\n"))
    // all further checks will fail with an error
    uuidCheckCount[uuid] = 255
    return
  }

  w.Write([]byte("OK\n"))
}

func VerifyHandler(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  uuid, err := uuid.Parse(vars["uuid"])
  if err != nil {
    http.Error(w, err.Error(), http.StatusBadRequest)
    return
  }
  if uuidCheckCount[uuid] > 1 {
    http.Error(w, errors.New("Too many checks").Error(), http.StatusBadRequest)
    return
  }

  // no more checks
  uuidCheckCount[uuid]++

  // uuid is OK
  w.Write([]byte("OK\n"))
}

func main() {
  r := mux.NewRouter()
  r.HandleFunc("/get/{uuid}", GetHandler)
  r.HandleFunc("/check/{uuid}", CheckHandler).Queries("word", "{word}")
  r.HandleFunc("/verify/{uuid}", VerifyHandler)

  // Bind to a port and pass our router in
  log.Fatal(http.ListenAndServe(":8000", r))
}
