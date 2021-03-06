// test2 project main.go
package main

import (
	"github.com/Ariemeth/mech"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	mech1 := mech.NewMech("Mech1")
	mech2 := mech.NewMech("Mech2")

	weapon1 := mech.Weapon{Range: 5, Damage: 1}
	weapon2 := mech.Weapon{Range: 3, Damage: 3}

	mech1.AddWeapon(weapon1)
	mech2.AddWeapon(weapon2)

	mech1.Fire(4, mech2)

}
