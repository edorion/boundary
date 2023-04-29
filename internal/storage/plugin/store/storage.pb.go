// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: controller/storage/storage/plugin/store/v1/storage.proto

// Package store provides protobufs for storing types in the storage plugin
// package.

package store

import (
	timestamp "github.com/hashicorp/boundary/internal/db/timestamp"
	_ "github.com/hashicorp/boundary/sdk/pbs/controller/protooptions"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type StorageBucket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// public_id is a surrogate key suitable for use in a public API.
	// @inject_tag: `gorm:"primary_key"`
	PublicId string `protobuf:"bytes,1,opt,name=public_id,json=publicId,proto3" json:"public_id,omitempty" gorm:"primary_key"`
	// The scope_id of the owning scope and must be set.
	// @inject_tag: `gorm:"not_null"`
	ScopeId string `protobuf:"bytes,2,opt,name=scope_id,json=scopeId,proto3" json:"scope_id,omitempty" gorm:"not_null"`
	// name is optional. If set, it must be globally unique.
	// @inject_tag: `gorm:"default:null"`
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty" gorm:"default:null"`
	// description is optional.
	// @inject_tag: `gorm:"default:null"`
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty" gorm:"default:null"`
	// The create_time is set by the database.
	// @inject_tag: `gorm:"default:current_timestamp"`
	CreateTime *timestamp.Timestamp `protobuf:"bytes,5,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty" gorm:"default:current_timestamp"`
	// The update_time is set by the database.
	// @inject_tag: `gorm:"default:current_timestamp"`
	UpdateTime *timestamp.Timestamp `protobuf:"bytes,6,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty" gorm:"default:current_timestamp"`
	// @inject_tag: `gorm:"default:null"`
	Version uint32 `protobuf:"varint,7,opt,name=version,proto3" json:"version,omitempty" gorm:"default:null"`
	// The plugin_id is the public id of the plugin managing this storage bucket.
	// @inject_tag: `gorm:"not_null"`
	PluginId string `protobuf:"bytes,8,opt,name=plugin_id,json=pluginId,proto3" json:"plugin_id,omitempty" gorm:"not_null"`
	// The name of the external storage bucket (as opposed to the name within boundary)
	// @inject_tag: `gorm:"not_null"`
	BucketName string `protobuf:"bytes,9,opt,name=bucket_name,json=bucketName,proto3" json:"bucket_name,omitempty" gorm:"not_null"`
	// @inject_tag: `gorm:"default:null"`
	BucketPrefix string `protobuf:"bytes,10,opt,name=bucket_prefix,json=bucketPrefix,proto3" json:"bucket_prefix,omitempty" gorm:"default:null"`
	// A boolean expression that allows filtering the workers that can handle a storage buckets
	// @inject_tag: `gorm:"not_null"`
	WorkerFilter string `protobuf:"bytes,11,opt,name=worker_filter,json=workerFilter,proto3" json:"worker_filter,omitempty" gorm:"not_null"`
	// attributes is a json formatted field.
	// @inject_tag: `gorm:"not_null"`
	Attributes []byte `protobuf:"bytes,12,opt,name=attributes,proto3" json:"attributes,omitempty" gorm:"not_null"`
	// secrets_hmac is a sha256-hmac of the unencrypted secrets stored in the db.
	// @inject_tag: `gorm:"not_null"`
	SecretsHmac []byte `protobuf:"bytes,13,opt,name=secrets_hmac,json=secretsHmac,proto3" json:"secrets_hmac,omitempty" gorm:"not_null"`
}

func (x *StorageBucket) Reset() {
	*x = StorageBucket{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_storage_storage_plugin_store_v1_storage_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StorageBucket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StorageBucket) ProtoMessage() {}

func (x *StorageBucket) ProtoReflect() protoreflect.Message {
	mi := &file_controller_storage_storage_plugin_store_v1_storage_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StorageBucket.ProtoReflect.Descriptor instead.
func (*StorageBucket) Descriptor() ([]byte, []int) {
	return file_controller_storage_storage_plugin_store_v1_storage_proto_rawDescGZIP(), []int{0}
}

func (x *StorageBucket) GetPublicId() string {
	if x != nil {
		return x.PublicId
	}
	return ""
}

func (x *StorageBucket) GetScopeId() string {
	if x != nil {
		return x.ScopeId
	}
	return ""
}

func (x *StorageBucket) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *StorageBucket) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *StorageBucket) GetCreateTime() *timestamp.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *StorageBucket) GetUpdateTime() *timestamp.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *StorageBucket) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *StorageBucket) GetPluginId() string {
	if x != nil {
		return x.PluginId
	}
	return ""
}

func (x *StorageBucket) GetBucketName() string {
	if x != nil {
		return x.BucketName
	}
	return ""
}

func (x *StorageBucket) GetBucketPrefix() string {
	if x != nil {
		return x.BucketPrefix
	}
	return ""
}

func (x *StorageBucket) GetWorkerFilter() string {
	if x != nil {
		return x.WorkerFilter
	}
	return ""
}

func (x *StorageBucket) GetAttributes() []byte {
	if x != nil {
		return x.Attributes
	}
	return nil
}

func (x *StorageBucket) GetSecretsHmac() []byte {
	if x != nil {
		return x.SecretsHmac
	}
	return nil
}

type StorageBucketSecret struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// storage_bucket_id is the public id of the storage bucket using this secret.
	// @inject_tag: `gorm:"primary_key"`
	StorageBucketId string `protobuf:"bytes,1,opt,name=storage_bucket_id,json=storageBucketId,proto3" json:"storage_bucket_id,omitempty" gorm:"primary_key"`
	// secrets is the plain-text of the secret data. We are not storing this plain-text
	// value in the database.
	// @inject_tag: `gorm:"-" wrapping:"pt,secrets_data"`
	Secrets []byte `protobuf:"bytes,2,opt,name=secrets,proto3" json:"secrets,omitempty" gorm:"-" wrapping:"pt,secrets_data"`
	// ct_secrets is the ciphertext of the secret data stored in the db.
	// @inject_tag: `gorm:"column:secrets_encrypted;not_null" wrapping:"ct,secrets_data"`
	CtSecrets []byte `protobuf:"bytes,3,opt,name=ct_secrets,json=ctSecrets,proto3" json:"ct_secrets,omitempty" gorm:"column:secrets_encrypted;not_null" wrapping:"ct,secrets_data"`
	// The key_id of the kms database key used for encrypting this entry.
	// It must be set.
	// @inject_tag: `gorm:"not_null"`
	KeyId string `protobuf:"bytes,4,opt,name=key_id,json=keyId,proto3" json:"key_id,omitempty" gorm:"not_null"`
}

func (x *StorageBucketSecret) Reset() {
	*x = StorageBucketSecret{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_storage_storage_plugin_store_v1_storage_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StorageBucketSecret) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StorageBucketSecret) ProtoMessage() {}

func (x *StorageBucketSecret) ProtoReflect() protoreflect.Message {
	mi := &file_controller_storage_storage_plugin_store_v1_storage_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StorageBucketSecret.ProtoReflect.Descriptor instead.
func (*StorageBucketSecret) Descriptor() ([]byte, []int) {
	return file_controller_storage_storage_plugin_store_v1_storage_proto_rawDescGZIP(), []int{1}
}

func (x *StorageBucketSecret) GetStorageBucketId() string {
	if x != nil {
		return x.StorageBucketId
	}
	return ""
}

func (x *StorageBucketSecret) GetSecrets() []byte {
	if x != nil {
		return x.Secrets
	}
	return nil
}

func (x *StorageBucketSecret) GetCtSecrets() []byte {
	if x != nil {
		return x.CtSecrets
	}
	return nil
}

func (x *StorageBucketSecret) GetKeyId() string {
	if x != nil {
		return x.KeyId
	}
	return ""
}

var File_controller_storage_storage_plugin_store_v1_storage_proto protoreflect.FileDescriptor

var file_controller_storage_storage_plugin_store_v1_storage_proto_rawDesc = []byte{
	0x0a, 0x38, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x73, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x70, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2a, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x2a, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c,
	0x65, 0x72, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2f, 0x76, 0x31, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x93, 0x05, 0x0a, 0x0d, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x42,
	0x75, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x49, 0x64, 0x12, 0x24, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x10, 0xc2, 0xdd, 0x29,
	0x0c, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x40, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1e, 0xc2, 0xdd, 0x29, 0x1a, 0x0a, 0x0b,
	0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x4b, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x4b, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x3e, 0x0a, 0x0b, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1d, 0xc2, 0xdd, 0x29,
	0x19, 0x0a, 0x0a, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x0b, 0x62,
	0x75, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x0a, 0x62, 0x75, 0x63, 0x6b,
	0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x46, 0x0a, 0x0d, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74,
	0x5f, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x42, 0x21, 0xc2,
	0xdd, 0x29, 0x1d, 0x0a, 0x0c, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x50, 0x72, 0x65, 0x66, 0x69,
	0x78, 0x12, 0x0d, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78,
	0x52, 0x0c, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x46,
	0x0a, 0x0d, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x09, 0x42, 0x21, 0xc2, 0xdd, 0x29, 0x1d, 0x0a, 0x0c, 0x57, 0x6f, 0x72,
	0x6b, 0x65, 0x72, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x0d, 0x77, 0x6f, 0x72, 0x6b, 0x65,
	0x72, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x0c, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72,
	0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x65, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x61, 0x74, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x73, 0x5f, 0x68, 0x6d, 0x61, 0x63, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x73, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x73, 0x48, 0x6d, 0x61, 0x63, 0x22, 0x91, 0x01, 0x0a, 0x13, 0x53, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x12, 0x2a, 0x0a, 0x11, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x62, 0x75, 0x63,
	0x6b, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x73, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07,
	0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x74, 0x5f, 0x73, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x63, 0x74, 0x53,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x12, 0x15, 0x0a, 0x06, 0x6b, 0x65, 0x79, 0x5f, 0x69, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6b, 0x65, 0x79, 0x49, 0x64, 0x42, 0x43, 0x5a,
	0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x61, 0x73, 0x68,
	0x69, 0x63, 0x6f, 0x72, 0x70, 0x2f, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x72, 0x79, 0x2f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f,
	0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x3b, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_controller_storage_storage_plugin_store_v1_storage_proto_rawDescOnce sync.Once
	file_controller_storage_storage_plugin_store_v1_storage_proto_rawDescData = file_controller_storage_storage_plugin_store_v1_storage_proto_rawDesc
)

func file_controller_storage_storage_plugin_store_v1_storage_proto_rawDescGZIP() []byte {
	file_controller_storage_storage_plugin_store_v1_storage_proto_rawDescOnce.Do(func() {
		file_controller_storage_storage_plugin_store_v1_storage_proto_rawDescData = protoimpl.X.CompressGZIP(file_controller_storage_storage_plugin_store_v1_storage_proto_rawDescData)
	})
	return file_controller_storage_storage_plugin_store_v1_storage_proto_rawDescData
}

var file_controller_storage_storage_plugin_store_v1_storage_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_controller_storage_storage_plugin_store_v1_storage_proto_goTypes = []interface{}{
	(*StorageBucket)(nil),       // 0: controller.storage.storage.plugin.store.v1.StorageBucket
	(*StorageBucketSecret)(nil), // 1: controller.storage.storage.plugin.store.v1.StorageBucketSecret
	(*timestamp.Timestamp)(nil), // 2: controller.storage.timestamp.v1.Timestamp
}
var file_controller_storage_storage_plugin_store_v1_storage_proto_depIdxs = []int32{
	2, // 0: controller.storage.storage.plugin.store.v1.StorageBucket.create_time:type_name -> controller.storage.timestamp.v1.Timestamp
	2, // 1: controller.storage.storage.plugin.store.v1.StorageBucket.update_time:type_name -> controller.storage.timestamp.v1.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_controller_storage_storage_plugin_store_v1_storage_proto_init() }
func file_controller_storage_storage_plugin_store_v1_storage_proto_init() {
	if File_controller_storage_storage_plugin_store_v1_storage_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_controller_storage_storage_plugin_store_v1_storage_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StorageBucket); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_controller_storage_storage_plugin_store_v1_storage_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StorageBucketSecret); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_controller_storage_storage_plugin_store_v1_storage_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_controller_storage_storage_plugin_store_v1_storage_proto_goTypes,
		DependencyIndexes: file_controller_storage_storage_plugin_store_v1_storage_proto_depIdxs,
		MessageInfos:      file_controller_storage_storage_plugin_store_v1_storage_proto_msgTypes,
	}.Build()
	File_controller_storage_storage_plugin_store_v1_storage_proto = out.File
	file_controller_storage_storage_plugin_store_v1_storage_proto_rawDesc = nil
	file_controller_storage_storage_plugin_store_v1_storage_proto_goTypes = nil
	file_controller_storage_storage_plugin_store_v1_storage_proto_depIdxs = nil
}
