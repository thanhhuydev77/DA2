package Read

import (
	"Project2/Model"
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
)

func ReadFileProductReviewCSV(path string) ([]Model.ProductReview, []string, []string) {
	csvFile, _ := os.Open(path)
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var ProductReviewList []Model.ProductReview

	users := []string{}
	items := []string{}
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal("err")
			log.Fatal(line)
		}

		rating, _ := strconv.Atoi(line[17])
		ID := line[0]
		Name:= line[9]
		ReviewUsername:= line[23]
		Rating:= rating
		ProductReviewList = append(ProductReviewList, Model.ProductReview{
			Id: ID,
			Name: Name,
			ReviewUsername: ReviewUsername,
			Rating: Rating,
		})

		_, userExist := findExist(users, ReviewUsername)
		if !userExist {
			users = append(users, ReviewUsername)
		}

		_, itemExist := findExist(items, ReviewUsername)
		if !itemExist {
			items = append(items, ID)
		}
	}
	return ProductReviewList[1:], users[1:], items[1:]
}


func findExist(list []string, item string) (int, bool)  {
	for i, value := range list {
		if (value == item){
			return i, true
		}
	}
	return -1, false
}