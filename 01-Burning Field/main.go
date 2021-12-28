package main

import(
	"fmt"
	"flag"
	"regexp"
	"os"
	"strconv"
	. "new/burn"
)

var propagationFlag = flag.Float64("p", 0.5, "A propagation probability between 0 and 1")
var fieldFlag = flag.String("field", "[0,0]", "The field's dimmensions [y, x]")
var fireFlag = flag.String("fire", "", "The initial fire positions ...[y, x]")

func main(){
	// parse input
	parseInput()

	// initialize field
	initialize()
	

	
	// simulate fire propagation
	var userInput string
	for (true) {
		fmt.Println(Env.String())
		fmt.Println("Press ENTER to process time cycle.")
		fmt.Scanln(&userInput)
		Env.Field.SimulateFires()
	}
}

func parseInput() {
	flag.Parse()

	fieldRegexp := `\[[0-9]+,[0-9]+\]`
	fireRegexp := `\[[0-9]+,[0-9]+\]*`

	field := regexp.MustCompile(fieldRegexp)
	if (!field.MatchString(*fieldFlag)) {
		os.Exit(1)
	}

	fire := regexp.MustCompile(fireRegexp)
	if (!fire.MatchString(*fireFlag)) {
		os.Exit(1)
	}
}

func initialize() {
	// init field
	fieldCompiled := regexp.MustCompile(`[0-9]+`)
	fieldFind := fieldCompiled.FindAll([]byte(*fieldFlag), -1)
	fieldY, errY := strconv.Atoi(string(fieldFind[0]))
	fieldX, errX := strconv.Atoi(string(fieldFind[1]))
	if (errY != nil && errX != nil ||
		fieldY < 1 ||
		fieldX < 1) {
		fmt.Println("Invalid field")
		os.Exit(1)
	}
	field := new(Field).InitField(fieldY, fieldX)

	// init fires
	fireCompiled := regexp.MustCompile(`[0-9]+`)
	fireFind := fireCompiled.FindAll([]byte(*fireFlag), -1)
	var fires [] Position
	for i := 0; i < len(fireFind); i+=2 {
		fireY, errY := strconv.Atoi(string(fireFind[i]))
		fireX, errX := strconv.Atoi(string(fireFind[i+1]))
		if (errY != nil && errX != nil ||
			fireY < 0 ||
			fireX < 0) {
			fmt.Println("Invalid position")
		}
		fires = append(fires, new(Position).InitPosition(fireY, fireX))
	}

	Env = new(Environment).InitEnvironment(*propagationFlag, field, fires)
	Env.Field.InitWeed()
	Env.Field.InitFires(Env.Fires)
}
