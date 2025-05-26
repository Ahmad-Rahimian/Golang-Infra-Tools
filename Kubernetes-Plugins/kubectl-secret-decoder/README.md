# kubectl-secret-decoder 🔐

A lightweight `kubectl` plugin written in Go to decode Kubernetes Secrets from the terminal, eliminating the need for manual base64 decoding.

---

## ✨ Features

- Connects to your Kubernetes cluster using kubeconfig.
- Decode one or all secrets in a specified namespace.
- Clean and colorized output for better readability.
- Supports custom kubeconfig file location.

---

## 🚀 Installation

```bash
go install github.com/Ahmad-Rahimian/Golang-Infra-Tools/Kubernetes-Plugins/kubectl-secret-decoder@latest

📦 Usage
Decode a specific secret:

kubectl-secret-decoder -n <namespace> -s <secret-name>

Decode all secrets in a namespace:

kubectl-secret-decoder -n <namespace>

Use custom kubeconfig file:

kubectl-secret-decoder -n <namespace> -k /path/to/kubeconfig

🧪 Examples

kubectl-secret-decoder -n default -s my-secret

Output:

🔐 Secret: default/my-secret
📌 username = admin
📌 password = secret123

⚙️ Flags
Flag	Short	Description	Required
--namespace	-n	Namespace of the secret(s)	Yes
--secretname	-s	Name of the secret (optional)	No
--kubeconfig	-k	Path to kubeconfig file (optional)	No
🛠 Development

Clone the repo:

git clone https://github.com/Ahmad-Rahimian/Golang-Infra-Tools.git
cd Golang-Plugins/kubectl-secret-decoder
go run main.go -n default

Build the binary:

go build -o kubectl-secret-decoder

📄 License

MIT License © Ahmad Rahimian
