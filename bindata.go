// Code generated for package main by go-bindata DO NOT EDIT. (@generated)
// sources:
// assets/tpl/curd.tpl
// assets/tpl/e.tpl
// assets/tpl/entity.tpl
// assets/tpl/example.tpl
// assets/tpl/init.tpl
// assets/tpl/markdown.tpl
// assets/tpl/models.tpl
// assets/tpl/gorm.tpl
// assets/tpl/tables.tpl
package main

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

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _assetsTplCurdTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x59\x6f\x6b\xdb\xda\x19\x7f\x2d\x81\xbe\xc3\x33\x73\x09\x52\x63\x94\xec\xde\xb1\x17\x01\xf7\x92\xd8\xce\xad\x59\xe2\xe4\xda\xca\xc6\x08\x21\xc8\xd2\x71\xa2\x45\x7f\x9c\xa3\xe3\xc4\x41\x08\x7a\x61\xe3\xb6\x70\x7b\x77\x5f\xac\x69\xe9\xca\xb6\x96\x76\x94\x41\x9b\xbe\xd8\x58\x5f\x0c\xf6\x65\x62\x37\x1f\x63\x9c\x23\xc9\x92\x6d\xd9\x96\xf3\xa7\x94\x91\x42\x69\x7d\x74\xf4\xfc\xfd\x3d\xbf\xe7\x9c\x47\xe4\xa4\x85\xc0\xf3\xe4\x3a\xc1\x6d\x8d\x28\x6a\xc3\x44\x55\xd5\x42\xbe\xbf\xee\xe8\xc8\x04\x97\x2d\x83\x27\xf0\x65\x81\x2f\xad\xc0\x1d\xf7\xd0\x94\x4b\x2b\x02\xaf\x74\x82\xff\x2b\x1d\x81\xf7\x05\x5e\xe0\x17\x16\xc0\x76\x08\x10\xac\xda\xae\xaa\x11\xc3\xb1\x05\xbe\xd9\xb6\x35\xa8\xa2\xe3\x34\xf9\xa2\xde\x00\x59\x96\x43\x81\x12\xdc\x19\x6f\x84\x27\xf0\x46\x13\x4c\x64\x8b\x7a\x43\x82\xbb\xb0\x48\x57\x00\x00\x30\x22\x6d\x6c\xc3\xdc\xd8\x57\xc3\x7d\xf4\x4f\x69\x65\x09\xf4\xc6\xf6\xe2\x4e\x3e\x58\xf3\x99\xdd\x59\x25\xd0\xb7\x2d\xd5\x25\x08\x97\x56\xf2\xec\x4d\x9f\xb9\x9c\x70\x17\x9c\xc6\xef\x90\x46\x26\x7b\xad\x74\x44\xd2\x8f\xdc\x34\xa7\xb9\x2c\xd6\x71\x9c\xd2\x59\x02\xd2\xc9\x0b\x3c\xe7\x87\xb9\x60\xa6\x5d\xfc\xf8\xef\xee\x1f\x4f\x7b\x0f\xef\xf7\x9e\x3f\xfc\xf8\xec\xf7\x17\x2f\xde\x74\xdf\x3e\xe9\xbd\xfb\x57\x68\xa0\x68\x4d\xd0\x2e\xc1\x1e\x22\x45\xc7\x6c\x5b\xb6\x2b\x4a\x14\x05\x86\xbd\x47\x4d\x0a\x2d\xca\x51\xc8\x2c\x9b\xe6\xaa\x81\x4c\x7d\xcd\x70\x89\xef\x43\x2e\x06\x42\xa0\xba\xfb\xea\xd9\xc5\x8b\x1f\x7a\x8f\xdf\xf7\x1e\xbd\x93\x33\x6b\xad\x39\xc7\xae\xe8\x1e\x9a\x4a\x87\x84\x7a\xf3\xd0\x52\xb1\x6a\xb9\x14\x2e\x86\x4d\x10\x6e\xaa\x1a\xf2\x7c\x09\x44\xec\x1c\xbb\x35\xe4\xb6\x4d\x02\xdb\x3b\x54\xee\xe6\xc1\x5e\xd9\x26\x06\x39\xf1\xfd\x34\x2d\x79\x40\x18\xd3\xbf\x0e\x96\xa8\x37\x87\x6d\x84\x4f\x82\xc5\xa5\x02\x58\x72\x69\x45\xfe\x96\x2e\x85\xfa\x23\xc5\xb2\x2c\x4b\x0c\x82\x74\xe3\xcf\x0a\x60\x1b\x2c\x3d\xf4\x17\x7d\xab\x2c\xd7\x89\xaa\x1d\x88\x08\x63\x29\x8a\x10\x8b\x85\x8e\x9a\x08\x03\x53\x22\x17\x4d\xc7\x45\xa2\x24\xf0\x4d\x27\x5a\xaa\xa2\x0e\x11\x99\x1d\xd8\x39\xa6\x06\x8c\x38\x50\x6d\x9b\xe6\x88\x13\x9e\x1f\xa9\x0e\xc4\xd4\x35\xd5\x16\x05\xde\xf3\xb0\x6a\xef\x21\x60\x2f\xb1\xbc\xb8\x15\xbb\xe9\xf8\xfe\x1c\x76\x8e\x65\xcf\x93\xef\xb5\xad\x56\x18\x86\x85\x05\xaa\xab\xe8\x58\x16\xb2\x89\xef\xd3\x97\x91\xad\xfb\x7e\x9a\x97\x14\xfd\xe9\x9e\xc6\x15\x18\x94\x52\x9c\x8c\x02\xa8\xad\x16\xb2\xf5\x44\x82\xf2\x0c\xc5\x53\xf2\xe3\x4d\x70\xc3\xf3\xa8\x69\x87\x20\x7f\xe3\x28\x94\xb3\x72\x4d\xd3\x51\xc9\x2f\x7f\x91\x63\xa2\x62\xe7\x96\x46\xdc\x95\x57\x83\x9d\xa9\x6e\x9b\x2e\x82\x21\xc1\x86\x9d\x4d\x6c\xc5\x9e\x41\x28\x31\x2c\x24\x2b\x86\x85\x32\x08\xa6\xdb\x66\x30\xf6\xab\x2f\xb3\x19\xfb\xd5\x97\xe3\x84\x4e\x7f\xbd\x1e\x54\xe2\x38\xe0\x84\xff\x50\x00\xf9\x03\x25\x90\xa4\x83\x47\x8f\xfb\x74\x30\x0b\x1b\xb0\x62\x44\xd9\xd8\x20\xc4\xdf\x65\xa9\x60\x90\x05\x12\xba\x07\x99\xe0\x73\x2e\xd7\xa9\xa4\x14\x47\xa9\xf0\x19\x94\x24\xdc\xd6\xe4\x4d\xd7\xe4\x68\x3d\xee\xf6\xfe\xfc\xcf\xde\xe9\xfb\x59\x4a\xb1\xae\x1e\xa1\xe1\xae\x7c\xa4\x9a\x6d\x34\x5a\x86\x0d\x68\x38\x8e\x39\x5c\x5f\x2e\xb1\xc8\x60\xa7\xdd\xc4\xa8\xa5\xe2\x48\xea\x15\x1a\x2c\x15\x1d\xf7\x57\x1c\xb6\x9b\x50\x15\x7b\x58\xee\x20\x4d\x64\xe6\x5e\xb6\x95\x1f\xa9\x18\xd4\x66\x13\x69\xa4\xe8\xb4\x6d\x02\x0c\x8f\x02\x9f\x58\xca\x87\x5d\x32\xd0\x2f\xd3\x63\xcc\x32\x7b\x8a\x74\xf1\x52\x2a\x1b\xb4\x8d\x26\x54\xde\x85\xc5\xa1\x54\x5e\x25\x99\x4a\xe7\x3a\xd3\xa9\x74\x6e\xd3\x79\xe5\x74\x2e\x2c\x40\xef\xf4\x7d\xf7\xe5\x5f\xce\xff\xfb\xa2\xf7\xdd\x59\xb6\x5c\x16\x31\x52\x09\x0a\xa2\x91\xa5\xef\x49\x20\x9a\xaa\x4b\x2a\x7a\xe0\xf3\x70\x5e\x35\xc7\x76\x09\x44\x1d\xb7\x00\xb9\x4a\xb5\x5e\xae\x29\x50\xa9\x2a\x1b\x90\x83\xf9\xb0\xe9\x31\x81\x4c\xc1\x56\xab\x85\x70\x42\x3e\xcc\x43\x0e\x44\xcf\x93\x2b\xb6\x8b\x30\x49\x5c\x10\x24\xf8\xf5\xf2\xda\x56\xb9\x9e\x78\xba\xae\xe2\x03\xdf\x97\x72\xd3\xe8\x01\x7d\x0a\x40\x25\x1a\x5d\x60\x5d\xd0\xe4\x02\x9c\xdd\x60\x43\x0e\xd2\x31\x84\xb7\x35\xba\x18\x58\x71\x49\xbc\x4d\x82\x16\x90\xce\x2c\xe8\x52\x3a\xff\x6f\xf8\x1a\xe2\xab\x5b\x7c\x5d\x03\xbe\x66\x6e\x43\x5b\x2d\x7d\x56\xea\x4a\x6f\x46\x23\xa0\xda\xda\x2c\x2d\x2b\xe5\xec\x78\xaa\x97\x15\x60\xcf\xa8\x41\xc9\x99\xc6\x6f\xee\x95\x6b\x65\x26\x03\x1b\x96\x8a\x4f\x7e\x85\x4e\x7c\x1f\x0a\xf0\x75\x4e\xe0\xc3\x6b\x08\xc5\x93\x7a\x80\xc4\xed\x9d\x44\xdf\xcc\xc3\xa2\x14\xe7\xfb\x0b\x23\x0f\x5f\x1c\xa9\x26\xdd\x1b\x2a\xa1\xf2\x99\x22\xdf\x0f\xe5\xf4\xef\xcc\xc1\xef\x3c\x78\x1e\x7d\x87\x25\x3d\x4c\x7f\x7f\xfc\x62\xc9\xfd\xf3\xd8\xc8\xe5\x24\x3e\x19\x24\x33\x92\xb9\xe2\x03\xf3\x66\xac\xf8\xdb\xb4\x24\xd2\x12\x9e\xac\x26\x25\xe6\xaf\xaf\x2f\xce\x5e\x26\xc7\x63\xd9\x72\xb3\x6a\xd8\xba\x48\x2f\xba\xc8\xa5\x86\x5e\x6e\xe6\x15\xe5\x63\xa9\x00\xb9\x7a\x79\xad\x5c\x54\x68\x3e\x2c\x79\x60\xe0\x37\x0f\xb9\xd5\xda\xc6\x7a\xa6\x54\x31\xca\xa3\xf6\xe4\xfb\x73\xa2\xe4\x18\x2f\xe0\xd6\xeb\x1a\x09\xac\x1a\xd8\x25\xa2\x11\x76\x97\x20\x16\x97\xbc\xf1\x5f\x7b\x20\x18\x66\xc7\x42\x13\xd6\x2a\xeb\x15\x05\x7e\x9e\x1b\x6c\x11\x71\xb8\x62\xd4\x18\xfa\xd8\x88\xf5\x9e\xdf\xef\xfe\xf4\xe3\xf9\x87\xfb\x33\xc6\x8d\xf2\xbd\xf8\x19\x86\x6b\xa3\x56\x2a\xd7\x60\xe5\xb7\x50\x29\x41\xa9\x5c\x2f\x66\x8e\x52\x4a\x84\xba\x8f\x1e\x77\x1f\x3c\x99\x25\x2a\x9b\x66\x5b\x3b\x48\x41\x93\xa5\xb6\xb6\xd9\xd2\x20\x79\x0c\x06\x84\x1b\xe1\xb7\x20\x2c\xc3\xc9\xcf\xb3\xcf\x2b\x48\x73\x6c\x3d\xa4\x16\x98\x2d\x46\xc7\xfb\x08\xa3\x31\x6c\xc7\x61\xe7\xd8\x1d\x37\xc1\x8e\xc1\xc4\x8d\xf4\x77\x8e\x4b\xef\xf0\x1c\x17\x85\x95\xf3\x05\x9e\x0b\x4e\x38\x54\x49\x7c\xc2\xe1\x70\x34\x3f\x62\x0c\x9b\x1a\x2c\xba\x8d\x5e\xb2\x44\x81\xe7\x00\x00\x76\xa3\x20\x47\xbf\x29\xef\x26\xf6\x0b\x3c\x47\x5f\x69\x3a\xa1\xb2\x78\x36\x1e\x19\xca\x96\xd9\xf0\x6c\x6e\xd7\xd0\xf3\x30\x47\x45\x30\x83\x53\x7c\x1b\x3f\xb4\xe6\x12\x0e\x32\x0f\x43\x6f\xb6\x77\x0d\x7d\x07\x0a\xcc\xb0\xd0\xf7\x49\xf8\x82\xc6\x09\xf4\xfe\x74\xd6\xfb\xe1\xbb\xee\x83\xef\x3f\x3e\xfc\xfe\xfc\x3f\x7f\xeb\xfe\xe1\x35\xdb\xf6\x0d\x22\xa0\xb1\x8a\x00\x5d\x25\xea\x0c\x48\x74\x45\x43\x77\x81\x75\xac\x99\xd1\x18\x0c\xe3\xb3\xe4\x85\xee\x0c\x3f\xa7\x19\xba\x2b\x41\xa1\x10\x7f\x50\x4b\x8e\xf4\xc3\x8f\x64\xdc\x68\xe1\xdf\x10\xc2\x73\x30\x4f\xf3\x91\x1b\x46\xba\x61\x83\x48\x45\xd5\x50\x0b\xa9\xe4\xdb\x36\x72\x89\xe1\xd8\xf4\x6c\x2f\x46\x4e\x30\xe6\xa1\xe7\x7c\x6e\x72\xfb\xef\xef\x0f\xb1\x66\xe8\x1d\x5a\x21\x74\x7b\x70\x26\xa0\x09\x60\x08\x0a\xe4\x6c\x1b\x7a\x87\xa2\xc2\xd0\x23\x4c\x4c\x2b\xb7\x64\xc7\xbf\xde\xaa\x8b\xca\x89\x4b\xd6\x12\x77\xc3\x85\x34\xd6\xde\x4b\x57\x51\xbf\xf3\x9f\x7f\xf8\x47\xc4\xd5\x61\xd5\x38\x36\x9a\xa1\x64\x36\x6c\x94\x42\xdd\xa9\xb7\xce\x0c\x18\xbe\x0e\x4e\x2e\x30\x46\x8e\x22\x96\xfa\x19\xc1\xd0\xa5\x30\xf6\x81\xbd\xa3\x28\x99\x9b\x8b\x7e\xb9\x87\xa6\x5c\xc6\xb8\xea\xd0\xb3\x14\x73\x63\x12\xad\x0d\x00\x28\x3d\xe8\xc1\x99\xa1\x1f\x6e\xd3\xb0\x11\x68\x4e\xdb\x26\x19\x87\x01\x74\x2b\x3d\x44\x68\xf1\xcc\x2c\x8d\x86\x52\x82\x5d\xdc\xd8\xaa\x2a\xe2\x1d\x29\x7b\x94\x07\x3f\x3c\xa6\xc4\x32\x8a\x23\x33\x26\xa6\xb5\xe9\x91\x4c\x50\xdd\xf8\xef\x9a\x29\x44\x38\xf4\xb9\x33\x68\x09\x0f\x5e\xf5\x4e\xdf\x06\x30\xee\x3d\x3d\xeb\xfe\xf4\xf7\xee\xdb\xa7\xdd\xe7\x6f\xd8\xc3\xe2\x3e\xd2\x0e\x80\xec\x07\xa0\x06\xc3\x85\x7d\xf5\x08\x7d\x9d\x2d\xd8\xf7\x54\x37\x89\xee\xf4\x1b\xd6\x27\xc2\x75\x74\xd6\xa0\x0c\xa4\x25\xe7\xa5\x53\x52\x94\x84\xfb\xcc\x69\x82\xab\xe7\x29\x34\xf6\x2e\x2c\xe6\xa9\x26\x81\xf7\xff\x17\x00\x00\xff\xff\x25\xf6\xd7\x9f\xee\x22\x00\x00")

func assetsTplCurdTplBytes() ([]byte, error) {
	return bindataRead(
		_assetsTplCurdTpl,
		"assets/tpl/curd.tpl",
	)
}

func assetsTplCurdTpl() (*asset, error) {
	bytes, err := assetsTplCurdTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/tpl/curd.tpl", size: 8942, mode: os.FileMode(511), modTime: time.Unix(1607498208, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsTplETpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x92\x41\x6b\xd4\x4e\x18\xc6\xcf\x1b\xc8\x77\x78\xd9\xc3\x9f\xe4\x4f\x98\xdc\x2b\x9e\x34\xde\xdc\x83\x0a\x9e\x27\xf1\x4d\x36\x34\x9b\xd9\xbe\x33\x61\x95\x25\x07\x65\x29\x2b\x28\xf4\xa0\xd2\xed\xa9\x42\x29\x78\xb0\xf1\xb8\x4a\xc5\x2f\xb3\x93\x6d\x4f\x7e\x05\x99\x64\xe3\xb6\x2a\x0a\x82\xa7\xf0\x3c\x79\x92\xe7\xf7\x0c\x93\x8e\xc6\x82\x14\x38\xb6\xd5\xeb\x3f\xe2\x8a\x87\x5c\xa2\x2f\xf7\xb2\xbe\x6d\x19\x2b\x49\xd5\xb0\x08\x59\x24\x46\xfe\x78\x37\xf1\x91\x48\x90\xec\xdb\x96\x6b\x5e\xfb\x3e\xe8\xf3\x67\x7a\xb9\xd4\x67\x47\xab\x8f\xcf\x6d\x4b\x3d\x19\x23\x04\x20\x15\x15\x91\x82\xa9\x6d\x95\x26\x16\x17\x79\x04\x03\x9c\x04\x8e\x0b\xff\x07\xc6\xee\x11\xaa\x82\x72\xf8\x2f\x98\x96\x9b\x90\xef\xc3\xc5\xd9\x07\xfd\xf9\xb5\x3e\xde\xaf\x8f\xe7\x9b\xaf\x1c\x84\xc0\x85\xfb\x8a\x47\xbb\x0e\x12\x41\xd3\xef\xb6\x8f\xe6\x47\x69\x0c\xa1\x07\x88\xb0\x73\x13\x90\xa5\x72\x50\x64\x99\x49\xba\x37\x20\x6c\x02\x5d\x15\xa2\x6d\xf5\x4a\xc0\x4c\xe2\x75\xbf\x59\xc4\x1e\xa6\x6a\xb8\xa9\x41\xd7\x24\xb7\x58\x2d\xd0\xe5\x62\xa6\x97\xa7\xab\x2f\x6f\xeb\xa7\xd5\xcf\x70\x77\x65\xb2\xe5\xf3\x60\x24\x13\x73\x08\x69\x9e\xfc\x1d\x2b\xd1\xef\x61\x89\x8f\x1d\xc4\xa6\xe7\x8f\xb0\x10\x0b\x1a\x71\xf5\x4b\xe6\x3b\x57\xa1\xdb\xdc\x86\xdb\x03\x4e\x89\x04\xc6\x58\x9a\x2b\xa4\x98\x47\x38\x2d\xff\xd9\x98\xb8\x59\xd3\x02\xb4\xcd\x8c\xb1\x1f\x97\xcd\x4f\xea\x37\xef\xeb\xc3\x4a\x1f\x9c\xd6\x87\xd5\xfa\xdd\xa7\xf5\xd1\xec\xf2\xd5\xe2\xa2\xaa\xbe\x9e\xbf\xd0\xf3\xc5\x6a\xf9\x52\x9f\xcc\xd6\x07\xfb\xd7\xa6\x6e\xc9\xba\xcb\xe3\x84\x42\x64\x5e\xa7\x0c\x91\x9c\xa4\x2a\x1a\x1a\xa7\x91\x11\x97\x08\x72\x2f\x63\x01\xd1\x40\xdc\x13\x13\xe9\x75\xf2\x96\xc8\xf3\xdb\x22\xc7\xef\xc6\x83\xc7\x46\xee\x5c\x99\xa5\xa8\x40\xaf\x5b\xbd\xbd\xea\x31\xcf\x64\xe7\x97\xb6\xf5\x2d\x00\x00\xff\xff\x3e\xa9\x79\x47\x77\x03\x00\x00")

func assetsTplETplBytes() ([]byte, error) {
	return bindataRead(
		_assetsTplETpl,
		"assets/tpl/e.tpl",
	)
}

func assetsTplETpl() (*asset, error) {
	bytes, err := assetsTplETplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/tpl/e.tpl", size: 887, mode: os.FileMode(511), modTime: time.Unix(1607498208, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsTplEntityTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x8e\x41\x0a\xc2\x40\x0c\x45\xf7\x85\xde\x21\x8b\x2e\xa5\x07\x10\x5c\x09\x5d\x76\x21\xbd\xc0\x68\x83\x54\x67\x3a\x25\x4d\x29\x25\xe4\xee\x12\x75\xa4\x54\xc1\x59\x4d\x5e\xfe\x0f\x2f\xcf\x44\xca\xc6\x9d\x3d\x1e\x63\x08\xd8\xb3\x6a\x9e\xf1\x32\x20\x24\xae\x0a\x23\xd3\x74\x61\x10\x0b\x93\xeb\xaf\x08\xc5\x6d\x07\x45\xc7\x18\x60\x7f\x80\xb2\xea\xd0\xb7\xa3\xaa\xc8\x93\x95\xb5\x0b\x56\x7b\xbd\x04\x9b\x65\x78\xc3\x44\xaa\x48\xc1\x71\x2a\xc3\x26\x7f\xc2\xe0\xe8\x6e\x3a\x22\xd8\xb7\xf6\xd1\x2f\xdb\x7a\xf2\x7e\x6d\x6c\xf3\x1f\x6b\x8a\xf3\x56\x9a\xe2\xbc\x72\x4e\xc0\x4e\x7d\x9c\x61\xb5\xf9\x65\xf6\x08\x00\x00\xff\xff\x6e\x2e\x32\xff\x48\x01\x00\x00")

func assetsTplEntityTplBytes() ([]byte, error) {
	return bindataRead(
		_assetsTplEntityTpl,
		"assets/tpl/entity.tpl",
	)
}

func assetsTplEntityTpl() (*asset, error) {
	bytes, err := assetsTplEntityTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/tpl/entity.tpl", size: 328, mode: os.FileMode(511), modTime: time.Unix(1607498208, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsTplExampleTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x54\x6d\x6b\xdc\x46\x10\xfe\x2e\xd0\x7f\xd8\x0a\xee\x22\x99\x8b\xe4\xd4\x49\x09\x82\xa3\xed\xbd\x84\x06\x92\xd4\xc5\x2e\xfd\x50\x4a\xd8\x93\xf6\xee\x16\x4b\xbb\xba\xdd\x55\x1a\xd7\x18\x2e\x50\x9a\x84\x24\x76\x20\x4d\x02\xad\xd3\x77\xf7\x85\x42\xed\x0f\xa5\x31\x97\x94\xfe\x19\x4b\x67\xff\x8b\x32\x92\xee\xcd\x75\x8a\x69\x09\xe4\xdb\xec\xec\xcc\xce\xcc\xf3\x3c\xb3\x11\xf6\x56\x70\x87\xa0\x70\x55\xf6\x02\x5d\xd3\x35\x1a\x46\x5c\x28\x64\xea\x9a\xe1\x63\x85\x5b\x58\x12\x47\xf6\x02\x43\xd7\x8c\x76\xa8\x0c\x5d\xbb\x8a\x8c\x0e\x55\xdd\xb8\x65\x7b\x3c\x74\x3a\xfc\xb4\xec\x05\xa7\x7d\x41\xaf\x11\xe1\x64\xaf\x40\xa8\x22\x52\x51\xd6\x01\x53\x2a\x41\x59\x47\x82\xc9\x88\x72\x62\x01\x01\x16\x94\x52\xab\x11\x41\x8d\x5a\x9d\xb3\x36\xed\x34\x99\xa2\x6a\x15\x49\x25\x62\x4f\xa1\x35\x5d\x7b\x87\x4b\x85\x10\x42\x79\x3e\x72\x9c\x64\x6b\x37\x79\xd2\xd7\xb5\x45\xe8\x0f\x21\x44\x99\x42\x8e\x33\xfc\x75\x27\xd9\xfc\x5e\xd7\xae\xe0\x90\xcc\x84\x0f\x3f\xff\x39\xbd\xf5\x54\xd7\x16\xb1\x94\xb3\xef\xec\x7c\x36\xfc\xe6\x86\xae\x35\x6a\x79\xce\xe4\x62\xf0\x20\xb9\x7f\x4f\xd7\xea\x5d\x2c\x24\x51\x53\x4f\x3d\x7f\x94\x65\x2c\xd3\x90\x7c\xc2\x19\x99\xdc\xa4\x8f\xff\x48\xee\x0e\x74\xed\x32\xbe\x7e\xd1\x0f\x48\xd1\x53\xba\xd5\x4f\x7e\xf8\x69\xf8\xcb\xe0\xf0\xf1\xef\x07\x7f\x7d\x95\x6e\x6c\x67\x11\xef\x46\x84\xcd\x44\xe4\x77\xe9\xc3\x5d\x5d\x5b\x07\x40\x1c\x07\x8d\x5d\xe9\xbd\xdf\x92\xc1\x03\x5d\x6b\xc7\xcc\x43\x17\x19\x55\x8d\x9a\xe9\xb5\x3b\x47\xf0\xb2\xd0\x9c\xec\x05\x76\xa3\x06\x88\xd1\x76\xd1\x98\xb4\x9b\xbd\x18\x07\x17\x78\xe0\x43\x8e\x3d\xea\xbb\x82\x0c\xc3\x82\xc8\x69\x27\xaa\x22\xe3\xd4\xdb\x92\x62\x67\xa9\x8b\x59\xa7\x8b\xe9\x29\x23\xeb\xc7\x97\x0c\xb9\x55\xd4\x0e\x95\xbd\x14\x09\xca\x54\xdb\x34\x4a\xd2\x2d\xc9\xb7\x94\x17\x99\x60\xf9\x96\x53\x92\x6f\x7a\x39\x5c\xd5\x92\x2c\x47\x60\xc1\xc3\x55\x25\x62\x52\x0e\xb8\x57\xbd\xc4\x3d\x1c\x94\x15\x0d\xc9\x55\xa8\x56\x2d\x49\xa3\x92\x37\x00\xe8\x17\x26\x70\x54\x98\x40\xfb\xc8\xcb\xc5\xc8\xcc\xb9\x2a\x0e\x05\x3f\x15\x5d\x8b\x45\x60\xbf\x17\x13\xb1\xda\x94\x1e\x8e\xc8\xcc\xb0\x56\x25\x93\x99\xc7\x19\x23\x9e\xa2\x9c\x55\x10\x11\x02\x26\x02\xc4\x80\x0b\xd3\xc8\xf5\x5a\x41\xbe\x64\x56\x86\x1f\x44\xbc\x56\x45\x8c\x06\x00\x53\x84\x19\xf5\x4c\x22\x84\x95\x01\x32\x79\xca\x5e\x22\xaa\x20\xb4\xce\x19\x93\x59\xe1\xc2\x61\x1d\x13\x08\xda\x98\x09\x04\x87\xa5\x6b\x82\xa8\x58\x30\x34\x89\xcf\xea\x5c\xc3\x02\x35\x6a\x23\x66\xf3\xf3\xda\x5a\x86\xd7\xfa\x7a\xa3\x85\xe6\xc6\x87\xcb\xdc\x27\xd9\xd6\x66\x2a\xa1\x8c\xaa\xfa\xf8\x29\x73\xc4\x34\x8c\x3c\xab\x9a\x62\xbb\x5c\x23\x00\x72\xba\x5c\x2a\xe0\x04\xe0\x76\x17\x16\xe6\xcf\x57\xf2\x6d\x72\x0d\xc1\x79\x7e\x83\xa5\x74\x8d\x33\xaf\x2f\x9c\x3d\xf7\x06\x9c\x73\x36\x5c\x63\x85\x32\x3f\x20\x3e\xb8\x0a\x4e\x5c\x23\x56\xed\xf3\x61\xeb\x2c\xf8\x0a\x40\xdc\x33\xf3\x95\xf1\x86\xb8\xe8\x5c\x25\x1b\xb2\x51\x43\xd5\x29\x5d\x5b\xba\x36\x3d\x62\x15\x5d\x21\x1f\x8f\x1d\x66\xa3\x66\x4d\x56\x24\xfd\x7a\xfb\x60\xe7\xbb\xf4\x76\x3f\xdd\xba\x3d\xfc\xe2\xd3\x7c\x57\x0a\x08\x96\x89\x54\xe3\xb4\x0b\x94\xf9\x1f\x74\x89\x20\xa6\x42\x73\xc5\x8f\x64\x2f\x67\xb0\x1c\x85\x4a\xd7\x5a\x9c\xaf\x00\x52\xe3\xec\x35\x28\xa6\xb8\xcf\xb3\xca\x82\xc8\x38\x50\x63\x09\x4d\xf5\x6a\x43\x19\xb3\x0c\xf9\x16\x7c\x22\x37\xff\x4c\x1f\xee\x1e\xde\xbc\x7b\xb8\xd5\x3f\xf8\xf1\x46\x7a\xe7\x79\xba\xb1\xbd\xbf\xd7\xdf\xdf\xbb\x93\xfd\xb0\xff\x10\x99\xb2\x9b\x42\x70\x31\x91\x19\x2c\xdc\x22\xec\x5b\xc0\xcc\xbc\xec\x87\xf3\x1f\xe5\x57\xf0\x41\x6c\x3c\x4d\x36\x1f\xc1\xff\x71\x7f\x63\x7f\xaf\x9f\x3e\xf9\xf6\xc5\x00\x5c\xc2\x52\x9d\x64\xf6\x7f\x1d\x4e\x48\x65\x32\x1a\x1c\xb7\x1f\x27\x68\xfd\x68\xdf\xfd\x67\x19\x3a\x9b\xc7\xb5\x5b\xe7\x31\x7b\x65\xfa\x4d\x6e\x7d\x99\x3c\x1b\xbc\x18\xdb\xba\x20\x58\x9d\x48\x59\xbe\xff\xdf\x74\x55\x54\x28\xfb\xfe\x4b\x9a\xa5\x79\x1d\x87\x51\x40\xde\x8f\xfc\x97\x3c\x49\x51\xe1\xff\x4d\xf2\x77\x00\x00\x00\xff\xff\x0f\x0d\x50\xf5\xaa\x08\x00\x00")

func assetsTplExampleTplBytes() ([]byte, error) {
	return bindataRead(
		_assetsTplExampleTpl,
		"assets/tpl/example.tpl",
	)
}

func assetsTplExampleTpl() (*asset, error) {
	bytes, err := assetsTplExampleTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/tpl/example.tpl", size: 2218, mode: os.FileMode(511), modTime: time.Unix(1607498208, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsTplGormTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x55\x5d\x6b\x13\x41\x14\x7d\xde\xf9\x15\xd7\x80\x65\xa6\x2c\x53\xf1\x31\xb2\x48\xd3\xb4\x22\xd8\x50\x88\xe2\x43\x29\x74\x37\x7b\x13\xd7\x4e\x36\x61\x66\x36\xa1\x2c\x0b\xbe\x69\xc1\x57\x0b\x7e\xe0\x47\xd1\xd7\xb6\x82\x2f\x2a\xfe\x9c\x84\xfa\x2f\x64\x66\xd3\x1a\xd7\x4d\x68\x0a\xe2\xd3\x32\x93\x7b\xee\x3d\xf7\x9c\xc3\x84\xe8\xfd\x3e\x42\x9a\xf2\xa6\x96\x49\x4b\xdf\xf7\x03\x81\x0d\xbf\x8b\x59\xb6\xd9\x0b\x51\x80\xb2\xd7\x90\x12\xa7\x5e\x83\xe5\x4e\x4f\x76\x79\xbd\x46\x9c\x2d\x80\xe5\x32\x10\xc9\x08\x19\xf8\x12\xa8\xa9\xbf\xab\x1a\x91\x58\x97\x12\x3c\x40\x29\x7b\x52\xf1\x06\x0e\x69\x25\x0c\x20\x52\x10\x47\xa2\xc2\x08\x23\xa4\x9d\xc4\x2d\xa0\xaa\xbc\x1f\x83\x8b\x03\x65\x86\x4c\x14\x77\x0c\x19\x89\x3a\x91\x31\x54\xd2\x94\x4f\x55\x57\xcc\xf8\xb9\xfd\xec\x52\x0c\xee\xa0\x5e\x15\x62\xad\x27\x92\x6e\xac\xca\x1a\xef\x1a\x49\x56\x85\xd8\x88\x50\x84\xf7\x22\xa5\xb3\x0c\x76\x4d\xf7\x95\x15\x68\xe0\x70\xb6\x5e\xe3\xc3\xd3\xd1\xf7\x6f\x9b\xbe\xd2\x28\x6f\x2a\xe1\x0f\xd0\xde\x8f\x4e\xbe\x9e\x7d\xfe\x90\x73\x9b\x87\xa7\x61\x70\xa1\x32\x9b\xb3\xc2\x14\xd5\xa5\x99\x45\x29\x71\x9c\x7a\xad\x0a\x61\xe0\x12\xc7\xd9\xaa\x42\x79\x6d\x9a\xb9\xc4\xc9\x26\xcb\xad\x49\xf4\x35\xda\x35\x8e\xde\x9e\x1d\x9f\x8e\x7e\xbc\xb8\x94\xa2\x39\x8e\x0e\x7c\x91\x94\xc7\x89\xe5\x19\x30\xc4\xa3\x36\x28\x5e\xaf\x81\xe7\x99\x14\x98\x9b\xf3\x5d\x7e\x67\xc6\x30\x32\x85\x28\x25\x54\x3d\x5b\xcf\xeb\x18\x24\x1d\xca\x72\xc7\xa9\xe2\x5b\x7c\x2a\x1c\x8c\x37\xfd\x01\xd2\x25\xcb\x80\xf1\x75\x33\xec\x96\x85\x5f\xfb\x6b\x0c\x4e\xfa\x4f\x8e\x71\x24\x26\xdb\xd7\x51\xa0\x46\x18\x3d\x7b\xff\xf3\xe5\xc7\xf1\xf3\xa7\xa3\xe3\x57\x0b\x68\x90\xa3\x69\x68\x3f\x0f\x1f\xa1\xc4\x4d\xbf\x0f\x5d\xbf\xbf\x9d\xe7\x6b\x27\x8a\x35\xca\xb6\xdf\xc2\xf4\x4a\x72\x84\xc1\x25\xa5\x20\x4e\xbb\x27\x61\x0f\xf7\x5d\xc8\x0d\xa9\x7a\x20\xfd\xb8\x83\x50\xe0\x66\x66\x0d\xcd\xa1\xa9\xad\xcc\xed\xae\xe6\xcd\xbe\x8c\x62\xdd\xa6\x95\xeb\x03\xf0\xe0\x76\xc5\x35\x8d\x18\x71\xcc\x78\x0f\xc2\x80\x5b\x34\x3d\x87\x4d\x46\xb0\x82\x61\x61\xc0\x27\x72\x94\x47\xee\xaa\x0e\x3d\xe8\x87\x36\x9f\xaf\xbf\x8c\x0f\x4f\x17\x76\x28\x47\xd3\xc4\x7e\x66\x9b\xe3\xc2\xf0\x7f\xba\x37\xdb\xbe\xe1\x3f\x32\xae\xe0\x5c\x51\xa6\xab\x9a\xd5\x44\x81\xad\xa9\x97\x16\xc6\xef\x3e\x9d\x9d\x1c\x8d\x0f\x9e\x8c\xdf\x1c\x2c\x60\x5b\xb1\x0f\x65\x40\xb7\x77\xca\x00\x6e\x6e\x0c\x9b\xef\x4c\x1c\x09\xb7\x68\x4f\x2f\x78\xac\xcc\xfa\x5d\x7f\x0f\x67\x36\xbf\xc1\x16\x7d\x94\x88\xe3\xe4\xec\xa9\xe2\x85\x7f\x1d\xfb\xe3\x46\x14\x87\x74\xc9\x0c\x9f\x25\xf2\x1f\xa4\x0b\x52\x1b\x9c\x9b\x0b\xfe\x2b\x00\x00\xff\xff\xc6\x35\xf2\x12\xc9\x07\x00\x00")

func assetsTplGormTplBytes() ([]byte, error) {
	return bindataRead(
		_assetsTplGormTpl,
		"assets/tpl/gorm.tpl",
	)
}

func assetsTplGormTpl() (*asset, error) {
	bytes, err := assetsTplGormTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/tpl/gorm.tpl", size: 1982, mode: os.FileMode(420), modTime: time.Unix(1624268847, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsTplInitTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x54\x5d\x8b\xe4\x44\x14\x7d\xee\x40\xfe\xc3\x35\xd0\xb3\x89\x84\xce\x7c\xec\x2e\xd2\xd2\xb4\xf6\xcc\x8a\x82\xa3\xee\xce\xbc\x4b\x4d\x52\xe9\x2e\x36\xa9\xca\x54\x55\x86\x19\xa5\x61\x7d\x11\x1f\x5c\x15\x56\xdd\x97\x55\x04\x15\x15\xc1\xdd\x07\x51\x18\x04\x7f\x4d\xf7\xe8\xbf\x90\x5b\x55\xe9\x4e\xda\x1d\xf6\x29\xa9\x73\x4f\xdd\x8f\x53\xa7\x8a\x95\x95\x90\x1a\x42\xdf\xeb\x05\x19\xd1\xe4\x84\x28\x9a\xa8\xd3\x22\x40\x20\x2f\xb5\xf9\x72\xaa\x93\x5a\x5a\x4c\x69\xc9\xf8\x54\x05\xbe\x17\xf9\x9e\xef\x9d\x11\x69\x36\x97\x44\x69\x2a\x0f\x26\xf0\xb2\x3a\x2d\x06\x07\x13\xdf\xeb\xa9\x82\x9c\xd1\x36\x62\x36\xe4\x35\x4f\x81\x71\xa6\xc3\x08\x3e\xf4\xbd\x5e\x9a\x4f\x61\x38\x82\xec\x64\x5f\xf0\x9c\x4d\x11\xea\xbd\x29\x94\x1e\x02\x00\x04\x85\x48\x49\x31\x13\x4a\x07\x31\x06\xde\x13\xd2\x06\xf6\xf6\xb6\x6f\x1b\xe4\x1d\x52\x52\x4b\x95\x42\x34\x2c\xa2\x94\xc5\x76\x76\xf7\x6e\xde\xba\x6d\xd1\x83\x89\xe3\x06\xf7\x19\xcf\x0a\x9a\x59\x78\x7f\x46\xa4\xa2\x7a\x08\x41\xad\xf3\x57\xca\x93\x9b\x16\x3e\x24\xe7\xef\x56\x94\x0f\x61\x67\x7b\xbb\x01\xde\xca\x0a\x3a\x84\x5b\x66\x3d\x6f\xcf\x3c\x32\x03\x1d\x4c\xc2\x34\x9f\x46\xbe\xd7\x4b\x92\x66\xf6\x8d\xc8\x1c\x15\xd0\x17\x15\x5d\xcd\x0b\x4a\xcb\x3a\xd5\x46\x0a\x1c\x1b\xdb\x06\xab\x31\x24\xc9\xe2\xc9\xb3\xc5\x37\x0f\x7c\xcf\x0c\x6e\x42\x8c\x9b\x6f\x92\x5c\xfd\xfa\x74\xf1\xf9\xf7\xbe\x67\x14\xe8\xee\xba\xfa\xf2\xe7\xe5\x27\x7f\xe2\x2e\xa2\xd4\x46\xc2\xa7\x1f\x5f\x7d\xf7\x91\xef\x39\x35\x3a\xa1\xcb\x47\x8b\x2f\x1e\xfa\x5e\xa3\x48\x3b\xe1\x5f\x5f\xdb\x5d\xc7\xac\xa4\x1f\x08\x4e\xd7\xa1\xe5\xe3\x3f\x16\x9f\x5e\xfa\x5e\xa3\x4f\xab\xc3\xe5\x93\x07\x8b\x1f\x7e\xba\xfa\xe5\xf2\xdf\xc7\xbf\xff\xf3\xf7\xb7\xcb\xcf\x7e\xb4\x34\xd4\xf5\x7f\x34\x4b\x58\x7e\xf5\xcc\x89\x94\x24\xb0\x82\x96\x0f\x7f\x5b\x5c\x3e\x6a\x59\xc7\xea\xb9\xd2\x30\x6a\x2c\x66\x54\x64\xb9\xeb\x4e\x0d\xee\x9c\xd6\xa4\x78\x43\x14\x19\xd2\x07\x4d\xf3\x31\x04\x81\xf5\x5e\xaf\x0d\xc3\x08\x82\x1b\xaf\x2b\x46\x92\xa3\x19\xe1\xd3\x19\x61\x37\x02\x77\xd0\x99\xe2\xe8\xd1\xbc\xd4\x83\xa3\x4a\x32\xae\xf3\x30\xe8\xab\x61\x5f\xbd\xa6\xd3\x2a\xc4\xbf\x2c\x4a\xfa\x6a\x9c\x5a\xe5\x46\x7d\xb5\x55\xe1\x1f\xa6\x1e\x69\x59\xd3\xad\x42\xa4\xa3\xb7\xd1\xcb\x5b\x9a\x95\xf4\x7d\xac\x37\xea\x2b\x6b\x35\x6c\x02\x0f\x63\xb5\xc0\x63\x5b\x2d\xd0\x14\xeb\x88\x90\xeb\x85\x3d\xc2\xd5\xd2\x1d\x9b\x59\xd7\xb2\x18\xdc\xad\xa9\xbc\xb8\xa3\x52\x52\xd1\xce\xf8\x11\x32\xd0\xa5\x19\xcd\xa9\x04\x14\xd5\x5d\x45\x94\x8e\x4a\x89\xa3\x4a\x9a\x8a\x33\x2a\xc3\xe8\x55\x83\xbc\x34\x02\xce\x0a\x4b\x72\xac\x1d\xa4\x1d\xba\x1b\xb0\x5f\x08\x45\x1d\x79\xa7\xcb\xee\x55\x84\xb3\x34\xc4\x40\x64\x80\x79\x2b\xc9\x2e\x26\x39\xb2\x77\xa5\x9d\x63\xf7\xba\x1c\xbb\xed\x1c\x2b\xd4\x80\x88\xcd\x43\xfc\x4d\x05\xe7\x34\xd5\x4c\xf0\xb8\x99\x07\xfd\x81\xbe\x0b\x83\xf2\x02\x5f\xb7\x18\x32\xc5\x23\xeb\x96\xcd\x01\x3b\x69\xe7\x9d\x7c\x83\x23\xaa\x9d\x83\xf7\x05\xe7\xca\x08\xeb\x80\xe8\x79\x4c\xbc\x12\x1d\x26\x02\xc8\x94\x54\xd7\x92\xc3\x7a\x83\xf1\x7d\x92\x40\x8a\x32\x40\x76\xe2\xec\xbe\xa1\x30\x36\x2b\x64\xe3\xf3\xd5\x03\xd4\x6e\xdf\x65\x6e\x62\x03\xb7\xd3\xcd\xe2\xa2\x9c\x15\xd7\x14\xec\x9e\x46\xb7\x5e\xf3\xaa\x3d\xa7\x9c\x0b\xbd\xb0\x9a\x96\x84\x2b\x62\x26\x06\xa5\x89\xd4\xae\xec\xf1\xf9\x84\x4e\x19\x0f\x23\x08\xcd\x65\x3e\x3e\x8f\x6d\x6d\x6b\xcd\xcd\xa1\x1c\xb9\xc9\x2a\x69\x45\x89\x06\x49\x55\x25\xb8\xa2\xa0\x05\x8c\xe3\x71\x3c\x76\xc9\xef\x99\xf0\xdd\x9a\x2a\xac\x7b\x48\xe4\xfd\x30\x15\x35\xd7\xf8\x08\x45\xcd\x5b\xd6\x2a\xd3\xbc\x1f\xc7\x92\x95\xf7\xd8\x74\xa6\xc3\x06\xb1\x99\xc2\x60\x1c\x07\x31\x98\x1c\x51\x0c\x41\x1c\x60\x27\xff\x05\x00\x00\xff\xff\xe3\x1c\x96\xb0\x4b\x07\x00\x00")

func assetsTplInitTplBytes() ([]byte, error) {
	return bindataRead(
		_assetsTplInitTpl,
		"assets/tpl/init.tpl",
	)
}

func assetsTplInitTpl() (*asset, error) {
	bytes, err := assetsTplInitTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/tpl/init.tpl", size: 1867, mode: os.FileMode(511), modTime: time.Unix(1607498208, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsTplMarkdownTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x92\xc1\x6a\xea\x40\x14\x86\xf7\x81\xbc\xc3\x81\xb8\x50\xb8\xc9\x03\x08\xf7\x6e\xae\x1b\xb9\x17\xeb\x42\xba\x11\x17\xb1\x9e\x16\x69\x92\x96\x68\xa0\x61\x66\xc0\x45\xa1\x85\xd2\xd6\x45\xc1\xac\x0a\x05\x11\xbb\xa8\x52\xda\x82\x98\x3e\x8e\x19\xf5\x2d\x4a\x66\x9c\xd4\x56\x9a\xc5\x84\xf9\xff\xc3\xf9\x4f\xbe\x13\xc3\x80\xd5\xc3\x98\xf7\x62\x3e\x3a\xd7\x35\x5d\xa3\xc9\xfc\x36\xb9\x1c\xd0\x54\x4d\xfa\xd7\x40\x61\x35\x7d\xe5\xd1\x0d\x50\x48\x86\x17\xfc\x65\x0c\x54\xd7\x68\xd1\x34\x4d\x71\xa8\x13\xe4\x4b\xd7\x08\xf1\x6d\xef\x08\xc1\xaa\xd9\x4d\x07\xff\xb7\x3b\x5d\xc6\x28\x10\x62\x95\xbd\x16\x9e\x31\x46\xeb\xe9\x45\x98\x15\xdb\x45\xc6\x1a\x79\xe3\xd3\xfd\x6a\x15\x40\x3e\x94\x10\xeb\xef\x89\xeb\xa2\x27\x9a\x89\x14\xf4\x5a\x8c\xa5\x03\xab\xc0\xdc\x31\x86\xbf\x20\xd7\xee\xa2\x0b\xc5\xdf\x60\x95\xb0\x73\x20\xe3\x75\xcd\x30\x0c\x20\x44\x78\x2a\xc9\x52\xf7\xad\x3c\x5d\xfb\x93\x95\x65\x79\x69\x86\x29\x68\x3c\x0d\xf8\xe4\x4d\xd2\x90\xa4\x84\x00\x14\x96\xcf\x71\x72\x7f\x05\x14\x78\x34\x4d\xfa\xa3\xc5\x6c\xbe\x7c\x9c\x53\x58\xc7\xd1\x6a\x32\x4c\x7a\xef\x99\xc3\xa3\xe9\x62\x16\xaf\xef\x26\xbb\x58\x15\x55\x73\x97\xed\x37\x61\x0b\xb2\x1c\x75\x03\x59\x40\x72\x02\xd7\x93\x1f\x03\x20\xb0\x4b\xa9\x16\x9e\x62\x25\x70\x9b\xe8\x33\x26\xf5\x72\xa7\x12\x38\xce\x66\x37\x25\x3c\xb4\x03\xa7\xbb\x6f\x3b\x01\xaa\x82\xaa\xdf\x76\x6d\x3f\xfc\x87\xa1\x52\x64\xab\x1f\x16\x51\xaf\xed\x55\x1b\x79\x23\xfb\x97\x0a\x99\xf7\x11\x00\x00\xff\xff\x54\x04\xaf\xad\x62\x02\x00\x00")

func assetsTplMarkdownTplBytes() ([]byte, error) {
	return bindataRead(
		_assetsTplMarkdownTpl,
		"assets/tpl/markdown.tpl",
	)
}

func assetsTplMarkdownTpl() (*asset, error) {
	bytes, err := assetsTplMarkdownTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/tpl/markdown.tpl", size: 610, mode: os.FileMode(511), modTime: time.Unix(1607498208, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsTplModelsTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\xcb\x3d\x0a\xc2\x40\x10\xc5\xf1\x7e\x4e\x31\x45\x4a\xd9\x03\x08\x56\x42\x4a\x0b\xc9\x05\x46\x33\x48\x34\xb3\x09\x9b\xb1\x58\x86\x77\x77\xf1\x63\x1b\x5f\xf9\xe3\xfd\x29\x22\x0d\x72\x99\xf5\xb8\x98\x69\x76\x80\xbc\xae\xca\x8d\x01\xde\xbc\x3c\xaf\xce\x41\x11\x45\xf2\x4d\xb9\xbb\xef\xb8\x9b\x5c\x8d\xf7\x07\x4e\xfd\xa4\xf3\xb8\x01\x11\x1f\x4b\x27\xb1\x77\xf5\x5d\xc3\xa1\xae\x3f\x6c\xd2\x2f\xc5\xc4\x5b\xcc\x7f\xff\xb3\x9a\x94\x07\x40\x11\x9a\x47\x80\xf0\x0a\x00\x00\xff\xff\x21\x54\x8c\xe7\xa9\x00\x00\x00")

func assetsTplModelsTplBytes() ([]byte, error) {
	return bindataRead(
		_assetsTplModelsTpl,
		"assets/tpl/models.tpl",
	)
}

func assetsTplModelsTpl() (*asset, error) {
	bytes, err := assetsTplModelsTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/tpl/models.tpl", size: 169, mode: os.FileMode(420), modTime: time.Unix(1624246413, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsTplTablesTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xe2\xe5\x4a\xce\xcf\x2b\x2e\x51\xd0\xe0\xe5\xaa\xae\x2e\x4a\xcc\x4b\x4f\x55\x50\xc9\xd4\x51\x50\xc9\x2c\x49\xcd\x55\xb0\xb2\x55\xd0\xab\xad\x05\xc9\x80\xf9\x7a\xa1\x05\x05\xa9\x45\x21\x89\x49\x39\xa9\x7e\x89\xb9\xa9\xb5\xb5\x0a\x0a\x0a\x0a\xb6\x0a\x4a\x30\x69\x24\x19\x25\x05\x7d\x7d\x05\x98\xb8\x73\x7e\x6e\x6e\x6a\x5e\x49\x6d\x6d\x75\x75\x6a\x5e\x0a\xc8\x40\x4d\x40\x00\x00\x00\xff\xff\xf6\x4d\x87\xcf\x77\x00\x00\x00")

func assetsTplTablesTplBytes() ([]byte, error) {
	return bindataRead(
		_assetsTplTablesTpl,
		"assets/tpl/tables.tpl",
	)
}

func assetsTplTablesTpl() (*asset, error) {
	bytes, err := assetsTplTablesTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/tpl/tables.tpl", size: 119, mode: os.FileMode(511), modTime: time.Unix(1607498208, 0)}
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
	"assets/tpl/curd.tpl":     assetsTplCurdTpl,
	"assets/tpl/e.tpl":        assetsTplETpl,
	"assets/tpl/entity.tpl":   assetsTplEntityTpl,
	"assets/tpl/example.tpl":  assetsTplExampleTpl,
	"assets/tpl/init.tpl":     assetsTplInitTpl,
	"assets/tpl/markdown.tpl": assetsTplMarkdownTpl,
	"assets/tpl/models.tpl":   assetsTplModelsTpl,
	"assets/tpl/gorm.tpl":     assetsTplGormTpl,
	"assets/tpl/tables.tpl":   assetsTplTablesTpl,
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
	"assets": &bintree{nil, map[string]*bintree{
		"tpl": &bintree{nil, map[string]*bintree{
			"curd.tpl":     &bintree{assetsTplCurdTpl, map[string]*bintree{}},
			"e.tpl":        &bintree{assetsTplETpl, map[string]*bintree{}},
			"entity.tpl":   &bintree{assetsTplEntityTpl, map[string]*bintree{}},
			"example.tpl":  &bintree{assetsTplExampleTpl, map[string]*bintree{}},
			"init.tpl":     &bintree{assetsTplInitTpl, map[string]*bintree{}},
			"markdown.tpl": &bintree{assetsTplMarkdownTpl, map[string]*bintree{}},
			"models.tpl":   &bintree{assetsTplModelsTpl, map[string]*bintree{}},
			"gorm.tpl":     &bintree{assetsTplGormTpl, map[string]*bintree{}},
			"tables.tpl":   &bintree{assetsTplTablesTpl, map[string]*bintree{}},
		}},
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
