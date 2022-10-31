#!/usr/bin/env bash
readonly PROJECT_ROOT="$(dirname "$(dirname "$(dirname "$0")")")"
readonly RELEASE_NAME=volume-selector
readonly AWS_ZONE=${1?' AWS zone must be 1st argument'}
helm_install() {
    pushd "${PROJECT_ROOT}" > /dev/null || exit
    local -r volumeId=$(aws ec2 create-volume \
        --availability-zone="${AWS_ZONE}" \
        --size=10 \
        --volume-type=gp3 \
        --tag-specifications 'ResourceType=volume,Tags=[{Key=volume,Value='"${RELEASE_NAME}"'}]' \
        --no-cli-pager \
        --output text \
        --query VolumeId)

    helm install "${RELEASE_NAME}"-disk neo4j-persistent-volume \
        --set neo4j.name="${RELEASE_NAME}" \
        --set data.driver=ebs.csi.aws.com \
        --set data.reclaimPolicy="Delete" \
        --set data.createPvc=false \
        --set data.createStorageClass=true \
        --set data.volumeHandle="${volumeId}" \
        --set data.capacity.storage=10Gi
    helm install  "${RELEASE_NAME}" neo4j -f examples/persistent-volume-selector/persistent-volume-selector.yaml
}

helm_install
