# go_errcheck_sample
errcheckというcliアプリの展覧

# Installation

1. clone the source
1. run `setup.sh`
1. run `errcheck.sh`


# Details

1. Install `errcheck`:
    > `go install github.com/kisielk/errcheck@latest`
1. Make sure gopath is in your $PATH
    > ``export PATH=${PATH}:`go env GOPATH\`/bin``
1. Run the error check:
    > `errcheck -blank -asserts -ignore 'Write' cmd/*`

# Results:
>cmd/main.go:19:15:      defer f.Close() // <-- error! We didn't use the err value  
cmd/main.go:26:5:       f, _ := os.Stat("non_existant_file.txt") // <-- error! Discarded  
cmd/main.go:50:9:       f.Chdir() // <-- error! Unchecked  
cmd/main.go:56:7:       i := obj.(int) // <-- error! Unchecked  

# Limitations

Cannot uncover non-use-after-reassign, aka, setting err twice but forgetting to check the second err value. That is covered by https://staticcheck.io/