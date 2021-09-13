package main

import "fmt"

type class struct {
	Id int
	Student int
}

func main()  {
	c := &class{2, 5}
	m := make(map[string]*class)
	m["c"] = c
	fmt.Println(m)
	//c2 := m["c"]
	m["c"].Id = 10
	//c2.Id = 10
	fmt.Println(*c)
	fmt.Println(m)
	c.Id = 20
	fmt.Println(*c)
	fmt.Println(*m["c"])
}
