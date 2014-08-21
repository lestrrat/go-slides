package main

import (
	"fmt"
	"log"
	"strconv"
)

func hoge() error {
	return nil
}

// START SIMPLE ERROR SAMPLE OMIT
func simpleError() {
	err := hoge()
	if err != nil {
		fmt.Printf("Error! %s", err)
		return
	}
	// 他になにか処理…
}
// END SIMPLE ERROR SAMPLE OMIT

// START STRCONV SAMPLE OMIT
func sampleParseInt(s string) (int64, error) {
	i, err := strconv.ParseInt(s, 10, 64) // HL
	if err != nil {                       // HL
		return 0, err
	}

	return i, nil
}
// END STRCONV SAMPLE OMIT

// START CHECK ERROR SAMPLE OMIT
func runParseInt() {
	for _, s := range []string{"100", "10XXX"} {
		i, err := sampleParseInt(s) // HL
		if err != nil {             // HL
			log.Printf("Failed to parse: %s", err)
			continue
		}

		log.Printf("parsed '%d'", i)
	}
}
// END CHECK ERROR SAMPLE OMIT

// START ERROR OMIT
type stringError string // 実体はただの文字列なエラー

func (se stringError) Error() string { // HL
	return string(se)
}

type structError struct { // 実体が構造体なエラー
	// Fields...
}

func (se structError) Error() string { // HL
	return fmt.Sprintf("error: ....")
}

// END ERROR OMIT
func mayReturnError() error {
	return nil
}

func usingErrorf() error {
	if err := mayReturnError(); err != nil {
		return fmt.Errorf("something went wrong: %s", err)
	}
	return nil
}
