#!/bin/sh -eu

REPOSITORY=registry.fillmore-labs.com
PROFILE=name-add

IMAGE=$(\
  env KO_DOCKER_REPO=$REPOSITORY/name-add \
  ko build --bare --sbom none ./main \
)

cat << _IMAGE > k8s/$PROFILE/image.yaml
---
apiVersion: builtin
kind: ImageTagTransformer
metadata:
  name: name-add
imageTag:
  name: name-add-image
  newName: ${IMAGE%%@*}
  digest: ${IMAGE#*@}
_IMAGE

kubectl apply -k k8s/$PROFILE
