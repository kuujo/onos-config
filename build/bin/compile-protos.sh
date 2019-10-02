#!/bin/sh

proto_imports=".:${GOPATH}/src/github.com/gogo/protobuf/protobuf:${GOPATH}/src/github.com/gogo/protobuf:${GOPATH}/src"

# admin.proto cannot be generated with fast marshaler/unmarshaler because it uses gnmi.ModelData
protoc -I=$proto_imports --gogo_out=Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,import_path=github.com/onosproject/onos-config/pkg/northbound/admin,plugins=grpc:. pkg/northbound/admin/*.proto
protoc -I=$proto_imports --gogo_out=Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,Mconfig/admin/admin.proto=github.com/onosproject/onos-config/pkg/northbound/admin,import_path=github.com/onosproject/onos-config/pkg/northbound/diags,plugins=grpc:. pkg/northbound/diags/*.proto
protoc -I=$proto_imports --gogofaster_out=Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,import_path=github.com/onosproject/onos-config/pkg/store/change,plugins=grpc:. pkg/store/change/*.proto
protoc -I=$proto_imports --gogofaster_out=Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,Mconfig/store/change/change.proto=github.com/onosproject/onos-config/pkg/store/change,import_path=github.com/onosproject/onos-config/pkg/store/network,plugins=grpc:. pkg/store/network/*.proto
protoc -I=$proto_imports --gogofaster_out=Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,Mconfig/store/change/change.proto=github.com/onosproject/onos-config/pkg/store/change,import_path=github.com/onosproject/onos-config/pkg/store/request,plugins=grpc:. pkg/store/request/*.proto
