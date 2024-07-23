package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

// structure for a specific book
type Book struct {
	Name          string `json:"name"`
	Author        string `json:"author"`
	AmountOfPages string `json:"amount_of_pages"`
}

// structure for all books
type Books struct {
	Books []Book `json:"books"`
}

// checking error
func CheckError(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("What do you want to do ?:\n1) Check available books\n2) take book\n3) give book\n4) Leave\n")

		filePath := "D:/io/books.json"
		jsonFile, err := os.Open(filePath)
		CheckError(err)
		defer jsonFile.Close()

		var answer int
		fmt.Scan(&answer)
		fmt.Println()
		var CheckBooks Books
		switch answer {
		case 1:

			ByteValue, err := ioutil.ReadAll(jsonFile)
			CheckError(err)

			err = json.Unmarshal(ByteValue, &CheckBooks)
			CheckError(err)

			for i, book := range CheckBooks.Books {
				fmt.Printf("%d. Book:\n", i+1)
				fmt.Printf("Name: %s\n", book.Name)
				fmt.Printf("Author: %s\n", book.Author)
				fmt.Printf("Amount of Pages: %s\n", book.AmountOfPages)
				fmt.Println()
			}

		case 2:
			ByteValue, err := ioutil.ReadAll(jsonFile)
			CheckError(err)

			fmt.Println("which book do you whant to take :")
			time.Sleep(time.Second)
			err = json.Unmarshal(ByteValue, &CheckBooks)
			CheckError(err)
			var count int
			for i, book := range CheckBooks.Books {
				count++
				fmt.Printf("%d. Book:\n", i+1)
				fmt.Printf("Name: %s\n", book.Name)
				fmt.Printf("Author: %s\n", book.Author)
				fmt.Printf("Amount of Pages: %s\n", book.AmountOfPages)
				fmt.Println()
			}
			var number int
			fmt.Scan(&number)
			if number > count {
				fmt.Println("number that you wrote not avaible ")
			} else {
				fmt.Println("here is your book")
				for i, book := range CheckBooks.Books {
					if i == number+1 {
						fmt.Printf("%d. Book:\n", i+1)
						fmt.Printf("Name: %s\n", book.Name)
						fmt.Printf("Author: %s\n", book.Author)
						fmt.Printf("Amount of Pages: %s\n", book.AmountOfPages)
						fmt.Println()
					}

				}
			}
		case 3:
			var SellBooks Books
			ByteValue, err := ioutil.ReadAll(jsonFile)
			CheckError(err)

			err = json.Unmarshal(ByteValue, &SellBooks)
			CheckError(err)

			fmt.Println("Name of the book:")
			var tt string
			tt, _ = reader.ReadString('\n')
			tt = strings.TrimSpace(tt)
			n, _ := reader.ReadString('\n') // Використовуємо reader для зчитування з пробілами
			n = strings.TrimSpace(n)

			fmt.Println("Author:")
			a, _ := reader.ReadString('\n') // Використовуємо reader для зчитування з пробілами
			a = strings.TrimSpace(a)

			fmt.Println("Amount of pages:")
			p, _ := reader.ReadString('\n') // Використовуємо reader для зчитування з пробілами
			p = strings.TrimSpace(p)

			NewBook := Book{
				Name:          n,
				Author:        a,
				AmountOfPages: p,
			}

			// adding a book to already existing slice of books
			SellBooks.Books = append(SellBooks.Books, NewBook)

			//marshaling existing JSON file
			updatedData, err := json.Marshal(SellBooks)
			CheckError(err)

			// adding new book to json file
			err = ioutil.WriteFile(filePath, updatedData, 0644)
			CheckError(err)

			fmt.Println("Book added successfully.")

		case 4:
			return

		default:
			fmt.Println("Something went wrong.")
		}
	}
}
