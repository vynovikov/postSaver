// Helper pachage for types and functions
package repo

import (
	"github.com/vynovikov/postSaver/internal/adapters/driver/rpc/pb"
)

type ReqData struct {
	IsFitst   bool
	IsLast    bool
	TS        string
	Name      string
	FileName  string
	Number    int
	ByteChunk []byte
}

type NameNumber struct {
	Name   string
	Number int
}

func NewNameNumber(name string, number int) NameNumber {
	return NameNumber{
		Name:   name,
		Number: number,
	}
}

type ReqUnary struct {
	R *pb.TextFieldReq
}

func NewReqUnary(r *pb.TextFieldReq) *ReqUnary {
	return &ReqUnary{
		R: r,
	}
}

type ReqStream struct {
	R *pb.FileUploadReq
}

func NewReqStream(r *pb.FileUploadReq) *ReqStream {
	return &ReqStream{
		R: r,
	}
}

type Request interface {
	IsFirst() bool       // Returns IsFirst
	IsLast() bool        // Returns IsFirst
	TS() string          // Returns TS
	Name() string        // Returns FieldName
	FileName() string    // Returns FileName
	Number() int         // Returns Number
	GetBody() []byte     // Returns request body
	IsUnary() bool       // Returns true if request is an unary
	IsStreamInfo() bool  // Returns true if request is a stream and has FileUploadReq_FileInfo field
	IsStreamData() bool  // Returns true if request is a stream and has FileUploadReq_FileData field
	Unwrap() interface{} // Returns initial request
}

func (u *ReqUnary) IsFirst() bool {
	return u.R.IsFirst
}
func (u *ReqUnary) IsLast() bool {
	return u.R.IsLast
}
func (u *ReqUnary) TS() string {
	return u.R.Ts
}
func (u *ReqUnary) Name() string {
	return u.R.Name
}
func (u *ReqUnary) FileName() string {
	return u.R.Filename
}
func (u *ReqUnary) Number() int {
	return 0
}
func (u *ReqUnary) GetBody() []byte {
	return u.R.ByteChunk
}
func (u *ReqUnary) IsUnary() bool {
	return true
}
func (u *ReqUnary) IsStreamInfo() bool {
	return false
}
func (u *ReqUnary) IsStreamData() bool {
	return false
}
func (u *ReqUnary) Unwrap() interface{} {
	return u.R
}

func (s *ReqStream) IsFirst() bool {
	if s.R.GetFileInfo() != nil {
		return s.R.GetFileInfo().IsFirst
	}
	return false
}
func (s *ReqStream) IsLast() bool {
	if s.R.GetFileData() != nil {
		return s.R.GetFileData().IsLast
	}
	return false
}
func (s *ReqStream) TS() string {
	if s.R.GetFileData() != nil {
		return s.R.GetFileData().Ts
	}
	if s.R.GetFileInfo() != nil {
		return s.R.GetFileInfo().Ts
	}
	return ""
}
func (s *ReqStream) Name() string {
	if s.R.GetFileData() != nil {
		return s.R.GetFileData().FieldName
	}
	return s.R.GetFileInfo().FieldName
}
func (s *ReqStream) FileName() string {
	if s.R.GetFileData() != nil {
		return ""
	}
	return s.R.GetFileInfo().FileName
}
func (s *ReqStream) Number() int {
	if s.R.GetFileData() != nil {
		return int(s.R.GetFileData().Number)
	}
	return 0
}
func (s *ReqStream) GetBody() []byte {
	if s.R.GetFileData() != nil {
		return s.R.GetFileData().ByteChunk
	}
	return nil
}
func (u *ReqStream) IsUnary() bool {
	return false
}

func (s *ReqStream) IsStreamInfo() bool {
	if s.R.GetFileData() != nil {
		return false
	}
	return true
}

func (s *ReqStream) IsStreamData() bool {
	return !s.IsStreamInfo()
}
func (s *ReqStream) Unwrap() interface{} {
	return s.R
}

type IDToRemove struct {
	TS string
	I  int
}

func NewIDToRemove(ts string, i int) IDToRemove {
	return IDToRemove{
		TS: ts,
		I:  i,
	}
}

type IDsToRemove struct {
	TS string
	I  []int
}

// Add appends new element to *IDsToRemove.
// Tested in models_test.go
func (i *IDsToRemove) Add(j IDToRemove) {
	if i.TS == "" {
		i.TS = j.TS
		i.I = append([]int{}, j.I)
		return
	}
	if i.TS == j.TS {
		i.I = append(i.I, j.I)
	}
}
