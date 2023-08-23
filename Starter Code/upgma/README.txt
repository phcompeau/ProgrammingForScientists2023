Copy the upgma folder into your "go/src" directory.

Then, write each of the missing functions in functions.go. folder.

To test a function, you should run

go test -run ^TestFunctionName$

where TestFunctionName is based on the specific function you are testing.

For example, after writing each appropriate function, you should call each of the following tests.

go test -run ^TestCountLeaves$
go test -run ^TestFindMinElement$
go test -run ^TestAddRowCol$
go test -run ^TestUPGMA$
go test -run ^TestBuildClustalTree$
go test -run ^TestGetMultipleAlignment$

To check that you have passed *all* tests, call "go test".

We will periodically navigate up ("cd ..") into the parent directory and run code calling our functions using main.go.