// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-2020 Datadog, Inc.

// +build linux_bpf

package probe

import (
	"github.com/DataDog/datadog-agent/pkg/security/rules"
)

var unlinkTables = []string{
	"unlink_policy",
	"unlink_path_inode_discarders",
}

func unlinkOnNewDiscarder(rs *rules.RuleSet, event *Event, probe *Probe, discarder Discarder) error {
	field := discarder.Field

	switch field {
	case "process.filename":
		return discardProcessFilename(probe, "unlink_process_discarders", event)
	case "unlink.filename":
		fsEvent := event.Unlink
		table := "unlink_path_inode_discarders"

		value, err := event.GetFieldValue(field)
		if err != nil {
			return err
		}

		isDiscarded, err := discardParentInode(probe, rs, "unlink", value.(string), fsEvent.MountID, fsEvent.Inode, table)
		if !isDiscarded || err != nil {
			// not able to discard the parent then only discard the filename
			_, err = discardInode(probe, fsEvent.MountID, fsEvent.Inode, table)
		}

		return err
	}
	return &ErrDiscarderNotSupported{Field: field}
}
