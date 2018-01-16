// Code generated by go-bindata.
// sources:
// templates/cf_lb.tf
// templates/concourse_lb.tf
// templates/network.tf
// templates/network_security_group.tf
// templates/output.tf
// templates/resource_group.tf
// templates/storage.tf
// templates/tls.tf
// templates/vars.tf
// DO NOT EDIT!

package azure

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

var _templatesCf_lbTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xe4\x57\x3d\x8f\xe3\x36\x10\xed\xfd\x2b\x08\x21\x45\x12\x9c\xbc\x9f\x38\x5c\xa3\x22\x41\x8a\xa4\x0b\x70\xe9\x09\x8a\x1c\xd9\xc4\xd2\x24\x8f\x1c\x79\xcf\x39\xf8\xbf\x07\x24\x25\xdb\x92\x28\x5b\xbe\x54\x41\xb6\x5b\xf3\xcd\x87\xde\xbc\x19\x0e\xf7\xcc\x49\x56\x2b\x20\x85\x3f\x78\x84\x1d\x15\x66\xc7\xa4\x2e\xc8\xb7\xe3\x6a\x75\x3e\xb4\xcd\x57\xca\xc1\x21\xad\x99\x87\x8f\xaf\xb9\x63\xcb\xbc\x7f\x37\x4e\xa4\x33\x07\xde\xb4\x8e\x03\x29\xd8\xdf\xad\x03\xb7\xa3\xbe\xad\x35\x60\x41\x0a\xde\x94\x3e\x04\x58\x11\xa2\xd9\x0e\xc8\xf8\xaf\x22\xc5\x0f\xdf\xf6\xcc\xad\x41\xef\xa9\x14\xc7\x32\x19\xac\x08\x61\x42\x38\xf0\x9e\x5a\x07\x8d\xfc\x7a\x09\xe7\x52\xb8\x14\xe0\xc7\x60\xa9\x01\xdf\x8d\x7b\xa3\xe1\xe7\x0f\xe4\xd3\x07\xf2\xf4\xd3\x31\x38\xe8\xb3\xa2\x1b\x67\x5a\x4b\x53\xf8\xe8\xa0\xcf\x72\x88\x58\xd7\xc6\x6f\xd7\x01\x16\xcd\xf7\xd2\x61\xcb\x14\xed\xdd\x47\xfb\x81\xf9\x08\x31\xb0\xcf\xb2\xd2\xbb\xf2\xc0\x5b\x27\xf1\x90\xe2\x46\x96\xe6\x29\xca\x30\x14\xd2\x53\x86\x33\x94\x46\x67\xa1\x0e\x36\xd2\xe8\x59\x16\x96\x92\xb0\x22\x04\xd9\xc6\xc7\xd4\x08\x01\xbd\x97\xce\xe8\x1d\x68\x9c\x24\x15\x22\x1d\x17\x7e\xb4\x6b\x15\x24\x65\x6c\x11\xad\xbf\x22\x8e\x39\x06\x84\xf6\x21\xa0\x75\xd2\x04\x8f\x79\x9b\xe7\xc7\xa7\x15\x21\x42\x3a\xe0\x63\x9e\xce\x7e\xff\xd0\xb5\x69\xb5\x88\x7a\xe3\x1c\xbc\x9f\xcd\xe0\x17\xa5\xcc\x7b\x8a\x6a\xd0\x70\xa3\x66\x70\x7f\x71\x1b\x50\x1d\xa7\xd6\x38\xa4\x8e\xe9\x0d\x0c\x51\x3f\x07\x8c\x00\x8f\x52\xc7\x2a\x4e\x80\x15\x29\x5e\x5f\x5f\x2e\x3c\xcd\xb5\xc3\xc4\xd3\x18\xd8\x63\xb2\xed\x70\xc9\xf0\xa2\xae\xc8\x4b\x38\x23\xab\x3c\x70\xcd\x9b\xfb\x3a\xe4\x2c\x16\x65\x36\xdf\x21\x95\xce\x70\x81\x5a\x9e\xff\xeb\x6a\xf9\x1f\xca\xc5\xb6\xb5\x92\x9c\xca\x30\x43\x55\x7d\x5b\x1e\x79\x7d\xd4\xa5\xb4\x73\x23\x75\x6a\x79\x63\xb6\x7e\x07\x49\xa7\xaf\x38\x15\x83\xa9\x53\x2e\x15\x29\xc4\x41\xb3\x9d\xe4\x33\x1c\x30\x6b\x95\x4c\x60\xba\x61\x08\xef\xec\x50\x90\xa2\xe3\xf3\x8e\x6b\x85\x59\x5b\xf6\xf6\xff\xf2\xde\x58\x7a\x3b\x05\xb9\xbe\xb5\xdd\xf5\x72\x4a\xb2\x22\xc5\x67\x64\x5a\x30\x27\xe8\xe7\x1d\x53\xaa\x88\xe7\x28\xc1\x8d\xcf\xd3\x09\x67\x96\xf1\xd0\xd9\x15\x79\x8e\xf7\x50\xea\xbb\x1a\xc6\x9e\x87\xc9\xfc\x19\x20\x8f\x4f\xc9\x47\xae\x4f\x2b\x52\xfc\x8e\x68\x3b\x00\xc3\x6d\xc6\xc9\x83\x32\x1b\xa9\x13\x64\x6b\x3c\x66\x20\x11\xb1\x4e\x9f\x3e\xd8\xbb\x8e\xc9\x4c\x6a\x04\xb7\x67\xa3\xd0\x1f\x1f\xbb\xaf\xde\x81\x69\x91\x64\x0f\x5b\xbd\x05\xa6\x70\x7b\xa0\xb8\x75\xe0\xb7\x46\x09\x52\x91\x97\x9e\x83\xae\x9a\x41\x58\xdc\xe8\x46\x6e\x5a\x97\x8a\x32\xa6\x25\xd7\x15\x9d\x71\x29\x6d\x39\x30\x4e\x39\xa7\xf5\x8b\x4a\xb1\x60\x23\x92\xe2\xf8\x90\xf0\xfe\xe1\x0c\x4d\xbf\xac\xe3\xba\x77\xd6\x4d\xcc\xbb\x71\x46\x23\x68\x11\xc7\xdc\x65\xb2\x15\x29\xfa\xb3\x70\x94\xd6\x87\x54\x9d\x80\xac\xc8\xeb\xeb\xcb\xbd\x4e\x94\xd9\x8c\x7d\x64\x9c\xdc\xa4\xf0\x5a\x67\xf1\xa6\xec\x1d\xcd\xd0\x39\x9d\x00\x63\x66\x4f\x88\xb5\xaa\xd7\xa7\x85\x6b\x45\x48\xcd\xf8\x5b\xc8\xf0\x34\xc7\x8d\x51\xa3\xcf\x9d\x64\xd3\xd9\x94\x9d\x4d\x19\x6c\x26\x0e\x03\xbb\xd4\x03\xa2\xd4\xa7\x15\x30\x3f\x5b\x17\x2e\xc5\x65\x0d\xe5\x16\x3d\x76\x4d\x6b\xcc\x9b\x84\xf8\xc4\x10\x94\x35\x8d\xd4\xa9\x83\x8b\xdf\xa4\x0f\xef\x0c\x71\x51\x94\x4c\xc4\x4f\x8f\xb3\x6d\x3b\x6a\x5c\x07\x5f\x5a\xf0\x48\x87\x9d\x54\x91\xa7\xde\x41\x0d\x74\xfc\x5d\xc3\xe9\x10\x69\xf1\x5e\xc5\x47\x91\x6c\xc2\xb0\x9d\x8c\x96\x8a\x14\xde\xab\x32\x20\x52\x58\xc1\x90\x0d\xe5\x30\x7a\x56\x1d\xfb\xb9\x92\x5e\x52\x43\x5c\xff\xeb\xb9\xce\xb1\x1c\x4a\x7a\x04\x0d\xee\x6a\x39\xee\xae\x4b\x70\xad\x3c\x76\x5a\x9c\xd5\x3c\x9d\xd5\xd3\x0d\x75\x0f\x5a\x71\xc2\xf5\xb5\xae\xbe\xb2\x3c\x9d\xcb\xdc\x81\x47\x05\x1a\xc7\x19\x15\x28\x72\xda\x4b\xc3\x99\x36\xa8\x3c\x2e\x9b\xb7\xa9\x5d\x4a\xab\xfb\xd2\x6b\x21\xf8\xa5\x78\xb0\xf9\xde\xf9\x95\xf9\x70\xb9\x87\xff\x06\x45\xbe\xb2\x73\x2d\x2e\x66\x6e\x3a\x5c\xbe\x84\x17\x0e\x86\x99\xa9\x70\xd7\x9b\xf8\xb2\xfd\xe3\x53\xd1\xb4\x68\x5b\x24\x45\x76\x81\x09\x55\xd8\x33\xd5\x8e\xdc\x67\xb0\xfd\xfb\xff\xbc\x28\xfe\x13\x00\x00\xff\xff\xaf\xda\x43\xab\xe7\x10\x00\x00")

func templatesCf_lbTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesCf_lbTf,
		"templates/cf_lb.tf",
	)
}

func templatesCf_lbTf() (*asset, error) {
	bytes, err := templatesCf_lbTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/cf_lb.tf", size: 4327, mode: os.FileMode(480), modTime: time.Unix(1516127023, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesConcourse_lbTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x92\xcf\x6e\xf3\x20\x10\xc4\xef\x7e\x8a\x15\xfa\x0e\x5f\xa5\xd8\x52\x6f\xbd\xe4\x59\x10\x06\x92\xa2\x92\x5d\xb4\x80\xfb\x27\xf2\xbb\x57\xd8\xb2\x1b\x92\x3a\x4a\x2d\xcb\x07\x6b\x67\x98\xf9\xb1\x6c\x23\x65\xd6\x16\x84\xfa\xca\x6c\xf9\x24\x63\xee\xd1\x26\x01\x42\x13\x6a\xca\x1c\x6d\x1b\x51\xc0\xb9\x01\x40\x75\xb2\x70\xfd\xec\x41\xfc\x3b\x0f\x8a\x3b\x8b\x83\x74\x66\x6c\x2b\x5d\x03\xa0\x8c\x61\x1b\xa3\x0c\x6c\x0f\xee\xe3\x52\xa5\x9d\xe1\xf9\xb8\xff\xc5\x00\x6d\x7a\x27\x7e\x93\xe5\xf7\x0e\x5e\x76\xf0\xfc\x34\x16\x83\x25\xa3\x3c\x32\xe5\x20\xe7\x14\x93\xc1\x92\xb9\x9e\xe8\x7a\x8a\xaf\x5d\x19\x9b\xe4\x83\xe3\x94\x95\x97\x8b\xfd\xa4\xaf\xe4\x57\x13\x95\x7e\x6c\x9a\x5b\x46\x21\xf7\xde\x69\xe9\x82\x00\x51\x3e\x9b\x70\x1e\x80\xe4\xfb\xd6\x85\x92\xd3\x93\x56\xc9\x11\xde\x37\x60\x7b\x74\x84\xdb\x5c\x2a\xc1\x43\x7c\xd6\x32\x72\xb9\x29\xe5\xd7\x2c\x7b\x10\xe6\x13\xd5\xc9\xe9\x0d\x14\xbe\xbf\x5c\x95\x6d\x14\x77\x09\x6c\xb5\x79\xbc\xc4\x6f\xf0\x6e\x99\x35\x00\x07\x26\x4c\x16\x4d\xe9\xab\x09\x0f\xee\x98\x79\x56\x96\xe4\x7f\xdf\xf1\xc5\xaf\x75\xa1\xad\xfc\xc4\x64\x77\x0b\xd7\x99\xba\xd6\x3a\xd1\x95\xd7\x4c\x6d\xc6\xc2\x9a\x72\x0a\x39\x81\xf0\xa4\x8c\xec\x95\x57\xa8\x2d\xcf\x80\x07\xe5\xf3\x15\x1d\xdf\x77\x6b\xa8\x9f\xdd\xfd\x0e\x00\x00\xff\xff\x7b\xdb\x7a\xd6\xdf\x03\x00\x00")

func templatesConcourse_lbTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesConcourse_lbTf,
		"templates/concourse_lb.tf",
	)
}

func templatesConcourse_lbTf() (*asset, error) {
	bytes, err := templatesConcourse_lbTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/concourse_lb.tf", size: 991, mode: os.FileMode(480), modTime: time.Unix(1516126949, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesNetworkTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x91\xcf\x4a\xc4\x30\x10\xc6\xef\x79\x8a\x61\xf0\xa0\xb0\x5b\x3c\x7a\xf1\x49\x44\x42\x36\x19\xd7\xe0\x36\x29\x93\x3f\x8a\x25\xef\x2e\xa9\xad\xd8\xd8\xe2\xe6\xfc\x9b\x99\xef\xf7\x85\x29\xf8\xc4\x9a\x00\xd5\x67\x62\xe2\x5e\x66\xcb\x31\xa9\x8b\x74\x14\xdf\x3d\xbf\x21\xe0\xc9\x87\x57\x84\x51\x00\x38\xd5\x13\x34\xef\x11\xf0\x66\xcc\x8a\x3b\x72\x59\x5a\x53\x8e\x15\x3f\x66\x87\x02\x40\x19\xc3\x14\x82\x0c\x83\xd2\xf4\xc3\x3f\xcd\x03\xf3\x05\xa9\xad\xe1\x82\xcf\x02\xe0\xe2\xb5\x8a\xd6\xbb\xcd\xfd\x4c\x67\xeb\x5d\xa9\x7b\x97\xd4\xf2\xcc\x3e\x0d\x72\x8a\x35\x71\x8b\xc4\x1a\xe8\x6a\xa4\xae\x52\x05\x45\x11\xe2\xaf\x74\x48\x27\x47\xf1\x5f\xd7\x1d\xd9\xb0\x92\x1d\x98\x5e\xec\xc7\xef\x81\x2a\xf8\x7d\xe1\xb6\xf5\x3e\xc0\xc3\x01\xee\xef\x76\xad\xae\xd6\x02\x68\x3e\x6e\xa3\x95\x86\x68\x6a\xf9\x0a\x00\x00\xff\xff\xdb\x9a\xf5\x1a\x0b\x02\x00\x00")

func templatesNetworkTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesNetworkTf,
		"templates/network.tf",
	)
}

func templatesNetworkTf() (*asset, error) {
	bytes, err := templatesNetworkTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/network.tf", size: 523, mode: os.FileMode(480), modTime: time.Unix(1513729670, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesNetwork_security_groupTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x94\xcf\x8e\x9b\x30\x10\x87\xef\x7e\x8a\x91\xd5\xd3\x4a\x1b\x6d\xd9\xb0\xe2\x92\x43\x8f\xbd\xf7\x8e\x1c\x7b\x4a\xac\x12\x0f\x1a\x9b\xa4\x6d\xc4\xbb\x57\x86\x92\x86\x84\x44\x14\x55\xaa\xb2\x0a\xe7\x6f\x7e\xf3\xc7\x9f\x60\xf4\x54\xb3\x46\x90\xea\x67\xcd\xc8\xdb\xdc\x61\xd8\x13\x7f\xcb\x3d\xea\x9a\x6d\xf8\x91\x17\x4c\x75\x25\x41\xae\xc9\x6f\x24\x1c\x04\x80\x53\x5b\x84\xb3\x6f\x05\xf2\xc3\x61\xa7\x78\x81\x6e\x97\x5b\xd3\x3c\xb7\xb8\x00\x28\x49\xab\x60\xc9\x8d\xc2\x8c\x85\x25\xd7\x44\xae\x9f\xa4\xeb\x97\xb7\x3d\x5a\xae\x1f\x6c\x08\x2c\x62\xfe\x22\x52\x8d\x14\x02\x20\xa8\xc2\xb7\xc3\x01\xa0\xdb\x59\x26\xb7\x45\x17\x2e\xc6\x8a\x9d\x1a\xd1\x08\x31\x61\x71\xae\x4b\x94\x20\xfd\xad\xb5\xaf\xae\xef\xbb\xed\x2b\xb6\x14\xc3\xc6\x6b\x92\x97\x17\x01\x60\x2c\xa3\x3e\x3f\xd1\x9f\xdc\xcf\x6e\x4d\xb5\x33\x31\x4d\x69\x8d\xde\x5f\x9d\xe0\x53\x59\xd2\xbe\xeb\x4a\x81\x34\x95\x57\xb8\x2f\xba\x8a\xd4\xef\x73\x56\xc4\x21\x67\xe5\x0a\x1c\x52\x4f\x91\x31\xe8\x83\x75\xed\x03\x5e\x80\x2b\x90\x49\x72\x12\xa4\x8c\x61\xf4\x3e\xaf\x18\xbf\xda\xef\x37\x82\xce\xc1\x9e\x19\x53\x60\x70\xe0\x09\x2a\x00\x8c\x0b\x3c\x22\xd4\x38\x38\x48\xfb\x2b\x51\x62\xe1\xb3\x2a\xd0\x85\x19\xbe\x9c\x14\x4f\xd0\xe6\xe3\x7d\x6b\xf3\x96\xbd\x65\x0f\x71\x86\xe2\x74\xcf\x49\x3c\xd7\x9d\x63\xfd\x04\x7d\x92\xfb\xd6\x27\x49\xd3\x34\x7d\xf8\x73\xf4\xc7\x38\x3f\xc3\x9a\x58\x35\xc1\x95\xd7\xff\xe1\xca\xd3\x3f\x32\x25\x7d\x7d\x68\x72\xd4\x44\x33\x9a\x4d\xbd\x9e\xa1\x4a\x5f\x39\x41\x97\xe5\x7d\xff\x5a\xb2\x6c\xb9\x7c\xef\xca\xfc\x0a\x00\x00\xff\xff\xea\xda\x59\x3f\xf4\x0b\x00\x00")

func templatesNetwork_security_groupTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesNetwork_security_groupTf,
		"templates/network_security_group.tf",
	)
}

func templatesNetwork_security_groupTf() (*asset, error) {
	bytes, err := templatesNetwork_security_groupTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/network_security_group.tf", size: 3060, mode: os.FileMode(480), modTime: time.Unix(1516127018, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesOutputTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x93\x51\x6f\xdb\x20\x14\x85\xdf\xfd\x2b\x90\xb5\x87\x55\xda\x52\x2d\x92\xf7\x10\x69\xbf\x05\x11\xfb\xd6\x61\xc5\x80\xee\xbd\x38\xed\xaa\xfc\xf7\x09\x13\x67\x26\x35\x4b\x5a\x3f\xc2\x39\xdf\x3d\x1c\xb0\x0b\xec\x03\x8b\x7a\xb4\xc0\xd2\xaa\x01\x6a\xf1\x56\x09\x31\x2a\x13\x40\xfc\x12\xf5\x97\x37\xf5\x27\x20\xe0\x20\x47\x8d\x1c\x94\x91\x16\xf8\xe8\xf0\x79\xb3\x77\x74\xd8\x44\xc7\xa9\xae\x4e\x55\x35\x83\x28\xec\x6f\xa2\x92\xa6\x44\x40\x20\x17\xb0\x05\xd9\xa3\x0b\xfe\xff\xa4\x5c\x5b\xcc\xc4\x0e\x55\x0f\x52\xb5\xad\x0b\xf6\x56\xb8\x5c\x5c\x62\x76\xf0\xa4\x82\x61\x49\xd0\x06\xd4\xfc\x9a\x12\x14\xa9\xe7\xd6\xae\xe4\x25\x38\xbc\x30\xa0\x55\x46\xea\x32\xd1\x87\xbd\xd1\xad\xd4\x67\x88\xf6\x52\x75\x1d\x02\xd1\x55\x4e\x8d\xd0\xb2\xc3\x79\xf7\x8a\x77\x60\xf6\xb4\x7b\x7c\xbc\x87\xbb\xdb\x36\x4d\xd3\x64\x74\x8f\x7a\x54\x0c\xf2\x19\x5e\x97\xe0\xf8\x4d\x61\xd9\x90\x5c\x68\x26\xa4\x1c\x07\xda\x2c\x16\xa5\x87\xe1\x54\x57\x42\x10\x58\xd2\xac\xc7\x18\x8c\x31\x40\x36\x28\xa5\xfa\xf8\x9c\x8b\x4f\x3a\x0f\x96\xe8\xf0\x6e\xd4\x93\x32\x94\xcd\xfa\x1d\x06\xbf\x77\x2f\x32\xa0\xf9\x44\xfb\xbb\xed\x36\xab\x68\xbe\xf9\x56\x77\xf8\x0e\x37\x2a\xdc\x2c\x05\xe9\xee\x8c\x6b\x95\xa1\x49\x7b\xb9\xbe\xf8\x48\xa2\x27\x8e\xfb\x9e\x8c\x60\x47\xa9\xbb\xe9\x3c\xda\x9e\x1f\x4c\x84\xfc\x43\x67\xcb\xb9\xb0\x3f\x26\x59\xdc\x39\x38\xe2\xaf\xd3\xd0\xdc\xf1\x4d\xfc\x78\x98\x5c\x73\x23\x97\x5d\xed\xef\x71\x37\xc9\x7d\x39\xc3\x07\xed\x3f\x1f\x0a\x4f\x79\xf5\xff\x4d\x88\x4c\x93\xdb\x33\x7a\xc1\x7e\x5d\xd8\x9a\xbd\x3f\xde\x32\xf7\xc7\xdc\x3a\xd7\xb7\x2c\xa0\xc0\x58\x69\xba\x50\xc2\x1d\xb0\xb5\xe2\x13\xed\x6f\x00\x00\x00\xff\xff\x2b\xd8\x75\x99\xf7\x05\x00\x00")

func templatesOutputTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesOutputTf,
		"templates/output.tf",
	)
}

func templatesOutputTf() (*asset, error) {
	bytes, err := templatesOutputTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/output.tf", size: 1527, mode: os.FileMode(480), modTime: time.Unix(1515632091, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesResource_groupTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x90\x41\x8a\xc4\x20\x10\x45\xf7\x9e\xe2\x23\xb3\x9d\xdc\x60\xce\x22\x15\x53\x64\x84\x44\x43\xa9\x59\x4c\xf0\xee\x83\x42\xa7\x3b\x24\xdd\x0d\x5d\x4b\xf9\xf5\xeb\xf9\x84\x63\xc8\x62\x19\x9a\xfe\xb2\xb0\xcc\xe6\xf6\x62\x46\x09\x79\xd1\xd0\x7d\x88\xbf\x1a\x9b\x02\x3c\xcd\x8c\x3a\x3f\xd0\x5f\xdb\x4a\xd2\xb1\x5f\x8d\x1b\xca\x77\xcb\x28\x60\x0a\x96\x92\x0b\xfe\x9e\x10\x1e\x5d\xf0\x45\x2b\x05\x24\x1a\x63\x2b\x02\xd8\xaf\x4e\x82\x9f\xd9\xa7\x53\x5b\x2d\x2a\xaa\x28\x75\x86\x5b\x72\x3f\x39\x6b\xdc\x13\xae\xab\x79\xcf\xfa\x72\x6b\xe7\x07\x8e\x66\xcc\xf1\x6a\x5b\xb8\x76\xd8\xd5\x8b\x5d\x8d\xb7\x9a\xfd\x0f\x86\x86\x41\x38\x46\x43\xd3\xa3\xb7\x98\x28\x39\xfb\x89\xb0\xff\x00\x00\x00\xff\xff\x58\xec\xbb\xb4\xcd\x01\x00\x00")

func templatesResource_groupTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesResource_groupTf,
		"templates/resource_group.tf",
	)
}

func templatesResource_groupTf() (*asset, error) {
	bytes, err := templatesResource_groupTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/resource_group.tf", size: 461, mode: os.FileMode(480), modTime: time.Unix(1513729670, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesStorageTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x91\xcd\x6a\xc3\x30\x0c\xc7\xef\x7e\x0a\x61\x76\xee\x1b\xec\xbc\xfb\xfa\x00\x46\x71\x44\x66\x70\xa4\x20\x2b\x81\xad\xe4\xdd\x47\xbc\x36\x6d\x0d\xa5\x3d\x2e\xd7\xfc\xf4\xff\xb2\x52\x91\x59\x23\x81\xc7\x9f\x59\x49\xc7\x50\x4c\x14\x07\x0a\x18\xa3\xcc\x6c\x1e\x7c\x27\xe5\xcb\xc3\xc9\x01\x30\x8e\x04\xcd\xf7\x0e\xfe\xed\xb4\xa0\x1e\x4a\x1a\xa7\x4c\x81\x78\x09\xa9\x5f\xbd\x03\xb8\x88\x87\x41\x65\x9e\x42\xbd\xae\xf8\xc5\xeb\x1e\x38\x6c\x46\x87\x8d\x5a\xbd\x73\x00\x59\x22\x5a\x12\x6e\x1d\xaf\x96\x4a\x43\x12\xae\x5e\xe7\xb8\xc1\x12\x69\x0b\x1f\x0d\xb9\x47\xed\x6f\x39\xa5\x29\xa7\x3f\xfd\x60\xdf\x53\x0d\xf6\xf1\x79\xac\xc6\x86\x43\xa9\x7d\x01\x88\x97\xa4\xc2\x23\xb1\x5d\x6d\x6f\x2a\xae\x6e\x75\xee\xf1\x88\x51\xd8\x30\x31\xe9\xd3\x19\x6b\xd0\x8a\x3c\x18\x0e\x5e\x9e\x0e\xa0\x79\xc3\xb3\xc0\xdd\x7d\x83\x34\x02\x7b\xee\xed\x3f\x95\xb2\x4f\x34\x69\x5a\xd0\xc8\xbf\x5e\xbb\x18\x8d\x91\x72\x7e\x52\x7d\xc7\xfe\x75\xfd\x2e\x4b\xb7\x75\xff\x0d\x00\x00\xff\xff\x23\xe4\x2a\x58\x37\x03\x00\x00")

func templatesStorageTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesStorageTf,
		"templates/storage.tf",
	)
}

func templatesStorageTf() (*asset, error) {
	bytes, err := templatesStorageTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/storage.tf", size: 823, mode: os.FileMode(480), modTime: time.Unix(1513729670, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesTlsTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x04\xc0\x41\x0a\x02\x31\x0c\x05\xd0\x7d\x4e\xf1\xc9\x09\x5c\x88\xe0\xa2\x0b\xaf\xa0\x07\x08\xad\x04\x5b\x6c\xa9\x24\xb1\x30\x0c\x73\xf7\x79\xa6\x3e\xff\xf6\x56\x70\x74\x97\x9f\xb5\x95\x43\xe5\xab\x1b\x83\xcb\xf4\x2a\x6b\x38\x63\x27\x20\xf7\xcf\xb4\x16\x75\x20\x81\x9f\xaf\x07\x13\x60\x9e\xa5\xb4\x70\x20\xe1\x7a\xb9\xdf\xe8\xa0\x33\x00\x00\xff\xff\x52\x4d\xac\xad\x51\x00\x00\x00")

func templatesTlsTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesTlsTf,
		"templates/tls.tf",
	)
}

func templatesTlsTf() (*asset, error) {
	bytes, err := templatesTlsTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/tls.tf", size: 81, mode: os.FileMode(480), modTime: time.Unix(1513729670, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesVarsTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x90\xc1\x6a\xc4\x20\x10\x40\xef\x7e\xc5\x20\x3d\xa7\xcd\xa5\xb7\x7e\x4b\x30\x3a\x2d\x43\xcd\x18\x46\x63\xa1\x21\xff\x5e\x12\xc1\xee\x8a\x2c\x9b\xdc\xe6\xbd\x07\xce\x64\x23\x64\x66\x8f\xa0\x91\xf3\x44\x4e\xc3\x7e\x28\xf5\x3f\x15\xfc\xa2\xc0\xed\x34\xd2\xb2\x7a\x9c\xfa\x49\xdc\xe6\x68\x85\xd6\x44\x81\x3b\x38\x21\x1b\x4e\x1d\x60\x3d\xe1\x23\x10\xd1\x0a\xa6\x16\x32\xa6\x9f\x20\xdf\x93\x25\x27\x1a\x76\x05\xe0\xf0\xd3\x6c\x3e\xc1\x07\xe8\xf1\x6d\xb8\xfe\xd7\xf1\x5d\xab\xbb\x8c\x38\xa1\xb0\xf1\xcf\x75\xab\x84\x4c\x0e\x05\xb4\xf9\xdd\x04\x65\x29\x45\xb3\xe9\x59\xbe\xec\xd9\xc8\xd0\x80\x43\x2b\x80\xba\x37\x94\xaf\xca\x15\x5c\x5a\xbd\x42\xab\x55\x70\xab\x95\x9b\x74\xb4\x02\x8e\xf3\xf5\x7f\x01\x00\x00\xff\xff\x59\x69\xa5\xda\xe3\x01\x00\x00")

func templatesVarsTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesVarsTf,
		"templates/vars.tf",
	)
}

func templatesVarsTf() (*asset, error) {
	bytes, err := templatesVarsTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/vars.tf", size: 483, mode: os.FileMode(480), modTime: time.Unix(1516130326, 0)}
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
	"templates/cf_lb.tf": templatesCf_lbTf,
	"templates/concourse_lb.tf": templatesConcourse_lbTf,
	"templates/network.tf": templatesNetworkTf,
	"templates/network_security_group.tf": templatesNetwork_security_groupTf,
	"templates/output.tf": templatesOutputTf,
	"templates/resource_group.tf": templatesResource_groupTf,
	"templates/storage.tf": templatesStorageTf,
	"templates/tls.tf": templatesTlsTf,
	"templates/vars.tf": templatesVarsTf,
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
	"templates": &bintree{nil, map[string]*bintree{
		"cf_lb.tf": &bintree{templatesCf_lbTf, map[string]*bintree{}},
		"concourse_lb.tf": &bintree{templatesConcourse_lbTf, map[string]*bintree{}},
		"network.tf": &bintree{templatesNetworkTf, map[string]*bintree{}},
		"network_security_group.tf": &bintree{templatesNetwork_security_groupTf, map[string]*bintree{}},
		"output.tf": &bintree{templatesOutputTf, map[string]*bintree{}},
		"resource_group.tf": &bintree{templatesResource_groupTf, map[string]*bintree{}},
		"storage.tf": &bintree{templatesStorageTf, map[string]*bintree{}},
		"tls.tf": &bintree{templatesTlsTf, map[string]*bintree{}},
		"vars.tf": &bintree{templatesVarsTf, map[string]*bintree{}},
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

