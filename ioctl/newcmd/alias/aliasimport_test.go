// Copyright (c) 2022 IoTeX Foundation
// This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.
package alias

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/iotexproject/iotex-core/ioctl/config"
	"github.com/iotexproject/iotex-core/ioctl/util"
	"github.com/iotexproject/iotex-core/test/mock/mock_ioctlclient"
)

func TestNewAliasImportCmd(t *testing.T) {
	require := require.New(t)
	ctrl := gomock.NewController(t)
	client := mock_ioctlclient.NewMockClient(ctrl)
	client.EXPECT().SelectTranslation(gomock.Any()).Return("mockTranslation", config.English).Times(24)
	client.EXPECT().Config().Return(config.Config{
		Aliases: map[string]string{
			"mhs2": "io19sdfxkwegeaenvxk2kjqf98al52gm56wa2eqks",
			"yqr":  "io1cl6rl2ev5dfa988qmgzg2x4hfazmp9vn2g66ng",
			"mhs":  "io1tyc2yt68qx7hmxl2rhssm99jxnhrccz9sneh08",
		},
	}).Times(4)
	client.EXPECT().AliasMap().Return(map[string]string{
		"mhs2": "io19sdfxkwegeaenvxk2kjqf98al52gm56wa2eqks",
		"yqr":  "io1cl6rl2ev5dfa988qmgzg2x4hfazmp9vn2g66ng",
		"mhs":  "io1tyc2yt68qx7hmxl2rhssm99jxnhrccz9sneh08",
	}).AnyTimes()

	t.Run("invalid flag", func(t *testing.T) {
		cmd := NewAliasImportCmd(client)
		result, err := util.ExecuteCmd(cmd, "-f", "test", "test")
		require.Error(err)
		require.Contains(err.Error(), "invalid flag")
		require.NotNil(result)
	})

	t.Run("import alias with json format", func(t *testing.T) {
		cmd := NewAliasImportCmd(client)
		result, err := util.ExecuteCmd(cmd, `{"name":"test","address":"io1uwnr55vqmhf3xeg5phgurlyl702af6eju542sx"}`)
		require.NoError(err)
		require.NotNil(result)
	})

	t.Run("import alias with yaml format", func(t *testing.T) {
		cmd := NewAliasImportCmd(client)
		result, err := util.ExecuteCmd(cmd, "-f", "yaml", `aliases:
- name: mhs2
  address: io19sdfxkwegeaenvxk2kjqf98al52gm56wa2eqks
- name: yqr
  address: io1cl6rl2ev5dfa988qmgzg2x4hfazmp9vn2g66ng
- name: mhs
  address: io1tyc2yt68qx7hmxl2rhssm99jxnhrccz9sneh08`)
		require.NoError(err)
		require.NotNil(result)
	})

	t.Run("force import", func(t *testing.T) {
		cmd := NewAliasImportCmd(client)
		result, err := util.ExecuteCmd(cmd, "-F", `{"name":"io1uwnr55vqmhf3xeg5phgurlyl702af6eju542sx","address":"io1uwnr55vqmhf3xeg5phgurlyl702af6eju542sx"}`)
		require.NoError(err)
		require.NotNil(result)
	})

	t.Run("failed to unmarshal json imported aliases", func(t *testing.T) {
		cmd := NewAliasImportCmd(client)
		result, err := util.ExecuteCmd(cmd, `""{"name":"d","address":""}""`)
		require.Error(err)
		require.Contains(err.Error(), "failed to unmarshal imported aliases")
		require.NotNil(result)
	})

	t.Run("failed to unmarshal yaml imported aliases", func(t *testing.T) {
		cmd := NewAliasImportCmd(client)
		result, err := util.ExecuteCmd(cmd, "-f", "yaml", `aliases:
			name: mhs2
			address: io19sdfxkwegeaenvxk2kjqf98al52gm56wa2eqks`)
		require.Error(err)
		require.Contains(err.Error(), "failed to unmarshal imported aliases")
		require.NotNil(result)
	})
}
