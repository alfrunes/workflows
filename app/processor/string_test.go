// Copyright 2022 Northern.tech AS
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.

package processor

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetOption(t *testing.T) {
	var expectedOption *Options
	var o Options

	expectedOption = &Options{Encoding: EncodingPlain}
	o = getOption()
	assert.Equal(t, *expectedOption, o)

	expectedOption = &Options{Encoding: EncodingPlain}
	o = getOption(expectedOption)
	assert.Equal(t, *expectedOption, o)

	expectedOption = &Options{Encoding: EncodingURL}
	o = getOption(expectedOption)
	assert.Equal(t, *expectedOption, o)

	expectedOption = &Options{Encoding: EncodingURL}
	o = getOption(expectedOption, &Options{Encoding: EncodingPlain})
	assert.Equal(t, *expectedOption, o)
}

func TestEncodingFromData(t *testing.T) {
	var data string
	var encoding Encoding
	var setEncoding string

	setEncoding = plainEncodingFlag
	data = "${encoding=" + setEncoding + ";workflow.input.variable0}"
	data, encoding = getEncodingFromData(data, EncodingURL)
	assert.Equal(t, "workflow.input.variable0", data)
	assert.Equal(t, EncodingPlain, encoding)

	setEncoding = urlEncodingFlag
	data = "${encoding=" + setEncoding + ";workflow.input.variable0}"
	data, encoding = getEncodingFromData(data, EncodingPlain)
	assert.Equal(t, "workflow.input.variable0", data)
	assert.Equal(t, EncodingURL, encoding)

	setEncoding = urlEncodingFlag
	data = "${encoding=" + setEncoding + ";workflow.input.variable0}"
	data, encoding = getEncodingFromData(data, EncodingURL)
	assert.Equal(t, "workflow.input.variable0", data)
	assert.Equal(t, EncodingURL, encoding)

	data = "${workflow.input.variable0}"
	data, encoding = getEncodingFromData(data, EncodingURL)
	assert.Equal(t, "${workflow.input.variable0}", data)
	assert.Equal(t, EncodingURL, encoding)

	data = "${workflow.input.variable0}"
	data, encoding = getEncodingFromData(data, EncodingPlain)
	assert.Equal(t, "${workflow.input.variable0}", data)
	assert.Equal(t, EncodingPlain, encoding)
}
