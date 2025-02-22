// Copyright 2019 The Kubernetes Authors.
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

package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"sigs.k8s.io/krew/internal/indexmigration"
	"sigs.k8s.io/krew/internal/receiptsmigration"
	"sigs.k8s.io/krew/pkg/constants"
)

// todo(corneliusweig) remove migration code with v0.4
// systemCmd represents the system command
var systemCmd = &cobra.Command{
	Use:   "system",
	Short: "Perform krew maintenance tasks",
	Long: `Perform krew maintenance tasks such as migrating to a new krew-home layout.

This command will be removed without further notice from future versions of krew.
`,
	Args:   cobra.NoArgs,
	Hidden: true,
}

// TODO(corneliusweig) remove migration code with v0.4
// systemCmd represents the system command
var receiptsUpgradeCmd = &cobra.Command{
	Use:   "receipts-upgrade",
	Short: "Perform a migration of the krew home",
	Long: `Krew became more awesome! To use the new features, you need to run this
one-time migration, which will reinstall all current plugins.

This command will be removed without further notice from future versions of krew.
`,
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		return receiptsmigration.Migrate(paths)
	},
	PreRunE: ensureIndexesUpdated,
}

var indexUpgradeCmd = &cobra.Command{
	Use:   "index-upgrade",
	Short: "Perform a migration of the krew index",
	Long: `Krew became more awesome! To use the new features, you need to run this
one-time migration, which will enable installing plugins from custom indexes.

This command will be removed without further notice from future versions of krew.
`,
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		return indexmigration.Migrate(paths)
	},
}

func init() {
	if _, ok := os.LookupEnv(constants.EnableMultiIndexSwitch); ok {
		systemCmd.AddCommand(indexUpgradeCmd)
	}
	systemCmd.AddCommand(receiptsUpgradeCmd)
	rootCmd.AddCommand(systemCmd)
}
