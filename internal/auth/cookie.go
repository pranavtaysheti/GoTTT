package auth

import (
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/pranavtaysheti/GoTTT/internal/server"
)

const ClientCookieName = "client_id"

func GetClientFromRequest(r *http.Request) (*server.Client, error) {
	c, err := getClientCookie(r)
	if err != nil {
		return nil, err 
	}

	cl, err := getClientFromCookie(c)
	if err != nil {
		return nil, err 
	}

	return cl, nil 
}

func getClientCookie(r *http.Request) (*http.Cookie, error) {
	clientCookie, err := r.Cookie(ClientCookieName)
	if err != nil {
		log.Println("Error reading player cookie, no client cookie found")
		return nil, err
	}

	return clientCookie, nil
}

func getClientFromCookie(c *http.Cookie) (*server.Client, error) {
	cl, err := server.GetClientByUUIDString(c.Value)
	if err != nil {
		log.Println("Client not found, corresponding to cookie")
		return nil, err
	}
	return cl, err
}

func MakeClientCookie(c uuid.UUID) *http.Cookie {
	cookie := http.Cookie{
		Name:  ClientCookieName,
		Value: c.String(),
		Path:  "/",
	}

	return &cookie
}
