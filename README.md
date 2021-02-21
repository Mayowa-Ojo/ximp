XIMP

> XIMP is a simple tool to generate sample env files from your development or production env.

### Installation
[Go 1.13]
```shell
$ mkdir <dir>
$ cd <dir>
$ git clone https://github.com/Mayowa-Ojo/ximp .
$ go build -o <dir> // add this to your $PATH
$ <OR>
$ go install
```

### Usage
First argument to the executable is the relative path to your target env. If no argument is provided, it looks for a <.env.development> in the current folder.

```shell
$ ximp ./env.dev
```
