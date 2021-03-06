// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package records

import (
	"encoding/json"
	"gopkg.in/check.v1"
)

func (t *RecordTest) TestARecord(c *check.C) {
	r := &ARecord{}
	c.Assert(r.Validate(), check.Not(check.IsNil))
	err := json.Unmarshal(t.aRecordJson, r)
	c.Assert(err, check.IsNil)
	c.Assert(string(r.Ascii()), check.Equals, string(t.aRecordAscii))
	c.Assert(r.Validate(), check.IsNil)
	err = r.Parse(t.aRecordAscii)
	c.Assert(err, check.IsNil)
	c.Assert(r.Validate(), check.IsNil)
	c.Assert(string(r.Ascii()), check.Equals, string(t.aRecordAscii))
}

func (t *RecordTest) TestARecordWithError(c *check.C) {
	r := &ARecord{}
	err := r.Parse(t.aRecordAscii[1:])
	c.Assert(err, check.Not(check.IsNil))
}
