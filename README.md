# araknet
A blockchain built with Golang

## Build Araknet CLI locally

* Have Go installed.
* Clone Araknet program inside your `$GOPATH/src`.
* `go get` relevant depedencies inside the `go.sum` file.
* Build the program by using this command inside your terminal `go build .`
* If on Linux you will have to make the Araknet program executable.

## CLI

List of commands in the program:
```bash

Get all commands
    - ./araknet help

Create a wallet
    - ./araknet createwallet

Create a blockchain (if blockchain does not exists)
    - ./araknet createblockchain -address <address>

Get balance for a specific address
    - ./araknet getbalance -address <address>

Send out Araknet coins
    - ./araknet send -from <your-address> -to <recipient-address> -amount <number of coins>

Print all blocks of the blockchain
    - ./araknet printchain

List all addresses created
    - ./araknet listaddresses
```

List of errors to encounter when using CLI
```bash

<Date(YYYY/MM/DD HH:MM:SS)> ERROR: Not enough funds
    - Occurs when you try to send Araknet coins but you lack enough funds

Blockchain already exists
    - You are trying to create a blockchain from a genesis block but one has already been created

```