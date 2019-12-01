package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
)

// todo get this from config file
const ClientID = "5958ddd634f2aea"
const AccessToken = "567d191b832101282951460d490181a4ca8eb3e9"

func UploadImageHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("uploadImageHandler hit")
	r.ParseMultipartForm(10 << 20)

	file, fileHeader, err := r.FormFile("image")
	if err != nil {
		fmt.Println("Error when upload file")
		fmt.Println(err)
		return
	}
	defer file.Close()

	fmt.Printf("Uploaded file %s\n", fileHeader.Filename)
	fmt.Printf("File size=%v", fileHeader.Size)

	//todo upload to imgur
}

func TestUploadImgur(w http.ResponseWriter, r *http.Request) {
	file, err := ioutil.ReadFile("assets/57232103_572879629870418_546118032422862848_n.jpg")
	check(err)
	w.Header().Set("Content-Type", "application/json")
	fmt.Println(w.Header())
	res := UploadImageToImgur(file)
	enc := json.NewEncoder(w)
	enc.Encode(&res)

}
func UploadImageToImgur(image []byte) map[string]interface{} {
	var requestBody bytes.Buffer

	multipartWriter := multipart.NewWriter(&requestBody)

	imageFieldWriter, err := multipartWriter.CreateFormField("image")

	check(err)
	_, err = imageFieldWriter.Write(image)
	check(err)

	multipartWriter.Close()

	req, err := http.NewRequest("POST", "https://api.imgur.com/3/upload", &requestBody)
	check(err)

	req.Header.Set("Authorization", fmt.Sprint("Client-ID ", ClientID))
	req.Header.Set("Content-Type", multipartWriter.FormDataContentType())
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", AccessToken))

	response, err := http.DefaultClient.Do(req)
	check(err)

	var result map[string]interface{}
	json.NewDecoder(response.Body).Decode(&result)

	fmt.Println(result)

	return result

}
