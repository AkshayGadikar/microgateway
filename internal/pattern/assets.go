// Code generated by go-bindata. DO NOT EDIT. @generated
// sources:
// DefaultHttpPattern.json
// DefaultChannelPattern.json
package pattern

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

var _defaulthttppatternJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xe4\x58\x5f\x6f\xdb\x36\x10\x7f\xf7\xa7\xb8\x0a\x41\x53\x03\xa9\x92\xfd\x79\x32\xe0\x87\x38\x0b\xd6\x75\x0b\x56\xb4\xee\xfa\x30\xec\x81\x96\xce\x16\x63\x8a\xd4\xc8\xa3\x33\xaf\xc9\x77\x1f\x44\x49\x0e\xf5\xc7\xb2\x9d\xc0\x01\x86\xbe\x24\x11\xef\x78\xfc\xdd\xef\x8e\x77\xbc\x7c\x1d\x00\x04\x92\xa5\x18\x8c\x20\xf8\x09\xe7\xcc\x0a\x7a\x47\x94\x7d\x60\x44\xa8\x65\x70\x96\xcb\x0d\x61\x66\x82\x11\xfc\x39\x00\x00\xf8\xea\x7e\x02\x04\x7c\x9e\x6f\x3a\x09\x23\x25\xe7\xa1\x35\xf8\x91\x11\xfe\xc6\x53\x4e\xa8\x61\x3c\x06\xd2\x16\xdd\x7e\xa7\x6c\x50\xaf\x78\xe4\x8e\xf1\xf4\x1e\xe5\x5c\x66\x96\x82\xd1\xc6\x3a\x40\x40\x6a\x89\x32\xdf\xb0\x10\x6a\xc6\x44\x50\x4a\x1e\xdc\xef\x87\xb3\x7e\x30\xef\xbf\x4c\x7b\x41\xbc\xff\x32\xfd\x83\x09\x1e\x33\x52\x7b\xa2\x18\x9f\x84\x19\x5b\x0b\xc5\xe2\x30\x41\x16\xa3\x36\xe1\xa5\xa5\x44\x69\xfe\x2f\x23\xae\xe4\xc6\x4a\x7e\x10\x5f\x48\x2e\x17\x37\x48\x89\x8a\xcb\xcd\x0e\xd9\xed\x1d\x7d\xaa\xc9\xbc\x4d\x4b\x5c\x37\x54\x7f\xc5\xb5\xaf\xc0\x6c\xd3\xd6\xa5\xad\x59\xe0\xc6\x34\x14\x7e\x31\xa6\x86\xcb\xce\x9a\x68\xec\xec\x30\x62\xaf\xb8\x8e\x2c\xa7\x89\x46\xb6\xdc\x11\xe8\xba\x6a\xd0\x63\xfe\x4d\x2b\x70\x73\x26\x0c\xc2\xfd\x3d\x9c\x84\x7e\xa8\x42\x65\x29\xb3\x64\xc2\x55\xbe\x52\x9d\x3e\x84\xd7\xaf\xa1\x34\x91\x3a\x5e\x73\xc9\xe9\xcf\xd7\xd3\xd3\x00\xba\x90\xe5\x29\x3e\x61\xd1\x12\x65\x7c\xf9\x08\x3d\x61\x82\x4a\x34\xbe\x42\x88\x5a\x2b\x0d\xaf\xc6\x20\xb9\x70\x47\xbd\x72\x2b\x21\x37\x12\xc9\xfd\xd9\xb9\x63\x78\x6c\x8f\xdf\xb4\x5d\xfe\xf0\x79\x7a\x5a\xd8\x68\x49\x7e\xff\xb4\x55\x74\x39\xbd\x7a\x77\x3a\xec\x8c\xa1\xe7\xd6\xa4\xf7\x9e\x44\x4a\x12\x4a\xf2\xd3\xab\x5a\xaa\xf2\x6b\x07\xd1\x93\x83\x89\x9e\xec\x41\xf4\xae\xcc\x75\x44\xf6\x45\x7c\x7b\x6e\x0d\xef\xef\xa1\xcf\x85\x23\xc4\x68\xd8\x1d\xa5\xc6\x4d\xeb\x8b\x93\xca\x50\x17\xf5\x6a\x94\x07\xcd\x4a\xaa\xae\xe6\xf3\x2a\x40\x0f\x8f\xe3\x27\xf3\x38\xfe\x9f\xf0\xa8\xd1\x20\x35\x58\x1c\x00\xfc\xe5\xda\xa7\x46\x93\x29\x69\xb0\xa7\x85\x7a\x2d\x71\x73\xdf\x45\xfe\xfd\x11\x59\x94\x60\xdc\x2e\xb3\x8e\x9f\x60\xe4\x56\x37\x8b\xc5\xd6\xe6\xbd\x8c\x73\xd7\x7e\xbc\xf8\xc1\xeb\x03\x31\x23\x56\x53\x73\x5d\x9e\x91\x35\x55\x83\x06\x07\x07\xae\xff\x89\x10\x63\x8c\xe1\x2d\x4c\x13\x84\x92\x2a\x58\x2b\x0b\x09\x5b\x21\x68\xfc\xdb\xa2\x21\x8c\x81\x1b\x50\x2b\xd4\x40\x09\x02\x13\x42\xdd\x61\x0c\xce\x85\x30\xd8\x9c\xf2\xf0\xd4\x16\x5e\xe4\x4e\x7f\x59\x74\xf5\xf3\x39\x04\x7d\xb7\x83\x20\x67\xf1\x06\x8d\x61\x0b\x2c\xcb\x5c\x27\xa2\x9a\xde\x99\x6f\x61\x55\xe8\x72\x25\xf7\x30\xd3\x56\x7e\x22\x8f\xdb\xef\xeb\x49\x58\x97\x6d\xce\x26\xcd\xb3\xec\x65\xd2\xae\xb2\x18\x44\x05\x14\x98\x95\x38\x4b\x0c\x87\x79\x7d\xc4\x87\x44\x8b\x03\x67\x79\x1f\x12\xbe\xbf\xb8\x68\x93\x90\xc7\xbd\x56\x2b\x2b\x40\x4e\xfe\x42\xae\x1e\xed\x05\x71\x14\x8e\x26\x07\x70\x74\x78\xa2\x76\x9d\x7f\xed\xac\x6c\xa9\xeb\x65\x31\xec\x28\xeb\xd5\x40\xd5\x39\xe9\xc4\x68\x22\xcd\xb3\xaa\x75\xb8\x62\x2b\x9a\x4a\x1a\x5d\x8c\x17\x9c\x12\x3b\x0b\x23\x95\x9e\x67\x5a\xdd\x62\x44\x6f\xe7\x42\x2d\xd4\x79\xca\x23\xad\x16\x8c\xf0\x8e\xad\xcf\x59\x44\x7c\xc5\x69\x7d\xae\x19\x61\xcb\x94\x41\x22\x2e\x17\xa6\xee\xb6\x53\xf3\x1f\x6b\xba\xc2\xda\xcf\x6b\xe5\x5a\xf7\xfc\xd4\xf0\xad\xd4\x40\x30\x2a\x45\x70\x93\x94\x79\xae\x8f\xb7\x77\xd4\xfd\xdc\xab\x90\x6d\xeb\xe1\x0d\x6c\x57\x8d\x72\x53\xbd\x02\x9e\x09\xaf\xac\x62\xb3\xe6\xe9\xdd\x51\x48\x8b\xe4\xdb\x04\xc1\x7d\x7b\x79\x48\x89\x46\x93\x28\x51\x9b\xfb\x1e\x17\x7d\x4d\x9e\xa2\xb2\xb5\x88\x56\x4b\x9e\x56\x86\x9a\xd7\x07\xd2\x72\x65\xaf\x98\x77\x4f\x4d\x0d\x5e\x6f\xd8\x12\x81\x49\x48\x88\x32\x88\x98\x10\x40\x2a\x7f\x2f\x68\x98\x15\x5b\xf7\x67\x38\x9f\x1f\x34\x9f\x79\xf9\x8d\x86\x76\x52\xda\x9a\xb9\xd3\xd6\xa4\x6d\x35\xf7\x15\x4a\x60\x9f\xf5\x8e\xff\x31\x74\xf0\x30\xf9\x26\x78\xc8\x4b\xde\xe0\x61\xf0\x5f\x00\x00\x00\xff\xff\x30\x31\xfa\x4b\x2d\x12\x00\x00")

func defaulthttppatternJsonBytes() ([]byte, error) {
	return bindataRead(
		_defaulthttppatternJson,
		"DefaultHttpPattern.json",
	)
}

func defaulthttppatternJson() (*asset, error) {
	bytes, err := defaulthttppatternJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "DefaultHttpPattern.json", size: 4653, mode: os.FileMode(420), modTime: time.Unix(1550003149, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _defaultchannelpatternJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x93\xc1\x6e\x1a\x31\x10\x86\xef\x3c\xc5\xd4\x42\x39\xa5\x0b\xad\x7a\x42\xe2\x90\xb6\xa7\x54\x91\x2a\x35\x6a\x0e\x55\x0f\x83\x77\x00\x83\xf1\xac\xec\x31\x68\x9b\xf0\xee\x95\xbd\xec\xb2\x45\x10\xd2\xcb\xae\x34\xfe\x67\xfc\xfd\x33\x9e\xe7\x01\x80\x72\xb8\x21\x35\x01\xf5\x95\xe6\x18\xad\x7c\x59\xa2\x73\x64\xbf\xa3\x08\x79\xa7\x6e\x93\x24\x08\x55\x41\x4d\xe0\xd7\x00\x00\xe0\x39\x7f\x01\x94\x99\xa7\xbc\x61\xa1\xd9\xcd\x8b\x18\xe8\xfe\xe9\x11\xa6\x53\x10\x1f\x29\xe7\x65\x51\x20\xbf\x35\x3a\xdf\x70\xff\xf4\xf8\x13\xad\x29\x51\xd8\x1f\x05\xc6\x55\x51\xd4\xa4\x2b\x0b\xa0\x84\xd7\xe4\x52\xc6\x74\x58\x54\x58\x5b\xc6\xb2\x58\x12\x96\xe4\x43\x71\x17\x65\xc9\xde\xfc\x41\x31\xec\xba\x2a\x00\x6a\x4d\xf5\x21\x25\xf3\xac\x76\xf2\x8d\x6a\x75\x38\xdf\xe7\xff\xfe\xf6\x4d\x06\xe6\x68\x03\xc1\xcb\x0b\x0c\x8b\x3e\x72\xc1\x51\xaa\x28\xa1\xd8\xa6\xc8\xab\x4e\xef\x74\xa2\xfb\x8c\x7a\x4d\xae\x7c\xd5\xaa\x6e\xda\xdd\x27\x6f\x43\x3d\x6f\x5b\xb4\x91\xfa\x9a\x26\xd0\x9a\xeb\x2e\x58\xa2\x4d\xf5\x55\x76\xd0\x1c\x27\xe7\xbf\xf3\x14\x3d\x85\x8a\x5d\xa0\xff\x9b\x24\xdc\xdc\x5c\xed\x43\x73\x5d\x47\x41\xde\xb3\x57\x93\x9c\xde\x05\x9b\xa4\x13\xf3\x5c\x26\x57\x9f\xc6\x1f\x7a\x5e\x4b\x14\xfc\x47\xd6\xab\x98\xfc\x9f\x45\xc9\x82\x07\x0a\x01\x17\x5d\x5b\xda\xa9\x5f\x9a\x7e\x5b\x34\xd3\xbf\x85\xf3\xe3\x78\x7c\x85\xb3\x6d\x71\x42\xfd\x11\xb5\xa6\x10\xde\x5d\xc4\xe9\x06\x73\x78\x38\x67\xe6\xd2\xee\xe6\xf9\xcd\x29\x29\x68\x6f\xaa\xbc\x08\x13\x50\x07\x05\x41\xe0\x0d\x41\xde\xa1\x70\x14\x7b\xca\x33\x5e\x18\x59\xc6\x59\xa1\x79\x33\xaa\x3c\xaf\x48\xcb\xfb\xb9\xe5\x05\x8f\x36\x46\x7b\x5e\xa0\xd0\x0e\xeb\x11\x6a\x31\x5b\x23\xf5\x68\xb5\x13\x75\xb6\x79\x2d\xd9\x85\x97\x7e\x82\xf6\x80\x6b\x02\x74\x80\x59\x0d\x1a\xad\x05\x61\xa8\x39\x7a\x98\x9d\xe6\x5e\x23\xd5\xec\xc4\x9b\xd9\x11\xb2\x5d\x98\x63\x5b\x07\xfb\xc1\xdf\x00\x00\x00\xff\xff\x5c\xb6\xd2\xe2\xdc\x04\x00\x00")

func defaultchannelpatternJsonBytes() ([]byte, error) {
	return bindataRead(
		_defaultchannelpatternJson,
		"DefaultChannelPattern.json",
	)
}

func defaultchannelpatternJson() (*asset, error) {
	bytes, err := defaultchannelpatternJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "DefaultChannelPattern.json", size: 1244, mode: os.FileMode(420), modTime: time.Unix(1550687842, 0)}
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
	"DefaultHttpPattern.json":    defaulthttppatternJson,
	"DefaultChannelPattern.json": defaultchannelpatternJson,
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
	"DefaultChannelPattern.json": &bintree{defaultchannelpatternJson, map[string]*bintree{}},
	"DefaultHttpPattern.json":    &bintree{defaulthttppatternJson, map[string]*bintree{}},
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
