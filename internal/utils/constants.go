package utils

const (
	ContextKeyUser contextKey = iota
	ContextKeyLogger
)

type contextKey int8
