# Calpoly-classes

Get the classes, name, and number of units in your terminal. This code is written in go language.

```go
go get .
go build -o class -ldflags="-s -w" main.go
```

## Usage:
Please choose acronyms from the school catalog to get the printed table of your classes on the terminal.

```bash
usage: ./class [acronyms]
```

![Sample Outp](./1.png)

## Compiled Version:

I added the compiled version of my code for windows x64 and mac so that if you don't have golang on your computer, you can use the compiled one.

**Windows:**

```bash
classw.exe csc
```
**Mac:**

```bash
./class csc
```
