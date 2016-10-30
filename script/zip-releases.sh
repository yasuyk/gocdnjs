#!/bin/bash

DIR=`pwd`
PKG_DIR=${DIR}/pkg
mkdir -p $PKG_DIR
for PLATFORM in $(find dist -mindepth 1 -maxdepth 1 -type d); do
    PLATFORM_NAME=$(basename ${PLATFORM})

    if [ ${PLATFORM_NAME} = "dist" ]; then
        continue
    fi
    cd ${PLATFORM}
    zip -r $PKG_DIR/${PLATFORM_NAME}.zip *
    cd ${DIR}
done
