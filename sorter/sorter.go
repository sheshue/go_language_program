/**
 * 获取命令行输入
 * 从对应文件中读取输入数据
 * 调用对应的排序函数
 * 将排序的结果输出到对应的文件中
 * 打印排序所花费时间的信息
 */
package main

import (
	"flag"
	"fmt"
)

var infile *string = flag.String("i", "infile", "file contains values for sorting")
var outfile *string = flag.String("o", "outfile", "file to receive sorted values")
var algorithm *string = flag.String("a", "qsort", "sort algorithm")

func main() {
	flag.Parse()
	if infile != nil {
		fmt.Println("infile = ", *infile, "outfile = ", *outfile, "algorithm =", *algorithm)
	}
}
