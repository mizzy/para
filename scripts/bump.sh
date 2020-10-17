#!/bin/sh

set -e

echo current version: "$(gobump show -r)"

# shellcheck disable=SC2039
read -r -p "input next version: " next_version

gobump set "$next_version" -w
ghch -w -N v"$next_version"

git commit -am "Checking in changes prior to tagging of version v$next_version"
git tag v"$next_version"
git push && git push --tags
