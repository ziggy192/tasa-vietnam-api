package controller

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

// todo get this from config file
const ClientID = "5958ddd634f2aea"
const AccessToken = "5eeae49394cd929e299785c8805bd168fc675280"

//handler to upload image and return url
func UploadImageHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("content-type", "application/json")
	enc := json.NewEncoder(w)

	fmt.Println("uploadImageHandler hit")
	r.ParseMultipartForm(10 << 20) // limit 10*2^20

	file, fileHeader, err := r.FormFile("image")
	check(err)
	defer file.Close()

	fmt.Printf("Uploaded file %s\n", fileHeader.Filename)
	fmt.Printf("File size=%v\n", fileHeader.Size)

	resJSON := UploadImageToImgur(file)

	//get url from resJson

	if resJSON["success"].(bool) == false {
		fmt.Println("imgur status == false")
		bufio.NewWriter(w).WriteString("imgur status == false")
		return
	}

	//have a separate struct for this
	link := resJSON["data"].(map[string]interface{})["link"].(string)

	//map to response dto
	resPostImage := UploadImageResponse{
		URL: link,
	}

	enc.Encode(resPostImage)
}

func TestUploadImgur(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("assets/57232103_572879629870418_546118032422862848_n.jpg")
	check(err)
	res := UploadImageToImgur(file)
	w.Header().Set("Content-Type", "application/json")
	fmt.Println(w.Header())
	enc := json.NewEncoder(w)
	enc.Encode(&res)

}
func UploadImageToImgur(image io.Reader) map[string]interface{} {
	//var requestBody bytes.Buffer
	//
	//multipartWriter := multipart.NewWriter(&requestBody)
	//
	//imageFieldWriter, err := multipartWriter.CreateFormField("image")
	//
	//check(err)
	//_, err = imageFieldWriter.Write(image)
	//check(err)
	//
	//err = multipartWriter.Close()
	//check(err)
	//
	//
	//fmt.Println(requestBody)

	requestBody := &bytes.Buffer{}
	multipartWriter := multipart.NewWriter(requestBody)
	//file, errFile1 := os.Open("/Users/ziggy192/Documents/Pics/rJP49tBP.png")
	//defer file.Close()
	part1, errFile1 := multipartWriter.CreateFormFile("image", filepath.Base("/Users/ziggy192/Documents/Pics/rJP49tBP.png"))
	_, errFile1 = io.Copy(part1, image)
	if errFile1 != nil {

		fmt.Println(errFile1)
	}
	err := multipartWriter.Close()
	if err != nil {
		fmt.Println(err)
	}

	req, err := http.NewRequest("POST", "https://api.imgur.com/3/upload", requestBody)
	check(err)
	req.Header.Set("Content-Type", multipartWriter.FormDataContentType())
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", AccessToken))

	response, err := http.DefaultClient.Do(req)
	check(err)
	defer response.Body.Close()

	var result map[string]interface{}
	json.NewDecoder(response.Body).Decode(&result)
	fmt.Println("imgurResponse=", result)

	return result

}
