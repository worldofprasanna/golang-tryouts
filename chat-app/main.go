package main

import (
  "net/http"
  "log"
  "sync"
  "html/template"
  "path/filepath"
  "flag"
  "chat/trace"
  "os"
  "github.com/stretchr/gomniauth"
  "github.com/stretchr/gomniauth/providers/facebook"
  "github.com/stretchr/objx"
  "fmt"
)

type templateHandler struct {
  once sync.Once
  filename string
  templ *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  t.once.Do(func() {
    t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
  })
  data := map[string]interface{}{
    "Host": r.Host,
  }
  if userName, err := r.Cookie("auth"); err == nil {
    data["User"] = objx.MustFromBase64(userName.Value)
  } else if err != nil {
    fmt.Println("Error occurred in main.go -> "+err.Error())
  }

  t.templ.Execute(w, data)
}

func main() {
  clientId := os.Getenv("CLIENT_ID")
  secretKey := os.Getenv("SECRET_KEY")
  addr := flag.String("addr", "8080", "Port in which application is running")
  flag.Parse()
  gomniauth.SetSecurityKey("test1234")
  gomniauth.WithProviders(
    facebook.New(clientId, secretKey,
      "http://localhost:3000/auth/callback/facebook"),
  )
  r := newRoom()
  r.tracer = trace.New(os.Stdout)
  http.Handle("/chat", MustAuth(&templateHandler{filename: "index.html"}))
  http.Handle("/login", &templateHandler{filename: "login.html"})
  http.Handle("/room", r)
  http.HandleFunc("/auth/", LoginHandler)
  http.HandleFunc("/logout", LogOutHandler)
  go r.run()
  log.Println("Application is running in port ", *addr)
  if err := http.ListenAndServe(*addr, nil); err != nil {
    log.Fatal("ListenAndServe:", err)
  }
}

