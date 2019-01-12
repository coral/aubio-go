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
	BlockSize  = 256
	Silence    = -90.0
	Threshold  = -10.0
)

func main() {
	fmt.Println("hello")

	//Init portaudio
	portaudio.Initialize()
	defer portaudio.Terminate()

	in := make([]float32, BufSize)
	stream, err := portaudio.OpenDefaultStream(Channels, 0, SampleRate, len(in), in)
	chk(err)
	defer stream.Close()

	//Initiate Aubio
	ta := aubio.TempoOrDie(aubio.SpecDiff, uint(BufSize), uint(BlockSize), uint(SampleRate))
	ta.SetSilence(Silence)
	ta.SetThreshold(Threshold)

	//ch := make(chan float64)

	chk(stream.Start())

	for {
		chk(stream.Read())
		fixb := convertTo64(in)
		b := aubio.NewSimpleBufferData(BufSize, fixb)
		ta.Do(b)
		for _, f := range ta.Buffer().Slice() {
			if f != 0 {
				fmt.Printf("Beat %.6f\n", f)
				fmt.Println(ta.GetBpm())
				fmt.Println(ta.GetConfidence())
			}
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
