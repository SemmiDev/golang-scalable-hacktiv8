package util

import (
	"encoding/csv"
	"os"
)

type Member struct {
	Nama, VoucherCode, Batch string
}

func createMembers(data [][]string) []Member {
	var members []Member
	// skip the first row because it's the header
	for _, row := range data[1:] {
		members = append(members, Member{
			Nama:        row[0],
			VoucherCode: row[1],
			Batch:       row[2],
		})
	}
	return members
}

func LoadMembers(filename string) ([]Member, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}

	members := createMembers(data)
	return members, nil
}

type FilterFunc func(Member) bool

// 604 - 543 = 61

func FilterMembers(members []Member, mustFilters []FilterFunc) []Member {
	var filteredMembers []Member
	for _, v := range members {
		status := true
		for _, filter := range mustFilters {
			if !filter(v) {
				status = false
				break
			}
		}
		if status {
			filteredMembers = append(filteredMembers, v)
		}
	}
	return filteredMembers
}
