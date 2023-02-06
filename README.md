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