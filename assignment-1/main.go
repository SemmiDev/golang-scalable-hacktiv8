package main

import (
	"assignment-1/util"
	"fmt"
	"log"
	"os"
	"strconv"
)

// friend is a struct that contains friend's details
type friend struct {
	No      uint8
	Name    string
	Address string
	Job     string
	Reason  string // the reason for choosing golang class
}

// friends is a struct that contains slice of friend
type friends struct {
	friends map[uint8]friend
}

// printFriendDetails prints friend's details to stdout
func printFriendDetails(data friends, id uint8) {
	friend, exist := data.friends[id]
	if !exist {
		fmt.Printf("friend with id %d doesn't exists\n", id)
		return
	}

	fmt.Printf("Nama\t\t\t\t: %s\n", friend.Name)
	fmt.Printf("Alamat\t\t\t\t: %s\n", friend.Address)
	fmt.Printf("Pekerjaan\t\t\t: %s\n", friend.Job)
	fmt.Printf("Alasan memilih kelas Golang\t: %s\n", friend.Reason)
}

// getFriendNoFromArgs gets friend's id from args
func getFriendNoFromArgs(index int, fallback int) uint8 {
	if os.Args[index] != "" {
		dataInInt, err := strconv.Atoi(os.Args[index])
		if err != nil {
			return uint8(fallback)
		}
		return uint8(dataInInt)
	}
	return uint8(fallback)
}

var friendsData = friends{
	friends: map[uint8]friend{
		1: {
			No:      1,
			Name:    "Sammi Aldhi Yanto",
			Address: "Padang",
			Job:     "Student",
			Reason:  "Karena ingin jago backend dengan golang",
		},
	},
}

func init() {
	members, err := util.LoadMembers("members.csv")
	if err != nil {
		log.Fatal(err)
	}

	filter := func(m util.Member) bool {
		return m.Batch == "10" && m.Name != "SAMMI ALDHI YANTO"
	}

	filteredMembers := util.FilterMembers(members, filter)

	var no uint8 = 2
	for _, v := range filteredMembers {
		friendsData.friends[no] = friend{
			No:      no,
			Name:    v.Name,
			Address: "-",
			Job:     "-",
			Reason:  "-",
		}
		no++
	}
}

func main() {
	friendID := getFriendNoFromArgs(1, 1)
	fmt.Print("\033[H\033[2J") // clear screen
	printFriendDetails(friendsData, friendID)
}
