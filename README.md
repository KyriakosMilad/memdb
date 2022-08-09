# memdb

Golang in-memory database.

## Installation

    $ go install github.com/KyriakosMilad/memdb

verify installation:

    $ memdb

## Start the server

You can start the server using `memdb` command

    $ memdb

The default listening port is 3636, But if you want you can specify the listening port using `port` option

    $ memdb -port 3636

## Connect to the server

    $ telnet 127.0.0.1 3636

127.0.0.1 is the host 

3636 is the listening port

## Commands

### set

Set value in the database.

syntax `set [key] [value]`

example

    $ set x 5

output:

     OK

### get

Get value from the database using key.

syntax: `get [key]`

if key exists it's value will be returned, if key does not exist `key [key] not found` will be returned.

example

    $ get x

output:

    5

example2

    $ get y

output2:

    key y not found


### delete
Delete value from the database using key.

syntax `delete [key]`

example

    $ delete x

output:

    OK


now if you try to get x's value

    $ get x

output:

    key x not found

### exit

close the connection to the server

    $ exit

## Stop the server

You can stop the server by stopping the running process (CTRL + C)

    $ ^C

output
```
stopping memdb
removing all clients
removed all clients successfully
saving database on the desk
successfully saved database on the disk
closing the tcp listener
closed tcp listener successfully
exiting
```
