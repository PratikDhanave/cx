// +build os

package cxos

import (
	"github.com/skycoin/cx/cxparsergenerator/actions"
)

const (
	// os
	CONST_OS_RUN_SUCCESS = iota + 0xFFFF
	CONST_OS_RUN_EMPTY_CMD
	CONST_OS_RUN_PANIC
	CONST_OS_RUN_START_FAILED
	CONST_OS_RUN_WAIT_FAILED
	CONST_OS_RUN_TIMEOUT

	// os.FileModes (FIXME: these are uint32 in Go, with _DIR & _TYPE exceeding math.MaxInt32)
	//CONST_OS_FILEMODE_DIR
	//CONST_OS_FILEMODE_TYPE
	CONST_OS_FILEMODE_APPEND
	CONST_OS_FILEMODE_EXCLUSIVE
	CONST_OS_FILEMODE_TEMPORARY
	CONST_OS_FILEMODE_SYMLINK
	CONST_OS_FILEMODE_DEVICE
	CONST_OS_FILEMODE_NAMED_PIPE
	CONST_OS_FILEMODE_SOCKET
	CONST_OS_FILEMODE_SETUID
	CONST_OS_FILEMODE_SETGID
	CONST_OS_FILEMODE_CHAR_DEVICE
	CONST_OS_FILEMODE_STICKY
	CONST_OS_FILEMODE_IRREGULAR
	CONST_OS_FILEMODE_PERM

	CONST_OS_SEEK_SET
	CONST_OS_SEEK_CUR
	CONST_OS_SEEK_END

	// json
	CONST_JSON_TOKEN_NULL
	CONST_JSON_TOKEN_DELIM
	CONST_JSON_TOKEN_BOOL
	CONST_JSON_TOKEN_F64
	CONST_JSON_TOKEN_NUMBER
	CONST_JSON_TOKEN_STR
	CONST_JSON_DELIM_CURLY_LEFT
	CONST_JSON_DELIM_CURLY_RIGHT
	CONST_JSON_DELIM_SQUARE_LEFT
	CONST_JSON_DELIM_SQUARE_RIGHT
)

const (
	OS_RUN_SUCCESS = iota
	OS_RUN_EMPTY_CMD
	OS_RUN_PANIC // 2
	OS_RUN_START_FAILED
	OS_RUN_WAIT_FAILED
	OS_RUN_TIMEOUT
)

func init() {
	// os
	actions.AddConstI32(CONST_OS_RUN_SUCCESS, "os.RUN_SUCCESS", OS_RUN_SUCCESS)
	actions.AddConstI32(CONST_OS_RUN_EMPTY_CMD, "os.RUN_EMPTY_CMD", OS_RUN_EMPTY_CMD)
	actions.AddConstI32(CONST_OS_RUN_PANIC, "os.RUN_PANIC", OS_RUN_PANIC)
	actions.AddConstI32(CONST_OS_RUN_START_FAILED, "os.RUN_START_FAILED", OS_RUN_START_FAILED)
	actions.AddConstI32(CONST_OS_RUN_WAIT_FAILED, "os.RUN_WAIT_FAILED", OS_RUN_WAIT_FAILED)
	actions.AddConstI32(CONST_OS_RUN_TIMEOUT, "os.RUN_TIMEOUT", OS_RUN_TIMEOUT)

	// os.FileModes (FIXME: these are uint32 in Go, with _DIR & _TYPE exceeding math.MaxInt32)
	//cxcore.AddConstI32(CONST_OS_FILEMODE_DIR, "os.ModeDir", 2147483648)
	//cxcore.AddConstI32(CONST_OS_FILEMODE_TYPE, "os.ModeType", 2399666176)
	actions.AddConstI32(CONST_OS_FILEMODE_APPEND, "os.ModeAppend", 1073741824)
	actions.AddConstI32(CONST_OS_FILEMODE_EXCLUSIVE, "os.ModeExclusive", 536870912)
	actions.AddConstI32(CONST_OS_FILEMODE_TEMPORARY, "os.ModeTemporary", 268435456)
	actions.AddConstI32(CONST_OS_FILEMODE_SYMLINK, "os.ModeSymlink", 134217728)
	actions.AddConstI32(CONST_OS_FILEMODE_DEVICE, "os.ModeDevice", 67108864)
	actions.AddConstI32(CONST_OS_FILEMODE_NAMED_PIPE, "os.ModeNamedPipe", 33554432)
	actions.AddConstI32(CONST_OS_FILEMODE_SOCKET, "os.ModeSocket", 16777216)
	actions.AddConstI32(CONST_OS_FILEMODE_SETUID, "os.ModeSetuid", 8388608)
	actions.AddConstI32(CONST_OS_FILEMODE_SETGID, "os.ModeSetgid", 4194304)
	actions.AddConstI32(CONST_OS_FILEMODE_CHAR_DEVICE, "os.ModeCharDevice", 2097152)
	actions.AddConstI32(CONST_OS_FILEMODE_STICKY, "os.ModeSticky", 1048576)
	actions.AddConstI32(CONST_OS_FILEMODE_IRREGULAR, "os.ModeIrregular", 524288)
	actions.AddConstI32(CONST_OS_FILEMODE_PERM, "os.ModePerm", 511)

	actions.AddConstI32(CONST_OS_SEEK_SET, "os.SEEK_SET", OS_SEEK_SET)
	actions.AddConstI32(CONST_OS_SEEK_CUR, "os.SEEK_CUR", OS_SEEK_CUR)
	actions.AddConstI32(CONST_OS_SEEK_END, "os.SEEK_END", OS_SEEK_END)

	// json
	actions.AddConstI32(CONST_JSON_TOKEN_NULL, "json.TOKEN_NULL", JSON_TOKEN_NULL)
	actions.AddConstI32(CONST_JSON_TOKEN_DELIM, "json.TOKEN_DELIM", JSON_TOKEN_DELIM)
	actions.AddConstI32(CONST_JSON_TOKEN_BOOL, "json.TOKEN_BOOL", JSON_TOKEN_BOOL)
	actions.AddConstI32(CONST_JSON_TOKEN_F64, "json.TOKEN_F64", JSON_TOKEN_F64)
	actions.AddConstI32(CONST_JSON_TOKEN_NUMBER, "json.TOKEN_NUMBER", JSON_TOKEN_NUMBER)
	actions.AddConstI32(CONST_JSON_TOKEN_STR, "json.TOKEN_STR", JSON_TOKEN_STR)
	actions.AddConstI32(CONST_JSON_DELIM_CURLY_LEFT, "json.DELIM_CURLY_LEFT", JSON_DELIM_CURLY_LEFT)
	actions.AddConstI32(CONST_JSON_DELIM_CURLY_RIGHT, "json.DELIM_CURLY_RIGHT", JSON_DELIM_CURLY_RIGHT)
	actions.AddConstI32(CONST_JSON_DELIM_SQUARE_LEFT, "json.DELIM_SQUARE_LEFT", JSON_DELIM_SQUARE_LEFT)
	actions.AddConstI32(CONST_JSON_DELIM_SQUARE_RIGHT, "json.DELIM_SQUARE_RIGHT", JSON_DELIM_SQUARE_RIGHT)
}
