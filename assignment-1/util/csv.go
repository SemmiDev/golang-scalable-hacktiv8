package util

import (
	"encoding/csv"
	"os"
)

type Member struct {
	Name, VoucherCode, Batch string
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

func createMembers(data [][]string) []Member {
	members := make([]Member, 0, len(data))
	for _, row := range data[1:] {
		members = append(members, Member{
			Name:        row[0],
			VoucherCode: row[1],
			Batch:       row[2],
		})
	}
	return members
}

type FilterFunc func(Member) bool

func FilterMembers(members []Member, filters FilterFunc) []Member {
	var filteredMembers []Member
	for _, v := range members {
		if filters(v) {
			filteredMembers = append(filteredMembers, v)
		}
	}
	return filteredMembers
}
