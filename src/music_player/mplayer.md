## 第三章音乐播放器Demo

### 目录结构
`go_language_program` 是我GOPATH路径  
```
<go_language_program>
    |--<src>
        |--<music_player>
            |--mplayer.go
            |--<library>
                |--manager.go
                |--manager_test.go
            |--<mp>
                |--play.go
                |--mp3.go
                |--wav.go
    |--<pkg>
    |--<bin>
```

### 构建与执行
#### 构建
```
$ go build music_player/library
$ go build music_player/mp
$ go test music_player/library
ok music_player/library 0.007s
$ go build music_player
$ go install sorter
```
如果没有出现任何问题，那么通过执行这些命令，我们应该能够在src的同一级目录下看到两
个目录——bin和pkg，其中pkg目录下放置的是library.a和mp.a, bin目录下放置的是music_player的二
进制可执行文件。

#### 执行
```
$ cd bin
$ ./music_player 
Enter following commands to control the player:
lib list -- View the existing music lib
lib add <name><artist><source><type> -- Add a music to the music lib
lib remove <name> -- Remove the specified music from the lib
play <name> -- Play the specified music
Enter command-> lib add HugeStone MJ ~/MusicLib/hs.mp3 MP3
Enter command-> play HugeStone
Playing MP3 music ~/MusicLib/hs.mp3
..........
Finished playing ~/MusicLib/hs.mp3
Enter command-> lib list
1 : HugeStone MJ ~/MusicLib/hs.mp3 MP3
Enter command-> lib view
Enter command-> q
```