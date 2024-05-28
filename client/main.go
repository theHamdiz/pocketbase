package main

import "github.com/theHamdiz/pocketbase"

func main() {
	pb := pocketbase.New()
	err := pb.Start()
	if err != nil {
		return
	}
}
