- Check the Go version
```
go version
```

- Run the Go Application
```
go run main.go
```

- Import the fmt package to print something on the screen.
```
package main
import "fmt"
func main() {
    fmt.Println("Hello, world!")
}
```

- Convert the Application to the module. Modules consist of many packages.
```
go mod init go-rest-service-example
```

- Removes unusued modules and arranges them.
```
go mod tidy
```