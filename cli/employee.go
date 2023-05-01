package cli

import (
	"log"

	"github.com/Improwised/golang-api/database/seeding"
	"github.com/spf13/cobra"
)

var seed = &cobra.Command{
	Use:   "seed",
	Short: "add employee data to postgresql",
	Long:  `this command tack employee data from erp-station and add to postgresql data-base`,

	Run: func(cmd *cobra.Command, args []string) {
		err := seeding.GetDetails()
		if err != nil {
			log.Println(err)
		}
	},
}
