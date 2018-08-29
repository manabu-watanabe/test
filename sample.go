package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
    "strings"
)

var sc = bufio.NewScanner(os.Stdin)

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	// go build sample.go
	// winfowsだと、ビルドするとexeファイルになるので、そのまま実行
	// winfows以外だと、単純に実行するには、下記
	// go run sample.go
	// ソースのフォーマット
	// go fmt sample.go
    //
    // 課題
    // りんごとみかんの箱詰め
    // りんご(A)、みかん(B)が、最初に指定された数と同じだけ並んでいます。
    // 連続したりんご、及びみかんごとに、箱詰めをした場合、何個箱が必要か？
    // > 10
    // > AAAABBBBAA
    // < 3
    

	cnt, _ := strconv.Atoi(nextLine())
    strs := strings.Split(nextLine(),"")
    out := 0
    es := ""
    for i, s := range strs { // range は、位置と、値を返す
        if (i >= cnt) {
            break
        }
        if es == s {
            continue
        }
        es = s
        out++
    }
    fmt.Printf(strconv.Itoa(out))
}
il
