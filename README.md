# xplat
Utility package contaning enviroment specific methods and variables

### Instalation
`go get github.com/fenwickelliott/xplat`

### Usage
`import "github.com/fenwickelliott/xplat`

### Methods
* `xplat.Appdir()`
Returns the appropriate directory for storing user specific application data, optionally joined to variadic subdirectories.

* `xplat.Openbrowser(url string)`
Opens the local enviroment's default browser to the given url