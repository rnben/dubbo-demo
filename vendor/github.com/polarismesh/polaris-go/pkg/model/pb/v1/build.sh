#!/bin/bash

# 安装protoc和protoc-gen-go插件
#
# 注意：
# grpc包引入github.com/golang/protobuf/proto v1.2.0
# protoc-gen-go插件和引入proto包的版本必须保持一致
#
# github.com/golang/protobuf/
#   protoc-gen-go：在pb.go文件中插入proto.ProtoPackageIsVersionX
#   proto：在lib.go中定义ProtoPackageIsVersionX
#
# ProtoPackageIsVersion并非表示proto2/proto3

PROTOC=../protoc

${PROTOC}/bin/protoc \
--plugin=protoc-gen-go=${PROTOC}/bin/protoc-gen-go \
--go_out=plugins=grpc:. \
--proto_path=${PROTOC}/include \
--proto_path=. \
polaris_model.proto polaris_client.proto polaris_service.proto polaris_routing.proto polaris_ratelimitrule.proto \
polaris_request.proto polaris_response.proto polaris_grpcapi.proto config_file.proto config_file_response.proto \
grpc_config_api.proto
