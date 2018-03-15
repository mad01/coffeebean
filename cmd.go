package coffeebean

import (
	"fmt"

	"github.com/spf13/cobra"
)

func cmdVersion() *cobra.Command {
	var command = &cobra.Command{
		Use:   "version",
		Short: "get version",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(getVersion())
		},
	}
	return command
}


func cmdApply() *cobra.Command {
	var filename string
	var kubeconfig string
	var command = &cobra.Command{
		Use:   "apply",
		Short: "apply k8s manifest file",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			client := NewKubectlClient(kubeconfig)
			err := client.Apply(filename)
			if err != nil {
				fmt.Println(err.Error())
			}
		},
	}
	command.Flags().StringVarP(&kubeconfig, "kube.config", "k", "", "outside cluster path to kube config")
	command.Flags().StringVarP(&filename, "filename", "f", "", "path to dir/file/url")
	return command
}

func runCmd() error {
	var rootCmd = &cobra.Command{Use: "kubectl-client"}
	rootCmd.AddCommand(cmdVersion())
	rootCmd.AddCommand(cmdApply())

	err := rootCmd.Execute()
	if err != nil {
		return fmt.Errorf("%v", err.Error())
	}
	return nil
}
