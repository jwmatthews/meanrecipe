// Code generated by go-bindata.
// sources:
// data/fruits.json
// data/herbs.json
// data/ingredient_corpus.txt
// data/ingredient_densities.json
// data/top_5k.txt
// data/vegetables.json
// DO NOT EDIT!

package meanrecipe

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _dataFruitsJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x53\xcd\x6e\xe6\x20\x0c\xbc\x57\xea\x3b\x44\x3d\xef\x1b\xad\xf6\xe0\x80\x4b\xfc\xc5\x60\x6a\x70\x3f\xe5\xed\x57\xed\xd6\x04\xb4\x37\xcf\xc4\x8c\xff\x26\xbf\x5f\x5f\xb6\xed\x0d\x6a\x65\x7c\xfb\xf5\x13\x2b\x05\xe9\x8e\x3e\x25\x40\x94\x1f\xb4\x43\x81\x02\x0e\x90\x79\xab\x58\x2b\xaa\x33\xc4\x3b\xaa\x5e\x0e\x19\xc2\xf9\x1f\x11\x4c\x15\x4a\x1f\x94\x48\xdc\x44\xa1\x24\x1c\x94\xe1\xf2\x4a\xae\x86\x65\x61\x14\x21\xbe\xab\x91\xab\x04\x28\xa0\xd7\x96\x91\xa5\xdc\x54\x07\x16\xab\x2e\x1b\x0e\x54\xca\x72\xc1\x84\x87\x62\x38\x88\x69\x1d\x26\x30\x66\x2c\x9d\xca\x10\x60\xb1\x38\xb7\x11\x24\x48\xb1\xd1\x83\xc2\xd2\x64\xb0\x60\x79\xbf\xe5\x96\xb1\x23\xe4\x36\x5a\x8d\xd0\xbd\x46\x54\x48\x52\xe6\xd1\xa2\x29\x81\x67\x62\x4a\x95\x6f\x15\xe4\x88\x3a\xd7\x7c\x47\x7a\x88\x0f\xf8\x4e\xe9\x27\x4a\xf2\xa0\x6d\xce\x4b\x22\x6d\xd9\x71\x52\x18\x8b\xfa\x8e\xe7\x16\x92\xc1\xa7\x8b\x1e\x52\xf0\x8a\xf8\x74\x68\xe1\xe4\x45\xe9\x01\xe1\x9c\x1f\x3f\x20\xef\xc6\x0e\xec\x61\xbb\xd7\x39\xe9\x49\xdb\x9c\x7a\x5a\xfe\x30\x70\xc4\x98\xc7\x86\x98\xb2\xbf\x62\x99\x73\xae\x70\xa0\x7f\xc9\x50\x22\xe8\x7d\xaf\x0c\x25\xb9\x71\xb3\x2d\xce\x2c\x18\xfa\x9c\x7a\x5f\x51\x98\x3e\x9d\x5d\x5c\x59\xe1\xcb\x5d\x03\x54\x18\x46\xaa\xd0\x1a\xad\x47\xab\x08\xe1\xb8\x63\x1d\xa1\x36\xca\xf7\x58\xf5\xb8\x1a\x30\x35\x87\x54\x70\xfe\x13\x2b\x5b\xf6\x50\x32\x26\x85\x72\x3b\xe5\x8b\xb9\xfb\x31\xad\x8c\xdb\xf7\xc4\xad\x23\x7a\x81\x0f\xa3\x12\xfc\x85\x02\x35\x2a\x03\xe4\xdd\x3a\xdc\xb0\xd5\x79\x41\x8a\x71\x35\xac\x4a\x38\x97\xff\xab\x01\x03\x2f\xa6\x6a\xd0\x9b\x65\xdf\x4a\xeb\xa0\xcb\x75\x5b\x57\x78\xce\xf9\x1d\x32\x28\xf1\x18\xa2\x7f\x6d\x7b\x3a\x8a\x25\x5e\xfd\xf1\x84\x8e\xfa\xaf\x87\xd7\x97\x3f\x7f\x03\x00\x00\xff\xff\xee\xe0\x8b\xf2\xba\x04\x00\x00")

func dataFruitsJsonBytes() ([]byte, error) {
	return bindataRead(
		_dataFruitsJson,
		"data/fruits.json",
	)
}

func dataFruitsJson() (*asset, error) {
	bytes, err := dataFruitsJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/fruits.json", size: 1210, mode: os.FileMode(438), modTime: time.Unix(1523666078, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dataHerbsJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x55\xc1\x72\xe3\x36\x0c\xbd\xef\xcc\xfe\x03\x67\xcf\xed\x47\xd8\xda\xdd\x6c\x1b\xa7\xf5\xc4\x69\x32\xdd\x4e\x0f\xcf\x12\x2c\x21\x26\x09\x15\x24\x9d\x51\xbe\xbe\xe3\xd8\xa4\x24\xeb\x86\xf7\x48\x82\x20\xf0\x00\xfe\xf3\xf9\x93\x31\x5f\x56\xbe\x25\xcb\x35\xbe\xfc\xf2\x01\xd7\x08\x6c\xaf\xf6\x0f\xb1\x83\x99\x12\x4f\x1d\x78\x46\xac\x31\x18\x4b\x38\x5c\xe1\x6f\xbe\x61\x78\x73\xc3\xae\xc5\x36\x52\x6c\x45\x4b\x57\x50\x75\xa4\xa7\xe2\xab\xea\xf8\x44\xe1\x0a\xee\xa0\x96\x6b\x33\xe3\x2a\xae\xc9\x0e\x19\x88\x32\x7c\x43\x3a\xbd\xa9\x62\x0b\x1f\x75\xbc\xcc\xf2\xe9\x1c\x50\xd9\x7c\x5d\x78\x66\x8a\x1e\x8e\x02\x2d\x96\xaa\x34\x73\x51\x29\x85\x72\x7f\x52\x9d\x3d\xec\x2b\xdb\x1c\xfc\xb7\x1e\xef\x12\xf3\xc3\x7e\x90\xeb\x4b\x0e\x5f\x61\x02\x7c\x44\x21\x52\x8c\x69\xf0\x0c\x53\x8b\x36\x18\x17\x86\x10\x24\x1f\xfb\x9d\xdd\x3e\x5d\xed\x7b\xf6\x9d\x69\x59\xf8\x8a\x37\x38\xd1\x24\xe2\x0d\x39\xf1\x66\x0f\xeb\x66\x44\xab\x28\x91\x5f\x18\x37\x68\xb4\x34\xa3\x4e\xa4\x7b\xf2\x39\x82\x0d\x3b\x2f\x7d\xc7\x16\x06\x2a\x0e\x71\x94\xc5\x46\x4e\x63\xdd\x1e\xa0\xaf\xa2\xc8\xf7\x3d\xb0\x8f\xd9\x4c\xed\x9b\x68\x41\x1c\x43\xda\x67\x17\x7f\x2a\xb5\xf0\x39\xb1\x5b\x68\xb0\x94\xab\xb9\x25\x65\x6b\xf3\xce\x47\x09\xe4\xa0\x79\xf1\x31\xe5\x8b\x77\x63\x0c\x3b\x9c\xa4\xec\xd8\xc1\x87\x2e\x7b\xde\x75\x1c\x8a\x2d\xaa\x54\xd4\x0b\x55\xb4\xe2\x8b\x98\x07\x97\xbd\xbd\x88\x34\x9a\x0e\xb9\xb0\x2b\xf1\xa2\x39\xdd\xab\xd7\x37\x70\x3e\xb4\xb2\x36\xf4\x5c\xe7\x73\x2b\x57\x77\x22\xb9\x12\x2b\xcf\xa1\xc4\x17\xa1\x66\x4a\xac\x02\x0e\x42\x91\x9b\xfc\xca\x0a\xae\xef\xca\xd9\x0a\x8a\x37\x0c\x23\x6a\xe0\x24\x67\x78\x6d\x51\x1f\xcd\x0d\x59\x21\x04\x2e\xce\xc8\x92\x0e\xa6\x97\xb7\x89\x94\x2f\x5c\x20\x6a\x4a\x8b\x41\xc5\xf2\xd8\x7e\xbe\x2f\x80\xbd\x87\x2b\xc9\xa9\xac\x9c\x68\xd1\x6a\x53\x57\x69\x4f\xfb\x62\xbb\x92\xa0\x3f\xb8\x25\x6b\xcf\x92\x8f\x7c\x2a\x43\x25\x79\x4e\xce\xf4\xa4\x81\xeb\xe4\x96\xed\x73\x36\xa7\xde\xbf\x93\xf7\xa5\x6e\xdf\xc9\xa7\x56\x89\x8e\x19\xb3\x6f\x49\x55\x24\xeb\xec\x4e\x09\x91\xd4\xdc\xc1\xc2\xb7\xb0\x45\xe2\x21\x2c\xd9\xcb\x64\xc9\xe0\xc3\x55\xae\xd0\x55\xf3\x66\xc6\xde\x89\xed\x51\x80\x82\x7d\x30\x72\x30\x5b\x28\x9a\xb1\xba\xe3\xc2\x8e\x2c\xbb\xd2\xeb\x1a\xe8\x63\x5f\x97\xfb\x3a\x79\xee\x49\xcd\x9e\xb4\xa8\xf7\x5e\x8e\x25\x27\xf7\xa2\x50\x76\x39\x71\x5f\x95\xa9\x31\x1b\x2e\x42\xdd\xf0\x7f\x49\x74\x14\xe0\x86\x63\x20\x98\xfa\x5c\x0d\x94\xf6\xac\xc7\x4e\xf5\xad\xfc\xda\x4e\xdf\xf3\x80\x10\xcb\xfb\x1f\xd0\x59\xec\x67\x2a\x7b\x48\x21\x42\x73\x21\xd6\x2a\x6f\xfe\x86\x7b\xe9\x38\xd2\x0d\x77\x2d\x7b\x46\xaf\xf0\x6d\x28\x28\x45\x47\xed\xd8\x40\xdc\x22\x8a\x9a\x2d\xf5\x7d\x89\x6a\xad\x78\x67\x7b\x9e\xd5\x33\xba\xea\xd8\xf2\x0d\x85\xe1\x2c\x0e\xd3\x4f\xc9\x2d\x7a\xe5\xe3\x38\xac\x7c\x3b\x3f\xb4\x25\x4d\xa7\x85\xf7\x6f\x08\xd1\xac\xc2\x82\xdf\x71\xdd\xa5\x05\x39\x9d\x31\x4f\x08\x0e\x7e\x71\xf0\x92\xc1\x0b\x55\x8b\xfa\x51\x9f\xe4\x97\xf4\x25\x8f\x0b\x7a\x2b\x8e\x5a\x85\x47\xa4\x69\x4b\x6c\xa5\xef\x67\xcd\xfc\x88\xa6\x4b\x9e\x27\x33\xb3\x44\x7a\x38\x68\xe9\xe4\x1d\x6c\x2c\xa6\x06\xf4\x98\x8e\xda\x1d\x42\xc0\x41\x91\x7f\x8a\x1d\x05\x14\xb1\xcd\x26\x69\x72\xa8\xcb\xeb\x1d\x94\x7d\x0e\xe4\x49\xfc\x11\x66\x4f\x28\x93\x35\xa9\x23\x2d\x22\xfb\xeb\x1d\xef\x39\xcc\x67\xf8\xc9\xed\xcf\x82\x18\xce\xfd\xc0\x07\x2a\xed\xf0\x82\x80\x7d\xde\xff\x77\x7a\xcf\xdf\xe0\x4f\x6a\x64\xfc\x14\x7e\x92\x52\xe8\x8e\x05\x85\xf8\xe5\xf3\xa7\x7f\xff\x0f\x00\x00\xff\xff\x0e\x4c\x07\x18\xd3\x08\x00\x00")

func dataHerbsJsonBytes() ([]byte, error) {
	return bindataRead(
		_dataHerbsJson,
		"data/herbs.json",
	)
}

func dataHerbsJson() (*asset, error) {
	bytes, err := dataHerbsJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/herbs.json", size: 2259, mode: os.FileMode(438), modTime: time.Unix(1523666078, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dataIngredient_corpusTxt = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\x56\xcd\xb2\xf4\xaa\x0d\xdc\x53\xc5\x3b\xb0\xca\x2e\x49\xe5\x3f\x79\x1c\x19\x6b\x6c\x3e\x03\xa2\x24\x98\xb9\x73\x9f\x3e\x25\x7e\x3c\x73\xee\xe6\xd4\xb1\x47\x16\x4d\xab\xd5\x92\x3f\xa9\x14\xdc\x9d\xc4\xe0\x71\x77\x7b\xff\xeb\x5b\x75\x1e\x72\xd6\xf7\x35\xf8\x4b\x1c\x3d\xfa\xcb\x90\x2b\xb9\x12\xd0\xa3\x58\xd3\xb2\x40\xac\x3d\xbc\xb8\x8a\x20\x85\x28\x3b\x69\x07\xb0\x7b\x44\x6a\xec\xf0\x38\xdc\x13\x72\x88\x11\xf4\x7f\x71\x25\x64\x7f\x5a\x73\xc4\x56\x31\xbb\x07\x23\x5a\xb3\x83\x9c\x2e\x02\x1f\xe8\x3c\x01\x0b\xba\x47\xc8\xe8\x24\x3c\x34\xf7\xc6\x74\x61\xfe\x73\x2b\x4e\x0f\x73\x55\x8a\xab\x9b\x14\x67\x0d\x85\xe8\xa8\xd4\x40\x19\xa2\xc3\xdf\x2a\x83\xaf\x6e\x83\x0c\x19\x5c\x02\x39\x71\x9f\x59\x25\x41\x8c\x0e\x62\xb4\x46\x33\xc7\xb7\x35\x1d\xc7\x3c\x07\xa1\xa7\xb6\x06\x9f\x50\x88\x41\x4f\x4d\x21\x5e\xd6\x24\xd4\xeb\xe9\x45\x5f\x0c\xca\x92\x35\x23\xb2\xdf\xd1\x9a\xad\xd5\x8a\x6c\xcd\x01\x1c\x83\xb7\xe6\x05\xfd\x91\x62\x78\xa2\xa3\x10\xad\x19\x79\x3a\x19\xd6\x50\x0e\x94\xad\x29\x58\x0a\xae\x47\xb1\x66\x8b\xe0\x2f\xb7\xde\x6e\x4c\xaf\x49\xa2\x35\xca\x99\x35\x3e\xe4\x0c\x49\x3f\xdd\xe0\x0a\xf9\x70\x85\x5e\xbb\xc6\x46\x4c\x94\xdd\xaf\x16\x3c\x5a\x53\x29\x41\x25\xad\xcb\x64\xdc\x9a\x02\x2c\x11\xdf\xf7\x77\x42\x3b\x58\x23\x5a\x19\xcf\x08\xc9\x9a\x27\x1e\x58\x61\x8b\x13\xae\xc7\x88\xfc\xb6\xe6\x08\xf9\xb8\x0f\xb0\xa6\x07\x3b\x7f\x22\x0a\x5a\xe3\x81\x99\xaa\xe2\x3a\x71\xdf\x81\xef\x1f\x36\xc4\x87\x35\x85\xea\xc4\xd1\x53\x9e\x94\x15\x41\x6e\x35\xe1\xd1\xbf\xe9\xb1\x42\x6f\x27\xd0\x14\x78\x82\x37\xe5\x0c\xa1\xe7\x3e\x83\xbf\x30\x6b\xd5\xeb\x69\x0d\x31\x1e\x90\xc9\x1a\xdf\x52\xc8\xd6\xd4\xf3\x9d\x70\xf1\x7d\xd3\xd0\x85\x01\x79\xbf\x39\x4c\x4d\x4e\x26\x4a\x9d\xbb\x08\xb9\x32\x29\x07\xa2\x78\x0a\x7a\xe8\xac\x83\xd7\xab\x9d\x08\xcf\xf7\x62\xe3\x73\x3a\x82\xe8\x0d\x5f\xc4\x1e\xa5\x22\xcb\x19\x18\x17\xe0\x02\x85\xc3\x05\x1a\x4f\x9e\x22\xd4\x0f\x70\xd5\x40\xcc\x6d\xb0\x13\x62\xb8\x31\x42\x4c\x94\x77\xb1\x26\x86\x84\xab\x64\x05\x38\xa1\x40\xbe\x19\x2c\x21\x23\x94\x12\xd1\x1a\xee\x11\xc4\x90\x8f\x3b\xfe\x60\xc4\x7c\xdf\x92\x21\x48\xd0\xab\x78\xf2\x94\x5b\xd5\xca\xbc\x31\x67\xbc\x23\x06\x90\x3d\xfc\xa2\xec\x52\x93\x0a\xbc\x6b\x34\x67\xfd\x57\xfb\x30\xfd\x0e\x8c\xda\x9c\x77\x09\xbb\xa2\x87\x6a\x9f\x21\x63\x57\x61\x47\x24\xd6\x30\x7e\x28\x1e\x6a\x5b\x8c\x28\x61\xbb\xf3\xdc\xd2\xa6\x75\x07\x3d\x56\x4a\xc8\xa0\x87\xc8\x49\x5c\x31\x87\x7c\x7c\xa7\x70\x8f\x08\x97\x66\x95\x13\x62\xec\x7a\x9a\x39\x0b\x88\x32\xaa\xa1\x1b\xc6\x78\x1f\xa9\xcc\x69\x38\x87\x54\xac\x11\x4c\x14\x43\x06\x6b\x7e\x6f\xde\x9f\x21\x07\x6b\xa4\x32\xbc\x36\x64\x0e\x1d\x2e\x09\x26\x50\x39\x7b\xc8\x14\x61\x48\x7c\x70\x78\xf7\x5e\xb7\x0b\x45\xe1\x21\xc6\xf1\xce\x47\x7a\xea\xf7\x1f\xc6\xa6\x2c\xa4\x92\xbf\xfa\x63\xff\xfd\x75\x86\x52\xb4\xad\xa6\x7a\x3a\x09\x2a\xe7\x12\xd1\xc9\x9b\x5b\x59\xf5\x1b\xa4\xbb\xc5\xfa\x06\x51\x20\x05\xef\x6e\x86\x77\x7e\xbb\xd7\x19\x2a\xba\x57\xc8\x23\x3c\x40\xee\xc2\xd9\xe0\xed\x22\xc2\xc3\x9a\x0b\xab\x3f\x35\xe9\x9b\x8e\xc6\x75\x30\xa4\xf1\x9f\x3c\xf0\x24\x0f\x3b\x29\x3b\x02\x69\x35\x35\x6c\x1b\x28\x88\x99\xaa\x83\xdf\x98\xbc\xa7\x18\xbe\xba\xe7\xa7\x0f\xfd\xec\x85\x21\x32\x58\xcd\x3f\xbb\xe9\xce\xd6\xe9\xbc\xb5\xee\x9b\x6f\x69\xd3\xff\x0a\x42\x6e\xd5\x2d\x9f\x84\x18\xa5\x74\x25\xeb\x85\x3f\xfc\x32\xe4\xbb\x6a\x29\xe4\x6a\xcd\xa9\x84\x8e\x52\xfd\xd0\xc0\x16\x1b\xde\xa1\xc3\xce\x0a\x6a\xfd\x7c\x63\x7e\x7f\x10\x10\xe7\xbb\x2f\xa6\x99\x47\xac\xb5\x8b\x75\x10\xbd\x32\x4e\xa6\x04\x51\xbb\xb3\x10\x5f\xd6\xd4\xc6\x09\x59\x1d\x5d\xb5\x08\xda\x42\x3a\x3c\xde\x18\x23\xbd\xdc\xb4\xf1\xef\x7a\x69\x25\xe6\xeb\x5f\x10\xa1\x60\x26\xd7\xfb\x5f\x61\x0e\xf7\xa4\x87\x5b\xb6\xe4\x84\xb4\x8c\x1b\x76\x23\x1a\xce\xb0\x06\xd8\x1a\x07\xf3\xc7\xe9\x74\x63\xea\x0c\x36\xbb\xa9\xed\xc8\x9f\xaa\x0f\x24\xf7\x63\x02\x3e\x80\x3b\xae\xc5\x60\xcf\x35\x65\x9a\x28\x82\x48\x07\x46\xf9\x81\x5e\x07\x28\xb2\xac\x91\x53\x5a\x2a\x57\xf8\x62\xaf\x6b\x68\xf6\x79\x6d\x7c\xa9\x9b\xbf\x87\x26\xfa\xa0\x93\xa5\xee\xa1\x78\xe9\x52\x53\xf7\xfa\x20\x62\x90\x72\x57\x4d\xa7\xc4\xf2\xf7\xae\xbd\x3f\x7d\x1c\x2d\x78\xaa\xf5\xe3\x46\xa2\x8d\xb2\x7c\x61\x36\xd8\x32\x19\x29\x0c\xef\x8f\xbf\xdf\x52\x9c\x33\xef\x7e\xee\xe6\xe5\xfe\xc0\x98\x87\x82\xdc\xb1\x7c\x69\xeb\xc0\x08\x55\x6f\x3e\x58\xbb\xeb\x37\x0a\xb2\xee\xfa\xc0\x6f\x80\xaf\x20\xf2\x19\x8c\x27\xaa\x08\x3f\x83\x78\x8c\xc2\xe5\x02\x9d\x28\x8a\xda\xe8\xe5\x4b\x15\x15\xb8\xc2\xd4\x6c\x42\x88\x5f\x53\x60\xf9\x7e\xa8\x10\x03\x64\x5d\x55\x84\x86\x91\xf6\xb3\x86\xbc\xa0\xc5\xf0\x88\xf4\x42\xfe\x16\xe5\x0f\x79\xf4\x0d\x6f\xd6\x7f\xb8\xf9\xa8\x27\x48\x01\x86\xa3\xc9\x1c\xad\x37\x6b\xd2\x55\xd9\x01\xf7\xcd\x61\x1a\x49\x1f\x2f\x73\x9f\x71\x95\xb8\xea\x9e\xa1\x23\x06\xaa\x7e\x15\x11\x2f\xbd\x76\x63\x05\x3f\x1b\xe2\xcb\x70\x47\x77\xdf\x87\xec\x8c\x22\xfd\x36\x72\x85\x34\xdb\x94\xa0\x0e\x16\x12\xe4\x83\xf4\x73\x38\x75\xf9\x60\xf0\x17\xf2\x3d\x63\x1e\x41\xce\x7b\x1e\x0f\x97\xe9\x18\x97\x27\x2a\xf4\x5a\x41\x77\xca\x8f\x98\x60\x1f\x41\x63\xe8\xff\xd1\xbf\x6b\xd3\x69\x32\xc7\xaf\x40\x1b\x4a\x5e\xeb\xab\x0a\x7a\xad\x09\x93\x94\x29\x41\x3c\x8e\xa2\x22\xb4\xa6\xc4\x96\xbe\xca\x5f\x81\x19\x0e\xea\x7b\x0b\xb7\x8e\xd3\x9f\xfa\xc3\x41\x50\x6f\x5c\xbd\xc2\x57\xd8\x33\xbe\xef\x4e\xd5\x95\x60\xe9\xe3\x9e\x3e\xdd\x34\xd6\x53\x3d\xc3\x71\xf6\x99\xf9\x68\xb3\x01\xbf\xea\xd1\xa3\x86\x31\x7e\x36\xbc\xd9\x73\x63\xc7\xeb\x86\x77\x0f\xdf\x39\xf4\xbe\x46\x5d\x3f\x4c\x0a\x1c\x27\xd6\x1a\xd6\xaa\x59\x10\xe3\xcf\x85\xb5\x6f\x13\xd6\xc4\xe3\x2f\xab\xc4\xd3\x25\x3d\x5c\xe8\x52\xf8\x6d\xcd\x58\x6b\xfe\x66\xcd\xdf\xad\xf9\x87\x35\xff\xb4\xe6\x5f\xd6\xfc\xdb\x9a\xff\x58\xf3\x5f\x6b\xfe\x67\xcd\x5f\xff\x1f\x00\x00\xff\xff\x27\x29\xa6\x0e\x86\x0c\x00\x00")

func dataIngredient_corpusTxtBytes() ([]byte, error) {
	return bindataRead(
		_dataIngredient_corpusTxt,
		"data/ingredient_corpus.txt",
	)
}

func dataIngredient_corpusTxt() (*asset, error) {
	bytes, err := dataIngredient_corpusTxtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/ingredient_corpus.txt", size: 3206, mode: os.FileMode(438), modTime: time.Unix(1523667476, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dataIngredient_densitiesJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x57\xd9\xae\xe3\x38\x0e\x7d\x6f\xa0\xff\xc1\xb8\xcf\x8d\xc0\x6b\xec\xd4\xdf\xd0\x36\x63\x6b\x22\x4b\x1e\x2d\x37\x48\x0d\xe6\xdf\x1b\xda\x28\xb9\xaa\x5e\x29\x8a\xe2\x72\x78\x48\xfd\xef\xef\xbf\xaa\xea\x6b\xb6\xc6\xa0\xfa\xfa\x51\xb5\xed\x78\xab\xff\x09\x32\x6e\xb1\x5a\x76\x44\x8d\x5f\x3f\xaa\xa6\x1b\xd2\x41\x96\xf5\xfd\x6d\x24\xd9\xba\x82\x2a\xf4\xc7\x07\xe9\x4b\x63\x60\x2b\x6c\xb5\xcd\xfd\xd6\xc5\x33\x85\x70\x14\x27\xdd\x98\x4e\x9e\x68\xa0\x30\x37\xd4\xc9\xdc\x09\xea\x40\x0d\xa2\x38\xac\xe9\x50\x31\xff\x5a\x61\xb1\xa7\x80\xf4\x9b\x69\x5d\x46\x34\x92\xf7\xce\x0b\x27\x9a\xc6\xdb\x14\x44\xef\x9d\x9d\x27\xae\x55\x3a\x1a\xe9\x09\x2d\xad\x22\xb1\x73\xb8\x0d\xf2\x55\xa1\xd6\x4c\x6c\xfe\xd1\xe1\xd6\x07\xe9\xc1\xf8\xcb\xe7\xb5\x4e\x81\x7d\xe3\x86\x06\x66\x8e\x95\x64\xdc\x67\xa3\xb9\x0d\xff\x14\x65\x48\x57\xfa\x22\xe3\x72\x91\x1c\x4c\x48\x6c\x93\x5e\xfc\xc8\xcd\x2a\x13\x4d\x44\xbf\xf5\x8b\x1d\xd5\xef\x16\xbe\x41\x30\xce\xc1\xdf\x1f\x6e\x8f\x20\xc4\x6f\x38\xa5\x02\x83\x2b\xdd\x18\x5a\x2a\x3f\x68\xef\x5e\xf3\x68\x92\x68\x65\x9c\x57\x6f\xc4\xf5\xeb\x47\x35\x25\x23\x87\xd5\x06\xd4\xea\x6b\xd4\x51\x8e\x80\x7b\xbf\x1e\xd1\xd3\x85\xad\xa8\xaa\x6f\x26\x70\x03\x0f\xb3\x8e\xd0\x51\x0a\x27\x2a\x23\xae\xd5\x9b\x09\xfc\xf3\x95\x19\xb8\x86\x83\x2d\xe5\xe9\x40\xa1\xce\x88\x4f\x9f\x93\x36\xe5\x64\x86\x45\x0a\xef\xe0\x3d\x89\xf6\x50\xbd\xa1\x4e\x75\x5a\x76\xb6\xbc\xd0\x6b\x3d\xfa\x64\xca\x58\xf5\xc2\x8f\x53\xac\xef\xb7\x7b\x90\x6d\x0a\x51\x54\x33\x82\xd0\xb1\xae\x23\xbd\x6b\xb4\xbf\x3f\xa4\x57\x16\x50\x4a\x06\xe1\x38\x25\x03\xfa\x8d\x68\xaa\x53\x1a\x30\x12\x83\x0d\x7a\xb0\x90\x36\x53\xf6\x4d\x2a\xef\xd8\xbd\xbf\x35\x51\x0f\x41\x47\x2c\x44\x89\x3e\x99\x80\x65\xbf\xaa\x49\x30\x07\x82\xab\xe2\xbd\x4e\x18\x3b\xe0\x23\x85\x00\x16\xfb\xad\x26\xb7\x76\xa9\x0c\x8a\x08\xe0\x9a\xfa\x50\x03\x87\x35\x8a\x72\x93\xb9\xcb\xd3\x23\x01\x31\x02\x39\x5f\x3a\x11\x84\x35\x04\xf0\x7b\xa6\x82\x45\x96\x07\x54\x70\xc9\xd9\x37\xfe\xae\xaf\x51\xc3\x81\xbf\xab\xc3\xb7\x5c\x60\x95\x57\xe1\x02\x42\x72\x88\xca\x2e\x7d\x03\x45\x70\x84\xfa\x8f\x14\xc2\x01\x6a\x03\xc5\x84\x4f\x42\x3b\xa4\x1a\xe6\xde\xfc\x25\x1d\xc3\x2f\x11\x38\xe1\xbd\x49\x3d\xe0\x73\x44\x5e\x52\x35\x41\x9f\xa0\x60\xb3\xbe\xc8\x75\x4f\xc9\x91\xea\xe5\xd9\xf3\x5e\x40\xe7\x59\xcd\x4a\x9a\x3d\xba\x93\xda\x06\x39\xaa\x8f\x8f\x86\xc0\x1c\x81\x9a\xd5\xeb\x2e\xa9\xcf\x4a\x2e\x8b\xe4\xcc\x43\xbd\x4f\x85\x4d\x17\xb4\xb4\xa7\x27\x85\xe9\x4a\xbe\xf2\x59\x1d\x56\xef\x4a\xca\x23\xeb\x3c\x52\x68\xde\x35\x6d\xe4\x92\x08\x25\xde\x95\x82\xc5\x9e\x6a\xd3\x43\x01\xbc\x91\x32\xc6\x84\x2a\xc5\x0e\x6f\xb2\xae\x93\x9b\x46\x1e\x41\xcf\xa1\x8f\x6a\x64\x3d\xa6\xda\xa1\x49\xa8\x9f\x41\xcd\xb8\x58\xac\xe8\x6c\x7c\x5c\x4d\x90\xbb\xf5\x9d\x32\x96\x82\x25\x8f\x89\xb7\xdf\x10\x47\x5c\x4f\x09\x0b\xbd\x7c\xe2\x79\xfa\x93\xa6\x2b\x86\x8c\x36\x9e\x2d\x33\x16\x37\x50\x9c\x2d\x71\xec\x51\xe1\xad\x86\x2d\xcc\x42\x22\xf8\x34\xa0\x9c\xb4\xa5\x1a\x3b\x4a\xcb\x2f\xb5\x05\x5d\x05\x26\x69\xa6\x07\xb5\x68\x2c\x87\x8e\x0e\x4c\xbf\x82\xb3\x84\x4a\x47\x81\x73\x96\x53\x35\x10\x59\xbe\xa5\x5a\x50\x1b\x54\x7a\x67\xaa\x4c\x26\x79\xe0\x4b\x19\xe8\xe4\x9e\x5c\x80\xf3\xe4\x69\xc4\xdf\x2f\x3d\x4d\xab\x82\x83\x58\x9f\x9b\xc2\x65\x2b\x17\x6f\x97\xc2\x13\xe7\xf0\x28\xb8\xc8\xbd\x31\xd1\x34\x59\x98\x10\x10\x1a\x73\xe8\xd2\xbb\x0a\x98\x66\xc1\x99\x9e\x00\xf3\x54\x96\x99\x18\x6e\x1a\x73\x36\xf2\x7e\x4d\xa2\x03\x4e\xd7\xb7\x1f\xe5\x31\xd1\xb5\xb9\x29\xde\xa2\x4a\xfa\xcd\xbd\x4e\xb5\x3b\xed\x71\xbe\x98\xaf\x52\x47\x8c\x39\x83\x00\xe1\x0b\x9f\x3d\x3f\x40\xe9\xfd\x00\xce\xe5\xdb\x39\xd6\x3d\x52\xce\x3d\xaf\xac\xa1\x4a\x43\xd2\xf6\x99\xd3\xb1\x33\xda\x42\x56\xfd\xc7\x06\xde\x74\x00\x2c\xd3\x4c\x25\xe9\x69\x97\x08\x6e\x04\x2b\xb9\x5d\x77\x54\x8a\x85\x91\x51\x4f\xa4\xca\x2d\xce\xc5\x01\x95\x70\x51\x20\xf2\x41\xd3\x76\xc9\xce\x0a\x26\x88\xf2\x6a\x74\x32\x81\xa9\xe4\x6d\x47\xa0\xe2\x78\x04\x64\xb8\x61\x5a\x17\xb2\x22\x16\x62\x3b\xce\x0e\xfc\x93\x5c\x2a\x10\x1b\x26\x7c\x5d\x84\x59\x7d\xa0\xa5\xe6\x44\x58\xf6\x38\x17\x89\x85\x4e\x04\x15\x66\x2b\xcd\x2b\x72\xb9\x30\x92\xd7\x40\xd0\x67\x11\xfb\xf8\x20\xdc\x9c\x0a\x3e\xd7\xb5\xf0\x00\xb1\xc9\xe8\x43\x14\xbd\xd8\x2a\xf0\x43\x43\xde\x6d\x04\x89\x4c\x61\x9e\x63\xc3\x37\x6d\x2a\xf9\x02\x96\xb3\x27\x97\xef\xd0\x17\x3d\xd5\x7d\x91\x8a\x81\x58\x83\x38\x13\x49\xc1\x3a\x3a\xee\x3d\xa9\x66\x76\xb1\xc7\x1c\xcd\xd0\x92\x8c\xdb\x76\x72\x10\xc6\xaf\xba\x94\xd7\x8d\x89\xcd\x6b\x3e\xc8\xf2\x0b\x42\xd3\x36\x63\xae\x16\xbe\xdc\x1b\xb9\xdb\x39\x1a\x13\xf0\xd6\x13\x9d\xea\x05\x38\x4f\x24\x90\x97\xe9\xe0\x28\xd1\xc3\xd8\x14\x2b\xb8\xe6\xbe\xbf\x3b\x9a\x4b\x3f\xed\xb2\xec\x4c\xb0\x98\xaf\x3e\xe5\x1b\xb6\x1d\x8d\x09\x33\x89\xe8\x21\xf0\x77\x2c\xce\x98\x92\x1b\x59\x3d\x63\xa8\xb9\x12\xb5\xe7\xb8\x00\x8d\x3e\xd3\xc8\xce\xbe\xbd\x8c\x58\x5d\xef\xae\x57\xc3\xda\x95\x29\xca\xaf\x62\xde\x8b\x3c\xcb\xc1\xe8\x38\xef\x28\xe6\x27\x0a\x81\x6e\x8e\x4f\x94\xc3\x1d\x8e\xd9\xaa\x2d\x0e\x8f\x3c\xdd\x95\xdd\xac\x5f\xa9\xdb\xba\x70\xc6\x57\xa0\x1b\x2f\xed\xe2\x73\x72\xcf\x60\xa6\x29\x40\xc6\x9e\x5c\x5a\x2f\x6a\x3a\x1a\x9d\x61\xf7\xd1\x88\x81\x62\xf2\xcb\xc2\x86\xe0\x06\x6a\x86\xb4\x57\xa5\x35\xbe\x23\x16\xd8\xe1\x27\xf2\x74\xa1\x9e\xae\x0b\x9a\x8e\x1f\xbb\x29\x09\x97\x08\xf8\xfa\xc2\x0c\x55\xa1\x4a\xc3\x34\x19\x7d\x90\x2c\xd0\x61\xf9\xf9\x48\xa8\x96\xcf\x27\x62\xdc\x3d\xf2\xca\xa3\xe2\xff\xe9\x42\x09\x51\x36\x5e\xe7\x7c\x62\xc8\x66\x4c\x3d\x67\x6c\xa0\xe9\xfe\x9e\xff\xab\xb0\xbc\x2e\x3d\x3b\x14\x3b\x41\xda\x98\x1f\x44\x83\x5a\x7e\xb2\xe1\x4c\xc9\x46\x3e\xed\x95\x4c\xde\x3b\x33\x6e\xe6\xa2\xdf\x83\x33\x28\xbd\xa4\x5a\x94\x3d\xe6\x30\x03\xa8\x40\x49\xb7\xe9\x29\xb7\x9b\x82\xdd\xfd\x74\x15\x2c\xaf\xd0\xf6\x53\x9f\x7f\x94\xab\x9f\x9e\x63\xf2\xe1\x64\xe8\xec\x6a\x8f\xd7\x96\x46\xf9\x86\x1c\x4c\x98\x58\x39\xeb\xa1\x57\xdc\xcf\xb5\xcb\x3c\xec\x55\x68\x31\x3e\x24\x07\xad\x43\x93\xe4\x7c\xbb\x5f\x45\x45\xb8\x6b\xbb\xf2\xb7\x11\x7f\x0d\xc5\x07\xff\x83\x6e\xfe\x55\x97\x43\x4a\xa5\x93\xba\x5f\x60\xf8\x7f\xb4\xf4\xf0\x7f\x2d\x13\x12\x62\x93\x0f\xe5\x3c\x8e\x5f\x88\x66\xec\x53\x36\x43\x96\x49\xde\xd2\x8a\x85\x87\xe4\x2c\x14\xbb\x26\xaa\xc0\x6d\xab\x84\x94\x6b\x1c\xb5\x0d\x99\xd1\xee\x6f\x98\xea\xf7\xf7\x5f\xff\xff\x37\x00\x00\xff\xff\x22\xe3\x9b\xfa\xe0\x10\x00\x00")

func dataIngredient_densitiesJsonBytes() ([]byte, error) {
	return bindataRead(
		_dataIngredient_densitiesJson,
		"data/ingredient_densities.json",
	)
}

func dataIngredient_densitiesJson() (*asset, error) {
	bytes, err := dataIngredient_densitiesJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/ingredient_densities.json", size: 4320, mode: os.FileMode(438), modTime: time.Unix(1523134194, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dataTop_5kTxt = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\x57\x5b\x96\xe3\x2a\x0c\xfc\xe7\x1c\xf6\xc0\xd6\x64\xac\xd8\x8c\x79\x5d\x09\x92\x49\xaf\xfe\x1e\xf1\xb2\x7b\xbe\x3a\x4e\x30\x88\x52\xa9\xaa\x9a\xc1\x17\xad\xb8\x1e\x40\x5a\x6d\xb5\x14\x24\xad\x0e\x20\xef\xac\x56\x1f\x68\x8f\xc9\xbb\x37\x9a\xe4\xbc\x56\xc1\xf9\x4b\xab\x97\x4f\x55\xbe\x8f\x2e\x45\xad\x32\xe6\x8c\xf3\x91\xb5\xda\x28\x7d\xa2\x19\x5b\xe2\x71\xb0\x56\xd6\xc5\x08\x41\x16\x6f\x70\xb9\x78\x98\x9c\x3e\xbb\xbc\xe3\x31\xa4\x68\xfe\x54\x67\x51\xab\x00\xe5\x27\x99\x80\xe0\xb5\x2a\x29\x40\x49\xc8\x5a\xbd\x21\x3a\xef\x41\xab\x0c\xc4\x1e\xbf\x6b\x0f\x4e\x3b\x68\xc5\xa9\x92\xb1\x84\x10\xb4\x7a\xe3\x81\x05\x36\x3f\x8a\xb5\xe8\x91\xbe\x5a\x1d\x2e\x1e\xeb\x30\xad\xda\x62\x63\x4f\x44\x46\xad\x2c\x10\xa5\x22\x35\x9e\xb8\xef\x40\xeb\x87\x0d\xf1\xa5\x55\x4e\x65\xd4\xd1\xb6\x3c\x53\x94\x0a\x62\x2d\x01\x8f\xf6\x4e\x5b\xcb\xe9\x6b\x18\x6a\xbf\xc4\x37\xc5\x08\xae\xed\x7d\x3a\x7b\x61\x34\x1b\xa5\x72\x6a\x95\x08\x0f\x88\x49\x2b\x5b\x83\x8b\x5a\x95\xf3\x1b\x70\xa2\xbd\x20\x09\x95\x4f\x4a\x29\x34\xd8\x3c\xc4\x42\x49\xae\xcc\x72\x7c\x46\x0b\x0d\x62\xb0\x72\x93\x13\xe1\xfd\x9d\x97\xbf\x0f\x43\x60\xb9\xd0\x27\x91\x45\x2e\x48\x7c\x3a\xc2\x59\x5f\x86\x4c\xee\x02\x59\x9f\x6c\xf2\x50\xee\x3a\xa5\xb1\xf0\x57\xda\xee\x63\x6d\x3b\xb4\x0f\x66\xb4\x7b\x07\xba\xcc\xef\xb7\xbc\x5b\x65\x83\x0f\x29\xee\xac\x95\x77\x01\x67\x47\x33\x50\x40\x86\xb8\x40\xcd\x2e\x22\xe4\xec\x51\x2b\x6a\x2b\x12\x41\x3c\xd6\xfa\x83\x10\xa3\x99\x84\x22\x70\xec\xe4\xba\x36\xd9\x14\x6b\x91\x66\x7d\x31\x46\x5c\x2b\x7a\x99\xbb\xfb\x93\xa2\x09\x95\x0b\xd0\x2e\xab\x29\xca\x47\x7b\x6a\x15\xd2\xcf\x0f\x10\x7a\x0f\x77\x5f\x1b\xc9\x3b\x91\xdf\x2e\x62\xa3\x69\xab\x89\xb5\x22\xdc\xd7\xe6\x9d\x82\x13\x37\x81\x75\x37\x96\x6a\xd8\x58\x2b\x2e\x08\x97\x56\x09\xe4\x7c\xce\x2e\x82\x9c\xc6\x67\xa2\x82\xd1\xc5\xe3\xb9\x93\x79\x79\xb8\x64\x73\x3e\xc1\xfb\xc6\xb5\xb1\x75\x06\x16\x20\x65\xe9\x86\xde\xaf\x93\x05\x42\x59\x4e\x2e\x64\xad\x18\x43\xf2\x2e\x82\x56\x3f\xd5\xda\xd3\x45\x27\xe7\x13\x7c\x36\x24\x72\xad\xea\xc4\x18\x40\xa8\x6e\x21\x26\x0f\x9d\xfe\x1d\xcc\x35\x95\x10\x21\x82\x54\x61\xc1\xfb\xfe\x9d\xf5\xe9\x2d\xef\xdf\xd0\x0d\x0e\x71\x49\xf6\x6a\x8f\xed\xf7\xcf\xe9\x72\x96\x91\x1b\x54\x0b\x90\x3d\x1a\xfe\x52\xcd\xb3\x83\x1d\xf6\x49\x95\xf6\x79\xf6\x60\x03\xcf\x10\x9c\x35\x0b\xed\x9d\xbe\xe6\x73\xba\x82\xe6\xe3\x62\x7f\xd5\x41\x6c\x34\xda\xe0\x6b\x3c\xc2\x4b\xab\x0b\x8b\x3d\xe5\x80\x6f\x3a\x2a\x95\x0e\x93\xac\xbf\xf7\x81\x77\xb2\xb0\x27\x81\x88\x21\xcc\xa9\x87\x6d\x03\x29\x68\x6c\xd5\x6e\xb0\x51\xb2\x36\x79\xf7\xef\x98\x74\x6e\xc1\x94\x81\x31\x68\xeb\xb5\x06\xde\xa2\xb8\xad\xb6\x86\x4d\x3e\x65\x04\x19\x8c\xa9\x97\xe0\x3d\xe7\x46\x60\xb9\xd9\x8d\x26\x41\x5c\x3d\x0a\x2e\x16\xad\x4e\x81\xaf\x37\xe6\x57\xc7\x37\x5f\x71\x2d\xed\xc2\x96\x51\xba\x65\x2b\xd1\xf7\xae\x20\x51\x5c\xe3\x60\x3a\x87\x3d\x96\xd2\x18\xda\x11\x9d\x3b\x0e\x48\x18\x51\x86\x32\x27\xba\xb4\x2a\x95\x02\x92\x28\xbb\x30\x0f\x64\x72\xbc\xd7\xea\x8b\xde\xa7\x8f\x19\x72\xfe\x6c\x8c\x40\x3e\xbe\xfe\x03\x1e\x32\xc6\x64\xda\xd8\x4b\x99\x5d\x47\xd3\xcb\x4c\xc5\x12\xcd\x6c\xf2\xd4\xb5\x60\x92\x61\x3c\xe1\xdf\x42\x60\x05\x62\x0f\xf6\x32\x63\xe9\x10\xbf\x6e\x43\x1d\xd6\x26\x7c\x3b\xd2\xdd\xe7\x5e\xd2\x7a\x0c\x40\x07\x50\x2b\x70\x42\xd9\xf6\x9a\xec\x4c\x1e\x98\xa5\xc2\x5c\x43\xbe\xdc\x03\xb1\x46\x90\x31\xd0\xa5\xd2\x25\x5a\xfe\xed\x3c\x68\x26\xc7\x93\xba\x9d\xda\xdc\x78\x24\x42\x75\x1f\x4e\xc0\x79\x75\x4a\x3c\x62\xaa\x3b\x39\x9b\x4a\xb9\x75\x86\x85\xf6\x73\xd4\x45\x24\x08\xbe\xb7\xa4\x2f\x8a\x0d\x57\x5b\xcf\x4d\x89\xcc\x3f\x00\x58\xc8\x48\xed\xbc\x07\x67\x0e\xf4\x50\xe4\x76\x1d\x84\xd5\x97\x8e\xef\xbc\xcf\x0b\x9f\x35\x7d\x1c\xf3\x6d\x7d\x27\x0a\xb9\x6e\xab\xed\x66\x37\xff\xae\x01\xee\x13\xde\xb0\x49\x5e\x06\x37\x3f\x9a\x5f\x80\x0a\x0c\x6a\x76\xeb\x5e\x1a\x3f\x55\xdd\x15\xf0\x0e\xa2\x61\x04\x4e\x5d\x1d\xdb\xd1\x9d\x45\x50\xbd\x7b\xf9\xf4\x41\x7a\x72\xef\x57\xf3\x73\xc6\x7d\x6a\x4f\x57\xea\xde\x42\xe0\x0c\x04\x47\xe5\xe1\xa5\x0b\x44\x6e\x9c\x6b\x05\xb7\xa8\x30\x84\xa1\x99\xc7\x88\x2f\xa6\x24\x2a\x12\x2c\xc4\x40\xa0\xc8\x5b\x1e\xf1\x6a\xac\x21\x29\x7e\xf0\xfe\xa1\xa2\x7d\x88\xd7\x21\x3b\x21\x73\xbb\x0d\x5f\x2e\x8c\x69\x0c\x10\x8f\x24\x6f\xc1\x29\x21\x83\xc0\x5e\x48\xcb\x36\x5e\x8e\xcf\x65\xc4\x5d\x43\x5a\x69\x53\xda\xa4\xe2\x52\xe0\xc0\x27\x8b\x60\xef\x8b\xba\xdb\xff\xab\xc5\xa5\x8a\x33\x0c\x4f\x65\xa8\x9d\xb3\x23\x34\x19\xa1\xee\xcc\x07\x03\x8b\x41\x44\x3c\x8e\x2c\x54\xd4\x2a\xfb\x1a\x1e\x24\x28\x40\x04\x47\x6a\xf9\x84\x6a\xab\xd3\x9e\xf2\xc3\x91\xa0\xac\xba\x5a\x63\x2f\xb7\x47\xfc\xce\xf1\x2b\xe9\x55\xc7\xfc\x3c\xb0\x6d\x7a\xdb\xb5\xec\x8e\x67\x63\x64\x7a\x40\x6b\x1a\xb5\xdc\x71\xb8\xd2\xc3\x8b\x52\x1b\xc5\x0c\xc7\x89\xa5\xb8\x99\x19\x33\xa2\x9f\x5c\xef\x13\xd1\x7c\x7f\xc6\xd1\x2e\xb7\xdd\xf4\x9a\xec\x6e\x95\x5a\x08\xb4\x40\x3b\x04\x11\x2b\x0b\x85\x65\xeb\x8d\x20\xee\xdf\x01\xb5\x44\x3c\x62\x24\xd8\x1d\x9f\x5a\xbd\xd3\x2e\x59\x89\x3f\x88\xc5\xdc\x59\x70\xc3\xae\xc6\xaf\x17\xae\x44\x21\xcd\xe4\xff\x2a\xc8\x6b\x9d\xc7\x83\x34\x1c\xd2\x25\x61\x60\x06\xaf\x5f\xb3\x21\x53\xdb\x72\x11\x1a\x3e\x51\xd4\x78\x88\x96\x69\x5f\x75\x57\x69\x37\xfe\x05\xe0\x30\x68\x3c\x0e\x13\x53\xda\xfd\xd3\xe3\x44\x77\x09\x45\x46\x65\x12\x4d\x57\xd4\x7e\x83\x95\xd8\x81\x56\x50\xa7\x99\xd4\x9a\x0c\xcd\xac\x72\xc7\x17\xd6\xea\xbf\xea\xa2\x98\xe4\x99\xca\xcc\x34\x53\x44\x47\x46\x6a\x3d\x12\xf3\xe3\xa6\x29\xdd\xd0\x16\x59\x80\xea\x51\x25\xc2\xcb\x8a\x27\xd7\x29\x9a\x0b\x29\xa2\xe7\xbe\xf9\xd8\xf5\x19\x03\x79\xc4\xe5\xc1\x92\x0d\x68\x43\x5b\x57\x9c\x05\x2a\xce\x9e\xe9\x42\x73\x22\x50\x0f\xad\x05\x5b\x8a\xe7\xd2\xaf\xde\xf9\x42\x2e\xee\x33\xb3\xce\xe2\x61\x03\xb6\x77\xc2\x13\x31\x1f\xc0\x8e\x58\xda\x49\x16\x80\xf8\x0c\xc2\xce\x0f\x6b\x75\x81\x5f\x71\x70\x38\xdc\xa3\x2f\x8f\x04\x08\xcd\x6b\x1f\xe1\x57\xdc\xd6\x7c\xda\x35\x5e\x54\x5d\x99\x3c\xe9\x1d\x38\xe1\x07\x47\xee\xb6\x04\x5b\x40\x28\x9d\x17\x96\xaa\x18\xd4\x86\xad\xa7\xc3\x4d\xbb\xda\xcc\x87\xa6\xbb\x32\x9e\xa3\xa0\x00\xf4\x27\x91\x4c\xfd\x06\xdb\xd7\xac\x80\xfa\x5b\x96\x84\x21\x94\xd8\xba\x5a\x44\xf5\x9b\x56\xd4\xad\x9b\x46\x8c\xb8\xfe\x0b\x9b\x5c\xc5\x37\xe4\x44\x50\x70\x1e\x3f\xe3\xfd\x8a\x10\xb7\x09\x2c\xad\x1f\xd6\x7e\x00\x41\x30\x01\x64\xc8\xfe\x0f\x00\x00\xff\xff\x1b\x8d\x4f\x9b\x74\x0e\x00\x00")

func dataTop_5kTxtBytes() ([]byte, error) {
	return bindataRead(
		_dataTop_5kTxt,
		"data/top_5k.txt",
	)
}

func dataTop_5kTxt() (*asset, error) {
	bytes, err := dataTop_5kTxtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/top_5k.txt", size: 3700, mode: os.FileMode(438), modTime: time.Unix(1523669317, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dataVegetablesJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\x55\xc1\x6e\x2c\x29\x0c\xbc\x3f\xe9\xfd\x03\xca\x79\xf7\x8b\x56\x7b\x30\xb4\x07\x9c\x06\xc3\xda\x90\x51\xe7\xeb\x57\x9d\x3c\xcc\x10\x69\x0e\xae\x82\xc6\x45\x51\x30\xff\xfc\xfe\xe5\xdc\x1b\x84\x2a\xec\xf4\xbf\x01\x9a\xde\xfe\xfa\xa6\xf2\xe3\xfe\x39\x6d\x52\x47\x9f\x64\x01\x01\xee\x36\x87\x49\x71\xd6\xd2\x29\xa4\x7a\x2e\x3c\xe2\xc8\x30\x91\x36\x10\x88\x43\x27\x1e\x1e\x25\x12\xdb\xec\xcf\x71\x92\xf3\x08\xfc\x87\xf0\xc0\xc0\xb0\x4b\xf2\xa0\x94\x67\x8d\xc0\xbb\x36\x8f\x68\x65\x86\x70\x6e\xab\xdd\xc4\xdf\x78\xe1\xe1\x1a\x4e\x4d\xbe\x9e\x2e\xa4\x7a\x19\x94\x5c\x7b\xdf\x55\x48\x85\xe3\x8b\xd0\xc5\x84\x50\x1f\xb9\x3e\x51\x36\x2a\x93\xc1\xa1\x8a\x59\x7f\xa8\x1b\xbd\xa3\xf0\xe8\xfb\x96\x02\x78\x0f\x11\x0d\x65\xf0\x82\xba\xb0\xc0\x13\xae\x85\xa4\x76\x03\x23\xd3\x26\x23\xc0\x85\xcc\xe8\x1a\xb6\xb6\x48\xcc\x28\x04\xe1\x15\xda\x7a\x09\x4a\x2d\x94\x71\x61\x39\x56\x7d\xd5\xbe\x46\x28\x9c\xcb\xb8\x90\xe8\x03\xa7\x21\x81\x32\x70\x97\x3a\x61\xcd\x19\xe4\x70\x51\x10\xd9\x38\x79\x2d\x9d\x42\x06\x6b\x54\x87\x44\xec\xab\xd5\x08\xa3\x78\xd3\x7f\x00\x9d\x75\x7e\x7c\x60\xa6\x00\x7d\xca\x38\x28\xcf\x34\x60\x8c\xed\x96\x31\x21\x1f\xf4\x31\x57\x7c\xdc\xae\xcc\x89\x0f\x3a\x8e\x8c\x09\xad\xff\x43\x48\x71\x4e\x8d\x20\x99\xa6\x57\x11\xcb\x7e\x56\x91\x38\x9a\xb0\xaf\xfd\xbd\x46\xe5\x9b\xd8\xcc\x4f\xe0\x81\xd1\xac\x49\x28\x5e\x1d\xf0\xe1\xb4\x51\x98\x3d\x53\x15\x45\x81\x83\xac\x4f\x1a\xde\xdf\x16\x6e\xcd\xdf\x21\x43\x43\x9e\x6b\xbd\xa3\x0c\x85\x8c\xc5\xfd\xbc\x77\xef\x14\xa0\x4c\x8b\x4e\xb0\xd3\x3d\xe9\x60\xbc\x5e\x15\x9f\x35\x65\x01\x3f\x73\x9b\xe1\x03\xf9\x30\xf1\x19\xf1\x74\x56\xc7\x51\xd0\x40\xa9\xec\xa2\x80\xaa\x31\xdc\x29\x2f\xd4\xfb\xb0\xed\x65\x2a\xf0\xda\xb4\x40\xc1\xcb\x6a\x8e\xd8\xd7\x1d\x29\x20\xef\x55\xa0\x4c\x38\x38\x6e\x9f\x0e\x4d\x52\xeb\x1a\xd6\xfe\x33\x69\x0c\x1f\xdb\x16\x19\x9f\xee\x13\x21\xff\x31\x9d\x21\x4c\x3f\xb9\xb6\xe5\x4d\x3d\x65\x3a\x56\x99\x2c\x6f\x55\x30\x82\x59\xde\xa0\x09\x9d\x60\x48\x34\xdb\x4e\x6e\xc4\xd4\x0c\xf5\x7e\xb9\x66\x2a\xd6\xcd\x69\xc4\xbd\xbe\x0a\x6c\xb5\x43\xb7\x0e\xa3\xb4\x93\xe6\xd0\x9d\x89\x10\x12\xd5\x17\x6c\x71\x90\x34\x3c\x88\x9f\xa8\x2a\x16\xb0\x8b\x2d\x83\x19\xe5\xb5\x8d\x8c\x0e\x1e\xe2\x94\xa1\xeb\xc9\xd1\x00\x39\xaf\x2d\x6b\x82\x9c\xed\x91\xd1\x93\x44\xec\x55\x55\x86\xf6\xf2\x7c\x6a\xdd\xac\xd6\x06\x31\xe1\xfd\x7e\x6e\xb9\xdd\x5d\xff\x1e\x9a\xb1\xd2\x27\x62\x77\x9b\x05\xb7\x4c\x0d\x75\xbf\x47\x1d\x64\x8d\x77\xa7\x75\x46\xb6\xa7\xcb\x62\xd9\xeb\xdd\xa9\xf8\x61\x5f\xdd\x7f\x31\x6a\x40\xd6\xf9\x3c\x41\x57\xec\x9f\xd0\x51\x5c\x48\xa8\x9d\x2d\x89\x5f\x64\x10\xb4\x88\x3f\x13\x75\x74\xdb\x11\x5c\x16\xd4\xcf\x71\x9f\x13\xd3\xdb\xef\x5f\xff\xfe\x1f\x00\x00\xff\xff\x1d\x08\x3f\x45\x50\x07\x00\x00")

func dataVegetablesJsonBytes() ([]byte, error) {
	return bindataRead(
		_dataVegetablesJson,
		"data/vegetables.json",
	)
}

func dataVegetablesJson() (*asset, error) {
	bytes, err := dataVegetablesJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/vegetables.json", size: 1872, mode: os.FileMode(438), modTime: time.Unix(1523666078, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"data/fruits.json": dataFruitsJson,
	"data/herbs.json": dataHerbsJson,
	"data/ingredient_corpus.txt": dataIngredient_corpusTxt,
	"data/ingredient_densities.json": dataIngredient_densitiesJson,
	"data/top_5k.txt": dataTop_5kTxt,
	"data/vegetables.json": dataVegetablesJson,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"data": &bintree{nil, map[string]*bintree{
		"fruits.json": &bintree{dataFruitsJson, map[string]*bintree{}},
		"herbs.json": &bintree{dataHerbsJson, map[string]*bintree{}},
		"ingredient_corpus.txt": &bintree{dataIngredient_corpusTxt, map[string]*bintree{}},
		"ingredient_densities.json": &bintree{dataIngredient_densitiesJson, map[string]*bintree{}},
		"top_5k.txt": &bintree{dataTop_5kTxt, map[string]*bintree{}},
		"vegetables.json": &bintree{dataVegetablesJson, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

