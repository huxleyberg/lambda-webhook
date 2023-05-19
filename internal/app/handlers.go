package app

import (
	"encoding/json"
	"net/http"
)

func GetAllUsersHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users := []struct {
			ID        string `json:"id"`
			Firstname string `json:"first_name"`
			Lastname  string `json:"last_name"`
		}{
			{
				ID:        "435564bggfh345",
				Firstname: "john",
				Lastname:  "Doe",
			},
			{
				ID:        "435564bggfh345",
				Firstname: "john",
				Lastname:  "Doe",
			},
		}

		SendJson(w, http.StatusOK, users)

	}
}

func AddUserHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		SendJson(w, http.StatusOK, map[string]interface{}{"message": "Users added successfully", "statusCode": 200})
	}
}

func SendJson(w http.ResponseWriter, status int, data interface{}) {
	jsonString, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(jsonString)
}
