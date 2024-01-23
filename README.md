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

- Convert the Application to the module. Modules consist of many packages. It creates the go.mod file that we can manage dependencies. It also creates the go.sum file so that we can see the history of the modules.
```
go mod init go-rest-service-example
```

- To get and add a module into your project.
```
go get github.com/twpayne/go-geom
```

- Removes unusued modules and arranges them.
```
go mod tidy
```