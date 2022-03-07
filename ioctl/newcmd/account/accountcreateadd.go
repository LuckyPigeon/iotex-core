// Copyright (c) 2022 IoTeX Foundation
// This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

package account

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"

	"github.com/iotexproject/go-pkgs/crypto"
	"github.com/iotexproject/iotex-core/ioctl"
	"github.com/iotexproject/iotex-core/ioctl/config"
	"github.com/iotexproject/iotex-core/ioctl/validator"
)

// Multi-language support
var (
	createAddCmdShorts = map[config.Language]string{
		config.English: "Create new account for ioctl",
		config.Chinese: "为ioctl创建新账户",
	}
	createAddCmdUses = map[config.Language]string{
		config.English: "createadd ALIAS",
		config.Chinese: "createadd 别名",
	}
	invalidAlias = map[config.Language]string{
		config.English: "invalid alias",
		config.Chinese: "无效别名",
	}
	aliasHasAlreadyUsed = map[config.Language]string{
		config.English: "** Alias \"%s\" has already used for %s\n" +
			"Overwriting the account will keep the previous keystore file stay, " +
			"but bind the alias to the new one.\nWould you like to continue?\n",
		config.Chinese: "** 这个别名 \"%s\" 已被 %s 使用!\n" +
			"复写帐户后先前的 keystore 文件将会留存!\n" +
			"但底下的别名将绑定为新的。您是否要继续？",
	}
	quit = map[config.Language]string{
		config.English: "quit",
		config.Chinese: "退出",
	}
	outputMessage = map[config.Language]string{
		config.English: "New account \"%s\" is created.\n" +
			"Please Keep your password, or you will lose your private key.",
		config.Chinese: "新帐户 \"%s\" 已建立。\n" +
			"请保护好您的密码，否则您会失去您的私钥。",
	}
)

// NewAccountCreateAdd represents the account createadd command
func NewAccountCreateAdd(client ioctl.Client) *cobra.Command {
	use, _ := client.SelectTranslation(createAddCmdUses)
	short, _ := client.SelectTranslation(createAddCmdShorts)
	invalidAlias, _ := client.SelectTranslation(invalidAlias)
	aliasHasAlreadyUsed, _ := client.SelectTranslation(aliasHasAlreadyUsed)
	quit, _ := client.SelectTranslation(quit)
	failToWriteToConfigFile, _ := client.SelectTranslation(failToWriteToConfigFile)
	failToGenerateNewPrivateKey, _ := client.SelectTranslation(failToGenerateNewPrivateKey)
	failToGenerateNewPrivateKeySm2, _ := client.SelectTranslation(failToGenerateNewPrivateKeySm2)
	outputMessage, _ := client.SelectTranslation(outputMessage)

	return &cobra.Command{
		Use:   use,
		Short: short,
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			aliases := client.AliasMap()

			// Validate inputs
			if err := validator.ValidateAlias(args[0]); err != nil {
				// return output.NewError(output.ValidationError, "invalid alias", err)
				return errors.Wrap(err, invalidAlias)
			}

			// if addr, ok := config.ReadConfig.Aliases[args[0]]; ok {
			if addr, ok := aliases[args[0]]; ok {
				if !client.AskToConfirm(fmt.Sprintf(aliasHasAlreadyUsed, args[0], addr)) {
					client.PrintInfo(quit)
					return nil
				}
			}

			var private crypto.PrivateKey
			var err error
			if client.IsCryptoSm2() {
				private, err = crypto.GenerateKey()
				if err != nil {
					return errors.Wrap(err, failToGenerateNewPrivateKey)
				}
			} else {
				private, err = crypto.GenerateKeySm2()
				if err != nil {
					return errors.Wrap(err, failToGenerateNewPrivateKeySm2)
				}
			}
			aliases[args[0]] = private.PublicKey().Address().String()
			out, err := yaml.Marshal(&config.ReadConfig)
			if err != nil {
				return errors.Wrapf(err, failToWriteToConfigFile, config.DefaultConfigFile)
			}
			if err := os.WriteFile(config.DefaultConfigFile, out, 0600); err != nil {
				return errors.Wrapf(err, failToWriteToConfigFile, config.DefaultConfigFile)
			}
			private.Zero()
			client.PrintInfo(fmt.Sprintf(outputMessage, args[0]))
			return nil
		},
	}
}
