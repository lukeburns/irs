package file

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/moov-io/irs/pkg/config"
	"github.com/moov-io/irs/pkg/records"
	"github.com/moov-io/irs/pkg/utils"
)

//  paymentPerson identifies the person making payments
type paymentPerson struct {
	Payer    records.Record   `json:"payer"`
	Payees   []records.Record `json:"payees"`
	EndPayer records.Record   `json:"end_payer"`
	States   []records.Record `json:"states"`
}

// Type returns type of “Person” record
func (p *paymentPerson) Type() string {
	return "Person"
}

// Ascii returns fire ascii of “Person” record
func (p *paymentPerson) Ascii() []byte {
	var buf bytes.Buffer

	buf.Grow(config.RecordLength)
	buf.Write(p.Payer.Ascii())

	for _, payee := range p.Payees {
		buf.Grow(config.RecordLength)
		buf.Write(payee.Ascii())
	}

	buf.Grow(config.RecordLength)
	buf.Write(p.EndPayer.Ascii())

	for _, state := range p.States {
		buf.Grow(config.RecordLength)
		buf.Write(state.Ascii())
	}

	return buf.Bytes()
}

// Validate performs some checks on the record and returns an error if not Validated
func (p *paymentPerson) Validate() error {
	if p.Payer == nil || p.EndPayer == nil {
		return utils.ErrInvalidFile
	}

	err := p.Payer.Validate()
	if err != nil {
		return err
	}

	for _, payee := range p.Payees {
		err = payee.Validate()
		if err != nil {
			return err
		}
	}

	err = p.EndPayer.Validate()
	if err != nil {
		return err
	}

	for _, state := range p.States {
		err = state.Validate()
		if err != nil {
			return err
		}
	}

	return nil
}

// SequenceNumber returns sequence number of the record
func (p *paymentPerson) SequenceNumber() int {
	if p.Payer == nil {
		return 0
	}
	return p.Payer.SequenceNumber()
}

// SequenceNumber set sequence number of the record
func (p *paymentPerson) SetSequenceNumber(int) {}

// Parse attempts to parse with raw data.
func (p *paymentPerson) Parse(buf []byte) (int, error) {
	bufSize := len(buf)
	readPtr := 0

	if string(buf[readPtr]) != config.ARecordType || bufSize < config.RecordLength {
		return readPtr, utils.ErrInvalidAscii
	}

	if p.Payer == nil {
		p.Payer = records.NewARecord()
	}
	err := p.Payer.Parse(buf[readPtr : readPtr+config.RecordLength])
	if err != nil {
		return readPtr, err
	}
	readPtr += config.RecordLength

	typeOfReturn := ""
	if p.Payer != nil {
		if arec, ok := p.Payer.(*records.ARecord); !ok {
			return -1, fmt.Errorf("unexpected Payer to be an ARecord, but got %T", p.Payer)
		} else {
			typeOfReturn = config.TypeOfReturns[arec.TypeOfReturn]
		}
	}

	p.Payees = []records.Record{}
	for string(buf[readPtr]) == config.BRecordType {
		if bufSize < readPtr+config.RecordLength {
			return readPtr, utils.ErrInvalidAscii
		}

		newPayee := records.NewBRecord(typeOfReturn)
		if err = newPayee.Parse(buf[readPtr : readPtr+config.RecordLength]); err != nil {
			return readPtr, err
		}

		readPtr += config.RecordLength
		p.Payees = append(p.Payees, newPayee)
	}

	if string(buf[readPtr]) == config.CRecordType || bufSize < readPtr+config.RecordLength {
		if p.EndPayer == nil {
			p.EndPayer = records.NewCRecord()
		}
		err := p.EndPayer.Parse(buf[readPtr : readPtr+config.RecordLength])
		if err != nil {
			return readPtr, err
		}
		readPtr += config.RecordLength
	}

	p.States = []records.Record{}
	for string(buf[readPtr]) == config.KRecordType {
		if bufSize < readPtr+config.RecordLength {
			return readPtr, utils.ErrInvalidAscii
		}

		newState := records.NewKRecord()
		if err = newState.Parse(buf[readPtr : readPtr+config.RecordLength]); err != nil {
			return readPtr, err
		}

		readPtr += config.RecordLength
		p.States = append(p.States, newState)
	}

	return readPtr, nil
}

// UnmarshalJSON parses a JSON blob
func (p *paymentPerson) UnmarshalJSON(data []byte) error {

	dummy := make(map[string]interface{})
	err := json.Unmarshal(data, &dummy)
	if err != nil {
		return err
	}

	for name, record := range dummy {
		if name != "payer" {
			continue
		}
		p.Payer = records.NewARecord()
		err := readJsonWithRecord(p.Payer, record)
		if err != nil {
			return err
		}
	}

	typeOfReturn := ""
	if p.Payer != nil {
		if arec, ok := p.Payer.(*records.ARecord); !ok {
			return fmt.Errorf("unexpected Payer to be an ARecord, but got %T", p.Payer)
		} else {
			typeOfReturn = config.TypeOfReturns[arec.TypeOfReturn]
		}
	}

	for name, record := range dummy {
		buf, err := json.Marshal(record)
		if err != nil {
			return err
		}

		switch name {
		case "payees":
			var list []interface{}
			err = json.Unmarshal(buf, &list)
			if err != nil {
				return nil
			}

			p.Payees = make([]records.Record, 0)
			for _, data := range list {
				newRecord := records.NewBRecord(typeOfReturn)
				err := readJsonWithRecord(newRecord, data)
				if err != nil {
					return err
				}
				p.Payees = append(p.Payees, newRecord)
			}
		case "states":
			var list []interface{}
			err = json.Unmarshal(buf, &list)
			if err != nil {
				return nil
			}
			p.States = make([]records.Record, 0)
			for _, data := range list {
				newRecord := records.NewKRecord()
				err := readJsonWithRecord(newRecord, data)
				if err != nil {
					return err
				}
				p.States = append(p.States, newRecord)
			}
		case "end_payer":
			p.EndPayer = records.NewCRecord()
			err := readJsonWithRecord(p.EndPayer, record)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
