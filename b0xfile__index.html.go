// Code generaTed by fileb0x at "2022-12-08 01:03:28.865623 +0300 MSK m=+0.430838251" from config file "b0x.yaml" DO NOT EDIT.
// modified(2022-12-08 00:20:58.273338806 +0300 MSK)
// original path: swagger-ui/dist/index.html

package swaggerFiles

import (
	"bytes"
	"compress/gzip"
	"io"

	"os"
)

// FileIndexHTML is "/index.html"
var FileIndexHTML = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x92\x4f\x8b\xdb\x30\x10\xc5\xef\xf9\x14\x13\xdd\x65\x91\x04\x42\x0f\xb2\x2f\xfd\x43\x0b\x2d\x2d\x34\x39\xf4\x28\x5b\x13\x7b\xba\x8a\x6c\xa4\x71\xd6\xc9\xa7\x5f\x6c\x79\x09\x61\x59\x02\x9b\x93\x87\xf7\xfc\x7e\x03\x6f\xa4\x97\x52\xc2\xf7\xdd\xaf\x9f\x70\x68\x03\x44\x36\x4c\x15\x58\x8a\x1c\xa8\xec\x99\x5a\x0f\x65\xef\xad\x43\x28\x7b\x72\x16\xa4\x2c\x16\x7a\xf9\xe5\xf7\xe7\xdd\xbf\x3f\x5f\xa1\xe1\xa3\x2b\x16\x7a\xfc\x80\x33\xbe\xce\x05\x7a\x51\x2c\x00\x74\x83\xc6\x8e\x03\x80\x3e\x22\x1b\xa8\x1a\x13\x22\x72\x2e\xf6\xbb\x6f\xf2\x93\x98\x2d\x26\x76\x58\xfc\x7d\x36\x75\x8d\x01\xf6\x3f\xb4\x4a\x4a\x72\x1d\xf9\x27\x08\xe8\x72\x11\xf9\xec\x30\x36\x88\x2c\x80\xcf\x1d\xe6\x82\x71\x60\x55\xc5\x28\xa0\x09\x78\xc8\x45\xa6\x62\xa2\xc8\x9e\xb2\x49\x57\x1f\xa0\x90\xb7\x38\xbc\x17\xa7\xaa\xf5\xaf\x41\x3a\x9a\x1a\x55\xe7\xeb\xeb\xfe\x83\x39\x8d\x7f\xc8\xcd\x7a\xd8\xac\xb3\xc9\x8a\x74\xc1\x98\x8b\x49\x79\x80\xb8\xda\x0e\xab\xed\x0d\x71\x52\x66\xa2\x56\xa9\xeb\x71\x2c\x5b\x7b\x9e\xb7\x58\x3a\x01\xd9\x5c\x5c\x6b\x11\x85\x56\x96\x4e\xb3\x1f\xab\x40\x1d\x43\x0c\xd5\x6d\x79\x32\x9d\x3b\xfb\x1f\xc5\x9b\xa3\x81\x56\x29\x76\x8f\x11\xd9\x78\x6b\x5c\xeb\x51\x76\x01\x23\xf2\x03\x38\xf2\xc4\x64\x1c\x5d\x30\xdc\xa7\x68\x95\x2a\xd0\x2a\x3d\xcd\x97\x00\x00\x00\xff\xff\xce\x44\x52\x76\xde\x02\x00\x00")

func init() {

	rb := bytes.NewReader(FileIndexHTML)
	r, err := gzip.NewReader(rb)
	if err != nil {
		panic(err)
	}

	err = r.Close()
	if err != nil {
		panic(err)
	}

	f, err := FS.OpenFile(CTX, "/index.html", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(f, r)
	if err != nil {
		panic(err)
	}

	err = f.Close()
	if err != nil {
		panic(err)
	}
}
