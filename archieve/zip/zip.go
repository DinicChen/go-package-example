// Copyright The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ingore

// Package zip provides support for reading and writing ZIP archives.

// zip 包提供了 zip 档案文件的读写服务.

// See: http://www.pkware.com/documents/casestudies/APPNOTE.TXT
//
// This package does not support disk spanning.

// 本包不支持跨硬盘的压缩.

// A note about ZIP64:
//
// To be backwards compatible the FileHeader has both 32 and 64 bit Size fields.
// The 64 bit fields will always contain the correct value and for normal archives
// both fields will be the same. For files requiring the ZIP64 format the 32 bit
// fields will be 0xffffffff and the 64 bit fields must be used instead.

// 关于ZIP64：
//
// 为了向下兼容, FileHeader 同时拥有 32 位和 64 位的 Size 字段. 64 位字段总是包含
// 正确的值, 对普通格式的档案来说它们的值是相同的. 对 zip64 格式的档案文件 32 位字
// 段将是 0xffffffff, 必须使用64位字段.
package zip

// Compression methods.

// 预定义压缩算法.
const (
	Store   uint16 = 0
	Deflate uint16 = 8
)

var (
	ErrFormat    = errors.New("zip: not a valid zip file")
	ErrAlgorithm = errors.New("zip: unsupported compression algorithm")
	ErrChecksum  = errors.New("zip: checksum error")
)

// RegisterCompressor registers custom compressors for a specified method ID. The
// common methods Store and Deflate are built in.

// RegisterCompressor 使用指定的方法 ID 注册一个 Compressor 类型函数. 常用的方法
// Store 和 Deflate 是内建的.
func RegisterCompressor(method uint16, comp Compressor)

// RegisterDecompressor allows custom decompressors for a specified method ID.

// RegisterDecompressor 使用指定的方法 ID 注册一个 Decompressor 类型函数.
func RegisterDecompressor(method uint16, d Decompressor)

// A Compressor returns a compressing writer, writing to the provided writer. On
// Close, any pending data should be flushed.

// Compressor 函数类型会返回一个 io.WriteCloser, 该接口会将数据压缩后写入提供的
// io.Writer. 关闭时, 应将缓冲中的数据刷新到下层接口中.
type Compressor func(io.Writer) (io.WriteCloser, error)

// Decompressor is a function that wraps a Reader with a decompressing Reader. The
// decompressed ReadCloser is returned to callers who open files from within the
// archive. These callers are responsible for closing this reader when they're
// finished reading.

// Decompressor 函数类型会把一个 io.Reader 包装成具有解压缩功能的 io.ReadCloser.
// io.ReadCloser 被返回给打开档案内文件的调用者. 这些调用者有责任在读取结束时关闭
// 该 io.ReadCloser.
type Decompressor func(io.Reader) io.ReadCloser

type File struct {
	FileHeader
	// contains filtered or unexported fields
}

// DataOffset returns the offset of the file's possibly-compressed data, relative
// to the beginning of the zip file.
//
// Most callers should instead use Open, which transparently decompresses data and
// verifies checksums.

// DataOffset 返回文件中可能存在的压缩数据相对于 zip 文件起始处的偏移量. 大多数调用
// 者应使用 Open 代替, 该方法会主动解压缩数据并验证校验和.
func (f *File) DataOffset() (offset int64, err error)

// Open returns a ReadCloser that provides access to the File's contents. Multiple
// files may be read concurrently.

// Open 方法返回一个 io.ReadCloser 接口, 提供读取文件内容的方法. 可以同时读取多个文件.
func (f *File) Open() (rc io.ReadCloser, err error)

// FileHeader describes a file within a zip file. See the zip spec for details.

// FileHeader 描述 zip 文件中的一个文件. 参见zip的定义获取细节.
//type FileHeader struct {
//	// Name is the name of the file.
//	// It must be a relative path: it must not start with a drive
//	// letter (e.g. C:) or leading slash, and only forward slashes
//	// are allowed.
//	Name string
//
//	CreatorVersion     uint16
//	ReaderVersion      uint16
//	Flags              uint16
//	Method             uint16
//	ModifiedTime       uint16 // MS-DOS time
//	ModifiedDate       uint16 // MS-DOS date
//	CRC32              uint32
//	CompressedSize     uint32 // deprecated; use CompressedSize64
//	UncompressedSize   uint32 // deprecated; use UncompressedSize64
//	CompressedSize64   uint64
//	UncompressedSize64 uint64
//	Extra              []byte
//	ExternalAttrs      uint32 // Meaning depends on CreatorVersion
//	Comment            string
//}
type FileHeader struct {
	// Name 是这个文件的名称
	// 该名称必须是相对路径，不可以以盘符(比如C:)或者反斜杠(\)开头。允许正斜杠(/)存在。
	Name string

	CreatorVersion     uint16
	ReaderVersion      uint16
	Flags              uint16
	Method             uint16
	ModifiedTime       uint16 // MS-DOS 时间
	ModifiedDate       uint16 // MS-DOS 日期
	CRC32              uint32
	CompressedSize     uint32 // 已废弃; 使用 CompressedSize64
	UncompressedSize   uint32 // 已废弃; 使用 UncompressedSize64
	CompressedSize64   uint64
	UncompressedSize64 uint64
	Extra              []byte
	ExternalAttrs      uint32 // 参数含义依赖于CreatorVersion
	Comment            string
}

// FileInfoHeader creates a partially-populated FileHeader from an os.FileInfo.
// Because os.FileInfo's Name method returns only the base name of the file it
// describes, it may be necessary to modify the Name field of the returned header
// to provide the full path name of the file.

// FileInfoHeader 创建一个根据 fi 填写了部分字段的 FileHeader. 因为 os.FileInfo
// 接口的 Name 方法只返回它描述的文件的无路径名, 有可能需要将返回值的 Name 字段修
// 改为文件的完整路径名.
func FileInfoHeader(fi os.FileInfo) (*FileHeader, error)

// FileInfo returns an os.FileInfo for the FileHeader.

// FileInfo 返回一个根据 h 的信息生成的 os.FileInfo.
func (h *FileHeader) FileInfo() os.FileInfo

// ModTime returns the modification time in UTC. The resolution is 2s.

// ModTime 返回最近一次修改的 UTC 时间. 精度2s.
func (h *FileHeader) ModTime() time.Time

// Mode returns the permission and mode bits for the FileHeader.

// Mode 返回 h 的权限和模式位.
func (h *FileHeader) Mode() (mode os.FileMode)

// SetModTime sets the ModifiedTime and ModifiedDate fields to the given time in
// UTC. The resolution is 2s.

// SetModTime 将 ModifiedTime 和 ModifiedDate 字段设置为给定的 UTC 时间. 精度2s.
func (h *FileHeader) SetModTime(t time.Time)

// SetMode changes the permission and mode bits for the FileHeader.

// SetMode 修改 h 的权限和模式位.
func (h *FileHeader) SetMode(mode os.FileMode)

type ReadCloser struct {
	Reader
	// contains filtered or unexported fields
}

// OpenReader will open the Zip file specified by name and return a ReadCloser.

// OpenReader 会打开 name 指定的 zip 文件并返回一个 ReadCloser.
func OpenReader(name string) (*ReadCloser, error)

// Close closes the Zip file, rendering it unusable for I/O.

// Close 关闭 zip 文件, 使它不能用于 I/O.
func (rc *ReadCloser) Close() error

type Reader struct {
	File    []*File
	Comment string
	// contains filtered or unexported fields
}

// NewReader returns a new Reader reading from r, which is assumed to have the
// given size in bytes.

// NewReader 创建并返回一个从 r 读取数据的 Reader, r 被假设其大小为 size 字节.
func NewReader(r io.ReaderAt, size int64) (*Reader, error)

// Writer implements a zip file writer.

// Writer 类型实现了 zip 文件的写入器.
type Writer struct {
	// contains filtered or unexported fields
}

// NewWriter returns a new Writer writing a zip file to w.

// NewWriter 创建并返回一个将 zip 文件写入 w 的 Writer.
func NewWriter(w io.Writer) *Writer

// Close finishes writing the zip file by writing the central directory. It does
// not (and can not) close the underlying writer.

// Close 方法通过写入中央目录来关闭该 Writer. 本方法不会也没办法关闭下层的 io.Writer 接口.
func (w *Writer) Close() error

// Create adds a file to the zip file using the provided name. It returns a Writer
// to which the file contents should be written. The name must be a relative path:
// it must not start with a drive letter (e.g. C:) or leading slash, and only
// forward slashes are allowed. The file's contents must be written to the
// io.Writer before the next call to Create, CreateHeader, or Close.

// Create 使用给定的文件名添加一个文件进 zip 文件. 本方法返回一个 io.Writer 接口(用于写入
// 新文件的内容). 文件名必须是相对路径, 不能以设备或斜杠开始, 只接受'/'作为路径
// 分隔. 新文件的内容必须在下一次调用 CreateHeader、Create 或 Close 方法之前全部写入.
func (w *Writer) Create(name string) (io.Writer, error)

// CreateHeader adds a file to the zip file using the provided FileHeader for the
// file metadata. It returns a Writer to which the file contents should be written.
// The file's contents must be written to the io.Writer before the next call to
// Create, CreateHeader, or Close.

// CreateHeader 使用给定的 FileHeader 来作为文件的元数据添加一个文件进 zip 文件.
// 本方法返回一个 io.Writer 接口(用于写入新文件的内容). 新文件的内容必须在下一次调
// 用 CreateHeader、Create 或 Close 方法之前全部写入.
func (w *Writer) CreateHeader(fh *FileHeader) (io.Writer, error)

// Flush flushes any buffered data to the underlying writer. Calling Flush is not
// normally necessary; calling Close is sufficient.

// Flush 将缓冲中的数据刷新到下层 Writer 中. 调用 Flush 不是必要的; 调用 Close 就足够了.
func (w *Writer) Flush() error
