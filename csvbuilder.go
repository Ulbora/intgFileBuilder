package filebuilder

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
)

/*
 Copyright (C) 2019 Ulbora Labs LLC. (www.ulboralabs.com)
 All rights reserved.

 Copyright (C) 2019 Ken Williamson
 All rights reserved.

 This program is free software: you can redistribute it and/or modify
 it under the terms of the GNU General Public License as published by
 the Free Software Foundation, either version 3 of the License, or
 (at your option) any later version.
 This program is distributed in the hope that it will be useful,
 but WITHOUT ANY WARRANTY; without even the implied warranty of
 MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 GNU General Public License for more details.
 You should have received a copy of the GNU General Public License
 along with this program.  If not, see <http://www.gnu.org/licenses/>.

*/

//ParseFiles ParseFiles
func (b *CsvFileBuilder) ParseFiles(dir string) *SourceFiles {
	var rtn SourceFiles
	//fmt.Println(dir)
	res := readFileDir(dir)
	//fmt.Println("parser res: ", res)
	for _, sfile := range *res {
		//fmt.Println("files in parserfile", sfile)
		for _, file := range sfile.Files {
			fmt.Println("source file in parserfile", file)
			sourceFileContent := readSourceFile(file.FullName)
			var sf Sourcefile
			sf.Name = file.Name
			sf.Content = sourceFileContent
			rtn.Files = append(rtn.Files, sf)
			//fmt.Println("source file in read file", sourceFileContent)
		}
	}
	return &rtn
}
func readFileDir(dir string) *[]SupplierDir {
	var rtn []SupplierDir
	//fmt.Println("dir: ", dir)
	files, err := ioutil.ReadDir(dir)
	if err == nil {
		for _, file := range files {
			//fmt.Println("file name: ", file.Name())
			if file.IsDir() {
				var sd SupplierDir
				sd.Name = file.Name()
				rtn = append(rtn, sd)
			}
		}
		for c, spd := range rtn {
			//fmt.Println("spd: ", spd)
			var sdirname = dir + string(filepath.Separator) + spd.Name
			//fmt.Println("dir name: ", sdirname)
			sfiles, err := ioutil.ReadDir(sdirname)
			if err == nil {
				for _, sfile := range sfiles {
					if !sfile.IsDir() {
						//fmt.Println("sfile: ", sfile)
						var spfile SupplierFile
						spfile.Name = sfile.Name()
						spfile.FullName = sdirname + string(filepath.Separator) + sfile.Name()
						spd.Files = append(spd.Files, spfile)
					}
				}
			}
			rtn[c].Files = spd.Files
			//fmt.Println("spd: ", spd)
		}
		//fmt.Println("rtn: ", rtn)
	}
	return &rtn
}

func readSourceFile(file string) [][]string {
	sourceFile, err := ioutil.ReadFile(file)
	if err != nil {
		log.Println("source file in read err", err)
	}
	r := csv.NewReader(strings.NewReader(string(sourceFile)))
	records := csvReader(r)
	//fmt.Println("records", records)
	return records
}

func csvReader(r *csv.Reader) [][]string {
	records, err := r.ReadAll()
	if err != nil {
		log.Println("csv error: ", err)
	}
	return records
}
