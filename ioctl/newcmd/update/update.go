// Copyright (c) 2022 IoTeX
// This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

package update

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/iotexproject/iotex-core/ioctl"
	"github.com/iotexproject/iotex-core/ioctl/config"
)

// Multi-language support
var (
	_shorts = map[config.Language]string{
		config.English: "Update password for IoTeX account",
		config.Chinese: "为IoTeX账户更新密码",
	}
	_uses = map[config.Language]string{
		config.English: "update [ALIAS|ADDRESS]",
		config.Chinese: "update [别名|地址]",
	}
	_flagUsages = map[config.Language]string{
		config.English: `set version type, "stable" or "unstable"`,
		config.Chinese: `设置版本类型, "稳定版" 或 "非稳定版"`,
	}
	_invalidVersionType = map[config.Language]string{
		config.English: "invalid version-type flag:%s",
		config.Chinese: "无效版本状态:%s",
	}
	_resultSuccess = map[config.Language]string{
		config.English: "ioctl is up-to-date now.",
		config.Chinese: "ioctl 现已更新完毕。",
	}
	_resultFail = map[config.Language]string{
		config.English: "failed to update ioctl",
		config.Chinese: "ioctl 更新失败",
	}
	_resultInfo = map[config.Language]string{
		config.English: "Downloading the latest %s version ...\n",
		config.Chinese: "正在下载最新的 %s 版本 ...\n",
	}
)

// NewUpdateCmd represents the update command
func NewUpdateCmd(c ioctl.Client) *cobra.Command {
	var versionType string
	MustSelectTranslation := func(in map[config.Language]string) string {
		translation, _ := c.SelectTranslation(in)
		return translation
	}

	use := MustSelectTranslation(_uses)
	short := MustSelectTranslation(_shorts)
	flagUsage := MustSelectTranslation(_flagUsages)
	success := MustSelectTranslation(_resultSuccess)
	fail := MustSelectTranslation(_resultFail)
	info := MustSelectTranslation(_resultInfo)
	invalidVersionType := MustSelectTranslation(_invalidVersionType)

	uc := &cobra.Command{
		Use:   use,
		Short: short,
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			var cmdString string
			switch versionType {
			case "stable":
				cmdString = "curl --silent https://raw.githubusercontent.com/iotexproject/" + "iotex-core/master/install-cli.sh | sh"
			case "unstable":
				cmdString = "curl --silent https://raw.githubusercontent.com/iotexproject/" + "iotex-core/master/install-cli.sh | sh -s \"unstable\""
			default:
				return errors.New(fmt.Sprintf(invalidVersionType, versionType))
			}
			_, err := c.ReadSecret()
			if err != nil {
				return errors.Wrap(err, fail)
			}
			cmd.Println(fmt.Sprintf(info, versionType))

			if err := c.Execute(cmdString); err != nil {
				return errors.Wrap(err, fail)
			}
			cmd.Println(success)
			return nil

		},
	}

	uc.Flags().StringVarP(&versionType, "version-type", "t", "stable",
		flagUsage)

	return uc
}
