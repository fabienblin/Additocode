package burn

import(
	"time"
	"math/rand"
)

const WEED = '*'
const FIRE = '$'
const ASHES = '~'

var Env Environment

type Environment struct {
	PropagationRate float64
	Field Field
	Fires [] Position
}

func (env *Environment) InitEnvironment(propagationRate float64, field Field, fires [] Position) Environment {
	env.PropagationRate = propagationRate
	env.Field = field
	env.Fires = fires

	rand.Seed(time.Now().UTC().UnixNano())
	
	return *env
}

func (env *Environment) String() string {
	var s string = ""
	for _, row := range env.Field.Grid {
		s += string(row) + "\n"
	}
	return s
}