package main

import (
	"fmt"

	"github.com/coral/aubio-go"

	"github.com/gordonklaus/portaudio"
)

const (
	SampleRate = 48000
	Channels   = 1
	BufSize    = 512
	BlockSize  = 512
	Silence    = -70.0
	Threshold  = 0.3
)

func main() {
	fmt.Println("hello")

	//Init portaudio
	portaudio.Initialize()
	defer portaudio.Terminate()
	nSamples := 0

	in := make([]float32, BufSize)
	stream, err := portaudio.OpenDefaultStream(Channels, 0, SampleRate, len(in), in)
	chk(err)
	defer stream.Close()

	//Initiate Aubio
	ta := aubio.TempoOrDie(aubio.SpecDiff, uint(BufSize), uint(BlockSize), uint(SampleRate))
	ta.SetSilence(Silence)
	//ta.SetThreshold(Threshold)

	oa := aubio.OnsetOrDie(aubio.SpecDiff, uint(BufSize), uint(BlockSize), uint(SampleRate))
	oa.SetSilence(Silence)
	//oa.SetThreshold(-10.0)

	//ba := aubio

	chk(stream.Start())

	for {
		chk(stream.Read())
		nSamples += len(in)
		fixb := convertTo64(in)
		b := aubio.NewSimpleBufferData(BufSize, fixb)
		go processTempo(ta, b)
		go processOnset(oa, b)
		//go lel(in)
	}
}

func lel(b []float32) {
	fmt.Println(b)
}

func processTempo(ta *aubio.Tempo, b *aubio.SimpleBuffer) {
	ta.Do(b)
	for _, f := range ta.Buffer().Slice() {
		if f != 0 {
			fmt.Printf("Beat %.6f\n", f)
			fmt.Println(ta.GetBpm())
			fmt.Println(ta.GetConfidence())
		}
	}
}

func processOnset(oa *aubio.Onset, b *aubio.SimpleBuffer) {
	oa.Do(b)

	for _, f := range oa.Buffer().Slice() {
		if f != 0 {
			fmt.Printf("Onset %.6f\n", f)
		}
	}
}

func chk(err error) {
	if err != nil {
		panic(err)
	}
}

func convertTo64(ar []float32) []float64 {
	newar := make([]float64, len(ar))
	var v float32
	var i int
	for i, v = range ar {
		newar[i] = float64(v)
	}
	return newar
}
