package main

import (
	"glojure/lang"
	"os"
)

var CLOJURE_MAIN *lang.Symbol = lang.InternSymbolByNsname("clojure.main")
var REQUIRE *lang.Var = lang.RT.Var("clojure.core", "require")
var MAIN *lang.Var = lang.RT.Var("clojure.main", "main")

func main() {
	REQUIRE.Invoke(CLOJURE_MAIN)
	MAIN.ApplyTo(lang.RT.Seq(os.Args...))
}
