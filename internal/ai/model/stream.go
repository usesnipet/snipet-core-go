package model

type StreamChunk struct {
	Text string
}

type TextStream struct {
	Chunks <-chan StreamChunk
	Done   <-chan Result
	Error  <-chan error
}
