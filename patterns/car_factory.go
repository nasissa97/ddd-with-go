package patterns

import (
	"errors"
	"fmt"
	"strings"
)

type Car interface {
	Honk()
}

type BMW struct {
	heatedSeatSubscriptionEnabled bool
}

func (b BMW) Honk() {
	fmt.Println("BEEEEEEEP")
}

type Tesla struct {
	autoPilotEnabled bool
}

func (t Tesla) Honk() {
	fmt.Println("EEEEELLLLLLLOOOOONNNN")
}

func BuildCar(carType string) (Car, error) {
	carType = strings.ToLower(carType)
	switch carType {
	case "bmw":
		return BMW{heatedSeatSubscriptionEnabled: true}, nil
	case "tesla":
		return Tesla{autoPilotEnabled: true}, nil
	default:
		return nil, errors.New("unknow car type")
	}
}
