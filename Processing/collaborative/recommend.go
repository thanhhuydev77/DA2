package collaborative

import (
	"Project2/Read"
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
	UserSimilarity(itemRating, 10, "./Storage/Clean"+handler.Filename)
	//resultString := ""
	//for key, element := range result {
	//	resultString += "\n" + key + ": ["
	//	for _, value := range element {
	//		score := strconv.FormatFloat(value.score, 'f', 6, 64)
	//		resultString += value.userId + ":" + score + ", "
	//	}
	//	resultString = resultString[:len(resultString)-2] + "],"
	//}
	io.WriteString(w, `{Filename: `+handler.Filename)
	//io.WriteString(w, `{Filename: `+handler.Filename+ ",\n" +`result: [`+resultString+`]}`)
}

func RecommendUserBased(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	filename := r.URL.Query().Get("filename")
	userId := r.URL.Query().Get("userId")

	result := Read.RecommendUserBasedFunc(filename, userId)
	io.WriteString(w, result)
}
