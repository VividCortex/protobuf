// Extensions for Protocol Buffers to create more go like structures.
//
// Copyright (c) 2013, Vastech SA (PTY) LTD. All rights reserved.
// http://github.com/gogo/protobuf/gogoproto
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are
// met:
//
//     * Redistributions of source code must retain the above copyright
// notice, this list of conditions and the following disclaimer.
//     * Redistributions in binary form must reproduce the above
// copyright notice, this list of conditions and the following disclaimer
// in the documentation and/or other materials provided with the
// distribution.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
// A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
// OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
// LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
// DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
// THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package fieldpath_test

import (
	"bytes"
	"github.com/VividCortex/protobuf/fieldpath"
	"github.com/VividCortex/protobuf/proto"
	"github.com/VividCortex/protobuf/test"
	"math/rand"
	"os"
	"testing"
	"time"
)

func TestString(t *testing.T) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	s := test.NewPopulatedNinOptStruct(r, false)
	data, err := proto.Marshal(s)
	if err != nil {
		panic(err)
	}
	buf := bytes.NewBuffer(nil)
	err = fieldpath.ToString("test", "NinOptStruct", test.ThetestDescription(), "", data, 0, os.Stdout)
	if err != nil {
		panic(err)
	}
	_ = buf
	t.Logf("%v", string(buf.Bytes()))
}
