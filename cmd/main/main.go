package main

import control_iq_takehome "githhub.com/morganhein/control-iq-takehome"

func main() {
	err := control_iq_takehome.Report("server_log.csv")
	if err != nil {
		panic(err)
	}
}
