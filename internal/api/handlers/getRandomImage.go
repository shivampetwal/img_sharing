//get data from db

package handlers

import (
	"database/sql"
	"encoding/json"
	//"errors"
	//"fmt"
	//"math/rand"
	"net/http"
	"time"
)

// DB is initialized by handlers.SetDB in main.go


type ImageDataResp struct {
	ID         int       `json:"id"`
	ImgURL     string    `json:"img_url"`
	CreatedAt  time.Time `json:"created_at"`
	Upvotes    int       `json:"upvotes"`
	Downvotes  int       `json:"downvotes"`
	TotalViews int       `json:"total_views"`
	Caption    string    `json:"caption"`
	Title      string    `json:"title"`
	SizeKB     int       `json:"size_kb"`
	Expiration int       `json:"expiration"`
}

func GetData(w http.ResponseWriter, r *http.Request) {
	// const query = `
	// 		SELECT *
	// 		FROM images
	// 		ORDER BY RANDOM()
	// 		LIMIT 1;`


	const query = `
		SELECT
			id,
			img_url,
			created_at,
			upvotes,
			downvotes,
			total_views,
			caption,
			title,
			size_kb,
			expiration
		FROM images
		ORDER BY random()
		LIMIT 1;
	`

	var img ImageDataResp
	var caption, title sql.NullString

	err := DB.QueryRow(query).Scan(
		&img.ID,
		&img.ImgURL,
		&img.CreatedAt,
		&img.Upvotes,
		&img.Downvotes,
		&img.TotalViews,
		&caption,
		&title,
		&img.SizeKB,
		&img.Expiration,
	)

	img.Caption = ""
	if caption.Valid {
		img.Caption = caption.String
	}

	img.Title = ""
	if title.Valid {
		img.Title = title.String
	}


	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "no images found", http.StatusNotFound)
			return
		}
		http.Error(w, "internal server error"+err.Error(), http.StatusInternalServerError)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(img)
}

//________________
//__________________________
//__________________________
// func GenerateRandomID() (randomID int) {
// 	// count(*) from table
// 	//generate radom number from 0 to count(*)

// 	var count int
// 	query := "SELECT COUNT(id) FROM images"
// 	err := DB.QueryRow(query).Scan(&count)
// 	if err != nil {
// 		fmt.Println("Error fetching count:", err)
// 		return
// 	}
// 	if count == 0 {
// 		fmt.Println("Error: No images found in the database")
// 		return
// 	}

// 	// Generate a random number between 0 and count
// 	randomID = rand.Intn(count) + 1 // Generates from 1 to rowCount

// 	return randomID
// }

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
