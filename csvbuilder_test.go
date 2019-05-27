package filebuilder

import (
	"encoding/csv"
	"fmt"
	"testing"
)

func TestParseFiles(t *testing.T) {
	files := ParseFiles("./testdir")
	fmt.Println("files len in parse files: ", len(files.Files))
	fmt.Println("source files in parse files", files)
	if len(files.Files) != 4 {
		t.Fail()
	}
}

func TestReadFileDir(t *testing.T) {
	res := readFileDir("./testdir")
	if (*res)[0].Name != "test1" {
		t.Fail()
	}
	if (*res)[0].Files[0].FullName != "./testdir/test1/test11.csv" {
		t.Fail()
	}
}

func TestReadFile(t *testing.T) {
	res := readSourceFile("./testdir/test1/test11.csv")
	fmt.Println("len: ", len(res))
	if len(res) == 0 {
		t.Fail()
	}
	//fmt.Println("source file in read file test", res)
}

func TestReadFileBadFile(t *testing.T) {
	res := readSourceFile("./testdir/test1/test11_bad.csv")
	fmt.Println("len: ", len(res))
	if len(res) != 0 {
		t.Fail()
	}
	fmt.Println("source file in read bad file test", res)
}

func TestReadFileBadCsvFile(t *testing.T) {
	//r := csv.NewReader(strings.NewReader("0"))
	var r csv.Reader
	res := csvReader(&r)
	fmt.Println("csv: ", res)
}
