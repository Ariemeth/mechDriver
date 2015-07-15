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

	mech1 := mech.NewMech(mech.Weapon{Range: 5, Damage: 1}, "Mech1")
	mech2 := mech.NewMech(mech.Weapon{Range: 3, Damage: 3}, "Mech2")

	mech1.Fire(2, mech2)
	
//	mech3 := mech.Mech{Weapons:mech.Weapon{Range: 5, Damage: 1}, Name:"testMech3"}
//	mech3.Fire(mech1)
}
