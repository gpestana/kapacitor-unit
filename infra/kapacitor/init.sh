#!/bin/bash

set -m

kapacitord -config /etc/kapacitor/kapacitor-unit.conf &

fg
exit 0
