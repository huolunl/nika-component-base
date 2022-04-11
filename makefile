mkfile_path := $(abspath $(lastword $(MAKEFILE_LIST)))
current_dir := $(notdir $(patsubst %/,%,$(dir $(mkfile_path))))
ROOT_DIR:=$(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))
GO := go
kitex := kitex
COMMIT_ID := $(shell git rev-parse HEAD)


.PHONY: app
app:
	@ cd pkg/nika_api/nika_application/v1/ && $(kitex) -module github.com/huolunl/nika-component-base -type protobuf nika_application.proto

.PHONY: user
user:
	@ cd pkg/nika_api/nika_user/v1/ && $(kitex) -module github.com/huolunl/nika-component-base -type protobuf nika_user.proto
.PHONY: cluster
cluster:
	@ cd pkg/nika_api/nika_cluster/v1/ && $(kitex) -module github.com/huolunl/nika-component-base -type protobuf nika_cluster.proto
.PHONY: operator
operator:
	@ cd pkg/nika_api/nika_operator/v1/ && $(kitex) -module github.com/huolunl/nika-component-base -type protobuf nika_operator.proto

.PHONY: chart
chart:
	@ cd pkg/nika_api/nika_chartmuseum/v1/ && $(kitex) -module github.com/huolunl/nika-component-base -type protobuf nika_chartmuseum.proto
