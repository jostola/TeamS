package bfrequence

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"unicode"
)

func LinjeTeller() {

	// Åpner og leser fra fil.
	linjeTeller := 0

	fil, err := os.Open("../Frequence/testtekst.txt")
	if err != nil {
		log.Fatal(err)
	}
	//buffer := bufio.NewReader(fil)
	scanner := bufio.NewScanner(fil)
	for scanner.Scan() {
		linjeTeller++

	}
	fmt.Println("Filen inneholder", linjeTeller, "linjer")
}

func RuneTeller() {
	// Hente inn tekstfilen som runer i en array-buffer.
	tekstFil, err := os.Open("../Frequence/testtekst.txt")
	if err != nil {
		log.Fatal(err)
	}

	bs := bufio.new()(tekstFil)
	m := make(map[rune]int)
	for _, r := range bs {
		m[r]++
	}
	// answer is now in m.  sort and format output:
	lfs := make(lfList, 0, len(m))
	for l, f := range m {
		lfs = append(lfs, &letterFreq{l, f})
	}
	sort.Sort(lfs)
	fmt.Println("rune frequency")
	teller := 0
	for _, lf := range lfs {
		if teller >= 5 {
			break
		}
		if unicode.IsGraphic(lf.rune) {
			fmt.Printf("   %c    %7d\n", lf.rune, lf.freq)

		} else {
			fmt.Printf("%U  %7d\n", lf.rune, lf.freq)
		}
		teller++
	}
}

type letterFreq struct {
	rune
	freq int
}
type lfList []*letterFreq

func (lfs lfList) Len() int { return len(lfs) }

func (lfs lfList) Less(i, j int) bool {
	switch fd := lfs[i].freq - lfs[j].freq; {
	case fd < 0:
		return false
	case fd > 0:
		return true
	}
	return lfs[i].rune < lfs[j].rune
}
func (lfs lfList) Swap(i, j int) {
	lfs[i], lfs[j] = lfs[j], lfs[i]
}
