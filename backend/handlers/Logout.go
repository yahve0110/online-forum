package handlers

import (
	"fmt"
	"forum/utils"
	"net/http"
	"time"
)

// Ваш обработчик logout
func logoutHandler(w http.ResponseWriter, r *http.Request) {
    // Получение идентификатора сессии из куки
	fmt.Println("Endpoint hit: logout")

	fmt.Println(r.Cookie("session_id"))
    sessionID, err := r.Cookie("session_id")
    if err != nil {
        http.Error(w, "No session ID found", http.StatusUnauthorized)
        return
    }

    // Удаление сессии из хранилища
    utils.DeleteSession(sessionID.Value)

    // Опционально: удаление куки с идентификатором сессии
    cookie := &http.Cookie{
        Name:    "session_id",
        Value:   "",
        Expires: time.Now(),
        Path:    "/",
    }
    http.SetCookie(w, cookie)

    // Редирект на главную страницу или другую страницу после выхода
    http.Redirect(w, r, "/", http.StatusSeeOther)
}
