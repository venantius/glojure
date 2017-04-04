# Glojure

Clojure on the Go runtime.

Status: **INCOMPLETE**

This is a compiler I spent a few months working on in early 2016. It is heavily based on the JVM Clojure compiler (version 1.8), although it deviates from it in notable ways due to the differences in threading and inheritance models between Go and Java.

With regard to the current state of the compiler: most of the core data structures (keywords, maps, vectors, etc.) have implementations that are nearly complete, and at the very least functional. The lexer and parser, however, still have a lot of work remaining, and due to the significant differences in threading models are likely to require a lot of consideration and thought.

I'd love to continue to work on it but at the moment it's not a focus. If you're interested in working on it, shoot me an email.

## Installation

Make sure you have Go installed (v. >=`1.5.0`)

1. Clone this repository inside `$GOPATH/src`
2. `cd` to inside the repository and run `git submodule update`
3. Add `export GO15VENDOREXPERIMENT=1` to your `~/.bash_profile`

# License

Copyright © 2017 David Jarvis

Distributed under the Eclipse Public License 1.0, the same as Clojure.
