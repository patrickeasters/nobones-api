package lookup

import (
	"errors"
	"fmt"
	"net"
)

const BonesTXTRecord = "api.nobones.today"

var InvalidDNSRecord = errors.New("indeterminate results from DNS")

func BonesDay() (bool, error) {
	records, err := net.LookupTXT(BonesTXTRecord)
	if err != nil {
		return false, fmt.Errorf("failed to lookup TXT record: %w", err)
	}
	if len(records) > 1 || len(records) == 0 {
		return false, errors.New("indeterminate DNS records found")
	}

	switch records[0] {
	case "bones":
		return true, nil
	case "nobones":
		return false, nil
	default:
		return false, InvalidDNSRecord
	}

}
