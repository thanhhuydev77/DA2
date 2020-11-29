package processing

import (
	"Project2/Model"
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

var SplitOk bool

func SaveCategoryTable(itemList []Model.Item, itemCategories *[]string, itemProperties *[]string) {

	SplitOk = false

	saveMapItemAndCategory(itemList)

	*itemProperties = append(*itemProperties, "id", "current_price", "raw_price", "likes_count", "is_new", "codCountry", "brand", "color")
	//save to csv
	for _, category := range *itemCategories {
		//create file
		csvFile, err := os.Create("Storage/" + category + ".csv")
		if err != nil {
			fmt.Println(err)
			return
		}
		csvWriter := csv.NewWriter(csvFile)
		//write header
		errWrite := csvWriter.Write(*itemProperties)
		if errWrite != nil {
			print(err)
			return
		}
		csvWriter.Flush()
		errClose := csvFile.Close()
		if errClose != nil {
			print(err)
			return
		}

	}
	// write content
	for _, item := range itemList {
		csvCategoryFile, err := os.OpenFile("Storage/"+item.Subcategory+".csv", os.O_APPEND, 0777)
		if err != nil {
			fmt.Println(err)
			return
		}

		csvWriter2 := csv.NewWriter(csvCategoryFile)
		itemId := strconv.Itoa(item.Id)
		itemCurrentPrice := fmt.Sprintf("%f", item.CurrentPrice)
		itemRawPrice := fmt.Sprintf("%f", item.RawPrice)
		itemLikeCount := strconv.Itoa(item.LikesCount)
		itemIsNew := strconv.FormatBool(item.IsNew)
		itemCodCountry := strings.Join(item.CodCountry, ",")
		itemColor := item.Variation0Color + "," + item.Variation1Color
		errWrite2 := csvWriter2.Write([]string{itemId, itemCurrentPrice, itemRawPrice, itemLikeCount, itemIsNew, itemCodCountry, item.Brand, itemColor})
		if errWrite2 != nil {
			print(errWrite2)
			return
		}
		csvWriter2.Flush()
		errClose := csvCategoryFile.Close()
		if errClose != nil {
			print(err)
			return
		}
	}
	fmt.Print("split file successfully\n")
	itemList = nil
	itemProperties = nil
	SaveAllUtilityTable(*itemCategories)
	SplitOk = true
}

func saveMapItemAndCategory(itemList []Model.Item) {
	os.Create("Storage/Mapping.csv")
	for _, item := range itemList {
		csvFile, err := os.OpenFile("Storage/Mapping.csv", os.O_APPEND, 0777)
		if err != nil {
			fmt.Println(err)
			return
		}
		csvWriter := csv.NewWriter(csvFile)
		errWrite2 := csvWriter.Write([]string{item.Subcategory, strconv.Itoa(item.Id)})
		if errWrite2 != nil {
			print(errWrite2)
			return
		}
		csvWriter.Flush()
		errClose := csvFile.Close()
		if errClose != nil {
			print(err)
			return
		}
	}
}
func FindCategory(Id int) string {
	csvCategoryFile, err := os.OpenFile("Storage/Mapping.csv", os.O_RDONLY, 0777)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	reader := csv.NewReader(bufio.NewReader(csvCategoryFile))
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		//line 0, chứa dãy ID của Item
		if line[1] == strconv.Itoa(Id) {
			return line[0]
		}
	}
	return ""
}
