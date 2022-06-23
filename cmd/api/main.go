package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type student struct {
	ID    int
	Name  string
	Grade int
}

var data = []student{
	student{3, "ethan", 21},
	student{1, "wick", 22},
	student{5, "bourne", 23},
	student{4, "bond", 23},
}

func user(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		var id = r.FormValue("id")
		if id == "" {
			var result, err = json.Marshal(data)

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			if err2, n := w.Write(result); n != nil {
				fmt.Println(err2)
			}
			return
		}
		var result []byte
		var err error

		idConv, _ := strconv.Atoi(id)
		for _, each := range data {
			if each.ID == idConv {
				result, err = json.Marshal(each)

				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				if err2, n := w.Write(result); n != nil {
					fmt.Println(err2)
				}
				return
			}
		}

		http.Error(w, "User not found", http.StatusBadRequest)
		return
	} else if r.Method == "POST" {
		var max = 0
		for _, each := range data {
			if each.ID > max {
				max = each.ID
			}
		}

		var name = r.FormValue("name")
		grade, _ := strconv.Atoi(r.FormValue("grade"))
		data = append(data, student{max + 1, name, grade})
		var result, _ = json.Marshal(data[len(data)-1])
		if err2, n := w.Write(result); n != nil {
			fmt.Println(err2)
		}
		return
	} else if r.Method == "PUT" {
		var id = r.FormValue("id")
		var name = r.FormValue("name")
		grade, _ := strconv.Atoi(r.FormValue("grade"))
		idConv, _ := strconv.Atoi(id)

		for k, v := range data {
			if v.ID == idConv {
				data[k].Grade = grade
				data[k].Name = name
			}
		}

		var result []byte
		var err error
		for _, each := range data {
			if each.ID == idConv {
				result, err = json.Marshal(each)

				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				if err2, n := w.Write(result); n != nil {
					fmt.Println(err2)
				}
				return
			}
		}
	} else if r.Method == "DELETE" {
		var id = r.FormValue("id")
		idConv, _ := strconv.Atoi(id)

		for k, v := range data {
			if v.ID == idConv {
				data = append(data[:k], data[k+1:]...)
			}
		}
		var result, err = json.Marshal(data)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err2, n := w.Write(result); n != nil {
			fmt.Println(err2)
		}
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}

func main() {
	http.HandleFunc("/user", user)
	// log.Fatal(http.ListenAndServe(":8080", nil))
	if err2 := http.ListenAndServe(":8000", nil); err2 != nil {
		fmt.Println(err2)
	}
}
