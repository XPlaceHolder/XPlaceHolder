// Code generated by go-bindata.
// sources:
// templates/db_mysql.tf
// templates/jumpbox.tf
// templates/load_balancer.tf
// templates/output.tf
// templates/provider.tf
// templates/resource_group.tf
// templates/storage_account.tf
// templates/subnet.tf
// templates/vars.tf
// templates/vm_web_server.tf
// templates/vmss_web_server.tf
// templates/vnet.tf
// DO NOT EDIT!

package terraform

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

var _templatesDb_mysqlTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x53\xcb\xae\x9b\x30\x10\xdd\xf3\x15\x23\xab\xab\x4a\xbd\xea\x2b\xea\x2a\xbb\x56\x77\x7b\xff\xc0\x1a\xf0\x04\x59\x31\x63\x3a\x63\x53\xa5\x11\xff\x5e\x61\x20\x24\xb9\xa8\x2c\xcf\x9c\xc7\x18\x1f\x0f\x28\x1e\xeb\x40\x60\x88\x5b\xcf\x64\x07\x12\xf5\x91\x0d\x5c\x2b\x00\x47\x27\xcc\x21\xc1\x11\xcc\xe1\xe5\x87\xa9\xc6\xaa\xda\x04\x4d\x14\x52\x03\xd7\x07\x50\x53\x14\x6c\xe9\x19\xae\xb1\x39\xe7\xde\x0a\x25\xe2\xe4\x23\x5b\x87\x17\x7d\xce\xf8\x76\x78\x0c\x70\x98\xb0\x46\x25\x9b\x95\x84\xb1\x7b\xe7\x7a\x23\xf4\xa8\xfa\x27\x8a\x9b\x09\x42\x1a\xb3\x34\x04\x06\xff\x66\x21\xe9\x6c\x77\xd1\xdf\xc1\x2a\xc9\x40\x62\xc0\xb8\x7a\x46\xe6\x05\x26\x67\xd8\xfd\x8e\x60\x3e\x5c\x07\x94\x17\xe2\xc1\x4e\xb4\xf1\xd3\xac\xab\x00\xd6\x10\xdb\x4a\xcc\xbd\xbd\x33\x29\xaa\x35\xf9\x91\xf6\x72\xce\x1c\x32\x3f\xa3\xc5\x7a\x32\x0d\xb1\xc1\xe9\xf7\xfc\x67\x95\x95\x32\x9a\xaa\x02\xd0\x73\x2e\x67\x58\x4e\x71\x04\xf3\xfa\x66\x5f\x89\x0f\x76\x66\x97\x3b\x2a\xd6\x00\x0d\xf6\xd8\xf8\x74\xd9\xbc\xee\xa7\xc9\x93\x14\x3d\x31\x09\x86\xb7\x2c\x7d\x54\x9a\x67\x27\xec\x7c\xb8\x2c\xd3\xc3\x84\x8d\x25\x7c\xbe\x6b\xdb\x4b\x3c\xf9\x40\xcb\x22\x2b\xda\xd5\x5b\xd0\x82\xc1\x47\xf8\xf2\xf9\xeb\xf7\x25\x71\xb7\x13\x9b\x66\x77\xbc\x48\x5b\x8a\x56\xc8\x65\x76\xc8\xc9\xce\xcc\x49\xf9\xd3\xeb\xd4\x0c\x77\x5b\x11\x5d\xe7\xd9\x6b\x12\x4c\x51\x6c\x88\xad\xe7\x2d\xe1\x5d\xbf\x8a\xfb\x8e\xe4\xd6\xaf\x1d\xed\x3a\x2a\xda\xe5\xf1\xdc\x17\xe7\xfe\x51\x15\x8e\x6a\xb0\xc4\xa7\x28\x0d\x75\xc4\xe5\x6d\xfd\xe2\x65\xeb\xb1\xfa\x17\x00\x00\xff\xff\x43\xf9\xfb\x23\x90\x03\x00\x00")

func templatesDb_mysqlTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesDb_mysqlTf,
		"templates/db_mysql.tf",
	)
}

func templatesDb_mysqlTf() (*asset, error) {
	bytes, err := templatesDb_mysqlTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/db_mysql.tf", size: 912, mode: os.FileMode(480), modTime: time.Unix(1541605382, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesJumpboxTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x56\x4d\x8f\xdb\x36\x10\xbd\xeb\x57\x4c\xd5\x9e\x0a\xd8\xd8\x75\x83\xde\x74\x28\x52\x14\x08\x90\xa2\x45\x9c\xb6\x87\xa2\x20\x68\x72\x24\x4f\x4d\x91\xc2\x90\xd4\xee\x26\xd8\xff\x5e\x50\x5f\x6b\x7b\x25\x7b\x91\x6e\x7c\x30\x04\xf1\x71\xde\xe8\xf1\xbd\x91\x18\xbd\x8b\xac\x10\x72\xf9\x29\x32\x72\x2d\x2c\x86\x3b\xc7\x07\xe1\x51\x45\xa6\xf0\x20\x2a\x76\xb1\xc9\x21\x3f\x44\x6b\xa2\x15\xff\xc6\xba\xd9\xb9\xfb\x45\xdc\xe7\x0c\xc0\xca\x1a\xe1\xec\x57\x40\xfe\xdd\xe7\x56\xf2\x1a\x6d\x2b\x12\xe0\x71\x35\x94\x5a\x59\x5f\xe5\x19\x80\x71\x4a\x06\x72\xf6\xd9\xae\xb1\xb5\xb1\xd9\x9e\x6a\x3d\x34\x74\x76\x77\xac\xf2\x98\x4a\x9e\xae\x75\xb4\x5f\x50\xb2\xeb\x36\xcf\x1e\xb3\xec\x05\x72\x71\x34\xf8\x02\xb5\x12\x4c\x78\xbf\x5f\x16\xec\x48\x82\x9f\x8c\x71\x77\xab\xad\xdf\xa7\x67\x6a\x98\x5c\x2a\x31\x8f\xbd\xbd\xb9\xc9\x00\x34\x31\xaa\x73\x2d\x9f\xea\xbd\xb3\x3b\x17\xad\x4e\xd5\xa4\x52\xe8\xfd\x65\xe6\x9e\xd5\x05\xa7\x9c\x59\xc0\x7d\x54\x4d\x42\x0d\xb2\x35\x8e\x83\x60\x69\x2b\x3c\x45\x7d\x9f\x30\x1a\x7d\x20\xdb\x9d\xd1\x33\x60\x01\xf9\x66\x73\x54\x48\x6a\xcd\xe8\xbd\x68\x18\x4b\xba\xbf\x50\xe8\x1c\x38\x62\xe6\x0c\xf0\xe5\xde\x1a\x8c\x00\x30\x6f\xfe\x19\x7f\xcd\x03\xd7\x2f\xca\xd2\x45\xdf\x35\x71\x67\x48\x09\x9a\x49\xe6\xd1\xd2\x35\x6f\x5d\x4a\x65\x5f\x66\x45\xcd\x52\x36\xff\x87\x90\x57\x43\xfa\x1a\x87\x34\xe9\x30\xb9\x43\x9a\xe9\x39\x0a\xc8\x7d\x90\x81\xd4\x95\x5c\x93\x0d\xc8\xa5\x54\x73\x99\x26\x75\x4d\xe1\x4b\x43\x2f\x51\x5f\x12\xf6\xab\xa9\xfa\xfa\xbe\x27\xfd\x8a\xae\x27\xdd\x51\x51\x23\x94\xb3\x25\x55\x91\x7b\x81\x92\xd2\x57\xdc\x7c\x45\xef\xde\xca\x00\x3e\xee\x2c\x86\xd4\xf5\x55\x37\xf7\xd0\xb5\x97\x76\xbc\x1c\xda\xeb\xe6\x70\x2b\x03\x2e\xfb\x4b\x3f\x58\x59\xf7\xa7\x3c\x67\xc6\x63\xfa\x13\xce\x09\x7a\x2e\xd8\xd3\xc2\xd0\xc4\xe3\xbc\x77\x5b\xe2\x10\xa5\x11\xb5\x54\x7b\xb2\xcf\x9d\x7b\xc9\xb5\xcb\x0a\x2e\xbb\xf5\xab\x38\xf5\x15\x5c\x3a\x65\x57\x90\xf6\x50\xc0\xdf\x33\x16\x9d\x30\xcf\xdc\x49\xaa\x93\xf9\x9f\x0c\xa0\xad\x85\xa7\x4f\xe7\x6a\x15\x90\x6f\x83\xb4\x5a\xb2\x16\x3f\x6f\x37\xa2\xdd\xe4\x59\x96\x01\x7c\x0b\x6f\x5d\x5d\xa3\x0d\x10\xf6\xe4\xc1\x90\x45\x08\x0e\x0e\x88\x0d\x84\x3d\xc2\x6f\x5b\xd0\xe4\x0f\x20\xcb\x80\x0c\x1a\x0d\x06\xb2\x55\xb7\xf4\xe7\xaf\xdd\x7b\xcd\x60\x40\xe1\xbc\x48\x30\xe1\xac\x08\xc8\xf5\xf0\xa6\x83\x02\x02\x47\x7c\x09\x8f\x96\x41\x76\x4c\xfe\x1a\x55\x42\x76\x64\x7e\x99\xcd\x07\xc7\xb2\x42\x41\x75\xfa\x67\x2c\x91\xd1\x2a\x1c\x62\xd9\x59\xd3\xef\x91\x93\x2a\x6f\xa5\x75\x96\x94\x34\xbd\xf5\x5d\x59\x22\x8f\x8a\xfd\xb1\x8b\x36\xc4\x2d\x72\x8b\x3c\xa4\xf1\x10\x9f\x04\xbd\xfd\x71\x7d\xf3\x66\xf5\xfe\xe3\xb6\x5f\x6b\x91\x7d\x6f\xb7\x02\x72\x23\x03\xfa\xd0\xfb\xfe\xa8\xa1\x41\xa7\xc5\xf9\x70\x61\x26\xa4\x7d\x3d\x91\xea\xa2\x52\x9d\x6e\xfb\x80\x52\xff\xc5\x14\x70\xc0\x30\xa6\xbc\xbb\x66\x4a\x40\x01\xf9\x2f\xec\xea\x77\x49\x92\x1e\x53\x4b\x2b\x2b\xd4\xfd\xc9\x85\x87\x06\x4f\x5c\xf2\xfe\xc3\x76\x6a\xdf\xa5\x8f\x15\x57\x92\x19\x25\x54\xae\x6e\x62\x40\x1e\x02\x50\x40\x7e\x14\x3c\x00\xa9\x6b\xb2\x22\x7a\xe4\xf1\x2b\x23\xa9\xd1\xdd\x3d\x06\x34\xd2\xfb\x3b\xc7\xdd\x40\xfe\x7d\xb8\xbe\xdd\xfc\xf0\xe6\x9b\x19\x62\x61\xc8\xc6\xfb\x61\xcc\x0e\x5d\x68\xf2\x72\x67\x70\xaa\x23\x64\x0c\x7b\xb4\x81\xa6\xb9\x56\x4a\xe3\x71\x2c\x16\x64\xe5\x87\x9d\x68\x5b\x62\x67\x3b\x43\xf6\x6f\xd7\x8a\x6c\x35\x8e\xa9\xff\x02\x00\x00\xff\xff\x84\x09\x2b\x1d\x69\x0c\x00\x00")

func templatesJumpboxTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesJumpboxTf,
		"templates/jumpbox.tf",
	)
}

func templatesJumpboxTf() (*asset, error) {
	bytes, err := templatesJumpboxTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/jumpbox.tf", size: 3177, mode: os.FileMode(480), modTime: time.Unix(1541605382, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesLoad_balancerTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x94\x3f\x8f\xd4\x30\x10\xc5\xfb\x7c\x0a\xcb\xa2\xcd\x42\x49\xb3\x05\x27\x21\x81\x44\xb1\x42\xf4\x96\x63\xfb\xf6\x4c\x7c\x1e\xe3\x3f\x5b\xb0\xca\x77\x47\x76\xe2\xbd\x4b\xe2\x39\x09\x0e\x0a\xb6\xdc\x79\x79\x9e\xf9\x3d\x7b\xbc\x0a\x90\xbc\x50\x84\xf2\x9f\xc9\x2b\xff\xc8\x5c\x1a\x8c\x16\x4c\x3b\x4a\xe8\x98\xac\x49\x96\x19\xe0\x92\x0d\xdc\x70\x2b\x94\x2f\x95\x6b\x47\x88\xe5\x8f\x8a\x60\xbf\x23\xa1\x6f\xae\x17\xee\x0f\xca\x5e\x58\x56\x4e\x7d\x76\xe9\xab\x4b\xaf\x1d\xed\x08\x31\x20\x78\xd4\x60\x71\x8f\xda\x56\x6d\x94\x9d\x3d\x24\x77\x58\x3a\xdb\xfc\x5b\xed\xa6\xec\xbd\xae\xb1\x75\xbb\xbf\xef\x5d\x86\xc8\xbe\x37\x40\x8c\x4b\xe9\x55\x08\x8c\x9b\xdb\x1c\x47\x42\x43\xe4\x51\x8b\xac\x0c\x63\x42\x01\x3d\x11\x5a\xd3\x0d\x63\x9a\x68\x37\x75\xdd\x3e\x19\x33\x20\x91\xe0\x79\xb4\x62\x18\x30\xf0\xff\x80\xf7\x2b\x30\xb7\xe0\xbd\xc8\x8c\x90\x7b\x0f\x36\x2a\x2b\x73\x36\x02\xec\xbd\x3e\x27\x3f\x4f\x99\xf9\x20\x37\xf6\x48\xe8\xa9\x24\xfa\xf9\xf4\x61\xce\x93\x16\xf1\x3e\x66\x2d\xd7\xf3\xdc\x14\x07\xe4\x9d\x1c\xb4\x2c\x7d\x4d\x58\x9e\x6c\xe0\x62\xcc\x0d\xd7\x23\x1c\x80\xc1\xde\x5d\x5b\x7b\xfd\xfb\xe4\xf3\xb1\x4f\x53\xc8\xfd\xdd\x30\x43\x73\xe2\x3a\x2e\x72\x11\xef\xb8\x18\x3f\x5a\xb9\x40\x3e\xe5\xee\x51\x2e\xce\xc3\xa0\x30\x10\x4b\xf1\xff\x99\xfc\x21\x46\xd7\x07\xe5\x2f\xda\x9e\xfb\xb9\xfb\xbc\x47\x3c\x44\x10\x60\x36\xe2\x4f\x31\xba\xf9\x39\xfd\x48\x2a\x44\xe6\x78\x7c\x78\x56\x7e\x5b\x3e\x05\x1f\xf7\xe7\xbc\x7f\x87\xf2\xf4\xc9\x64\x9c\xdf\xc1\x4a\x48\x1b\x9c\x73\x0d\xa3\xf9\xaa\xfd\xf0\x22\x58\xcc\xf9\x0f\x18\xaf\xad\xbe\xdc\x7d\xcd\x23\x21\x88\xd7\xda\x6f\xc2\xad\x36\x47\x8b\xec\xc2\x96\x90\xfa\x02\xdb\xa2\xaa\x42\xb7\xd0\xed\x7a\x36\xf6\x4d\xeb\x71\x3f\x03\xb5\x01\xd4\xdc\x05\xed\x2d\xd4\x54\x2e\x28\xcb\x55\x6c\xc4\x81\x1c\x5a\xe4\xed\x53\xe6\x52\xb1\x9d\x7e\x05\x00\x00\xff\xff\x55\x3f\x0e\x54\x53\x08\x00\x00")

func templatesLoad_balancerTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesLoad_balancerTf,
		"templates/load_balancer.tf",
	)
}

func templatesLoad_balancerTf() (*asset, error) {
	bytes, err := templatesLoad_balancerTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/load_balancer.tf", size: 2131, mode: os.FileMode(480), modTime: time.Unix(1541605382, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesOutputTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\xcd\x41\x0a\xc2\x40\x0c\x85\xe1\xfd\x9c\x22\x0c\xae\x7b\x03\xcf\x12\x82\x04\x29\x9d\x49\x4a\x4c\x22\x58\x7a\x77\x51\xa8\xd8\x0a\x6e\x1f\x7f\xbe\x68\xf8\x1c\x0e\x95\x25\x47\x53\xe9\x2c\x5e\x61\x29\x00\x49\x2d\x18\xce\x50\x4f\x4b\x92\x0d\xf4\x08\x63\xfc\xaa\xd6\x5a\xd6\x52\xb6\x73\xe3\x9b\x86\x5d\x18\xaf\xa6\x31\xa3\x50\xe7\x1f\xe6\x4d\x58\xc7\x7d\x3b\x4c\x21\x2d\xe4\xb8\xbe\x84\xfd\x8b\x14\xf6\xff\x70\x8e\xe6\x41\x0d\x85\xfd\xae\x36\x6d\xf2\x71\xfe\xd0\xcf\x00\x00\x00\xff\xff\xb3\xf3\x67\xee\xfc\x00\x00\x00")

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

	info := bindataFileInfo{name: "templates/output.tf", size: 252, mode: os.FileMode(480), modTime: time.Unix(1541605382, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesProviderTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x8e\xc1\x0a\xc3\x20\x0c\x40\xef\x7e\x45\x90\x1d\x36\x18\xfb\x83\x7d\x4b\x71\x36\x87\xb0\x36\x4a\x8c\x1e\x56\xfc\xf7\x51\x0b\xb5\x14\x36\x8f\xbe\xf7\x92\x44\x09\x85\x46\x14\xb0\xee\x93\x05\x65\xb6\xb0\x18\x00\x3f\x11\xb2\x0e\x34\xc2\xf6\x9e\x60\x2f\x4b\x71\xf2\x68\xd6\xb0\xe3\x6a\xbb\x9c\xd0\x0b\xea\x4f\x79\xc3\x2d\x48\xf9\x95\xbc\x50\x54\x0a\xbc\xee\x38\x05\x27\xdc\x12\x45\x76\x7f\x0e\xda\x71\x93\x91\x0b\x49\xe0\x19\x59\xbb\x3c\x85\xf0\xce\xf1\xba\x36\x07\x9e\xee\xd0\xa7\x1c\xfe\x6f\xd5\x9a\x6a\xbe\x01\x00\x00\xff\xff\xcd\xed\xb3\xce\x1e\x01\x00\x00")

func templatesProviderTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesProviderTf,
		"templates/provider.tf",
	)
}

func templatesProviderTf() (*asset, error) {
	bytes, err := templatesProviderTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/provider.tf", size: 286, mode: os.FileMode(480), modTime: time.Unix(1541605382, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesResource_groupTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x4a\x2d\xce\x2f\x2d\x4a\x4e\x55\x50\x4a\xac\x2a\x2d\x4a\x2d\xca\x8d\x87\x89\xc4\xa7\x17\xe5\x97\x16\x28\x29\x28\x65\x97\xe6\xe5\x94\xe6\x61\x88\x57\x73\x29\x28\xe4\x25\xe6\xa6\x2a\x80\x80\xad\x82\x92\x4a\x75\x59\x62\x91\x5e\x6a\x5e\x59\x3c\x48\xb4\x56\x17\xa2\x4d\x89\x4b\x41\x21\x27\x3f\x39\xb1\x24\x33\x3f\x0f\xa1\x0a\x26\x52\xab\xc4\x55\xcb\x05\x08\x00\x00\xff\xff\x26\x45\x5a\x39\x83\x00\x00\x00")

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

	info := bindataFileInfo{name: "templates/resource_group.tf", size: 131, mode: os.FileMode(480), modTime: time.Unix(1541605382, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesStorage_accountTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\xd0\x51\x4e\xc3\x30\x0c\x06\xe0\xf7\x9e\xc2\x8a\x78\x00\x69\xdb\x0d\x76\x03\x9e\xd8\x01\x2c\x37\x75\x47\xb4\xcc\xa9\x5c\x67\x52\x99\x7a\x77\x94\x8e\x88\xc2\x86\xc8\x63\xf2\xe5\xcf\xef\x28\x8f\x29\xab\x67\x70\xf4\x91\x95\xf5\x8c\xa3\x25\xa5\x23\x23\x79\x9f\xb2\x98\x03\xd7\x71\x4f\x39\xda\xfd\xc9\xb5\x01\x10\x3a\x33\x3c\x5c\x7b\x70\x4f\x57\xe5\x21\x92\xe7\xe7\x0b\xe9\x8e\xe5\x82\x85\x6f\xc0\x6d\xdd\x06\x9c\x7b\x99\xb7\x23\xb9\x06\xa0\xb6\xc0\xa3\xa6\x3c\xe0\x2a\x73\x09\xa9\xd5\x7e\xb2\xdd\x29\x4b\xcc\xf2\x7b\xb7\x5c\x9e\x4b\x68\x4c\x9e\x2c\x24\xf9\xa3\x59\x69\x54\xc9\xe2\xbf\xc6\xc2\x53\x90\xee\xce\x1f\x6e\xb3\xaf\x9d\x05\xd6\x07\x8e\xa4\x23\xed\xd6\xb0\xfc\x41\xb8\x3d\x84\x36\x0d\xbc\xc0\xd7\xb7\x43\x31\x2c\xd4\x46\xc6\x36\xa6\x16\x59\xbc\x4e\x43\xad\xbc\x07\xd3\xcc\xdf\xa4\x0f\x91\xff\x21\xef\x66\xc3\x88\xa6\xd4\xf7\xc1\x63\x92\x38\x55\x32\x37\x9f\x01\x00\x00\xff\xff\x9c\x2c\x1c\x64\xe9\x01\x00\x00")

func templatesStorage_accountTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesStorage_accountTf,
		"templates/storage_account.tf",
	)
}

func templatesStorage_accountTf() (*asset, error) {
	bytes, err := templatesStorage_accountTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/storage_account.tf", size: 489, mode: os.FileMode(480), modTime: time.Unix(1541605382, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesSubnetTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\xcf\xc1\x0a\x82\x40\x10\x06\xe0\xbb\x4f\x31\x2c\x1d\x0a\x4c\x3c\x76\xe9\x49\x22\x96\x49\xa7\x10\x75\x94\x59\xc7\x22\xf1\xdd\x43\x6d\x25\xf7\xd6\xde\x76\xf8\xe7\x1b\x7e\x21\xd7\xa8\x64\x04\x06\xdf\x2a\x24\xb5\x75\x7a\x63\xea\x0c\x18\x87\xbc\x7e\x86\x08\x80\xb1\x26\x08\xdf\x19\xcc\x6e\xe8\x51\x12\xe2\xde\x4e\x89\xf1\xf8\xdd\x89\x00\x3c\x6e\x1f\xd2\x68\x6b\x17\x60\xde\xf0\xc7\xb6\x89\xa4\x54\xae\x94\xc3\xe9\xcc\x4e\x5e\x5f\x48\xa7\x58\x59\xa6\xee\xd9\x48\xb9\x80\x1b\x2f\x48\x78\x30\x1c\xaf\x22\xe6\xb9\x90\x73\xb6\x15\xba\x17\xaf\xdf\x4e\x59\x91\xcb\xd2\x64\xff\x27\xee\x4d\xd7\x62\x46\x97\xf4\x1a\xc3\x29\x86\xf4\x30\x9a\x68\xfc\x04\x00\x00\xff\xff\x39\x0c\xf3\xab\x6e\x01\x00\x00")

func templatesSubnetTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesSubnetTf,
		"templates/subnet.tf",
	)
}

func templatesSubnetTf() (*asset, error) {
	bytes, err := templatesSubnetTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/subnet.tf", size: 366, mode: os.FileMode(480), modTime: time.Unix(1541605382, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesVarsTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x91\x41\x6e\xe3\x30\x0c\x45\xf7\x3e\x05\xa1\x03\xcc\x4c\x32\xc0\xec\xbc\x18\xa4\x40\x0e\x10\x74\x55\x14\x02\x2d\xb1\xa9\x10\x99\x32\x48\xc9\x41\x13\xe4\xee\x85\x9d\xb4\xb1\x93\xb6\xf6\x8a\xfa\x0f\x9f\xfa\x5f\x3d\x4a\xc0\x26\x12\x18\xe2\xde\x32\xb6\x64\xe0\x78\xaa\xaa\xd9\x79\x90\xc4\x2d\x71\x56\x03\xc7\x0a\x20\xbf\x75\x04\x35\x98\x16\x3b\x53\x01\x78\x7a\xc1\x12\x33\xd4\xa3\x08\xf0\xff\x50\x84\x56\x31\x15\x0f\x97\xaf\x06\xd3\x95\x26\x06\x67\xae\xc0\xe3\x66\x9d\x7a\x92\xd1\x77\x00\x8a\x6e\x3f\xe7\x09\xb6\x26\x69\x91\x2f\x6e\x35\x98\xed\x38\x4f\x80\xd5\x6b\x60\xfc\xd8\x56\x83\x71\xc3\x3c\xe8\xa7\xea\x34\x49\x81\x03\x6b\x27\x59\x6e\x63\x9e\x81\x4c\x8c\x9c\x6d\xf0\x5f\xcb\x5a\x1a\x75\x12\xba\x1c\x12\x7f\x0b\xb9\x18\xe8\x07\x8f\x8b\xac\xe4\x84\xee\x6e\x11\x93\xc3\xc1\xfc\x5c\xf4\xb5\x59\x43\xa8\xb9\xe8\xd2\x54\x33\x7c\x57\x38\x16\xb6\x7d\x90\x5c\x30\x5a\xa6\xbc\x4f\xb2\xb3\xe8\xbd\x90\xaa\xd5\x0e\x1d\x4d\xde\xec\xdc\x50\x0c\x9a\xe7\xef\xf6\x64\x16\x7f\x7e\x8d\xff\xef\xc5\x3f\xf3\x3c\xdf\xb1\xa7\xc6\x2a\x49\x4f\x62\xfb\xd6\xba\x54\xee\xbb\x9b\x23\x1a\x0e\x74\x77\xff\x4d\x46\xf6\x28\xde\x3e\x2c\xd5\xf6\x7f\x6f\x72\xc4\x84\xde\x36\x18\x91\x1d\x89\xd5\x5d\x19\x37\xbc\x07\x00\x00\xff\xff\xa9\xf4\x23\x83\x9c\x02\x00\x00")

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

	info := bindataFileInfo{name: "templates/vars.tf", size: 668, mode: os.FileMode(480), modTime: time.Unix(1541605382, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesVm_web_serverTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x56\x4f\x6f\x1b\xb7\x13\xbd\xeb\x53\xf0\xb7\xbf\x9c\x02\x58\xb0\xdb\xa0\xe8\x45\x87\x22\x40\xd1\x00\x29\x5a\xc4\x69\x7b\x08\x02\x62\x44\x8e\x56\x03\x71\x87\x0b\xfe\x59\xdb\x31\xf4\xdd\x0b\x72\x77\x65\xfd\xa1\xd6\x4a\x10\x1f\xea\x83\x21\x69\x1f\x67\x86\x6f\xe6\xbd\x59\x87\xde\x46\xa7\x50\x54\xf0\x25\x3a\x74\x8d\x64\x0c\x77\xd6\x6d\xa4\x47\x15\x1d\x85\x07\x59\x3b\x1b\xdb\x4a\x54\x9b\xc8\x26\xb2\xec\x1a\xe9\xd1\x75\xe8\xce\x22\x1f\x67\x42\x30\x34\x28\x8e\xfe\x16\xa2\x7a\xf5\xd8\x81\x9b\x23\x77\x32\x01\xb6\x57\x7d\xa4\x2b\xf6\x75\x35\x13\xc2\x58\x05\x81\x2c\x9f\x1c\x1a\x6b\x1b\xab\xed\x33\xcd\x87\x8a\x8e\x7e\x1d\xa3\x6c\x53\xc8\xc3\x67\x39\xeb\x37\x84\xcc\xc5\x56\xb3\xed\x6c\x76\x01\x5f\x2e\x1a\xbc\x88\xae\x04\x94\xeb\x10\x26\x28\xdb\x63\xe1\x17\x63\xec\xdd\xd5\x6f\x09\x3e\x13\xa2\x75\x64\x53\x90\x32\xf8\xe6\xfa\x7a\x26\x84\x26\x87\xea\x98\xcf\xa7\x80\xef\x78\x69\x23\xeb\x14\x0d\x94\x42\xef\xa7\x53\xf7\x59\x6d\xb0\xca\x9a\x33\xb8\x8f\x2a\xd7\x36\x50\xd7\x5a\x17\xa4\x03\xae\xf1\x10\xf5\x3a\x61\x34\xfa\x40\x9c\xfb\x74\x02\x5c\x88\xea\xe7\xeb\xbd\x40\xa0\xb5\x43\xef\x65\xeb\x70\x45\xf7\x13\x81\x8e\x81\x23\xa6\x34\x04\xdf\x3e\x5f\xc3\x30\x08\x51\x9e\xff\xc2\x8c\x95\x81\xf3\x0b\x05\x35\x39\x7d\xd0\x01\x19\x58\x92\x49\x70\x8f\xa1\x34\x78\xa7\x98\xaf\x50\x68\xd7\x8c\x22\x05\xff\x9f\xd7\x28\x71\x40\xb7\x02\xb5\xa7\x4f\x26\xf5\x9c\xfc\x0a\xac\x30\xa9\xab\x57\x8f\xca\x46\x0e\x73\x62\x8d\xf7\xf9\x22\xf9\xfb\x74\x88\x3b\x5c\x8e\x6d\xe9\x1a\x99\x0f\x6c\xcf\xd1\xfa\x62\xe4\xbe\xd4\xe4\x93\xfe\xae\x73\x4f\x3a\x27\xa3\x56\x2a\xcb\x2b\xaa\xa3\xeb\x29\x4a\xdd\x9a\xb6\xcb\x73\x3d\xa3\xb6\xd0\x35\x21\x7c\x5c\x32\x86\x54\xfe\xd9\x50\xe3\x95\x7a\xe8\xdc\x03\x8f\x1f\x87\x2a\xb3\x25\x77\x10\x50\x52\xbb\x33\x22\x30\xbb\xbe\x2e\x44\xa5\x1f\x18\x1a\x52\x09\xbd\x2d\xcf\x69\x47\x2e\x44\x30\xb2\x01\xb5\x26\x3e\xd8\x22\x53\x43\x5a\x96\xed\xe9\x4d\xcb\x43\xf6\x22\x03\xf6\x1d\x86\x6b\x27\x56\x49\xda\x8b\x85\xf8\x54\x98\xac\x1d\x66\xfe\x24\xe7\xf9\xeb\x39\xe9\x4f\x7b\x77\xff\xbc\xad\x3e\xa7\x25\x77\x64\x83\x7d\xbf\x0f\x0a\x3d\x86\x9c\x8e\xea\x09\x62\x68\x7f\x42\xd0\x97\xe3\xde\x9c\x91\x7d\x42\x4e\x18\xc6\x33\x66\x31\x13\xe2\xff\xe2\xad\x6d\x1a\xe4\x20\xc2\x9a\xbc\x30\xc4\x28\x82\x15\x1b\xc4\x56\x84\x35\x8a\x3f\x6e\x85\x26\xbf\x11\xb0\x0a\xe8\x84\x46\x83\x81\xb8\xce\x8f\xfe\xfe\x3d\x2f\x4d\x83\x01\xa5\xf5\x32\xc1\xa4\x65\x19\xd0\x35\xc3\x1a\x15\x0b\x11\x5c\xc4\x4b\xf2\x68\x08\x90\x33\xf9\xe7\x52\x25\x64\x4e\xe6\xcf\x67\xf3\xc1\x3a\xa8\x51\x52\x93\xfe\x3b\x5c\xa1\x43\x56\x38\xe8\xbd\x8d\x4b\x43\x7e\x8d\x2e\xf1\xf3\x16\xd8\x32\x29\x30\xbd\xf4\xec\x6a\x85\x6e\xe4\xee\xaf\x65\xe4\x10\x6f\x33\x6f\x83\xbe\x37\xf1\x89\xda\x9b\x9f\xe6\xd7\x6f\xae\xde\x7f\xbc\xed\x9f\x75\xe8\x7c\x2f\x89\x85\xa8\x0c\x04\xf4\xa1\x17\xe8\x5e\x41\x03\x4f\x67\x8d\x67\x31\xea\xb4\x6b\x12\xae\x64\x31\x2a\x2b\xba\x3e\x3c\xf4\x01\x41\xff\xe3\x28\xe0\x80\x71\x98\xfc\xc3\xb6\x3b\x8d\x2e\x44\xf5\xab\xb3\xcd\xbb\x44\x48\x8f\x69\x80\xa1\x46\xdd\xf7\x2d\x3c\xb4\x79\x53\xde\x06\x60\x0d\x4e\xcb\xf7\x1f\x6e\x77\xc5\xdb\xf4\x1e\x64\x57\x64\x46\x02\x95\x6d\xda\x18\x92\xdd\xe6\x0b\x2c\x44\xb5\xb6\x3e\xa4\x2f\xa5\x82\x41\x37\xc4\x32\x7a\x74\xe3\x3e\x4e\xd4\xe4\x5f\xf7\x01\x2d\x78\x7f\x67\x5d\x36\xfe\x3f\x87\xcf\x37\x3f\xfc\xf8\xe6\x7f\x85\x3a\xa4\x21\x8e\xf7\x83\x99\x0f\x45\x69\xf2\xb0\x34\xb8\x8b\x23\x21\x86\x35\x72\xa0\x9d\x6d\xae\xc0\x78\x1c\x83\x05\xa8\xfd\x70\x12\xb9\x23\x67\x39\x4f\xe7\x42\x54\x3e\x40\x4d\x5c\x4f\x98\xeb\xa9\xaf\x2c\x41\x6d\x90\xf5\xd3\x7b\xa3\xb5\x46\x82\xf7\x56\x51\xce\x7e\xf0\x32\xc5\xa4\x76\x07\x4e\x81\x8f\x13\xeb\xff\xd9\xe5\x5f\x72\xbc\x13\x67\xfe\x5a\xcb\x2b\x6c\xce\x03\x73\xbe\x6c\x39\x16\x29\x3a\xde\xf3\x66\x59\xa4\x72\x2c\xd0\x58\xd0\x72\x09\x06\x58\xa1\x2b\x23\xb3\x8b\x6e\x67\xff\x06\x00\x00\xff\xff\x5d\x3e\x41\x66\x8e\x0e\x00\x00")

func templatesVm_web_serverTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesVm_web_serverTf,
		"templates/vm_web_server.tf",
	)
}

func templatesVm_web_serverTf() (*asset, error) {
	bytes, err := templatesVm_web_serverTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/vm_web_server.tf", size: 3726, mode: os.FileMode(480), modTime: time.Unix(1541605382, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesVmss_web_serverTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x56\x51\x8b\xe3\x36\x10\x7e\xf7\xaf\x10\xa6\x4f\x85\x84\xe4\x7a\x57\xda\x87\x3c\x1c\xa5\xa5\x0b\x57\x58\x6e\xaf\xf4\xa1\x14\x31\x96\x26\x89\x88\x2c\x99\x91\x94\xbd\xdc\x92\xff\x5e\x64\x59\x89\xed\x38\xc9\xb2\x6d\x9e\x76\xad\x4f\xdf\x7c\x1a\x7d\x33\x23\x42\x67\x03\x09\x64\x25\x7c\x0b\x84\x54\x73\x83\xfe\xd9\xd2\x8e\x3b\x14\x81\x94\x3f\xf0\x0d\xd9\xd0\x94\xac\xdc\x05\xa3\x83\xe1\xfb\xda\x39\xee\x90\xf6\x48\x57\xb1\x2f\x05\x63\x06\x6a\x64\xa3\xdf\x8a\x95\xdf\xbd\xec\x81\xe6\x68\xf6\x3c\x02\x8e\xb3\xc4\x34\x33\x6e\x53\x16\x8c\x69\x2b\xc0\x2b\x6b\x2e\x36\x65\x75\x59\x6f\x8a\x34\xef\x34\x8d\xbe\x66\x96\x63\xa4\x1c\xae\xb5\x51\xdf\x40\xd9\x8a\x2d\x8b\x63\x51\xbc\x22\x63\x14\x34\xbe\x32\x61\x11\xca\xb7\xde\xdf\x48\x5a\x2f\x0f\x1f\xb5\xb6\xcf\xb3\xdf\x23\xbc\x60\xac\x21\x65\x23\xc9\x34\x78\xb9\x58\x14\x8c\x49\x45\x28\xc6\x19\x3d\x13\x3e\x98\xca\x06\x23\x23\x1b\x08\x81\xce\xdd\x0e\x9d\xa2\x5a\x6f\x85\xd5\x57\x70\x5f\x44\xab\xad\x4b\x5e\x63\xc9\x73\x02\xb3\xc1\x21\xea\xfb\x88\x91\xe8\xbc\x32\xed\x4d\x5d\x00\x57\xac\xfc\x69\xd1\x23\x02\x29\x09\x9d\xe3\x0d\xe1\x5a\x7d\xbd\x41\x34\x06\x66\xcc\x94\x0d\xde\xee\xb0\xce\x0e\x8c\x4d\x57\xc0\x84\xcb\xa6\x81\xf3\x57\x17\xd5\x4d\x07\xea\x8a\x1b\xf0\xbc\xb1\x56\x8f\x7c\x77\xfe\xfc\x72\x27\x09\xff\x31\x0f\x37\x8c\x9b\x98\x13\xc1\xcc\xb9\x6d\x2a\x74\x90\x15\x68\x30\x02\x89\x2b\x39\x81\x3f\x2b\xd1\x55\x8e\x1e\x77\xf1\xbc\x6d\xae\xe4\xf1\x9e\x21\x07\x9e\x5c\x93\x35\x1e\x8d\x4c\x66\x73\x1e\xc8\x0f\x81\x1f\x16\x8b\xb6\x68\x86\x40\x34\x72\xcc\xf8\x61\xb1\x5c\xfe\x5c\x30\x56\x81\xd8\x65\xdc\x64\xe8\x77\xef\xfa\x74\xaa\xe1\xc2\x9a\xb5\xda\x04\x4a\x66\xcd\x36\x79\x0c\x95\x56\xe2\xe1\xf1\x63\x32\xef\x95\x5b\xde\x2b\xf2\x01\x34\xaf\x41\x6c\x95\x41\xee\x04\x68\xe4\x0e\xfd\xf0\xce\xef\x75\x92\x89\x26\xdc\x6e\xbb\xd2\x7e\xdf\x6a\x8d\xbb\x4d\xf8\xad\xc4\x27\xcf\x85\x66\x43\x20\x63\x93\xd1\x4a\x1c\x78\x6d\x65\x8f\xf4\x0f\x30\x01\x74\xdb\x41\x76\xa1\xcd\x48\x2f\x27\xa7\x14\x3c\x63\x95\xcb\x6e\x5f\x73\xa7\xbe\x25\x66\xc6\xbc\x42\xca\xd0\x27\x0f\x46\x02\xc9\xb4\x22\xa0\x01\x11\xbb\xee\x15\x12\x61\x83\xf1\x2d\xcb\xb1\x88\xd1\xbd\x25\xd8\x20\x6f\xc8\xae\x95\x46\xae\xea\xf8\x1f\xe1\x1a\x09\x8d\xc0\x4e\x59\x13\x0d\xe0\xb6\x48\x91\xf5\x17\x30\xd6\x28\x91\xd4\x33\x66\xd7\xeb\xb3\x96\x3f\xab\x60\x7c\x78\x6a\xa3\xa5\xe5\x78\xbc\xd3\xa9\x97\x3f\xce\x17\xef\x67\x9f\xbe\x3c\xa5\xb5\x3d\x92\x4b\xd7\xb9\x62\xa5\x06\x8f\xce\x5f\x15\x66\x1d\x97\xca\xed\xc6\xa9\xea\x3b\x20\x27\x20\xda\x6f\x33\x5c\xfa\x8c\x20\xff\x22\xe5\xb1\xc3\x10\x82\x47\x6e\x9b\x93\x9b\x56\xac\xfc\x8d\x6c\xfd\x10\x8f\x9f\x30\x35\x18\xd8\xa0\x6c\xa3\x72\x7f\x68\xb0\x9f\x6b\xfe\xe9\xf3\xd3\x55\xa9\x12\x3c\xf4\xc5\xea\x30\x7a\x30\x2c\xa6\x84\xde\x55\xb9\x62\xe5\xaf\x75\xe3\x0f\x69\xb1\x95\x15\x1d\xc1\x37\x55\x37\x50\xb3\x1e\xeb\xb2\x94\x4e\x80\xb0\x75\x13\x7c\xec\xdd\x50\x63\x1e\x3d\x93\x75\x96\xb8\x41\xd6\xca\xf0\xe0\x90\x86\x95\x50\x1f\xda\x95\x3e\xa8\x01\xe7\x9e\x2d\xc9\x33\xe8\x31\x7e\x89\x9f\x96\xef\x7e\x78\x5f\x5e\x6a\xe2\x5a\x99\xf0\xb5\xeb\x34\x9d\x40\xa9\x1c\x54\x1a\x4f\x6c\x1c\x82\xdf\xa2\xf1\xaa\xab\xf7\x15\x5b\x83\x76\x98\xc9\xf2\x08\x1a\x9e\x32\x6b\x3d\x35\xf3\xd8\x36\x3a\x68\x87\x4c\xca\x1b\x52\x35\x50\xac\x10\x4f\x01\xd3\xde\xe9\x39\xa9\xe4\xff\x3c\x25\xdb\x99\xd0\x46\x1c\x37\xdb\xee\x10\x77\x47\x55\xcf\x2e\x17\xa7\x54\xcd\x80\xb2\xec\x18\xf3\x79\xef\x32\x9e\xd2\xc1\x98\x0b\x95\x41\x3f\x31\xfa\x2e\x54\x9c\xb3\x93\xf6\xcc\x1d\x98\xfc\x67\x37\x02\xe3\x6f\x30\x1b\x79\x9e\x4c\xa7\xe7\x90\xb5\x9a\x2b\xe9\xd8\x8a\xfd\x3d\x18\xad\x93\xc8\xc9\x79\x3b\x8d\x8c\x0a\xfe\x99\x94\xa0\xd2\xdb\xb2\x7d\x82\xc4\x67\xae\x6b\xe3\xb7\x67\x1a\x49\xc8\x8f\x94\xf9\xd4\xcb\xa5\x17\xe1\xd8\xfa\xf3\x58\xfc\x1b\x00\x00\xff\xff\x65\xcf\xb3\xe3\xb5\x0c\x00\x00")

func templatesVmss_web_serverTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesVmss_web_serverTf,
		"templates/vmss_web_server.tf",
	)
}

func templatesVmss_web_serverTf() (*asset, error) {
	bytes, err := templatesVmss_web_serverTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/vmss_web_server.tf", size: 3253, mode: os.FileMode(480), modTime: time.Unix(1541605382, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesVnetTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8e\x41\x0a\xc2\x40\x0c\x45\xf7\x73\x8a\x10\xdc\x5a\xed\xc6\x9d\x27\x11\x19\x42\x1b\xa4\xb4\xcd\x94\xcc\x64\x04\x4b\xef\x2e\xad\x56\xb0\x9a\xec\xf2\x1f\xff\x45\x39\x06\xd3\x8a\x01\xe9\x61\xca\xda\xfb\xdc\x68\x32\xea\xbc\x70\xba\x07\x6d\x11\xb0\x35\xe9\x4c\x7e\x83\xd1\x01\x08\xf5\x0c\x9b\x39\x03\xee\xc6\x4c\x5a\xb0\x64\x3f\x03\xd3\x3e\x0b\x27\x74\x00\xab\xce\xdf\x34\xd8\xb0\x84\x2f\x7c\xb5\x7f\x03\xc5\x5b\xbd\xb9\x2e\x9d\x73\x1d\xd5\xb5\x72\x8c\x3e\x0e\x54\xf1\xc7\x7e\xc1\xf2\x58\x2c\x7b\x28\x4f\x78\x75\x00\x5d\xa8\x28\x35\x41\xfe\x7e\xb9\x86\x13\xba\xc9\x3d\x03\x00\x00\xff\xff\x11\x4f\x1f\xbb\x10\x01\x00\x00")

func templatesVnetTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesVnetTf,
		"templates/vnet.tf",
	)
}

func templatesVnetTf() (*asset, error) {
	bytes, err := templatesVnetTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/vnet.tf", size: 272, mode: os.FileMode(480), modTime: time.Unix(1541605382, 0)}
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
	"templates/db_mysql.tf": templatesDb_mysqlTf,
	"templates/jumpbox.tf": templatesJumpboxTf,
	"templates/load_balancer.tf": templatesLoad_balancerTf,
	"templates/output.tf": templatesOutputTf,
	"templates/provider.tf": templatesProviderTf,
	"templates/resource_group.tf": templatesResource_groupTf,
	"templates/storage_account.tf": templatesStorage_accountTf,
	"templates/subnet.tf": templatesSubnetTf,
	"templates/vars.tf": templatesVarsTf,
	"templates/vm_web_server.tf": templatesVm_web_serverTf,
	"templates/vmss_web_server.tf": templatesVmss_web_serverTf,
	"templates/vnet.tf": templatesVnetTf,
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
		"db_mysql.tf": &bintree{templatesDb_mysqlTf, map[string]*bintree{}},
		"jumpbox.tf": &bintree{templatesJumpboxTf, map[string]*bintree{}},
		"load_balancer.tf": &bintree{templatesLoad_balancerTf, map[string]*bintree{}},
		"output.tf": &bintree{templatesOutputTf, map[string]*bintree{}},
		"provider.tf": &bintree{templatesProviderTf, map[string]*bintree{}},
		"resource_group.tf": &bintree{templatesResource_groupTf, map[string]*bintree{}},
		"storage_account.tf": &bintree{templatesStorage_accountTf, map[string]*bintree{}},
		"subnet.tf": &bintree{templatesSubnetTf, map[string]*bintree{}},
		"vars.tf": &bintree{templatesVarsTf, map[string]*bintree{}},
		"vm_web_server.tf": &bintree{templatesVm_web_serverTf, map[string]*bintree{}},
		"vmss_web_server.tf": &bintree{templatesVmss_web_serverTf, map[string]*bintree{}},
		"vnet.tf": &bintree{templatesVnetTf, map[string]*bintree{}},
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
