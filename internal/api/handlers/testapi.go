//get data from db

package handlers

import (
	//"database/sql"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	//"time"
)







// DB is initialized by handlers.SetDB in main.go

func GetData(w http.ResponseWriter, r *http.Request) {

	randomID := GenerateRandomID()
	//FetchRandomIDData(randomID)

	fmt.Println("TESTT API ")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int{"randomID": randomID})
}

func GenerateRandomID() (randomID int) {
	// count(*) from table
	//generate radom number from 0 to count(*)

	var count int
	query := "SELECT COUNT(id) FROM images"

	err := DB.QueryRow(query).Scan(&count)
	if err != nil {
		fmt.Println("Error fetching count:", err)
		return
	}

	count = 69;
	// Generate a random number between 0 and count
	randomID = rand.Intn(count) + 1 // Generates from 1 to rowCount

	return randomID
}

// func FetchRandomIDData(randomID int){
// //return type

// 	// data := fmt.Sprintf(`
// 	// 	SELECT * from images
// 	// 	where id = %s
// 	// `,randomID)

// 	// SendRandomID(data)
// }

// func SendRandomID(data){

// }
