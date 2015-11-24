//This file implement some md5 function
package main

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func ReadFile(path string) ([]byte, error) {
	fp, err := os.Open(path)
	if err != nil {
		fmt.Printf("Open file %s error\n", path)
		return nil, err
	}
	defer fp.Close()

	fd, err := ioutil.ReadAll(fp)
	if err != nil {
		fmt.Printf("Read from file:%s error\n", path)
		return nil, err
	}
	return fd, nil
}

//Compute md5 of a small file. Read whole file in memory
func MD5SmallFile(path string) (string, error) {
	fc, err := ReadFile(path)
	if err != nil {
		return "", err
	}
	h := md5.New()
	h.Write(fc)
	cipherStr := h.Sum(nil)
	return hex.EncodeToString(cipherStr), nil
}

//Compute md5 of large file
func Md5(path string) (string, error) {
	if info, err := os.Stat(path); err != nil {
		return "", err
	} else if info.IsDir() {
		fmt.Printf("Md5 args cannot be a directory\n")
		return "", nil
	}

	fp, err := os.Open(path)
	if err != nil {
		fmt.Printf("MD5 Open file: %s failed\n", path)
		return "", err
	}
	defer fp.Close()

	r := bufio.NewReader(fp)
	buf := make([]byte, 1024*1024) //Buffer 512K

	h := md5.New()
	for {
		n, err := r.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			return "", err
		}
		h.Write(buf[0:n])
	}
	chksum := hex.EncodeToString(h.Sum(nil))
	return chksum, nil
}

/*Read md5 file and check sum,it's function's like md5sum -c file.md5
  file content's like:
a63258692d102095d8fd1b322e317d52 system.go
a6d0bf86016972ea59b64bfecc4301dd md5.go
*/
func Md5CheckSum(md5file string) bool {
	md5content, err := ReadFile(md5file)
	if err != nil {
		fmt.Printf("md5 checksum open file: %s failed\n", md5file)
		return false
	}
	lines := strings.Split(string(md5content), string('\n'))
	for _, line := range lines {
		if len(line) < 10 {
			continue
		}
		tline := strings.Split(line, " ")
		tmd5, err := Md5(tline[1])
		if err != nil {
			fmt.Printf("Check md5 failed, file: %s\n", tline[1])
			return false
		}
		if tmd5 != tline[0] {
			fmt.Printf("%s md5 not match\n", tline[1])
			return false
		}
	}
	return true
}
