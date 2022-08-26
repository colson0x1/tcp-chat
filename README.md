# TCP Chat

> I'm using go net package as well as core routines and channels.

> Transmission control protocol (TCP) is a major protocol of the internet. It sits above the network layer and provides transport mechanism for application layer protocols such as http, smtp, irc etc.

### Core logic of the app
> Once the client connects to the server, following commands can be used. Each commands starts with / <br/>
`/username <name>` - sets the name else user stays anonymous. <br/>
`/join <name>` - join a room. if room doesn't exist, new room will be created. user can only be at one room at one room concurrently. <br/>
`/rooms` - shows list of available rooms to join. <br/>
`/msg <msg>` - broadcast message to all members in a room. <br/>
`/exit` - disconnects from chat server.


> The application consists of the following parts:
> First, client holds a name of the user, current connection, current room. Then,  there's a room type which contains the list of the members on the room. And there's, comment, which means comment from client to server such as join colson and all other comments from the list. Then there's a centralized server which will be responsible for processing the incoming commands as well as storing rooms and clients. And there's a TCP server on top of that which accepts incoming connections.

***

### To run the program:
<ul>
  <li>
    <span style="color: #f03c15;"> Make sure go is installed on your machine. https://go.dev/dl/ </span> 
  </li>
</ul>

```go
$ go run main.go
```
