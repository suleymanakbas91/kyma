#!/bin/bash

###
# Following script deploys Kyma on Minikube without a VM driver, and runs the integrations tests. 
#

set -o errexit

CURRENT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
SCRIPTS_DIR=${CURRENT_DIR}/../
INSTALLATION_DIR=${CURRENT_DIR}/../../

sudo ${INSTALLATION_DIR}/cmd/run.sh --vm-driver "none"
sudo ${SCRIPTS_DIR}/is-installed.sh
sudo ${SCRIPTS_DIR}/watch-pods.sh
sudo ${SCRIPTS_DIR}/testing.sh