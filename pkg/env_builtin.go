package jolt

import _ "embed"

//go:embed builtin/c++.jolt
var BuiltinCpp string

//go:embed builtin/gcc.jolt
var BuiltinGcc string

//go:embed builtin/clang.jolt
var BuiltinClang string

//go:embed builtin/clang-tidy.jolt
var BuiltinClangTidy string

//go:embed builtin/cppcheck.jolt
var BuiltinCppcheck string

//go:embed builtin/protobuf.jolt
var BuiltinProtobuf string

var Builtin = []*string{
	&BuiltinCpp,
	&BuiltinGcc,
	&BuiltinCppcheck,
	&BuiltinClang,
	&BuiltinClangTidy,
	&BuiltinProtobuf,
}
