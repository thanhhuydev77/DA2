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
)

func ReadUtilityTable(itemid string, category string) []Model.ItemUtility {

	var ListItemInCategory []Model.ItemUtility
	var ListUtility []float64
	var listItemId Model.ItemIdListResult
	csvCategoryFile, err := os.OpenFile("Storage/"+category+"_Utility.csv", os.O_RDONLY, 0777)
	if err != nil {
		fmt.Println(err)
		return nil
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
		if line[0] == "" {
			for _, itemId := range line {
				if itemId != "" {
					listItemId.ItemIds = append(listItemId.ItemIds, itemId)
				}
			}
			//line chứa ultility của Item đó
		} else if line[0] == itemid {
			for index, utility := range line {
				//skip cell chứa ID
				if index == 0 {
					continue
				}
				fltUtility, _ := strconv.ParseFloat(utility, 5)
				ListUtility = append(ListUtility, fltUtility)
				ListItemInCategory = append(ListItemInCategory, Model.ItemUtility{
					ItemId:       listItemId.ItemIds[index-1],
					UtilityValue: ListUtility[index-1],
				})
			}
		}
	}
	return ListItemInCategory
}
