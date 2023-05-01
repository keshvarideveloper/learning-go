# Declare variables with ```var```
Go has a lot of ways to declare variables. There’s a reason for
this: each declaration style communicates something about how the variable is used.

* The most verbose way to declare a variable in Go uses the ```var``` keyword, an **explicit
type, and an assignment.** It looks like this:
  * ```go
     var x int = 10
* If the type on the righthand side of the = is the **expected type of your variable**, you can
  **leave off the type from the left side of the =**. Since the **default type of an integer literal
  is int**, the following declares x to be a variable of type int:
  * ```go 
    var x = 10
*  If you want to **declare a variable and assign it the zero value**, you can keep
   the type and **drop the =** on the righthand side:
  * ```go
    var x int
* You can declare multiple variables at once with var, and they can be of the same type:
  * ```go
    var x, y int = 10, 20
* All zero values of the same type:
  * ```go
    var x, y int
* Or of different types:
  * ```go
    var x, y = 10, "hello"

## Declaring multiple variables at once:
If you are declaring multiple variables at once, you
can wrap them in a declaration list:
```go
can wrap them in a declaration list:
var (
    x int
    y = 20
    z int = 30
    d, e = 40, "hello"
    f, g string
)
```
## Short declaration:
When you are **within a function**, you can
use the ```:=``` operator to replace a var declaration that uses type inference. The following two statements do exactly the same thing: 

they declare x to be an int with the
value of 10:
```go
var x = 10
// OR
x := 10
```

Like var, you can declare **multiple variables at once** using ```:=``` :
```go
var x, y = 10, "hello"
// OR
x, y := 10, "hello"

```
## How do you know which style to use:
As always, choose what makes your intent
clearest. The most common declaration style **within functions is ```:=```**

There are some situations **within functions where you should avoid ```:=```***:
* When initializing a variable to its **zero value**, use ```var x int```. This makes it clear
that the zero value is intended.
* When assigning an untyped constant or a literal to a variable and the default type
for the constant or literal isn’t the type you want for the variable, use the long var
form with the type specified. While it is legal to use a type conversion to specify
the type of the value and use := to write ```x := byte(20)```, it is idiomatic to write
```var x byte = 20```.
* Because ```:=``` allows you to assign to both new and existing variables, it sometimes
creates new variables when you think you are reusing existing ones. In those situations, explicitly declare all
of your new variables with var to make it clear which variables are new, and then
use the assignment operator (=) to assign values to both new and old variables.
* While var and ```:=``` allow you to **declare multiple variables** on the same line, only use
  this style **when assigning multiple values returned from a function**.
* Avoid declaring variables outside of functions because they **complicate data flow analysis**:
  * **You should ***rarely*** declare variables outside of functions**, in what’s called the package
    block. **Package-level variables** whose values change are a
    **bad idea**. When you have a **variable outside of a function**, it can **be difficult to track
    the changes made to it**, which makes it hard to understand how data is flowing
    through your program. This can lead to subtle bugs. **As a general rule, you should
    only declare variables in the package block that are effectively immutable.**
## Tip
* **The ```:=``` operator can do one trick that you cannot do with var:** it **allows you to assign
values to existing variables**, too. As long as there is **one new variable on the lefthand
side of the ```:=```**, then any of the **other variables can already exist**:

  *   ```go 
      x := 10
      x, y := 30, "hello" // You can assign value to x because y is a new variable


* There is one **limitation on ```:=```**. If you are declaring a variable at package level, you
 must use var because **```:=``` is not legal outside of functions.**

* Avoid declaring variables outside of functions because they complicate data flow analysis.