# vernam
This little package can encrypt and decrypt a message through vernam's cipher. The message must be lowercase and without spaces
___
To install, in a terminal type:

```
go get github.com/Leox-06/vernam
```

Function you can call
```
# RandomKey('length of massage')
# Encrypt('message', 'key')
# Decrypt('encrypted message', 'key')
```

___
Example usage:

```go
package main

import (
	"fmt"

	"https://github.com/Leox-06/vernam"
)

func main(){
	msg := "message"
	key := vernam.RandomKey(len(msg))
	fmt.Println(vernam.Encrypt(msg, key))
}
```
