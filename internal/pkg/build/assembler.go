// Copyright (c) Contributors to the Apptainer project, established as
//   Apptainer a Series of LF Projects LLC.
//   For website terms of use, trademark policy, privacy policy and other
//   project policies see https://lfprojects.org/policies
// Copyright (c) 2018-2019, Sylabs Inc. All rights reserved.
// This software is licensed under a 3-clause BSD license. Please consult the
// LICENSE.md file distributed with the sources of this project regarding your
// rights to use or distribute this software.

package build

import (
	"github.com/apptainer/apptainer/pkg/build/types"
)

// Assembler is responsible for assembling an image from a bundle.
// For example a bundle may be holding multiple file systems intended
// to be separate partitions within a SIF image. The assembler would need
// to detect these directories and make sure it properly assembles the SIF
// with them as partitions.
type Assembler interface {
	Assemble(*types.Bundle, string) error
}
