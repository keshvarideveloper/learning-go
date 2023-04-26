# Numeric Data Type

Go has a large number of numeric types: **12** different types (and a few special names)
that are grouped into three categories:
* Integer
* Floating point
* Complrx

If you are coming from a language like Java-Script that gets along with only a single numeric type, this might seem like a lot. And
in fact, some types are used frequently while others are more esoteric.

## Integer 
Go provides both **signed** and **unsigned** integers in a **variety of sizes**, from one to eight
bytes.

| Type name |                 Value range                 |           Description           |
|:---------:|:-------------------------------------------:|:-------------------------------:|
|   int8    |                 -128 to 127                 |      8-bit signed integer       |
|   int16   |               -32768 to 32767               |      16-bit signed integer      |
|   int32   |          -2147483648 to 2147483647          |      32-bit signed integer      |
|   int64   | -9223372036854775808 to 9223372036854775807 |      64-bit signed integer      |
|   uint8   |                  0 to 255                   |     8-bit unsigned integer      |
|  uint16   |                 0 to 65535                  |     16-bit unsigned integer     | 
|  uint32   |               0 to 4294967295               |     32-bit unsigned integer     |
|  uint64   |        0 to  to 18446744073709551615        |     64-bit unsigned integer     |
|    int    |       32- or 64-bit signed int range        |  32- or 64-bit signed integer   |
|   uint    |      32- or 64-bit unsigned int range       | 32- or 64-bit unsigned integer  |

Please read the tip about int and uint type for more details in below.

### how to get those numbers from type
Any of those types tell us what size can be stored in an integer variable. To remember the size you can think in terms of: 1 beer (1 byte), 2 beers (2 bytes), 4 beers (4 bytes) or 8 beers (8 bytes).

#### For example:
* Int8 : 1 beer/byte = 8 bit = 2^8 = 256 = 256/2 = -128 to 127

* Int16 : 2 beers/bytes = 16 bit = 2^16 = 65536 = 65536/2 = -32768 to 32767

* Int32 : 4 beers/bytes = 32 bit = 2^32 = 4294967296 = 4294967296/2 = -2147483648 to 2147483647

* Int64 : 8 beers/bytes = 64 bit = 2^64 = 18446744073709551616 = 18446744073709551616/2 = -9223372036854775808 to 9223372036854775807

### zero value:
zero value for all the integer types is **0**.

### Tip for *int* type:
The **int** and **uint** data types are special as they are the most efficient sizes for **signed
and unsigned integers** on a given platform and **can be either 32 or 64 bits each—their
size is defined by Go itself**. The int data type is the most widely used data type in
Go due to its versatility

On a **32-bit CPU, int is a 32-bit signed integer** like
an int32. On most **64-bit CPUs, int is a 64-bit signed integer**, just like an int64.
Because int isn’t consistent from platform to platform, **it is a compile-time error to
assign, compare, or perform mathematical operations between an int and an int32
or int64 without a type conversion**

Go does have some special names for integer types. **A byte is an alias for uint8**; it is
legal to assign, compare, or perform mathematical operations between a byte and a
uint8. However, you rarely see uint8 used in Go code; just call it a byte.




## Floating point 
There are two floating point types in Go:


| Type name |              Largest absolute value               |        Smallest (nonzero) absolute value        |           Description           |
|:---------:|:-------------------------------------------------:|:-----------------------------------------------:|:-------------------------------:|
|  float32  |   3.40282346638528859811704183484516925440e+38    |  1.401298464324817070923729583289916131280e-45  |  32-bit floating-point number   |
|  float64  | 1.797693134862315708145274237317043567981e+308    | 4.940656458412465441765687928682213723651e-324  |  64-bit floating-point number   |


### Zero value:
Zero value for all the floating point types is **0**.

### Tip:
If you don't give type (float32 OR float64) to variable, its type will be float64