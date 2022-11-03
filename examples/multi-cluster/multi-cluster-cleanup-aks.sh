#!/usr/bin/env bash
readonly AKS_GROUP=multiClusterGroup
readonly AKS_LOCATION=${1?' Azure location must be 1st argument'}

cleanup() {
    az aks delete -l ${AKS_LOCATION} -y --name clusterone -g ${AKS_GROUP}
    az aks delete -l ${AKS_LOCATION} -y --name clustertwo -g ${AKS_GROUP}
    az aks delete -l ${AKS_LOCATION} -y --name clusterthree -g ${AKS_GROUP}
    az network application-gateway delete -l ${AKS_LOCATION} --name multiClusterGateway -g ${AKS_GROUP}
    az network vnet delete -l ${AKS_LOCATION} --name multiClusterVnet  --resource-group ${AKS_GROUP}
    az network public-ip delete -l ${AKS_LOCATION} -n appGatewayIp -g ${AKS_GROUP}
    az group delete -l ${AKS_LOCATION} -g ${AKS_GROUP} -y
}

cleanup
