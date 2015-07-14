// test2 project main.go
package main

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/Ariemeth/mech"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	f, err := os.Open("C:\\Users\\Edward\\OneDrive\\My_Projects\\Go\\src\\github.com\\Ariemeth\\mechDriver\\data.txt")
	check(err)

	b1 := make([]byte, 25)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: \n%s", n1, b1)

	//	r4 := bufio.NewReader(f)
	//	b4, err := r4.ReadBytes('\n')
	//	fmt.Printf("reader:%s \n %d bytes read", b1, b4)

	f.Close()
	mech1 := mech.NewMech(mech.Weapon{Range: 5, Damage: 1}, "Mech1")
	mech2 := mech.NewMech(mech.Weapon{Range: 3, Damage: 3}, "Mech2")

	mech1.Weapons.Fire(2, mech2)

	buf := new(bytes.Buffer)
	var data = []interface{}{
		uint16(61374),
		int8(-54),
		uint8(254),
	}
	for _, v := range data {
		err := binary.Write(buf, binary.LittleEndian, v)
		if err != nil {
			fmt.Println("binary.Write failed:", err)
		}
	}
	fmt.Printf("%d\n", buf.Bytes())

	type SmallWeapon struct {
		Name   string
		Damage string
		Range  string
	}

	type SmallMech struct {
		Name      string
		Structure int `json:"structure"`
		Weapons   []SmallWeapon
	}

	type Configuration struct {
		Mechs []string
		HP    string
	}

	file, _ := os.Open("conf.json")
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err2 := decoder.Decode(&configuration)
	if err2 != nil {
		fmt.Println("error:", err2)
	}
	fmt.Println(configuration.Mechs)

	///mechHP1 int
	mechHP1, err3 := strconv.ParseInt(configuration.HP, 0, 0)
	if err3 == nil {
		fmt.Println("HP:", mechHP1)
	}

	smallMech := SmallMech{}
	//	weapons := []SmallWeapon{}
	//	smallMech.Weapons = [2]SmallWeapon{}
	smallMech.Weapons = append(smallMech.Weapons, SmallWeapon{})
	smallMech.Weapons = append(smallMech.Weapons, SmallWeapon{})

	b, err5 := json.Marshal(&smallMech)
	if err5 == nil {
		fmt.Println(b)
	} else {
		fmt.Println("error:", err5)
	}
	//	n := bytes.Index(b, []byte{0})
	s := string(b[:])
	//	os.Stdout.Write(b)
	fmt.Println(s)

	//file2, _ := os.Open("test.json")
	//encoder := json.NewDecoder(file2)
	//smallMech := SmallMech{}
	//err4 := encoder.Decode(&smallMech)
	//if err4 != nil {
	//	fmt.Println("error:", err4)
	//}

}
