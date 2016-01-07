// Copyright The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ingore

// Package tar implements access to tar archives. It aims to cover most of the
// variations, including those produced by GNU and BSD tars.

// tar 包实现了 tar 格式压缩文件的存取. 本包目标是覆盖大多数 tar 的变种, 包括
// GNU 和 BSD 生成的 tar 文件.

// References:
//
//	http://www.freebsd.org/cgi/man.cgi?query=tar&sektion=5
//	http://www.gnu.org/software/tar/manual/html_node/Standard.html
//	http://pubs.opengroup.org/onlinepubs/9699919799/utilities/pax.html
package tar

//const (
//	// Types
//	TypeReg           = '0'    // regular file
//	TypeRegA          = '\x00' // regular file
//	TypeLink          = '1'    // hard link
//	TypeSymlink       = '2'    // symbolic link
//	TypeChar          = '3'    // character device node
//	TypeBlock         = '4'    // block device node
//	TypeDir           = '5'    // directory
//	TypeFifo          = '6'    // fifo node
//	TypeCont          = '7'    // reserved
//	TypeXHeader       = 'x'    // extended header
//	TypeXGlobalHeader = 'g'    // global extended header
//	TypeGNULongName   = 'L'    // Next file has a long name
//	TypeGNULongLink   = 'K'    // Next file symlinks to a file w/ a long name
//	TypeGNUSparse     = 'S'    // sparse file
//)

const (
	// 文件类型
	TypeReg           = '0'    // 常规文件
	TypeRegA          = '\x00' // 常规文件
	TypeLink          = '1'    // 硬链接
	TypeSymlink       = '2'    // 符号链接
	TypeChar          = '3'    // 字符设备节点
	TypeBlock         = '4'    // 块设备节点
	TypeDir           = '5'    // 目录
	TypeFifo          = '6'    // 先入先出节点
	TypeCont          = '7'    // 保留
	TypeXHeader       = 'x'    // 扩展头部
	TypeXGlobalHeader = 'g'    // 全局扩展头部
	TypeGNULongName   = 'L'    // 下一个文件文件名很长
	TypeGNULongLink   = 'K'    // 下一个符号链接链接到的文件名称很长
	TypeGNUSparse     = 'S'    // 稀疏文件
)

var (
	ErrWriteTooLong    = errors.New("archive/tar: write too long")
	ErrFieldTooLong    = errors.New("archive/tar: header field too long")
	ErrWriteAfterClose = errors.New("archive/tar: write after close")
)

var (
	ErrHeader = errors.New("archive/tar: invalid tar header")
)

// A Header represents a single header in a tar archive. Some fields may not be
// populated.
//type Header struct {
//	Name       string    // name of header file entry
//	Mode       int64     // permission and mode bits
//	Uid        int       // user id of owner
//	Gid        int       // group id of owner
//	Size       int64     // length in bytes
//	ModTime    time.Time // modified time
//	Typeflag   byte      // type of header entry
//	Linkname   string    // target name of link
//	Uname      string    // user name of owner
//	Gname      string    // group name of owner
//	Devmajor   int64     // major number of character or block device
//	Devminor   int64     // minor number of character or block device
//	AccessTime time.Time // access time
//	ChangeTime time.Time // status change time
//	Xattrs     map[string]string
//}

// Header 代表 tar 档案文件里的单个头. 某些字段可能未被填充使用.
type Header struct {
	Name       string    // 头部名称, 一般设置为文件名全路径
	Mode       int64     // 权限和模式位
	Uid        int       // 用户id
	Gid        int       // 用户组id
	Size       int64     // 按字节表示长度
	ModTime    time.Time // 修改时间
	Typeflag   byte      // 头部条目类型
	Linkname   string    // 链接的目标名称
	Uname      string    // 用户名
	Gname      string    // 用户组名
	Devmajor   int64     // 字符或块主设备号
	Devminor   int64     // 字符或块次设备号
	AccessTime time.Time // 访问时间
	ChangeTime time.Time // 状态改变时间
	Xattrs     map[string]string
}

// FileInfoHeader creates a partially-populated Header from fi. If fi describes a
// symlink, FileInfoHeader records link as the link target. If fi describes a
// directory, a slash is appended to the name. Because os.FileInfo's Name method
// returns only the base name of the file it describes, it may be necessary to
// modify the Name field of the returned header to provide the full path name of
// the file.

// FileInfoHeader 创建一个根据 fi 填写了部分字段的 Header. 如果 fi 描述一个符号链接,
// FileInfoHeader 将 link 参数作为链接目标. 如果 fi 描述一个目录, 则会在名字后面添加
// 斜杠. 因为 os.FileInfo 接口的 Name 方法只返回它描述的文件的无路径名, 有可能需要将
// 返回值的 Name 字段修改为文件的完整路径名.
func FileInfoHeader(fi os.FileInfo, link string) (*Header, error)

// FileInfo returns an os.FileInfo for the Header.

// FileInfo 返回 Header 对应的 os.FileInfo.
func (h *Header) FileInfo() os.FileInfo

// A Reader provides sequential access to the contents of a tar archive. A tar
// archive consists of a sequence of files. The Next method advances to the next
// file in the archive (including the first), and then it can be treated as an
// io.Reader to access the file's data.

// Reader 提供了对一个 tar 档案文件的顺序读取. 一个 tar 档案文件包含一系列文件.
// Next 方法前进到档案中的下一个文件(包括第一个), 返回值可以被视为一个 io.Reader
// 来获取文件的数据.
type Reader struct {
	// contains filtered or unexported fields
}

// NewReader creates a new Reader reading from r.

// NewReader 创建一个从 r 读取数据的 Reader.
func NewReader(r io.Reader) *Reader

// Next advances to the next entry in the tar archive.

// Next 将前进到 tar 归档文件中的下一条记录.
func (tr *Reader) Next() (*Header, error)

// Read reads from the current entry in the tar archive. It returns 0, io.EOF when
// it reaches the end of that entry, until Next is called to advance to the next
// entry.

// 从档案文件的当前记录读取数据, 到达记录末端时返回(0, io.EOF), 直到调用 Next 方法
// 前进到下一记录.
func (tr *Reader) Read(b []byte) (n int, err error)

// A Writer provides sequential writing of a tar archive in POSIX.1 format. A tar
// archive consists of a sequence of files. Call WriteHeader to begin a new file,
// and then call Write to supply that file's data, writing at most hdr.Size bytes
// in total.

// Writer 类型提供了 POSIX.1 格式的 tar 档案文件的顺序写入. 一个 tar 档案文件包含一
// 系列文件. 调用 WriteHeader 来创建一个新的文件, 然后调用 Write 写入文件的数据, 该
// 记录写入的数据不能超过 hdr.Size 字节.
type Writer struct {
	// contains filtered or unexported fields
}

// NewWriter creates a new Writer writing to w.

// NewWriter 创建一个写入数据到 w 的 Writer.
func NewWriter(w io.Writer) *Writer

// Close closes the tar archive, flushing any unwritten data to the underlying
// writer.

// Close 关闭 tar 档案文件, 并将缓冲中未写入下层 io.Writer 接口的数据刷新到下层.
func (tw *Writer) Close() error

// Flush finishes writing the current file (optional).

// Flush 结束当前文件的写入(可选的).
func (tw *Writer) Flush() error

// Write writes to the current entry in the tar archive. Write returns the error
// ErrWriteTooLong if more than hdr.Size bytes are written after WriteHeader.

// Write 向 tar 档案文件的当前记录中写入数据. 如果写入的数据总长度超出上一次调用
// WriteHeader 时的参数 hdr.Size, 会返回 ErrWriteTooLong 错误.
func (tw *Writer) Write(b []byte) (n int, err error)

// WriteHeader writes hdr and prepares to accept the file's contents. WriteHeader
// calls Flush if it is not the first header. Calling after a Close will return
// ErrWriteAfterClose.

// WriteHeader 写入 hdr 并准备接受文件内容. 如果不是第一次调用本方法, 会调用 Flush.
// 在 Close 之后调用本方法会返回 ErrWriteAfterClose.
func (tw *Writer) WriteHeader(hdr *Header) error
