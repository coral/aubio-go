package aubio

/*
#cgo LDFLAGS: -laubio
#include <aubio/aubio.h>
#include <aubio/musicutils.h>
*/
import "C"

// Tempo is a wrapper for the aubio_tempo_t tempo detection object.
type MusicUtils struct {
	o   *C.aubio_onset_t
	buf *SimpleBuffer
}

//uint_t aubio_silence_detection (const fvec_t * v, smpl_t threshold);

func Do(input *SimpleBuffer) {

	//C.aubio_silence_detection(t.o, input.vec, t.buf.vec)
}
