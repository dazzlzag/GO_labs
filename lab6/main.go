package main

import "fmt"

type contract struct {
	ID     int
	Number string
	Date   string
}
type contacts struct {
	id           int
	Addss, Phone string
}
type user struct {
	ID   int
	Name string
	contacts
}
type employee struct {
	ID   int
	Name string
	contacts
}

func main() {
	// 1
	cont := contract{ID: 1,
		Number: "#000A\\n101",
		Date:   "2024-01-31",
	}

	fmt.Printf("%+v\n", cont)
	//fmt.Printf("{ID:%d Number:%s Date:%s}", cont2.ID, cont2.Number, cont2.Date)
	// 2
	cont2 := contract{ID: 1,
		Number: "#000A101\t01",
		Date:   "2024-01-31",
	}
	fmt.Printf("%+v\n", cont2)
	//fmt.Printf("{ID:%d Number:%s Date:%s}", cont2.ID, cont2.Number, cont2.Date)
	// 3
	cont3 := contract{ID: 1,
		Number: "#000A101\\n01",
		Date:   "2024-01-31",
	}
	fmt.Println(contractPrint(cont3))

	// 4
	u := user{
		ID:   0,
		Name: "name",
		contacts: contacts{
			Addss: "addr",
			Phone: "phone",
		},
	}
	e := employee{
		ID:   0,
		Name: "name1",
		contacts: contacts{
			Addss: "addr1",
			Phone: "phone1",
		},
	}
	fmt.Println(u.Addss, u.Phone, e.Addss, e.Phone)
}
func contractPrint(s contract) string {
	return fmt.Sprintf("ДоговорNo %s от %s", s.Number, s.Date)
}
