// Copyright 2015 The appc Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"github.com/spf13/cobra"
)

var (
	readOnly bool
	cmdMount = &cobra.Command{
		Use:     "mount [command]",
		Aliases: []string{"volume"},
		Short:   "Manage mount points",
	}
	cmdAddMount = &cobra.Command{
		Use:     "add NAME PATH",
		Short:   "Add a mount point",
		Example: "acbuild mount add htmlfiles /usr/share/nginx/html --read-only",
		Run:     runWrapper(runAddMount),
	}
	cmdRmMount = &cobra.Command{
		Use:     "remove NAME/PATH",
		Aliases: []string{"rm"},
		Short:   "Remove a mount point with the given name or path",
		Example: "acbuild mount remove /usr/share/nginx/html",
		Run:     runWrapper(runRmMount),
	}
)

func init() {
	cmdAcbuild.AddCommand(cmdMount)
	cmdMount.AddCommand(cmdAddMount)
	cmdMount.AddCommand(cmdRmMount)

	cmdAddMount.Flags().BoolVar(&readOnly, "read-only", false, "(appc only) Set the mount point to be read only")
}

func runAddMount(cmd *cobra.Command, args []string) (exit int) {
	if len(args) == 0 {
		cmd.Usage()
		return 1
	}
	if len(args) != 2 {
		stderr("mount add: incorrect number of arguments")
		return 1
	}

	if debug {
		if readOnly {
			stderr("Adding read only mount point %q=%q", args[0], args[1])
		} else {
			stderr("Adding mount point %q=%q", args[0], args[1])
		}
	}

	a, err := newACBuild()
	if err != nil {
		stderr("%v", err)
		return 1
	}
	err = a.AddMount(args[0], args[1], readOnly)

	if err != nil {
		stderr("mount add: %v", err)
		return getErrorCode(err)
	}

	return 0
}

func runRmMount(cmd *cobra.Command, args []string) (exit int) {
	if len(args) == 0 {
		cmd.Usage()
		return 1
	}
	if len(args) != 1 {
		stderr("mount remove: incorrect number of arguments")
		return 1
	}

	if debug {
		stderr("Removing mount point %q", args[0])
	}

	a, err := newACBuild()
	if err != nil {
		stderr("%v", err)
		return 1
	}
	err = a.RemoveMount(args[0])

	if err != nil {
		stderr("mount remove: %v", err)
		return getErrorCode(err)
	}

	return 0
}
