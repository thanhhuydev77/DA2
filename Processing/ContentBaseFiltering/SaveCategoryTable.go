package processing

import (
	"Project2/Model"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)
var splitOk bool
func SaveCategoryTable(itemList []Model.Item, itemCategories *[]string, itemproperties *[]string) {

	splitOk = false
	*itemproperties = append(*itemproperties, "id", "current_price", "raw_price", "likes_count", "is_new", "codCountry","brand")
	//save to csv
	for _, category := range *itemCategories {
		//create file
		csvFile, err := os.Create("Storage/" + category + ".csv")
		if err != nil {
			fmt.Println(err)
			return
		}
		csvwriter := csv.NewWriter(csvFile)
		//write header
		csvwriter.Write(*itemproperties)
		csvwriter.Flush()
		csvFile.Close()
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

		csvWriter2.Write([]string{itemId, itemCurrentPrice, itemRawPrice, itemLikeCount, itemIsNew, itemCodCountry,item.Brand})
		csvWriter2.Flush()
		csvCategoryFile.Close()
	}
	fmt.Print("split file successfully\n")
	splitOk = true
	SaveAllUtilityTable(*itemCategories)

}

