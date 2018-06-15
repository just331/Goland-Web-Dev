# Techniques we will learn 

- concatentae 
- CLI pipleline - output to a file with >

# Code we will use from the standard library 

## [os.Create](https://godoc,org/os#Create) 
This allows us to create a file.

``` Go 
func Create(name String) (*File, error)
```

***

## [defer](https://golang.org/ref/spec#Defer_statements)
The defer keyword allows us to defer the execution of a statement until the function in which we have placed the defer statement returns.
 
 ***
 
 ## [io.copy](https://godoc.org/io#copy)
 This allows us to copy from a source to a destination. 
 ``` Go
 func Copy(dst Writer, src Reader) (written int64, err error)
 ```
 
 ## [strings.NewReader](https://godoc.org/strings#NewReader)
 NewReader returns a new Reader reading from s.
 ``` Go
 func NewReader(s String) *Reader
 ```
 
 ##[os.Args](https://godco.org/os#pkg-variables)
 Args is a variable from package os. Args hold the command-line arguments, start with the program name.
 ``` Go
 var Args []string
 ```