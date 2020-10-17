#!/bin/bash
set -e

dir=$(dirname "${0}")

DIR=$(cd "$dir"/.. && pwd)
cd "${DIR}"

test -d pkg && rm -rf ./pkg
make crossbuild

VERSION=$(gobump show -r)

echo "$VERSION"

# Generate shasum
pushd ./pkg/dist/v"${VERSION}"
shasum -a 256 ./* >./v"${VERSION}"_SHASUMS
popd
