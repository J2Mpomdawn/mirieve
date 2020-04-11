package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/360EntSecGroup-Skylar/excelize"
)

//where to put the scanned text
var linesbox []string = make([]string, 0)
var lines *[]string = &linesbox

func cut(filename string) {

	//open file
	file, err := os.Open(filename)

	//close after use
	defer file.Close()

	//error handling
	if err != nil {
		fmt.Println("can't open the file :\\")
	}

	//prepare scanner
	scanner := bufio.NewScanner(file)

	//read and pack
	for scanner.Scan() {
		*lines = append(*lines, scanner.Text())
	}

	//remove before comma
	for i, line := range *lines {
		for j, str := range line {
			if str == 44 {
				(*lines)[i] = (*lines)[i][j+1:]
			}
		}
	}

	//box for what converted to byte array
	bytelinesbox := make([][]byte, 0)
	blsbp := &bytelinesbox

	//connect elements while putting indention
	for _, element := range *lines {
		for _, ele := range element {
			*blsbp = append(*blsbp, []byte(string(ele)))
		}
		*blsbp = append(*blsbp, []byte{10})
	}

	//connect elements of bytelinesbox
	bytelinebox := make([]byte, 0)
	blbp := &bytelinebox

	for _, e := range *blsbp {
		*blbp = append(*blbp, e...)
	}

	//close once
	file.Close()

	//open again for writing
	file, err = os.Create(filename)

	//error handling
	if err != nil {
		fmt.Println("can't create the file")
	}

	//rewrite
	file.Write(*blbp)
}

func exceloperare(eventname string) {

	//exel-file open
	f, err := excelize.OpenFile("C:/Users/ryogen/OneDrive/ボーダー推移.xlsx")

	//error handling
	if err != nil {
		fmt.Println(err)
	}

	//get recorded events
	rows := f.GetRows("2500")
	row := rows[0]

	//next to last column
	next := len(row) + 1

	//byte array conversion for convert next to column alphabet
	colbyte := make([]byte, 0, 6)
	cbp := &colbyte

	//first digit
	fst := next % 26
	if fst == 0 {
		fst = 26
	}
	fb := byte(fst + 64)
	*cbp = append(*cbp, fb)

	//second digit
	scd1 := next / 26
	if fst == 26 {
		scd1--
	}
	scd2 := scd1
	if next > 26 {
		scd2 = scd2 % 26
		if scd2 == 0 {
			scd2 = 26
		}
	}
	sb := byte(scd2 + 64)
	if sb != 64 {
		*cbp, (*cbp)[0] = append((*cbp)[:1], (*cbp)[0:]...), sb
	}

	//third digit
	trd := scd1 / 26
	if scd1 == 26 {
		trd--
	}
	tb := byte(trd + 64)
	if tb != 64 {
		*cbp, (*cbp)[0] = append((*cbp)[:1], (*cbp)[0:]...), tb
	}

	//convert byte array to string
	col := string(*cbp)

	//record the event name
	f.SetCellValue("2500", col+"1", eventname)

	//record borders
	f.SetCellValue("2500", col+"2", 0)
	i := 2
	for _, border := range *lines {
		i++
		f.SetCellValue("2500", fmt.Sprintf("%s%d", col, i), border)
	}

	//save changes
	f.Save()
}

func main() {
	filename := ""
	print("イベント名を入力 → ")
	fmt.Scan(&filename)
	path := "C:/Users/ryogen/Desktop/ミリグラ/" + filename + ".txt"
	cut(path)
	exceloperare(filename)
}
