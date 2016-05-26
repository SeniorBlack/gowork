package main

import "bufio"
import "flag"
import "fmt"
import "io"
import "os"
import "strconv"

import "algorithm/bubblesort"

var infile *string = flag.String("i", "infile", "File contains values for sorting")
var outfile *string = flag.String("o", "outfile", "File to receive sorte values")
var algorithm *string = flag.String("a", "qsort", "Sort algorithm")
var help *string = flag.String("h", "", "USAGE: sorter –i <in> –o <out> –a <qsort|bubblesort>")

func readValues(infile string) (values []int, err error) {
	file, err := os.Open(infile)
	if err != nil {
		fmt.Println("Failed to open the input file", infile)

	}

	defer file.Close() //遇到错误，关闭资源

	br := bufio.NewReader(file) //将文件读入缓存，得到对象

	values = make([]int, 0) //创建切片,初始化切片，长度为0,容量为0

	for {
		line, isPrefix, err1 := br.ReadLine()

		if err1 != nil {
			if err1 != io.EOF {
				err = err1
			}
			break
		}

		if isPrefix {
			fmt.Println("A too lang line, seems unexpected.")
			return
		}

		str := string(line) //将字符数组为字符串

		value, err1 := strconv.Atoi(str) // 将字符串转换为整数

		if err1 != nil {
			err = err1
			return
		}
		values = append(values, value) //append方法想切片values中添加元素

	}
	return
}

func writeValues(values []int, outfile string) error {
	file, err := os.Create(outfile)
	if err != nil {
		fmt.Println("Failed to create the output file ", outfile)
		return err
	}

	defer file.Close()

	for _, value := range values {
		str := strconv.Itoa(value)
		file.WriteString(str + "\n")
	}
	return nil
}

func main() {
	flag.Parse()

	if infile != nil {
		fmt.Println("infile =", *infile, "outfile =", *outfile, "algorithm =",
			*algorithm)
	}

	//读取
	values, err := readValues(*infile)
	if err == nil {
		fmt.Println("Read values:", values)
	} else {
		fmt.Println(err)
	}
	//排序
	bubblesort.BubbleSort(values)
	//写入
	err1 := writeValues(values, *outfile)
	if err1 == nil {
		fmt.Println("write successfully!")
	} else {
		fmt.Println(err1)
	}
}
