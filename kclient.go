package main

import (
	"fmt"
	"os"
	"time"

	"github.com/mad01/node-terminator/pkg/kutil"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/spf13/cobra"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/kubernetes/pkg/kubectl/cmd"
	cmdutil "k8s.io/kubernetes/pkg/kubectl/cmd/util"
	"k8s.io/kubernetes/pkg/kubectl/resource"
)

func newKubectlClient(kubeconfig string) *kubectlClient {
	config, err := k8sGetClientConfig(kubeconfig)
	if err != nil {
		panic(fmt.Sprintf("failed to get kube rest config: %v", err.Error()))
	}

	e := kubectlClient{
		waitInterval: 1 * time.Minute,
		ClientConfig: kutil.NewClientConfig(config, metav1.NamespaceAll),
	}

	return &e
}

// kubectlClient struct
type kubectlClient struct {
	waitInterval time.Duration
	ClientConfig clientcmd.ClientConfig
}

func (e *kubectlClient) apply(path string) error {
	f := cmdutil.NewFactory(e.ClientConfig)

	options := &cmd.ApplyOptions{
		FilenameOptions: resource.FilenameOptions{
			Filenames: []string{path},
		},
		Cascade: true,
	}

	// schema, err := f.Validator(cmdutil.GetFlagBool(cmd, "validate"), cmdutil.GetFlagBool(cmd, "openapi-validation"), cmdutil.GetFlagString(cmd, "schema-cache-dir"))
	cobraCmd := &cobra.Command{
		Use: "apply",
	}
	cobraCmd.Flags().Bool("validate", true, "")
	cobraCmd.Flags().Bool("openapi-validation", true, "")
	cobraCmd.Flags().Bool("dry-run", false, "")
	cobraCmd.Flags().Bool("overwrite", true, "")
	cobraCmd.Flags().Bool("record", false, "")
	cobraCmd.Flags().String("schema-cache-dir", "", "")
	cobraCmd.Flags().String("output", "", "")

	err := cmd.RunApply(
		f,
		cobraCmd,
		os.Stdout,
		os.Stderr,
		options,
	)
	if err != nil {
		return fmt.Errorf("failed to run apply: %v", err.Error())
	}

	return nil
}

