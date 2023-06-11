package userServices

import (
	"encoding/json"
	"fmt"
	"net/http"
	userModel "github.com/atticus64/users-api/src/models/user"
	user "github.com/atticus64/users-api/src/controllers/user/dto"
	"github.com/atticus64/users-api/src/db"
)

type User struct {
	ID int
	Age int
	Name string
	Language string
}

type UserData struct {
	Age int
	Name string
	Language string
}

func CreateUser(w http.ResponseWriter, r *http.Request) {

	var user user.CreateUser

	if (r.Method == "POST") {

		err := json.NewDecoder(r.Body).Decode(&user)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Server Error")
			return
		}

		var model userModel.Model

		model.Name = user.Name
		model.Age = user.Age
		model.Language = user.Language

		result := db.Ctx.Create(&model)

		if result.Error != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Bad json")
			return 
		}

		w.WriteHeader(http.StatusCreated)
		data, errorJson := json.Marshal(&user)

		if errorJson != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Bad json")
			return 
		}

		fmt.Fprintf(w, "User: %s\n", data)

	} else  {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Bad Request, method not supported")

	}

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "DELETE" {
		keys, ok := r.URL.Query()["id"]

		if !ok || len(keys[0]) < 1 {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Printf("Url Param 'id' is missing")
			return
		}

		key := keys[0]

		var user User

		result := db.Ctx.First(&user, key)

		if result.Error != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error: %v\n", result.Error)
			return
		}


		res := db.Ctx.Delete(&user)

		if res.Error != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error: %v\n", res.Error)
			return
		}

		w.WriteHeader(http.StatusAccepted)
		fmt.Fprintf(w, "User deleted successfully")

	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Method not supported")
	}
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var data UserData

		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Server Error")
			return
		}


		keys, ok := r.URL.Query()["id"]

		if !ok || len(keys[0]) < 1 {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Printf("Url Param 'id' is missing")
			return
		}

		// Query()["key"] will return an array of items, 
		// we only want the single item.
		key := keys[0]

		var user User

		result := db.Ctx.First(&user, key)

		if result.Error != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error: %v\n", result.Error)
			return
		}

		user.Name = data.Name
		user.Age = data.Age
		user.Language = data.Language

		res := db.Ctx.Updates(&user)

		if res.Error != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error: %v\n", res.Error)
			return
		}

		dataJson, errorJson := json.Marshal(&user)

		if errorJson != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Bad json")
			return 
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "%s", dataJson)

	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Method not supported")
	}
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	var users []User

	result := db.Ctx.Find(&users)

	if result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Server Error")
		return
	}


	data, errorJson := json.Marshal(&users)

	if errorJson != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Bad json")
		return 
	}

	fmt.Fprintf(w, "%s", data)
}


