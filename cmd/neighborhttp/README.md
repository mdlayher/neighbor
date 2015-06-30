neighborhttp
============

Command `neighborhttp` provides a simple HTTP server which retrieves neighbor
information for any device which makes a request to it.

Usage
-----

On host machine:

```
$ ./neighborhttp -host :8080
```

On neighbor machine:

```
$ curl -6 -g "[2002:db8::212:7fff:feeb:6b40]:8080"
2002:db8::212:7fff:feeb:6b40 -> 00:12:7f:eb:6b:40 [StateReachable] Flags(0)
$ curl 192.168.1.1:8080
192.168.1.1 -> 00:12:7f:eb:6b:40 [StateReachable] Flags(0)
```
