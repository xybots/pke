// Copyright © 2019 Banzai Cloud
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

package app

import (
	"flag"
	"syscall"

	"github.com/banzaicloud/pke/cmd/pke/app/cmd"
	"github.com/banzaicloud/pke/cmd/pke/app/constants"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func Run(gitVersion, gitCommit, gitTreeState, buildDate string) error {
	// set global umask
	syscall.Umask(constants.Umask)

	cobra.EnableCommandSorting = false
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)

	c := cmd.NewPKECommand(gitVersion, gitCommit, gitTreeState, buildDate)
	return c.Execute()
}
