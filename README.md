# net-cat
## About Project
This project consists on recreating the `NetCat` in a `Server-Client` Architecture that can run in a server mode on a specified port listening for incoming connections, and it can be used in client mode, trying to connect to a specified port and transmitting information to the server.

NetCat, nc system command, is a command-line utility that reads and writes data across network connections using TCP or UDP. It is used for anything involving TCP, UDP, or UNIX-domain sockets, it is able to open `TCP` connections, send UDP packages, listen on arbitrary `TCP` and UDP ports and many more.

To see more information about NetCat inspect the manual `man nc`.
## Usage
```
$ go run . $port 
```
### Example
- `1st` Terminal
```bash
$ go run main.go
Listening on the port :8989
$ 
```
- `2nd` Terminal
```bash
$ nc localhost 8989
Welcome to TCP-Chat!
         _nnnn_
        dGGGGMMb
       @p~qp~~qMb
       M|@||@) M|
       @,----.JM|
      JS^\__/  qKL
     dZP        qKRb
    dZP          qKKb
   fZP            SMMb
   HZM            MMMM
   FqM            MMMM
 __| ".        |\dS"qML
 |    `.       | `' \Zq
_)      \.___.,|     .'
\____   )MMMMMP|   .'
     `-'       `--'
[ENTER YOUR NAME]: Dias
[2021-10-27 19:14:10][Dias]:Hello 
Adilet has joined our chat...
[2021-10-27 19:14:43][Adilet]:Hello
[2021-10-27 19:14:43][Dias]:How Are you?
[2021-10-27 19:15:31][Adilet]:I am fine, and you?
[2021-10-27 19:15:51][Dias]:Good
Adilet has left our chat...
[2021-10-27 19:16:10][Dias]:^C
$ 
```
- `3rd` Terminal
```bash
$ nc localhost 8989
Welcome to TCP-Chat!
         _nnnn_
        dGGGGMMb
       @p~qp~~qMb
       M|@||@) M|
       @,----.JM|
      JS^\__/  qKL
     dZP        qKRb
    dZP          qKKb
   fZP            SMMb
   HZM            MMMM
   FqM            MMMM
 __| ".        |\dS"qML
 |    `.       | `' \Zq
_)      \.___.,|     .'
\____   )MMMMMP|   .'
     `-'       `--'
[ENTER YOUR NAME]: Adilet
[2021-10-27 19:14:20][Dias]:Hello
[2021-10-27 19:14:32][Adilet]:Hello
[2021-10-27 19:15:04][Dias]:How Are you?
[2021-10-27 19:15:04][Adilet]:I am fine, and you?
[2021-10-27 19:15:31][Dias]:Good
[2021-10-27 19:16:03][Adilet]:^C
$ 
```