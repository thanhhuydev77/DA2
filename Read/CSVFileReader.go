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

func ReadFileProductReviewCSV(path string) []Model.ProductReview {
	csvFile, _ := os.Open(path)
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var ProductReviewList []Model.ProductReview
	reader.Read()
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal("err")
			log.Fatal(line)
		}

		rating, _ := strconv.Atoi(line[17])
		ProductReviewList = append(ProductReviewList, Model.ProductReview{
			Id:             line[0],
			Name:           line[9],
			ReviewUsername: line[23],
			Rating:         rating,
		})
		// file, err := os.OpenFile(path+"Clean.csv", os.O_CREATE, 0644)
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// file.Close()

		// for _, item := range ProductReviewList {
		// 	go writeCleanData([]string{item.Id, item.Name, item.ReviewUsername, strconv.Itoa(item.Rating)}, path + "Clean.csv")
		// }

	}
	//return len(ProductReviewList)
	return ProductReviewList[: 10]
}


