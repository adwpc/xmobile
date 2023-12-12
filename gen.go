// gen.go
package main

import (
	"io/ioutil"
	"math/rand"
	"os"
	"regexp"
	"strings"
	"time"
)

func replace(srcFile, dstFile, prefix string, keyMap []string) {
	content, err := ioutil.ReadFile(srcFile)
	if err != nil {
		panic(err)
	}
	s := string(content)
	for _, key := range keyMap {
		dstStr := regexpReplace(key, prefix)
		s = strings.ReplaceAll(s, key, dstStr)
	}

	f, err := os.Create(dstFile)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	_, err = f.WriteString(s)
	if err != nil {
		panic(err)
	}
}

func regexpReplace(str, prefix string) string {
	re, err := regexp.Compile("go")
	if err != nil {
		panic(err)
	}
	return re.ReplaceAllString(str, prefix+"_go")
}

func random(length int) string {
	rand.Seed(time.Now().UnixNano())
	chars := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	return b.String()
}

func replaceAll() {
	prefix := random(6)

	replace("bind/java/context_android.c", "bind/java/context_android.c", prefix, []string{"Java_go_Seq_setContext"})
	replace("bind/java/Seq.java", "bind/java/Seq.java", prefix, []string{"package go;", "import go.Universe;", "System.loadLibrary(\"gojni\");"})
	replace("bind/java/seq_android.c.support", "bind/java/seq_android.c.support", prefix, []string{"Java_go_Seq", "go/Seq"})
	replace("bind/java/SeqBench.java", "bind/java/SeqBench.java", prefix, []string{"package go;"})
	replace("bind/genjava.go", "bind/genjava.go", prefix, []string{"return \"go\""})
	replace("cmd/gobind/gen.go", "cmd/gobind/gen.go", prefix, []string{"\"java\", \"go\""})
}

func main() {
	// for java symbol conflix
	replaceAll()
}

//go:generate go run gen.go
