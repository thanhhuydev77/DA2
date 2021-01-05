package Read

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

func ReadFileProductReviewCSV(path string) (map[string][]Model.ItemUserRating, map[string][]Model.UserItemRating) {
	csvFile, err := os.Open(path)
	if err != nil {
		fmt.Print(err)
	}
	lines := csv.NewReader(bufio.NewReader(csvFile))

	itemRating := make(map[string][]Model.ItemUserRating)
	userRating := make(map[string][]Model.UserItemRating)

	for {
		line, error := lines.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}

		rating, err := strconv.Atoi(line[17])
		if err != nil {
			continue
		}
		ID := line[0]
		//Name:= line[9]
		ReviewUsername := line[23]

		// next if id product or
		if ID == "" || ReviewUsername == "" || rating < 1 || rating > 5 {
			continue
		}

		if len(itemRating[ReviewUsername]) > 0 {
			itemRating[ReviewUsername] = append(itemRating[ReviewUsername], Model.ItemUserRating{
				Item:   ID,
				Rating: rating,
			})
		} else {
			itemRating[ReviewUsername] = []Model.ItemUserRating{
				{
					Item:   ID,
					Rating: rating,
				},
			}
		}

		if len(userRating[ID]) > 0 {
			userRating[ID] = append(userRating[ID], Model.UserItemRating{
				User:   ReviewUsername,
				Rating: rating,
			})
		} else {
			userRating[ID] = []Model.UserItemRating{
				{
					User:   ReviewUsername,
					Rating: rating,
				},
			}
		}
	}
	return itemRating, userRating
}

func RecommendUserBasedFunc(filename string, userId string) string {
	csvFile, err := os.Open("./Storage/" + filename)
	if err != nil {
		fmt.Print(err)
	}
	lines := csv.NewReader(bufio.NewReader(csvFile))
	lines.Comma = ','
	lines.FieldsPerRecord = -1
	result := "["
	for {
		line, error := lines.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}

		ID := line[0]
		if ID != userId {
			continue
		} else {
			length := len(line)
			for i := 1; i < length; i++ {
				result += line[i] + ","
			}
		}
		break
	}
	length := len(result)
	if length > 1 {
		length -= 1
	}
	return result[:length] + "]"
}
