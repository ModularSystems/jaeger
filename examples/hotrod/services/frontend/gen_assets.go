// Code generated by "esc -pkg frontend -o examples/hotrod/services/frontend/gen_assets.go -prefix examples/hotrod/services/frontend/web_assets examples/hotrod/services/frontend/web_assets"; DO NOT EDIT.

package frontend

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"sync"
	"time"
)

type _escLocalFS struct{}

var _escLocal _escLocalFS

type _escStaticFS struct{}

var _escStatic _escStaticFS

type _escDirectory struct {
	fs   http.FileSystem
	name string
}

type _escFile struct {
	compressed string
	size       int64
	modtime    int64
	local      string
	isDir      bool

	once sync.Once
	data []byte
	name string
}

func (_escLocalFS) Open(name string) (http.File, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	return os.Open(f.local)
}

func (_escStaticFS) prepare(name string) (*_escFile, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	var err error
	f.once.Do(func() {
		f.name = path.Base(name)
		if f.size == 0 {
			return
		}
		var gr *gzip.Reader
		b64 := base64.NewDecoder(base64.StdEncoding, bytes.NewBufferString(f.compressed))
		gr, err = gzip.NewReader(b64)
		if err != nil {
			return
		}
		f.data, err = ioutil.ReadAll(gr)
	})
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (fs _escStaticFS) Open(name string) (http.File, error) {
	f, err := fs.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.File()
}

func (dir _escDirectory) Open(name string) (http.File, error) {
	return dir.fs.Open(dir.name + name)
}

func (f *_escFile) File() (http.File, error) {
	type httpFile struct {
		*bytes.Reader
		*_escFile
	}
	return &httpFile{
		Reader:   bytes.NewReader(f.data),
		_escFile: f,
	}, nil
}

func (f *_escFile) Close() error {
	return nil
}

func (f *_escFile) Readdir(count int) ([]os.FileInfo, error) {
	if !f.isDir {
		return nil, fmt.Errorf(" escFile.Readdir: '%s' is not directory", f.name)
	}

	fis, ok := _escDirs[f.local]
	if !ok {
		return nil, fmt.Errorf(" escFile.Readdir: '%s' is directory, but we have no info about content of this dir, local=%s", f.name, f.local)
	}
	limit := count
	if count <= 0 || limit > len(fis) {
		limit = len(fis)
	}

	if len(fis) == 0 && count > 0 {
		return nil, io.EOF
	}

	return fis[0:limit], nil
}

func (f *_escFile) Stat() (os.FileInfo, error) {
	return f, nil
}

func (f *_escFile) Name() string {
	return f.name
}

func (f *_escFile) Size() int64 {
	return f.size
}

func (f *_escFile) Mode() os.FileMode {
	return 0
}

func (f *_escFile) ModTime() time.Time {
	return time.Unix(f.modtime, 0)
}

func (f *_escFile) IsDir() bool {
	return f.isDir
}

func (f *_escFile) Sys() interface{} {
	return f
}

// FS returns a http.Filesystem for the embedded assets. If useLocal is true,
// the filesystem's contents are instead used.
func FS(useLocal bool) http.FileSystem {
	if useLocal {
		return _escLocal
	}
	return _escStatic
}

// Dir returns a http.Filesystem for the embedded assets on a given prefix dir.
// If useLocal is true, the filesystem's contents are instead used.
func Dir(useLocal bool, name string) http.FileSystem {
	if useLocal {
		return _escDirectory{fs: _escLocal, name: name}
	}
	return _escDirectory{fs: _escStatic, name: name}
}

// FSByte returns the named file from the embedded assets. If useLocal is
// true, the filesystem's contents are instead used.
func FSByte(useLocal bool, name string) ([]byte, error) {
	if useLocal {
		f, err := _escLocal.Open(name)
		if err != nil {
			return nil, err
		}
		b, err := ioutil.ReadAll(f)
		_ = f.Close()
		return b, err
	}
	f, err := _escStatic.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.data, nil
}

// FSMustByte is the same as FSByte, but panics if name is not present.
func FSMustByte(useLocal bool, name string) []byte {
	b, err := FSByte(useLocal, name)
	if err != nil {
		panic(err)
	}
	return b
}

// FSString is the string version of FSByte.
func FSString(useLocal bool, name string) (string, error) {
	b, err := FSByte(useLocal, name)
	return string(b), err
}

// FSMustString is the string version of FSMustByte.
func FSMustString(useLocal bool, name string) string {
	return string(FSMustByte(useLocal, name))
}

var _escData = map[string]*_escFile{

	"/index.html": {
		name:    "index.html",
		local:   "examples/hotrod/services/frontend/web_assets/index.html",
		size:    3384,
		modtime: 1550511159,
		compressed: `
H4sIAAAAAAAC/9RX/1PbOBb/PX/FG23v7FywnRAoNMS54UiX0r0uvUC709vpD7L0YivYkivJgSzD/34j
2wkJlJv9cZsZQHrfP++LXhhntsgnHYBxgZYCy6g2aGNycXUZHB8fvgkG5JEraYExWQq8LZW2BJiSFqWN
ya3gNos5LgXDoL7sgZDCCpoHhtEc40HY34OC3omiKrZJlUFd32mSY9xvnGVIuTsAjK2wOU7eKTu7nEIA
M8HRwKWEKRZU8nHU8BtZw7QoLRjNYpJZW5pRFDHFMVx8q1CvQqaKqDkGw3AQDsJCyHBhyGQcNaqtnVzI
G9CYx8TYVY4mQ7QEMo3zR7sFvWNcholS1lhNS3dx9jeEaBgOw6OIGfNIqx0yYwgIaTHVwq5iYjI6PD4I
/vX5ixBXFz/jLwN+Xryfnd6sWPXu9N0sHe5fFp/Y7e2RksPZF54efKa9j8XVtfkj+uX18TLhbxfZQUWA
aWWM0iIVMiZUKrkqVGXI/0nOnwWxeIph8V0I1+zw4j8i6e8ffVuuFlcf5u8Wlx/ov2/m1W+f7/579+mj
PHt/epTvF2e//XpRnr8pzs+mx7fnv16wj9Oj6zv6MoTHArVgXF0mnbCqBId7KKhOhQysKkcwOCzvTuCh
E2bKasWDpLJWSbiHknIuZDqC/b6TYJU2So+gVA6IPtk10v+ekVGmlqjh/rnuXOQW9QgSLdLMSjTGPz78
W9eZ+Kk1kav0hUh/sqJ8gVWDjVq0bjKi9WiME8VXbWm5WALLqTExcRNJhUTdln2XW6eL5qht8zsQcq5c
drlYbuQZOkzrq5vGgZs/mIWX4TQcR9lgm3cwGWMxeTaWWEzGUXawJbkVhla35JHzHEIeFDwYgjuYInj9
RLZpgJLKZ1T3aY0kVkJiZQ2wPiS5YjewU07yXQOcWhqwylhVoI7JYH9IJjPKMsw9Az/nStMcpmhEKs04
cmE8QbKdy786uOGbfTK51qqAs0wxlVMrUP/wqI6GAzJ5T0sq0aCrFWr749fq8PURmZwW9A8hUzhT8zki
zBQ1FvWfAff06nAKHhMrSjI5ywW7ASVh7a7e9EATtUSwCpTmqIECozp8ydDjM0fW2HOk/OnrEm0/LxvW
OGqes85mUU06nXklmRVKwlzpgtpppam7+rw9dOG+A7CkGjjEsKZCBP6gX3/gHzBo/rzud09a2UoKayAG
rxDSc0SNttISPlCbhVpVkvu8C71G7qTz0Ok4LZYLlPbTp4spxNuizZFKrgq/2/pzvpxOTo2d4bcKja3V
+iedziuf1EuLdEP3xcsnX1Sl4RaT1oNnQPCRW3BayXRCoLftugfEbYOG1W3N7bZSN2SumP46eT4ubZOo
nXB6vXU+9FaEu64C531Ha60z12iyM6ohhlf+K59s7TjSDUuNJUrue9vDVKsEjGpSL4ypMCW1LHPN3PRV
GP6u8dsIvN4mop73td0krk28bsgykXON0u/+3v+6qeimaWPApQ0t1Sna0I2PQRuuuU7aLU/Urvr3bT96
C4op6iChaUpT9EbgGTRGKBl7T3Pv7a2TVfM2YXYAHpx1pqRROYa5Sv3W0ybGBOdKI8QwpRZDqW79mvUq
pAt653sRb/Pxz83A1+7XyHrg/V0qabAm7zTdXoul9ThaH/ZqaoE2U3wE3vnba68hmYoxNGYEmxZxqdoD
i3f2ylJbme4mPS50Ord1bncjr5+4LcDOxoZRz+R6HuNn80stDd9en27E1+3UjITXftkfJxMHtpaearFs
0jCOkglQrcXStY6QUMusffXAg7aNtivUFC+nFiVbNTy/hhU0dXHz7hXmq9eG9OAy9dA96bif+nnaPErj
qPmH6X8BAAD//7STX3s4DQAA
`,
	},

	"/": {
		name:  "/",
		local: `examples/hotrod/services/frontend/web_assets`,
		isDir: true,
	},
}

var _escDirs = map[string][]os.FileInfo{

	"examples/hotrod/services/frontend/web_assets": {
		_escData["/index.html"],
	},
}
