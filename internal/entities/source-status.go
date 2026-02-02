package entities

type SourceStatus string

const (
	SourceActive SourceStatus = "active"
	SourcePaused SourceStatus = "paused"
	SourceError  SourceStatus = "error"
)
