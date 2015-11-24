```
                                                 ,-.
                                                  ) \
                                              .--'   |
                                             /       /
                                             |_______|
                                            (  O   O  )
                                             {'-(_)-'}
                                           .-{   ^   }-.
                                          /   '.___.'   \
                                         /  |    o    |  \
                                         |__|    o    |__|
                                         (((\_________/)))
                                             \___|___/
                                        jgs.--' | | '--.
                                           \__._| |_.__/
```

[Warden](https://github.com/cloudfoundry/warden) in Go, because why not?

[![Build Status](https://travis-ci.org/cloudfoundry-incubator/garden.png?branch=master)](https://travis-ci.org/cloudfoundry-incubator/garden)
[![Coverage Status](https://coveralls.io/repos/cloudfoundry-incubator/garden/badge.png?branch=HEAD)](https://coveralls.io/r/cloudfoundry-incubator/garden?branch=HEAD)

# Backends

Garden provides a platform-neutral API for containerization. Backends implement support for various specific platforms. So far, the list of backends is as follows:

 - [Garden Linux](https://github.com/cloudfoundry-incubator/garden-linux/) - Linux backend
 - [Guardian](https://github.com/cloudfoundry-incubator/guardian/) - Linux backend using [runc](https://github.com/opencontainers/runc)
 - [Greenhouse](https://github.com/cloudfoundry-incubator/garden-windows) - Windows backend

# Client API

The canonical API for Garden is defined as a collection of Go interfaces. See the [godoc documentation](http://godoc.org/github.com/cloudfoundry-incubator/garden) for details.

# Development

## Prerequisites

* [go](https://golang.org)
* [git](http://git-scm.com/) (for garden and its dependencies)
* [mercurial](http://mercurial.selenic.com/) (for some other dependencies not using git)

## Running the tests

Assuming go is installed and `$GOPATH` is set:
```
mkdir -p $GOPATH/src/github.com/cloudfoundry-incubator
cd $GOPATH/src/github.com/cloudfoundry-incubator
git clone git@github.com:cloudfoundry-incubator/garden
cd garden
go get -t -u ./...
go install github.com/onsi/ginkgo/ginkgo
ginkgo -r
```
