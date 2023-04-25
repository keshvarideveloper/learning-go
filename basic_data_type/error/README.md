# Error Data Type
Go provides a special data type for representing error conditions and error messages
named error, in practice, this means that **Go treats errors as value.**

## nil value
if the value of an error variable is nil, then there was no error.

## Error message
When error occurred, you can get the message from error:
```go
_, err := strconv.Atoi("ABC123")
if err != nil {
// Print err message
fmt.Println(err)
}
```
## Custom error message

This usually happens inside a function **other than main()**,
because **main() does not return anything to any other function.** Additionally, a good
place to define your custom errors is inside the Go packages you create.

### You can create your own error message with two ways:
#### 1- errors.New()
#### 2- fmt.Errorf()

### errors.New()

You can use **errors.New()** from the errors package.
```go 
errors.New("custom error ...")
```

###  fmt.Errorf()
You can use the **fmt.Errorf()** function, which simplifies the creation of custom error
messages, the fmt.Errorf() function returns an error value just like errors.New().

```go
// Suppose a=1 b=2 c=3
 fmt.Errorf("a: %d and b: %d, c: %d", a, b, c) 
```

## Global error handling tactic:
You should have a global error handling tactic in each application that should not change. In practice, this
means the following:

* All error messages should be handled at the same level, which means that
**all errors should** either **be returned to the calling function** or **be handled at
the place they occurred.**
* It should be clearly documented how to handle critical errors. This means
that there will be situations **where a critical error should terminate the
program** and other times **where a critical error might just create a warning
message onscreen.**
* It is considered a good practice to **send all error messages to the ***log service***** of
your machine because this way the error messages can be examined at a later
time. However, **this is not always true**, so exercise caution when setting this
up, for example, **cloud native apps do not work that way.**

## Tip
Usually you do not need to define custom error messages unless you are
creating big applications or packages.