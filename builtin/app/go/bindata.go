// Code generated by go-bindata.
// sources:
// data/aws-vpc-public-private/build/template.json
// data/common/dev/Vagrantfile.tpl
// data/common/dev-dep/build/Vagrantfile.fragment.tpl
// data/common/dev-dep/build/Vagrantfile.tpl
// data/common/dev-dep/build/build.sh.tpl
// data/common/dev-dep/build/upstart.conf.tpl
// DO NOT EDIT!

package goapp

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"reflect"
	"strings"
	"unsafe"
	"os"
	"time"
	"io/ioutil"
	"path"
	"path/filepath"
)

func bindataRead(data, name string) ([]byte, error) {
	var empty [0]byte
	sx := (*reflect.StringHeader)(unsafe.Pointer(&data))
	b := empty[:]
	bx := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	bx.Data = sx.Data
	bx.Len = len(data)
	bx.Cap = bx.Len

	gz, err := gzip.NewReader(bytes.NewBuffer(b))
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
	name string
	size int64
	mode os.FileMode
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

var _dataAwsVpcPublicPrivateBuildTemplateJson = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xaa\xe6\x52\x00\x02\xa5\xdc\xcc\xbc\xf8\x82\xc4\xe4\xec\xd4\xa2\xf8\xb2\xd4\xa2\xe2\xcc\xfc\x3c\x25\x2b\x05\x25\x03\x3d\x0b\x3d\x03\x25\xae\x5a\x2e\x40\x00\x00\x00\xff\xff\x1b\x82\xd0\x08\x26\x00\x00\x00"

func dataAwsVpcPublicPrivateBuildTemplateJsonBytes() ([]byte, error) {
	return bindataRead(
		_dataAwsVpcPublicPrivateBuildTemplateJson,
		"data/aws-vpc-public-private/build/template.json",
	)
}

func dataAwsVpcPublicPrivateBuildTemplateJson() (*asset, error) {
	bytes, err := dataAwsVpcPublicPrivateBuildTemplateJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/aws-vpc-public-private/build/template.json", size: 38, mode: os.FileMode(420), modTime: time.Unix(1435862031, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _dataCommonDevVagrantfileTpl = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x94\x5b\x6b\xeb\x46\x10\xc7\xdf\xf5\x29\xa6\xf2\x69\x4f\x0b\xd1\x8a\x13\x42\x1e\x4c\x62\x08\x29\xa4\x79\x28\x49\x1b\x37\x2f\xa5\xc4\x6b\x69\x24\x2d\x59\xed\x6c\x77\x57\xbe\xe4\xf2\xdd\x3b\x2b\xc9\x89\x1d\x48\x39\x60\xb0\x34\x3b\xf3\xfb\xcf\x4d\x3b\x81\x2b\x34\xe8\x64\xc0\x12\x96\x5b\xb8\x09\x81\x8e\xa0\x24\x30\x14\x00\x4b\x15\x7e\x48\x26\xc9\x04\xe6\x8d\xf2\xc0\xbf\xd0\x20\xdc\xcb\xda\x49\x13\x2a\xa5\x11\xea\x8f\xb1\x50\x91\xeb\xbd\x4a\x5c\xa1\x26\xdb\xa2\x09\x40\x15\x23\x42\x44\x48\x6b\xb5\x2a\x64\x50\x64\x72\x8f\x6e\xa5\x0a\x14\x70\x1d\xc0\x37\xd4\xe9\xb2\x17\x5d\x22\x34\xd2\x94\x59\x14\xc7\x52\xc0\x9c\xa0\xa5\x52\x55\xdb\x88\x65\xce\x9e\xfc\x11\x74\x1e\x7b\xb5\x0b\x6b\xa3\x41\x24\xc9\x78\x2c\x0a\x32\x95\xaa\x3b\x87\x3f\xa7\xc7\xe9\x2f\xb1\xa2\x97\xc1\xf4\x92\x00\x0c\x4f\x62\xd5\x8a\x25\x6d\xe0\x1c\xd2\x46\xfa\x46\x15\xe4\x6c\x6e\x1d\x16\xca\xe3\xe9\x49\x9a\xb0\xe3\x04\xee\x30\x74\x16\x24\xf8\xad\x29\xb8\xcc\x8a\x74\x89\x0e\x2a\x47\x2d\x50\xe7\x60\x4d\xee\x51\x99\x1a\x4a\xc5\x71\x81\x1c\x67\x49\x90\xaf\x86\x24\x0e\x94\x06\xc0\xc3\x08\x48\x9f\x9f\xc1\xca\xd0\x88\x1d\xe0\xf5\x35\x3d\x82\x74\x17\x39\x8a\x5f\x1b\x1f\xa4\xd6\x70\x45\xb0\xec\x14\x37\x08\xcd\x4a\x39\x32\xb1\xab\x07\x70\xeb\x68\xa5\x3c\x77\x15\x52\xdf\xa0\xd6\xcc\x52\x46\x2b\x83\x53\xf8\xe2\x0b\xa7\x6c\x78\xa8\x49\x4b\x53\x0f\xdc\xdf\xe5\x23\x82\xe2\xb6\x13\x77\x4f\x06\x58\x8c\xb2\xe0\x7d\xb3\x80\x9a\xd0\x8f\x05\xe9\xbe\x9e\xd8\x61\x6e\x4e\x34\x44\xfb\x77\x2a\xb3\x1b\xc0\x8f\x7f\xfc\x8d\x45\x43\x90\x16\xe5\x5b\x5b\x52\x98\xcd\x20\x6f\xa8\xc5\x9d\x25\x17\x4b\x1e\x80\x2b\xfe\x49\xd0\x94\x49\x72\x98\x32\xcf\xe7\xec\xec\xee\xf2\xcf\xeb\xdb\x79\xe2\x31\x40\x86\x49\x32\x30\x7f\xa5\xb5\xd1\x24\xcb\xd8\xbf\x2b\x12\x42\xa4\xc9\xba\x8e\x1e\xff\x42\x76\xf3\x41\xa1\x26\x11\xa4\x13\xf5\x13\x34\x21\x58\x3f\xcd\x73\xcf\xe3\x92\x35\x8a\x9a\xa8\xd6\x28\xad\xf2\xbc\x34\x6d\x3e\x88\xf2\xdf\x37\x71\x22\x8e\x05\x97\xd2\x6d\x32\xd9\x96\xa7\x27\x23\x60\xa7\xfe\x97\xe1\x77\xb7\xa7\xed\x3b\xde\x32\xb6\x41\x76\x09\x79\xe7\x5d\xae\xa9\x90\x1a\xb2\xcd\x53\xf5\x59\x32\x3b\x16\x0f\xa4\x07\xdd\xdc\x5e\xcc\x7f\x7b\x87\xb5\x8f\xdc\x6e\xc8\x2c\xe4\x64\x63\x54\xdc\x98\xe1\x84\xa3\xd6\x06\x16\xeb\x86\x64\xab\x16\xd3\xdd\xc3\x81\xe3\xc8\xe6\x0d\x0e\x11\xce\x6b\xfc\x46\xef\x4f\xbe\xe2\xc6\x92\x0b\xbd\xf5\x7c\x2f\x30\x5f\x2a\x33\x7d\x2f\x80\xad\xbd\xe5\x4b\xf4\xfb\xfa\xe9\xec\x0e\x99\x43\x25\xfb\xd4\xff\x89\x1c\x13\x1d\xb7\x3d\xe6\x7a\x7f\x79\xe7\xfb\xab\xa4\x26\xbe\x63\xc2\x7b\x47\xa4\x0d\x59\x1c\x71\x67\x4b\xbe\x76\x20\xdb\xc2\x2c\xe7\xab\x26\x37\x1d\x7f\x26\xc7\xb3\x9f\xbe\x1d\xba\xa9\xf1\x03\x62\xbf\x9a\xf7\x7d\xf9\xe4\xa0\x45\x57\x74\x4e\x49\x9d\x8c\x3b\xf5\x5f\x00\x00\x00\xff\xff\xd6\x6f\x4d\xca\x02\x05\x00\x00"

func dataCommonDevVagrantfileTplBytes() ([]byte, error) {
	return bindataRead(
		_dataCommonDevVagrantfileTpl,
		"data/common/dev/Vagrantfile.tpl",
	)
}

func dataCommonDevVagrantfileTpl() (*asset, error) {
	bytes, err := dataCommonDevVagrantfileTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/common/dev/Vagrantfile.tpl", size: 1282, mode: os.FileMode(420), modTime: time.Unix(1439846905, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _dataCommonDevDepBuildVagrantfileFragmentTpl = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x91\xbd\x6a\xeb\x40\x10\x85\xfb\x7d\x8a\x83\x7c\xcb\x6b\x6d\x6f\x7c\x2b\x17\xb7\x4c\x8a\x34\xa9\x82\xbc\x1a\x47\x03\xfb\xc7\xee\x48\x41\x18\xbf\x7b\xc6\x22\xc1\x31\xb1\x21\x45\x4a\x69\x67\xbe\xf3\x0d\x67\x85\xff\x14\xa9\x74\x42\x3d\xf6\x33\x7a\xca\x14\x7b\x8a\x6e\xc6\x21\x15\xfd\x9c\xc8\xa7\x1c\x28\xca\x06\xc7\x23\x62\x17\x08\xa7\x93\x31\x2b\xec\x52\xd6\x19\xf6\x54\xc1\x51\x12\x3a\x08\x85\x8c\x9e\x0b\x39\x49\x65\x6e\xf1\x34\x10\xaa\x2b\x9c\x05\x6f\xec\x3d\x42\x9a\x08\x32\x50\x68\x8d\x4b\xf1\xc0\xaf\xed\x14\xda\x5c\xd2\xc4\x95\x53\x44\x73\x86\x35\x7f\x0d\x50\xd3\x58\x1c\x6d\xd0\x68\x62\xee\x64\x68\x5d\x0a\x59\x1f\x7b\x8d\xb6\xaa\xb4\x56\x4b\xbb\x1f\xd9\xf7\x36\x8d\x92\x47\x59\xb6\x7a\xaa\xc2\xb1\x13\x65\xe9\xaa\x95\x90\x75\x36\xaf\x2f\xd6\x8d\xf9\xbd\xdc\x31\x57\xe9\x8a\xb4\x67\xe0\xcf\xd2\xdb\xab\x95\x3b\x2a\x75\x20\xef\x17\x1e\x47\xcf\x51\x5d\xfe\x5c\x08\x2f\x95\x64\xcc\xc6\x7c\xfb\x85\x7f\xd8\x6e\x77\x0f\x8f\xcf\x9f\xb5\xe8\x1d\xd8\xab\x4d\x99\x4d\x98\x70\x43\x06\x76\xac\xc5\xfa\xe4\x3a\x6f\x75\xd0\xde\xa8\x56\x8b\xc2\x87\xf2\x52\xf3\x1d\xd2\xd5\x59\xb0\x24\xce\x72\x64\xf9\x42\x5c\x5e\xcc\xa2\xf7\x1e\x00\x00\xff\xff\x50\x3a\x96\x31\x6d\x02\x00\x00"

func dataCommonDevDepBuildVagrantfileFragmentTplBytes() ([]byte, error) {
	return bindataRead(
		_dataCommonDevDepBuildVagrantfileFragmentTpl,
		"data/common/dev-dep/build/Vagrantfile.fragment.tpl",
	)
}

func dataCommonDevDepBuildVagrantfileFragmentTpl() (*asset, error) {
	bytes, err := dataCommonDevDepBuildVagrantfileFragmentTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/common/dev-dep/build/Vagrantfile.fragment.tpl", size: 621, mode: os.FileMode(420), modTime: time.Unix(1440179251, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _dataCommonDevDepBuildVagrantfileTpl = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x54\x5f\x4f\xfb\x46\x10\x7c\xf7\xa7\xd8\x3a\xa8\xb4\x12\x3e\x0b\x84\x78\x88\x20\x12\x02\x29\xe5\xa1\x02\x95\x94\x3e\xc2\xc5\xde\xd8\x27\xec\xdb\xeb\xdd\x39\x7f\x80\x7c\xf7\xee\x9d\x1d\x20\x48\xf4\xf7\x93\x22\xc5\x5e\xef\xce\xcc\xde\x8c\x3d\x82\x29\x6a\xb4\xd2\x63\x09\xf3\x0d\xdc\x7a\x4f\x47\x50\x12\x68\xf2\x80\xa5\xf2\xbf\x24\xa3\x64\x04\xb3\x5a\x39\xe0\x9f\xaf\x11\x1e\x64\x65\xa5\xf6\x0b\xd5\x20\x54\x5f\x67\x61\x41\x16\xe6\x9d\x6a\x4a\xa5\xab\xd0\xce\xc3\x73\xa5\xa5\xdd\xf0\x8d\xf4\x01\xa3\x73\xdc\x2e\x1d\x48\x28\xd1\xa0\x2e\x51\x17\x9b\x38\x56\xe2\x12\x1b\x32\x2d\x6a\x2f\x22\xeb\x75\x2f\xa3\x96\xba\xcc\x82\x16\x86\xe0\xf9\x40\x2c\x60\x46\xd0\x52\xa9\x16\x9b\x58\x3c\x0a\xa8\x51\xdd\xa5\x31\xb1\x21\x49\x06\x9d\xa2\x20\xbd\x50\x55\x67\xf1\xb7\xf4\x24\xfd\x3d\xec\xf6\xd6\x97\xde\x12\x80\xfe\x4a\x2c\x5b\x31\xa7\x35\x5c\x40\x5a\x4b\x57\xab\x82\xac\xc9\x8d\xc5\x42\x39\x3c\x3b\x4d\x13\x6e\x1c\xc1\x3d\xfa\xce\xb0\x6a\xb7\xd1\x05\x6f\xb0\xa0\xa6\x44\x0b\x0b\x4b\x2d\x50\x67\x61\x45\xf6\x39\xec\x5c\x2a\x9e\xf3\x14\x16\x26\xc8\x97\xbd\x88\x3d\xa6\x1e\xe0\x71\x00\x48\x5f\x5f\xc1\x48\x5f\x8b\x1d\xc0\x76\x9b\x1e\x41\xba\x9b\xfc\x31\xf9\xaa\x46\x8b\x51\x42\x41\xad\xe1\xdd\x4b\x28\xa5\x97\xd1\x2e\x8a\xc3\x39\xb1\x33\x02\xfe\xc1\xb0\x7c\x3c\x43\x96\x26\x8b\x02\x5d\xef\x68\xf4\x0b\x5c\x61\x95\xe1\x93\xff\x09\xa9\xef\x44\xdb\x6d\xce\xae\x65\x6c\x64\x1e\x41\xa2\xf2\xc0\x36\xc8\xbe\xd1\xce\xcb\xa6\x81\x29\x0d\x24\xa8\x97\xca\x92\x0e\x1e\xef\x11\x19\x4b\x4b\xe5\x14\x69\x48\x5d\x8d\x4d\xc3\x40\x4a\x37\x4a\xe3\x18\x0e\x7a\x61\x8f\x15\x35\x52\x57\x09\x27\x26\x49\xf6\x6b\xec\xdb\xf9\xf9\xfd\xd5\x5f\x37\x77\xb3\xc4\xa1\x87\x0c\x93\x04\x8b\x9a\x20\xbd\xa6\x95\x6e\x48\xc6\x30\x4e\x49\x08\x91\x26\xab\x2a\x74\xfc\x0b\xd9\x2d\xe4\x35\xb5\xb8\x3b\xe8\xbc\x22\xe1\xa5\x15\xd5\x0b\xd4\xde\x1b\x37\xce\x73\xc7\x36\xca\x0a\x45\x45\x54\x35\x28\x8d\x72\x61\xf3\xbc\x27\xe5\xbf\x63\x71\x2a\x4e\x04\xab\xec\xd6\x99\x6c\xcb\xb3\xd3\x01\x60\xc7\xfe\xb7\xe6\x7b\xfb\x89\xdb\x75\xc1\x00\x69\x21\xbb\x82\xbc\x73\x36\x6f\xa8\x90\x0d\x64\xeb\x97\xc5\x77\x62\x12\xdb\x7e\xfb\x68\xa0\xf9\x53\xc6\xdc\x4c\x6f\xef\x2e\x67\x7f\x7c\xf0\xb4\xcf\x1c\x45\xc8\x0c\xdb\x6f\xc2\x54\x70\xae\x7f\xc2\x53\x2b\x0d\x4f\xab\x9a\x64\xab\x9e\xc6\xbb\x8b\xbd\xc6\x01\x9b\x73\xe7\x03\x38\x87\xef\x1d\x3d\x3e\x39\xc4\xb5\x21\xeb\x63\xf5\xe2\xd3\x60\xce\x6f\xfb\xf8\x63\x37\xae\xc6\xca\x41\xe8\x3b\x84\xc9\xe4\xcb\x32\x62\xce\xef\x9b\x2d\xf6\x31\xfb\x4d\x3e\xa3\xfe\xcf\xe4\x20\x74\x48\x5a\xd0\xfa\x70\x75\xef\xe2\x07\xa5\x22\xfe\x40\xf9\x8f\x13\x91\xc6\x67\xc1\xfd\xce\xf0\x0b\x82\x90\x6d\x60\x12\xe2\x9b\xeb\x8e\x23\x7a\x32\xf9\xf5\x78\xbf\x4d\x0d\xe1\xe5\xbe\x8a\x3f\x3e\xf3\x17\x0b\x2d\xda\xa2\xb3\x4a\x36\xc9\x10\xb7\xff\x02\x00\x00\xff\xff\x17\x58\x96\x17\x3f\x05\x00\x00"

func dataCommonDevDepBuildVagrantfileTplBytes() ([]byte, error) {
	return bindataRead(
		_dataCommonDevDepBuildVagrantfileTpl,
		"data/common/dev-dep/build/Vagrantfile.tpl",
	)
}

func dataCommonDevDepBuildVagrantfileTpl() (*asset, error) {
	bytes, err := dataCommonDevDepBuildVagrantfileTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/common/dev-dep/build/Vagrantfile.tpl", size: 1343, mode: os.FileMode(420), modTime: time.Unix(1440118643, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _dataCommonDevDepBuildBuildShTpl = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x4c\x91\xc1\x6e\xdb\x30\x10\x44\xef\xfa\x8a\xa9\x7d\xc8\xc9\x12\xfa\x01\x2d\xd0\x5e\x72\x6c\x81\xb4\x1f\xb0\xa2\x56\x12\x1b\x99\x4b\xac\x96\x51\xfc\xf7\x5d\x4a\x0e\x62\xc0\x30\x04\x2e\x67\xe6\xed\xf0\xfc\xa5\xeb\x63\xea\x7a\x5a\xe7\xe6\xdc\x9c\xf1\xa3\x98\x5c\x26\x4e\xac\x64\x3c\xa0\xbf\xe1\x97\x99\xb4\xfb\xec\xcf\x1c\x57\xf8\xcf\x66\x46\x5f\xe2\x32\x60\x0d\x1a\xb3\x61\x14\x05\xe1\x59\x2e\x6e\xe3\xa2\xac\xf2\x8f\x83\xb5\xcd\xca\x86\x0b\x37\x2e\x7d\xf1\x2f\xab\x72\x13\x8c\xf4\xca\x87\x87\x87\x6a\x80\x24\xac\x72\x65\xe4\x85\xcc\x9d\xae\x35\x80\x0c\x1b\x3f\x29\x23\x26\x73\x94\x60\xf1\x8d\x1d\x02\x7f\xfb\x92\xac\xf8\x29\x32\xa9\xc5\x50\x16\xd2\x0a\x39\xf0\x48\x65\x71\x91\xa4\x27\xc3\x22\x34\x3c\x26\xc4\xf1\x08\x8f\xab\x4f\xdd\xc5\xb9\xda\x86\xdf\xb3\xa8\xe1\xf7\xcb\xd7\x6f\xa7\xef\x38\xed\x94\x52\x34\x30\xfc\xbf\xee\x30\xc6\x85\x9d\xcc\x41\x30\x39\xfe\x7e\x4a\x36\xd7\x51\x66\x5d\x6e\xd5\xa6\xe4\xa6\x45\x37\x3b\x7e\xf7\x46\x93\x52\xb2\xae\x3d\x42\xab\xdf\xb3\x54\x7e\xd9\xa5\x9b\xe8\x6b\x4c\x13\x86\xa8\xde\x8d\xe8\xad\x09\x03\x3e\x44\xf5\xf2\xcf\xbd\xd1\x4a\x7d\xef\x0f\x94\x06\x6c\x1a\xed\x68\x4b\x8a\xe5\x62\x9f\x86\xeb\x4c\xea\x65\x7f\xfa\x9d\xb1\x45\xc7\xab\x77\x83\x5c\xb3\xd3\x3f\x4c\xeb\x22\xf7\x5a\x11\x28\x81\x69\x8d\xbe\x02\xbf\x5b\x6d\x17\xd1\x0b\x99\xe4\xfe\xaa\x17\xc1\xa9\x13\x7f\xf6\xee\xc8\x3c\x35\xff\x03\x00\x00\xff\xff\xf2\xf1\x39\xcc\x26\x02\x00\x00"

func dataCommonDevDepBuildBuildShTplBytes() ([]byte, error) {
	return bindataRead(
		_dataCommonDevDepBuildBuildShTpl,
		"data/common/dev-dep/build/build.sh.tpl",
	)
}

func dataCommonDevDepBuildBuildShTpl() (*asset, error) {
	bytes, err := dataCommonDevDepBuildBuildShTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/common/dev-dep/build/build.sh.tpl", size: 550, mode: os.FileMode(493), modTime: time.Unix(1440178707, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _dataCommonDevDepBuildUpstartConfTpl = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x54\xcf\x31\x4f\x03\x31\x0c\x05\xe0\xdd\xbf\xe2\xa9\x48\x30\x95\xd0\xc2\x31\xde\xca\xc8\xc8\x80\x3a\xa4\x17\x17\x45\x4a\xed\x28\x76\x5b\x50\xd5\xff\xce\x11\x81\x54\xa6\x28\xef\x49\x9f\xed\xc4\x36\xb5\x5c\x3d\xab\x60\x71\x3e\x43\xe2\x9e\x71\xb9\x60\x89\x17\x16\x6e\xd1\x39\x61\xfb\x85\x57\x77\x5d\x10\x35\xb6\x1a\x4f\xf2\xf7\xa2\xe4\x7d\x76\xac\x06\x0c\x44\xe6\xb1\x39\x66\xa6\x1d\xa4\xf0\x91\x0b\xde\xd7\x8f\x4f\xc3\x66\x2e\xb4\xfe\xcf\x1f\x9e\x37\x44\x37\x78\x63\x24\x95\x3b\xc7\x29\x8a\xc3\x15\x33\xdb\x11\x57\xc5\x2e\x9a\x63\xa7\x0d\x89\xab\x51\x55\xf3\x65\x87\xf8\x93\x27\x58\x61\xae\x7d\x68\x5f\x9e\x80\x70\xb0\x16\x8a\x4e\xb1\x84\x6d\x96\x70\x75\xc9\x38\x86\x63\xfc\xe9\x3e\xae\xd2\xfb\xf9\x8b\xf5\x78\xbb\x22\x96\x84\x5f\xe5\x3b\x00\x00\xff\xff\x9a\xf6\x9d\x7c\x0c\x01\x00\x00"

func dataCommonDevDepBuildUpstartConfTplBytes() ([]byte, error) {
	return bindataRead(
		_dataCommonDevDepBuildUpstartConfTpl,
		"data/common/dev-dep/build/upstart.conf.tpl",
	)
}

func dataCommonDevDepBuildUpstartConfTpl() (*asset, error) {
	bytes, err := dataCommonDevDepBuildUpstartConfTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/common/dev-dep/build/upstart.conf.tpl", size: 268, mode: os.FileMode(420), modTime: time.Unix(1440179369, 0)}
	a := &asset{bytes: bytes, info:  info}
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
	if (err != nil) {
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
	"data/aws-vpc-public-private/build/template.json": dataAwsVpcPublicPrivateBuildTemplateJson,
	"data/common/dev/Vagrantfile.tpl": dataCommonDevVagrantfileTpl,
	"data/common/dev-dep/build/Vagrantfile.fragment.tpl": dataCommonDevDepBuildVagrantfileFragmentTpl,
	"data/common/dev-dep/build/Vagrantfile.tpl": dataCommonDevDepBuildVagrantfileTpl,
	"data/common/dev-dep/build/build.sh.tpl": dataCommonDevDepBuildBuildShTpl,
	"data/common/dev-dep/build/upstart.conf.tpl": dataCommonDevDepBuildUpstartConfTpl,
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
	Func func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"data": &bintree{nil, map[string]*bintree{
		"aws-vpc-public-private": &bintree{nil, map[string]*bintree{
			"build": &bintree{nil, map[string]*bintree{
				"template.json": &bintree{dataAwsVpcPublicPrivateBuildTemplateJson, map[string]*bintree{
				}},
			}},
		}},
		"common": &bintree{nil, map[string]*bintree{
			"dev": &bintree{nil, map[string]*bintree{
				"Vagrantfile.tpl": &bintree{dataCommonDevVagrantfileTpl, map[string]*bintree{
				}},
			}},
			"dev-dep": &bintree{nil, map[string]*bintree{
				"build": &bintree{nil, map[string]*bintree{
					"Vagrantfile.fragment.tpl": &bintree{dataCommonDevDepBuildVagrantfileFragmentTpl, map[string]*bintree{
					}},
					"Vagrantfile.tpl": &bintree{dataCommonDevDepBuildVagrantfileTpl, map[string]*bintree{
					}},
					"build.sh.tpl": &bintree{dataCommonDevDepBuildBuildShTpl, map[string]*bintree{
					}},
					"upstart.conf.tpl": &bintree{dataCommonDevDepBuildUpstartConfTpl, map[string]*bintree{
					}},
				}},
			}},
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
        err = os.MkdirAll(_filePath(dir, path.Dir(name)), os.FileMode(0755))
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
                err = RestoreAssets(dir, path.Join(name, child))
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

