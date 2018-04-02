#!/bin/bash

glide update --strip-vendor

make build

./_output/mockProbe -v=3
