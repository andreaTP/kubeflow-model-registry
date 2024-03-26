#! /bin/bash
set -euxo pipefail

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )

VERSION="24.3"
# TODO: support for linux
OS="osx"
# TODO: support for arm
ARCH="x86_64"

wget -q https://github.com/protocolbuffers/protobuf/releases/download/v${VERSION}/protoc-${VERSION}-${OS}-${ARCH}.zip -O ${SCRIPT_DIR}/../protoc.zip && \
  unzip -qo ${SCRIPT_DIR}/../protoc.zip -d ${SCRIPT_DIR}/.. && \
  bin/protoc --version && \
  rm ${SCRIPT_DIR}/../protoc.zip
