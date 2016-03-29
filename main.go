package main

import "fmt"
import "github.com/tsandall/vertest/version"

func main() {
	fmt.Printf("Goodbye Cruel World (%v-%v)", version.Version, version.Vcs)
}
