package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func main() {
	data, err := ioutil.ReadFile("files/hard.json")
	if err != nil {
		fmt.Println("read file error:", err)
		return
	}

	var triangle [][]int
	err = json.Unmarshal(data, &triangle)
	if err != nil {
		fmt.Println("error unmarshal:", err)
		return
	}

	maxPath := 0

	for i := len(triangle) - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {

			if triangle[i+1][j] > triangle[i+1][j+1] {
				triangle[i][j] += triangle[i+1][j]
			} else {
				triangle[i][j] += triangle[i+1][j+1]
			}

			if maxPath < triangle[i][j] {
				maxPath = triangle[i][j]
			}
		}
	}
	fmt.Println("max result is:", maxPath)
}
