package cmd

import (
	"github.com/0xERR0R/go-cache"
	"github.com/spf13/cobra"
)

func newClientCommand() *cobra.Command {
	c := &cobra.Command{
		Use:		"clients",
		Aliases: 	[]string{"client"},
		Short: 		"Change ip of client name",
	}

	changeCommand := &cobra.Command{
		Use: 	"change",
		Args: 	cobra.NoArgs,
		Short:  "Change ip of client name",
		Run: 	changeIpOnCache,
	}
	changeCommand.Flags().IPP("ip", "i", nil, "client ip")
	changeCommand.Flags().StringP("name", "n", "", "client name")
	c.AddCommand(changeCommand)

	return c
}

func changeIpOnCache(cmd *cobra.Command, _ []string) {
	ip, _ := cmd.Flags().GetIP("ip")
	name, _ := cmd.Flags().GetString("name")

	var cache *cache.Cache

	c, found := cache.Get(ip.String())
	if found {
		if c == nil {

		}
		cache.Delete(ip.String())
	}

	cache.Set(ip.String(), name, 0)
}