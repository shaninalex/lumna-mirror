#!/usr/bin/env bash
set -e

# NOTE: cz - commitizen cli versioning tool
cz bump

TAG=$(git describe --tags)

sed -i "s/const Version = \".*\"/const Version = \"v${TAG}\"/" ./version.go

git add .
git commit --amend
