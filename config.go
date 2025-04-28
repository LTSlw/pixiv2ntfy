package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
)

func bindStr(p *string, nameEnv, nameFlag string, value string, usage string) {
	if env, ok := os.LookupEnv(nameEnv); ok {
		flag.StringVar(p, nameFlag, env, usage)
	} else {
		flag.StringVar(p, nameFlag, value, usage)
	}
}

func bindUint64(p *uint64, nameEnv, nameFlag string, value uint64, usage string) {
	if env, ok := os.LookupEnv(nameEnv); ok {
		v, err := strconv.ParseUint(env, 10, 64)
		if err != nil {
			failf(os.Stderr, "invalid value %q for env %s: parse error\n", env, nameEnv)
		}
		flag.Uint64Var(p, nameFlag, v, usage)
	} else {
		flag.Uint64Var(p, nameFlag, value, usage)
	}
}

func failf(w io.Writer, format string, a ...any) {
	fmt.Fprintf(w, format, a...)
	os.Exit(2)
}
