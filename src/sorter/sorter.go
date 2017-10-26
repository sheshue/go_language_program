/**
 * 获取命令行输入
 * 从对应文件中读取输入数据
 * 调用对应的排序函数
 * 将排序的结果输出到对应的文件中
 * 打印排序所花费时间的信息
 */
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"

	"sorter/algorithm/bubblesort"
	"sorter/algorithm/qsort"
)

var infile *string = flag.String("i", "infile", "file contains values for sorting")
var outfile *string = flag.String("o", "outfile", "file to receive sorted values")
var algorithm *string = flag.String("a", "qsort", "sort algorithm")

func main() {
	flag.Parse()
	if infile != nil {
		fmt.Println("infile = ", *infile, "outfile = ", *outfile, "algorithm =", *algorithm)
	}
	values, err := readValues(*infile)
	if err == nil {
		t1 := time.Now()
		switch *algorithm {
		case "qsort":
			qsort.QuickSort(values)
		case "bubblesort":
			bubblesort.BubbleSort(values)
		default:
			fmt.Println("sorting algorithm", *algorithm, "is either unknow or unsu")
		}
		t2 := time.Now()
		fmt.Println("the sorting process costs", t2.Sub(t1), "to complete.")
		writeValues(values, *outfile)
	} else {
		fmt.Println(err)
	}
}

func readValues(infile string) (values []int, err error) {
	file, err := os.Open(infile)
	if err != nil {
		fmt.Println("fail to open the input file", infile)
		return
	}
	defer file.Close()
	br := bufio.NewReader(file)
	values = make([]int, 0)
	for {
		line, isPrefix, err1 := br.ReadLine()
		if err1 != nil {
			if err1 != io.EOF {
				err = err1
			}
			break
		}

		if isPrefix {
			fmt.Println("a too long line, seems unexprected.")
			return
		}
		str := string(line)
		value, err1 := strconv.Atoi(str)
		if err1 != nil {
			err = err1
			return
		}
		values = append(values, value)
	}
	return
}

func writeValues(values []int, outfile string) error {
	file, err := os.Create(outfile)
	if err != nil {
		fmt.Println("fail to create the output file ", outfile)
		return err
	}
	defer file.Close()
	for _, value := range values {
		str := strconv.Itoa(value)
		file.WriteString(str + "\n")
	}
	return nil
}
