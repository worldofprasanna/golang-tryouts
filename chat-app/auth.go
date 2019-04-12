package main

import (
  "net/http"
  "strings"
  "github.com/stretchr/gomniauth"
  "github.com/stretchr/objx"
)

type authHandler struct {
  next http.Handler
}

func (authHandler *authHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  if _, err := r.Cookie("auth"); err == http.ErrNoCookie {
    w.Header().Set("Location", "/login")
    w.WriteHeader(http.StatusTemporaryRedirect)
  } else if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
  authHandler.next.ServeHTTP(w, r)
}

func MustAuth(httpHandler http.Handler) http.Handler {
  return &authHandler{next:httpHandler}
}

func LogOutHandler(w http.ResponseWriter, r *http.Request) {
  http.SetCookie(w, &http.Cookie{
    Name: "auth",
    Value: "",
    MaxAge: -1,
    Path: "/",
  })
  w.Header().Set("Location", "/chat")
  w.WriteHeader(http.StatusTemporaryRedirect)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
  segs := strings.Split(r.URL.Path, "/")
  action := segs[2]
  provider := segs[3]

  switch action {
  case "login":
    provider, err := gomniauth.Provider(provider)
    if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
    }
    loginUrl, err := provider.GetBeginAuthURL(nil, nil)
    if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
    }
    w.Header().Set("Location", loginUrl)
    w.WriteHeader(http.StatusTemporaryRedirect)
  case "callback":
    provider, err := gomniauth.Provider(provider)
    if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }
    credentials, err := provider.CompleteAuth(objx.MustFromURLQuery(r.URL.RawQuery))
    if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }
    user, err := provider.GetUser(credentials)
    if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }
    cookieValue := objx.New(map[string]interface{}{
      "name": user.Name(),
      "AvatarURL": user.AvatarURL(),
    }).MustBase64()
    http.SetCookie(w, &http.Cookie{
      Name: "auth",
      Value: cookieValue,
      Path: "/",
    })
    w.Header().Set("Location", "/chat")
    w.WriteHeader(http.StatusTemporaryRedirect)
  default:
    w.WriteHeader(http.StatusNotFound)
  }
}
