// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package MStableContract

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// AmpData is an auto generated low-level Go binding around an user-defined struct.
type AmpData struct {
	InitialA      uint64
	TargetA       uint64
	RampStartTime uint64
	RampEndTime   uint64
}

// BassetData is an auto generated low-level Go binding around an user-defined struct.
type BassetData struct {
	Ratio        *big.Int
	VaultBalance *big.Int
}

// BassetPersonal is an auto generated low-level Go binding around an user-defined struct.
type BassetPersonal struct {
	Addr       common.Address
	Integrator common.Address
	HasTxFee   bool
	Status     uint8
}

// FeederConfig is an auto generated low-level Go binding around an user-defined struct.
type FeederConfig struct {
	Supply *big.Int
	A      *big.Int
	Limits WeightLimits
}

// InvariantConfig is an auto generated low-level Go binding around an user-defined struct.
type InvariantConfig struct {
	A      *big.Int
	Limits WeightLimits
}

// WeightLimits is an auto generated low-level Go binding around an user-defined struct.
type WeightLimits struct {
	Min *big.Int
	Max *big.Int
}

// MStableABI is the input ABI used to generate the binding from.
const MStableABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nexus\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_mAsset\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"bAssets\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newIntegrator\",\"type\":\"address\"}],\"name\":\"BassetsMigrated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cacheSize\",\"type\":\"uint256\"}],\"name\":\"CacheSizeChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"swapFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"redemptionFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"govFee\",\"type\":\"uint256\"}],\"name\":\"FeesChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"output\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"input\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"inputQuantity\",\"type\":\"uint256\"}],\"name\":\"Minted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"output\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"inputs\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"inputQuantities\",\"type\":\"uint256[]\"}],\"name\":\"MintedMulti\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"redeemer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"mAssetQuantity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"output\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"outputQuantity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"scaledFee\",\"type\":\"uint256\"}],\"name\":\"Redeemed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"redeemer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"mAssetQuantity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"outputs\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"outputQuantity\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"scaledFee\",\"type\":\"uint256\"}],\"name\":\"RedeemedMulti\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"currentA\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"targetA\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rampEndTime\",\"type\":\"uint256\"}],\"name\":\"StartRampA\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"currentA\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"StopRampA\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"swapper\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"input\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"output\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"outputAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"Swapped\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"min\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"max\",\"type\":\"uint128\"}],\"name\":\"WeightLimitsChanged\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"collectPendingFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"collectPlatformInterest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"mintAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newSupply\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"data\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"swapFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"redemptionFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"govFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pendingFees\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cacheSize\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"initialA\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"targetA\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"rampStartTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"rampEndTime\",\"type\":\"uint64\"}],\"internalType\":\"structAmpData\",\"name\":\"ampData\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint128\",\"name\":\"min\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"max\",\"type\":\"uint128\"}],\"internalType\":\"structWeightLimits\",\"name\":\"weightLimits\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bAsset\",\"type\":\"address\"}],\"name\":\"getBasset\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"integrator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"hasTxFee\",\"type\":\"bool\"},{\"internalType\":\"enumBassetStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"internalType\":\"structBassetPersonal\",\"name\":\"personal\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint128\",\"name\":\"ratio\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"vaultBalance\",\"type\":\"uint128\"}],\"internalType\":\"structBassetData\",\"name\":\"vaultData\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBassets\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"integrator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"hasTxFee\",\"type\":\"bool\"},{\"internalType\":\"enumBassetStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"internalType\":\"structBassetPersonal[]\",\"name\":\"\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint128\",\"name\":\"ratio\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"vaultBalance\",\"type\":\"uint128\"}],\"internalType\":\"structBassetData[]\",\"name\":\"vaultData\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"supply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint128\",\"name\":\"min\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"max\",\"type\":\"uint128\"}],\"internalType\":\"structWeightLimits\",\"name\":\"limits\",\"type\":\"tuple\"}],\"internalType\":\"structFeederConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_inputs\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_inputQuantities\",\"type\":\"uint256[]\"}],\"name\":\"getMintMultiOutput\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"mintOutput\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_input\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_inputQuantity\",\"type\":\"uint256\"}],\"name\":\"getMintOutput\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"mintOutput\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"k\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_outputs\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_outputQuantities\",\"type\":\"uint256[]\"}],\"name\":\"getRedeemExactBassetsOutput\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"fpTokenQuantity\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_output\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_fpTokenQuantity\",\"type\":\"uint256\"}],\"name\":\"getRedeemOutput\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"bAssetOutput\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_input\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_output\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_inputQuantity\",\"type\":\"uint256\"}],\"name\":\"getSwapOutput\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"swapOutput\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_nameArg\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbolArg\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"integrator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"hasTxFee\",\"type\":\"bool\"},{\"internalType\":\"enumBassetStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"internalType\":\"structBassetPersonal\",\"name\":\"_mAsset\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"integrator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"hasTxFee\",\"type\":\"bool\"},{\"internalType\":\"enumBassetStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"internalType\":\"structBassetPersonal\",\"name\":\"_fAsset\",\"type\":\"tuple\"},{\"internalType\":\"address[]\",\"name\":\"_mpAssets\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint128\",\"name\":\"min\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"max\",\"type\":\"uint128\"}],\"internalType\":\"structWeightLimits\",\"name\":\"limits\",\"type\":\"tuple\"}],\"internalType\":\"structInvariantConfig\",\"name\":\"_config\",\"type\":\"tuple\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mAsset\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_bAssets\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"_newIntegration\",\"type\":\"address\"}],\"name\":\"migrateBassets\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_input\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_inputQuantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minOutputQuantity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"mintOutput\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_inputs\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_inputQuantities\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"_minOutputQuantity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"}],\"name\":\"mintMulti\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"mintOutput\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nexus\",\"outputs\":[{\"internalType\":\"contractINexus\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_output\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_fpTokenQuantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minOutputQuantity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"}],\"name\":\"redeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"outputQuantity\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_outputs\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_outputQuantities\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"_maxInputQuantity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"}],\"name\":\"redeemExactBassets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"fpTokenQuantity\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_inputQuantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"_minOutputQuantities\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"}],\"name\":\"redeemProportionately\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"outputQuantities\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_cacheSize\",\"type\":\"uint256\"}],\"name\":\"setCacheSize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_swapFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_redemptionFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_govFee\",\"type\":\"uint256\"}],\"name\":\"setFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"_min\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"_max\",\"type\":\"uint128\"}],\"name\":\"setWeightLimits\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_targetA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_rampEndTime\",\"type\":\"uint256\"}],\"name\":\"startRampA\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stopRampA\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_input\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_output\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_inputQuantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minOutputQuantity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"}],\"name\":\"swap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"swapOutput\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// MStable is an auto generated Go binding around an Ethereum contract.
type MStable struct {
	MStableCaller     // Read-only binding to the contract
	MStableTransactor // Write-only binding to the contract
	MStableFilterer   // Log filterer for contract events
}

// MStableCaller is an auto generated read-only Go binding around an Ethereum contract.
type MStableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MStableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MStableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MStableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MStableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MStableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MStableSession struct {
	Contract     *MStable          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MStableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MStableCallerSession struct {
	Contract *MStableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// MStableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MStableTransactorSession struct {
	Contract     *MStableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// MStableRaw is an auto generated low-level Go binding around an Ethereum contract.
type MStableRaw struct {
	Contract *MStable // Generic contract binding to access the raw methods on
}

// MStableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MStableCallerRaw struct {
	Contract *MStableCaller // Generic read-only contract binding to access the raw methods on
}

// MStableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MStableTransactorRaw struct {
	Contract *MStableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMStable creates a new instance of MStable, bound to a specific deployed contract.
func NewMStable(address common.Address, backend bind.ContractBackend) (*MStable, error) {
	contract, err := bindMStable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MStable{MStableCaller: MStableCaller{contract: contract}, MStableTransactor: MStableTransactor{contract: contract}, MStableFilterer: MStableFilterer{contract: contract}}, nil
}

// NewMStableCaller creates a new read-only instance of MStable, bound to a specific deployed contract.
func NewMStableCaller(address common.Address, caller bind.ContractCaller) (*MStableCaller, error) {
	contract, err := bindMStable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MStableCaller{contract: contract}, nil
}

// NewMStableTransactor creates a new write-only instance of MStable, bound to a specific deployed contract.
func NewMStableTransactor(address common.Address, transactor bind.ContractTransactor) (*MStableTransactor, error) {
	contract, err := bindMStable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MStableTransactor{contract: contract}, nil
}

// NewMStableFilterer creates a new log filterer instance of MStable, bound to a specific deployed contract.
func NewMStableFilterer(address common.Address, filterer bind.ContractFilterer) (*MStableFilterer, error) {
	contract, err := bindMStable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MStableFilterer{contract: contract}, nil
}

// bindMStable binds a generic wrapper to an already deployed contract.
func bindMStable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MStableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MStable *MStableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MStable.Contract.MStableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MStable *MStableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MStable.Contract.MStableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MStable *MStableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MStable.Contract.MStableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MStable *MStableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MStable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MStable *MStableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MStable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MStable *MStableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MStable.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_MStable *MStableCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MStable.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_MStable *MStableSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _MStable.Contract.Allowance(&_MStable.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_MStable *MStableCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _MStable.Contract.Allowance(&_MStable.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_MStable *MStableCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MStable.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_MStable *MStableSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _MStable.Contract.BalanceOf(&_MStable.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_MStable *MStableCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _MStable.Contract.BalanceOf(&_MStable.CallOpts, account)
}

// Data is a free data retrieval call binding the contract method 0x73d4a13a.
//
// Solidity: function data() view returns(uint256 swapFee, uint256 redemptionFee, uint256 govFee, uint256 pendingFees, uint256 cacheSize, (uint64,uint64,uint64,uint64) ampData, (uint128,uint128) weightLimits)
func (_MStable *MStableCaller) Data(opts *bind.CallOpts) (struct {
	SwapFee       *big.Int
	RedemptionFee *big.Int
	GovFee        *big.Int
	PendingFees   *big.Int
	CacheSize     *big.Int
	AmpData       AmpData
	WeightLimits  WeightLimits
}, error) {
	var out []interface{}
	err := _MStable.contract.Call(opts, &out, "data")

	outstruct := new(struct {
		SwapFee       *big.Int
		RedemptionFee *big.Int
		GovFee        *big.Int
		PendingFees   *big.Int
		CacheSize     *big.Int
		AmpData       AmpData
		WeightLimits  WeightLimits
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.SwapFee = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.RedemptionFee = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.GovFee = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.PendingFees = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.CacheSize = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.AmpData = *abi.ConvertType(out[5], new(AmpData)).(*AmpData)
	outstruct.WeightLimits = *abi.ConvertType(out[6], new(WeightLimits)).(*WeightLimits)

	return *outstruct, err

}

// Data is a free data retrieval call binding the contract method 0x73d4a13a.
//
// Solidity: function data() view returns(uint256 swapFee, uint256 redemptionFee, uint256 govFee, uint256 pendingFees, uint256 cacheSize, (uint64,uint64,uint64,uint64) ampData, (uint128,uint128) weightLimits)
func (_MStable *MStableSession) Data() (struct {
	SwapFee       *big.Int
	RedemptionFee *big.Int
	GovFee        *big.Int
	PendingFees   *big.Int
	CacheSize     *big.Int
	AmpData       AmpData
	WeightLimits  WeightLimits
}, error) {
	return _MStable.Contract.Data(&_MStable.CallOpts)
}

// Data is a free data retrieval call binding the contract method 0x73d4a13a.
//
// Solidity: function data() view returns(uint256 swapFee, uint256 redemptionFee, uint256 govFee, uint256 pendingFees, uint256 cacheSize, (uint64,uint64,uint64,uint64) ampData, (uint128,uint128) weightLimits)
func (_MStable *MStableCallerSession) Data() (struct {
	SwapFee       *big.Int
	RedemptionFee *big.Int
	GovFee        *big.Int
	PendingFees   *big.Int
	CacheSize     *big.Int
	AmpData       AmpData
	WeightLimits  WeightLimits
}, error) {
	return _MStable.Contract.Data(&_MStable.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_MStable *MStableCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _MStable.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_MStable *MStableSession) Decimals() (uint8, error) {
	return _MStable.Contract.Decimals(&_MStable.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_MStable *MStableCallerSession) Decimals() (uint8, error) {
	return _MStable.Contract.Decimals(&_MStable.CallOpts)
}

// GetBasset is a free data retrieval call binding the contract method 0x3e37bcbc.
//
// Solidity: function getBasset(address _bAsset) view returns((address,address,bool,uint8) personal, (uint128,uint128) vaultData)
func (_MStable *MStableCaller) GetBasset(opts *bind.CallOpts, _bAsset common.Address) (struct {
	Personal  BassetPersonal
	VaultData BassetData
}, error) {
	var out []interface{}
	err := _MStable.contract.Call(opts, &out, "getBasset", _bAsset)

	outstruct := new(struct {
		Personal  BassetPersonal
		VaultData BassetData
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Personal = *abi.ConvertType(out[0], new(BassetPersonal)).(*BassetPersonal)
	outstruct.VaultData = *abi.ConvertType(out[1], new(BassetData)).(*BassetData)

	return *outstruct, err

}

// GetBasset is a free data retrieval call binding the contract method 0x3e37bcbc.
//
// Solidity: function getBasset(address _bAsset) view returns((address,address,bool,uint8) personal, (uint128,uint128) vaultData)
func (_MStable *MStableSession) GetBasset(_bAsset common.Address) (struct {
	Personal  BassetPersonal
	VaultData BassetData
}, error) {
	return _MStable.Contract.GetBasset(&_MStable.CallOpts, _bAsset)
}

// GetBasset is a free data retrieval call binding the contract method 0x3e37bcbc.
//
// Solidity: function getBasset(address _bAsset) view returns((address,address,bool,uint8) personal, (uint128,uint128) vaultData)
func (_MStable *MStableCallerSession) GetBasset(_bAsset common.Address) (struct {
	Personal  BassetPersonal
	VaultData BassetData
}, error) {
	return _MStable.Contract.GetBasset(&_MStable.CallOpts, _bAsset)
}

// GetBassets is a free data retrieval call binding the contract method 0x1d3ce398.
//
// Solidity: function getBassets() view returns((address,address,bool,uint8)[], (uint128,uint128)[] vaultData)
func (_MStable *MStableCaller) GetBassets(opts *bind.CallOpts) ([]BassetPersonal, []BassetData, error) {
	var out []interface{}
	err := _MStable.contract.Call(opts, &out, "getBassets")

	if err != nil {
		return *new([]BassetPersonal), *new([]BassetData), err
	}

	out0 := *abi.ConvertType(out[0], new([]BassetPersonal)).(*[]BassetPersonal)
	out1 := *abi.ConvertType(out[1], new([]BassetData)).(*[]BassetData)

	return out0, out1, err

}

// GetBassets is a free data retrieval call binding the contract method 0x1d3ce398.
//
// Solidity: function getBassets() view returns((address,address,bool,uint8)[], (uint128,uint128)[] vaultData)
func (_MStable *MStableSession) GetBassets() ([]BassetPersonal, []BassetData, error) {
	return _MStable.Contract.GetBassets(&_MStable.CallOpts)
}

// GetBassets is a free data retrieval call binding the contract method 0x1d3ce398.
//
// Solidity: function getBassets() view returns((address,address,bool,uint8)[], (uint128,uint128)[] vaultData)
func (_MStable *MStableCallerSession) GetBassets() ([]BassetPersonal, []BassetData, error) {
	return _MStable.Contract.GetBassets(&_MStable.CallOpts)
}

// GetConfig is a free data retrieval call binding the contract method 0xc3f909d4.
//
// Solidity: function getConfig() view returns((uint256,uint256,(uint128,uint128)) config)
func (_MStable *MStableCaller) GetConfig(opts *bind.CallOpts) (FeederConfig, error) {
	var out []interface{}
	err := _MStable.contract.Call(opts, &out, "getConfig")

	if err != nil {
		return *new(FeederConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(FeederConfig)).(*FeederConfig)

	return out0, err

}

// GetConfig is a free data retrieval call binding the contract method 0xc3f909d4.
//
// Solidity: function getConfig() view returns((uint256,uint256,(uint128,uint128)) config)
func (_MStable *MStableSession) GetConfig() (FeederConfig, error) {
	return _MStable.Contract.GetConfig(&_MStable.CallOpts)
}

// GetConfig is a free data retrieval call binding the contract method 0xc3f909d4.
//
// Solidity: function getConfig() view returns((uint256,uint256,(uint128,uint128)) config)
func (_MStable *MStableCallerSession) GetConfig() (FeederConfig, error) {
	return _MStable.Contract.GetConfig(&_MStable.CallOpts)
}

// GetMintMultiOutput is a free data retrieval call binding the contract method 0xe5555b8f.
//
// Solidity: function getMintMultiOutput(address[] _inputs, uint256[] _inputQuantities) view returns(uint256 mintOutput)
func (_MStable *MStableCaller) GetMintMultiOutput(opts *bind.CallOpts, _inputs []common.Address, _inputQuantities []*big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MStable.contract.Call(opts, &out, "getMintMultiOutput", _inputs, _inputQuantities)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMintMultiOutput is a free data retrieval call binding the contract method 0xe5555b8f.
//
// Solidity: function getMintMultiOutput(address[] _inputs, uint256[] _inputQuantities) view returns(uint256 mintOutput)
func (_MStable *MStableSession) GetMintMultiOutput(_inputs []common.Address, _inputQuantities []*big.Int) (*big.Int, error) {
	return _MStable.Contract.GetMintMultiOutput(&_MStable.CallOpts, _inputs, _inputQuantities)
}

// GetMintMultiOutput is a free data retrieval call binding the contract method 0xe5555b8f.
//
// Solidity: function getMintMultiOutput(address[] _inputs, uint256[] _inputQuantities) view returns(uint256 mintOutput)
func (_MStable *MStableCallerSession) GetMintMultiOutput(_inputs []common.Address, _inputQuantities []*big.Int) (*big.Int, error) {
	return _MStable.Contract.GetMintMultiOutput(&_MStable.CallOpts, _inputs, _inputQuantities)
}

// GetMintOutput is a free data retrieval call binding the contract method 0x119849cf.
//
// Solidity: function getMintOutput(address _input, uint256 _inputQuantity) view returns(uint256 mintOutput)
func (_MStable *MStableCaller) GetMintOutput(opts *bind.CallOpts, _input common.Address, _inputQuantity *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MStable.contract.Call(opts, &out, "getMintOutput", _input, _inputQuantity)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMintOutput is a free data retrieval call binding the contract method 0x119849cf.
//
// Solidity: function getMintOutput(address _input, uint256 _inputQuantity) view returns(uint256 mintOutput)
func (_MStable *MStableSession) GetMintOutput(_input common.Address, _inputQuantity *big.Int) (*big.Int, error) {
	return _MStable.Contract.GetMintOutput(&_MStable.CallOpts, _input, _inputQuantity)
}

// GetMintOutput is a free data retrieval call binding the contract method 0x119849cf.
//
// Solidity: function getMintOutput(address _input, uint256 _inputQuantity) view returns(uint256 mintOutput)
func (_MStable *MStableCallerSession) GetMintOutput(_input common.Address, _inputQuantity *big.Int) (*big.Int, error) {
	return _MStable.Contract.GetMintOutput(&_MStable.CallOpts, _input, _inputQuantity)
}

// GetPrice is a free data retrieval call binding the contract method 0x98d5fdca.
//
// Solidity: function getPrice() view returns(uint256 price, uint256 k)
func (_MStable *MStableCaller) GetPrice(opts *bind.CallOpts) (struct {
	Price *big.Int
	K     *big.Int
}, error) {
	var out []interface{}
	err := _MStable.contract.Call(opts, &out, "getPrice")

	outstruct := new(struct {
		Price *big.Int
		K     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Price = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.K = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetPrice is a free data retrieval call binding the contract method 0x98d5fdca.
//
// Solidity: function getPrice() view returns(uint256 price, uint256 k)
func (_MStable *MStableSession) GetPrice() (struct {
	Price *big.Int
	K     *big.Int
}, error) {
	return _MStable.Contract.GetPrice(&_MStable.CallOpts)
}

// GetPrice is a free data retrieval call binding the contract method 0x98d5fdca.
//
// Solidity: function getPrice() view returns(uint256 price, uint256 k)
func (_MStable *MStableCallerSession) GetPrice() (struct {
	Price *big.Int
	K     *big.Int
}, error) {
	return _MStable.Contract.GetPrice(&_MStable.CallOpts)
}

// GetRedeemExactBassetsOutput is a free data retrieval call binding the contract method 0x04de5a73.
//
// Solidity: function getRedeemExactBassetsOutput(address[] _outputs, uint256[] _outputQuantities) view returns(uint256 fpTokenQuantity)
func (_MStable *MStableCaller) GetRedeemExactBassetsOutput(opts *bind.CallOpts, _outputs []common.Address, _outputQuantities []*big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MStable.contract.Call(opts, &out, "getRedeemExactBassetsOutput", _outputs, _outputQuantities)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRedeemExactBassetsOutput is a free data retrieval call binding the contract method 0x04de5a73.
//
// Solidity: function getRedeemExactBassetsOutput(address[] _outputs, uint256[] _outputQuantities) view returns(uint256 fpTokenQuantity)
func (_MStable *MStableSession) GetRedeemExactBassetsOutput(_outputs []common.Address, _outputQuantities []*big.Int) (*big.Int, error) {
	return _MStable.Contract.GetRedeemExactBassetsOutput(&_MStable.CallOpts, _outputs, _outputQuantities)
}

// GetRedeemExactBassetsOutput is a free data retrieval call binding the contract method 0x04de5a73.
//
// Solidity: function getRedeemExactBassetsOutput(address[] _outputs, uint256[] _outputQuantities) view returns(uint256 fpTokenQuantity)
func (_MStable *MStableCallerSession) GetRedeemExactBassetsOutput(_outputs []common.Address, _outputQuantities []*big.Int) (*big.Int, error) {
	return _MStable.Contract.GetRedeemExactBassetsOutput(&_MStable.CallOpts, _outputs, _outputQuantities)
}

// GetRedeemOutput is a free data retrieval call binding the contract method 0x78aa987e.
//
// Solidity: function getRedeemOutput(address _output, uint256 _fpTokenQuantity) view returns(uint256 bAssetOutput)
func (_MStable *MStableCaller) GetRedeemOutput(opts *bind.CallOpts, _output common.Address, _fpTokenQuantity *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MStable.contract.Call(opts, &out, "getRedeemOutput", _output, _fpTokenQuantity)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRedeemOutput is a free data retrieval call binding the contract method 0x78aa987e.
//
// Solidity: function getRedeemOutput(address _output, uint256 _fpTokenQuantity) view returns(uint256 bAssetOutput)
func (_MStable *MStableSession) GetRedeemOutput(_output common.Address, _fpTokenQuantity *big.Int) (*big.Int, error) {
	return _MStable.Contract.GetRedeemOutput(&_MStable.CallOpts, _output, _fpTokenQuantity)
}

// GetRedeemOutput is a free data retrieval call binding the contract method 0x78aa987e.
//
// Solidity: function getRedeemOutput(address _output, uint256 _fpTokenQuantity) view returns(uint256 bAssetOutput)
func (_MStable *MStableCallerSession) GetRedeemOutput(_output common.Address, _fpTokenQuantity *big.Int) (*big.Int, error) {
	return _MStable.Contract.GetRedeemOutput(&_MStable.CallOpts, _output, _fpTokenQuantity)
}

// GetSwapOutput is a free data retrieval call binding the contract method 0x72ea9076.
//
// Solidity: function getSwapOutput(address _input, address _output, uint256 _inputQuantity) view returns(uint256 swapOutput)
func (_MStable *MStableCaller) GetSwapOutput(opts *bind.CallOpts, _input common.Address, _output common.Address, _inputQuantity *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MStable.contract.Call(opts, &out, "getSwapOutput", _input, _output, _inputQuantity)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSwapOutput is a free data retrieval call binding the contract method 0x72ea9076.
//
// Solidity: function getSwapOutput(address _input, address _output, uint256 _inputQuantity) view returns(uint256 swapOutput)
func (_MStable *MStableSession) GetSwapOutput(_input common.Address, _output common.Address, _inputQuantity *big.Int) (*big.Int, error) {
	return _MStable.Contract.GetSwapOutput(&_MStable.CallOpts, _input, _output, _inputQuantity)
}

// GetSwapOutput is a free data retrieval call binding the contract method 0x72ea9076.
//
// Solidity: function getSwapOutput(address _input, address _output, uint256 _inputQuantity) view returns(uint256 swapOutput)
func (_MStable *MStableCallerSession) GetSwapOutput(_input common.Address, _output common.Address, _inputQuantity *big.Int) (*big.Int, error) {
	return _MStable.Contract.GetSwapOutput(&_MStable.CallOpts, _input, _output, _inputQuantity)
}

// MAsset is a free data retrieval call binding the contract method 0x178d341f.
//
// Solidity: function mAsset() view returns(address)
func (_MStable *MStableCaller) MAsset(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MStable.contract.Call(opts, &out, "mAsset")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MAsset is a free data retrieval call binding the contract method 0x178d341f.
//
// Solidity: function mAsset() view returns(address)
func (_MStable *MStableSession) MAsset() (common.Address, error) {
	return _MStable.Contract.MAsset(&_MStable.CallOpts)
}

// MAsset is a free data retrieval call binding the contract method 0x178d341f.
//
// Solidity: function mAsset() view returns(address)
func (_MStable *MStableCallerSession) MAsset() (common.Address, error) {
	return _MStable.Contract.MAsset(&_MStable.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MStable *MStableCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MStable.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MStable *MStableSession) Name() (string, error) {
	return _MStable.Contract.Name(&_MStable.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MStable *MStableCallerSession) Name() (string, error) {
	return _MStable.Contract.Name(&_MStable.CallOpts)
}

// Nexus is a free data retrieval call binding the contract method 0xa3f5c1d2.
//
// Solidity: function nexus() view returns(address)
func (_MStable *MStableCaller) Nexus(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MStable.contract.Call(opts, &out, "nexus")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Nexus is a free data retrieval call binding the contract method 0xa3f5c1d2.
//
// Solidity: function nexus() view returns(address)
func (_MStable *MStableSession) Nexus() (common.Address, error) {
	return _MStable.Contract.Nexus(&_MStable.CallOpts)
}

// Nexus is a free data retrieval call binding the contract method 0xa3f5c1d2.
//
// Solidity: function nexus() view returns(address)
func (_MStable *MStableCallerSession) Nexus() (common.Address, error) {
	return _MStable.Contract.Nexus(&_MStable.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_MStable *MStableCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _MStable.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_MStable *MStableSession) Paused() (bool, error) {
	return _MStable.Contract.Paused(&_MStable.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_MStable *MStableCallerSession) Paused() (bool, error) {
	return _MStable.Contract.Paused(&_MStable.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MStable *MStableCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MStable.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MStable *MStableSession) Symbol() (string, error) {
	return _MStable.Contract.Symbol(&_MStable.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MStable *MStableCallerSession) Symbol() (string, error) {
	return _MStable.Contract.Symbol(&_MStable.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_MStable *MStableCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MStable.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_MStable *MStableSession) TotalSupply() (*big.Int, error) {
	return _MStable.Contract.TotalSupply(&_MStable.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_MStable *MStableCallerSession) TotalSupply() (*big.Int, error) {
	return _MStable.Contract.TotalSupply(&_MStable.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_MStable *MStableTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MStable.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_MStable *MStableSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MStable.Contract.Approve(&_MStable.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_MStable *MStableTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MStable.Contract.Approve(&_MStable.TransactOpts, spender, amount)
}

// CollectPendingFees is a paid mutator transaction binding the contract method 0x98fec3af.
//
// Solidity: function collectPendingFees() returns()
func (_MStable *MStableTransactor) CollectPendingFees(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MStable.contract.Transact(opts, "collectPendingFees")
}

// CollectPendingFees is a paid mutator transaction binding the contract method 0x98fec3af.
//
// Solidity: function collectPendingFees() returns()
func (_MStable *MStableSession) CollectPendingFees() (*types.Transaction, error) {
	return _MStable.Contract.CollectPendingFees(&_MStable.TransactOpts)
}

// CollectPendingFees is a paid mutator transaction binding the contract method 0x98fec3af.
//
// Solidity: function collectPendingFees() returns()
func (_MStable *MStableTransactorSession) CollectPendingFees() (*types.Transaction, error) {
	return _MStable.Contract.CollectPendingFees(&_MStable.TransactOpts)
}

// CollectPlatformInterest is a paid mutator transaction binding the contract method 0x7e8901ea.
//
// Solidity: function collectPlatformInterest() returns(uint256 mintAmount, uint256 newSupply)
func (_MStable *MStableTransactor) CollectPlatformInterest(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MStable.contract.Transact(opts, "collectPlatformInterest")
}

// CollectPlatformInterest is a paid mutator transaction binding the contract method 0x7e8901ea.
//
// Solidity: function collectPlatformInterest() returns(uint256 mintAmount, uint256 newSupply)
func (_MStable *MStableSession) CollectPlatformInterest() (*types.Transaction, error) {
	return _MStable.Contract.CollectPlatformInterest(&_MStable.TransactOpts)
}

// CollectPlatformInterest is a paid mutator transaction binding the contract method 0x7e8901ea.
//
// Solidity: function collectPlatformInterest() returns(uint256 mintAmount, uint256 newSupply)
func (_MStable *MStableTransactorSession) CollectPlatformInterest() (*types.Transaction, error) {
	return _MStable.Contract.CollectPlatformInterest(&_MStable.TransactOpts)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_MStable *MStableTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _MStable.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_MStable *MStableSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _MStable.Contract.DecreaseAllowance(&_MStable.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_MStable *MStableTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _MStable.Contract.DecreaseAllowance(&_MStable.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_MStable *MStableTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _MStable.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_MStable *MStableSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _MStable.Contract.IncreaseAllowance(&_MStable.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_MStable *MStableTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _MStable.Contract.IncreaseAllowance(&_MStable.TransactOpts, spender, addedValue)
}

// Initialize is a paid mutator transaction binding the contract method 0x23fb4bd3.
//
// Solidity: function initialize(string _nameArg, string _symbolArg, (address,address,bool,uint8) _mAsset, (address,address,bool,uint8) _fAsset, address[] _mpAssets, (uint256,(uint128,uint128)) _config) returns()
func (_MStable *MStableTransactor) Initialize(opts *bind.TransactOpts, _nameArg string, _symbolArg string, _mAsset BassetPersonal, _fAsset BassetPersonal, _mpAssets []common.Address, _config InvariantConfig) (*types.Transaction, error) {
	return _MStable.contract.Transact(opts, "initialize", _nameArg, _symbolArg, _mAsset, _fAsset, _mpAssets, _config)
}

// Initialize is a paid mutator transaction binding the contract method 0x23fb4bd3.
//
// Solidity: function initialize(string _nameArg, string _symbolArg, (address,address,bool,uint8) _mAsset, (address,address,bool,uint8) _fAsset, address[] _mpAssets, (uint256,(uint128,uint128)) _config) returns()
func (_MStable *MStableSession) Initialize(_nameArg string, _symbolArg string, _mAsset BassetPersonal, _fAsset BassetPersonal, _mpAssets []common.Address, _config InvariantConfig) (*types.Transaction, error) {
	return _MStable.Contract.Initialize(&_MStable.TransactOpts, _nameArg, _symbolArg, _mAsset, _fAsset, _mpAssets, _config)
}

// Initialize is a paid mutator transaction binding the contract method 0x23fb4bd3.
//
// Solidity: function initialize(string _nameArg, string _symbolArg, (address,address,bool,uint8) _mAsset, (address,address,bool,uint8) _fAsset, address[] _mpAssets, (uint256,(uint128,uint128)) _config) returns()
func (_MStable *MStableTransactorSession) Initialize(_nameArg string, _symbolArg string, _mAsset BassetPersonal, _fAsset BassetPersonal, _mpAssets []common.Address, _config InvariantConfig) (*types.Transaction, error) {
	return _MStable.Contract.Initialize(&_MStable.TransactOpts, _nameArg, _symbolArg, _mAsset, _fAsset, _mpAssets, _config)
}

// MigrateBassets is a paid mutator transaction binding the contract method 0x6fb3e89c.
//
// Solidity: function migrateBassets(address[] _bAssets, address _newIntegration) returns()
func (_MStable *MStableTransactor) MigrateBassets(opts *bind.TransactOpts, _bAssets []common.Address, _newIntegration common.Address) (*types.Transaction, error) {
	return _MStable.contract.Transact(opts, "migrateBassets", _bAssets, _newIntegration)
}

// MigrateBassets is a paid mutator transaction binding the contract method 0x6fb3e89c.
//
// Solidity: function migrateBassets(address[] _bAssets, address _newIntegration) returns()
func (_MStable *MStableSession) MigrateBassets(_bAssets []common.Address, _newIntegration common.Address) (*types.Transaction, error) {
	return _MStable.Contract.MigrateBassets(&_MStable.TransactOpts, _bAssets, _newIntegration)
}

// MigrateBassets is a paid mutator transaction binding the contract method 0x6fb3e89c.
//
// Solidity: function migrateBassets(address[] _bAssets, address _newIntegration) returns()
func (_MStable *MStableTransactorSession) MigrateBassets(_bAssets []common.Address, _newIntegration common.Address) (*types.Transaction, error) {
	return _MStable.Contract.MigrateBassets(&_MStable.TransactOpts, _bAssets, _newIntegration)
}

// Mint is a paid mutator transaction binding the contract method 0xf74bfe8e.
//
// Solidity: function mint(address _input, uint256 _inputQuantity, uint256 _minOutputQuantity, address _recipient) returns(uint256 mintOutput)
func (_MStable *MStableTransactor) Mint(opts *bind.TransactOpts, _input common.Address, _inputQuantity *big.Int, _minOutputQuantity *big.Int, _recipient common.Address) (*types.Transaction, error) {
	return _MStable.contract.Transact(opts, "mint", _input, _inputQuantity, _minOutputQuantity, _recipient)
}

// Mint is a paid mutator transaction binding the contract method 0xf74bfe8e.
//
// Solidity: function mint(address _input, uint256 _inputQuantity, uint256 _minOutputQuantity, address _recipient) returns(uint256 mintOutput)
func (_MStable *MStableSession) Mint(_input common.Address, _inputQuantity *big.Int, _minOutputQuantity *big.Int, _recipient common.Address) (*types.Transaction, error) {
	return _MStable.Contract.Mint(&_MStable.TransactOpts, _input, _inputQuantity, _minOutputQuantity, _recipient)
}

// Mint is a paid mutator transaction binding the contract method 0xf74bfe8e.
//
// Solidity: function mint(address _input, uint256 _inputQuantity, uint256 _minOutputQuantity, address _recipient) returns(uint256 mintOutput)
func (_MStable *MStableTransactorSession) Mint(_input common.Address, _inputQuantity *big.Int, _minOutputQuantity *big.Int, _recipient common.Address) (*types.Transaction, error) {
	return _MStable.Contract.Mint(&_MStable.TransactOpts, _input, _inputQuantity, _minOutputQuantity, _recipient)
}

// MintMulti is a paid mutator transaction binding the contract method 0xaf290bd4.
//
// Solidity: function mintMulti(address[] _inputs, uint256[] _inputQuantities, uint256 _minOutputQuantity, address _recipient) returns(uint256 mintOutput)
func (_MStable *MStableTransactor) MintMulti(opts *bind.TransactOpts, _inputs []common.Address, _inputQuantities []*big.Int, _minOutputQuantity *big.Int, _recipient common.Address) (*types.Transaction, error) {
	return _MStable.contract.Transact(opts, "mintMulti", _inputs, _inputQuantities, _minOutputQuantity, _recipient)
}

// MintMulti is a paid mutator transaction binding the contract method 0xaf290bd4.
//
// Solidity: function mintMulti(address[] _inputs, uint256[] _inputQuantities, uint256 _minOutputQuantity, address _recipient) returns(uint256 mintOutput)
func (_MStable *MStableSession) MintMulti(_inputs []common.Address, _inputQuantities []*big.Int, _minOutputQuantity *big.Int, _recipient common.Address) (*types.Transaction, error) {
	return _MStable.Contract.MintMulti(&_MStable.TransactOpts, _inputs, _inputQuantities, _minOutputQuantity, _recipient)
}

// MintMulti is a paid mutator transaction binding the contract method 0xaf290bd4.
//
// Solidity: function mintMulti(address[] _inputs, uint256[] _inputQuantities, uint256 _minOutputQuantity, address _recipient) returns(uint256 mintOutput)
func (_MStable *MStableTransactorSession) MintMulti(_inputs []common.Address, _inputQuantities []*big.Int, _minOutputQuantity *big.Int, _recipient common.Address) (*types.Transaction, error) {
	return _MStable.Contract.MintMulti(&_MStable.TransactOpts, _inputs, _inputQuantities, _minOutputQuantity, _recipient)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_MStable *MStableTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MStable.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_MStable *MStableSession) Pause() (*types.Transaction, error) {
	return _MStable.Contract.Pause(&_MStable.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_MStable *MStableTransactorSession) Pause() (*types.Transaction, error) {
	return _MStable.Contract.Pause(&_MStable.TransactOpts)
}

// Redeem is a paid mutator transaction binding the contract method 0x43bcfab6.
//
// Solidity: function redeem(address _output, uint256 _fpTokenQuantity, uint256 _minOutputQuantity, address _recipient) returns(uint256 outputQuantity)
func (_MStable *MStableTransactor) Redeem(opts *bind.TransactOpts, _output common.Address, _fpTokenQuantity *big.Int, _minOutputQuantity *big.Int, _recipient common.Address) (*types.Transaction, error) {
	return _MStable.contract.Transact(opts, "redeem", _output, _fpTokenQuantity, _minOutputQuantity, _recipient)
}

// Redeem is a paid mutator transaction binding the contract method 0x43bcfab6.
//
// Solidity: function redeem(address _output, uint256 _fpTokenQuantity, uint256 _minOutputQuantity, address _recipient) returns(uint256 outputQuantity)
func (_MStable *MStableSession) Redeem(_output common.Address, _fpTokenQuantity *big.Int, _minOutputQuantity *big.Int, _recipient common.Address) (*types.Transaction, error) {
	return _MStable.Contract.Redeem(&_MStable.TransactOpts, _output, _fpTokenQuantity, _minOutputQuantity, _recipient)
}

// Redeem is a paid mutator transaction binding the contract method 0x43bcfab6.
//
// Solidity: function redeem(address _output, uint256 _fpTokenQuantity, uint256 _minOutputQuantity, address _recipient) returns(uint256 outputQuantity)
func (_MStable *MStableTransactorSession) Redeem(_output common.Address, _fpTokenQuantity *big.Int, _minOutputQuantity *big.Int, _recipient common.Address) (*types.Transaction, error) {
	return _MStable.Contract.Redeem(&_MStable.TransactOpts, _output, _fpTokenQuantity, _minOutputQuantity, _recipient)
}

// RedeemExactBassets is a paid mutator transaction binding the contract method 0xaa5d27e9.
//
// Solidity: function redeemExactBassets(address[] _outputs, uint256[] _outputQuantities, uint256 _maxInputQuantity, address _recipient) returns(uint256 fpTokenQuantity)
func (_MStable *MStableTransactor) RedeemExactBassets(opts *bind.TransactOpts, _outputs []common.Address, _outputQuantities []*big.Int, _maxInputQuantity *big.Int, _recipient common.Address) (*types.Transaction, error) {
	return _MStable.contract.Transact(opts, "redeemExactBassets", _outputs, _outputQuantities, _maxInputQuantity, _recipient)
}

// RedeemExactBassets is a paid mutator transaction binding the contract method 0xaa5d27e9.
//
// Solidity: function redeemExactBassets(address[] _outputs, uint256[] _outputQuantities, uint256 _maxInputQuantity, address _recipient) returns(uint256 fpTokenQuantity)
func (_MStable *MStableSession) RedeemExactBassets(_outputs []common.Address, _outputQuantities []*big.Int, _maxInputQuantity *big.Int, _recipient common.Address) (*types.Transaction, error) {
	return _MStable.Contract.RedeemExactBassets(&_MStable.TransactOpts, _outputs, _outputQuantities, _maxInputQuantity, _recipient)
}

// RedeemExactBassets is a paid mutator transaction binding the contract method 0xaa5d27e9.
//
// Solidity: function redeemExactBassets(address[] _outputs, uint256[] _outputQuantities, uint256 _maxInputQuantity, address _recipient) returns(uint256 fpTokenQuantity)
func (_MStable *MStableTransactorSession) RedeemExactBassets(_outputs []common.Address, _outputQuantities []*big.Int, _maxInputQuantity *big.Int, _recipient common.Address) (*types.Transaction, error) {
	return _MStable.Contract.RedeemExactBassets(&_MStable.TransactOpts, _outputs, _outputQuantities, _maxInputQuantity, _recipient)
}

// RedeemProportionately is a paid mutator transaction binding the contract method 0xc1db77ec.
//
// Solidity: function redeemProportionately(uint256 _inputQuantity, uint256[] _minOutputQuantities, address _recipient) returns(uint256[] outputQuantities)
func (_MStable *MStableTransactor) RedeemProportionately(opts *bind.TransactOpts, _inputQuantity *big.Int, _minOutputQuantities []*big.Int, _recipient common.Address) (*types.Transaction, error) {
	return _MStable.contract.Transact(opts, "redeemProportionately", _inputQuantity, _minOutputQuantities, _recipient)
}

// RedeemProportionately is a paid mutator transaction binding the contract method 0xc1db77ec.
//
// Solidity: function redeemProportionately(uint256 _inputQuantity, uint256[] _minOutputQuantities, address _recipient) returns(uint256[] outputQuantities)
func (_MStable *MStableSession) RedeemProportionately(_inputQuantity *big.Int, _minOutputQuantities []*big.Int, _recipient common.Address) (*types.Transaction, error) {
	return _MStable.Contract.RedeemProportionately(&_MStable.TransactOpts, _inputQuantity, _minOutputQuantities, _recipient)
}

// RedeemProportionately is a paid mutator transaction binding the contract method 0xc1db77ec.
//
// Solidity: function redeemProportionately(uint256 _inputQuantity, uint256[] _minOutputQuantities, address _recipient) returns(uint256[] outputQuantities)
func (_MStable *MStableTransactorSession) RedeemProportionately(_inputQuantity *big.Int, _minOutputQuantities []*big.Int, _recipient common.Address) (*types.Transaction, error) {
	return _MStable.Contract.RedeemProportionately(&_MStable.TransactOpts, _inputQuantity, _minOutputQuantities, _recipient)
}

// SetCacheSize is a paid mutator transaction binding the contract method 0x1820783d.
//
// Solidity: function setCacheSize(uint256 _cacheSize) returns()
func (_MStable *MStableTransactor) SetCacheSize(opts *bind.TransactOpts, _cacheSize *big.Int) (*types.Transaction, error) {
	return _MStable.contract.Transact(opts, "setCacheSize", _cacheSize)
}

// SetCacheSize is a paid mutator transaction binding the contract method 0x1820783d.
//
// Solidity: function setCacheSize(uint256 _cacheSize) returns()
func (_MStable *MStableSession) SetCacheSize(_cacheSize *big.Int) (*types.Transaction, error) {
	return _MStable.Contract.SetCacheSize(&_MStable.TransactOpts, _cacheSize)
}

// SetCacheSize is a paid mutator transaction binding the contract method 0x1820783d.
//
// Solidity: function setCacheSize(uint256 _cacheSize) returns()
func (_MStable *MStableTransactorSession) SetCacheSize(_cacheSize *big.Int) (*types.Transaction, error) {
	return _MStable.Contract.SetCacheSize(&_MStable.TransactOpts, _cacheSize)
}

// SetFees is a paid mutator transaction binding the contract method 0xcec10c11.
//
// Solidity: function setFees(uint256 _swapFee, uint256 _redemptionFee, uint256 _govFee) returns()
func (_MStable *MStableTransactor) SetFees(opts *bind.TransactOpts, _swapFee *big.Int, _redemptionFee *big.Int, _govFee *big.Int) (*types.Transaction, error) {
	return _MStable.contract.Transact(opts, "setFees", _swapFee, _redemptionFee, _govFee)
}

// SetFees is a paid mutator transaction binding the contract method 0xcec10c11.
//
// Solidity: function setFees(uint256 _swapFee, uint256 _redemptionFee, uint256 _govFee) returns()
func (_MStable *MStableSession) SetFees(_swapFee *big.Int, _redemptionFee *big.Int, _govFee *big.Int) (*types.Transaction, error) {
	return _MStable.Contract.SetFees(&_MStable.TransactOpts, _swapFee, _redemptionFee, _govFee)
}

// SetFees is a paid mutator transaction binding the contract method 0xcec10c11.
//
// Solidity: function setFees(uint256 _swapFee, uint256 _redemptionFee, uint256 _govFee) returns()
func (_MStable *MStableTransactorSession) SetFees(_swapFee *big.Int, _redemptionFee *big.Int, _govFee *big.Int) (*types.Transaction, error) {
	return _MStable.Contract.SetFees(&_MStable.TransactOpts, _swapFee, _redemptionFee, _govFee)
}

// SetWeightLimits is a paid mutator transaction binding the contract method 0xa1c25151.
//
// Solidity: function setWeightLimits(uint128 _min, uint128 _max) returns()
func (_MStable *MStableTransactor) SetWeightLimits(opts *bind.TransactOpts, _min *big.Int, _max *big.Int) (*types.Transaction, error) {
	return _MStable.contract.Transact(opts, "setWeightLimits", _min, _max)
}

// SetWeightLimits is a paid mutator transaction binding the contract method 0xa1c25151.
//
// Solidity: function setWeightLimits(uint128 _min, uint128 _max) returns()
func (_MStable *MStableSession) SetWeightLimits(_min *big.Int, _max *big.Int) (*types.Transaction, error) {
	return _MStable.Contract.SetWeightLimits(&_MStable.TransactOpts, _min, _max)
}

// SetWeightLimits is a paid mutator transaction binding the contract method 0xa1c25151.
//
// Solidity: function setWeightLimits(uint128 _min, uint128 _max) returns()
func (_MStable *MStableTransactorSession) SetWeightLimits(_min *big.Int, _max *big.Int) (*types.Transaction, error) {
	return _MStable.Contract.SetWeightLimits(&_MStable.TransactOpts, _min, _max)
}

// StartRampA is a paid mutator transaction binding the contract method 0x44e3fa3c.
//
// Solidity: function startRampA(uint256 _targetA, uint256 _rampEndTime) returns()
func (_MStable *MStableTransactor) StartRampA(opts *bind.TransactOpts, _targetA *big.Int, _rampEndTime *big.Int) (*types.Transaction, error) {
	return _MStable.contract.Transact(opts, "startRampA", _targetA, _rampEndTime)
}

// StartRampA is a paid mutator transaction binding the contract method 0x44e3fa3c.
//
// Solidity: function startRampA(uint256 _targetA, uint256 _rampEndTime) returns()
func (_MStable *MStableSession) StartRampA(_targetA *big.Int, _rampEndTime *big.Int) (*types.Transaction, error) {
	return _MStable.Contract.StartRampA(&_MStable.TransactOpts, _targetA, _rampEndTime)
}

// StartRampA is a paid mutator transaction binding the contract method 0x44e3fa3c.
//
// Solidity: function startRampA(uint256 _targetA, uint256 _rampEndTime) returns()
func (_MStable *MStableTransactorSession) StartRampA(_targetA *big.Int, _rampEndTime *big.Int) (*types.Transaction, error) {
	return _MStable.Contract.StartRampA(&_MStable.TransactOpts, _targetA, _rampEndTime)
}

// StopRampA is a paid mutator transaction binding the contract method 0xc4db7fa0.
//
// Solidity: function stopRampA() returns()
func (_MStable *MStableTransactor) StopRampA(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MStable.contract.Transact(opts, "stopRampA")
}

// StopRampA is a paid mutator transaction binding the contract method 0xc4db7fa0.
//
// Solidity: function stopRampA() returns()
func (_MStable *MStableSession) StopRampA() (*types.Transaction, error) {
	return _MStable.Contract.StopRampA(&_MStable.TransactOpts)
}

// StopRampA is a paid mutator transaction binding the contract method 0xc4db7fa0.
//
// Solidity: function stopRampA() returns()
func (_MStable *MStableTransactorSession) StopRampA() (*types.Transaction, error) {
	return _MStable.Contract.StopRampA(&_MStable.TransactOpts)
}

// Swap is a paid mutator transaction binding the contract method 0xd5bcb9b5.
//
// Solidity: function swap(address _input, address _output, uint256 _inputQuantity, uint256 _minOutputQuantity, address _recipient) returns(uint256 swapOutput)
func (_MStable *MStableTransactor) Swap(opts *bind.TransactOpts, _input common.Address, _output common.Address, _inputQuantity *big.Int, _minOutputQuantity *big.Int, _recipient common.Address) (*types.Transaction, error) {
	return _MStable.contract.Transact(opts, "swap", _input, _output, _inputQuantity, _minOutputQuantity, _recipient)
}

// Swap is a paid mutator transaction binding the contract method 0xd5bcb9b5.
//
// Solidity: function swap(address _input, address _output, uint256 _inputQuantity, uint256 _minOutputQuantity, address _recipient) returns(uint256 swapOutput)
func (_MStable *MStableSession) Swap(_input common.Address, _output common.Address, _inputQuantity *big.Int, _minOutputQuantity *big.Int, _recipient common.Address) (*types.Transaction, error) {
	return _MStable.Contract.Swap(&_MStable.TransactOpts, _input, _output, _inputQuantity, _minOutputQuantity, _recipient)
}

// Swap is a paid mutator transaction binding the contract method 0xd5bcb9b5.
//
// Solidity: function swap(address _input, address _output, uint256 _inputQuantity, uint256 _minOutputQuantity, address _recipient) returns(uint256 swapOutput)
func (_MStable *MStableTransactorSession) Swap(_input common.Address, _output common.Address, _inputQuantity *big.Int, _minOutputQuantity *big.Int, _recipient common.Address) (*types.Transaction, error) {
	return _MStable.Contract.Swap(&_MStable.TransactOpts, _input, _output, _inputQuantity, _minOutputQuantity, _recipient)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_MStable *MStableTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MStable.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_MStable *MStableSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MStable.Contract.Transfer(&_MStable.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_MStable *MStableTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MStable.Contract.Transfer(&_MStable.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_MStable *MStableTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MStable.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_MStable *MStableSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MStable.Contract.TransferFrom(&_MStable.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_MStable *MStableTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MStable.Contract.TransferFrom(&_MStable.TransactOpts, sender, recipient, amount)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_MStable *MStableTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MStable.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_MStable *MStableSession) Unpause() (*types.Transaction, error) {
	return _MStable.Contract.Unpause(&_MStable.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_MStable *MStableTransactorSession) Unpause() (*types.Transaction, error) {
	return _MStable.Contract.Unpause(&_MStable.TransactOpts)
}

// MStableApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the MStable contract.
type MStableApprovalIterator struct {
	Event *MStableApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MStableApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MStableApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MStableApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MStableApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MStableApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MStableApproval represents a Approval event raised by the MStable contract.
type MStableApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_MStable *MStableFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*MStableApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _MStable.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &MStableApprovalIterator{contract: _MStable.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_MStable *MStableFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *MStableApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _MStable.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MStableApproval)
				if err := _MStable.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_MStable *MStableFilterer) ParseApproval(log types.Log) (*MStableApproval, error) {
	event := new(MStableApproval)
	if err := _MStable.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MStableBassetsMigratedIterator is returned from FilterBassetsMigrated and is used to iterate over the raw logs and unpacked data for BassetsMigrated events raised by the MStable contract.
type MStableBassetsMigratedIterator struct {
	Event *MStableBassetsMigrated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MStableBassetsMigratedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MStableBassetsMigrated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MStableBassetsMigrated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MStableBassetsMigratedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MStableBassetsMigratedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MStableBassetsMigrated represents a BassetsMigrated event raised by the MStable contract.
type MStableBassetsMigrated struct {
	BAssets       []common.Address
	NewIntegrator common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterBassetsMigrated is a free log retrieval operation binding the contract event 0x407ca8e939a25b34f290fb7f4d3b0d85d03f13313dc34a6a15bbd91492cfa249.
//
// Solidity: event BassetsMigrated(address[] bAssets, address newIntegrator)
func (_MStable *MStableFilterer) FilterBassetsMigrated(opts *bind.FilterOpts) (*MStableBassetsMigratedIterator, error) {

	logs, sub, err := _MStable.contract.FilterLogs(opts, "BassetsMigrated")
	if err != nil {
		return nil, err
	}
	return &MStableBassetsMigratedIterator{contract: _MStable.contract, event: "BassetsMigrated", logs: logs, sub: sub}, nil
}

// WatchBassetsMigrated is a free log subscription operation binding the contract event 0x407ca8e939a25b34f290fb7f4d3b0d85d03f13313dc34a6a15bbd91492cfa249.
//
// Solidity: event BassetsMigrated(address[] bAssets, address newIntegrator)
func (_MStable *MStableFilterer) WatchBassetsMigrated(opts *bind.WatchOpts, sink chan<- *MStableBassetsMigrated) (event.Subscription, error) {

	logs, sub, err := _MStable.contract.WatchLogs(opts, "BassetsMigrated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MStableBassetsMigrated)
				if err := _MStable.contract.UnpackLog(event, "BassetsMigrated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBassetsMigrated is a log parse operation binding the contract event 0x407ca8e939a25b34f290fb7f4d3b0d85d03f13313dc34a6a15bbd91492cfa249.
//
// Solidity: event BassetsMigrated(address[] bAssets, address newIntegrator)
func (_MStable *MStableFilterer) ParseBassetsMigrated(log types.Log) (*MStableBassetsMigrated, error) {
	event := new(MStableBassetsMigrated)
	if err := _MStable.contract.UnpackLog(event, "BassetsMigrated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MStableCacheSizeChangedIterator is returned from FilterCacheSizeChanged and is used to iterate over the raw logs and unpacked data for CacheSizeChanged events raised by the MStable contract.
type MStableCacheSizeChangedIterator struct {
	Event *MStableCacheSizeChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MStableCacheSizeChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MStableCacheSizeChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MStableCacheSizeChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MStableCacheSizeChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MStableCacheSizeChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MStableCacheSizeChanged represents a CacheSizeChanged event raised by the MStable contract.
type MStableCacheSizeChanged struct {
	CacheSize *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCacheSizeChanged is a free log retrieval operation binding the contract event 0x2f5a6b1defeafd30e7568ea5c176aa0702b0af5b00ba41fa20e58b2c72e8afe7.
//
// Solidity: event CacheSizeChanged(uint256 cacheSize)
func (_MStable *MStableFilterer) FilterCacheSizeChanged(opts *bind.FilterOpts) (*MStableCacheSizeChangedIterator, error) {

	logs, sub, err := _MStable.contract.FilterLogs(opts, "CacheSizeChanged")
	if err != nil {
		return nil, err
	}
	return &MStableCacheSizeChangedIterator{contract: _MStable.contract, event: "CacheSizeChanged", logs: logs, sub: sub}, nil
}

// WatchCacheSizeChanged is a free log subscription operation binding the contract event 0x2f5a6b1defeafd30e7568ea5c176aa0702b0af5b00ba41fa20e58b2c72e8afe7.
//
// Solidity: event CacheSizeChanged(uint256 cacheSize)
func (_MStable *MStableFilterer) WatchCacheSizeChanged(opts *bind.WatchOpts, sink chan<- *MStableCacheSizeChanged) (event.Subscription, error) {

	logs, sub, err := _MStable.contract.WatchLogs(opts, "CacheSizeChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MStableCacheSizeChanged)
				if err := _MStable.contract.UnpackLog(event, "CacheSizeChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCacheSizeChanged is a log parse operation binding the contract event 0x2f5a6b1defeafd30e7568ea5c176aa0702b0af5b00ba41fa20e58b2c72e8afe7.
//
// Solidity: event CacheSizeChanged(uint256 cacheSize)
func (_MStable *MStableFilterer) ParseCacheSizeChanged(log types.Log) (*MStableCacheSizeChanged, error) {
	event := new(MStableCacheSizeChanged)
	if err := _MStable.contract.UnpackLog(event, "CacheSizeChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MStableFeesChangedIterator is returned from FilterFeesChanged and is used to iterate over the raw logs and unpacked data for FeesChanged events raised by the MStable contract.
type MStableFeesChangedIterator struct {
	Event *MStableFeesChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MStableFeesChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MStableFeesChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MStableFeesChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MStableFeesChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MStableFeesChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MStableFeesChanged represents a FeesChanged event raised by the MStable contract.
type MStableFeesChanged struct {
	SwapFee       *big.Int
	RedemptionFee *big.Int
	GovFee        *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterFeesChanged is a free log retrieval operation binding the contract event 0xe06a46af1c04656f68e4f75cbbb23baa176651c7f99930a378ef9f1616dc2b8c.
//
// Solidity: event FeesChanged(uint256 swapFee, uint256 redemptionFee, uint256 govFee)
func (_MStable *MStableFilterer) FilterFeesChanged(opts *bind.FilterOpts) (*MStableFeesChangedIterator, error) {

	logs, sub, err := _MStable.contract.FilterLogs(opts, "FeesChanged")
	if err != nil {
		return nil, err
	}
	return &MStableFeesChangedIterator{contract: _MStable.contract, event: "FeesChanged", logs: logs, sub: sub}, nil
}

// WatchFeesChanged is a free log subscription operation binding the contract event 0xe06a46af1c04656f68e4f75cbbb23baa176651c7f99930a378ef9f1616dc2b8c.
//
// Solidity: event FeesChanged(uint256 swapFee, uint256 redemptionFee, uint256 govFee)
func (_MStable *MStableFilterer) WatchFeesChanged(opts *bind.WatchOpts, sink chan<- *MStableFeesChanged) (event.Subscription, error) {

	logs, sub, err := _MStable.contract.WatchLogs(opts, "FeesChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MStableFeesChanged)
				if err := _MStable.contract.UnpackLog(event, "FeesChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFeesChanged is a log parse operation binding the contract event 0xe06a46af1c04656f68e4f75cbbb23baa176651c7f99930a378ef9f1616dc2b8c.
//
// Solidity: event FeesChanged(uint256 swapFee, uint256 redemptionFee, uint256 govFee)
func (_MStable *MStableFilterer) ParseFeesChanged(log types.Log) (*MStableFeesChanged, error) {
	event := new(MStableFeesChanged)
	if err := _MStable.contract.UnpackLog(event, "FeesChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MStableMintedIterator is returned from FilterMinted and is used to iterate over the raw logs and unpacked data for Minted events raised by the MStable contract.
type MStableMintedIterator struct {
	Event *MStableMinted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MStableMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MStableMinted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MStableMinted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MStableMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MStableMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MStableMinted represents a Minted event raised by the MStable contract.
type MStableMinted struct {
	Minter        common.Address
	Recipient     common.Address
	Output        *big.Int
	Input         common.Address
	InputQuantity *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterMinted is a free log retrieval operation binding the contract event 0x30873c596f54a2e2e09894670d7e1a48b2433c00204f81fbedf557353c36e7c7.
//
// Solidity: event Minted(address indexed minter, address recipient, uint256 output, address input, uint256 inputQuantity)
func (_MStable *MStableFilterer) FilterMinted(opts *bind.FilterOpts, minter []common.Address) (*MStableMintedIterator, error) {

	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}

	logs, sub, err := _MStable.contract.FilterLogs(opts, "Minted", minterRule)
	if err != nil {
		return nil, err
	}
	return &MStableMintedIterator{contract: _MStable.contract, event: "Minted", logs: logs, sub: sub}, nil
}

// WatchMinted is a free log subscription operation binding the contract event 0x30873c596f54a2e2e09894670d7e1a48b2433c00204f81fbedf557353c36e7c7.
//
// Solidity: event Minted(address indexed minter, address recipient, uint256 output, address input, uint256 inputQuantity)
func (_MStable *MStableFilterer) WatchMinted(opts *bind.WatchOpts, sink chan<- *MStableMinted, minter []common.Address) (event.Subscription, error) {

	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}

	logs, sub, err := _MStable.contract.WatchLogs(opts, "Minted", minterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MStableMinted)
				if err := _MStable.contract.UnpackLog(event, "Minted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMinted is a log parse operation binding the contract event 0x30873c596f54a2e2e09894670d7e1a48b2433c00204f81fbedf557353c36e7c7.
//
// Solidity: event Minted(address indexed minter, address recipient, uint256 output, address input, uint256 inputQuantity)
func (_MStable *MStableFilterer) ParseMinted(log types.Log) (*MStableMinted, error) {
	event := new(MStableMinted)
	if err := _MStable.contract.UnpackLog(event, "Minted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MStableMintedMultiIterator is returned from FilterMintedMulti and is used to iterate over the raw logs and unpacked data for MintedMulti events raised by the MStable contract.
type MStableMintedMultiIterator struct {
	Event *MStableMintedMulti // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MStableMintedMultiIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MStableMintedMulti)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MStableMintedMulti)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MStableMintedMultiIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MStableMintedMultiIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MStableMintedMulti represents a MintedMulti event raised by the MStable contract.
type MStableMintedMulti struct {
	Minter          common.Address
	Recipient       common.Address
	Output          *big.Int
	Inputs          []common.Address
	InputQuantities []*big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterMintedMulti is a free log retrieval operation binding the contract event 0x7d3ff197e9071095bd36b627028ef523ecf46fcbf17cbde745a4b65aec88b6bc.
//
// Solidity: event MintedMulti(address indexed minter, address recipient, uint256 output, address[] inputs, uint256[] inputQuantities)
func (_MStable *MStableFilterer) FilterMintedMulti(opts *bind.FilterOpts, minter []common.Address) (*MStableMintedMultiIterator, error) {

	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}

	logs, sub, err := _MStable.contract.FilterLogs(opts, "MintedMulti", minterRule)
	if err != nil {
		return nil, err
	}
	return &MStableMintedMultiIterator{contract: _MStable.contract, event: "MintedMulti", logs: logs, sub: sub}, nil
}

// WatchMintedMulti is a free log subscription operation binding the contract event 0x7d3ff197e9071095bd36b627028ef523ecf46fcbf17cbde745a4b65aec88b6bc.
//
// Solidity: event MintedMulti(address indexed minter, address recipient, uint256 output, address[] inputs, uint256[] inputQuantities)
func (_MStable *MStableFilterer) WatchMintedMulti(opts *bind.WatchOpts, sink chan<- *MStableMintedMulti, minter []common.Address) (event.Subscription, error) {

	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}

	logs, sub, err := _MStable.contract.WatchLogs(opts, "MintedMulti", minterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MStableMintedMulti)
				if err := _MStable.contract.UnpackLog(event, "MintedMulti", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMintedMulti is a log parse operation binding the contract event 0x7d3ff197e9071095bd36b627028ef523ecf46fcbf17cbde745a4b65aec88b6bc.
//
// Solidity: event MintedMulti(address indexed minter, address recipient, uint256 output, address[] inputs, uint256[] inputQuantities)
func (_MStable *MStableFilterer) ParseMintedMulti(log types.Log) (*MStableMintedMulti, error) {
	event := new(MStableMintedMulti)
	if err := _MStable.contract.UnpackLog(event, "MintedMulti", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MStablePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the MStable contract.
type MStablePausedIterator struct {
	Event *MStablePaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MStablePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MStablePaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MStablePaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MStablePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MStablePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MStablePaused represents a Paused event raised by the MStable contract.
type MStablePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_MStable *MStableFilterer) FilterPaused(opts *bind.FilterOpts) (*MStablePausedIterator, error) {

	logs, sub, err := _MStable.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &MStablePausedIterator{contract: _MStable.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_MStable *MStableFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *MStablePaused) (event.Subscription, error) {

	logs, sub, err := _MStable.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MStablePaused)
				if err := _MStable.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_MStable *MStableFilterer) ParsePaused(log types.Log) (*MStablePaused, error) {
	event := new(MStablePaused)
	if err := _MStable.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MStableRedeemedIterator is returned from FilterRedeemed and is used to iterate over the raw logs and unpacked data for Redeemed events raised by the MStable contract.
type MStableRedeemedIterator struct {
	Event *MStableRedeemed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MStableRedeemedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MStableRedeemed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MStableRedeemed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MStableRedeemedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MStableRedeemedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MStableRedeemed represents a Redeemed event raised by the MStable contract.
type MStableRedeemed struct {
	Redeemer       common.Address
	Recipient      common.Address
	MAssetQuantity *big.Int
	Output         common.Address
	OutputQuantity *big.Int
	ScaledFee      *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRedeemed is a free log retrieval operation binding the contract event 0x105ffcfe6c5fa767ff5a53039fdd9ba80ce97196d4daa0beb1bfbfa0ed838ad8.
//
// Solidity: event Redeemed(address indexed redeemer, address recipient, uint256 mAssetQuantity, address output, uint256 outputQuantity, uint256 scaledFee)
func (_MStable *MStableFilterer) FilterRedeemed(opts *bind.FilterOpts, redeemer []common.Address) (*MStableRedeemedIterator, error) {

	var redeemerRule []interface{}
	for _, redeemerItem := range redeemer {
		redeemerRule = append(redeemerRule, redeemerItem)
	}

	logs, sub, err := _MStable.contract.FilterLogs(opts, "Redeemed", redeemerRule)
	if err != nil {
		return nil, err
	}
	return &MStableRedeemedIterator{contract: _MStable.contract, event: "Redeemed", logs: logs, sub: sub}, nil
}

// WatchRedeemed is a free log subscription operation binding the contract event 0x105ffcfe6c5fa767ff5a53039fdd9ba80ce97196d4daa0beb1bfbfa0ed838ad8.
//
// Solidity: event Redeemed(address indexed redeemer, address recipient, uint256 mAssetQuantity, address output, uint256 outputQuantity, uint256 scaledFee)
func (_MStable *MStableFilterer) WatchRedeemed(opts *bind.WatchOpts, sink chan<- *MStableRedeemed, redeemer []common.Address) (event.Subscription, error) {

	var redeemerRule []interface{}
	for _, redeemerItem := range redeemer {
		redeemerRule = append(redeemerRule, redeemerItem)
	}

	logs, sub, err := _MStable.contract.WatchLogs(opts, "Redeemed", redeemerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MStableRedeemed)
				if err := _MStable.contract.UnpackLog(event, "Redeemed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRedeemed is a log parse operation binding the contract event 0x105ffcfe6c5fa767ff5a53039fdd9ba80ce97196d4daa0beb1bfbfa0ed838ad8.
//
// Solidity: event Redeemed(address indexed redeemer, address recipient, uint256 mAssetQuantity, address output, uint256 outputQuantity, uint256 scaledFee)
func (_MStable *MStableFilterer) ParseRedeemed(log types.Log) (*MStableRedeemed, error) {
	event := new(MStableRedeemed)
	if err := _MStable.contract.UnpackLog(event, "Redeemed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MStableRedeemedMultiIterator is returned from FilterRedeemedMulti and is used to iterate over the raw logs and unpacked data for RedeemedMulti events raised by the MStable contract.
type MStableRedeemedMultiIterator struct {
	Event *MStableRedeemedMulti // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MStableRedeemedMultiIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MStableRedeemedMulti)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MStableRedeemedMulti)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MStableRedeemedMultiIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MStableRedeemedMultiIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MStableRedeemedMulti represents a RedeemedMulti event raised by the MStable contract.
type MStableRedeemedMulti struct {
	Redeemer       common.Address
	Recipient      common.Address
	MAssetQuantity *big.Int
	Outputs        []common.Address
	OutputQuantity []*big.Int
	ScaledFee      *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRedeemedMulti is a free log retrieval operation binding the contract event 0xdf5bf345d5d6eedd0e8667d799d60af65d7127e3b0417e04bef4d954d2a4750e.
//
// Solidity: event RedeemedMulti(address indexed redeemer, address recipient, uint256 mAssetQuantity, address[] outputs, uint256[] outputQuantity, uint256 scaledFee)
func (_MStable *MStableFilterer) FilterRedeemedMulti(opts *bind.FilterOpts, redeemer []common.Address) (*MStableRedeemedMultiIterator, error) {

	var redeemerRule []interface{}
	for _, redeemerItem := range redeemer {
		redeemerRule = append(redeemerRule, redeemerItem)
	}

	logs, sub, err := _MStable.contract.FilterLogs(opts, "RedeemedMulti", redeemerRule)
	if err != nil {
		return nil, err
	}
	return &MStableRedeemedMultiIterator{contract: _MStable.contract, event: "RedeemedMulti", logs: logs, sub: sub}, nil
}

// WatchRedeemedMulti is a free log subscription operation binding the contract event 0xdf5bf345d5d6eedd0e8667d799d60af65d7127e3b0417e04bef4d954d2a4750e.
//
// Solidity: event RedeemedMulti(address indexed redeemer, address recipient, uint256 mAssetQuantity, address[] outputs, uint256[] outputQuantity, uint256 scaledFee)
func (_MStable *MStableFilterer) WatchRedeemedMulti(opts *bind.WatchOpts, sink chan<- *MStableRedeemedMulti, redeemer []common.Address) (event.Subscription, error) {

	var redeemerRule []interface{}
	for _, redeemerItem := range redeemer {
		redeemerRule = append(redeemerRule, redeemerItem)
	}

	logs, sub, err := _MStable.contract.WatchLogs(opts, "RedeemedMulti", redeemerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MStableRedeemedMulti)
				if err := _MStable.contract.UnpackLog(event, "RedeemedMulti", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRedeemedMulti is a log parse operation binding the contract event 0xdf5bf345d5d6eedd0e8667d799d60af65d7127e3b0417e04bef4d954d2a4750e.
//
// Solidity: event RedeemedMulti(address indexed redeemer, address recipient, uint256 mAssetQuantity, address[] outputs, uint256[] outputQuantity, uint256 scaledFee)
func (_MStable *MStableFilterer) ParseRedeemedMulti(log types.Log) (*MStableRedeemedMulti, error) {
	event := new(MStableRedeemedMulti)
	if err := _MStable.contract.UnpackLog(event, "RedeemedMulti", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MStableStartRampAIterator is returned from FilterStartRampA and is used to iterate over the raw logs and unpacked data for StartRampA events raised by the MStable contract.
type MStableStartRampAIterator struct {
	Event *MStableStartRampA // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MStableStartRampAIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MStableStartRampA)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MStableStartRampA)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MStableStartRampAIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MStableStartRampAIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MStableStartRampA represents a StartRampA event raised by the MStable contract.
type MStableStartRampA struct {
	CurrentA    *big.Int
	TargetA     *big.Int
	StartTime   *big.Int
	RampEndTime *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterStartRampA is a free log retrieval operation binding the contract event 0xbed20f105bdd1ca3336acef7422ceb8a840b29ffa3411edd10279120b372d6c1.
//
// Solidity: event StartRampA(uint256 currentA, uint256 targetA, uint256 startTime, uint256 rampEndTime)
func (_MStable *MStableFilterer) FilterStartRampA(opts *bind.FilterOpts) (*MStableStartRampAIterator, error) {

	logs, sub, err := _MStable.contract.FilterLogs(opts, "StartRampA")
	if err != nil {
		return nil, err
	}
	return &MStableStartRampAIterator{contract: _MStable.contract, event: "StartRampA", logs: logs, sub: sub}, nil
}

// WatchStartRampA is a free log subscription operation binding the contract event 0xbed20f105bdd1ca3336acef7422ceb8a840b29ffa3411edd10279120b372d6c1.
//
// Solidity: event StartRampA(uint256 currentA, uint256 targetA, uint256 startTime, uint256 rampEndTime)
func (_MStable *MStableFilterer) WatchStartRampA(opts *bind.WatchOpts, sink chan<- *MStableStartRampA) (event.Subscription, error) {

	logs, sub, err := _MStable.contract.WatchLogs(opts, "StartRampA")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MStableStartRampA)
				if err := _MStable.contract.UnpackLog(event, "StartRampA", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStartRampA is a log parse operation binding the contract event 0xbed20f105bdd1ca3336acef7422ceb8a840b29ffa3411edd10279120b372d6c1.
//
// Solidity: event StartRampA(uint256 currentA, uint256 targetA, uint256 startTime, uint256 rampEndTime)
func (_MStable *MStableFilterer) ParseStartRampA(log types.Log) (*MStableStartRampA, error) {
	event := new(MStableStartRampA)
	if err := _MStable.contract.UnpackLog(event, "StartRampA", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MStableStopRampAIterator is returned from FilterStopRampA and is used to iterate over the raw logs and unpacked data for StopRampA events raised by the MStable contract.
type MStableStopRampAIterator struct {
	Event *MStableStopRampA // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MStableStopRampAIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MStableStopRampA)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MStableStopRampA)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MStableStopRampAIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MStableStopRampAIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MStableStopRampA represents a StopRampA event raised by the MStable contract.
type MStableStopRampA struct {
	CurrentA *big.Int
	Time     *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStopRampA is a free log retrieval operation binding the contract event 0x46e22fb3709ad289f62ce63d469248536dbc78d82b84a3d7e74ad606dc201938.
//
// Solidity: event StopRampA(uint256 currentA, uint256 time)
func (_MStable *MStableFilterer) FilterStopRampA(opts *bind.FilterOpts) (*MStableStopRampAIterator, error) {

	logs, sub, err := _MStable.contract.FilterLogs(opts, "StopRampA")
	if err != nil {
		return nil, err
	}
	return &MStableStopRampAIterator{contract: _MStable.contract, event: "StopRampA", logs: logs, sub: sub}, nil
}

// WatchStopRampA is a free log subscription operation binding the contract event 0x46e22fb3709ad289f62ce63d469248536dbc78d82b84a3d7e74ad606dc201938.
//
// Solidity: event StopRampA(uint256 currentA, uint256 time)
func (_MStable *MStableFilterer) WatchStopRampA(opts *bind.WatchOpts, sink chan<- *MStableStopRampA) (event.Subscription, error) {

	logs, sub, err := _MStable.contract.WatchLogs(opts, "StopRampA")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MStableStopRampA)
				if err := _MStable.contract.UnpackLog(event, "StopRampA", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStopRampA is a log parse operation binding the contract event 0x46e22fb3709ad289f62ce63d469248536dbc78d82b84a3d7e74ad606dc201938.
//
// Solidity: event StopRampA(uint256 currentA, uint256 time)
func (_MStable *MStableFilterer) ParseStopRampA(log types.Log) (*MStableStopRampA, error) {
	event := new(MStableStopRampA)
	if err := _MStable.contract.UnpackLog(event, "StopRampA", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MStableSwappedIterator is returned from FilterSwapped and is used to iterate over the raw logs and unpacked data for Swapped events raised by the MStable contract.
type MStableSwappedIterator struct {
	Event *MStableSwapped // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MStableSwappedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MStableSwapped)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MStableSwapped)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MStableSwappedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MStableSwappedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MStableSwapped represents a Swapped event raised by the MStable contract.
type MStableSwapped struct {
	Swapper      common.Address
	Input        common.Address
	Output       common.Address
	OutputAmount *big.Int
	Fee          *big.Int
	Recipient    common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSwapped is a free log retrieval operation binding the contract event 0x1eeaa4acf3c225a4033105c2647625dbb298dec93b14e16253c4231e26c02b1d.
//
// Solidity: event Swapped(address indexed swapper, address input, address output, uint256 outputAmount, uint256 fee, address recipient)
func (_MStable *MStableFilterer) FilterSwapped(opts *bind.FilterOpts, swapper []common.Address) (*MStableSwappedIterator, error) {

	var swapperRule []interface{}
	for _, swapperItem := range swapper {
		swapperRule = append(swapperRule, swapperItem)
	}

	logs, sub, err := _MStable.contract.FilterLogs(opts, "Swapped", swapperRule)
	if err != nil {
		return nil, err
	}
	return &MStableSwappedIterator{contract: _MStable.contract, event: "Swapped", logs: logs, sub: sub}, nil
}

// WatchSwapped is a free log subscription operation binding the contract event 0x1eeaa4acf3c225a4033105c2647625dbb298dec93b14e16253c4231e26c02b1d.
//
// Solidity: event Swapped(address indexed swapper, address input, address output, uint256 outputAmount, uint256 fee, address recipient)
func (_MStable *MStableFilterer) WatchSwapped(opts *bind.WatchOpts, sink chan<- *MStableSwapped, swapper []common.Address) (event.Subscription, error) {

	var swapperRule []interface{}
	for _, swapperItem := range swapper {
		swapperRule = append(swapperRule, swapperItem)
	}

	logs, sub, err := _MStable.contract.WatchLogs(opts, "Swapped", swapperRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MStableSwapped)
				if err := _MStable.contract.UnpackLog(event, "Swapped", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSwapped is a log parse operation binding the contract event 0x1eeaa4acf3c225a4033105c2647625dbb298dec93b14e16253c4231e26c02b1d.
//
// Solidity: event Swapped(address indexed swapper, address input, address output, uint256 outputAmount, uint256 fee, address recipient)
func (_MStable *MStableFilterer) ParseSwapped(log types.Log) (*MStableSwapped, error) {
	event := new(MStableSwapped)
	if err := _MStable.contract.UnpackLog(event, "Swapped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MStableTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the MStable contract.
type MStableTransferIterator struct {
	Event *MStableTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MStableTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MStableTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MStableTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MStableTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MStableTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MStableTransfer represents a Transfer event raised by the MStable contract.
type MStableTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_MStable *MStableFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MStableTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MStable.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MStableTransferIterator{contract: _MStable.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_MStable *MStableFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *MStableTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MStable.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MStableTransfer)
				if err := _MStable.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_MStable *MStableFilterer) ParseTransfer(log types.Log) (*MStableTransfer, error) {
	event := new(MStableTransfer)
	if err := _MStable.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MStableUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the MStable contract.
type MStableUnpausedIterator struct {
	Event *MStableUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MStableUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MStableUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MStableUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MStableUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MStableUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MStableUnpaused represents a Unpaused event raised by the MStable contract.
type MStableUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_MStable *MStableFilterer) FilterUnpaused(opts *bind.FilterOpts) (*MStableUnpausedIterator, error) {

	logs, sub, err := _MStable.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &MStableUnpausedIterator{contract: _MStable.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_MStable *MStableFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *MStableUnpaused) (event.Subscription, error) {

	logs, sub, err := _MStable.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MStableUnpaused)
				if err := _MStable.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_MStable *MStableFilterer) ParseUnpaused(log types.Log) (*MStableUnpaused, error) {
	event := new(MStableUnpaused)
	if err := _MStable.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MStableWeightLimitsChangedIterator is returned from FilterWeightLimitsChanged and is used to iterate over the raw logs and unpacked data for WeightLimitsChanged events raised by the MStable contract.
type MStableWeightLimitsChangedIterator struct {
	Event *MStableWeightLimitsChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MStableWeightLimitsChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MStableWeightLimitsChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MStableWeightLimitsChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MStableWeightLimitsChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MStableWeightLimitsChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MStableWeightLimitsChanged represents a WeightLimitsChanged event raised by the MStable contract.
type MStableWeightLimitsChanged struct {
	Min *big.Int
	Max *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterWeightLimitsChanged is a free log retrieval operation binding the contract event 0x1633022fee8dcf5a3cdeb5f1b49d5b734a3cfef7fc093e30cfdd28ddde8cd136.
//
// Solidity: event WeightLimitsChanged(uint128 min, uint128 max)
func (_MStable *MStableFilterer) FilterWeightLimitsChanged(opts *bind.FilterOpts) (*MStableWeightLimitsChangedIterator, error) {

	logs, sub, err := _MStable.contract.FilterLogs(opts, "WeightLimitsChanged")
	if err != nil {
		return nil, err
	}
	return &MStableWeightLimitsChangedIterator{contract: _MStable.contract, event: "WeightLimitsChanged", logs: logs, sub: sub}, nil
}

// WatchWeightLimitsChanged is a free log subscription operation binding the contract event 0x1633022fee8dcf5a3cdeb5f1b49d5b734a3cfef7fc093e30cfdd28ddde8cd136.
//
// Solidity: event WeightLimitsChanged(uint128 min, uint128 max)
func (_MStable *MStableFilterer) WatchWeightLimitsChanged(opts *bind.WatchOpts, sink chan<- *MStableWeightLimitsChanged) (event.Subscription, error) {

	logs, sub, err := _MStable.contract.WatchLogs(opts, "WeightLimitsChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MStableWeightLimitsChanged)
				if err := _MStable.contract.UnpackLog(event, "WeightLimitsChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWeightLimitsChanged is a log parse operation binding the contract event 0x1633022fee8dcf5a3cdeb5f1b49d5b734a3cfef7fc093e30cfdd28ddde8cd136.
//
// Solidity: event WeightLimitsChanged(uint128 min, uint128 max)
func (_MStable *MStableFilterer) ParseWeightLimitsChanged(log types.Log) (*MStableWeightLimitsChanged, error) {
	event := new(MStableWeightLimitsChanged)
	if err := _MStable.contract.UnpackLog(event, "WeightLimitsChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
