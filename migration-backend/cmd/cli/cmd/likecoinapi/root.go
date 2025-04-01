package likecoinapi

import (
	"github.com/spf13/cobra"
)

var LikecoinAPICmd = &cobra.Command{
	Use:   "likecoin-api",
	Short: "Likecoin API CLI",
	Long:  `Likecoin API CLI`,
}
