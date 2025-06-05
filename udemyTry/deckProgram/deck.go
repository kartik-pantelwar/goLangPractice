package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

// declaring deck type
type deck []string

func (d deck) print() {
	for i, e := range d {
		fmt.Println(i, "=", e)
	}
}

func NewDeck() deck {
	d := deck{}
	suits := []string{"club", "diamond", "heart", "spade"}
	for _, s := range suits {
		for v := 2; v < 11; v++ {
			d = append(d, fmt.Sprintf("%s-%d", s, v))
		}
		faceCards := []string{"king", "queen", "jack", "ace"}
		for _, f := range faceCards {
			d = append(d, fmt.Sprintf("%s-%s", s, f))
		}
	}
	return d
}

func deal(d deck, i int) (deck, deck) {
	d1 := d[0:i] //0 is starting index
	d2 := d[i:]
	return d1, d2
}

func (d deck) saveToFile(filename string) error {
	dString := strings.Join(d, ",")
	dByte := []byte(dString)
	err := os.WriteFile((fmt.Sprintf("%s.txt", filename)), dByte, 0666)
	return err
}

func newDeckFromFile(filename string) deck {
	byteData, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}
	stringData := string(byteData)
	return strings.Split(stringData, ",")
}

func shuffleDeck(d deck) deck {
	l := len(d)
	for i := 0; i < l; i++ {
		random := rand.Intn(l)
		var temp string
		temp = d[i]
		d[i] = d[random]
		d[random] = temp
		
		//* Another method of swapping
		//d[i],d[random]= d[random],d[i]
	}
	return d
}
