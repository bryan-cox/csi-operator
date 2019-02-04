// Code generated by go-bindata.
// sources:
// pkg/generated/manifests/default-config.yaml
// DO NOT EDIT!

package generated

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

var _defaultConfigYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x54\xc1\x8e\xdb\x46\x0c\xbd\xeb\x2b\x88\xe8\x92\x00\xb6\x9c\xa0\x3d\x14\xbe\x15\x9b\x1c\x82\x06\x41\xb1\xc9\x0f\xd0\x33\xb4\x45\xec\x68\xa8\x90\x23\xb9\xfa\xfb\x82\x23\x79\xd7\x2d\xba\x45\x7b\xb2\x2d\x91\x7c\x8f\xef\x3d\xba\x85\xef\x3d\x1b\xb0\x41\xe9\x09\x22\x9d\x71\x4a\x05\x82\xe4\x33\x5f\x26\xc5\xc2\x92\xe1\xcc\x89\x76\x10\x64\x18\x39\x51\x04\xce\xb5\x56\x46\x52\x2c\xa2\x5d\xd3\xc2\xaf\x29\x81\x4d\xe3\x28\x5a\xbc\xa0\xd0\x60\x80\x4a\x90\xd8\xfc\xc1\x89\x92\x5c\x77\x50\xe4\x42\xa5\x27\x85\x2b\x97\xde\x67\xb0\x42\x24\x0b\xca\xa3\xe3\x74\x4d\xd3\xc2\xc7\x8d\x01\x0f\x78\x21\x83\x22\x30\x19\xc1\x59\x14\x66\x54\x96\xc9\x80\xfe\x28\xa4\x19\x13\xfc\x36\x9d\x48\x33\x15\xb2\xca\x4d\x32\xe5\x62\x5d\xb3\xed\xf0\xb9\x0e\x38\x36\x00\x58\x0a\x86\x9e\xb4\x3e\x39\xc2\x9b\x1f\x13\x2e\x1d\xcb\xe1\xe9\x17\x0b\xc6\x87\x60\xbc\xbf\x95\x1c\xe7\xf7\xdd\x4f\xdd\xfb\x37\x0d\xc0\xa8\x32\xb3\xb1\xe4\x7f\x6d\xbc\xab\x5a\x7b\x3f\x78\x6f\x54\x9e\x49\x1f\xe9\xc2\x56\x14\x5f\xeb\x5f\xab\xf6\x7a\x2b\xbb\x03\x4f\x3c\x53\x26\xb3\xdf\x55\x4e\xf4\x4a\xfb\xad\x66\xf4\x1a\xef\xfd\xd9\xc1\x9b\x16\xbe\xe0\x89\x12\x18\x25\x0a\x45\x14\x4a\x8f\x05\x06\x2c\xa1\x27\x83\x2c\x91\x0c\xae\x3d\x29\xc1\x28\xd1\x56\x2b\x82\xe4\xa2\x92\x12\xe9\x9d\x92\x4d\x0b\x6f\xef\xd6\xdb\x3d\xeb\xf8\xae\x5a\x8b\x29\xc9\x95\xa2\x3b\xa4\x53\xee\xe0\x7b\x4f\xb6\xcd\x2c\xcb\xc8\x01\x53\x5a\xa0\xc7\x99\xc0\x64\xa0\xa6\x85\x51\x79\xe6\x44\x17\x8a\x80\x21\x90\x55\x73\x3d\x48\x56\x44\xf1\x42\x70\xc2\xf0\x44\xb9\x4e\x5c\xa1\x0e\x91\xfc\x03\x30\x47\x08\x4a\x58\xe8\x10\x29\x51\xf1\x69\xb3\xa4\x69\x20\xab\xef\x1c\x00\x22\x8d\x49\x96\xc1\x99\x83\xd2\x8f\x89\x95\xc0\xa6\xd0\x6f\x94\x2a\x4b\x90\x0c\x91\x22\x07\xf4\x54\x56\x2d\x76\x4d\x0b\x26\x80\xf5\x57\xdd\x5e\x65\x60\xf3\xd0\x2e\x80\xd9\xd3\xa7\x10\xa5\x4a\x57\xd6\x7d\xd0\x47\xae\xb2\xf6\x68\xce\x2c\x52\x2e\x8c\xc9\x25\xfb\xe7\x9d\xba\x86\xf3\x59\xd1\x8a\x4e\xa1\x4c\x4a\x5f\x25\xd2\xb7\xcd\x20\x8f\x68\x0b\x9f\x86\xb1\x2c\x8e\xb9\xc5\x77\x57\xf9\x62\x5e\xaa\x57\x6e\xeb\xd7\x69\x38\x91\x82\x9c\x41\x69\x4c\x1c\xd0\xfc\xfb\x9d\x75\x75\xd1\x57\x2c\xf3\x23\x7d\xfb\x85\x30\x92\x42\xc5\xf5\xbb\xe6\x75\xa9\x97\xc3\xcd\x72\x7d\xe7\x07\x74\x53\xf2\x71\x03\x3a\xc2\x87\xca\x00\x07\x72\xcc\x87\x34\x59\x21\x7d\x94\x44\xab\x0c\x17\x9e\xc9\xe0\xe1\xdb\xe7\x2d\xf9\x2b\x15\xb7\xe6\xee\x46\x8d\x23\x05\xd4\xa6\xad\x9c\x91\xb3\x53\xc9\x32\x5d\x7a\x18\x49\x07\x36\x27\x7d\x73\xaa\x6b\xc2\x0b\x88\xe3\x1e\xc1\x16\x2b\x34\x1c\x65\xa4\x6c\x3d\x9f\xcb\xd1\xef\x6f\xc5\xfb\x8f\xe4\x9e\xff\x39\x6e\xb2\x54\x8a\x77\x82\x79\x4a\xff\x4a\x25\x0a\xa4\xbf\xab\x96\x81\x8b\x41\xc6\x81\x6c\xc4\x40\x5d\xb3\x56\x7c\xda\x0a\x1e\xfe\x07\xf1\xfd\x8b\x7d\xfb\x75\xca\xfe\x86\xe3\x3b\x3d\x8a\x14\xdf\xe9\x69\x3a\x79\xee\x21\x62\x41\x88\xac\x35\x37\x8b\xa7\x19\x53\x5a\x73\xdc\x35\x5b\x91\xf7\x7c\x64\x3d\xc2\x61\x46\x3d\x24\x3e\x1d\xb6\x17\xcd\x9f\x01\x00\x00\xff\xff\x28\x17\x94\x27\xeb\x05\x00\x00")

func defaultConfigYamlBytes() ([]byte, error) {
	return bindataRead(
		_defaultConfigYaml,
		"default-config.yaml",
	)
}

func defaultConfigYaml() (*asset, error) {
	bytes, err := defaultConfigYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "default-config.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"default-config.yaml": defaultConfigYaml,
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
	"default-config.yaml": {defaultConfigYaml, map[string]*bintree{}},
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