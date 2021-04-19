# network-scanner
A simple network scanner in go

[![forthebadge](https://forthebadge.com/images/badges/made-with-go.svg)](https://forthebadge.com)

# Installation:
```
go get -u https://github.com/krishpranav/network-scanner
git clone https://github.com/krishpranav/network-scanner
cd network-scanner
go build cmd/portscanner/portscanner.go
```

# Help
```
 portscanner -h
 Usage of portscanner:
   -A   Scans all ports from port 1 to 1024
   -p   Port or ports to scan (default "80")
```

# Local Host Portscanner
```
$ ./portscanner
```

# Network Machine Full Port Scan
```
$ ./portscanner -p 1-65535 192.168.0.8
```

# Network Discovery
```
$ ./portscanner -p 21,80 192.168.0.1-192.168.0.254
```
