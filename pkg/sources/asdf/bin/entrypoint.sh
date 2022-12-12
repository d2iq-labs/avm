#!/usr/bin/env bash

# export ASDF_DIR and ASDF_DATA_DIR so that asdf can be installed in a custom location
export ASDF_DIR=${ASDF_DIR:-$HOME/.local/share/avm/plugins/asdf}
export ASDF_DATA_DIR=${ASDF_DATA_DIR:-$ASDF_DIR}

# remove existing asdf installation from PATH to avoid conflicts. This script assumes asdf is installed in the default
# location
EXISTING_ASDF_BIN="$HOME/.asdf/bin"
EXISTING_ASDF_USER_SHIMS="$HOME/.asdf/shims"
[[ ":$PATH:" == *":${EXISTING_ASDF_BIN}:"* ]] && PATH="${PATH//$EXISTING_ASDF_BIN:/}"
[[ ":$PATH:" == *":${EXISTING_ASDF_USER_SHIMS}:"* ]] && PATH="${PATH//$EXISTING_ASDF_USER_SHIMS:/}"

# source asdf.sh to add asdf to PATH
source "${ASDF_DIR}/asdf.sh"

# execute the command passed to this script
asdf "$@"