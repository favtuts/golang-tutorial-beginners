# golang-tutorial-beginners
The source code for blog post: https://www.favtuts.com/golang-tutorial-learn-go-programming-language-for-beginners/

# How to run Go program

```
go run <filename>
```

# Declare variable

```
# Variable can be declared using the syntax
var <variable_name> <type>

# give an initial value to a variable during the declaration itself 
var <variable_name> <type> = <value>

# omit the type during the declaration using the syntax
var <variable_name> = <value>

# declare multiple variables with the syntax
var <variable_name1>, <variable_name2>  = <value1>, <value2>

# declaring the variables with value by omitting the var keyword using
<variable_name> := <value>
```

# Declare constant

```
# eclared by using the keyword “const”
const <constant_name> =<value>
```

# For Loop syntax

```
for initialisation_expression; evaluation_expression; iteration_expression{
   // one or more statement
}
```

# If else syntax

```
if condition{
// statements_1
}else{
// statements_2
}
```

# Switch syntax

```
switch expression {
    case value_1:
        statements_1
    case value_2:
        statements_2
    case value_n:
        statements_n
    default:
        statements_default
    }
```

# Arrays

```
# The syntax for declaring an array is
var arrayname [size] type

# Each array element can be assigned value using the syntax
arrayname [index] = value

# Can assign values to array elements during declaration using the syntax
arrayname := [size] type {value_0,value_1,…,value_size-1} 

# Can also ignore the size parameter while declaring the array with values by replacing size with … 
# and the compiler will find the length from the number of values
arrayname :=  […] type {value_0,value_1,…,value_size-1}

# Can find the length of the array by using the syntax
len(arrayname)
```

# Slice and Append

```
# The syntax for creating a slice is
var slice_name [] type = array_name[start:end]

len(slice_name) – returns the length of the slice

append(slice_name, value_1, value_2) – Golang append is used to append value_1 and value_2 to an existing slice.

append(slice_nale1,slice_name2…) – appends slice_name2 to slice_name1
```


# Functions syntax

```
func function_name(parameter_1 type, parameter_n type) return_type {
//statements
}
```

# Packages syntax

```
# can import other packages in our program using the syntax
import package_name
```

# Pointers syntax

A pointer variable stores the memory address of another variable.
The asterisk(*) represents the variable is a pointer
```
var variable_name *type
```

# Structures syntax

A Structure is a user defined datatype which itself contains one more element of the same or different type.
The syntax for declaring a structure is
```
type struct_name struct {
   variable_1 variable_1_type
   variable_2 variable_2_type
   variable_n variable_n_type
}
```

create variables of the type
```
var variable_name struct_name
```

# Methods(not functions)

A method is a function with a receiver argument. Architecturally, it’s between the func keyword and method name. The syntax of a method is
```
func (variable variabletype) methodName(parameter1 paramether1type) {  
}
```

Go is not an object oriented language and it doesn’t have the concept of class. Methods give a feel of what you do in object oriented programs where the functions of a class are invoked using the syntax `objectname.functionname()`


# Goroutines

A goroutine is a function which can run concurrently with other functions, the calling function will not wait for the execution of the invoked function to complete. It will continue to execute with the next statements. You can have multiple goroutines in a program.

Goroutine is invoked using keyword go followed by a function call.

```
go add(x,y)
go function_name(parameter list)
```


# Channels

Channels are a way for functions to communicate with each other. It can be thought as a medium to where one routine places data and is accessed by another routine in Golang server.

A channel can be declared with the syntax
```
channel_variable := make(chan datatype)
```
You can send data to a channel using the syntax
```
channel_variable <- variable_name
```
You can receive data from a channel using the syntax
```
variable_name := <- channel_variable
```

The sender who pushes data to channel can inform the receivers that no more data will be added to the channel by closing the channel. This is mainly used when you use a loop to push data to a channel. A channel can be closed using
```
close(channel_name)
```
And at the receiver end, it is possible to check whether the channel is closed using an additional variable while fetching data from channel using
```
variable_name, status := <- channel_variable
```
If the status is True it means you received data from the channel. If false, it means you are trying to read from a closed channel


# Select

Select can be viewed as a switch statement which works on channels.

Add a default case to the select in the same program and see the output. Here, on reaching select block, if no case is having data ready on the channel, it will execute the default block without waiting for data to be available on any channel.

# Go for VS Code extension

```
Tools environment: GOPATH=C:\Users\tranvt\go
Installing 7 tools at C:\Users\tranvt\go\bin in module mode.
  gotests
  gomodifytags
  impl
  goplay
  dlv
  staticcheck
  gopls

Installing github.com/cweill/gotests/gotests@v1.6.0 (C:\Users\tranvt\go\bin\gotests.exe) SUCCEEDED
Installing github.com/fatih/gomodifytags@v1.16.0 (C:\Users\tranvt\go\bin\gomodifytags.exe) SUCCEEDED
Installing github.com/josharian/impl@v1.1.0 (C:\Users\tranvt\go\bin\impl.exe) SUCCEEDED
Installing github.com/haya14busa/goplay/cmd/goplay@v1.0.0 (C:\Users\tranvt\go\bin\goplay.exe) SUCCEEDED
Installing github.com/go-delve/delve/cmd/dlv@latest (C:\Users\tranvt\go\bin\dlv.exe) SUCCEEDED
Installing honnef.co/go/tools/cmd/staticcheck@latest (C:\Users\tranvt\go\bin\staticcheck.exe) SUCCEEDED
Installing golang.org/x/tools/gopls@latest (C:\Users\tranvt\go\bin\gopls.exe) SUCCEEDED

All tools successfully installed. You are ready to Go. :)
```