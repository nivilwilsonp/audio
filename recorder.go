package audio

type AudioFormat struct{}

type State int

const (
	StoppedState   State = iota //The recorder is not active.
	RecordingState              //The recording is requested.
	PausedState                 //The recorder is paused
)

type Quality int

const (
	VeryLowQuality Quality = iota
	LowQuality
	NormalQuality
	HighQuality
	VeryHighQuality
)

type Errors int

const (
	NoError             Errors = iota //No Errors.
	ResourceError                     //Device is not ready or not available.
	FormatError                       //Current format is not supported.
	OutOfSpaceError                   //No space left on device.
	LocationNotWritable               //The output location is not writable.
)

type Recorder struct {
	format AudioFormat
	state  State
	
}

func (recorder *Recorder) StartRecording()  {}
func (recorder *Recorder) StopRecording()   {}
func (recorder *Recorder) PauseRecording()  {}
func (recorder *Recorder) ResumeRecording() {}
