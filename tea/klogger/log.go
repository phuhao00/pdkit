package klogger

import (
	"go.uber.org/zap"
	"log"
)

var (
	DefaultLogger =NewDefaultLogger()
)

type KLogger struct {
	zap *zap.Logger
	sys  *log.Logger

}

//NewDefaultLogger get default logger.
func NewDefaultLogger() *KLogger {
	zL,err:=zap.NewProduction(zap.AddCallerSkip(1))
	if err != nil {
		log.Fatalf("NewDefaultLogger err:%s",err.Error())
	}
	return &KLogger{
		zap:    zL,
		sys: nil,
	}
}

//ErrorF printf err
func (k *KLogger)ErrorF(template string, args ...interface{})  {
	sug:=k.zap.Sugar()
	sug.Errorf(template,args)
}

//InfoF printf info
func (k *KLogger)InfoF(template string, args ...interface{})  {
	sug:=k.zap.Sugar()
	sug.Infof(template,args)
}

