# Go test label
Minimal and simplest way to add colored label to golang console text for test cases.
## Instructions

### Install
Add module to the project by running

```bash
go get github.com/sajibsrs/go-test-label
```

### Import
```go
import label "github.com/sajibsrs/go-test-label"
```

### Usage
Boolean param defines whether label will have background.

```go
label := label.NewLabel(false) // param boolean
t.Logf("%sTV turned off", label.Pass)
```
