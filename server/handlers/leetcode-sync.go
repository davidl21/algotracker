package handlers

import (
	"log"
)

type LeetcodeSync struct {
	l *log.Logger
}

func NewLeetcodeSync(l *log.Logger) *LeetcodeSync {
	return &LeetcodeSync{l}
}