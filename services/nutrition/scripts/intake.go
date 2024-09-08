package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
)

func shagram(s string) []int {
	result := make([]int, 26)
	for i := range s {
		result[int(s[i])-int('a')]++
	}
	return result
}

func kShagram(s string, k int) []int {
	shagram0 := shagram(s)
	shagramk := make([]int, 26)
	for i := range shagram0 {
		shagramk[(i+k)%26] = shagram0[i]
	}
	return shagramk
}

func getWeight(s string) int {
	r := shagram(s)
	result := 0
	for i, value := range r {
		result += i * value
	}
	return result
}

func main() {
	file, err := os.Open("selected_strings.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	defer func() {
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}()
	// optionally, resize scanner's capacity for lines over 64K, see next example
	a := make([]string, 0)
	for scanner.Scan() {
		cur := scanner.Text()
		a = append(a, cur)
	}
	// fmt.Println(a)
	fmt.Println(kShagram("gather", 13))
	fmt.Println(shagram("urgent"))
	fmt.Println(reflect.DeepEqual(kShagram("gather", 13), shagram("urgent")))

	for k := 1; k <= 26; k++ {
		if reflect.DeepEqual(kShagram("accepts", k), shagram("courage")) {
			fmt.Println("true: ", k)
		}
	}
	fmt.Println("false")

	fmt.Println(kShagram("while", 4))

	maxWeightString := "a"
	for _, item := range a {
		if reflect.DeepEqual(kShagram(item, 4), (shagram("while"))) {
			fmt.Println(item)
		}
		if len(item) != 10 {
			continue
		}

		if getWeight(maxWeightString) < getWeight(item) {
			maxWeightString = item
		}
	}
	fmt.Println(maxWeightString, getWeight(maxWeightString))

	maxWeightString = "a"
	for i := range a {
		for j := range a {
			if i >= j {
				continue
			}
			if len(a[i]) != len(a[j]) {
				continue
			}
			for k := 0; k < 26; k++ {
				if reflect.DeepEqual(kShagram(a[i], k), shagram(a[j])) {
					if getWeight(a[i]) > getWeight(maxWeightString) {
						maxWeightString = a[i]
					}
					if getWeight(a[j]) > getWeight(maxWeightString) {
						maxWeightString = a[j]
					}
					break
				}
			}
		}
	}
	fmt.Println(getWeight(maxWeightString))
	fmt.Println(maxWeightString)
}
