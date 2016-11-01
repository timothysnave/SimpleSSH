# SimpleSSH

SimpleSSH is a convenience wrapper around the golang.org/x/crypto/ssh library. It allows for easily connecting to servers through SSH, executing commands, and returning the results.

Example:
```go
package main

import "SimpleSSH"

func main() {
	user := "myuser"
	pass := "mypassword"
	serv := "myserver.mydomain.com"

	myssh := SimpleSSH.New(serv, user, pass)
	defer myssh.Cleanup()

	out := myssh.Run("date")
	println(out)
}
```
