// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package subrecords

import (
	"bytes"
	"reflect"
	"unicode/utf8"

	"github.com/moov-io/irs/pkg/config"
	"github.com/moov-io/irs/pkg/utils"
)

type Sub1099OID struct {
	// Enter “2” (two) to indicate notification by the IRS twice within
	// three calendar years that the payee provided an incorrect
	// name and/or TIN combination. Otherwise, enter a blank.
	SecondTinNotice string `json:"second_tin_notice"`

	// Required. Enter the CUSIP number, if any. If there is no
	// CUSIP number, enter the abbreviation for the stock
	// exchange and issuer, the coupon rate, and year of maturity
	// (must be four-digit year). For example, NYSE XYZ 12/2019.
	// Show the name of the issuer if other than the payer. If fewer
	// than 39 characters are required, left justify the information
	// and fill unused positions with blanks.
	Description string `json:"direct_sales_indicator" validate:"required"`

	// Enter "1" (one) if there is FATCA filing requirement.
	// Otherwise, enter a blank.
	FATCA string `json:"fatca_requirement_indicator"`

	// This portion of the “B” Record may be used to record
	// information for state or local government reporting or for the
	// filer’s own purposes. Payers should contact the state or local
	// revenue departments for filing requirements. If this field is not
	// used, enter blanks.
	SpecialDataEntries string `json:"special_data_entries"`

	// State income tax withheld is for the convenience of the filers.
	// This information does not need to be reported to the IRS.
	// The payment amount must be right justified and unused
	// positions must be zero-filed. If not reporting state income tax
	// withheld, this field may be used as a continuation of the
	// Special Data Entries field.
	StateIncomeTaxWithheld int `json:"state_income_tax_withheld"`

	// Local income tax withheld is for the convenience of the filers.
	// This information does not need to be reported to the IRS.
	// The payment amount must be right justified and unused
	// positions must be zero-filled. If not reporting local tax
	// withheld, this field may be used as a continuation of the
	// Special Data Entries Field.
	LocalIncomeTaxWithheld int `json:"local_income_tax_withheld"`

	// Enter the valid CF/SF code if this payee record is to be
	// forwarded to a state agency as part of the CF/SF Program.
	CombinedFSCode int `json:"combined_federal_state_code"`
}

// Type returns type of “1099-OID” record
func (r *Sub1099OID) Type() string {
	return config.Sub1099OidType
}

// Parse parses the “1099-OID” record from fire ascii
func (r *Sub1099OID) Parse(buf []byte) error {
	record := string(buf)
	if utf8.RuneCountInString(record) != config.SubRecordLength {
		return utils.ErrRecordLength
	}

	fields := reflect.ValueOf(r).Elem()
	if !fields.IsValid() {
		return utils.ErrValidField
	}

	return utils.ParseValue(fields, config.Sub1099OIDLayout, record)
}

// Ascii returns fire ascii of “1099-OID” record
func (r *Sub1099OID) Ascii() []byte {
	var buf bytes.Buffer
	records := config.ToSpecifications(config.Sub1099OIDLayout)
	fields := reflect.ValueOf(r).Elem()
	if !fields.IsValid() {
		return nil
	}

	buf.Grow(config.SubRecordLength)
	for _, spec := range records {
		value := utils.ToString(spec.Field, fields.FieldByName(spec.Name))
		buf.WriteString(value)
	}

	return buf.Bytes()
}

// Validate performs some checks on the record and returns an error if not Validated
func (r *Sub1099OID) Validate() error {
	return utils.Validate(r, config.Sub1099OIDLayout)
}

// customized field validation functions
// function name should be "Validate" + field name

func (r *Sub1099OID) ValidateFATCA() error {
	if r.FATCA == config.FatcaFilingRequirementIndicator || len(r.FATCA) == 0 {
		return nil
	}
	return utils.NewErrValidValue("fatca filing requirement indicator")
}

func (r *Sub1099OID) ValidateCombinedFSCode() error {
	if _, ok := config.ParticipateStateCodes[r.CombinedFSCode]; ok {
		return nil
	}
	return utils.NewErrValidValue("combined federal state code")
}
