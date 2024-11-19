# GoLang - Comandos Sistema

| Idioma - Language |
| ----------------  |
| <a href="#start">Português - Portuguese</a>
| <a href="https://github.com/jefferson-developer-it/go-cmd/blob/main/README.md">Inglês - English</a>

<div id="start">

# Sobre

A ideia de criar comandos veio do objetivo de me readaptar ao windows, usando o terminal, afinal, usei Linux um bom tempo, e com isso me adaptei a alguns comandos que não temos no windows, como touch, ls(tem pelo powershell), etc. Alguns foram criados por <strong>diversão</strong>.

Criei o <a href="https://github.com/jefferson-developer-it/sh-commands">SH Commands</a> enquanto usava o Linux, e o  <a href="https://github.com/jefferson-developer-it/commands_system">Commands System</a> quando tava no windows.

## Problemas

- O Sh Commands foi feito para abreviar comandos do Linux, próprio para Linux.
- O Commands System foi feito para Windows, usando C++ e NodeJS, então <em>alguns</em> commandos não funcionaria em máquinas Linux, tão pouco máquinas sem o NodeJS.

Essa versão feita em Go é feita para suportar em qualquer máquina, sem outra linguagem intermediaria.
Só precisa do GO para compilar(download da versão compilada será disponível.)

## Por que não C++?
Sim, pensei em fazer inteiramente em C++, por outro lado... Eu estou estudando o Go, e o que seria melhor para praticar se não um projeto em Go?
O GO foi escolhido por este motivo.
</div>

# Commandos

## file
Comando inspirado no "touch" e no "cat", ele tem a função de escrever e arquivos.

```shell
$ file <flag?> <filename> <content>
```
### Valor de flag

| Valor | Função |
| ------|--------  |
| -r | Ler arquivos | 
| -w | Escrever arquivo|
| -ow | Sobrescrever arquivo|
| -h | Ajuda|
| * | Valor vai para file name|

Qualquer valor que não seja estes da tabela, é classificado como *, e ele vai para filename, enquanto o flag fica como "auto", então ele irá ler o arquivo, se existir, se não, criará ele.

```shell
$ file texto.txt
```
```txt
Se existir - ler ex.: Olá, mundo!
Se não - escreve(como não há valor de content):  
// This file is writing for file(Go Commands), github.com/jefferson-developer-it/go-cmd
```

```shell
$ file texto.txt "Hello, World!"
```

Se existir - ler ex.: Hello, World!
Se não - escreve ex.: "Hello, World!"

```shell
$ file -w texto.txt "Hello, Go!"
```

```txt
Se existir:
"This file exist!! Use 'file -ow <filename> <content>' for overwriting!"
```

```shell
$ file -ow texto.txt "Hello, Go!"
```

Sobrescreve o arquivo existente.

```shell
$ file -r texto.txt "Hello, Go!"
```
```
Valor de "<content>" é ignorado e ler o conteúdo do arquivo.
```

## c

Limpa a tela do terminal, equivalente ao cls do cmd(windows) e ao clear do Linux/Mac.

## byte e bytestr

Converte texto string para byte array:

```shell
$ byte  byte "Hello, Go"
output: [72 101 108 108 111 44 32 71 111]
```

Temos o bytestr que converte byte para string:
```shell
$ bytestr [ 79, 108, 196, 161, 32, 74, 101, 102, 102, 33]
outpu: Olá Jeff!
```
Ou
```shell
$ bytestr 79, 108, 196, 161, 32, 74, 101, 102, 102, 33
outpu: Olá Jeff!
```

## lj

Equivalente ao ls(só não tão completo), ler o diretório:

```shell
$ lj ./commands
Name       Type  Size
-----      ----- -----
byte.go    file  466kb
bytestr.go file  899kb
c.go       file  478kb
file.go    file  1758kb
lj.go      file  888kb
percent.go file  570kb
```

## percent

Faz cálculos de porcentagens:
```shell
$ percent <base?> <percent?> 
```

Exemplos

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