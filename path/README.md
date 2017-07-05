# path

## file.go

#### CreateFile(path string) error

- `path` the path of file

#### AppendToFile(path string, s ...string) error

- `path` the path of file
- `s` the content strings

#### OverwriteFile(path string, s ...string) error

- `path` the path of file
- `s` the content strings

#### ReadFile(path string) (s string,err error)

- `path` the path of file
- `s` the content of the file

#### RemoveFile(path string, dir bool) error

- `path` the path of file or directory
- `dir` if the path directory

## home.go

#### HomeDir() (home string,err error)

- `home` the path to the user's home directory

## path.go

#### IsExist(path string) (ok bool, err error)

- `path` the path of file or directory
- `ok` exists the path of file or directory
- `err` if not match not exist errors

#### MustExist(_path string) (ok bool)

- `ok` exists the path of file or directory

#### Ensure(path string, dir bool) error

- `path` the path of file or directory
- `dir` the path is directory

#### ExecPath() (path string)

- `path` the file path of program

#### ExecDir() (path string)

- `path` the directory path of program
