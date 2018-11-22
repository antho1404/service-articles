#!/bin/bash -e

# generate mocks for Storage.
# TODO(ilgooz) have mocks as private.
mockery -name=Storage -dir ./articles -inpkg
