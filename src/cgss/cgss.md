## 一个棋牌游戏服务器Demo
从命令行指定输入指定命令:  
```
Commands:
    login <username><level><exp>
    logout <username>
    send <message>
    listplayer
    quit(q)
    help(h)
```

一个具体的执行过程如下：  
```
$ go run cgss.go
Casual Game Server Solution
A new session has been created successfully.
Commands:
login <username><level><exp>
logout <username>
send <message>
listplayer
quit(q)
help(h)
Command> login Tom 1 101
Command> login Jerry 2 321
Command> listplayer
1 : &{Tom 1 101 0 <nil>}
2 : &{Jerry 2 321 0 <nil>}
Command> send Hello everybody.
Tom received message: Hello everybody.
Jerry received message: Hello everybody.
Command> logout Tom
Command> listplayer
1 : &{Jerry 2 321 0 <nil>}
Command> send Hello the people online.
Jerry received message: Hello the people online.
Command> logout Jerry
Command> listplayer
Failed. No player online.
Command> q
$
```

### 目录结构
`go_language_program` 是我GOPATH路径  
```
<go_language_program>
    |--<src>
        |--<cgss>
            |--cgss.go
            |--<cg>
                |--center.go
                |--centerclient.go
                |--player.go
            |--<ipc>       
                |--client.go
                |--server.go    
                |--ipc_test.go    
    |--<pkg>
    |--<bin>
```