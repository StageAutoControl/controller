package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/apinnecke/dmx-auto-control/metronome"
	"github.com/apinnecke/dmx-auto-control/output"
)

func main() {
	sig := getSignal()
	//out := output.NewBufferOutput(os.Stdout)
	out := output.NewAudioOutput(1776, 1332)
	if err := out.Start(); err != nil {
		panic(err)
	}

	m := metronome.NewPlayer(out)
	b := metronome.NewBar(4, 4, 120)

	//if err := m.PlayBarUntilSignalOrLimit(b, sig, 4); err != nil {
	if err := m.PlayBarUntilSignal(b, sig); err != nil {
		fmt.Println(err)
	}
}

func getSignal() chan os.Signal {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, os.Kill)
	return sig
}
