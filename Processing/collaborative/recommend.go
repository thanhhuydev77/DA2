package collaborative

import (
	"Project2/Read"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func RecommendUploadFile(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	defer file.Close()

	f, err1 := os.OpenFile("./Storage/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err1 != nil {
		fmt.Println("Error: ", err1.Error())
	}
	defer f.Close()

	io.Copy(f, file)

	itemRating, _ := Read.ReadFileProductReviewCSV("./Storage/" + handler.Filename)
	//result :=
	userIds := UserSimilarity(itemRating, 10, "./Storage/Clean"+handler.Filename)
	jsonresult, _ := json.Marshal(userIds)
	io.WriteString(w, `{"userIds": `+string(jsonresult)+"}")
}

func RecommendUserBased(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	filename := r.URL.Query().Get("filename")
	userId := r.URL.Query().Get("userId")

	result := Read.RecommendUserBasedFunc(filename, userId)
	jsonresult, _ := json.Marshal(result)
	io.WriteString(w, `{ "products" : `+string(jsonresult)+"}")
}
