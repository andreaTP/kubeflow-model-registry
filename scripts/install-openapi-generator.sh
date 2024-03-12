#! /bin/bash
set -euxo pipefail

PROJECT_BIN=${1}

if ! command -v npm &> /dev/null
then
	echo "npm is not available please install it to be able to install openapi-generator"
	exit 1
fi

if ! command -v ${PROJECT_BIN}/openapi-generator-cli &> /dev/null
then
	mkdir -p ${PROJECT_BIN}
	mkdir -p ${PROJECT_BIN}/openapi-generator-installation
	{
    cd ${PROJECT_BIN}
	  npm install -g --prefix ${PROJECT_BIN}/openapi-generator-installation @openapitools/openapi-generator-cli
	  ln -s openapi-generator-installation/bin/openapi-generator-cli openapi-generator-cli
	}
fi
