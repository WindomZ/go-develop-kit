# path

## home.go

#### HomeDir() (home string,err error)

- `home` the path to the user's home directory

## path.go

#### IsExist(path string) (ok bool, err error)

- `path` the path of file or directory
- `ok` exists the path of file or directory
- `err` if not match not exist errors

#### Ensure(path string, dir bool) error

- `path` the path of file or directory
- `dir` the path is directory
