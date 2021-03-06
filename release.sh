#!/bin/bash

read -p "[WARNING] This will bump the version to $VERSION. Are you sure? Hit [Enter] to proceed"

git add .
git tag -a v${VERSION} -m "v${VERSION} release"
git push origin v${VERSION}

echo "[INFO] version bumped to v${VERSION}"