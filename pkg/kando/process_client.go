// Copyright 2020 The Kanister Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package kando

import (
	"github.com/spf13/cobra"
)

const (
	processAsJSONFlagName = "as-json"
)

func newProcessClientCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "client <command>",
		Short: "Send commands to the process server",
	}
	cmd.AddCommand(newProcessClientCreateCommand())
	cmd.AddCommand(newProcessClientGetCommand())
	cmd.AddCommand(newProcessClientListCommand())
	cmd.AddCommand(newProcessClientSignalCommand())
	cmd.AddCommand(newProcessClientOutputCommand())
	cmd.PersistentFlags().BoolP(processAsJSONFlagName, "j", false, "Display output as json")
	return cmd
}

func processAsJSONFlagValue(cmd *cobra.Command) bool {
	b, err := cmd.Flags().GetBool(processAsJSONFlagName)
	if err != nil {
		panic(err.Error())
	}
	return b
}
