// Code generated by go-bindata.
// sources:
// latest.sql
// migrations/1_initial_schema.sql
// migrations/2_index_participants_by_toid.sql
// migrations/3_use_sequence_in_history_accounts.sql
// migrations/4_add_protocol_version.sql
// migrations/5_create_trades_table.sql
// migrations/6_create_assets_table.sql
// migrations/7_modify_trades_table.sql
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

var _latestSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x5c\x5d\x6f\xdb\xb8\xd2\xbe\xcf\xaf\x20\xf6\xc6\x31\x60\x07\x71\xbe\xda\x38\x6f\x0b\xb8\x8e\xba\x35\xea\xca\x5b\xdb\x79\xbb\xc5\x62\x41\xd0\x12\x6d\xf3\x54\x12\x55\x92\x4e\x93\x3d\x38\xff\xfd\x40\x5f\xd6\x17\x29\x4a\xb6\xd2\x9e\xcb\x98\xa3\x99\xe7\x19\x0e\x39\xe4\x0c\x91\x7e\xff\xa4\xdf\x07\x7f\x50\x2e\x36\x0c\x2f\x3e\x4f\x81\x8d\x04\x5a\x21\x8e\x81\xbd\x73\xfd\x93\x7e\xff\x24\x18\xbf\xdf\xb9\x3e\xb6\xc1\x9a\x51\x37\x15\x78\xc4\x8c\x13\xea\x81\xdb\xb3\x9b\xb3\xeb\x8c\xd4\xea\x19\xf8\x1b\x18\x7c\x5e\x10\x39\x59\x18\x4b\xc0\x05\x12\xd8\xc5\x9e\x80\x82\xb8\x98\xee\x04\x78\x03\xce\xef\xc2\x21\x87\x5a\xdf\xca\xbf\x5a\x0e\x09\xa4\xb1\x67\x51\x9b\x78\x1b\xf0\x06\x74\x1e\x96\xef\x5f\x77\xee\x12\x75\x9e\x8d\x98\x0d\x2d\xea\xad\x29\x73\x89\xb7\x81\x5c\x30\xe2\x6d\x38\x78\x03\xa8\x17\xeb\xd8\x62\xeb\x1b\x5c\xef\x3c\x4b\x10\xea\xc1\x15\xb5\x09\x0e\xc6\xd7\xc8\xe1\x38\x67\xc6\x25\x1e\x74\x31\xe7\x68\x13\x0a\xfc\x40\xcc\x23\xde\xe6\x2e\xc6\x8e\x11\xb3\xb6\xd0\x47\x62\x0b\xde\x00\x7f\xb7\x72\x88\xd5\x0b\xc8\x5a\x48\x20\x87\x26\x62\x36\x5e\xa3\x9d\x23\xa0\x40\x2b\x07\x73\x1f\x59\x38\x00\xdd\x29\x8c\xfe\x20\x62\x0b\x29\xb1\x33\x38\x4e\xa2\xd9\x30\x91\x8b\x87\x60\x43\x99\x0f\x5d\xb2\x61\x28\xc0\xcc\xef\xc0\xf2\xd9\xc7\x43\xb0\x1c\xbd\x9b\x1a\x77\x60\x61\x6d\xb1\x8b\x86\x31\x88\x3b\x30\xfb\xe1\x61\x36\x04\xfd\x70\xc6\xc6\x73\x63\xb4\x34\x22\xd1\xa2\x1e\x70\x7a\x02\x00\x00\xc4\x06\x02\x3f\x09\x60\xce\x96\xc0\x7c\x98\x4e\x7b\xe1\xaf\xc8\xf7\x1d\x82\x6d\x88\x04\x08\xe6\x81\x0b\xe4\xfa\x20\x00\x1a\xfe\x09\xfe\xa1\x1e\x3e\xe9\xde\x9d\xe4\x81\x6e\x09\x17\x94\x3d\x43\x64\x59\x74\xe7\x09\x0e\x89\x0d\x39\xfe\x9e\x00\x5e\x18\x9f\x1f\x0c\x73\x5c\x13\x73\x22\xad\xd2\x1a\xc2\x5c\x2c\x47\xf3\x25\xf8\x32\x59\x7e\x00\x83\xf0\x87\x89\x39\x9e\x1b\x9f\x0c\x73\x09\xde\x7d\x8d\x7f\x32\x67\xe0\xd3\xc4\xfc\xff\xd1\xf4\xc1\xd8\xff\x3d\xfa\x33\xfd\x7b\x3c\x1a\x7f\x30\xc0\x40\x47\xe6\x60\xb7\x17\x15\xa5\x7e\x5f\x91\x0d\xf1\x04\xb8\x37\xde\x8f\x1e\xa6\x4b\xe0\xe1\x27\xf1\x88\x9c\xd3\x8e\x82\x71\x67\x38\x64\x78\x63\x39\x88\xf3\x6e\x71\xba\x6c\x9b\x61\xce\x81\xb5\x45\x0c\x59\x02\x33\xf0\x88\xd8\x33\xf1\x36\xa7\x37\x57\xdd\x8a\x89\xe2\x1c\xb7\xc1\x2c\x54\x93\xf2\x22\x9e\xc0\x1b\xcc\x8a\x18\x03\x29\x28\x9e\x7d\x2c\x87\x29\x15\xb7\xa8\x2d\x13\x1f\x5c\xc8\xc5\x09\xe7\x3b\xcc\x24\x1f\x5c\xdf\xa4\x1f\xe8\xfc\xd1\x72\xd8\x66\x75\xfe\xb4\xa0\xad\x22\x02\x66\x5f\x4c\xe3\x1e\xbc\xfb\xaa\x61\x34\x9a\x2e\x8d\xb9\x86\xd0\x5e\x57\x61\xf8\x8c\xd8\x2a\x6c\x78\xbd\xc6\x56\x0b\x51\x17\xeb\x89\xc3\xae\xb0\x66\x60\xba\xbc\xf2\x71\x92\xc8\x51\x1f\x47\xfb\xa0\x52\xf2\x37\xca\x6c\xcc\x7e\x53\x44\x73\x18\xc7\xf2\x21\x1b\x0b\x44\x1c\x0e\xfe\xc5\xa9\xb7\x52\x07\x9b\x83\xed\x0d\x66\xc7\xfb\x21\xd6\x13\xfb\x81\xe3\xef\x3b\xec\x59\x2a\x6c\x91\x30\xdc\x22\xbe\xad\xb5\x0a\x7d\x86\x1f\x09\xdd\x71\xa8\xfd\x30\x76\x0b\x43\x1e\x47\x51\x72\x0d\x27\x62\x8f\x23\xd9\xe5\xce\x0b\x16\xd2\x89\xa8\x27\x6f\x39\x94\xcb\x12\x53\x70\x54\xd8\xe7\xa6\xe2\x37\x0c\x23\xa1\xfd\x28\x92\xdd\xf9\x76\x6d\xd9\x7d\xe8\xc4\x7f\xba\x3e\x65\x02\x33\x98\x9c\x76\x8a\x5c\x06\xc5\x20\xa2\x02\x39\xd0\xa2\xc4\xe3\xf2\x18\x5c\x63\x0c\x7d\x4a\x1d\xf9\x68\x70\xf8\x82\x6b\xac\x9a\xeb\x70\x98\x61\x8e\xd9\xa3\x4a\xc4\x45\x4f\x50\x3c\xc1\x60\xeb\xe4\xe4\x1f\x95\x94\xcf\xa8\xa0\x16\x75\x94\xbc\xce\x6b\xec\xad\xe9\x3c\xfb\x88\x09\x62\x11\x1f\xb5\x91\x55\xe5\x6a\x75\xb9\xa8\xfe\x2e\xa0\xdf\x57\x9a\x52\x6e\x37\xbd\x54\xda\xf8\x59\xe9\xa6\x11\xd1\x23\xd3\x4f\xa5\xad\x72\x3a\x92\x8b\x57\xa4\xa7\xfd\x07\x2d\xc6\x66\xf9\xcc\x57\xd8\x07\x32\xbb\xa6\x4a\x26\x3c\x91\x5b\x11\x95\x30\x33\x1d\x99\x98\xa2\x9f\x38\xdd\x31\x0b\x27\xd1\xad\x48\x09\xc9\x32\xef\x74\x86\xc3\x92\x44\x8d\x75\x20\x18\xb2\xf1\xf1\xee\x8c\xd4\x14\xf2\xfd\xb1\x79\x3c\xce\x6b\x87\x64\x15\xba\x5e\x63\xa6\x34\x1b\xee\xbe\xba\xd3\x48\x24\x14\x1d\x5d\x2b\x45\xdc\x70\x7a\xa4\x02\xa1\x05\xcc\xb4\xb6\xf6\x72\x95\xe6\xf6\x52\x15\x16\x43\x48\x84\x43\x8e\x1d\x07\x33\xb0\xa2\xd4\xc1\xc8\x8b\xc6\xc6\x33\x73\xb1\x9c\x8f\x26\xe6\xb2\x30\x6f\x30\x43\x04\x86\xd7\x70\x30\xfe\x60\x8c\x3f\x82\xd3\xd3\x2c\xc5\xb7\xe0\xbc\xdb\xd5\xa9\x92\x7d\x9e\xb0\xfa\xbf\x12\xd1\x1a\xfa\x72\xa4\x0b\xea\x0b\x1e\x09\x01\x56\xc6\xfa\x7e\x29\xb7\x9a\xe8\x54\x8a\xeb\xa6\xba\x3a\x7b\xcc\x31\xc9\x4e\x85\xaf\xdd\x74\xa7\xb1\xf2\xb3\x12\x5e\x43\xb2\x47\xa6\x3c\x8d\xb5\x72\xd2\x53\x7d\x50\x91\xf6\x32\x9f\xb4\x1a\xab\x49\x7c\x66\x21\xd5\xbe\x7d\xc4\x9b\xb3\xe6\x4e\x53\x37\x33\x56\x27\x39\xa9\x6c\x6a\x5a\x7d\x3c\x47\xca\xa5\xa7\xba\xda\xfc\x92\xcb\x89\x78\x82\xd8\x7b\xc4\x0e\xf5\xb1\xac\xe0\x27\x9e\x82\xab\xc2\xce\x11\x8a\x41\x17\x0b\xa4\x18\x0a\xbc\xa0\x1a\xe6\x64\xe3\x21\xb1\x63\x58\x56\x9b\xba\xbd\xe9\xfe\xf5\x77\x7a\xba\xf8\xf7\x7f\x64\xe7\x8b\xbf\xfe\x2e\xde\x59\xb0\x4b\x15\x65\xa4\x54\x97\x47\x3d\x5c\x79\x5a\x49\x75\x95\xd5\xc4\xcc\x88\x8b\xe1\x8a\xee\x3c\x9b\x07\x33\xf7\x9a\x21\x6f\x53\x55\xf4\x8c\x8a\x60\xc4\x4e\x56\x4f\x8c\xa5\xd6\x92\x8f\x96\xcf\xcc\x9c\x16\xeb\x28\x20\x1a\x1f\xcf\xa6\x0f\x9f\xcc\x60\x4a\x17\xc6\xb2\xa2\x60\x98\x2d\xcd\x64\xcb\x85\xcd\x0e\xee\xed\x91\x50\xe8\x6f\x44\xaa\xf2\xc0\x5f\x87\xa4\x32\x73\xb6\x46\x53\x69\xa1\x11\x51\xcd\x36\x2f\xa7\x7a\x8f\x04\x02\x6b\xca\x6a\xf4\x0a\xc0\xfd\x68\x39\xd2\x50\x9c\x98\x0b\x63\xbe\x04\x13\x73\x39\x2b\xf5\x0b\xc2\xec\xb8\x00\xa7\x9d\x01\x24\x1e\x11\x04\x39\x90\x87\xba\xce\xf8\x77\xa7\xd3\x03\x9d\x8b\xf3\xc1\xab\xfe\xe0\xbc\x7f\x71\x0d\x06\x17\xc3\xf3\x8b\xe1\xd5\xe0\xec\xf2\xfa\xfa\xf5\xe0\xba\x7f\xfe\xaa\xd3\xbd\xab\xa7\xfd\x02\x12\xcf\xc6\x4f\x79\x17\xac\x9e\xa1\xa0\xc4\xae\xb6\x74\x3b\xb8\x6c\x62\xe8\x12\xee\x38\xde\xef\xf0\x90\x78\xb0\x58\x78\xaf\x34\x77\x33\x18\x0c\x6e\x9b\xd8\xbb\x82\xc8\xb6\x61\xb1\x98\x52\x6d\xe3\xfa\xf6\xf6\x75\x13\x1b\xd7\x30\x4a\x27\xc9\xc9\x36\xec\x3c\x55\x9a\x78\x75\x7e\x75\xd5\xc8\x6d\x37\x89\x89\x78\xb7\xa9\x61\xe2\xf2\xd5\xd5\x4d\x13\x13\xaf\xa0\x4b\x6d\xb2\x7e\xae\xcf\xe2\xf5\xe0\xf6\xfc\x22\x36\xa1\x58\x1a\x95\x0d\x9d\x3a\x6b\xe3\xa0\x66\x57\xb0\xe4\x35\x7a\x17\xc6\xd4\x18\x2f\x33\xdd\xc3\x33\x8e\xab\x1b\x41\x3d\x30\xe8\x45\xad\xc2\x1a\x74\xcb\x3d\x9e\x23\xc8\x56\xf6\x15\x5a\xa1\x9a\x4b\x61\x4d\x88\xca\xfa\x0a\x4d\x98\x2a\xd4\xca\xca\xf4\x2d\xa8\xad\x51\x0e\x3d\x7c\x9a\x9a\xd5\xe3\xda\x98\xb6\xea\x24\xdd\x64\x1a\x15\xf5\xb7\x16\x5c\x2e\x29\x43\xb5\xa3\x55\x7f\xe1\x3f\x7c\x2a\x9b\xde\x34\xdb\x98\x4c\xdd\x41\xa4\xc9\x74\x2a\xef\x95\xcd\x5d\x52\x4c\x17\x85\xbf\xa1\xff\x0d\x3f\x27\x26\xd2\x5a\x4f\xd3\x33\x5d\x41\x6b\x78\x15\x18\xdd\xdf\x67\xab\x47\x32\xc3\xe0\x8f\xf9\xe4\xd3\x68\xfe\x15\x7c\x34\xbe\x82\x53\x62\xeb\xee\x09\x85\xfd\x2e\xed\x76\xc3\xb4\x4f\x0e\xb3\x4d\x6d\xd8\x0a\xbb\xbc\x59\x19\xb9\x83\x80\x81\x07\x73\xf2\xf9\xc1\x00\xa7\xa9\x78\x2f\xd3\xf0\xef\xe5\xda\xf3\x0d\x5d\xd3\xce\xb4\x36\x26\xde\x68\x52\x15\xf7\x1c\xcd\xee\xd8\x2e\x33\xb9\x91\x2a\xa6\x15\xb0\x6a\x33\x57\x5e\x7d\xb4\x9b\x49\xbb\xec\x55\x66\xaa\xf8\x57\x42\xd3\x7a\x20\x0a\xe9\xd5\x73\x1c\xd5\x09\x95\x89\x79\x6f\xfc\x59\xaf\x68\x16\x8a\x16\xf5\x80\x99\x59\x5c\x10\x0f\x8b\x89\xf9\x3b\x58\x09\x86\x71\xb2\xc2\x14\x2b\x69\xb5\x3f\x2b\x1e\x0c\x27\x55\x91\x45\x92\xab\xe8\xe5\xf1\x44\xc2\xbd\x52\xc9\x4c\x06\x6e\x8b\xf8\xf6\x18\x64\x61\xe5\xb0\x16\xac\x62\xbd\x51\x86\x26\x3a\xda\x1d\x83\x27\xd2\x50\x0f\x51\xa1\x98\xd9\x2b\xd7\x2d\xa5\x8b\x0c\xe2\x20\x36\xc2\xf1\x03\x90\xc6\xfb\x72\x04\xb8\xa0\x2e\x0b\x3b\x79\x4c\x93\x43\x2c\xeb\xb1\xf5\x92\x7e\x9a\x0a\x6c\x5a\x54\x39\x12\x26\xb1\x6b\x03\x4c\xfb\x15\x3d\x69\x63\x50\x03\x9a\xfa\xd0\x6f\x0b\x77\xac\x2b\x0b\x5d\x91\x1c\x0e\x62\x22\x27\x20\x9e\xda\x23\x10\xeb\x52\xc4\xf4\x81\x14\xf2\xcd\xa7\x32\x09\xea\x07\x51\xb9\xa5\x07\x71\x88\xc1\xa7\x3a\x0e\x75\x7e\xb5\xa3\xf7\x6f\xa0\x82\xad\xfa\x78\x5f\xe7\xd5\x65\x21\x27\x0f\xba\x72\x18\xe5\x88\xb2\x7e\x6d\x0b\x56\x49\x67\xbd\xed\x4d\x06\x50\x44\x53\x22\x8e\x99\xd6\x54\xc7\xe1\x21\xa9\x0b\x3f\xc1\xec\x70\x57\x5c\xaf\x8f\x48\x07\x39\x2d\x05\xac\x36\x2e\x20\x4b\xde\x0e\xc8\xb1\x24\x2d\x67\x87\xd2\x6f\x3b\xff\x38\x44\x79\x5d\x3a\x5c\xa5\xde\xb9\x14\x9f\x8f\x08\x0b\xdf\xe7\xb7\x82\xb0\xa8\x4d\x87\x31\xd7\xef\xef\x95\xda\xfd\xbd\xd2\x9b\x0e\x05\x89\x16\x56\x4b\xac\x47\x87\xb8\x61\x4e\x0a\xb4\xb6\xe6\xdd\x06\x8e\xd5\xfa\x2d\x2a\xcc\x97\xaa\x92\xd4\x83\xf1\x03\xf4\x63\x1d\xaa\x35\x90\x3b\x1d\x27\x0f\xea\xf3\xe7\xd1\x48\xb0\x01\xf6\xe3\xe3\xa0\x4a\xb7\x1e\xb1\x64\x95\xe5\x15\xc6\x67\x9f\x40\x5f\x70\x9b\x3e\x38\x1e\x2a\xb5\x6a\x0f\x5b\x81\x90\x06\x68\x9c\xb9\x02\x95\xfb\x20\x6a\x09\xad\x4c\xb5\x36\x69\xd6\x8d\xe4\x8c\xf2\xb6\x83\x21\xa7\xfa\x90\x2c\xaf\x56\x57\x78\x6d\xdc\xbe\xa3\x4b\xef\x99\xb5\xf0\x0b\x1f\xd4\x27\x93\x79\x5e\xfe\x62\xfe\xcf\x3e\x61\xd7\x31\xc9\xc8\xd6\x27\x21\x7b\x2c\xff\x62\x6c\xa4\x2f\xf3\x75\xb4\x64\x1f\xd5\xe7\x97\x5c\x5d\x5f\x8c\xd3\xfe\xb5\x8d\x8e\x87\xb2\xc6\x90\x57\x9d\xf6\x12\x5e\x62\x69\x17\xb5\x4b\xaf\x1d\x4d\x17\x78\x5e\x69\xfe\xe0\xda\xd2\x0a\xaf\x32\x51\x87\x83\xe6\x34\x5d\x69\xac\xbd\xf4\x55\x56\x5c\x0b\xbb\x3e\x89\x65\xaf\x38\x2f\x11\x36\x65\xfd\x07\x5f\xb0\xc2\x43\xdc\x3e\x91\x27\x75\x1d\xb8\xa2\xf4\xdb\xc1\x5e\xae\xd0\xa9\x3d\x22\x9c\x9e\x26\x4f\xcc\xfb\x6f\xdf\x82\x0e\xa7\x8e\x9d\xe9\x1a\x74\x86\x43\x81\x9f\x44\xb7\xdb\x03\x6a\x41\x8b\xda\xf5\x04\xa3\x0a\xa8\x5a\x74\x45\x77\x9b\xad\xa8\x65\x3e\x27\x5a\x0d\x20\x27\x5a\x80\xd0\x05\x5f\x3e\x18\x73\x23\x0a\x32\xf0\x06\x5c\x5e\x6a\x5e\xc4\xcb\x5f\x48\xef\xcb\x27\x70\x9d\xa9\x90\xbf\xff\xd8\x4e\x91\xdc\xc6\xba\x92\xb8\x0a\x09\x78\x3f\x9b\x1b\x93\xdf\xcd\xa8\x20\x5e\x90\xe8\x82\xb9\xf1\xde\x98\x1b\xe6\xd8\x58\x94\x8e\xbb\x9a\xfe\x81\xd2\x13\xf1\x85\xee\x97\xfb\x21\x8b\x43\xe6\x85\xe4\xae\x2c\xf5\x41\x58\xc2\x6f\xee\x81\xf2\xab\xfe\x5f\xe8\x06\x05\x98\xbc\x2f\xca\x42\x2d\x07\x45\xf1\xa2\xff\xbf\xe0\x10\x75\x68\x94\x2a\x29\x75\xa3\x43\xf5\x2f\x18\x80\x45\x5d\xdf\xc1\x02\x87\x1c\xfe\x1b\x00\x00\xff\xff\x83\x5e\xe2\xb7\xaf\x41\x00\x00")

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

	info := bindataFileInfo{name: "latest.sql", size: 16815, mode: os.FileMode(420), modTime: time.Unix(1508958161, 0)}
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

	info := bindataFileInfo{name: "migrations/1_initial_schema.sql", size: 10559, mode: os.FileMode(420), modTime: time.Unix(1508537629, 0)}
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

	info := bindataFileInfo{name: "migrations/2_index_participants_by_toid.sql", size: 277, mode: os.FileMode(420), modTime: time.Unix(1508537629, 0)}
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

	info := bindataFileInfo{name: "migrations/3_use_sequence_in_history_accounts.sql", size: 447, mode: os.FileMode(420), modTime: time.Unix(1508537629, 0)}
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

	info := bindataFileInfo{name: "migrations/4_add_protocol_version.sql", size: 188, mode: os.FileMode(420), modTime: time.Unix(1508537629, 0)}
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

	info := bindataFileInfo{name: "migrations/5_create_trades_table.sql", size: 1100, mode: os.FileMode(420), modTime: time.Unix(1508537629, 0)}
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

	info := bindataFileInfo{name: "migrations/6_create_assets_table.sql", size: 366, mode: os.FileMode(420), modTime: time.Unix(1508537629, 0)}
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

	info := bindataFileInfo{name: "migrations/7_modify_trades_table.sql", size: 2303, mode: os.FileMode(420), modTime: time.Unix(1508958134, 0)}
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
	"latest.sql": latestSql,
	"migrations/1_initial_schema.sql": migrations1_initial_schemaSql,
	"migrations/2_index_participants_by_toid.sql": migrations2_index_participants_by_toidSql,
	"migrations/3_use_sequence_in_history_accounts.sql": migrations3_use_sequence_in_history_accountsSql,
	"migrations/4_add_protocol_version.sql": migrations4_add_protocol_versionSql,
	"migrations/5_create_trades_table.sql": migrations5_create_trades_tableSql,
	"migrations/6_create_assets_table.sql": migrations6_create_assets_tableSql,
	"migrations/7_modify_trades_table.sql": migrations7_modify_trades_tableSql,
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
	"latest.sql": &bintree{latestSql, map[string]*bintree{}},
	"migrations": &bintree{nil, map[string]*bintree{
		"1_initial_schema.sql": &bintree{migrations1_initial_schemaSql, map[string]*bintree{}},
		"2_index_participants_by_toid.sql": &bintree{migrations2_index_participants_by_toidSql, map[string]*bintree{}},
		"3_use_sequence_in_history_accounts.sql": &bintree{migrations3_use_sequence_in_history_accountsSql, map[string]*bintree{}},
		"4_add_protocol_version.sql": &bintree{migrations4_add_protocol_versionSql, map[string]*bintree{}},
		"5_create_trades_table.sql": &bintree{migrations5_create_trades_tableSql, map[string]*bintree{}},
		"6_create_assets_table.sql": &bintree{migrations6_create_assets_tableSql, map[string]*bintree{}},
		"7_modify_trades_table.sql": &bintree{migrations7_modify_trades_tableSql, map[string]*bintree{}},
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

