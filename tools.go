// +build tools

package main

import (
	_ "4d63.com/gochecknoglobals"
	_ "4d63.com/gochecknoinits"
	_ "github.com/alecthomas/gocyclo"
	_ "github.com/alexkohler/nakedret"
	_ "github.com/client9/misspell/cmd/misspell"
	_ "github.com/golang/lint/golint"
	_ "github.com/gordonklaus/ineffassign"
	_ "github.com/jgautheron/goconst/cmd/goconst"
	_ "github.com/kisielk/errcheck"
	_ "github.com/mdempsky/maligned"
	_ "github.com/mdempsky/unconvert"
	_ "github.com/mibk/dupl"
	_ "github.com/opennota/check/cmd/structcheck"
	_ "github.com/opennota/check/cmd/varcheck"
	_ "github.com/securego/gosec/cmd/gosec"
	_ "github.com/stripe/safesql"
	_ "github.com/tsenart/deadcode"
	_ "github.com/walle/lll/cmd/lll"
	_ "golang.org/x/tools/cmd/goimports"
	_ "golang.org/x/tools/cmd/gotype"
	_ "golang.org/x/tools/cmd/gotype"
	_ "honnef.co/go/tools/cmd/staticcheck"
	_ "mvdan.cc/interfacer"
	_ "mvdan.cc/unparam"
)
