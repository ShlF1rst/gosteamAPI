package main

import (
	"fmt"
	"steamAPI"
)

func main() {
	client := steamAPI.GetClient()
	interfaces, _ := client.GetSupportedAPIList()
	for _, inter := range interfaces {
		if inter.Name == "ISteamWebAPIUtil" {
			fmt.Println(inter)
		}
	}
	si, _ := client.GetServerInfo()
	fmt.Println(si)
}
