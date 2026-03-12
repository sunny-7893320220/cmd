#!/usr/bin/env bash
set -euo pipefail

TEMPLATE_FILE="${TEMPLATE_FILE:-"$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)/deployment-template.yaml"}"
KUBECTL_ARGS=()

if [[ "${DRY_RUN:-}" == "1" ]]; then
  KUBECTL_ARGS+=(--dry-run=client)
fi

require() {
  local name="$1"
  if [[ -z "${!name:-}" ]]; then
    echo "Missing required env var: $name" >&2
    exit 1
  fi
}

# Required
require APP_NAME
require NAMESPACE
require IMAGE

# Defaults
: "${COMPONENT:=api}"
: "${REPLICAS:=1}"
: "${CONTAINER_PORT:=8080}"
: "${IMAGE_PULL_POLICY:=IfNotPresent}"
: "${HEALTH_PATH:=/}"

: "${CPU_REQUEST:=100m}"
: "${MEMORY_REQUEST:=128Mi}"
: "${CPU_LIMIT:=500m}"
: "${MEMORY_LIMIT:=512Mi}"

command -v envsubst >/dev/null 2>&1 || { echo "envsubst not found (install gettext-base)." >&2; exit 1; }
command -v kubectl  >/dev/null 2>&1 || { echo "kubectl not found." >&2; exit 1; }

# Only substitute the variables we expect (safer than substituting everything).
VARS='$APP_NAME $NAMESPACE $IMAGE $COMPONENT $REPLICAS $CONTAINER_PORT $IMAGE_PULL_POLICY $HEALTH_PATH $CPU_REQUEST $MEMORY_REQUEST $CPU_LIMIT $MEMORY_LIMIT'

envsubst "$VARS" < "$TEMPLATE_FILE" | kubectl apply "${KUBECTL_ARGS[@]}" -f -

