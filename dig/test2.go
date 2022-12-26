package main

import (
	"fmt"

	"go.uber.org/dig"
)

func TestName2() {
	type DSN struct {
		Addr string
	}
	c := dig.New()

	type DSNRev struct {
		dig.Out
		PrimaryDSN   *DSN `name:"primary"`
		SecondaryDSN *DSN `name:"secondary"`
	}
	p1 := func() (DSNRev, error) {
		return DSNRev{PrimaryDSN: &DSN{Addr: "Primary DSN"},
			SecondaryDSN: &DSN{Addr: "Secondary DSN"}}, nil
	}

	if err := c.Provide(p1); err != nil {
		fmt.Println(err)
	}

	type DBInfo struct {
		dig.In
		PrimaryDSN   *DSN `name:"primary"`
		SecondaryDSN *DSN `name:"secondary"`
	}
	inv1 := func(db DBInfo) {
		fmt.Println(db.PrimaryDSN)
		fmt.Println(db.SecondaryDSN)
	}

	if err := c.Invoke(inv1); err != nil {
		fmt.Println(err)
	}
}
