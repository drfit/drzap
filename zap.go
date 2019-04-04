package zap

import (
	"github.com/go-dr/dr"

	log "github.com/unisx/logus"
)

func init() {
	dr.SetLogger(drZap{})
}

type drZap struct{}

const (
	logDr = "Dr.Log"
	msgDr = "Dr.Msg"
)

func (drZap) Info(msg string) {
	log.Info(logDr, log.String(msgDr, msg))
}

func (drZap) Debug(msg string) {
	log.Debug(logDr, log.String(msgDr, msg))
}

func (drZap) Warn(msg string) {
	log.Warn(logDr, log.String(msgDr, msg))
}

func (drZap) Error(msg string) {
	log.Error(logDr, log.String(msgDr, msg))
}

func (drZap) Panic(msg string) {
	log.Panic(logDr, log.String(msgDr, msg))
}
