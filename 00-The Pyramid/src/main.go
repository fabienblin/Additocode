package main

import(
	"fmt"
	"os"
	"strconv"
	"bytes"
)

const FILL = 'H'
const LEFT = '/'
const RIGHT = '\\'
const SPACE = ' '

func main(){
	nbStage, err := strconv.Atoi(os.Args[1])
	if(err != nil){
		os.Exit(1)
	}


	var pyramidWidth = iterativeStageWidth(nbStage-1)
	var pyramid string

	for line := 0; line < nbStage * 4; line++ {
		pyramid += getSpaces(line, pyramidWidth) + getStructure(line) + "\n"
	}
	
	fmt.Print(pyramid);
	os.Exit(0)
}

func getSpaces(line int, pyramidWidth int) string {
	var spaces bytes.Buffer

	for i := 0; i < spaceLineWidth(line, pyramidWidth); i++ {
		spaces.WriteRune(SPACE)
	}

	return spaces.String()
}

func getStructure(line int) string {
	var structure bytes.Buffer

	for i := 0; i < structureLineWidth(line); i++ {
		if (i == 0){
			structure.WriteRune(LEFT)
		} else if (i == structureLineWidth(line) - 1) {
			structure.WriteRune(RIGHT)
		} else {
			structure.WriteRune(FILL)
		}
	}

	return structure.String()
}

func spaceLineWidth(line int, pyramidWidth int) int {
	return pyramidWidth / 2 - structureLineWidth(line) / 2
}

func structureLineWidth(line int) int {
	return 2 + (2 * line) + (4 * (line / 4))
}

func explicitStageWidth(n int) int {
	return 12 * n + 8
}

func recursiveStageWidth(n int) int {
	if (n == 0){
		return 8
	}

	return 12 + recursiveStageWidth(n - 1)
}

func iterativeStageWidth(n int) int {
	var width int = 8

	for i := 0; i < n; i++ {
		width += 12
	}

	return width
}