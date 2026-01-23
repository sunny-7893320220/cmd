| Command             | Description                                                                                                                               |
| ------------------- | ----------------------------------------------------------------------------------------------------------------------------------------- |
| `helmfile sync`     | 🔁 **Install or upgrade** all releases to match your Helmfile definitions. Equivalent to `helm upgrade --install`. *(Most commonly used)* |
| `helmfile apply`    | 🧩 Shows a **diff** first, then applies changes (asks for confirmation). Best for prod changes.                                           |
| `helmfile diff`     | 🕵️‍♂️ Shows what would change if you ran `helmfile sync`. Requires `helm diff` plugin.                                                   |
| `helmfile template` | 📄 Renders all Helm templates with your Helmfile values (no install). Great for debugging.                                                |
| `helmfile lint`     | 🧹 Validates all Helm charts defined in the Helmfile for syntax and best practices.                                                       |
| `helmfile list`     | 📋 Lists all releases managed by Helmfile.                                                                                                |
| `helmfile destroy`  | 💣 Deletes all Helm releases defined in the Helmfile.                                                                                     |
| `helmfile deps`     | 📦 Fetches and updates all chart dependencies (like `helm dependency update`).                                                            |



| Command                      | Description                                                              |
| ---------------------------- | ------------------------------------------------------------------------ |
| `helmfile -e <env> sync`     | Run Helmfile for a specific environment (defined under `environments:`). |
| `helmfile -l name=<release>` | Apply or sync a **specific release** by label selector.                  |
| `helmfile -f <path>`         | Use a **specific Helmfile** (e.g., for staging or dev clusters).         |


helmfile -e <environment> diff



1. helm repo list

2. helm search repo bitnami/moodle