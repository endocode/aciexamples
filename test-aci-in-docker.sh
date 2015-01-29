#!/bin/bash
docker run --env=ACI_ENABLE=1 --env=ACI_ROOTFS=$PWD/bin/image-aci/rootfs --env=ACI_MANIFEST=$PWD/bin/image-aci/manifest ubuntu
