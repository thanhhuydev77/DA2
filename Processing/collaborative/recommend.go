package collaborative

import (
	"Project2/Read"
	"encoding/csv"
	"fmt"
	"io"
	"log"
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

	f, _ := os.OpenFile("./storage/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	defer f.Close()

	io.Copy(f, file)

	_, _, items := Read.ReadFileProductReviewCSV("./storage/" + handler.Filename)

	data := [][]string{}
	data = append(data, items)

	// for _, user := range users {

	// 	data = append(data, []string{
	// 		user, 
	// 	})
	// }
	fmt.Print(len(items))
	go writeCleanData(data, "./storage/Clean"+handler.Filename)
	io.WriteString(w, `{Filename: `+handler.Filename+`}`)
}

func writeCleanData(data [][]string, path string) {
	csvData, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		log.Fatal(err)
	}
	defer csvData.Close()

	csv := csv.NewWriter(csvData)
	defer csv.Flush()

	for _, value := range data {
		errWrite := csv.Write(value)
		if errWrite != nil {
			log.Fatal("Write file err")
		}
	}
}
