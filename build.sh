#!/bin/bash
set -e

TAG=$(git tag --sort=-v:refname | head -n 1)

echo -e "\nLatest Git tag: $TAG"

build_and_push() {
    local name=$1
    local dockerfile=$2

    echo -e "\nBUILD: $name"
    docker build -t lumna_"$name" -f "$dockerfile" .
    docker tag lumna_"$name" shanart/lumna_"$name":"$TAG"
    docker push shanart/lumna_"$name":"$TAG"
}

build_and_push "auth" "./apps/auth/Dockerfile"
build_and_push "org" "./apps/org/Dockerfile"
build_and_push "project" "./apps/project/Dockerfile"
build_and_push "user" "./apps/user/Dockerfile"
build_and_push "frontend" "./frontend/Dockerfile"
