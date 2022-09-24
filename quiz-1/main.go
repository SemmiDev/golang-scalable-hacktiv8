package main

import (
	"fmt"
	"sort"
)

// fruits => ["apel", "mangga", "apel", "jeruk", "mangga", "jeruk", "apel", "mangga", "mangga"]
// masukin ke maps, key = namaBuah, value = jumlah buah => {"apel": 3, "mangga": 4, "jeruk": 2}
// bkin slice yg values nya cuman keys dari maps -> ["apel", "mangga", "jeruk"]
// sort berdasarkan value dari map, menggunakan fungsi sort.Slice / sort.SliceTable yg ada di package sort
// jadi sekarang slice nya => ["mangga", "apel", "jeruk"] => dari buah yang paling banyak ke yang paling sedikit
// print data di maps berdasarkan key dari value slice tsb, otomatis akan berurutan dari yang paling banyak ke yang paling sedikit
// referensi : https://zetcode.com/golang/sort/

func SortFruits(fruits []string) (sortedFruitsWithCounts []string) {
	fruitsCounts := make(map[string]int)
	for _, fruit := range fruits {
		fruitsCounts[fruit]++
	}

	keys := make([]string, 0, len(fruitsCounts))
	for key := range fruitsCounts {
		keys = append(keys, key)
	}
	// keys = ["apel", "mangga", "jeruk"]

	// sort by value
	sort.Slice(keys, func(i, j int) bool {
		return fruitsCounts[keys[i]] > fruitsCounts[keys[j]]
	})
	// keys = ["mangga", "apel", "jeruk"]

	/*
		sample output
		namaBuah counts

		mangga 4
		apel 3
		jeruk 2
	*/

	for i := 0; i < len(keys); i++ {
		keys[i] = keys[i] + " " + fmt.Sprint(fruitsCounts[keys[i]])
	}

	return keys
}
