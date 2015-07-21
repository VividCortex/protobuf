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

package fieldpath

import (
	"fmt"
	descriptor "github.com/VividCortex/protobuf/protoc-gen-gogo/descriptor"
	"strings"
)

type errUndefinedMsg struct {
	pkg string
	msg string
}

func (this *errUndefinedMsg) Error() string {
	return fmt.Sprintf("undefined msg %s.%s", this.pkg, this.msg)
}

func ToMessages(path string, fileDescriptorSet *descriptor.FileDescriptorSet) ([]string, error) {
	elems := strings.Split(path, ".")
	return toMessages(elems, fileDescriptorSet, make([]string, 0, len(elems)))
}

func toMessages(elems []string, fileDescriptorSet *descriptor.FileDescriptorSet, messages []string) ([]string, error) {
	if len(elems) < 2 {
		return nil, &errUndefinedMsg{elems[0], "???"}
	}
	msg := fileDescriptorSet.GetMessage(elems[0], elems[1])
	if msg == nil {
		return nil, &errUndefinedMsg{elems[0], elems[1]}
	}
	if len(elems) == 2 {
		messages = append(messages, elems[0]+"."+elems[1])
		return messages, nil
	}
	newPkg, newMsg := fileDescriptorSet.FindMessage(elems[0], elems[1], elems[2])
	if len(newMsg) == 0 {
		return nil, &errUndefined{elems[0], elems[1], elems[2]}
	}
	messages = append(messages, elems[0]+"."+elems[1])
	elems[1] = newPkg
	elems[2] = newMsg
	return toMessages(elems[1:], fileDescriptorSet, messages)
}
