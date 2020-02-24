package api

import (
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"jwtauth/auth"
	"jwtauth/json"
	js "encoding/json"
	"log"
	"net/http"
	"time"
)

type AuthHandler interface {
	LoginPost(w http.ResponseWriter, r *http.Request)
	RegisterPost(w http.ResponseWriter, r *http.Request)
}

type handler struct{
	authService auth.Service
}

func (h *handler) LoginPost(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil{
		http.Error(w, http.StatusText(http.StatusInternalServerError),http.StatusInternalServerError)
		return
	}
	userRequested , err := h.serializer(contentType).Decode(requestBody)
	if err != nil{
		http.Error(w, http.StatusText(http.StatusInternalServerError),http.StatusInternalServerError)
		return
	}
	user, err  := h.authService.Login(userRequested.Username)
	if err != nil{
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(userRequested.Password))
	log.Println(err)
	if err != nil{
		type Err struct {
			Error string `json:"error"`
		}
		errResponse, _ := js.Marshal(&Err{Error:"password doesn't match!"})
		setupResponse(w, contentType, errResponse, http.StatusCreated)
		return
	}
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = user.Username
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return
	}
	responseBody, _ := js.Marshal(map[string]string{"token":t})
	setupResponse(w, contentType,responseBody , http.StatusOK)
}

func (h *handler) RegisterPost(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	log.Println(contentType)
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil{
		http.Error(w, http.StatusText(http.StatusInternalServerError),http.StatusInternalServerError)
		return
	}
	user , err := h.serializer(contentType).Decode(requestBody)
	if err != nil{
		http.Error(w, http.StatusText(http.StatusInternalServerError),http.StatusInternalServerError)
		return
	}
	start := time.Now()
	err = h.authService.Register(user)
	elapsed := time.Since(start)
	log.Println(elapsed)
	if err != nil{
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	responseBody, err := h.serializer(contentType).Encode(user)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	setupResponse(w, contentType, responseBody, http.StatusCreated)
}

func NewHandler(service auth.Service) AuthHandler{
	return &handler{authService:service}
}

func(h *handler) serializer(contentType string) auth.Serializer {
	return &json.User{}
}

func setupResponse(w http.ResponseWriter, contentType string, body []byte, statusCode int) {
	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(statusCode)
	_, err := w.Write(body)
	if err != nil {
		log.Println(err)
	}
}

