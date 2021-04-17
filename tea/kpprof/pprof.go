package kpprof

import (
	"net/http"
	"net/http/pprof"
	"pdkit/tea/klogger"
)

// InitPprof start http pprof.
func InitPprof(logger klogger.KLogger) {
	pprofServeMux := http.NewServeMux()
	pprofServeMux.HandleFunc("/debug/pprof/", pprof.Index)
	pprofServeMux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	pprofServeMux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	pprofServeMux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	for _, addr := range []string{":8088"} {
		go func() {
			if err := http.ListenAndServe(addr, pprofServeMux); err != nil {
				logger.ErrorF("http.ListenAndServe(\"%s\", pproServeMux) error(%v)", addr)
				panic(err)
			}
		}()
	}
}

