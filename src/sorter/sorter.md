## 排序算法比较Demo
从命令行指定输入的数据文件和输出的数据文件，并指定对应的排序算法  
`USAGE: sorter –i <in> –o <out> –a <qsort|bubblesort>`  

一个具体的执行过程如下：  

`$ ./sorter –I in.dat –o out.dat –a qsort`  
>`The sorting process costs 10us to complete.`  

### 目录结构
`go_language_program` 是我GOPATH路径  

其中sorter.go是主程序，qsort.go用于实现快速排序，bubblesort.go用于实现冒泡排序。  
__注意：由于我的目录结构跟作者的不同(作者是将`sorter`作为GOPATH),因此我将主程序(`sorter.go`)引入的两个排序包做了改变__  
*作者代码：*
```
"algorithm/bubblesort"
"algorithm/qsort"
```
*我的代码：*
```
"sorter/algorithm/bubblesort"
"sorter/algorithm/qsort"
```
```
<go_language_program>
    |--<src>
        |--<sorter>
            |--sorter.go
            |--<algorithm>
                |--<qsort>
                    |--qsort.go
                    |--qsort_test.go
                |--<bubblesort>
                    |--bubblesort.go
                    |--bubblesort_test.go
    |--<pkg>
    |--<bin>
```

### 构建与执行
#### 构建
```
$ go build sorter/algorithm/qsort
$ go build sorter/algorithm/bubblesort
$ go test sorter/algorithm/qsort
ok sorter/algorithm/qsort 0.007s
$ go test sorter/algorithm/bubblesort
ok sorter/algorithm/bubblesort 0.013s
$ go install sorter/algorithm/qsort
$ go install sorter/algorithm/bubblesort
$ go build sorter
$ go install sorter
```
如果没有出现任何问题，那么通过执行这些命令，我们应该能够在src的同一级目录下看到两
个目录——bin和pkg，其中pkg目录下放置的是bubblesort.a和qsort.a, bin目录下放置的是sorter的二
进制可执行文件。

#### 执行
因为sorter接受的是一个文件格式的输入，所以需要准备这样的一个文件。我们可以在sorter
所在的bin目录内创建一个unsorted.dat文本文件，按一行一个整数的方式填入一些数据后保存。
sorted.dat会由程序自动创建，因此不需要事先创建。  
```
$ cd bin
$ cat unsorted.dat
123
3064
3
64
490
1
23
5331
2
7
4
2
132
$ ./sorter -i unsorted.dat -o sorted.dat -a qsort
infile = unsorted.dat outfile = sorted.dat algorithm = qsort
The sorting process costs 3us to complete.
$ ./sorter -i unsorted.dat -o sorted.dat -a bubblesort
infile = unsorted.dat outfile = sorted.dat algorithm = bubblesort
The sorting process costs 2us to complete.
$ cat sorted.dat
1
2
2
3
4
7
23
64
123
132
490
3064
5331
```