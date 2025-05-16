package handlers

import (
	"bytes" 
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart" //for creating multipart request
	"net/http"
	"os"
)

type UploadSuccessResp struct {
	Data struct {
		DisplayURL string `json:"display_url"`
		Title      string `json:"title"`
		URL        string `json:"url"`
		Size       int    `json:"size"`
		Expiration int    `json:"expiration"`
	} `json:"data"`
	Success bool `json:"success"`
	Status  int  `json:"status"`
}

//ImgBB responses
type ImgBBErrorResp struct {
	Error struct {
		Message string `json:"message"`
		Code    int    `json:"code"`
	} `json:"error"`
	StatusTxt string `json:"status_txt"`
	Success   bool   `json:"success"`
	Status    int    `json:"status"`
}

func UploadImage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// 1. Parse the incoming multipart form from the client
	if err := r.ParseMultipartForm(32 << 20); err != nil { // 32MB max
		http.Error(w, "Failed to parse multipart form: "+err.Error(), http.StatusBadRequest)
		return
	}

	// 2. Get the file from the form data
	file, handler, err := r.FormFile("image") // "image" is the form field name
	if err != nil {
		http.Error(w, "Error retrieving the file from form-data: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	apiKey := os.Getenv("IMGBB_API_KEY")
	if apiKey == "" {
		http.Error(w, "Server configuration error: IMGBB_API_KEY not set", http.StatusInternalServerError)
		return
	}
	imgbbURL := fmt.Sprintf("https://api.imgbb.com/1/upload?key=%s", apiKey)

	// 3. Create a new multipart request body for ImgBB
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Create a form file field for the image
	part, err := writer.CreateFormFile("image", handler.Filename)
	if err != nil {
		http.Error(w, "Failed to create form file for ImgBB request: "+err.Error(), http.StatusInternalServerError)
		return
	}
	// Copy the uploaded file content to the new form file field
	if _, err = io.Copy(part, file); err != nil {
		http.Error(w, "Failed to copy file content for ImgBB request: "+err.Error(), http.StatusInternalServerError)
		return
	}


	// Close the multipart writer. This is important as it writes the trailing boundary.
	if err = writer.Close(); err != nil {
		http.Error(w, "Failed to close multipart writer for ImgBB: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// 4. Make the POST request to ImgBB
	req, err := http.NewRequest("POST", imgbbURL, body)
	if err != nil {
		http.Error(w, "Failed to create request to ImgBB: "+err.Error(), http.StatusInternalServerError)
		return
	}
	// Set the Content-Type header with the correct boundary, which is set by the writer
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Failed to send request to ImgBB: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// 5. Read the response body from ImgBB to get more details
	respBodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Failed to read ImgBB response body: "+err.Error(), http.StatusInternalServerError)
		return
	}

	//raw response for debugging
	fmt.Printf("ImgBB Raw Response (Status: %d): %s\n", resp.StatusCode, string(respBodyBytes))

	// 6. Decode the ImgBB JSON response
	var successResp UploadSuccessResp
	if err := json.Unmarshal(respBodyBytes, &successResp); err == nil && successResp.Success {
		fmt.Fprintf(w, "Image uploaded successfully: %s", successResp.Data.URL)
		// 7. Save to database
		if DB == nil {
			http.Error(w, "Database connection not initialized", http.StatusInternalServerError)
			return
		}

		imageURL := successResp.Data.DisplayURL
		if imageURL == "" {
			imageURL = successResp.Data.URL
		}

		query := `INSERT INTO images (title, img_url, size_kb, expiration) VALUES ($1, $2, $3, $4) RETURNING id`
		var imageID int
		// ImgBB size is in bytes, convert to KB for DB
		sizeKB := successResp.Data.Size / 1024
		if successResp.Data.Size > 0 && sizeKB == 0 { // handle very small files
			sizeKB = 1
		}

		errDb := DB.QueryRow(query, successResp.Data.Title, imageURL, sizeKB, successResp.Data.Expiration).Scan(&imageID)
		if errDb != nil {
			fmt.Printf("DB Insert Error: %v\n", errDb)
			return
		}
		fmt.Printf("Image metadata saved to DB with ID: %d\n", imageID)
		fmt.Fprintf(w, "Image metadata saved to DB with ID: %d %s", imageID, successResp.Data.URL)

		return
	}

	// If not a success response, try to parse as an ImgBB error response
	var errorResp ImgBBErrorResp
	if err := json.Unmarshal(respBodyBytes, &errorResp); err == nil && !errorResp.Success {
		// ImgBB returned a structured error
		errMsg := fmt.Sprintf("ImgBB Error (Code: %d): %s", errorResp.Error.Code, errorResp.Error.Message)
		fmt.Println(errMsg)                                   // Log it
		http.Error(w, errMsg, http.StatusInternalServerError) // Send ImgBB's error to your client
		return
	}

	// If decoding into specific structs failed or it's an unexpected format
	// but ImgBB returned a non-2xx status
	if resp.StatusCode >= 400 {
		errMsg := fmt.Sprintf("ImgBB returned HTTP status %d. Response: %s", resp.StatusCode, string(respBodyBytes))
		fmt.Println(errMsg)
		http.Error(w, "Failed to upload image to ImgBB. Service returned an error.", http.StatusInternalServerError)
		return
	}

	http.Error(w, "An unexpected error occurred after attempting to upload to ImgBB.", http.StatusInternalServerError)
}
