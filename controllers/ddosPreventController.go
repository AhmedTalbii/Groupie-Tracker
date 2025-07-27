package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"groupietracker/config"
	"groupietracker/models"
)

func DdosPreventController(w http.ResponseWriter, r *http.Request) error {
	// adress ip
	ClientIp := strings.Split(r.RemoteAddr, ":")[0]
	// Initialize map if nil (only needed once globally — add this in setup, not per request)
	if models.Clients == nil {
		models.Clients = make(map[string]*models.Client)
	}
	// take the informations about the client from the struct
	Client, Exists := models.Clients[ClientIp]
	// delet the bann after one houre
	if Exists && Client.Banned {
		if time.Since(Client.LastTime) > config.BannTime*time.Hour {
			// Ban expired — unban
			Client.Banned = false
			Client.Warning = false
			Client.Requests = 0
			Client.LastTime = time.Now()
		} else {
			// Still banned
			w.WriteHeader(http.StatusTooManyRequests)
			w.Write([]byte("You are banned for 1 hour"))
			return errors.New("client is banned")
		}
	}

	// if does not exist client set to 0
	if !Exists {
		fmt.Println(4)
		models.Clients[ClientIp] = &models.Client{
			Requests: 0,
			LastTime: time.Now(),
			Warning:  false,
			Banned:   false,
		}
	}

	// if time is less than time between addd 1 to the req else make it 0
	if Client.LastTime.Add(config.TimeBetweenEatchReq * time.Second).After(time.Now()) {
		Client.Requests++
	} else {
		Client.Requests = 0
	}
	Client.LastTime = time.Now()

	// check if the client check the web site before
	if models.Clients[ClientIp].Requests > config.RequestsNumbertToGetWarningOrBan { // if (N) of req > (RequestsNumbertToGetWarningOrBan) Show Error + return
		if !Client.Warning && !Client.Banned {
			fmt.Println(1)
			models.Clients[ClientIp] = &models.Client{
				Requests: 0,
				LastTime: time.Now(),
				Warning:  true,
				Banned:   false,
			}
			w.WriteHeader(http.StatusTooManyRequests)
			w.Write([]byte("429 Status Too Many Requests || Warning next you will be banned for 1 hour"))
			return errors.New("429 Status Too Many Requests")
		} else if Client.Warning && !Client.Banned {
			Client.Banned = true
            Client.Warning = false
            Client.Requests = 0
			w.WriteHeader(http.StatusTooManyRequests)
			w.Write([]byte("You Are banned for 1 hour"))
			return errors.New("client is banned")
		}
	}
	return nil
}
