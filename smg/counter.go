package smg

import "io"

// JustCounterWriter todo
type JustCounterWriter struct {
	count int64
}

// Write writes to io.Discard to get the bytes number
func (cw *JustCounterWriter) Write(p []byte) (n int, err error) {
	n, err = io.Discard.Write(p)
	cw.count += int64(n)
	return n, err
}

// Count Counts the bytes number
func (cw *JustCounterWriter) Count() (n int64) {
	return cw.count
}
