package Read

import (
	. "Project2/Model"
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func ReadFileitemPropertyCSV(path string) int {

	csvFile, _ := os.Open(path)
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var ItemList []Item
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		var rawPrice float64
		var discount int
		var likesCount int
		var id int
		var currentPrice float64
		if tempcurrentPrice, Error := strconv.ParseFloat(line[3], 2); Error != nil {
			continue
		} else {
			currentPrice = tempcurrentPrice
		}
		if temprawPrice, Error1 := strconv.ParseFloat(line[4], 2); Error1 != nil {
			continue
		} else {
			rawPrice = temprawPrice
		}

		if tempdiscount, Error2 := strconv.Atoi(line[6]); Error2 != nil {
			continue
		} else {
			discount = tempdiscount
		}

		if templikesCount, Error3 := strconv.Atoi(line[7]); Error3 != nil {
			continue
		} else {
			likesCount = templikesCount
		}
		isNew := line[8] == "TRUE"
		codCountry := strings.Split(line[11], ",")
		if tempid, Error4 := strconv.Atoi(line[20]); Error4 != nil {
			continue
		} else {
			id = tempid
		}

		ItemList = append(ItemList, Item{
			Category:            line[0],
			Subcategory:         line[1],
			Name:                line[2],
			CurrentPrice:        currentPrice,
			RawPrice:            rawPrice,
			Currency:            line[5],
			Discount:            discount,
			LikesCount:          likesCount,
			IsNew:               isNew,
			Brand:               line[9],
			BrandUrl:            line[10],
			CodCountry:          codCountry,
			Variation0Color:     line[12],
			Variation1Color:     line[13],
			Variation0Thumbnail: line[14],
			Variation0Image:     line[15],
			Variation1Thumbnail: line[16],
			Variation1Image:     line[17],
			ImageUrl:            line[18],
			Url:                 line[19],
			Id:                  id,
			Model:               line[21],
		})
	}
	return len(ItemList)
}

func ReadFileCSV(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	fileItemProperty, handlerItemProperty, err := r.FormFile("fileItemProperty")
	if err != nil {
		io.WriteString(w, `{"message":"Upload File Item Property Fail!"}`)
		return
	}
	defer fileItemProperty.Close()

	f, _ := os.OpenFile("./storage/"+handlerItemProperty.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	defer f.Close()
	io.Copy(f, fileItemProperty)
	len := ReadFileitemPropertyCSV("./storage/" + handlerItemProperty.Filename)
	io.WriteString(w, `{"message":"Upload File Item Property Successful with `+strconv.Itoa(len)+` lines"}`)
}
