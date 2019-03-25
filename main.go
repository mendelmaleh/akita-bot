package main

import (
	"github.com/mendelmaleh/akita"
	_ "github.com/mendelmaleh/akita/commands/start"
	_ "github.com/mendelmaleh/akita/commands/help"
	_ "github.com/mendelmaleh/akita/commands/cancel"
)

func main() {
	a := akita.NewAkita("")
	a.API.Debug = true
	a.Start()
}
