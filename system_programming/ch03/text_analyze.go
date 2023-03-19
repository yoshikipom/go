package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

var source = `line 1
line 2
line 3
`

var source2 = "123 1.234 1.0e4 test"

func textAnalyze() {
	// reader := bufio.NewReader(strings.NewReader(source))
	// for {
	// 	line, err := reader.ReadString('\n')
	// 	fmt.Printf("%#v\n", line)
	// 	if err == io.EOF {
	// 		break
	// 	}
	// }

	scanner := bufio.NewScanner(strings.NewReader(source))
	for scanner.Scan() {
		fmt.Printf("%#v\n", scanner.Text())
	}
}

func fscan() {
	reader := strings.NewReader(source2)
	var i int
	var f, g float64
	var s string
	fmt.Fscan(reader, &i, &f, &g, &s)
	fmt.Printf("i=%#v f=%#v g=%#v s=%#v\n", i, f, g, s)
}

var csvSource = `
13101,"100 ","1000003","トウキョウト","チヨダク","ヒトツバシ(1チョウメ)","東京都","千代田区"," 一ツ橋(1丁目)",1,0,1,0,0,0 
13101,"101 ","1010003","トウキョウト","チヨダク","ヒトツバシ(2チョウメ)","東京都","千代田区"," 一ツ橋(2丁目)",1,0,1,0,0,0 
13101,"100 ","1000012","トウキョウト","チヨダク","ヒビヤコウエン","東京都","千代田区"," 日比谷公園 ",0,0,0,0,0,0 
13101,"102 ","1020093","トウキョウト","チヨダク","ヒラカワチョウ","東京都","千代田区","平河町",0,0,1,0,0,0 
13101,"102 ","1020071","トウキョウト","チヨダク","フジミ","東京都","千代田区","富士見",0,0,1,0,0,0
`

func csvRead() {
	reader := strings.NewReader(csvSource)
	csvReader := csv.NewReader(reader)
	for {
		line, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		fmt.Println(line[2], line[6:9])
	}
}
