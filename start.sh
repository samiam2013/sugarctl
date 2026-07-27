#!/bin/bash

entry=$(gopass find libreview.com)
export LIBRE_LINKUP_USER=$(gopass show "$entry" username)
export LIBRE_LINKUP_PASS=$(gopass show -o "$entry")

go build .
./sugarctl
