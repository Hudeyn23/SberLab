package CLI

import (
	"CLI/CLI/Request"
	"fmt"
	"github.com/spf13/cobra"
)

var cmdECS = &cobra.Command{
	Use:   "ecs",
	Short: "Operations with ecs",
	Long:  `Operations with elastic cloud servers on SberCloud`,
}

var cmdShowECS = &cobra.Command{
	Use:   "show",
	Short: "Show list of ecs",
	Long:  `Prints list of elastic cloud servers in project`,
	Run: func(cmd *cobra.Command, args []string) {
		var reqUrl = fmt.Sprintf("https://ecs.ru-moscow-1.hc.sbercloud.ru/v1/%s/todos/detail?offset=%s&limit=%s", config.ProjectID, offset, limit)
		Request.MakeRequest(reqUrl, config.AccessKey, config.SecretKey)
	},
}
