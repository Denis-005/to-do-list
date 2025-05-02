package main

import (
	"fmt"
	"sort"
	"strings"
)

type Achievement struct {
	Name    string
	Scores  int
	Rewards []string
}

func main() {
	var size int
	fmt.Scan(&size)
	players := make(map[string]Achievement)

	for i := 0; i < size; i++ {
		var name string
		var scores int
		fmt.Scan(&name, &scores)

		player := players[name]
		player.Scores += scores

		if player.Scores >= 10 && !contains(player.Rewards, "junior") {
			player.Rewards = append(player.Rewards, "junior")
		}

		if player.Scores >= 50 && !contains(player.Rewards, "middle") {
			player.Rewards = append(player.Rewards, "middle")
		}

		if player.Scores >= 100 && !contains(player.Rewards, "senior") {
			player.Rewards = append(player.Rewards, "senior")
		}

		players[name] = player
	}

	keys := make([]string, 0, len(players))
	for k := range players {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, name := range keys {
		player := players[name]
		fmt.Printf("%s [%s]\n", name, strings.Join(player.Rewards, " "))
	}
}

func contains(slice []string, value string) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}
