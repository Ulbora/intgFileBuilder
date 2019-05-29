package filebuilder

import (
	"encoding/csv"
	"fmt"
	"testing"
)

func TestAllSupplierDirs(t *testing.T) {
	var b Builder
	var cb CsvFileBuilder
	b = &cb
	res := b.ReadAllSupplierDirs("./testdir")
	if (*res)[0].Name != "test1" {
		t.Fail()
	}
	if (*res)[0].Files[0].FullName != "./testdir/test1/test11.csv" {
		t.Fail()
	}
	fcont := b.ReadSourceFile((*res)[0].Files[0].FullName)
	fmt.Println("len: ", len(fcont))
	if len(fcont) == 0 {
		t.Fail()
	}
	//fmt.Println("source file in TestAllSupplierDirs", fcont)
}

func TestReadSourceFile(t *testing.T) {
	var b Builder
	var cb CsvFileBuilder
	b = &cb
	res := b.ReadSourceFile("./testdir/test1/test11.csv")
	fmt.Println("len: ", len(res))
	if len(res) == 0 {
		t.Fail()
	}
	//fmt.Println("source file in read file test", res)
}

func TestReadFileBadFile(t *testing.T) {
	var b Builder
	var cb CsvFileBuilder
	b = &cb
	res := b.ReadSourceFile("./testdir/test1/test11_bad.csv")
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
	if len(res) != 0 {
		t.Fail()
	}
}

func TestSaveCartFile(t *testing.T) {
	var b Builder
	var cb CsvFileBuilder
	cb.OutputDir = "./cartFileTest"
	b = &cb
	files := b.ReadAllSupplierDirs("./testdir")
	fmt.Println("AllSupplierDirs: ", files)
	fcont := b.ReadSourceFile((*files)[0].Files[0].FullName)
	var cf CartCsvFile
	cf.SupplierDir = "test1"
	cf.FileName = (*files)[0].Files[0].Name
	cf.Content = fcont
	fmt.Println("CartCsvFile: ", cf)
	// var cb CsvFileBuilder
	suc := b.SaveCartFile(cf)
	if suc != true {
		t.Fail()
	}

}

func TestCreateBadFile(t *testing.T) {
	//r := csv.NewReader(strings.NewReader("0"))
	f, err := createFile("./cartFileTest/test5/test.csv")
	fmt.Println(f)
	fmt.Println(err)
	if err == nil {
		t.Fail()
	}
}

func TestCreateBadDir(t *testing.T) {
	suc := createDir("./cartFileTest/test6/test.csv")
	if suc != true {
		t.Fail()
	}
}

func TestLogWriteAllError(t *testing.T) {
	//var fileName = "cartFileTest/test,"
	f, _ := createFile("")
	defer f.Close()
	w := csv.NewWriter(f)
	var content = [][]string{
		{"first_name", "last_name", "username"},
		{"Rob", "Pike", "rob"},
		{"Ken", "Thompson", "ken"},
		{"Robert", "Griesemer", "gri"},
	}
	fmt.Println("w bin: ", w)
	w.WriteAll(content)
	logError := logWriteAllError(w)
	if !logError {
		t.Fail()
	}
}
