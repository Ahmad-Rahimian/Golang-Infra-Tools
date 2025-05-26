package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var decodeCmd = &cobra.Command{
	Use:   "decode",
	Short: "Decode Kubernetes Secrets",
	Long:  `Connects to your Kubernetes cluster and decodes one or more Secrets from a specified namespace.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Read flags
		namespace, _ := cmd.Flags().GetString("namespace")
		secretName, _ := cmd.Flags().GetString("name")
		kubeconfig, _ := cmd.Flags().GetString("kubeconfig")

		if namespace == "" {
			log.Fatal("❌ --namespace (-n) is required")
		}

		// Default kubeconfig
		if kubeconfig == "" {
			home, err := os.UserHomeDir()
			if err != nil {
				log.Fatalf("⚠️ Cannot determine home directory: %v", err)
			}
			kubeconfig = filepath.Join(home, ".kube", "config")
		}

		// Load kubeconfig
		config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
		if err != nil {
			log.Fatalf("❌ Error loading kubeconfig: %v", err)
		}

		// Create client
		clientset, err := kubernetes.NewForConfig(config)
		if err != nil {
			log.Fatalf("❌ Error creating Kubernetes client: %v", err)
		}

		ctx := context.TODO()

		if secretName != "" {
			secret, err := clientset.CoreV1().Secrets(namespace).Get(ctx, secretName, metav1.GetOptions{})
			if err != nil {
				log.Fatalf("❌ Error fetching secret %q: %v", secretName, err)
			}
			printHeader(namespace, secret.Name)
			decodeAndPrint(secret.Data)
		} else {
			secrets, err := clientset.CoreV1().Secrets(namespace).List(ctx, metav1.ListOptions{})
			if err != nil {
				log.Fatalf("❌ Error listing secrets in namespace %q: %v", namespace, err)
			}
			if len(secrets.Items) == 0 {
				fmt.Printf("ℹ️ No secrets found in namespace %q\n", namespace)
				return
			}
			for _, secret := range secrets.Items {
				printHeader(namespace, secret.Name)
				decodeAndPrint(secret.Data)
				fmt.Println("─────")
			}
		}

		fmt.Println("✅ Done decoding secrets.")
	},
}

func printHeader(namespace, secretname string) {
	fmt.Printf("\n🔐 Secret: \033[1;34m%s/%s\033[0m\n", namespace, secretname)
}

func decodeAndPrint(data map[string][]byte) {
	for key, val := range data {
		fmt.Printf("📌 \033[1m%s\033[0m = %s\n", key, string(val))
	}
}

func init() {
	rootCmd.AddCommand(decodeCmd)

	decodeCmd.Flags().StringP("namespace", "n", "", "Namespace of the Secret (required)")
	decodeCmd.Flags().StringP("name", "s", "", "Name of the Secret (optional; if omitted, decode all secrets in the namespace)")
	decodeCmd.Flags().StringP("kubeconfig", "k", "", "Path to the kubeconfig file (optional)")
}
