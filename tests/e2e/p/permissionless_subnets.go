// Copyright (C) 2019-2022, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package p

import (
	"context"
	"fmt"
	"time"

	ginkgo "github.com/onsi/ginkgo/v2"

	"github.com/onsi/gomega"

	"github.com/lasthyphen/dijetsnodego/genesis"
	"github.com/lasthyphen/dijetsnodego/ids"
	"github.com/lasthyphen/dijetsnodego/snow/choices"
	"github.com/lasthyphen/dijetsnodego/tests"
	"github.com/lasthyphen/dijetsnodego/tests/e2e"
	"github.com/lasthyphen/dijetsnodego/utils/constants"
	"github.com/lasthyphen/dijetsnodego/utils/units"
	"github.com/lasthyphen/dijetsnodego/vms/avm"
	"github.com/lasthyphen/dijetsnodego/vms/components/djtx"
	"github.com/lasthyphen/dijetsnodego/vms/components/verify"
	"github.com/lasthyphen/dijetsnodego/vms/platformvm"
	"github.com/lasthyphen/dijetsnodego/vms/platformvm/reward"
	"github.com/lasthyphen/dijetsnodego/vms/platformvm/signer"
	"github.com/lasthyphen/dijetsnodego/vms/platformvm/status"
	"github.com/lasthyphen/dijetsnodego/vms/platformvm/validator"
	"github.com/lasthyphen/dijetsnodego/vms/secp256k1fx"
	"github.com/lasthyphen/dijetsnodego/wallet/subnet/primary"
	"github.com/lasthyphen/dijetsnodego/wallet/subnet/primary/common"
)

var _ = e2e.DescribePChain("[Permissionless Subnets]", func() {
	ginkgo.It("subnets operations",
		// use this for filtering tests by labels
		// ref. https://onsi.github.io/ginkgo/#spec-labels
		ginkgo.Label(
			"require-network-runner",
			"xp",
			"permissionless-subnets",
		),
		func() {
			ginkgo.By("reload initial snapshot for test independence", func() {
				err := e2e.Env.RestoreInitialState(true /*switchOffNetworkFirst*/)
				gomega.Expect(err).Should(gomega.BeNil())
			})

			rpcEps := e2e.Env.GetURIs()
			gomega.Expect(rpcEps).ShouldNot(gomega.BeEmpty())
			nodeURI := rpcEps[0]

			tests.Outf("{{blue}} setting up keys {{/}}\n")
			testKey := genesis.EWOQKey
			keyChain := secp256k1fx.NewKeychain(testKey)

			var baseWallet primary.Wallet
			ginkgo.By("setup wallet", func() {
				var err error
				ctx, cancel := context.WithTimeout(context.Background(), e2e.DefaultWalletCreationTimeout)
				baseWallet, err = primary.NewWalletFromURI(ctx, nodeURI, keyChain)
				cancel()
				gomega.Expect(err).Should(gomega.BeNil())
			})

			pWallet := baseWallet.P()
			pChainClient := platformvm.NewClient(nodeURI)
			xWallet := baseWallet.X()
			xChainClient := avm.NewClient(nodeURI, xWallet.BlockchainID().String())
			xChainID := xWallet.BlockchainID()

			owner := &secp256k1fx.OutputOwners{
				Threshold: 1,
				Addrs: []ids.ShortID{
					testKey.PublicKey().Address(),
				},
			}

			var subnetID ids.ID
			ginkgo.By("create a permissioned subnet", func() {
				var err error
				ctx, cancel := context.WithTimeout(context.Background(), e2e.DefaultConfirmTxTimeout)
				subnetID, err = pWallet.IssueCreateSubnetTx(
					owner,
					common.WithContext(ctx),
				)
				cancel()
				gomega.Expect(subnetID, err).Should(gomega.Not(gomega.Equal(constants.PrimaryNetworkID)))

				ctx, cancel = context.WithTimeout(context.Background(), e2e.DefaultConfirmTxTimeout)
				txStatus, err := pChainClient.GetTxStatus(ctx, subnetID)
				cancel()
				gomega.Expect(txStatus.Status, err).To(gomega.Equal(status.Committed))
			})

			var subnetAssetID ids.ID
			ginkgo.By("create a custom asset for the permissionless subnet", func() {
				var err error
				ctx, cancel := context.WithTimeout(context.Background(), e2e.DefaultWalletCreationTimeout)
				subnetAssetID, err = xWallet.IssueCreateAssetTx(
					"RnM",
					"RNM",
					9,
					map[uint32][]verify.State{
						0: {
							&secp256k1fx.TransferOutput{
								Amt:          100 * units.MegaDjtx,
								OutputOwners: *owner,
							},
						},
					},
					common.WithContext(ctx),
				)
				cancel()
				gomega.Expect(err).Should(gomega.BeNil())

				ctx, cancel = context.WithTimeout(context.Background(), e2e.DefaultConfirmTxTimeout)
				txStatus, err := xChainClient.GetTxStatus(ctx, subnetAssetID)
				cancel()
				gomega.Expect(txStatus, err).To(gomega.Equal(choices.Accepted))
			})

			ginkgo.By(fmt.Sprintf("Send 100 MegaDjtx of asset %s to the P-chain", subnetAssetID), func() {
				ctx, cancel := context.WithTimeout(context.Background(), e2e.DefaultWalletCreationTimeout)
				exportTxID, err := xWallet.IssueExportTx(
					constants.PlatformChainID,
					[]*djtx.TransferableOutput{
						{
							Asset: djtx.Asset{
								ID: subnetAssetID,
							},
							Out: &secp256k1fx.TransferOutput{
								Amt:          100 * units.MegaDjtx,
								OutputOwners: *owner,
							},
						},
					},
					common.WithContext(ctx),
				)
				cancel()
				gomega.Expect(err).Should(gomega.BeNil())

				ctx, cancel = context.WithTimeout(context.Background(), e2e.DefaultConfirmTxTimeout)
				txStatus, err := xChainClient.GetTxStatus(ctx, exportTxID)
				cancel()
				gomega.Expect(txStatus, err).To(gomega.Equal(choices.Accepted))
			})

			ginkgo.By(fmt.Sprintf("Import the 100 MegaDjtx of asset %s from the X-chain into the P-chain", subnetAssetID), func() {
				ctx, cancel := context.WithTimeout(context.Background(), e2e.DefaultWalletCreationTimeout)
				importTxID, err := pWallet.IssueImportTx(
					xChainID,
					owner,
					common.WithContext(ctx),
				)
				cancel()
				gomega.Expect(err).Should(gomega.BeNil())

				ctx, cancel = context.WithTimeout(context.Background(), e2e.DefaultConfirmTxTimeout)
				txStatus, err := pChainClient.GetTxStatus(ctx, importTxID)
				cancel()
				gomega.Expect(txStatus.Status, err).To(gomega.Equal(status.Committed))
			})

			ginkgo.By("make subnet permissionless", func() {
				ctx, cancel := context.WithTimeout(context.Background(), e2e.DefaultConfirmTxTimeout)
				transformSubnetTxID, err := pWallet.IssueTransformSubnetTx(
					subnetID,
					subnetAssetID,
					50*units.MegaDjtx,
					100*units.MegaDjtx,
					reward.PercentDenominator,
					reward.PercentDenominator,
					1,
					100*units.MegaDjtx,
					time.Second,
					365*24*time.Hour,
					0,
					1,
					5,
					.80*reward.PercentDenominator,
					common.WithContext(ctx),
				)
				cancel()
				gomega.Expect(err).Should(gomega.BeNil())

				ctx, cancel = context.WithTimeout(context.Background(), e2e.DefaultConfirmTxTimeout)
				txStatus, err := pChainClient.GetTxStatus(ctx, transformSubnetTxID)
				cancel()
				gomega.Expect(txStatus.Status, err).To(gomega.Equal(status.Committed))
			})

			validatorStartTime := time.Now().Add(time.Minute)
			ginkgo.By("add permissionless validator", func() {
				ctx, cancel := context.WithTimeout(context.Background(), e2e.DefaultConfirmTxTimeout)
				addSubnetValidatorTxID, err := pWallet.IssueAddPermissionlessValidatorTx(
					&validator.SubnetValidator{
						Validator: validator.Validator{
							NodeID: genesis.LocalConfig.InitialStakers[0].NodeID,
							Start:  uint64(validatorStartTime.Unix()),
							End:    uint64(validatorStartTime.Add(5 * time.Second).Unix()),
							Wght:   25 * units.MegaDjtx,
						},
						Subnet: subnetID,
					},
					&signer.Empty{},
					subnetAssetID,
					&secp256k1fx.OutputOwners{},
					&secp256k1fx.OutputOwners{},
					reward.PercentDenominator,
					common.WithContext(ctx),
				)
				cancel()
				gomega.Expect(err).Should(gomega.BeNil())

				ctx, cancel = context.WithTimeout(context.Background(), e2e.DefaultConfirmTxTimeout)
				txStatus, err := pChainClient.GetTxStatus(ctx, addSubnetValidatorTxID)
				cancel()
				gomega.Expect(txStatus.Status, err).To(gomega.Equal(status.Committed))
			})

			delegatorStartTime := validatorStartTime
			ginkgo.By("add permissionless delegator", func() {
				ctx, cancel := context.WithTimeout(context.Background(), e2e.DefaultConfirmTxTimeout)
				addSubnetDelegatorTxID, err := pWallet.IssueAddPermissionlessDelegatorTx(
					&validator.SubnetValidator{
						Validator: validator.Validator{
							NodeID: genesis.LocalConfig.InitialStakers[0].NodeID,
							Start:  uint64(delegatorStartTime.Unix()),
							End:    uint64(delegatorStartTime.Add(5 * time.Second).Unix()),
							Wght:   25 * units.MegaDjtx,
						},
						Subnet: subnetID,
					},
					subnetAssetID,
					&secp256k1fx.OutputOwners{},
					common.WithContext(ctx),
				)
				cancel()
				gomega.Expect(err).Should(gomega.BeNil())

				ctx, cancel = context.WithTimeout(context.Background(), e2e.DefaultConfirmTxTimeout)
				txStatus, err := pChainClient.GetTxStatus(ctx, addSubnetDelegatorTxID)
				cancel()
				gomega.Expect(txStatus.Status, err).To(gomega.Equal(status.Committed))
			})
		})
})
