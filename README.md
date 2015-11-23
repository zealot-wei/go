package main

import (
        "crypto/md5"
        "encoding/hex"
        "fmt"
        "io/ioutil"
        "os"
)

func readFile(path string) []byte {
        fp, err := os.Open(path)
        if err != nil {
                fmt.Printf("Open file %s error.", path)
                return nil
        }
        defer fp.Close()

        fd, err := ioutil.ReadAll(fp)
        if err != nil {
                fmt.Printf("Read from file:%s error", path)
                return nil
        }
        return fd
}

func md5sum(path string) string {
        fc := readFile(path)
        h := md5.New()
        h.Write(fc)
        cipherStr := h.Sum(nil)
        return hex.EncodeToString(cipherStr)
}

func main() {
        c := readFile("/home/work/.vimrc")

        fmt.Println("Read from file:")
        fmt.Println(string(c))
        fmt.Println(md5sum("/data/centos7.tgz"))

}
