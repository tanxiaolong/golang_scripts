package main

import "github.com/albshin/go-pubg"
import "fmt"

func main() {
	client, err := pubg.New("EE94F226013FC988B8978D8320611986")
	if err != nil {
		fmt.Println(err)
	}
	info, err := client.GetPlayer("EATUTONIGHT") // Returns JSON unfiltered for player "JohnDoe"
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", info)

	steaminfo, err := client.GetSteamInfo("76561198383600183") // Returns Steam information for user based on SteamId
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", steaminfo)
}
