neighbor [![Build Status](https://travis-ci.org/mdlayher/neighbor.svg?branch=master)](https://travis-ci.org/mdlayher/neighbor) [![GoDoc](http://godoc.org/github.com/mdlayher/neighbor?status.svg)](http://godoc.org/github.com/mdlayher/neighbor)
========

Package `neighbor` enables network neighbor detection using operating system
specific facilities.  MIT Licensed.

At this time, `neighbor` only works on Linux.  Ports to other platforms should
be possible through use of other operating system specific APIs.

Two example programs which demonstrate the library can be found in
[cmd/](https://github.com/mdlayher/neighbor/tree/master/cmd/).

Some code in this package was taken directly from the Go standard library,
with slight modifications.  The Go standard library is Copyright (c) 2012
The Go Authors. All rights reserved.  The Go license can be found at
https://golang.org/LICENSE.
