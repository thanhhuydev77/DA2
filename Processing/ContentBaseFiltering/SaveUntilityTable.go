package processing

import (
	"Project2/Model"
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

var numberfilecomplete int
var categorylen int

func SaveAllUtilityTable(ItemListCategory []string) bool {
	numberfilecomplete = 0
	categorylen = len(ItemListCategory)
	for _, category := range ItemListCategory {

		ListItemInCategory := []Model.ItemCategory{}
		csvCategoryFile, err := os.OpenFile("Storage/"+category+".csv", os.O_RDONLY, 0777)
		if err != nil {
			fmt.Println(err)
			return false
		}
		reader := csv.NewReader(bufio.NewReader(csvCategoryFile))
		headerline := true
		for {
			line, error := reader.Read()
			if error == io.EOF {
				break
			} else if error != nil {
				log.Fatal(error)
			}

			if headerline {
				headerline = false
				continue
			}
			id, _ := strconv.Atoi(line[0])
			tempcurrentPrice, _ := strconv.ParseFloat(line[1], 2)
			temprawPrice, _ := strconv.ParseFloat(line[2], 2)
			templikesCount, _ := strconv.Atoi(line[3])
			isNew := line[4] == "TRUE"
			codCountry := strings.Split(line[5], ",")
			color := strings.Split(line[6], ",")
			itemcategory := Model.ItemCategory{
				Id:           id,
				CurrentPrice: tempcurrentPrice,
				RawPrice:     temprawPrice,
				LikesCount:   templikesCount,
				IsNew:        isNew,
				CodCountry:   codCountry,
				Brand:        line[6],
				Color:        color,
			}
			ListItemInCategory = append(ListItemInCategory, itemcategory)
		}
		go SaveUntility1Table(ListItemInCategory, category)
	}
	return true
}
func SaveUntility1Table(itemlist []Model.ItemCategory, category string) bool {

	csvFile, err := os.Create("Storage/" + category + "_Utility.csv")
	if err != nil {
		fmt.Println(err)
		return false
	}
	listItemheader := []string{}
	listItemheader = append(listItemheader, "")
	utititytable := [][]string{}
	for indexId, item1 := range itemlist {
		utitityRow := []string{}
		utitityRow = append(utitityRow, strconv.Itoa(item1.Id))
		listItemheader = append(listItemheader, strconv.Itoa(item1.Id))

		for indexId2, item2 := range itemlist {
			if indexId == indexId2 {
				utitityRow = append(utitityRow, "")
			} else {
				utility := calcDeltaPrice(item1.CurrentPrice, item1.RawPrice, item2.CurrentPrice, item2.RawPrice)
				utility += calcsimilarCodCountry(item1.CodCountry, item2.CodCountry)
				utility += checkBrand(item1.Brand, item2.Brand)
				utility += calcColorSimilar(item1.Color, item2.Color)
				utitityRow = append(utitityRow, fmt.Sprintf("%f", utility))
			}
		}
		utititytable = append(utititytable, utitityRow)
	}
	csvwriter := csv.NewWriter(csvFile)
	//write header
	csvwriter.Write(listItemheader)
	//write matrix
	csvwriter.WriteAll(utititytable)
	csvwriter.Flush()
	csvFile.Close()

	numberfilecomplete += 1
	fmt.Print(" number file complete :" + strconv.Itoa(numberfilecomplete) + "\n")
	if numberfilecomplete == categorylen {
		fmt.Print("All done")
	}

	defer os.Remove("Storage/" + category + ".csv")
	return true
}
func calcDeltaPrice(currentprice1 float64, rawprice1 float64, currentprice2 float64, rawprice2 float64) float64 {
	deltaRawPrice := rawprice2 - rawprice1
	deltaCurrentPrice := math.Abs(currentprice2 - currentprice1)
	result := deltaRawPrice / math.Pow(math.E, deltaCurrentPrice)
	return result
}
func calcsimilarCodCountry(codCountry1 []string, codCountry2 []string) float64 {
	if len(codCountry1) == 0 {
		return float64(len(codCountry2))
	}
	if len(codCountry2) == 0 {
		return 0.0
	}
	var result float64
	result = 0
	for _, valuecountry1 := range codCountry1 {
		for _, valuecountry2 := range codCountry2 {
			if valuecountry1 == valuecountry2 {
				result += 1
			}
		}
	}
	return result
}
func checkBrand(brand1 string, brand2 string) float64 {
	if brand1 == brand2 && brand1 != "" {
		return 5
	}
	if brand1 == "" && brand2 != "" {
		return 2
	}
	return 0
}
func calcColorSimilar(color1 []string, color2 []string) float64 {
	result := float64(0)
	for _, value1 := range color1 {
		for _, value2 := range color2 {
			if value1 == value2 {
				result += 1
			}
		}
	}
	return result
}
