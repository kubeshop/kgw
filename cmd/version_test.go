/*
The MIT License (MIT)

Copyright © 2022 Kubeshop

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
.
*/
package cmd

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Version(t *testing.T) {
	t.Parallel()

	writer := bytes.NewBufferString("")
	version := "_some-version_"
	date := "_some-date_"
	time := "_some-time-2022-05-20T10:55:55Z_"
	tag := "_some-tag-v2.10.1-19-ge837b7bc_"
	command := NewVersionCommand(writer, version, date, time, tag)
	command.Run(nil, []string{})

	expected := `
_some-version_
_some-tag-v2.10.1-19-ge837b7bc_
_some-time-2022-05-20T10:55:55Z_

https://github.com/kubeshop/kusk/releases/latest
`
	actual := writer.String()

	assert := assert.New(t)

	assert.Equal(expected, actual)
}