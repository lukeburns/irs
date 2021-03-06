// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package records

import (
	"encoding/json"
	"gopkg.in/check.v1"
)

func (t *RecordTest) TestFRecord(c *check.C) {
	r := NewFRecord()
	c.Assert(r.Validate(), check.Not(check.IsNil))
	err := json.Unmarshal(t.fRecordJson, r)
	c.Assert(err, check.IsNil)
	c.Assert(string(r.Ascii()), check.Equals, string(t.fRecordAscii))
	c.Assert(r.Validate(), check.IsNil)
	err = r.Parse(t.fRecordAscii)
	c.Assert(err, check.IsNil)
	c.Assert(r.Validate(), check.IsNil)
	c.Assert(string(r.Ascii()), check.Equals, string(t.fRecordAscii))
}

func (t *RecordTest) TestFRecordWithError(c *check.C) {
	r := &FRecord{}
	err := r.Parse(t.fRecordAscii[1:])
	c.Assert(err, check.Not(check.IsNil))
}
