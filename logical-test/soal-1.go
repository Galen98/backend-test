package main

import (
	"fmt"
	"strings"
)

func balikKata(kata string) string {
	runes := []rune(kata)
	n := len(runes)
	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}
	return string(runes)
}

func ubahKalimat(kalimat string) string {
	katas := strings.Split(kalimat, " ")
	for i, kata := range katas {
		katas[i] = balikKata(kata)
	}
	return strings.Join(katas, " ")
}

func main() {
	kalimat1 := "italem irad irigayaj"
	kalimat2 := "iadab itsap ulalreb"
	kalimat3 := "nalub kusutret gnalali"
	fmt.Println(ubahKalimat(kalimat1))
	fmt.Println(ubahKalimat(kalimat2))
	fmt.Println(ubahKalimat(kalimat3))
}
