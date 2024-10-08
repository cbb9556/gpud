// Package sxid provides the NVIDIA SXID error details.
package sxid

import "github.com/leptonai/gpud/components/common"

// Defines the SXID error type.
// ref. https://docs.nvidia.com/datacenter/tesla/pdf/fabric-manager-user-guide.pdf
type Detail struct {
	DocumentVersion string `json:"documentation_version"`

	SXID           int    `json:"sxid"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	PotentialFatal bool   `json:"potential_fatal"`
	AlwaysFatal    bool   `json:"always_fatal"`
	Impact         string `json:"impact"`
	Recovery       string `json:"recovery"`
	OtherImpact    string `json:"other_impact"`

	RequiredActions common.RequiredActions `json:"required_actions"`
}

// Returns the error if found.
// Otherwise, returns false.
func GetDetail(id int) (*Detail, bool) {
	e, ok := details[id]
	return &e, ok
}

// These are copied from:
// "D.4 Non-Fatal NVSwitch SXid Errors"
// "D.5 Fatal NVSwitch SXid Errors"
// "D.6 Always Fatal NVSwitch SXid Errors"
// "D.7 Other Notable NVSwitch SXid Errors"
// ref. https://docs.nvidia.com/datacenter/tesla/pdf/fabric-manager-user-guide.pdf
var details = map[int]Detail{
	// D.4 Non-Fatal NVSwitch SXid Errors
	11004: {
		DocumentVersion: "DU-09883-001_v1.3 (October 2023)",

		SXID:           11004,
		Name:           "Ingress invalid ACL",
		Description:    "This SXid error can happen only because of an incorrect FM partition configuration and is expected not to occur in the field.",
		PotentialFatal: false,
		AlwaysFatal:    false,
		Impact:         "Corresponding GPU NVLink traffic will be stalled, and the subsequent GPU access will hang. The GPU driver on the guest VM will abort CUDA jobs with Xid 45.",
		Recovery:       "Validate GPU/NVSwitch fabric partition routing information using the NVSwitch-audit tool. Restart the guest VM.",
		OtherImpact:    "If the error is observed on a Trunk port, partitions that are using NVSwitch trunk ports will be affected.",
		RequiredActions: common.RequiredActions{
			RebootSystem: true,
		},
	},
	11012: {
		DocumentVersion: "DU-09883-001_v1.3 (October 2023)",

		SXID:           11012,
		Name:           "Single bit ECC errors",
		Description:    "",
		PotentialFatal: false,
		AlwaysFatal:    false,
		Impact:         "No guest VM impact because the NVSwitch hardware will auto correct the ECC errors.",
		Recovery:       "Not Applicable.",
		OtherImpact:    "No Impact.",
	},
	11021: {
		DocumentVersion: "DU-09883-001_v1.3 (October 2023)",

		SXID:           11021,
		Name:           "Single bit ECC errors",
		Description:    "",
		PotentialFatal: false,
		AlwaysFatal:    false,
		Impact:         "No guest VM impact because the NVSwitch hardware will auto correct the ECC errors.",
		Recovery:       "Not Applicable.",
		OtherImpact:    "No Impact.",
	},
	11022: {
		DocumentVersion: "DU-09883-001_v1.3 (October 2023)",

		SXID:           11022,
		Name:           "Single bit ECC errors",
		Description:    "",
		PotentialFatal: false,
		AlwaysFatal:    false,
		Impact:         "No guest VM impact because the NVSwitch hardware will auto correct the ECC errors.",
		Recovery:       "Not Applicable.",
		OtherImpact:    "No Impact.",
	},
	11023: {
		DocumentVersion: "DU-09883-001_v1.3 (October 2023)",

		SXID:           11023,
		Name:           "Single bit ECC errors",
		Description:    "",
		PotentialFatal: false,
		AlwaysFatal:    false,
		Impact:         "No guest VM impact because the NVSwitch hardware will auto correct the ECC errors.",
		Recovery:       "Not Applicable.",
		OtherImpact:    "No Impact.",
	},
	12021: {
		DocumentVersion: "DU-09883-001_v1.3 (October 2023)",

		SXID:           12021,
		Name:           "Single bit ECC errors",
		Description:    "",
		PotentialFatal: false,
		AlwaysFatal:    false,
		Impact:         "No guest VM impact because the NVSwitch hardware will auto correct the ECC errors.",
		Recovery:       "Not Applicable.",
		OtherImpact:    "No Impact.",
	},
	12023: {
		DocumentVersion: "DU-09883-001_v1.3 (October 2023)",

		SXID:           12023,
		Name:           "Single bit ECC errors",
		Description:    "",
		PotentialFatal: false,
		AlwaysFatal:    false,
		Impact:         "No guest VM impact because the NVSwitch hardware will auto correct the ECC errors.",
		Recovery:       "Not Applicable.",
		OtherImpact:    "No Impact.",
	},
	15008: {
		DocumentVersion: "DU-09883-001_v1.3 (October 2023)",

		SXID:           15008,
		Name:           "Single bit ECC errors",
		Description:    "",
		PotentialFatal: false,
		AlwaysFatal:    false,
		Impact:         "No guest VM impact because the NVSwitch hardware will auto correct the ECC errors.",
		Recovery:       "Not Applicable.",
		OtherImpact:    "No Impact.",
	},
	15011: {
		DocumentVersion: "DU-09883-001_v1.3 (October 2023)",

		SXID:           15011,
		Name:           "Single bit ECC errors",
		Description:    "",
		PotentialFatal: false,
		AlwaysFatal:    false,
		Impact:         "No guest VM impact because the NVSwitch hardware will auto correct the ECC errors.",
		Recovery:       "Not Applicable.",
		OtherImpact:    "No Impact.",
	},
	19049: {
		DocumentVersion: "DU-09883-001_v1.3 (October 2023)",

		SXID:           19049,
		Name:           "Single bit ECC errors",
		Description:    "",
		PotentialFatal: false,
		AlwaysFatal:    false,
		Impact:         "No guest VM impact because the NVSwitch hardware will auto correct the ECC errors.",
		Recovery:       "Not Applicable.",
		OtherImpact:    "No Impact.",
	},
	19055: {
		DocumentVersion: "DU-09883-001_v1.3 (October 2023)",

		SXID:           19055,
		Name:           "Single bit ECC errors",
		Description:    "",
		PotentialFatal: false,
		AlwaysFatal:    false,
		Impact:         "No guest VM impact because the NVSwitch hardware will auto correct the ECC errors.",
		Recovery:       "Not Applicable.",
		OtherImpact:    "No Impact.",
	},
	19057: {
		DocumentVersion: "DU-09883-001_v1.3 (October 2023)",

		SXID:           19057,
		Name:           "Single bit ECC errors",
		Description:    "",
		PotentialFatal: false,
		AlwaysFatal:    false,
		Impact:         "No guest VM impact because the NVSwitch hardware will auto correct the ECC errors.",
		Recovery:       "Not Applicable.",
		OtherImpact:    "No Impact.",
	},
	19059: {
		DocumentVersion: "DU-09883-001_v1.3 (October 2023)",

		SXID:           19059,
		Name:           "Single bit ECC errors",
		Description:    "",
		PotentialFatal: false,
		AlwaysFatal:    false,
		Impact:         "No guest VM impact because the NVSwitch hardware will auto correct the ECC errors.",
		Recovery:       "Not Applicable.",
		OtherImpact:    "No Impact.",
	},
	19062: {
		DocumentVersion: "DU-09883-001_v1.3 (October 2023)",

		SXID:           19062,
		Name:           "Single bit ECC errors",
		Description:    "",
		PotentialFatal: false,
		AlwaysFatal:    false,
		Impact:         "No guest VM impact because the NVSwitch hardware will auto correct the ECC errors.",
		Recovery:       "Not Applicable.",
		OtherImpact:    "No Impact.",
	},
	19065: {
		DocumentVersion: "DU-09883-001_v1.3 (October 2023)",

		SXID:           19065,
		Name:           "Single bit ECC errors",
		Description:    "",
		PotentialFatal: false,
		AlwaysFatal:    false,
		Impact:         "No guest VM impact because the NVSwitch hardware will auto correct the ECC errors.",
		Recovery:       "Not Applicable.",
		OtherImpact:    "No Impact.",
	},
	19068: {
		DocumentVersion: "DU-09883-001_v1.3 (October 2023)",

		SXID:           19068,
		Name:           "Single bit ECC errors",
		Description:    "",
		PotentialFatal: false,
		AlwaysFatal:    false,
		Impact:         "No guest VM impact because the NVSwitch hardware will auto correct the ECC errors.",
		Recovery:       "Not Applicable.",
		OtherImpact:    "No Impact.",
	},
	19071: {
		DocumentVersion: "DU-09883-001_v1.3 (October 2023)",

		SXID:           19071,
		Name:           "Single bit ECC errors",
		Description:    "",
		PotentialFatal: false,
		AlwaysFatal:    false,
		Impact:         "No guest VM impact because the NVSwitch hardware will auto correct the ECC errors.",
		Recovery:       "Not Applicable.",
		OtherImpact:    "No Impact.",
	},
	24001: {
		DocumentVersion: "DU-09883-001_v1.3 (October 2023)",

		SXID:           24001,
		Name:           "Single bit ECC errors",
		Description:    "",
		PotentialFatal: false,
		AlwaysFatal:    false,
		Impact:         "No guest VM impact because the NVSwitch hardware will auto correct the ECC errors.",
		Recovery:       "Not Applicable.",
		OtherImpact:    "No Impact.",
	},
	24002: {
		DocumentVersion: "DU-09883-001_v1.3 (October 2023)",

		SXID:           24002,
		Name:           "Single bit ECC errors",
		Description:    "",
		PotentialFatal: false,
		AlwaysFatal:    false,
		Impact:         "No guest VM impact because the NVSwitch hardware will auto correct the ECC errors.",
		Recovery:       "Not Applicable.",
		OtherImpact:    "No Impact.",
	},
	24003: {
		DocumentVersion: "DU-09883-001_v1.3 (October 2023)",

		SXID:           24003,
		Name:           "Single bit ECC errors",
		Description:    "",
		PotentialFatal: false,
		AlwaysFatal:    false,
		Impact:         "No guest VM impact because the NVSwitch hardware will auto correct the ECC errors.",
		Recovery:       "Not Applicable.",
		OtherImpact:    "No Impact.",
	},
	20001: {
		DocumentVersion: "DU-09883-001_v1.3 (October 2023)",

		SXID:           20001,
		Name:           "TX Replay Error",
		Description:    "",
		PotentialFatal: false,
		AlwaysFatal:    false,
		Impact:         "NVLink packet needs to be retransmitted. This error might impact the NVLink throughput of the specified port.",
		Recovery:       "Not Applicable.",
		OtherImpact:    "If the error is observed on a Trunk port, the partitions that are using NVSwitch trunk ports might see throughput impact.",
	},
	12028: {
		DocumentVersion: "DU-09883-001_v1.3 (October 2023)",

		SXID:           12028,
		Name:           "egress nonposted PRIV error",
		Description:    "",
		PotentialFatal: false,
		AlwaysFatal:    false,
		Impact:         "Corresponding GPU NVLink traffic will be stalled, and subsequent GPU access will hang. The GPU driver on the guest VM will abort CUDA jobs with Xid 45.",
		Recovery:       "Restart Guest VM.",
		OtherImpact:    "If the error is observed on a Trunk port, the partitions that are using NVSwitch trunk ports will be affected.",
		RequiredActions: common.RequiredActions{
			RebootSystem: true,
		},
	},
	19084: {
		DocumentVersion: "DU-09883-001_v1.3 (October 2023)",

		SXID:           19084,
		Name:           "AN1 Heartbeat Timeout Error",
		Description:    "",
		PotentialFatal: false,
		AlwaysFatal:    false,
		Impact:         "This error is usually accompanied by a fatal SXid error that will affect the corresponding GPU NVLink traffic.",
		Recovery:       "Reset all GPUs and all NVSwitches (refer to section D.9).",
		OtherImpact:    "If the error is observed on a Trunk port, the partitions that are using NVSwitch trunk ports will be affected.",
	},
	20012: {
		DocumentVersion: "DU-09883-001_v1.3 (October 2023)",

		SXID:           20012,
		Name:           "Broken/inconsistent connection",
		Description:    "",
		PotentialFatal: false,
		AlwaysFatal:    false,
		Impact:         "This error could occur due to a broken/inconsistent connection or uncoordinated shutdown.",
		Recovery:       "If this issue was not due to an uncoordinated shutdown, check link mechanical connections.",
		OtherImpact:    "No impact if error is confined to a single GPU.",
	},
	22013: {
		DocumentVersion: "DU-09883-001_v1.3 (October 2023)",

		SXID:           22013,
		Name:           "Minion Link DLREQ interrupt",
		Description:    "",
		PotentialFatal: false,
		AlwaysFatal:    false,
		Impact:         "This SXid can be safely ignored.",
		Recovery:       "Not Applicable.",
		OtherImpact:    "No Impact.",
	},

	// D.5 Potential Fatal NVSwitch SXid Errors
	11001: {
		DocumentVersion: "DU-09883-001_v1.3 (October 2023)",

		SXID:            11001,
		Name:            "ingress invalid command",
		Description:     defaultPotentialFatalErr.Description,
		PotentialFatal:  defaultPotentialFatalErr.PotentialFatal,
		AlwaysFatal:     defaultPotentialFatalErr.AlwaysFatal,
		Impact:          defaultPotentialFatalErr.Impact,
		Recovery:        defaultPotentialFatalErr.Recovery,
		OtherImpact:     defaultPotentialFatalErr.OtherImpact,
		RequiredActions: defaultPotentialFatalErr.RequiredActions,
	},
	11009: {
		DocumentVersion: "DU-09883-001_v1.3 (October 2023)",

		SXID:            11009,
		Name:            "ingress invalid VCSet",
		Description:     defaultPotentialFatalErr.Description,
		PotentialFatal:  defaultPotentialFatalErr.PotentialFatal,
		AlwaysFatal:     defaultPotentialFatalErr.AlwaysFatal,
		Impact:          defaultPotentialFatalErr.Impact,
		Recovery:        defaultPotentialFatalErr.Recovery,
		OtherImpact:     defaultPotentialFatalErr.OtherImpact,
		RequiredActions: defaultPotentialFatalErr.RequiredActions,
	},
	11013: {
		DocumentVersion: "DU-09883-001_v1.3 (October 2023)",

		SXID:            11013,
		Name:            "ingress header DBE",
		Description:     defaultPotentialFatalErr.Description,
		PotentialFatal:  defaultPotentialFatalErr.PotentialFatal,
		AlwaysFatal:     defaultPotentialFatalErr.AlwaysFatal,
		Impact:          defaultPotentialFatalErr.Impact,
		Recovery:        defaultPotentialFatalErr.Recovery,
		OtherImpact:     defaultPotentialFatalErr.OtherImpact,
		RequiredActions: defaultPotentialFatalErr.RequiredActions,
	},
	11018: {
		DocumentVersion: "DU-09883-001_v1.3 (October 2023)",

		SXID:            11018,
		Name:            "ingress RID DBE",
		Description:     defaultPotentialFatalErr.Description,
		PotentialFatal:  defaultPotentialFatalErr.PotentialFatal,
		AlwaysFatal:     defaultPotentialFatalErr.AlwaysFatal,
		Impact:          defaultPotentialFatalErr.Impact,
		Recovery:        defaultPotentialFatalErr.Recovery,
		OtherImpact:     defaultPotentialFatalErr.OtherImpact,
		RequiredActions: defaultPotentialFatalErr.RequiredActions,
	},
	11019: {
		DocumentVersion: "DU-09883-001_v1.3 (October 2023)",

		SXID:            11019,
		Name:            "ingress RLAN DBE",
		Description:     defaultPotentialFatalErr.Description,
		PotentialFatal:  defaultPotentialFatalErr.PotentialFatal,
		AlwaysFatal:     defaultPotentialFatalErr.AlwaysFatal,
		Impact:          defaultPotentialFatalErr.Impact,
		Recovery:        defaultPotentialFatalErr.Recovery,
		OtherImpact:     defaultPotentialFatalErr.OtherImpact,
		RequiredActions: defaultPotentialFatalErr.RequiredActions,
	},
	11020: {
		DocumentVersion: "DU-09883-001_v1.3 (October 2023)",

		SXID:            11020,
		Name:            "ingress control parity",
		Description:     defaultPotentialFatalErr.Description,
		PotentialFatal:  defaultPotentialFatalErr.PotentialFatal,
		AlwaysFatal:     defaultPotentialFatalErr.AlwaysFatal,
		Impact:          defaultPotentialFatalErr.Impact,
		Recovery:        defaultPotentialFatalErr.Recovery,
		OtherImpact:     defaultPotentialFatalErr.OtherImpact,
		RequiredActions: defaultPotentialFatalErr.RequiredActions,
	},
	12001: {
		DocumentVersion: "DU-09883-001_v1.3 (October 2023)",

		SXID:            12001,
		Name:            "egress crossbar overflow",
		Description:     defaultPotentialFatalErr.Description,
		PotentialFatal:  defaultPotentialFatalErr.PotentialFatal,
		AlwaysFatal:     defaultPotentialFatalErr.AlwaysFatal,
		Impact:          defaultPotentialFatalErr.Impact,
		Recovery:        defaultPotentialFatalErr.Recovery,
		OtherImpact:     defaultPotentialFatalErr.OtherImpact,
		RequiredActions: defaultPotentialFatalErr.RequiredActions,
	},
	12002: {
		DocumentVersion: "DU-09883-001_v1.3 (October 2023)",

		SXID:            12002,
		Name:            "egress packet route",
		Description:     defaultPotentialFatalErr.Description,
		PotentialFatal:  defaultPotentialFatalErr.PotentialFatal,
		AlwaysFatal:     defaultPotentialFatalErr.AlwaysFatal,
		Impact:          defaultPotentialFatalErr.Impact,
		Recovery:        defaultPotentialFatalErr.Recovery,
		OtherImpact:     defaultPotentialFatalErr.OtherImpact,
		RequiredActions: defaultPotentialFatalErr.RequiredActions,
	},
	12022: {
		DocumentVersion: "DU-09883-001_v1.3 (October 2023)",

		SXID:            12022,
		Name:            "egress input ECC DBE error",
		Description:     defaultPotentialFatalErr.Description,
		PotentialFatal:  defaultPotentialFatalErr.PotentialFatal,
		AlwaysFatal:     defaultPotentialFatalErr.AlwaysFatal,
		Impact:          defaultPotentialFatalErr.Impact,
		Recovery:        defaultPotentialFatalErr.Recovery,
		OtherImpact:     defaultPotentialFatalErr.OtherImpact,
		RequiredActions: defaultPotentialFatalErr.RequiredActions,
	},
	12024: {
		DocumentVersion: "DU-09883-001_v1.3 (October 2023)",

		SXID:            12024,
		Name:            "egress output ECC DBE error",
		Description:     defaultPotentialFatalErr.Description,
		PotentialFatal:  defaultPotentialFatalErr.PotentialFatal,
		AlwaysFatal:     defaultPotentialFatalErr.AlwaysFatal,
		Impact:          defaultPotentialFatalErr.Impact,
		Recovery:        defaultPotentialFatalErr.Recovery,
		OtherImpact:     defaultPotentialFatalErr.OtherImpact,
		RequiredActions: defaultPotentialFatalErr.RequiredActions,
	},
	12025: {
		DocumentVersion: "DU-09883-001_v1.3 (October 2023)",

		SXID:            12025,
		Name:            "egress credit overflow",
		Description:     defaultPotentialFatalErr.Description,
		PotentialFatal:  defaultPotentialFatalErr.PotentialFatal,
		AlwaysFatal:     defaultPotentialFatalErr.AlwaysFatal,
		Impact:          defaultPotentialFatalErr.Impact,
		Recovery:        defaultPotentialFatalErr.Recovery,
		OtherImpact:     defaultPotentialFatalErr.OtherImpact,
		RequiredActions: defaultPotentialFatalErr.RequiredActions,
	},
	12026: {
		DocumentVersion: "DU-09883-001_v1.3 (October 2023)",

		SXID:            12026,
		Name:            "egress destination request ID error",
		Description:     defaultPotentialFatalErr.Description,
		PotentialFatal:  defaultPotentialFatalErr.PotentialFatal,
		AlwaysFatal:     defaultPotentialFatalErr.AlwaysFatal,
		Impact:          defaultPotentialFatalErr.Impact,
		Recovery:        defaultPotentialFatalErr.Recovery,
		OtherImpact:     defaultPotentialFatalErr.OtherImpact,
		RequiredActions: defaultPotentialFatalErr.RequiredActions,
	},
	12027: {
		DocumentVersion: "DU-09883-001_v1.3 (October 2023)",

		SXID:            12027,
		Name:            "egress destination response ID error",
		Description:     defaultPotentialFatalErr.Description,
		PotentialFatal:  defaultPotentialFatalErr.PotentialFatal,
		AlwaysFatal:     defaultPotentialFatalErr.AlwaysFatal,
		Impact:          defaultPotentialFatalErr.Impact,
		Recovery:        defaultPotentialFatalErr.Recovery,
		OtherImpact:     defaultPotentialFatalErr.OtherImpact,
		RequiredActions: defaultPotentialFatalErr.RequiredActions,
	},
	12030: {
		DocumentVersion: "DU-09883-001_v1.3 (October 2023)",

		SXID:            12030,
		Name:            "egress control parity error",
		Description:     defaultPotentialFatalErr.Description,
		PotentialFatal:  defaultPotentialFatalErr.PotentialFatal,
		AlwaysFatal:     defaultPotentialFatalErr.AlwaysFatal,
		Impact:          defaultPotentialFatalErr.Impact,
		Recovery:        defaultPotentialFatalErr.Recovery,
		OtherImpact:     defaultPotentialFatalErr.OtherImpact,
		RequiredActions: defaultPotentialFatalErr.RequiredActions,
	},
	12031: {
		DocumentVersion: "DU-09883-001_v1.3 (October 2023)",

		SXID:            12031,
		Name:            "egress credit parity error",
		Description:     defaultPotentialFatalErr.Description,
		PotentialFatal:  defaultPotentialFatalErr.PotentialFatal,
		AlwaysFatal:     defaultPotentialFatalErr.AlwaysFatal,
		Impact:          defaultPotentialFatalErr.Impact,
		Recovery:        defaultPotentialFatalErr.Recovery,
		OtherImpact:     defaultPotentialFatalErr.OtherImpact,
		RequiredActions: defaultPotentialFatalErr.RequiredActions,
	},
	12032: {
		DocumentVersion: "DU-09883-001_v1.3 (October 2023)",

		SXID:            12032,
		Name:            "egress flit type mismatch",
		Description:     defaultPotentialFatalErr.Description,
		PotentialFatal:  defaultPotentialFatalErr.PotentialFatal,
		AlwaysFatal:     defaultPotentialFatalErr.AlwaysFatal,
		Impact:          defaultPotentialFatalErr.Impact,
		Recovery:        defaultPotentialFatalErr.Recovery,
		OtherImpact:     defaultPotentialFatalErr.OtherImpact,
		RequiredActions: defaultPotentialFatalErr.RequiredActions,
	},
	14017: {
		DocumentVersion: "DU-09883-001_v1.3 (October 2023)",

		SXID:            14017,
		Name:            "TS ATO timeout",
		Description:     defaultPotentialFatalErr.Description,
		PotentialFatal:  defaultPotentialFatalErr.PotentialFatal,
		AlwaysFatal:     defaultPotentialFatalErr.AlwaysFatal,
		Impact:          defaultPotentialFatalErr.Impact,
		Recovery:        defaultPotentialFatalErr.Recovery,
		OtherImpact:     defaultPotentialFatalErr.OtherImpact,
		RequiredActions: defaultPotentialFatalErr.RequiredActions,
	},
	15001: {
		DocumentVersion: "DU-09883-001_v1.3 (October 2023)",

		SXID:            15001,
		Name:            "route buffer over/underflow",
		Description:     defaultPotentialFatalErr.Description,
		PotentialFatal:  defaultPotentialFatalErr.PotentialFatal,
		AlwaysFatal:     defaultPotentialFatalErr.AlwaysFatal,
		Impact:          defaultPotentialFatalErr.Impact,
		Recovery:        defaultPotentialFatalErr.Recovery,
		OtherImpact:     defaultPotentialFatalErr.OtherImpact,
		RequiredActions: defaultPotentialFatalErr.RequiredActions,
	},
	15006: {
		DocumentVersion: "DU-09883-001_v1.3 (October 2023)",

		SXID:            15006,
		Name:            "route transdone over/underflow",
		Description:     defaultPotentialFatalErr.Description,
		PotentialFatal:  defaultPotentialFatalErr.PotentialFatal,
		AlwaysFatal:     defaultPotentialFatalErr.AlwaysFatal,
		Impact:          defaultPotentialFatalErr.Impact,
		Recovery:        defaultPotentialFatalErr.Recovery,
		OtherImpact:     defaultPotentialFatalErr.OtherImpact,
		RequiredActions: defaultPotentialFatalErr.RequiredActions,
	},
	15009: {
		DocumentVersion: "DU-09883-001_v1.3 (October 2023)",

		SXID:            15009,
		Name:            "route GLT DBE",
		Description:     defaultPotentialFatalErr.Description,
		PotentialFatal:  defaultPotentialFatalErr.PotentialFatal,
		AlwaysFatal:     defaultPotentialFatalErr.AlwaysFatal,
		Impact:          defaultPotentialFatalErr.Impact,
		Recovery:        defaultPotentialFatalErr.Recovery,
		OtherImpact:     defaultPotentialFatalErr.OtherImpact,
		RequiredActions: defaultPotentialFatalErr.RequiredActions,
	},
	15010: {
		DocumentVersion: "DU-09883-001_v1.3 (October 2023)",

		SXID:            15010,
		Name:            "route parity",
		Description:     defaultPotentialFatalErr.Description,
		PotentialFatal:  defaultPotentialFatalErr.PotentialFatal,
		AlwaysFatal:     defaultPotentialFatalErr.AlwaysFatal,
		Impact:          defaultPotentialFatalErr.Impact,
		Recovery:        defaultPotentialFatalErr.Recovery,
		OtherImpact:     defaultPotentialFatalErr.OtherImpact,
		RequiredActions: defaultPotentialFatalErr.RequiredActions,
	},
	15012: {
		DocumentVersion: "DU-09883-001_v1.3 (October 2023)",

		SXID:            15012,
		Name:            "route incoming DBE",
		Description:     defaultPotentialFatalErr.Description,
		PotentialFatal:  defaultPotentialFatalErr.PotentialFatal,
		AlwaysFatal:     defaultPotentialFatalErr.AlwaysFatal,
		Impact:          defaultPotentialFatalErr.Impact,
		Recovery:        defaultPotentialFatalErr.Recovery,
		OtherImpact:     defaultPotentialFatalErr.OtherImpact,
		RequiredActions: defaultPotentialFatalErr.RequiredActions,
	},
	15013: {
		DocumentVersion: "DU-09883-001_v1.3 (October 2023)",

		SXID:            15013,
		Name:            "route credit parity",
		Description:     defaultPotentialFatalErr.Description,
		PotentialFatal:  defaultPotentialFatalErr.PotentialFatal,
		AlwaysFatal:     defaultPotentialFatalErr.AlwaysFatal,
		Impact:          defaultPotentialFatalErr.Impact,
		Recovery:        defaultPotentialFatalErr.Recovery,
		OtherImpact:     defaultPotentialFatalErr.OtherImpact,
		RequiredActions: defaultPotentialFatalErr.RequiredActions,
	},
	19047: {
		DocumentVersion: "DU-09883-001_v1.3 (October 2023)",

		SXID:            19047,
		Name:            "NCISOC HDR ECC DBE Error",
		Description:     defaultPotentialFatalErr.Description,
		PotentialFatal:  defaultPotentialFatalErr.PotentialFatal,
		AlwaysFatal:     defaultPotentialFatalErr.AlwaysFatal,
		Impact:          defaultPotentialFatalErr.Impact,
		Recovery:        defaultPotentialFatalErr.Recovery,
		OtherImpact:     defaultPotentialFatalErr.OtherImpact,
		RequiredActions: defaultPotentialFatalErr.RequiredActions,
	},
	19048: {
		DocumentVersion: "DU-09883-001_v1.3 (October 2023)",

		SXID:            19048,
		Name:            "NCISOC DAT ECC DBE Error",
		Description:     defaultPotentialFatalErr.Description,
		PotentialFatal:  defaultPotentialFatalErr.PotentialFatal,
		AlwaysFatal:     defaultPotentialFatalErr.AlwaysFatal,
		Impact:          defaultPotentialFatalErr.Impact,
		Recovery:        defaultPotentialFatalErr.Recovery,
		OtherImpact:     defaultPotentialFatalErr.OtherImpact,
		RequiredActions: defaultPotentialFatalErr.RequiredActions,
	},
	19054: {
		DocumentVersion: "DU-09883-001_v1.3 (October 2023)",

		SXID:            19054,
		Name:            "HDR RAM ECC DBE Error",
		Description:     defaultPotentialFatalErr.Description,
		PotentialFatal:  defaultPotentialFatalErr.PotentialFatal,
		AlwaysFatal:     defaultPotentialFatalErr.AlwaysFatal,
		Impact:          defaultPotentialFatalErr.Impact,
		Recovery:        defaultPotentialFatalErr.Recovery,
		OtherImpact:     defaultPotentialFatalErr.OtherImpact,
		RequiredActions: defaultPotentialFatalErr.RequiredActions,
	},
	19056: {
		DocumentVersion: "DU-09883-001_v1.3 (October 2023)",

		SXID:            19056,
		Name:            "DAT0 RAM ECC DBE Error",
		Description:     defaultPotentialFatalErr.Description,
		PotentialFatal:  defaultPotentialFatalErr.PotentialFatal,
		AlwaysFatal:     defaultPotentialFatalErr.AlwaysFatal,
		Impact:          defaultPotentialFatalErr.Impact,
		Recovery:        defaultPotentialFatalErr.Recovery,
		OtherImpact:     defaultPotentialFatalErr.OtherImpact,
		RequiredActions: defaultPotentialFatalErr.RequiredActions,
	},
	19058: {
		DocumentVersion: "DU-09883-001_v1.3 (October 2023)",

		SXID:            19058,
		Name:            "DAT1 RAM ECC DBE Error",
		Description:     defaultPotentialFatalErr.Description,
		PotentialFatal:  defaultPotentialFatalErr.PotentialFatal,
		AlwaysFatal:     defaultPotentialFatalErr.AlwaysFatal,
		Impact:          defaultPotentialFatalErr.Impact,
		Recovery:        defaultPotentialFatalErr.Recovery,
		OtherImpact:     defaultPotentialFatalErr.OtherImpact,
		RequiredActions: defaultPotentialFatalErr.RequiredActions,
	},
	19060: {
		DocumentVersion: "DU-09883-001_v1.3 (October 2023)",

		SXID:            19060,
		Name:            "CREQ RAM HDR ECC DBE Error",
		Description:     defaultPotentialFatalErr.Description,
		PotentialFatal:  defaultPotentialFatalErr.PotentialFatal,
		AlwaysFatal:     defaultPotentialFatalErr.AlwaysFatal,
		Impact:          defaultPotentialFatalErr.Impact,
		Recovery:        defaultPotentialFatalErr.Recovery,
		OtherImpact:     defaultPotentialFatalErr.OtherImpact,
		RequiredActions: defaultPotentialFatalErr.RequiredActions,
	},
	19061: {
		DocumentVersion: "DU-09883-001_v1.3 (October 2023)",

		SXID:            19061,
		Name:            "CREQ RAM DAT ECC DBE Error",
		Description:     defaultPotentialFatalErr.Description,
		PotentialFatal:  defaultPotentialFatalErr.PotentialFatal,
		AlwaysFatal:     defaultPotentialFatalErr.AlwaysFatal,
		Impact:          defaultPotentialFatalErr.Impact,
		Recovery:        defaultPotentialFatalErr.Recovery,
		OtherImpact:     defaultPotentialFatalErr.OtherImpact,
		RequiredActions: defaultPotentialFatalErr.RequiredActions,
	},
	19063: {
		DocumentVersion: "DU-09883-001_v1.3 (October 2023)",

		SXID:            19063,
		Name:            "Response RAM HDR ECC DBE Error",
		Description:     defaultPotentialFatalErr.Description,
		PotentialFatal:  defaultPotentialFatalErr.PotentialFatal,
		AlwaysFatal:     defaultPotentialFatalErr.AlwaysFatal,
		Impact:          defaultPotentialFatalErr.Impact,
		Recovery:        defaultPotentialFatalErr.Recovery,
		OtherImpact:     defaultPotentialFatalErr.OtherImpact,
		RequiredActions: defaultPotentialFatalErr.RequiredActions,
	},
	19064: {
		DocumentVersion: "DU-09883-001_v1.3 (October 2023)",

		SXID:            19064,
		Name:            "Response RAM DAT ECC DBE Error",
		Description:     defaultPotentialFatalErr.Description,
		PotentialFatal:  defaultPotentialFatalErr.PotentialFatal,
		AlwaysFatal:     defaultPotentialFatalErr.AlwaysFatal,
		Impact:          defaultPotentialFatalErr.Impact,
		Recovery:        defaultPotentialFatalErr.Recovery,
		OtherImpact:     defaultPotentialFatalErr.OtherImpact,
		RequiredActions: defaultPotentialFatalErr.RequiredActions,
	},
	19066: {
		DocumentVersion: "DU-09883-001_v1.3 (October 2023)",

		SXID:            19066,
		Name:            "COM RAM HDR ECC DBE Error",
		Description:     defaultPotentialFatalErr.Description,
		PotentialFatal:  defaultPotentialFatalErr.PotentialFatal,
		AlwaysFatal:     defaultPotentialFatalErr.AlwaysFatal,
		Impact:          defaultPotentialFatalErr.Impact,
		Recovery:        defaultPotentialFatalErr.Recovery,
		OtherImpact:     defaultPotentialFatalErr.OtherImpact,
		RequiredActions: defaultPotentialFatalErr.RequiredActions,
	},
	19067: {
		DocumentVersion: "DU-09883-001_v1.3 (October 2023)",

		SXID:            19067,
		Name:            "COM RAM DAT ECC DBE Error",
		Description:     defaultPotentialFatalErr.Description,
		PotentialFatal:  defaultPotentialFatalErr.PotentialFatal,
		AlwaysFatal:     defaultPotentialFatalErr.AlwaysFatal,
		Impact:          defaultPotentialFatalErr.Impact,
		Recovery:        defaultPotentialFatalErr.Recovery,
		OtherImpact:     defaultPotentialFatalErr.OtherImpact,
		RequiredActions: defaultPotentialFatalErr.RequiredActions,
	},
	19069: {
		DocumentVersion: "DU-09883-001_v1.3 (October 2023)",

		SXID:            19069,
		Name:            "RSP1 RAM HDR ECC DBE Error",
		Description:     defaultPotentialFatalErr.Description,
		PotentialFatal:  defaultPotentialFatalErr.PotentialFatal,
		AlwaysFatal:     defaultPotentialFatalErr.AlwaysFatal,
		Impact:          defaultPotentialFatalErr.Impact,
		Recovery:        defaultPotentialFatalErr.Recovery,
		OtherImpact:     defaultPotentialFatalErr.OtherImpact,
		RequiredActions: defaultPotentialFatalErr.RequiredActions,
	},
	19070: {
		DocumentVersion: "DU-09883-001_v1.3 (October 2023)",

		SXID:            19070,
		Name:            "RSP1 RAM DAT ECC DBE Error",
		Description:     defaultPotentialFatalErr.Description,
		PotentialFatal:  defaultPotentialFatalErr.PotentialFatal,
		AlwaysFatal:     defaultPotentialFatalErr.AlwaysFatal,
		Impact:          defaultPotentialFatalErr.Impact,
		Recovery:        defaultPotentialFatalErr.Recovery,
		OtherImpact:     defaultPotentialFatalErr.OtherImpact,
		RequiredActions: defaultPotentialFatalErr.RequiredActions,
	},
	20034: {
		DocumentVersion: "DU-09883-001_v1.3 (October 2023)",

		SXID:           20034,
		Name:           "LTSSM Fault Up",
		Description:    defaultPotentialFatalErr.Description,
		PotentialFatal: defaultPotentialFatalErr.PotentialFatal,
		AlwaysFatal:    defaultPotentialFatalErr.AlwaysFatal,
		Impact: defaultPotentialFatalErr.Impact + `
Guest VM Impact: This SXid is triggered whenever the associated link has gone
down from active. This interrupt is usually associated with other NVLink errors.
`,
		Recovery: defaultPotentialFatalErr.Recovery + `
Guest VM Recovery: In case of A100, restart the VM. In case of H100, reset the
GPU (refer to section D.9). If issue persists, report GPU issues.
`,
		OtherImpact: defaultPotentialFatalErr.OtherImpact + `
Other Guest VM Impact: No impact if error is confined to a single GPU.
`,
		RequiredActions: common.RequiredActions{
			RebootSystem: true,
		},
	},
	22012: { // in both D.4 and D.5, treat it as potential fatal
		DocumentVersion: "DU-09883-001_v1.3 (October 2023)",

		SXID:           22012,
		Name:           "Minion Link NA interrupt",
		Description:    "",
		PotentialFatal: true,
		AlwaysFatal:    false,
		Impact:         "This error could occur due to a broken/inconsistent connection or uncoordinated shutdown.",
		Recovery:       "If this issue was not due to an uncoordinated shutdown, check link mechanical connections.",
		OtherImpact:    "No impact if error is confined to a single GPU.",
	},
	24004: {
		DocumentVersion: "DU-09883-001_v1.3 (October 2023)",

		SXID:            24004,
		Name:            "sourcetrack TCEN0 crubmstore DBE",
		Description:     defaultPotentialFatalErr.Description,
		PotentialFatal:  defaultPotentialFatalErr.PotentialFatal,
		AlwaysFatal:     defaultPotentialFatalErr.AlwaysFatal,
		Impact:          defaultPotentialFatalErr.Impact,
		Recovery:        defaultPotentialFatalErr.Recovery,
		OtherImpact:     defaultPotentialFatalErr.OtherImpact,
		RequiredActions: defaultPotentialFatalErr.RequiredActions,
	},
	24005: {
		DocumentVersion: "DU-09883-001_v1.3 (October 2023)",

		SXID:            24005,
		Name:            "sourcetrack TCEN0 TD crubmstore DBE",
		Description:     defaultPotentialFatalErr.Description,
		PotentialFatal:  defaultPotentialFatalErr.PotentialFatal,
		AlwaysFatal:     defaultPotentialFatalErr.AlwaysFatal,
		Impact:          defaultPotentialFatalErr.Impact,
		Recovery:        defaultPotentialFatalErr.Recovery,
		OtherImpact:     defaultPotentialFatalErr.OtherImpact,
		RequiredActions: defaultPotentialFatalErr.RequiredActions,
	},
	24006: {
		DocumentVersion: "DU-09883-001_v1.3 (October 2023)",

		SXID:            24006,
		Name:            "sourcetrack TCEN1 crubmstore DBE",
		Description:     defaultPotentialFatalErr.Description,
		PotentialFatal:  defaultPotentialFatalErr.PotentialFatal,
		AlwaysFatal:     defaultPotentialFatalErr.AlwaysFatal,
		Impact:          defaultPotentialFatalErr.Impact,
		Recovery:        defaultPotentialFatalErr.Recovery,
		OtherImpact:     defaultPotentialFatalErr.OtherImpact,
		RequiredActions: defaultPotentialFatalErr.RequiredActions,
	},
	24007: {
		DocumentVersion: "DU-09883-001_v1.3 (October 2023)",

		SXID:            24007,
		Name:            "sourcetrack timeout error",
		Description:     defaultPotentialFatalErr.Description,
		PotentialFatal:  defaultPotentialFatalErr.PotentialFatal,
		AlwaysFatal:     defaultPotentialFatalErr.AlwaysFatal,
		Impact:          defaultPotentialFatalErr.Impact,
		Recovery:        defaultPotentialFatalErr.Recovery,
		OtherImpact:     defaultPotentialFatalErr.OtherImpact,
		RequiredActions: defaultPotentialFatalErr.RequiredActions,
	},

	// D.6 Always Fatal NVSwitch SXid Errors
	12020: {
		DocumentVersion: "DU-09883-001_v1.3 (October 2023)",

		SXID:            12020,
		Name:            "egress sequence ID error",
		Description:     defaultAlwaysFatalErr.Description,
		PotentialFatal:  defaultAlwaysFatalErr.PotentialFatal,
		AlwaysFatal:     defaultAlwaysFatalErr.AlwaysFatal,
		Impact:          defaultAlwaysFatalErr.Impact,
		Recovery:        defaultAlwaysFatalErr.Recovery,
		OtherImpact:     defaultAlwaysFatalErr.OtherImpact,
		RequiredActions: defaultAlwaysFatalErr.RequiredActions,
	},
	22003: {
		DocumentVersion: "DU-09883-001_v1.3 (October 2023)",

		SXID:            22003,
		Name:            "Minion Halt",
		Description:     defaultAlwaysFatalErr.Description,
		PotentialFatal:  defaultAlwaysFatalErr.PotentialFatal,
		AlwaysFatal:     defaultAlwaysFatalErr.AlwaysFatal,
		Impact:          defaultAlwaysFatalErr.Impact,
		Recovery:        defaultAlwaysFatalErr.Recovery,
		OtherImpact:     defaultAlwaysFatalErr.OtherImpact,
		RequiredActions: defaultAlwaysFatalErr.RequiredActions,
	},
	22011: {
		DocumentVersion: "DU-09883-001_v1.3 (October 2023)",

		SXID:            22011,
		Name:            "Minion exterror",
		Description:     defaultAlwaysFatalErr.Description,
		PotentialFatal:  defaultAlwaysFatalErr.PotentialFatal,
		AlwaysFatal:     defaultAlwaysFatalErr.AlwaysFatal,
		Impact:          defaultAlwaysFatalErr.Impact,
		Recovery:        defaultAlwaysFatalErr.Recovery,
		OtherImpact:     defaultAlwaysFatalErr.OtherImpact,
		RequiredActions: defaultAlwaysFatalErr.RequiredActions,
	},
	23001: {
		DocumentVersion: "DU-09883-001_v1.3 (October 2023)",

		SXID:            23001,
		Name:            "ingress SRC-VC buffer overflow",
		Description:     defaultAlwaysFatalErr.Description,
		PotentialFatal:  defaultAlwaysFatalErr.PotentialFatal,
		AlwaysFatal:     defaultAlwaysFatalErr.AlwaysFatal,
		Impact:          defaultAlwaysFatalErr.Impact,
		Recovery:        defaultAlwaysFatalErr.Recovery,
		OtherImpact:     defaultAlwaysFatalErr.OtherImpact,
		RequiredActions: defaultAlwaysFatalErr.RequiredActions,
	},
	23002: {
		DocumentVersion: "DU-09883-001_v1.3 (October 2023)",

		SXID:            23002,
		Name:            "ingress SRC-VC buffer underflow",
		Description:     defaultAlwaysFatalErr.Description,
		PotentialFatal:  defaultAlwaysFatalErr.PotentialFatal,
		AlwaysFatal:     defaultAlwaysFatalErr.AlwaysFatal,
		Impact:          defaultAlwaysFatalErr.Impact,
		Recovery:        defaultAlwaysFatalErr.Recovery,
		OtherImpact:     defaultAlwaysFatalErr.OtherImpact,
		RequiredActions: defaultAlwaysFatalErr.RequiredActions,
	},
	23003: {
		DocumentVersion: "DU-09883-001_v1.3 (October 2023)",

		SXID:            23003,
		Name:            "egress DST-VC credit overflow",
		Description:     defaultAlwaysFatalErr.Description,
		PotentialFatal:  defaultAlwaysFatalErr.PotentialFatal,
		AlwaysFatal:     defaultAlwaysFatalErr.AlwaysFatal,
		Impact:          defaultAlwaysFatalErr.Impact,
		Recovery:        defaultAlwaysFatalErr.Recovery,
		OtherImpact:     defaultAlwaysFatalErr.OtherImpact,
		RequiredActions: defaultAlwaysFatalErr.RequiredActions,
	},
	23004: {
		DocumentVersion: "DU-09883-001_v1.3 (October 2023)",

		SXID:            23004,
		Name:            "egress DST-VC credit underflow",
		Description:     defaultAlwaysFatalErr.Description,
		PotentialFatal:  defaultAlwaysFatalErr.PotentialFatal,
		AlwaysFatal:     defaultAlwaysFatalErr.AlwaysFatal,
		Impact:          defaultAlwaysFatalErr.Impact,
		Recovery:        defaultAlwaysFatalErr.Recovery,
		OtherImpact:     defaultAlwaysFatalErr.OtherImpact,
		RequiredActions: defaultAlwaysFatalErr.RequiredActions,
	},
	23005: {
		DocumentVersion: "DU-09883-001_v1.3 (October 2023)",

		SXID:            23005,
		Name:            "ingress packet burst error",
		Description:     defaultAlwaysFatalErr.Description,
		PotentialFatal:  defaultAlwaysFatalErr.PotentialFatal,
		AlwaysFatal:     defaultAlwaysFatalErr.AlwaysFatal,
		Impact:          defaultAlwaysFatalErr.Impact,
		Recovery:        defaultAlwaysFatalErr.Recovery,
		OtherImpact:     defaultAlwaysFatalErr.OtherImpact,
		RequiredActions: defaultAlwaysFatalErr.RequiredActions,
	},
	23006: {
		DocumentVersion: "DU-09883-001_v1.3 (October 2023)",

		SXID:            23006,
		Name:            "ingress packet sticky error",
		Description:     defaultAlwaysFatalErr.Description,
		PotentialFatal:  defaultAlwaysFatalErr.PotentialFatal,
		AlwaysFatal:     defaultAlwaysFatalErr.AlwaysFatal,
		Impact:          defaultAlwaysFatalErr.Impact,
		Recovery:        defaultAlwaysFatalErr.Recovery,
		OtherImpact:     defaultAlwaysFatalErr.OtherImpact,
		RequiredActions: defaultAlwaysFatalErr.RequiredActions,
	},
	23007: {
		DocumentVersion: "DU-09883-001_v1.3 (October 2023)",

		SXID:            23007,
		Name:            "possible bubbles at ingress",
		Description:     defaultAlwaysFatalErr.Description,
		PotentialFatal:  defaultAlwaysFatalErr.PotentialFatal,
		AlwaysFatal:     defaultAlwaysFatalErr.AlwaysFatal,
		Impact:          defaultAlwaysFatalErr.Impact,
		Recovery:        defaultAlwaysFatalErr.Recovery,
		OtherImpact:     defaultAlwaysFatalErr.OtherImpact,
		RequiredActions: defaultAlwaysFatalErr.RequiredActions,
	},
	23008: {
		DocumentVersion: "DU-09883-001_v1.3 (October 2023)",

		SXID:            23008,
		Name:            "ingress packet invalid dst error",
		Description:     defaultAlwaysFatalErr.Description,
		PotentialFatal:  defaultAlwaysFatalErr.PotentialFatal,
		AlwaysFatal:     defaultAlwaysFatalErr.AlwaysFatal,
		Impact:          defaultAlwaysFatalErr.Impact,
		Recovery:        defaultAlwaysFatalErr.Recovery,
		OtherImpact:     defaultAlwaysFatalErr.OtherImpact,
		RequiredActions: defaultAlwaysFatalErr.RequiredActions,
	},
	23009: {
		DocumentVersion: "DU-09883-001_v1.3 (October 2023)",

		SXID:            23009,
		Name:            "ingress packet parity error",
		Description:     defaultAlwaysFatalErr.Description,
		PotentialFatal:  defaultAlwaysFatalErr.PotentialFatal,
		AlwaysFatal:     defaultAlwaysFatalErr.AlwaysFatal,
		Impact:          defaultAlwaysFatalErr.Impact,
		Recovery:        defaultAlwaysFatalErr.Recovery,
		OtherImpact:     defaultAlwaysFatalErr.OtherImpact,
		RequiredActions: defaultAlwaysFatalErr.RequiredActions,
	},
	23010: {
		DocumentVersion: "DU-09883-001_v1.3 (October 2023)",

		SXID:            23010,
		Name:            "ingress SRC-VC buffer overflow",
		Description:     defaultAlwaysFatalErr.Description,
		PotentialFatal:  defaultAlwaysFatalErr.PotentialFatal,
		AlwaysFatal:     defaultAlwaysFatalErr.AlwaysFatal,
		Impact:          defaultAlwaysFatalErr.Impact,
		Recovery:        defaultAlwaysFatalErr.Recovery,
		OtherImpact:     defaultAlwaysFatalErr.OtherImpact,
		RequiredActions: defaultAlwaysFatalErr.RequiredActions,
	},
	23011: {
		DocumentVersion: "DU-09883-001_v1.3 (October 2023)",

		SXID:            23011,
		Name:            "ingress SRC-VC buffer underflow",
		Description:     defaultAlwaysFatalErr.Description,
		PotentialFatal:  defaultAlwaysFatalErr.PotentialFatal,
		AlwaysFatal:     defaultAlwaysFatalErr.AlwaysFatal,
		Impact:          defaultAlwaysFatalErr.Impact,
		Recovery:        defaultAlwaysFatalErr.Recovery,
		OtherImpact:     defaultAlwaysFatalErr.OtherImpact,
		RequiredActions: defaultAlwaysFatalErr.RequiredActions,
	},
	23012: {
		DocumentVersion: "DU-09883-001_v1.3 (October 2023)",

		SXID:            23012,
		Name:            "egress DST-VC credit overflow",
		Description:     defaultAlwaysFatalErr.Description,
		PotentialFatal:  defaultAlwaysFatalErr.PotentialFatal,
		AlwaysFatal:     defaultAlwaysFatalErr.AlwaysFatal,
		Impact:          defaultAlwaysFatalErr.Impact,
		Recovery:        defaultAlwaysFatalErr.Recovery,
		OtherImpact:     defaultAlwaysFatalErr.OtherImpact,
		RequiredActions: defaultAlwaysFatalErr.RequiredActions,
	},
	23013: {
		DocumentVersion: "DU-09883-001_v1.3 (October 2023)",

		SXID:            23013,
		Name:            "egress DST-VC credit underflow",
		Description:     defaultAlwaysFatalErr.Description,
		PotentialFatal:  defaultAlwaysFatalErr.PotentialFatal,
		AlwaysFatal:     defaultAlwaysFatalErr.AlwaysFatal,
		Impact:          defaultAlwaysFatalErr.Impact,
		Recovery:        defaultAlwaysFatalErr.Recovery,
		OtherImpact:     defaultAlwaysFatalErr.OtherImpact,
		RequiredActions: defaultAlwaysFatalErr.RequiredActions,
	},
	23014: {
		DocumentVersion: "DU-09883-001_v1.3 (October 2023)",

		SXID:            23014,
		Name:            "ingress packet burst error",
		Description:     defaultAlwaysFatalErr.Description,
		PotentialFatal:  defaultAlwaysFatalErr.PotentialFatal,
		AlwaysFatal:     defaultAlwaysFatalErr.AlwaysFatal,
		Impact:          defaultAlwaysFatalErr.Impact,
		Recovery:        defaultAlwaysFatalErr.Recovery,
		OtherImpact:     defaultAlwaysFatalErr.OtherImpact,
		RequiredActions: defaultAlwaysFatalErr.RequiredActions,
	},
	23015: {
		DocumentVersion: "DU-09883-001_v1.3 (October 2023)",

		SXID:            23015,
		Name:            "ingress packet sticky error",
		Description:     defaultAlwaysFatalErr.Description,
		PotentialFatal:  defaultAlwaysFatalErr.PotentialFatal,
		AlwaysFatal:     defaultAlwaysFatalErr.AlwaysFatal,
		Impact:          defaultAlwaysFatalErr.Impact,
		Recovery:        defaultAlwaysFatalErr.Recovery,
		OtherImpact:     defaultAlwaysFatalErr.OtherImpact,
		RequiredActions: defaultAlwaysFatalErr.RequiredActions,
	},
	23016: {
		DocumentVersion: "DU-09883-001_v1.3 (October 2023)",

		SXID:            23016,
		Name:            "possible bubbles at ingress",
		Description:     defaultAlwaysFatalErr.Description,
		PotentialFatal:  defaultAlwaysFatalErr.PotentialFatal,
		AlwaysFatal:     defaultAlwaysFatalErr.AlwaysFatal,
		Impact:          defaultAlwaysFatalErr.Impact,
		Recovery:        defaultAlwaysFatalErr.Recovery,
		OtherImpact:     defaultAlwaysFatalErr.OtherImpact,
		RequiredActions: defaultAlwaysFatalErr.RequiredActions,
	},
	23017: {
		DocumentVersion: "DU-09883-001_v1.3 (October 2023)",

		SXID:            23017,
		Name:            "ingress credit parity error",
		Description:     defaultAlwaysFatalErr.Description,
		PotentialFatal:  defaultAlwaysFatalErr.PotentialFatal,
		AlwaysFatal:     defaultAlwaysFatalErr.AlwaysFatal,
		Impact:          defaultAlwaysFatalErr.Impact,
		Recovery:        defaultAlwaysFatalErr.Recovery,
		OtherImpact:     defaultAlwaysFatalErr.OtherImpact,
		RequiredActions: defaultAlwaysFatalErr.RequiredActions,
	},

	// D.7 Other Notable NVSwitch SXid Error
	10001: {
		DocumentVersion: "DU-09883-001_v1.3 (October 2023)",

		SXID:           10001,
		Name:           "Host_priv_error",
		Description:    "The errors are not fatal to the fabric/system, but they might be followed by other fatal events.",
		PotentialFatal: true,
		AlwaysFatal:    false,
		Impact:         "",
		Recovery:       "",
		OtherImpact:    "",
	},
	10002: {
		DocumentVersion: "DU-09883-001_v1.3 (October 2023)",

		SXID:           10002,
		Name:           "Host_priv_timeout",
		Description:    "The errors are not fatal to the fabric/system, but they might be followed by other fatal events.",
		PotentialFatal: true,
		AlwaysFatal:    false,
		Impact:         "",
		Recovery:       "",
		OtherImpact:    "",
	},
	10003: {
		DocumentVersion: "DU-09883-001_v1.3 (October 2023)",

		SXID:           10003,
		Name:           "Host_unhandled_interrupt",
		Description:    "This SXid error is never expected to occur.",
		PotentialFatal: true,
		AlwaysFatal:    true,
		Impact:         "If it occurs, it is fatal to the fabric/system.",
		Recovery:       "To recover, it will require a reset to all GPUs and NVSwitches (refer to section D.9).",
		OtherImpact:    "If the error is observed on a Trunk port, the partitions that are using NVSwitch trunk ports will be affected.",
	},
	10004: {
		DocumentVersion: "DU-09883-001_v1.3 (October 2023)",

		SXID:           10004,
		Name:           "Host_thermal_event_start",
		Description:    "Related to thermal events, which are not directly fatal to the fabric/system, but they indicate that system cooling might be insufficient.",
		PotentialFatal: true,
		AlwaysFatal:    false,
		Impact:         "This error might force the specified NVSwitch Links to enter power saving mode (Single Lane Mode) and impact over the NVLink throughput.",
		Recovery:       "Ensure that the system cooling is sufficient.",
		OtherImpact:    "",
	},
	10005: {
		DocumentVersion: "DU-09883-001_v1.3 (October 2023)",

		SXID:           10005,
		Name:           "Host_thermal_event_end",
		Description:    "Related to thermal events, which are not directly fatal to the fabric/system, but they indicate that system cooling might be insufficient.",
		PotentialFatal: true,
		AlwaysFatal:    false,
		Impact:         "",
		Recovery:       "Ensure that the system cooling is sufficient.",
		OtherImpact:    "",
	},
}

// D.5 Fatal NVSwitch SXid Errors
var defaultPotentialFatalErr = Detail{
	Description:    "The hypervisor must track these SXid source ports (NVLink) to determine whether the error occurred on an NVSwitch trunk port or NVSwitch access port. The fatal SXid will be propagated to the GPU as Xid 74 when applicable.",
	PotentialFatal: true,
	AlwaysFatal:    false,
	Impact: `If the error occurred on an NVSwitch access port, the impact will be limited to the corresponding guest VM. To recover, shut down the guest VM.

If the errors occurred on an NVSwitch trunk port, to reset the trunk ports and recover, shut down the guest VM partitions that are crossing the trunk port. The partitions can be recreated. Currently, the partitions that are using NVSwitch trunk ports are the 16x GPU partition and the 8x GPU partitions with four GPUs per baseboard.
`,
	Recovery:    "Restart the guest VM.",
	OtherImpact: "",

	RequiredActions: common.RequiredActions{
		RebootSystem: true,
	},
}

// D.6 Always Fatal NVSwitch SXid Errors
var defaultAlwaysFatalErr = Detail{
	Description: `Always fatal to the entire fabric/system. After an always fatal SXid error has occurred, the guest VM partitions need to be shut down and one of the following tasks must occur:

1. The host needs to be restarted.
2. After the NVSwitches and GPUs are SBRed, restart the Service VM restart.

`,
	PotentialFatal: true,
	AlwaysFatal:    true,
	Impact:         `Always fatal to the entire fabric/system.`,
	Recovery:       "Restart the guest VM.",
	OtherImpact:    "",

	RequiredActions: common.RequiredActions{
		RebootSystem: true,
	},
}
