package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"forum/utils"
)

type MainPageResponse struct {
	Message     string `json:"message"`
	UserDetails string `json:"userDetails,omitempty"`
	IsLoggedIn  bool   `json:"isLoggedIn"`
}

func MainPageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: mainpage")

	// Извлекаем идентификатор сессии из куки
	sessionID, err := r.Cookie("session_id")
	if err != nil {
		// Если кука не найдена, отправляем сообщение для неавторизованных пользователей
		response := MainPageResponse{
			Message:    "Hello, user!",
			IsLoggedIn: false,
		}

		sendJSONResponse(w, response)
		return
	}

	// Получаем сессию по идентификатору
	session := utils.GetSession(sessionID.Value)
	if session == nil {
		// Если сессия не найдена, отправляем сообщение для неавторизованных пользователей
		response := MainPageResponse{
			Message:    "Hello, user!",
			IsLoggedIn: false,
		}

		sendJSONResponse(w, response)
		return
	}

	// Если пользователь авторизован, отправляем сообщение с дополнительной информацией
	user, ok := session.Values["user"].(utils.User)
	if !ok {
		http.Error(w, "User information not found in session", http.StatusInternalServerError)
		return
	}

	response := MainPageResponse{
		Message:     "Hello!",
		UserDetails: fmt.Sprintf("Nickname: %s", user.Nickname),
		IsLoggedIn:  true,
	}

	sendJSONResponse(w, response)
}

// sendJSONResponse отправляет JSON-ответ
func sendJSONResponse(w http.ResponseWriter, response MainPageResponse) {
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
    fmt.Println(string(jsonResponse))
	w.Write(jsonResponse)
}