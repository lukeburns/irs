// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package records

import (
	"io/ioutil"
	"path/filepath"
	"testing"

	"gopkg.in/check.v1"
)

func Test(t *testing.T) { check.TestingT(t) }

type RecordTest struct {
	tRecordJson          []byte
	tRecordAscii         []byte
	aRecordJson          []byte
	aRecordAscii         []byte
	bRecord1099MiscJson  []byte
	bRecord1099MiscAscii []byte
	bRecord1099IntJson   []byte
	bRecord1099IntAscii  []byte
	bRecord1099OidJson   []byte
	bRecord1099OidAscii  []byte
	bRecord1099PatrJson  []byte
	bRecord1099PatrAscii []byte
	bRecord1097BtcJson   []byte
	bRecord1097BtcAscii  []byte
	cRecordJson          []byte
	cRecordAscii         []byte
	kRecordJson          []byte
	kRecordAscii         []byte
	fRecordJson          []byte
	fRecordAscii         []byte
}

var _ = check.Suite(&RecordTest{})

func (t *RecordTest) SetUpSuite(c *check.C) {
	var err error

	t.tRecordJson, err = ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "transmitterRecord.json"))
	c.Assert(err, check.IsNil)

	t.tRecordAscii, err = ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "transmitterRecord.ascii"))
	c.Assert(err, check.IsNil)

	t.aRecordJson, err = ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "payerRecord.json"))
	c.Assert(err, check.IsNil)

	t.aRecordAscii, err = ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "payerRecord.ascii"))
	c.Assert(err, check.IsNil)

	t.bRecord1099MiscJson, err = ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "payeeRecordWith1099Misc.json"))
	c.Assert(err, check.IsNil)

	t.bRecord1099MiscAscii, err = ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "payeeRecordWith1099Misc.ascii"))
	c.Assert(err, check.IsNil)

	t.bRecord1099IntJson, err = ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "payeeRecordWith1099Int.json"))
	c.Assert(err, check.IsNil)

	t.bRecord1099IntAscii, err = ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "payeeRecordWith1099Int.ascii"))
	c.Assert(err, check.IsNil)

	t.bRecord1099OidJson, err = ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "payeeRecordWith1099Oid.json"))
	c.Assert(err, check.IsNil)

	t.bRecord1099OidAscii, err = ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "payeeRecordWith1099Oid.ascii"))
	c.Assert(err, check.IsNil)

	t.bRecord1099PatrJson, err = ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "payeeRecordWith1099Patr.json"))
	c.Assert(err, check.IsNil)

	t.bRecord1099PatrAscii, err = ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "payeeRecordWith1099Patr.ascii"))
	c.Assert(err, check.IsNil)

	t.bRecord1097BtcJson, err = ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "payeeRecordWith1097Btc.json"))
	c.Assert(err, check.IsNil)

	t.bRecord1097BtcAscii, err = ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "payeeRecordWith1097Btc.ascii"))
	c.Assert(err, check.IsNil)

	t.cRecordJson, err = ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "endPayerRecord.json"))
	c.Assert(err, check.IsNil)

	t.cRecordAscii, err = ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "endPayerRecord.ascii"))
	c.Assert(err, check.IsNil)

	t.kRecordJson, err = ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "stateRecord.json"))
	c.Assert(err, check.IsNil)

	t.kRecordAscii, err = ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "stateRecord.ascii"))
	c.Assert(err, check.IsNil)

	t.fRecordJson, err = ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "endTransmitterRecord.json"))
	c.Assert(err, check.IsNil)

	t.fRecordAscii, err = ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "endTransmitterRecord.ascii"))
	c.Assert(err, check.IsNil)
}
