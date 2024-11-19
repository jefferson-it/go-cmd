# Commands System - GoLang

| Idioma - Language |
| ----------------  |
| <a href="https://github.com/jefferson-developer-it/go-cmd/blob/main/README-PT.md">Português - Portuguese</a>
| <a href="#start">Inglês - English</a>

<div id="start">

# About

</div>

The idea of ​​creating commands came from the objective of readapting me to windows, using the terminal, after all, I used Linux for a long time, and with that I adapted to some commands that we don't have in windows, such as touch, ls (it has powershell), etc. Some were created by <strong>fun</strong>.

I created <a href="https://github.com/jefferson-developer-it/sh-commands">SH Commands</a> while using Linux, and <a href="https://github. com/jefferson-developer-it/commands_system">Commands System</a> when I was on windows.

## Problems

- Sh Commands was made to abbreviate Linux commands, suitable for Linux.
- Commands System was made for Windows, using C++ and NodeJS, so <em>some</em> commands would not work on Linux machines, nor machines without NodeJS.

This version made in Go is made to support on any machine, without any other intermediate language.
Just need GO to compile(download of compiled version will be available.)

## Why not C++?
Yes, I thought of doing it entirely in C++, on the other hand... I'm studying Go, and what better to practice if not a Go project?
GO was chosen for this reason.
</div>

# Commandos

## filet
Command inspired by "touch" and "cat", it has the function of writing and files.

```shell
$ file <flag?> <filename> <content>
```
### flag value

| Value | Function |
| ------|-------- |
| -r | Read files |
| -w | Write file|
| -ow | Overwrite file|
| -h | Help|
| * | Value goes to file name|

Any value other than these from the table is classified as *, and it goes to filename, while the flag is set to "auto", so it will read the file, if it exists, if not, create it.

```shell
$ file text.txt
```
```txt
If there is - read e.g: Hello world!
If not - write(as there is no content value):
// This file is writing for file(Go Commands), github.com/jefferson-developer-it/go-cmd
```

```shell
$ file text.txt "Hello, World!"
```

If there is - read eg Hello, World!
If not - write eg "Hello, World!"

```shell
$ file -w text.txt "Hello, Go!"
```

```txt
If it exists:
"This file exists!! Use 'file -ow <filename> <content>' for overwriting!"
```

```shell
$ file -ow text.txt "Hello, Go!"
```

Overwrite the existing file.

```shell
$ file -r text.txt "Hello, Go!"
```
```
Value of "<content>" is ignored and read the contents of the file.
```

## ç

Clears the terminal screen, equivalent to cmd(windows) cls and Linux/Mac clear.

## byte and bytestr

Convert string text to byte array:

```shell
$ byte byte "Hello, Go"
output: [72 101 108 108 111 44 ​​32 71 111]
```

We have bytestr which converts byte to string:
```shell
$bytestr [ 79, 108, 196, 161, 32, 74, 101, 102, 102, 33]
outpu: Hello Jeff!
```
Or
```shell
$bytestr 79, 108, 196, 161, 32, 74, 101, 102, 102, 33
outpu: Hello Jeff!
```

## lj

Equivalent to ls (just not as complete), read the directory:

```shell
$ lj ./commands
Name Type Size
----- ----- -----
byte.go file 466kb
bytestr.go file 899kb
c.go file 478kb
file.go file 1758kb
lj.go file 888kb
percent.go file 570kb
```

## percent

Calculates percentages:
```shell
$ percent <base?> <percent?>
```

Sorry, but here it is in Portuguese, Examples.

```shell
$ percent

Output:
Qual o valor total?
> 100

Quantos porcento a ser descontado?
> 50

Você também pode usar 'percent <base> <percent>'.

50% de 100 é 50.
E 100 - 50% = 50.
```

```shell
$ percent 75 25

Output:
32% de 75 é 24.
E 75 - 32% = 51.
```