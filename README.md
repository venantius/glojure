# Clojure-Go

Clojure on the Go runtime.

## Environment Configuration

I find Go's model of development environment configuration to be confusing. Since I find I have to re-learn it each time I come back to it, I might as well document the fact that the expectation is that this package sits in $GOPATH/src.

Hence, go/main.go would sit at $GOPATH/src/clojure-go/go/main.go.

### GO15VENDOREXPERIMENT

We're using the `vendor` directory with the experimental vendoring flag. More details can be found [here](https://github.com/golang/go/wiki/PackageManagementTools#go15vendorexperiment). The TL;DR is: 
 * set `GO15VENDOREXPERIMENT`=1
 * make sure you're using at least Go version >=1.5.0
