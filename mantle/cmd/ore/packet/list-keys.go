// Copyright 2017 CoreOS, Inc.
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

package packet

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	cmdListKeys = &cobra.Command{
		Use:   "list-keys",
		Short: "List Packet SSH keys",
		RunE:  runListKeys,
	}
)

func init() {
	Packet.AddCommand(cmdListKeys)
}

func runListKeys(cmd *cobra.Command, args []string) error {
	if len(args) != 0 {
		fmt.Fprintf(os.Stderr, "Unrecognized args in packet list-keys cmd: %v\n", args)
		os.Exit(2)
	}

	keys, err := API.ListKeys()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't list keys: %v\n", err)
		os.Exit(1)
	}

	for _, key := range keys {
		fmt.Println(key.Label)
	}
	return nil
}
