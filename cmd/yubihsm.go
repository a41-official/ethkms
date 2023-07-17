package cmd

import (
	"a41-official/ethkms/yubihsm"
	"fmt"
	"strconv"
	"syscall"

	"github.com/spf13/cobra"
	"golang.org/x/term"
)

var yubihsmCmd = &cobra.Command{
	Use:   "yubihsm",
	Short: "YubiHSM2 commands",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		fmt.Print("Enter Authkey ID: ")

		var authKeyID uint16
		fmt.Scanln(&authKeyID)

		fmt.Print("Enter password: ")

		bytePassword, err := term.ReadPassword(int(syscall.Stdin))
		if err != nil {
			cmd.Println("Invalid password input")
		}
		fmt.Println()

		password := string(bytePassword)

		// Connector Init
		yubihsm.Init(authKeyID, password)
	},
}

var getOpaqueCmd = &cobra.Command{
	Use:   "get-opaque",
	Short: "Get opaque from the device",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		objectID, err := strconv.ParseUint(args[0], 10, 16)
		if err != nil {
			cmd.Println("Invalid object ID")
		}

		resp, err := yubihsm.GetOpaque(uint16(objectID))
		if err != nil {
			cmd.Println("Failed to get opaque")
		}

		cmd.Printf("Response: %v\n", resp)
	},
}

func init() {
	rootCmd.AddCommand(yubihsmCmd)

	yubihsmCmd.AddCommand(getOpaqueCmd)
}
