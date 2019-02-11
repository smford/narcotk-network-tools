
# Simple reverse lookup tool

Note, the number of results will depend on the type of resolver you use.  Typically only 1 result will be returned by your resolver.

```
When using the host C library resolver, at most one result will be returned. To bypass the host resolver, use a custom Resolver.
```
(source)[https://golang.org/pkg/net/#LookupAddr]

## Compiling

Linux
```
GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o assets/reverselookup-linux-x86_64 reverselookup.go
```
