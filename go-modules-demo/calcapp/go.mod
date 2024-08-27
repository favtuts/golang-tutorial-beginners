module calcapp

go 1.18

// Can be github.com/yourname/operations
replace favtuts.com/operations => ../operations

require github.com/ttacon/chalk v0.0.0-20160626202418-22c06c80ed31

// This is a random version number I added. You can actually put any semantic version here
require favtuts.com/operations v0.0.0
