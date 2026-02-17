package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type TODO struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func h2(w http.ResponseWriter, r *http.Request) {
	url := "https://jsonplaceholder.typicode.com/todos/2"

	resp, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(resp)

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		var todoitem TODO
		err := json.NewDecoder(resp.Body).Decode(&todoitem)
		if err != nil {
			log.Printf("Error decoding JSON: %v", err)
			http.Error(w, "Error decoding JSON", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(todoitem)
	} else {
		http.Error(w, fmt.Sprintf("API returned status: %d", resp.StatusCode), http.StatusInternalServerError)
	}
}
func main() {

	// todoitem := TODO{
	// 	UserID:    1,
	// 	ID:        1,
	// 	Title:     "delectus aut autem",
	// 	Completed: false,
	// }

	// list, err := json.MarshalIndent(todoitem, "", "\t")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(string(list))

	// json.Unmarshal(list, &todoitem)
	// fmt.Println(todoitem)

	http.ListenAndServe(":8080", http.HandlerFunc(h2))
}
