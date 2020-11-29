package collaborative

import (
	"Project2/Read"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	//"strconv"
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

	length := Read.ReadFileProductReviewCSV("./storage/" + handler.Filename)
	
	for _, item := range length {
		str := []string{}
		str = append(str, item.Id)
		str = append(str, item.Name)
		str = append(str, item.ReviewUsername)
		str = append(str, strconv.Itoa(item.Rating))
			go writeCleanData(str, "./storage/Clean" + handler.Filename)
	}

	io.WriteString(w, `{Filename: `+ handler.Filename + `}`)
}


func writeCleanData(data []string, path string) {
	csvData, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	csv := csv.NewWriter(csvData)
	errWrite := csv.Write(data)
	if errWrite != nil {
		log.Fatal("Write file err")
	}
	csv.Flush()
	errClose := csvData.Close()
	if errClose != nil {
		log.Fatal("close file error")
	}
}