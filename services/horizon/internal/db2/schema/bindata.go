// Code generated by go-bindata.
// sources:
// .DS_Store
// latest.sql
// migrations/1_initial_schema.sql
// migrations/2_index_participants_by_toid.sql
// migrations/3_use_sequence_in_history_accounts.sql
// migrations/4_add_protocol_version.sql
// migrations/5_create_trades_table.sql
// migrations/6_create_assets_table.sql
// migrations/7_modify_trades_table.sql
// migrations/8_create_asset_stats_table.sql
// DO NOT EDIT!

package schema

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

var _Ds_store = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x98\x4d\x6a\x84\x40\x10\x85\x5f\x75\x5c\x34\xc9\xa6\x97\x59\xf6\x15\x72\x83\x46\xcc\x09\xbc\x40\x02\x01\x11\x12\x05\xf3\xb3\x76\x95\x73\xe5\x68\x41\xfa\x85\x08\xea\xac\x66\x18\x67\x78\x1f\x34\xdf\xc2\x2a\x95\x5e\x74\x57\x15\x00\x2b\x3f\x5f\x1e\x80\x00\xc0\x23\x1b\xdf\x58\xc5\x73\x2d\x70\xb4\xe5\x15\x80\x67\x7c\xa0\x45\x8f\x0e\xef\xeb\xef\x5a\x30\xe5\xde\xe2\x0d\x2d\x1a\x0c\xf3\xfc\xaf\x7a\xe8\x5e\xfb\xae\x61\x8c\x10\x42\x08\x21\x8e\x03\xef\x55\x7f\x77\xee\x1f\x11\x42\xec\x8e\xe9\x7c\x88\x74\xa2\xc7\x6c\xe3\x73\x47\x17\xb3\x9c\x40\x47\x3a\xd1\x63\xb6\x31\xce\xd1\x05\xed\xe9\x40\x47\x3a\xd1\x63\x36\x0f\x2d\x63\xf3\x61\xfc\xb2\xb1\x43\xb1\x40\x47\x3a\x9d\x66\x6f\x84\xb8\x74\x6e\xb2\xc2\x74\xff\x3f\x6e\xf7\xff\x42\x88\x2b\xc6\x8a\xaa\xae\xca\x03\x83\x36\xc7\x42\xe0\x89\x31\x3f\x7f\x89\x1b\x85\x80\xcb\x03\xc3\x7b\xfc\xc7\xa9\x18\x10\x62\x47\xfc\x06\x00\x00\xff\xff\x4d\xd6\x4f\xb3\x04\x18\x00\x00")

func Ds_storeBytes() ([]byte, error) {
	return bindataRead(
		_Ds_store,
		".DS_Store",
	)
}

func Ds_store() (*asset, error) {
	bytes, err := Ds_storeBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: ".DS_Store", size: 6148, mode: os.FileMode(420), modTime: time.Unix(1509392365, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _latestSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x5c\x6d\x6f\xdb\x38\x12\xfe\x9e\x5f\x41\xec\x17\x27\x80\x13\xc4\x69\x9c\xa4\xce\xb5\x80\x6b\xab\x5b\xa3\xae\xdc\xda\xca\x75\x8b\xc5\x82\xa0\x25\xda\xe6\x55\x12\x55\x92\x4e\x93\x3d\xdc\x7f\x3f\xe8\xcd\xd6\x0b\x29\x4a\xb6\xda\xee\xc7\x58\xa3\x99\xe7\x19\x0e\x39\xc3\x21\x95\xf3\xf3\x93\xf3\x73\xf0\x91\x72\xb1\x66\x78\xf1\x69\x0a\x1c\x24\xd0\x12\x71\x0c\x9c\xad\x17\x9c\x9c\x9f\x9f\x84\xcf\xc7\x5b\x2f\xc0\x0e\x58\x31\xea\xed\x05\x1e\x31\xe3\x84\xfa\xe0\xe5\xc5\xcd\x45\x3f\x23\xb5\x7c\x06\xc1\x1a\x86\xaf\x17\x44\x4e\x16\x86\x05\xb8\x40\x02\x7b\xd8\x17\x50\x10\x0f\xd3\xad\x00\xaf\xc0\xe5\x7d\xf4\xc8\xa5\xf6\xd7\xf2\xaf\xb6\x4b\x42\x69\xec\xdb\xd4\x21\xfe\x1a\xbc\x02\x9d\x07\xeb\xed\x5d\xe7\x3e\x55\xe7\x3b\x88\x39\xd0\xa6\xfe\x8a\x32\x8f\xf8\x6b\xc8\x05\x23\xfe\x9a\x83\x57\x80\xfa\x89\x8e\x0d\xb6\xbf\xc2\xd5\xd6\xb7\x05\xa1\x3e\x5c\x52\x87\xe0\xf0\xf9\x0a\xb9\x1c\xe7\xcc\x78\xc4\x87\x1e\xe6\x1c\xad\x23\x81\xef\x88\xf9\xc4\x5f\xdf\x27\xd8\x31\x62\xf6\x06\x06\x48\x6c\xc0\x2b\x10\x6c\x97\x2e\xb1\xbb\x21\x59\x1b\x09\xe4\xd2\x54\xcc\xc1\x2b\xb4\x75\x05\x14\x68\xe9\x62\x1e\x20\x1b\x87\xa0\x3b\x85\xa7\xdf\x89\xd8\x40\x4a\x9c\x0c\x8e\x93\x78\x34\x4c\xe4\xe1\x01\x40\x9c\x63\x01\x43\x77\xf1\x7b\x60\x3d\x07\x78\x00\xac\xe1\x9b\xa9\x71\x0f\x16\xf6\x06\x7b\x68\x90\x00\xb8\x07\xb3\xef\x3e\x66\x03\x70\x1e\x8d\xd6\x68\x6e\x0c\x2d\x23\x16\xcd\xea\x00\xa7\x27\x00\x00\x40\x1c\xb0\x24\x6b\xe2\x0b\x60\xce\x2c\x60\x3e\x4c\xa7\xdd\xe8\x77\xe4\xd1\xad\x2f\xe4\xcf\xfc\xad\x07\x91\x6d\x87\x02\x1c\x10\x5f\xe0\x35\x66\x05\x91\x95\x8b\xd6\x1c\x70\x0f\xb9\x6e\xf9\x7d\x41\x3d\x17\xd8\x1b\xc4\x90\x2d\x30\x03\x8f\x88\x3d\x13\x7f\x7d\x7a\x73\x7d\xb6\x13\x3c\x39\xbb\x3f\xc9\xd3\x5f\x53\x16\x40\x8f\xac\x19\x0a\x87\xec\x70\x17\x14\xf4\xec\xdd\x20\xf0\x53\xc9\x09\x41\xe0\x12\xec\x40\x24\x40\x18\x86\x5c\x20\x2f\x00\xe1\x38\x45\x7f\x82\xbf\xa9\x8f\xcb\x40\x37\x84\x0b\xca\x9e\x77\x1e\x82\xc4\x81\x1c\x7f\x4b\x01\x2f\x8c\x4f\x0f\x86\x39\xaa\x89\x39\x95\x56\x69\x8d\x60\x2e\xac\xe1\xdc\x02\x9f\x27\xd6\x3b\xd0\x8b\x7e\x98\x98\xa3\xb9\xf1\xc1\x30\x2d\xf0\xe6\x4b\xf2\x93\x39\x03\x1f\x26\xe6\xbf\x87\xd3\x07\x63\xf7\xf7\xf0\x8f\xfd\xdf\xa3\xe1\xe8\x9d\x01\x7a\x3a\x32\x07\xbb\xbd\xa8\xa8\x14\x7e\x63\xe3\xed\xf0\x61\x6a\x01\x1f\x3f\x89\x47\xe4\x9e\x76\x14\x8c\x3b\x83\x01\xc3\x6b\xdb\x45\x9c\x9f\x15\x87\xcb\x71\x18\xe6\x5c\x1e\x5a\x15\x03\x15\x4e\x8a\x16\x98\x45\x6a\xf6\xbc\xe4\x13\x23\x9e\x81\xe2\x39\xc0\x9a\x19\x90\x15\xb7\xa9\x23\x13\xef\x5d\xc9\xc5\x09\xe7\x5b\xcc\x24\x2f\xf4\x6f\xaa\x66\x58\x9e\x48\xcb\x61\x9b\xd5\xf9\xd3\x82\xb6\x8a\x08\x98\x7d\x36\x8d\x31\x78\xf3\x45\xc3\x68\x38\xb5\x8c\xb9\x86\xd0\x4e\x57\xe1\xf1\x05\x71\x54\xd8\xf0\x6a\x85\xed\x16\xa2\x2e\xd1\x93\x84\x5d\x61\xce\x40\xd5\xea\x9e\xca\xd1\x00\xc7\xeb\xa0\x52\xf2\x37\xca\x1c\xcc\x7e\x53\x44\x73\x14\xc7\xf2\x47\x0e\x16\x88\xb8\x1c\xfc\x87\x53\x7f\xa9\x0e\x36\x17\x3b\x6b\xcc\x8e\xf7\x43\xa2\x27\xf1\x03\xc7\xdf\xb6\xd8\xb7\x55\xd8\x62\x61\xb8\x41\x7c\x53\x6b\x16\x06\x0c\x3f\x12\xba\xe5\x50\xfb\x62\xe2\x16\x86\x7c\x8e\xe2\xda\x22\x1a\x88\x1d\x8e\x74\x95\xbb\x2c\x58\xd8\x0f\x44\x3d\x79\xdb\xa5\x5c\x96\x98\xc2\x4a\x69\x97\x9b\x8a\xef\x30\x8c\x84\xf6\xa5\x58\x76\x1b\x38\xb5\x65\x77\xa1\x93\xfc\xe9\x05\x94\x09\xcc\x60\x5a\xec\x15\xb9\xf4\x4a\xe5\x80\x40\x2e\xb4\x29\xf1\xb9\x3c\x06\x57\x18\xc3\x80\x52\x57\xfe\x34\xac\x3d\xe1\x0a\xab\xc6\x3a\x7a\xcc\x30\xc7\xec\x51\x25\xe2\xa1\x27\x28\x9e\x60\x54\x1a\x91\xbf\x55\x52\x01\xa3\x82\xda\xd4\x55\xf2\xba\xac\xb1\xb6\xee\xc7\x39\x40\x4c\x10\x9b\x04\xa8\x8d\xac\x2a\x57\xab\xcb\x45\xf5\x57\x01\xfd\xba\xd2\x94\x72\xbb\xe9\xa5\xd2\xc6\xcf\x4a\x37\x8d\x88\x1e\x99\x7e\x2a\x6d\x95\xd3\x91\x5c\xbc\x22\x3d\xed\x5e\x68\x31\x36\x75\x5b\x8e\xec\xaa\xa9\xdc\x96\x84\x15\xb9\x1d\x53\x89\x32\xd3\x91\x89\x29\xfe\x89\xd3\x2d\xb3\x71\x1a\xdd\x8a\x94\x90\x4e\xf3\x4e\x67\x30\x28\x49\xd4\x98\x07\x82\x21\x07\x1f\xef\xce\x58\x4d\x21\xdf\x1f\x9b\xc7\x93\xbc\x76\x48\x56\xa1\xab\x15\x66\x4a\xb3\xd1\xea\xab\xab\x46\x62\xa1\xb8\x74\xad\x14\xa9\xd8\x93\x46\x16\x30\xd3\xda\xda\xc9\x55\x9a\xdb\x49\x55\x58\x8c\x20\x11\x0e\x39\x76\x5d\xcc\xc0\x92\x52\x17\x23\x3f\x7e\x36\x9a\x99\x0b\x6b\x3e\x9c\x98\x56\x61\xdc\x60\x86\x08\x8c\xba\x10\x60\xf4\xce\x18\xbd\x07\xa7\xa7\x59\x8a\xaf\xc1\xe5\xd9\x99\x4e\x95\xec\xf5\x94\xd5\xbf\x4a\x44\x6b\xe8\xcb\x91\x2e\xa8\x2f\x78\x24\x02\x58\x19\xeb\xbb\xa9\xdc\x6a\xa2\x53\x29\xae\x9b\xea\xea\xac\x31\xc7\x24\x3b\x15\xbe\x76\xd3\x9d\xc6\xca\xcf\x4a\x78\x0d\xc9\x1e\x99\xf2\x34\xd6\xca\x49\x4f\xf5\x42\x45\xda\xcb\xbc\xd2\x6a\xac\xa6\xf1\x99\x85\x54\x7b\xf7\x91\x2c\xce\x9a\x3d\x4d\xdd\xcc\x58\x9d\xe4\xa4\xb2\x7b\xd3\xea\xf2\x1c\x29\xa7\x9e\x6a\x6b\xf3\x4b\x36\x27\xe2\x09\x62\xff\x11\xbb\x34\xc0\xb2\x86\x9f\x78\x0a\xb7\x0a\x5b\x57\x28\x1e\x7a\x58\x20\xc5\xa3\xd0\x0b\xaa\xc7\x9c\xac\x7d\x24\xb6\x0c\xcb\x7a\x53\x2f\x6f\xce\xfe\xfc\x6b\x5f\x5d\xfc\xf7\x7f\xb2\xfa\xe2\xcf\xbf\x8a\x7b\x16\xec\x51\x45\x1b\x69\xaf\xcb\xa7\x3e\xae\xac\x56\xf6\xba\xca\x6a\x12\x66\xc4\xc3\x70\x49\xb7\xbe\x13\xb5\x7a\xef\x18\xf2\xd7\x55\x4d\xcf\xb8\x09\x46\x9c\x74\xf6\x24\x58\x6a\x4d\xf9\x78\xfa\xcc\xcc\x69\xb1\x8f\x02\xe2\xe7\xa3\xd9\xf4\xe1\x83\x19\x0e\xe9\xc2\xb0\x2a\x1a\x86\xd9\xd6\x4c\xb6\x5d\xd8\xac\x70\x6f\x8f\x84\x42\x7f\x23\x52\x95\x05\x7f\x1d\x92\xca\xcc\xd9\x1a\x4d\xa5\x85\x46\x44\x35\xcb\xbc\x9c\xea\x18\x09\x04\x56\x94\x69\x8e\x4a\xc0\x78\x68\x0d\x35\xf4\x14\x2a\xab\x8e\x1f\xea\xa8\x9d\x98\x0b\x63\x6e\x81\x89\x69\xcd\x4a\x47\x10\x51\xc2\x5d\x80\xd3\x4e\x0f\x12\x9f\x08\x82\x5c\xc8\x23\x5d\x17\xfc\x9b\xdb\xe9\x82\xce\xd5\x65\xef\xf6\xbc\xd7\x3b\xef\x5d\x83\xde\xe5\xa0\x7f\x3b\xb8\xbe\xbc\xe8\x5f\x5f\xf5\xaf\x6e\xce\x2f\xef\x3a\x67\xf7\xf5\xb4\x5f\x41\xe2\x3b\xf8\x29\xef\xd5\xe5\x33\x14\x94\x38\xd5\x96\xfa\x77\xfd\x46\x96\x5e\xc0\x2d\xc7\xbb\xac\x01\x89\x0f\x8b\xcd\xfc\x6a\x7b\xb7\x2f\x2f\x6f\x9b\xd8\xbb\x86\xc8\x71\x60\xb1\x41\x53\x69\xa3\x7f\x7d\x75\xdb\x6b\x62\xa3\x0f\xe3\x14\x95\x56\xcb\xd1\x61\x5e\xb5\x89\x97\xd7\x37\xfd\x26\x26\x6e\x52\x13\xc9\x0a\xa6\x37\x71\x73\x75\xd7\x6b\xe4\xa9\x5b\xe8\x51\x87\xac\x9e\xeb\xb3\xb8\xed\xf5\x5e\x34\x32\x71\x97\x63\x11\xcf\xc2\x3a\x76\xae\xaf\xfb\xa9\x1d\xc5\x1c\xac\x3c\x8c\x6a\x32\xb7\x1b\x1d\xd4\x85\xcb\x95\x46\xef\xc2\x98\x1a\x23\x2b\x73\xf0\x7b\xc1\x71\xf5\x21\x56\x17\xf4\xba\xf1\x29\x6f\x0d\xba\xe5\xf3\xa9\x23\xc8\x56\x9e\x89\xb4\x42\x35\x97\x7e\x9b\x10\x95\x9d\x89\x1c\xb1\x64\x57\x1d\x31\xb4\xa0\xb6\x46\x2b\xf7\xf0\x61\x6a\xd6\x4b\x6c\x63\xd8\xaa\x0b\x8c\x26\xc3\xa8\xe8\x1d\xb6\xe0\x72\x49\x0b\xad\x1d\xad\xfa\x66\xc5\xe1\x43\xd9\x74\x97\xdc\xc6\x60\xea\x8a\xa8\x26\xc3\xa9\xdc\x13\x37\x77\x49\xf6\xfa\x49\x36\x3b\x04\x5f\xf1\x73\xaa\x7a\xdf\x9f\x6a\x5a\x87\x66\x34\x46\x5b\x97\xe1\x78\x9c\xed\x76\x15\x0d\x82\x8f\xf3\xc9\x87\xe1\xfc\x0b\x78\x6f\x7c\x01\xa7\xc4\xd1\xdd\x38\x29\xfe\xdd\x12\xea\x82\x56\x19\x72\x99\x61\x2d\xfa\xc2\x0e\xaa\xb0\x3a\xef\xef\x15\xc0\xfd\x8d\x04\x98\xbd\x3e\x00\x5b\x61\x97\x37\x2b\x23\x77\x10\x30\xf0\x60\x4e\x3e\x3d\x18\xe0\x74\x2f\xde\xcd\x5c\xad\xe8\xe6\x2e\x42\x34\x74\x4d\x3b\xc3\xda\x98\x78\xa3\x41\x55\xec\x28\x35\x6b\x79\xbb\xcc\xe4\x46\xaa\x98\x56\xc0\xaa\xcd\x5c\xb9\xc9\xd4\x2e\x7d\xed\xb2\x57\x99\xa9\xe2\x5f\x09\x4d\xeb\x81\x38\xa4\x97\xcf\x51\xb4\xa7\x44\x26\xe6\xd8\xf8\xa3\x5e\x73\x32\x12\xcd\x6b\x01\x33\xb3\x38\x19\x1e\x16\x13\xf3\x77\xb0\x14\x0c\xe3\xec\xec\x52\xa3\x89\xe7\xd8\xf1\x78\x92\x4b\x4b\xb5\x10\x29\xe6\xf5\x72\x57\x67\x1f\x0c\x67\xaf\x22\x8b\x24\xd7\xc9\xcd\xe3\x89\x85\xbb\xa5\x56\xa9\x0c\xdc\x06\xf1\xcd\x31\xc8\xa2\x8e\x71\x2d\x58\xc5\x3e\xb3\x0c\x4d\x5c\x16\x1f\x83\x27\xd6\x50\x0f\x51\xa1\x89\xdd\x2d\xf7\xab\xa5\x53\x1e\xe2\x30\x36\xa2\xe7\x07\x20\x4d\xb2\x44\x0c\xb8\xa0\x2e\x0b\x3b\xbd\x44\x95\x43\x2c\x3b\x5b\xed\xa6\xe7\xa8\x2a\xb0\xfb\x66\xda\x91\x30\x89\x53\x1b\xe0\xfe\x9c\xaa\x2b\x3d\x10\xd6\x80\xa6\x01\x0c\xda\xc2\x9d\xe8\xca\x42\x57\xa4\xaa\x83\x98\xc8\x09\x88\xa7\xf6\x08\x24\xba\x14\x31\x7d\x20\x85\xfc\xa1\x63\x99\x04\x0d\xc2\xa8\xdc\xd0\x83\x38\x24\xe0\xf7\x3a\x0e\x75\x7e\xb5\xa3\x77\x77\xdf\xc2\xa5\xfa\x78\x5f\xe7\xd5\x65\x21\xa7\x17\xf9\x72\x18\xe5\x88\xb2\x7e\x6d\x0b\x56\x49\x67\xbd\xe5\x4d\x06\x50\xc4\x43\x22\x8e\x19\xd6\xbd\x8e\xc3\x43\x52\x17\x7e\x82\x39\xd1\xaa\xb8\x5a\x1d\x91\x0e\x72\x5a\x0a\x58\x1d\x5c\x40\x96\xde\x19\x91\x63\x49\xaf\x1a\xb8\x94\x7e\xdd\x06\xc7\x21\xca\xeb\xd2\xe1\x2a\xdd\x99\x90\xe2\x0b\x10\x61\xd1\x67\x29\xad\x20\x2c\x6a\xd3\x61\xcc\xdd\xf3\xe8\x96\xae\x79\x74\x4b\x77\x79\x14\x24\x5a\x98\x2d\x89\x1e\x1d\xe2\x86\x39\x29\xd4\xda\x9a\x77\x1b\x38\x56\xeb\xb7\xf8\xf4\xa4\xd4\xd1\xa5\x3e\x4c\x3e\x3c\x38\xd6\xa1\x5a\x03\xb9\xea\x38\xfd\x90\x22\x5f\x8f\xc6\x82\x0d\xb0\x1f\x1f\x07\x55\xba\xf5\x88\x25\xb3\x2c\xaf\x30\xa9\x7d\x42\x7d\xe1\xde\xfe\xe0\x78\xa8\xd4\xaa\x2d\xb6\x42\x21\x0d\xd0\x24\x73\x85\x2a\x77\x41\xd4\x12\x5a\x99\x6a\x6d\xd2\xac\x1b\xc9\x19\xe5\x6d\x07\x43\x4e\xf5\x21\x59\x5e\xad\xae\x70\xcb\xbc\x7d\x47\x97\xee\xb1\x6b\xe1\x17\x5e\xa8\x4f\x26\xf3\x59\xc1\x0f\xf3\x7f\xf6\xd3\x05\x1d\x93\x8c\x6c\x7d\x12\xb2\x8f\x24\x7e\x18\x1b\xe9\x17\x19\x3a\x5a\xb2\x97\xea\xf3\x4b\xb7\xae\x3f\x8c\xd3\xee\x96\x95\x8e\x87\xb2\xc7\x90\x57\xbd\x3f\x87\xf9\x11\x53\xbb\xa8\x5d\xba\xed\x68\x3a\xc1\xf3\x4a\xf3\x85\x6b\x4b\x33\xbc\xca\x44\x1d\x0e\x9a\x6a\xba\xd2\x58\x7b\xe9\xab\xac\xb8\x16\x76\x7d\x12\xcb\x6e\x71\x7e\x44\xd8\x94\xf5\x1f\xbc\xc1\x8a\x8a\xb8\x5d\x22\x4f\xfb\x3a\x70\x49\xe9\xd7\x83\xbd\x5c\xa1\x53\x5b\x22\x9c\x9e\xa6\x9f\x16\x9c\xbf\x7e\x0d\x3a\x9c\xba\x4e\xe6\x0c\xa3\x33\x18\x08\xfc\x24\xce\xce\xba\x40\x2d\x68\x53\xa7\x9e\x60\xdc\x01\x55\x8b\x2e\xe9\x76\xbd\x11\xb5\xcc\xe7\x44\xab\x01\xe4\x44\x0b\x10\xce\xc0\xe7\x77\xc6\xdc\x88\x83\x0c\xbc\x02\x2f\x5e\x28\x7a\xc6\xe5\xe3\x3f\xe2\xc0\x55\xa6\x39\xff\xf6\xfd\xcf\x39\x04\x4c\xcc\x82\xb7\xb3\xb9\x31\xf9\xdd\xdc\x35\xde\xc1\xdc\x78\x6b\xcc\x0d\x73\x64\x2c\x0a\xbd\xe8\xe8\xe9\xcc\x04\x0f\x1f\xc7\x61\xc8\xcc\x8d\x85\x35\x9f\x8c\xac\xf0\xa7\xb1\x31\x35\x2c\x03\x8c\x86\x8b\xd1\x70\x6c\x54\x7f\x03\x22\xff\x26\x60\xd7\x38\x6a\xcf\x19\x79\x3b\x9a\xa3\x09\x15\x92\xbc\x7f\x0a\x12\x72\x67\x25\x85\xbe\xe6\x1c\x47\xe9\x89\x64\x2b\xfb\xcb\xfd\x90\xc5\x21\xf3\x42\xda\x25\xa8\x0e\x98\x66\x1e\x28\x7f\xc7\xf2\x0b\xdd\xa0\x00\x93\xf7\x45\x59\xa8\xe5\xa0\x28\xb6\x38\xfe\x09\x0e\x51\x87\x46\xa9\x87\x54\x37\x3a\x54\xff\x73\x05\xd8\xd4\x0b\x5c\x2c\x70\xc4\xe1\xff\x01\x00\x00\xff\xff\xfb\x8b\x92\x21\xa0\x45\x00\x00")

func latestSqlBytes() ([]byte, error) {
	return bindataRead(
		_latestSql,
		"latest.sql",
	)
}

func latestSql() (*asset, error) {
	bytes, err := latestSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "latest.sql", size: 17824, mode: os.FileMode(420), modTime: time.Unix(1510685860, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrations1_initial_schemaSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x5a\x5f\x6f\xdb\xc8\x11\x7f\xf7\xa7\x18\xdc\x8b\x6c\xd4\x6a\x2f\xb8\xe2\x70\x95\xe1\x03\x14\x99\x69\x84\xca\x54\x22\x51\x4d\x82\xc3\x61\xb1\x22\x47\xd4\xd6\xe4\x2e\xb3\xbb\x74\xa4\x2b\xfa\xdd\x0b\x52\x24\xc5\xff\xa4\x1c\xc9\xf7\x28\xee\xec\xcc\xfc\x66\x66\x7f\x33\x5c\x6a\x38\x84\xbf\xf8\xcc\x95\x54\x23\xac\x82\xab\xe1\xf0\x6a\x38\x84\x0f\x42\x69\x57\xe2\xf2\xe3\x0c\x1c\xaa\xe9\x9a\x2a\x04\x27\xf4\xe3\xe5\xab\xa5\x61\x81\xd2\x54\xa3\x8f\x5c\x13\xcd\x7c\x14\xa1\x86\x7b\xf8\xf1\x2e\x5e\xf2\x84\xfd\x54\x7d\x6a\x7b\x2c\x92\x46\x6e\x0b\x87\x71\x17\xee\x61\xb0\xb2\xde\xfd\x32\xb8\x4b\xd5\x71\x87\x4a\x87\xd8\x82\x6f\x84\xf4\x19\x77\x89\xd2\x92\x71\x57\xc1\x3d\x08\x9e\xe8\xd8\xa2\xfd\x44\x36\x21\xb7\x35\x13\x9c\xac\x85\xc3\x30\x5a\xdf\x50\x4f\x61\xc1\x8c\xcf\x38\xf1\x51\x29\xea\xc6\x02\xdf\xa8\xe4\x8c\xbb\x77\x57\x09\x3c\x93\xfa\x38\x82\xc0\x0b\x5c\xf5\xd5\xbb\x03\x6b\x1f\xe0\x08\x8c\xcf\x96\x61\x2e\xa7\x73\xf3\x0e\x96\xf6\x16\x7d\x3a\x82\xe1\x1d\xcc\xbf\x71\x94\x23\x18\xc6\xc8\x27\x0b\x63\x6c\x19\x47\x49\x98\xbe\x03\x73\x6e\x81\xf1\x79\xba\xb4\x96\xa9\x42\xf8\x34\xb5\xde\xc3\x72\xf2\xde\x78\x1c\x43\xe0\x12\x9b\x6a\xea\x89\xc8\x7a\xc1\xfc\x51\x4b\xc9\x91\xc9\xfc\xf1\xd1\x30\xad\x16\x37\x0e\x02\x30\x37\xab\x4a\x60\xba\x84\xc1\x87\xd9\xdf\x02\x37\x4a\x5e\x20\x85\x8d\x4e\x28\xa9\x07\x1e\xe5\x6e\x48\x5d\x1c\x94\xfd\xd8\x2a\x2d\x24\x9e\x2f\x0a\x07\x7d\xc5\x20\x84\x6b\x8f\xd9\xcd\x01\x28\xba\xf0\x32\xfc\x89\xd9\x08\x7e\x54\xb2\xa0\xf7\x01\xc2\x46\x48\x88\x9e\x47\x15\xa7\x50\x2b\x10\x1b\xb8\x7e\xc2\xfd\x2d\x3c\x53\x2f\xc4\x1b\x08\x28\x93\x2a\x0e\x49\x5c\x86\x48\xa5\xbd\x25\x01\xd5\x5b\xb8\x4f\xbc\xbe\x2d\xa6\x30\x12\x73\x70\x43\x43\x4f\x13\x4d\xd7\x1e\xaa\x80\xda\x18\x95\xf3\xa0\xb4\xfa\x8d\xe9\x2d\x11\xcc\xc9\x55\x68\x31\xee\x2c\xf2\x6c\x4f\xa8\x6d\x8b\x90\x6b\x95\xc2\xb7\xc6\x6f\x67\xc6\x11\x7c\x12\xbb\x2c\x02\x77\x60\x65\x66\x47\xf9\x7c\xc4\xfb\x2a\x5a\xe1\xfa\x0a\x00\x80\x39\xb0\x66\x2e\xe3\x3a\xce\x94\xb9\x9a\xcd\x6e\xe3\xe7\xd4\x71\x24\x2a\x05\xf6\x96\x4a\x6a\x6b\x94\xf0\x4c\xe5\x9e\x71\xf7\xfa\xe7\xbf\xdf\x5c\xdd\x54\x6a\x25\xd1\x8e\x9b\x0d\xda\xe7\x76\x39\x51\x9a\x78\x5c\x02\x42\x9a\x10\xa4\x72\x22\x40\x49\x63\x5e\x68\x92\xfc\x41\x48\x07\xe5\x0f\xc0\xb8\x46\x17\x65\x69\x35\xae\x97\xfa\x25\x07\x35\x65\x9e\x82\xff\x28\xc1\xd7\xcd\x41\xf1\xd0\x71\x51\x9e\x39\x28\x89\xd2\x24\x28\x0a\xbf\x86\xc8\xed\x26\x47\x0f\xc2\x64\x4b\xd5\xb6\x3e\xa3\x25\xf9\x40\xe2\x33\x13\xa1\x22\x9d\x1b\x93\x18\x49\xca\x15\x3d\xb0\x6f\x9c\x95\xcc\x8f\x07\xe3\xdd\x78\x35\xb3\xe0\xc7\x92\x85\x63\x56\xfa\xc9\xdb\x9e\x50\xe8\x10\xaa\x21\xea\x20\x4a\x53\x3f\x80\xe8\x20\x45\xbd\x24\x7a\x02\x7f\x08\x8e\xe5\x3d\x12\xa9\xee\xdc\x74\x90\x0d\x03\xa7\xb7\x6c\x56\x47\xc9\x4f\x3f\x10\x52\xa3\x24\xcf\x28\x15\x13\xbc\x82\xe5\x4d\xb9\xa2\x84\xa6\x1e\xb1\x05\xe3\xaa\xbe\x20\x37\x88\x24\x10\xc2\xab\x5f\x8d\x9a\x2e\xd9\x60\x53\xae\xe3\x65\x89\x0a\xe5\x73\x93\x88\x4f\x77\x44\xef\x88\x42\x4d\x14\xfb\xa3\x2a\xd5\x5c\xca\xc7\xb4\x05\x54\x6a\x66\xb3\x80\x9e\x9d\xa1\xea\x6d\x1c\xf9\xaa\x1e\x53\xff\xe3\xde\x4d\x20\xa7\xe2\x27\xcc\x21\x0a\xbf\xa6\x61\x58\x1a\x1f\x57\x86\x39\x69\x89\x44\x1e\x7c\x2a\xdd\xcf\x46\x8c\x60\x69\x8d\x17\xd6\xa1\x91\xbe\x89\x1f\x4c\xcd\xc9\xc2\x88\x5b\xdf\xdb\x2f\xc9\x23\x73\x0e\x8f\x53\xf3\xdf\xe3\xd9\xca\xc8\x7e\x8f\x3f\x1f\x7f\x4f\xc6\x93\xf7\x06\xbc\x39\x0b\x50\x98\x7f\x32\x8d\x07\x78\xfb\xa5\x03\xf1\x78\x66\x19\x8b\x13\x01\x67\xba\x3b\xc4\xff\xca\x9c\x4e\x2c\x97\x2a\xd4\xae\x66\x9a\xa7\xc7\xc6\x86\x1b\x04\x1e\xb3\x0f\xb8\xe2\x7e\xf4\x9d\xed\xe8\xf0\x48\x89\x50\xda\x98\x96\x7a\x03\xf7\xa7\x3c\x35\x18\x8c\x46\x15\x89\x1e\x87\x22\x0f\xef\x72\xb4\xd0\x64\x25\x8e\x7d\x03\x2d\xd4\xed\xad\x4f\xc0\xf7\x90\x42\x93\x67\xe7\xa5\x85\x0e\x2b\xaf\x45\x0c\x27\x82\xfd\x4e\x6a\xe8\xb0\x56\x25\x87\xa6\x0d\x2d\xf4\x90\xdb\x72\xb9\x92\x4d\x29\x22\xef\x5f\xef\x71\x2c\x99\xc2\x3a\x86\xbc\xbe\x0c\xd2\x4e\x06\xb5\xb2\x47\xd3\xcd\xf3\x0a\x6d\x6c\xcd\x4d\xb3\xde\x9f\x32\xad\xe9\x1d\x41\xfe\x8c\x9e\x08\x10\x34\xee\x2a\x54\xbd\x8b\x66\xa7\xd0\xd3\x0d\x8b\x3e\x46\xaf\x90\xb5\x4b\x51\x14\x9a\x96\x15\x73\x39\xd5\xa1\xc4\xba\x37\xaa\x7f\xfc\x7c\xf3\xdb\xef\x47\x16\xfe\xef\xff\xea\x78\xf8\xb7\xdf\xcb\x43\x1c\xfa\x82\xc4\xdd\xa0\xca\xd9\x99\x2e\x2e\x38\xb6\xb2\xfa\x51\x57\x55\x4d\x82\x8c\xf9\x48\xd6\x22\xe4\x8e\x8a\x32\xf7\x8b\xa4\xdc\xc5\x98\x0c\xf3\x87\x89\x39\xe9\xd1\x49\x6c\xf7\x3a\xef\x87\xe3\x32\x37\x67\x5d\xdd\x1d\x0e\xf2\x93\xf9\x6c\xf5\x68\x46\x29\x8d\x5e\xa8\x53\x94\x1c\x77\xfa\x99\x7a\xd7\x83\x5e\x03\xc5\x60\x34\x92\xe8\xda\x1e\x55\xaa\xc2\xe8\x67\x43\xd1\xd8\xac\x4e\xc2\xd1\xc1\x7e\x6d\x48\x3a\x42\x11\x3c\xe1\xfe\x78\xad\x62\x2e\xad\xc5\x78\x6a\xb6\xa0\xad\x12\xde\x89\x09\x8c\x4b\x69\xfc\xf0\x90\xb3\xd6\xc7\x47\xf8\xb0\x98\x3e\x8e\x17\x5f\xe0\x5f\xc6\x17\xb8\x66\xce\xe9\x3d\xf8\x82\x48\x9b\x6c\xb6\x61\x6d\xf5\xb3\x13\xed\x3a\x1b\x50\x52\x48\x53\xf3\xc1\xf8\xfc\x82\x46\x15\xef\xcb\xe9\x83\xb9\x59\xdf\xb6\x56\xcb\xa9\xf9\x4f\x58\x6b\x89\x08\xd7\x89\xf0\x6d\xa5\x2f\xd4\x79\x1a\xb5\xb7\xb3\xb9\x19\xf7\xca\x5e\x3e\x96\x3b\x6c\x9d\x6b\x87\x86\x7a\x36\xe7\x0e\xea\xfa\xb9\x57\xea\xe5\xb7\xd5\xb6\x5d\x5b\xe3\x04\xc9\x7a\x7f\x58\xff\x5e\xb7\x57\xe6\xf4\xe3\x2a\xf5\xbe\xa4\x3b\x8f\x21\xbd\x76\x2b\xb8\x5f\xf7\x9a\x7d\x9b\xde\xa0\x35\x79\x7e\xa4\xd5\x73\xfa\xcc\x9c\xde\xde\x1e\xa7\xfa\xdb\xda\x8b\x82\x0e\x04\x22\x20\xc1\x45\x40\x24\x8a\xf3\x38\x1a\xfa\xdf\x8b\x60\x55\xd1\x64\x37\x7a\xeb\xfd\xd9\x01\x15\x75\xe7\x31\xa5\x77\x95\x05\x10\xf5\xee\xe5\x4f\xef\x45\x7c\xac\x18\xe8\x77\x6c\x6b\xbc\x65\xdc\xc1\x1d\x29\xdf\xab\x13\xc1\x49\x72\x79\x7e\x56\xd7\x3b\xad\xe5\x71\x64\x97\xfc\x45\xf6\x3e\x08\x9e\x00\xe4\xcc\xe1\x6f\x33\xd4\xed\x7e\x67\x0a\x12\x0a\x88\xf4\x45\x73\xf1\x79\xe8\xbd\xd5\x44\x27\x01\x45\x42\x1d\x5e\x27\x87\x23\x52\x99\x5d\x72\x5f\xc2\xf5\x3a\x3b\x9d\x87\x34\x93\xec\x0f\xe2\xa2\x35\x53\xb0\xf3\x12\x8a\x69\x56\x57\xba\xc5\xbf\x70\x0a\x2a\x1f\x0d\x3a\xb1\x94\x36\xf4\x47\x96\xfb\x86\xf3\x3a\x99\xc9\x7f\x34\xea\x82\x95\x93\xed\x8f\xa8\xee\xf3\xd4\xeb\x40\xab\xfd\x30\xd6\x85\xb1\x6e\x53\x7f\xb0\xe9\xa4\xf8\x3a\x00\xb3\x8b\x9e\x2e\x50\x8d\x93\x7f\x51\xf5\xf1\x8e\xfc\xe2\xdc\x50\x36\x55\x3b\x55\x9d\xca\x10\x45\xa5\xc5\x7b\xe4\x4b\x50\x44\x9b\xbd\x3e\x80\x8a\x3b\x4e\x03\x77\xa1\x9e\x59\xb5\xd2\x0b\x48\x5d\xe7\x8c\x87\x66\xbd\xbb\xd0\x34\x9e\x28\x6e\x18\x08\x5f\x38\x8f\x57\x13\xd2\x9c\x8f\xfc\xf8\x79\xf1\xe3\x52\x35\xf6\xe2\x49\x58\x4b\xea\x60\x36\x1b\xa5\xef\x92\x64\x2d\xc4\xd3\x79\x0a\xaa\xc5\x40\xe7\x08\x76\x7d\x9d\x7e\x17\x1b\xfe\xfa\x2b\x0c\x94\xf0\x1c\x42\x95\x42\x1d\x97\xe2\x60\x34\xd2\xb8\xd3\x37\x37\xb7\xd0\x2c\x68\x0b\xa7\x9f\x20\x53\x2a\x44\xd9\x2c\xba\x16\xa1\xbb\xd5\xbd\xcc\x17\x44\xdb\x1d\x28\x88\x96\x5c\xb8\x81\x4f\xef\x8d\x85\x71\x38\x4f\x70\x0f\x3f\xfd\x94\xcb\x5e\xd3\xbf\xf9\xc0\x16\x7e\xe0\xa1\xc6\x38\x13\xf9\x3f\x02\x3e\x88\x6f\xfc\xca\x91\x22\x80\xf8\x3f\x4e\xf5\xe5\x62\x53\x65\x53\x07\xef\x3a\x04\x8b\x07\xaa\x6d\x53\x8e\x23\x7a\x89\xf5\xd7\x9c\xb6\xb6\x36\x99\xb4\xaa\xda\x64\xb2\x37\x96\x4c\xe8\xff\x01\x00\x00\xff\xff\x5d\xb2\x1f\x7d\x3f\x29\x00\x00")

func migrations1_initial_schemaSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations1_initial_schemaSql,
		"migrations/1_initial_schema.sql",
	)
}

func migrations1_initial_schemaSql() (*asset, error) {
	bytes, err := migrations1_initial_schemaSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/1_initial_schema.sql", size: 10559, mode: os.FileMode(420), modTime: time.Unix(1508347796, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrations2_index_participants_by_toidSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x8f\xb1\xca\xc2\x50\x0c\x46\xf7\x3c\x45\xc6\xff\x47\xfa\x04\x9d\xc4\x16\xe9\xd2\x4a\xb5\xe0\x76\x49\xdb\x8b\xcd\xe0\xcd\x25\x37\x20\x7d\x7b\x41\x07\x5b\xbb\xb8\x86\x8f\x73\x72\xb2\x0c\x77\x77\xbe\x29\x99\xc7\x2e\x02\x1c\xda\x72\x7f\x29\xb1\xaa\x8b\xf2\x8a\x93\x44\xd7\xcf\x6e\x12\x1e\xb1\xa9\x71\xe2\x64\xa2\xb3\x93\xe8\x95\x8c\x25\xb8\x48\x6a\x3c\x70\xa4\x60\x09\xbb\x73\x55\x1f\xb1\x37\xf5\x1e\xff\xb6\x5b\x1e\xff\xf3\x2f\xbc\xbd\xf1\xb6\xc6\x9b\x52\x48\x34\xfc\x28\x58\xae\x5f\x0a\x58\x26\x15\xf2\x08\x00\x45\xdb\x9c\xb6\x49\xf9\xea\xfe\xf9\x25\x87\x67\x00\x00\x00\xff\xff\x33\xec\x54\x7a\x15\x01\x00\x00")

func migrations2_index_participants_by_toidSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations2_index_participants_by_toidSql,
		"migrations/2_index_participants_by_toid.sql",
	)
}

func migrations2_index_participants_by_toidSql() (*asset, error) {
	bytes, err := migrations2_index_participants_by_toidSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/2_index_participants_by_toid.sql", size: 277, mode: os.FileMode(420), modTime: time.Unix(1508347796, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrations3_use_sequence_in_history_accountsSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x91\x4d\x6b\xb3\x40\x14\x85\xf7\xf3\x2b\xce\x2e\xca\xfb\x66\x91\x6d\x5c\x4d\xc6\x1b\x22\x8c\x63\x3b\x5e\xdb\x64\x25\xa2\x43\x3a\x90\x6a\xeb\xd8\xaf\x7f\x5f\x48\xd3\x0f\x08\x6d\xa1\xcb\x73\x78\xe0\x39\xdc\x3b\x9f\xe3\xdf\xad\xdf\x8f\xcd\xe4\x50\xdd\x09\x65\x49\x32\xa1\xa4\xcb\x8a\x8c\x22\xdc\xf8\x30\x0d\xe3\x4b\xdd\xb4\xed\xf0\xd0\x4f\xa1\xf6\x5d\x1d\xdc\xbd\x00\x80\x92\xa5\x65\x5c\x67\xbc\xc1\xe2\x58\x64\x46\x59\xca\xc9\x30\x56\xbb\x53\x65\x0a\xe4\x99\xb9\x92\xba\xa2\x8f\x2c\xb7\x9f\x59\x49\xb5\x21\x2c\x12\x51\x92\x26\xc5\x08\x6e\x7a\x6c\x0e\xd1\xec\x1b\xef\xec\x3f\xa2\x13\x99\xcb\x6d\xe4\xbb\x18\x6b\x5b\xe4\x67\x33\xe3\x38\x11\x52\x33\x59\xb0\x5c\x69\x42\x61\xf4\xee\x0c\xc2\x1b\xa1\x0a\x5d\xe5\x06\xbe\x43\x49\x8c\x94\xd6\xb2\xd2\x8c\xde\x3d\xff\xbc\x64\xb9\x1c\xdd\xbe\x3d\x34\x21\xc4\x89\x10\x5f\xcf\x98\x0e\x4f\xfd\x1f\xec\xa9\x2d\x2e\xde\xf5\x89\x38\xa6\xdf\xde\x90\x88\xd7\x00\x00\x00\xff\xff\x55\xe2\xdd\x2c\xbf\x01\x00\x00")

func migrations3_use_sequence_in_history_accountsSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations3_use_sequence_in_history_accountsSql,
		"migrations/3_use_sequence_in_history_accounts.sql",
	)
}

func migrations3_use_sequence_in_history_accountsSql() (*asset, error) {
	bytes, err := migrations3_use_sequence_in_history_accountsSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/3_use_sequence_in_history_accounts.sql", size: 447, mode: os.FileMode(420), modTime: time.Unix(1508347796, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrations4_add_protocol_versionSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\xcd\xb1\x0a\xc2\x30\x10\x06\xe0\x3d\x4f\xf1\xef\x52\x70\xef\x14\x4d\x9d\xce\x44\x4a\x32\x38\x15\xd1\xa3\x06\x6a\xae\x5c\x82\xe2\xdb\xbb\xba\x88\x4f\xf0\x75\x1d\x36\x8f\x3c\xeb\xa5\x31\xd2\x6a\x2c\xc5\x61\x44\xb4\x3b\x1a\x10\x3c\x9d\x71\xcf\xb5\x89\xbe\xa7\x85\x6f\x33\x6b\x85\x01\xac\x73\xd8\x07\x4a\x47\x8f\x55\xa5\xc9\x55\x96\xe9\xc9\x5a\xb3\x14\xe4\xd2\x78\x66\x85\x1b\x0e\x36\x51\xc4\x16\x3e\x44\xf8\x44\xd4\x1b\xf3\x6d\x39\x79\x95\xff\x9a\x1b\xc3\xe9\x97\xd5\x9b\x4f\x00\x00\x00\xff\xff\x83\xbb\x30\x2e\xbc\x00\x00\x00")

func migrations4_add_protocol_versionSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations4_add_protocol_versionSql,
		"migrations/4_add_protocol_version.sql",
	)
}

func migrations4_add_protocol_versionSql() (*asset, error) {
	bytes, err := migrations4_add_protocol_versionSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/4_add_protocol_version.sql", size: 188, mode: os.FileMode(420), modTime: time.Unix(1508347796, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrations5_create_trades_tableSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x94\x51\x6f\xaa\x40\x10\x85\xdf\xf9\x15\x13\x9f\x30\x17\x93\x7b\x6f\x5a\x5f\x4c\x9a\x58\x25\xad\xa9\xc1\xd6\x4a\xd2\x37\xb2\xb0\x23\x6c\xa2\x2c\x99\x1d\xda\xf0\xef\x1b\x68\x69\x10\x57\xad\xaf\x9c\x39\x67\x38\xbb\x5f\x76\x34\x82\x3f\x7b\x95\x92\x60\x84\xb0\x70\x66\x6b\x7f\xba\xf1\x61\x33\xbd\x5f\xfa\x90\x29\xc3\x9a\xaa\x88\x49\x48\x34\xe0\x3a\x00\xf0\xf3\x51\x17\x48\x82\x95\xce\x23\x25\x21\x56\xa9\xca\x19\x82\xd5\x06\x82\x70\xb9\xf4\x9a\xc9\x81\x26\x89\x34\x00\x95\x33\xa6\x48\x1d\xb5\x91\xf5\x76\x8b\x64\x35\x37\xb2\xc1\xdd\xee\x84\x5e\xcb\x71\x59\x9d\x75\xeb\x9d\x8c\x84\x31\xc8\x11\x57\x05\x42\x92\x09\x12\x09\x23\xc1\xbb\xa0\x4a\xe5\xa9\x3b\xbe\x19\xf6\x22\x3b\x1e\x65\x4c\x89\x64\x71\xdd\x8e\xcf\xb8\x12\x2d\x6d\x9b\xfe\xfd\xb7\x7b\xf6\xba\xcc\xb9\xff\xff\x30\x7b\xf4\x67\x4f\xe0\x76\x47\xee\xe0\xef\xf0\xbb\x57\xac\xcb\x34\xe3\x6b\x9b\x1d\xb8\xae\xe8\x76\xe0\xfb\x75\xbb\xd6\x75\xb6\xdf\xe1\x50\xdd\xd0\x19\x4e\x9c\x96\xbf\x30\x58\xbc\x84\x3e\x2c\x82\xb9\xff\x06\x19\x93\x8c\x0a\x25\x61\x15\xf4\x91\x0c\x5f\x17\xc1\x03\xc4\x4c\x88\xe0\xda\xc8\xf4\x5a\x0a\x3b\xe1\x9d\xd4\xb8\x8a\x1a\x0c\x2f\x45\xb7\xac\xda\x52\xea\x90\xfa\xb6\x2e\x65\xf4\x90\xf4\xfa\xe4\x78\xc7\x00\x9e\x5a\xf7\x75\x78\x97\x16\x1e\xb1\xe2\x1d\x5f\xa8\x67\x63\xa3\x5e\xdb\x7d\x17\xe6\xfa\x23\x77\xe6\xeb\xd5\xb3\xfd\x5d\x48\x84\x49\x84\xc4\x89\xf3\x19\x00\x00\xff\xff\x79\x87\x24\x6b\x4c\x04\x00\x00")

func migrations5_create_trades_tableSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations5_create_trades_tableSql,
		"migrations/5_create_trades_table.sql",
	)
}

func migrations5_create_trades_tableSql() (*asset, error) {
	bytes, err := migrations5_create_trades_tableSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/5_create_trades_table.sql", size: 1100, mode: os.FileMode(420), modTime: time.Unix(1508347796, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrations6_create_assets_tableSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x90\x3d\x4f\xc3\x30\x18\x84\x77\xff\x8a\x1b\x1d\x91\x0e\x20\xe8\x92\xc9\x34\x16\x58\x18\xa7\xb8\x31\xa2\x53\xe5\x26\x16\x78\x80\x54\xb6\x11\xca\xbf\x47\xaa\x28\xf9\x50\xe6\x7b\xf4\xbc\xef\xdd\x6a\x85\xab\x4f\xff\x1e\x6c\x72\x30\x27\xb2\xd1\x9c\xd5\x1c\x35\xbb\x97\x1c\x1f\x3e\xa6\x2e\xf4\x07\x1b\xa3\x4b\x11\x94\x00\x80\x6f\xb1\xe3\x5a\x30\x89\xad\x16\xcf\x4c\xef\xf1\xc4\xf7\xc8\xcf\xd9\x19\x3c\xa4\xfe\xe4\xf0\xca\xf4\xe6\x91\x69\xba\xbe\xcd\xa0\xaa\x1a\xca\x48\x39\x86\x9a\xae\x1d\xa0\xeb\x9b\x65\xc8\xc7\xf8\xed\xc2\x3f\x76\xb7\x9e\x63\x46\x89\x17\xc3\xe9\xa0\xcc\x47\x3f\xe4\x13\x4b\x46\xb2\x82\x5c\xfa\x09\x55\xf2\xb7\xbf\xf8\xd8\x5f\xee\x54\x6a\x5e\xd9\xec\x84\x7a\xc0\x31\x05\xe7\x40\x27\xb6\x82\x90\xf1\x74\x65\xf7\xf3\x45\x4a\x5d\x6d\x97\xa7\x6b\x6c\x6c\x6c\xeb\x8a\xdf\x00\x00\x00\xff\xff\xfb\x53\x3e\x81\x6e\x01\x00\x00")

func migrations6_create_assets_tableSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations6_create_assets_tableSql,
		"migrations/6_create_assets_table.sql",
	)
}

func migrations6_create_assets_tableSql() (*asset, error) {
	bytes, err := migrations6_create_assets_tableSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/6_create_assets_table.sql", size: 366, mode: os.FileMode(420), modTime: time.Unix(1508372276, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrations7_modify_trades_tableSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x54\x4d\x8f\xda\x30\x14\xbc\xe7\x57\x3c\xed\x29\x51\xc3\xaa\xad\xda\xbd\x6c\x55\x09\x58\x97\x46\x65\xc3\x36\x04\xa9\xb7\xc8\x89\xdf\x06\xab\xc1\x8e\x6c\xa7\x88\x7f\x5f\x05\x08\xcd\x27\xb0\xbb\x87\x5e\x93\x99\x79\x6f\xec\xf1\x8c\x46\xf0\x6e\xc3\x53\x45\x0d\xc2\x2a\xb7\x46\x23\x60\x4a\xe6\x60\xd6\x08\x32\x63\x60\x14\x65\xa8\xc1\xd0\x38\xc3\x5b\xc8\x0b\x03\x14\x04\x6e\x41\x0a\x04\x2e\x20\xcf\x68\x82\xd6\x43\xb0\x78\x82\x70\x3c\x99\x13\x58\x73\x6d\xa4\xda\x45\x07\xde\xbd\x35\x0d\xc8\x38\x24\xbd\x3f\xc1\xb6\x00\xe0\xf4\x51\xe6\xa8\xa8\xe1\x52\x44\x9c\xc1\xc4\x9b\x79\x7e\x08\xfe\x22\x04\x7f\x35\x9f\xbb\x7b\xe4\x8d\x54\x0c\xd5\x0d\x78\x7e\x48\x66\x24\x68\xfd\xcd\x90\xa5\xa8\xa2\x24\x93\x1a\x59\x44\x0d\x84\xde\x23\x59\x86\xe3\xc7\xa7\x16\x50\x3e\x3f\xa3\x1a\x1c\x12\x53\x8d\x11\x4d\x12\x59\x08\xd3\x03\x82\x80\x7c\x23\x01\xf1\xa7\x64\x79\xda\xfc\x88\xd6\x36\x67\x4e\x5d\x44\x6b\xbc\x5a\xa2\xc4\x76\x04\x36\xa5\x6c\x87\x3e\xfd\x4e\xa6\x3f\xc0\xae\x43\xbe\xc2\xfb\x23\x71\xbf\x09\xaa\x37\x3b\x38\xe9\xbc\xc1\xc4\x49\xe3\xac\x8f\x16\xea\x9f\x95\xbd\x41\xae\x23\x8d\x59\x86\x0a\x26\x8b\xc5\x9c\x8c\xfd\xc3\xbf\x3d\xd7\x6e\x1e\xf3\x97\xce\xd2\x8e\xe5\xdc\x5b\x55\x04\x57\xbe\xf7\x73\x45\xc0\xf3\x1f\xc8\x2f\x58\x1b\xc5\xa2\x9c\x33\x58\xf8\xed\x54\xae\x96\x9e\x3f\x83\xd8\x28\x44\xb0\xfb\xc2\xe9\x56\x41\x74\x4e\xf1\xae\x8b\x52\xae\x22\xc3\x37\x18\x65\x52\xfe\x2e\xf2\xc1\x09\x93\x30\x20\xa4\x69\xc1\xed\x38\x70\x3b\xb1\xee\x1d\x5a\xd1\xae\x1a\xd9\x39\xa5\x3e\xc5\xeb\x1d\x5c\xb5\x60\xbc\x8b\xf6\xcf\xee\xd2\x79\x57\x6f\xb3\xbc\x37\xab\x5e\x4d\x0f\x72\x2b\x1a\xe5\x24\x70\x8b\xaa\xea\x25\x85\x5c\x68\x53\xe2\xaa\xde\x92\x02\x6f\x87\x7b\x09\x12\xaa\x13\xca\xf0\xd5\xfd\x14\xf3\x94\x0b\x33\xd0\x4f\x5c\x18\x4c\x51\x0d\xd5\x4e\x2f\xf7\x10\xf2\xc1\xdf\x71\xb1\x3b\x47\x96\x19\x3b\x5e\xa7\xd9\xe5\x08\xc9\x9a\x2a\x9a\x18\x54\xf0\x87\xaa\x1d\x17\xa9\x7d\xf7\xc9\x19\xe6\x70\xad\x0b\x54\x3d\xac\xcf\x77\x67\x58\x89\x64\x7d\x93\x3e\x7c\xec\xe7\x1c\x5e\x77\x6b\xfd\xaa\x03\xea\x90\x5a\x01\xc8\x22\x5d\x9b\x97\x1a\x6b\xb0\x5e\x60\xad\xc1\xbb\xda\x5c\xc5\x3a\x6b\xaf\x09\x2a\x0d\xfe\x87\x62\x7a\xc5\x13\x6c\x8b\x94\x1a\xe5\x55\x5d\x92\x68\xe5\xd1\x6d\xc7\xc6\xed\xa6\x6f\x60\xda\xe1\xe4\x2e\xcd\xeb\x04\xc5\xed\xde\xa6\xdb\x17\x0c\xe7\xfe\x6f\x00\x00\x00\xff\xff\x2a\xff\xe8\x4a\xff\x08\x00\x00")

func migrations7_modify_trades_tableSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations7_modify_trades_tableSql,
		"migrations/7_modify_trades_table.sql",
	)
}

func migrations7_modify_trades_tableSql() (*asset, error) {
	bytes, err := migrations7_modify_trades_tableSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/7_modify_trades_table.sql", size: 2303, mode: os.FileMode(420), modTime: time.Unix(1509125270, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrations8_create_asset_stats_tableSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x91\xc1\x6e\xf2\x30\x10\x84\xef\x79\x8a\x3d\x82\xfe\x9f\x5b\xd5\x0b\x27\x63\x6f\xa9\x55\xe3\x44\x1b\xa7\x2a\xa7\xc8\x24\x2e\x8d\x44\x92\x2a\x36\xaa\x78\xfb\xca\x84\x56\xa1\x8a\x4f\xb6\xf5\x69\x66\x67\x67\xb5\x82\x7f\x6d\x73\x1c\x6c\x70\x50\x7c\x26\x9c\x90\x19\x04\xc3\x36\x0a\xc1\x7a\xef\x42\xe9\x83\x0d\x1e\x16\x09\x00\x40\x53\xc3\xdd\xd9\xc8\xad\xd4\x66\xbc\x67\x24\x77\x8c\xf6\xf0\x82\x7b\x20\x7c\x42\x42\xcd\x31\x87\x8f\xc6\x87\x7e\xb8\x94\x57\x31\x0f\xa9\x06\x81\x0a\x0d\x02\x67\x39\x67\x02\xe3\x4f\x91\x89\xe8\x4a\x98\x1b\x92\xdc\xfc\xbf\x7a\xd9\xb6\x3f\x77\x61\xde\x4b\xa7\x06\x74\xa1\xd4\x48\x76\xe7\xb6\xb4\x55\x15\x71\x1f\xdf\x52\x1b\xdc\x22\xcd\x90\xef\x27\x7b\xf4\x93\xf9\xf3\x1d\x53\xea\xa6\x7a\x4f\x86\xbe\x3d\x4d\x93\xbe\x32\xe2\xcf\x8c\x16\x8f\x0f\xcb\x5f\x32\x59\xae\x93\x9f\x8d\x49\x2d\xf0\xed\xb6\xb1\xc3\xa5\xac\xfa\xda\xc5\x64\x7f\xd2\x17\xb9\xd4\x5b\x38\x84\xc1\x39\x58\x8c\x70\x24\xa3\xce\xb4\x08\xd1\x7f\x75\x89\xa0\x34\x9b\x29\xa2\xb2\xbe\xb2\xb5\x5b\x7f\x07\x00\x00\xff\xff\xa9\x7e\x10\x6f\xb9\x01\x00\x00")

func migrations8_create_asset_stats_tableSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations8_create_asset_stats_tableSql,
		"migrations/8_create_asset_stats_table.sql",
	)
}

func migrations8_create_asset_stats_tableSql() (*asset, error) {
	bytes, err := migrations8_create_asset_stats_tableSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/8_create_asset_stats_table.sql", size: 441, mode: os.FileMode(420), modTime: time.Unix(1510554929, 0)}
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
	".DS_Store": Ds_store,
	"latest.sql": latestSql,
	"migrations/1_initial_schema.sql": migrations1_initial_schemaSql,
	"migrations/2_index_participants_by_toid.sql": migrations2_index_participants_by_toidSql,
	"migrations/3_use_sequence_in_history_accounts.sql": migrations3_use_sequence_in_history_accountsSql,
	"migrations/4_add_protocol_version.sql": migrations4_add_protocol_versionSql,
	"migrations/5_create_trades_table.sql": migrations5_create_trades_tableSql,
	"migrations/6_create_assets_table.sql": migrations6_create_assets_tableSql,
	"migrations/7_modify_trades_table.sql": migrations7_modify_trades_tableSql,
	"migrations/8_create_asset_stats_table.sql": migrations8_create_asset_stats_tableSql,
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
	".DS_Store": &bintree{Ds_store, map[string]*bintree{}},
	"latest.sql": &bintree{latestSql, map[string]*bintree{}},
	"migrations": &bintree{nil, map[string]*bintree{
		"1_initial_schema.sql": &bintree{migrations1_initial_schemaSql, map[string]*bintree{}},
		"2_index_participants_by_toid.sql": &bintree{migrations2_index_participants_by_toidSql, map[string]*bintree{}},
		"3_use_sequence_in_history_accounts.sql": &bintree{migrations3_use_sequence_in_history_accountsSql, map[string]*bintree{}},
		"4_add_protocol_version.sql": &bintree{migrations4_add_protocol_versionSql, map[string]*bintree{}},
		"5_create_trades_table.sql": &bintree{migrations5_create_trades_tableSql, map[string]*bintree{}},
		"6_create_assets_table.sql": &bintree{migrations6_create_assets_tableSql, map[string]*bintree{}},
		"7_modify_trades_table.sql": &bintree{migrations7_modify_trades_tableSql, map[string]*bintree{}},
		"8_create_asset_stats_table.sql": &bintree{migrations8_create_asset_stats_tableSql, map[string]*bintree{}},
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

