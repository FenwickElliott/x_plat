# xplat
Utility package contaning enviroment specific methods and variables

### Instalation
`go get github.com/fenwickelliott/xplat`

### Usage
`import "github.com/fenwickelliott/xplat`

### Methods
* `xplat.Appdir(subdirs ...string) (string, error)`
Returns the appropriate directory for storing user specific application data, with optional substrings.

* `xplat.Openbrowser(url string) error`
Opens the local enviroment's default browser to the given url