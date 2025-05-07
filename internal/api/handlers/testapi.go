
//get data from db

package handlers

import (
	"fmt"
	"net/http"
	"database/sql"
	"math/rand"
	"time"
)

// Global database connection variable; ensure it is properly initialized elsewhere.
var db *sql.DB



func getData(){

	randomID  := GenerateRandomID()
	FetchRandomIDData(randomID)
	//SendRandomID()
//it will be called from fetch randomeif
}








func GenerateRandomID()(randomID int){
	// count(*) from table
	//generate radom number from 0 to count(*)

	var count int
	query := "SELECT COUNT(id) FROM images"

	rowCount := db.QueryRow(query)
		err := rowCount.Scan(&count)
		if err != nil {
			fmt.Println("Error fetching count:", err)
			return
		}
		// Generate a random number between 0 and count
		randomID = rand.Intn(count) + 1   // Generates from 1 to rowCount

		return randomID
}


func FetchRandomIDData(randomID int){
//return type

	query := fmt.Sprintf(`
		SELECT * from images
		where id = %s
	`,randomID)



	SendRandomID(data)
}


func SendRandomID(data){

}