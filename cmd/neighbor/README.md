neighbor
========

Command `neighbor` provides neighbor detection akin to `ip neighbor show`.

Usage
-----

```
$ ./neighbor -h
Usage of neighbor:
  -4=false: IPv4 only?
  -6=false: IPv6 only?
  -ip="": IP address to lookup specific neighbor
```

Show all IPv4 and IPv6 neighbors:

```
$ ./neighbor
fe80::212:7fff:feeb:6b40 dev eth0 lladdr 00:12:7f:eb:6b:40 StateReachable FlagsRouter
2002:db8::212:7fff:feeb:6b40 dev eth0 lladdr 00:11:22:33:44:55 StateDelay
192.168.1.1 dev eth0 lladdr 00:12:7f:eb:6b:40 StateReachable
192.168.1.2 dev eth0 lladdr 00:11:22:33:44:55 StateDelay
```

Show all IPv4 neighbors only:

```
$ ./neighbor -4
192.168.1.1 dev eth0 lladdr 00:12:7f:eb:6b:40 StateReachable
192.168.1.2 dev eth0 lladdr 00:11:22:33:44:55 StateDelay
```

Show all IPv6 neighbors only:

```
$ ./neighbor -6
fe80::212:7fff:feeb:6b40 dev eth0 lladdr 00:12:7f:eb:6b:40 StateReachable FlagsRouter
2002:db8::212:7fff:feeb:6b40 dev eth0 lladdr 00:11:22:33:44:55 StateDelay
```

Find a specific neighbor by IP address:

```
$ ./neighbor -ip fe80::212:7fff:feeb:6b40
fe80::212:7fff:feeb:6b40 dev eth0 lladdr 00:12:7f:eb:6b:40 StateReachable FlagsRouter
$ ./neighbor -ip 192.168.1.1
192.168.1.1 dev eth0 lladdr 00:12:7f:eb:6b:40 StateReachable
```
