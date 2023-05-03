## Const
Go supports constants, which are variables that **cannot change their values.**
Constants in Go are defined with the help of the ```const``` keyword. Generally speaking,
constants can be either **global or local variables.** 

The **main benefit** you get from **using
constants** in your programs is the guarantee that their value **will not change during
program execution.**

The value of a constant variable is **defined at
compile time, not at runtime—this** means that it is included **in the binary executable.**

## Const declarations:
To declare a constant, we can use the following syntax:
```go
const shark = "Sammy"
```
Or multiple declare:

```go
const (
 idKey = "id"
 nameKey = "name"
)
```
**Can't reassign to the const variable:**
```go
const year = 365

year = year + 1
fmt.Println("year with new value: ", year)
// Result: cannot assign to year
```
## Typed and Untyped Constants:
Constants can be **typed or untyped.** An **untyped constant** works exactly **like a literal**;
it **has no type of its own**, but does **have a default type that is used when no other type
can be inferred.**

A **typed constant** can only be **directly assigned to a variable of that
type.**

Whether or not to make a constant typed **depends on why the constant was declared.**

If you are giving a name to a mathematical constant that could be used with multiple
numeric types, then keep the constant untyped. 

In general, leaving a **constant untyped gives you more flexibility.**

**Untyped const:**
```go
const x = 10
```

There are situations where you want a constant to
enforce a type. We’ll **use typed constants** when we look at **creating enumerations with
```iota```.**

**Typed const:**
```go
//All of the following assignments are legal:
var y int = x
var z float64 = x
var d byte = x
```

### For more information about typed and untyped const, check this sample:

```go
const (
	year     = 365
	leapYear = int32(366)
)

func main() {
	hours := 24
	minutes := int32(60)
	fmt.Println(hours * year)    
	fmt.Println(minutes * year)   
	fmt.Println(minutes * leapYear)
}
// Result
//8760
//21900
//21960
```
If you declare a constant with a **type**, it **will be that exact type**. Here when we declare the constant **leapYear**, we **define** it as data type **int32**.

Therefore it is a **typed constant**, which **means it can only operate with int32 data types**. 

The **year constant** we declare with **no type**, so it is considered **untyped**. Because of this, you **can use it with any integer data type.**

**If we try to multiply two types that are typed and not compatible, the program will not compile:**
```go
fmt.Println(hours * leapYear)
// Result 
// invalid operation: hours * leapYear (mismatched types int and int32)
```

In this case, **hours was inferred as an int**, and **leapYear was explicitly declared as an int32**. Because **Go is a typed language**, an **int** and an **int32** are **not compatible** for mathematical operations. 
**To multiply them**, you would **need to convert one to a int32 or an int.**

## Tip:
Constants in Go are a way to **give names** to literals. There is no way
in Go to declare that a variable is immutable.
