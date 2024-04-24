package command

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/go-woox/woox/config"
	"github.com/spf13/cobra"
)

var UpgradeCmd = &cobra.Command{
	Use:     "upgrade",
	Short:   "Upgrade the woox command.",
	Long:    "Upgrade the woox command.",
	Example: "woox upgrade",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Printf("go install %s\n", config.WooxCmd)
		cmd := exec.Command("go", "install", config.WooxCmd)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			log.Fatalf("go install %s error\n", err)
		}
		fmt.Printf("\n🎉 Woox upgrade successfully!\n\n")
	},
}
