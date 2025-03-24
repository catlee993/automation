package main

import (
	sort_service "automation/pkg"
	"fmt"
)

func main() {
	fmt.Println(fmt.Sprintf("A standard pkg: %s", sort_service.Sort(149, 149, 149, 19.9)))
	fmt.Println(fmt.Sprintf("A special pkg: %s", sort_service.Sort(200, 300, 400, 10)))
	fmt.Println(fmt.Sprintf("A rejected pkg: %s", sort_service.Sort(150, 150, 150, 20.01)))
	return
}
