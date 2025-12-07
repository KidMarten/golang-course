package main

import (
	"fmt"
)

type bookmarkMap = map[string]string

func main() {

	bookmarks := make(bookmarkMap)
Menu:
	for {
		input := getMenu()
		switch input {
		case 1:
			printBookmarks(bookmarks)
		case 2:
			addBookmark(bookmarks)
		case 3:
			deleteBookmark(bookmarks)
		case 4:
			break Menu
		}
	}
}

func getMenu() int {
	var input int
	fmt.Println("1. List bookmarks")
	fmt.Println("2. Add")
	fmt.Println("3. Delete")
	fmt.Println("4. Exit")
	fmt.Scan(&input)
	return input
}

func printBookmarks(bookmarks bookmarkMap) {
	if len(bookmarks) == 0 {
		fmt.Println("No bookmarks yet")
	}
	for key, value := range bookmarks {
		fmt.Println(key, ":", value)
	}
}

func addBookmark(bookmarks bookmarkMap) {
	var newBookmarkKey string
	var newBookmarkValue string

	fmt.Print("Enter name: ")
	fmt.Scan(&newBookmarkKey)
	fmt.Print("Enter value: ")
	fmt.Scan(&newBookmarkValue)

	bookmarks[newBookmarkKey] = newBookmarkValue
}

func deleteBookmark(bookmarks bookmarkMap) {
	var bookmarkToDelete string
	fmt.Print("Enter name to delete: ")
	fmt.Scan(&bookmarkToDelete)
	delete(bookmarks, bookmarkToDelete)
}
