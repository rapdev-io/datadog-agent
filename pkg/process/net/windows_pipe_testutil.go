// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2024-present Datadog, Inc.

//go:build test && windows

package net

// OverrideSystemProbeNamedPipeConfig sets the active named pipe path and its DACL for
// System Probe connections.
// This is used by tests only to avoid conflicts with an existing locally installed Datadog agent.
func OverrideSystemProbeNamedPipeConfig(path string, securityDescriptor string) {
	if path != "" {
		systemProbePipeName = path
	}

	if securityDescriptor != "" {
		systemProbePipSecurityDescriptor = securityDescriptor
	}
}