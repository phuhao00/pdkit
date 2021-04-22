package xgo

import (
	"fmt"
	//"github.com/pkg/errors"
)

func try2(fn func(), cleaner func()) (ret error) {
	if cleaner != nil {
		defer cleaner()
	}
	defer func() {
		//_, file, line, _ := runtime.Caller(5)
		if err := recover(); err != nil {
			//_logger.Error("recover", zap.Any("err", err), zap.String("line", fmt.Sprintf("%s:%d", file, line)))
			if _, ok := err.(error); ok {
				ret = err.(error)
			} else {
				ret = fmt.Errorf("%+v", err)
			}
		}
	}()
	fn()
	return nil
}

func try(fn func() error, cleaner func()) (ret error) {
	if cleaner != nil {
		defer cleaner()
	}
	defer func() {
		if err := recover(); err != nil {
			//_, file, line, _ := runtime.Caller(2)
			//_logger.Error("recover", zap.Any("err", err), zap.String("line", fmt.Sprintf("%s:%d", file, line)))
			if _, ok := err.(error); ok {
				ret = err.(error)
			} else {
				ret = fmt.Errorf("%+v", err)
			}
			//ret = errors.Wrap(ret, fmt.Sprintf("%s:%d", xstring.FunctionName(fn), line))
		}
	}()
	return fn()
}