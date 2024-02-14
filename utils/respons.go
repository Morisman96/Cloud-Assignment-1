package utils

import (
	_ "encoding/json"
	_ "fmt"
	_ "log"
	_ "net/http"
	_ "os"
	_ "strconv"
)

type BookLanguage struct {
	Language string
	Books    int
	Authors  int
	Fraction float64
}

type BookLangReaders struct {
	Country    string
	IsoCode    int
	Books      int
	Authors    int
	Readership int
}

func GetBookLanguage() {

}

func GetBookLangReaders() {

}
