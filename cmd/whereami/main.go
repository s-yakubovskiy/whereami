package main

import (
	"github.com/s-yakubovskiy/whereami/cmd/whereami/cmd"
	"github.com/s-yakubovskiy/whereami/config"
)

var Version string

func init() {
	config.Init()
}

// TODO:
// 1. Update documentation and how-to install guide. Also how-to setup IP Quality with api key
// 2. Adding experimental flag --gps to add posibility enrich whrmi location output with precise lat & long, altimeter and some other optional stuff
// 3. curl -v http://api.airvisual.com/v2/nearest_city\?lat=59.655065167\&lon=56.787904833\&key=API_KEY

func main() {
	// fmt.Printf("%+v\n", config.Cfg)
	// os.Exit(0)
	cmd.Execute()
}
