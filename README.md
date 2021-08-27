# Go test label
Minimal and simplest way to add colored label to golang console text for test cases.

![go-test-label](https://user-images.githubusercontent.com/2642811/131183449-d3f274c3-9cb8-4d82-9fd9-da2b64fa62ec.png)

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

### Test

Run test with in terminal with verbose mode

```bash
go test -v
```
