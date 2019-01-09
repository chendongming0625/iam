// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: openpitrix/iam/im/im.proto

/*
Package pbim is a generated protocol buffer package.

It is generated from these files:
	openpitrix/iam/im/im.proto
	openpitrix/iam/im/filter.proto

It has these top-level messages:
	User
	UserList
	Group
	Empty
	UserId
	UserIdList
	GroupId
	GroupIdList
	Range
	Password
	ListUesrsResponse
	ListGroupsResponse
	JoinGroupRequest
	LeaveGroupRequest
	FieldValidator
*/
package pbim

import regexp "regexp"
import fmt "fmt"
import go_proto_validators "github.com/mwitkow/go-proto-validators"
import proto "github.com/golang/protobuf/proto"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import _ "github.com/golang/protobuf/ptypes/timestamp"
import _ "github.com/mwitkow/go-proto-validators"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

var _regex_User_Uid = regexp.MustCompile("^[a-z0-1_-]{2,32}$")

func (this *User) Validate() error {
	if !_regex_User_Uid.MatchString(this.Uid) {
		return go_proto_validators.FieldError("Uid", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-z0-1_-]{2,32}$"`, this.Uid))
	}
	// Validation of proto3 map<> fields is unsupported.
	if this.CreateTime != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.CreateTime); err != nil {
			return go_proto_validators.FieldError("CreateTime", err)
		}
	}
	if this.UpdateTime != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.UpdateTime); err != nil {
			return go_proto_validators.FieldError("UpdateTime", err)
		}
	}
	if this.StatusTime != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.StatusTime); err != nil {
			return go_proto_validators.FieldError("StatusTime", err)
		}
	}
	return nil
}
func (this *UserList) Validate() error {
	for _, item := range this.Value {
		if item != nil {
			if err := go_proto_validators.CallValidatorIfExists(item); err != nil {
				return go_proto_validators.FieldError("Value", err)
			}
		}
	}
	return nil
}

var _regex_Group_Gid = regexp.MustCompile("^[a-z0-1_-]{2,32}$")
var _regex_Group_GroupPath = regexp.MustCompile("^[a-z0-1_-.]{2,255}$")

func (this *Group) Validate() error {
	if !_regex_Group_Gid.MatchString(this.Gid) {
		return go_proto_validators.FieldError("Gid", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-z0-1_-]{2,32}$"`, this.Gid))
	}
	if !_regex_Group_GroupPath.MatchString(this.GroupPath) {
		return go_proto_validators.FieldError("GroupPath", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-z0-1_-.]{2,255}$"`, this.GroupPath))
	}
	// Validation of proto3 map<> fields is unsupported.
	if this.CreateTime != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.CreateTime); err != nil {
			return go_proto_validators.FieldError("CreateTime", err)
		}
	}
	if this.UpdateTime != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.UpdateTime); err != nil {
			return go_proto_validators.FieldError("UpdateTime", err)
		}
	}
	if this.StatusTime != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.StatusTime); err != nil {
			return go_proto_validators.FieldError("StatusTime", err)
		}
	}
	return nil
}
func (this *Empty) Validate() error {
	return nil
}
func (this *UserId) Validate() error {
	return nil
}
func (this *UserIdList) Validate() error {
	return nil
}
func (this *GroupId) Validate() error {
	return nil
}
func (this *GroupIdList) Validate() error {
	return nil
}
func (this *Range) Validate() error {
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *Password) Validate() error {
	return nil
}
func (this *ListUesrsResponse) Validate() error {
	for _, item := range this.User {
		if item != nil {
			if err := go_proto_validators.CallValidatorIfExists(item); err != nil {
				return go_proto_validators.FieldError("User", err)
			}
		}
	}
	return nil
}
func (this *ListGroupsResponse) Validate() error {
	for _, item := range this.Group {
		if item != nil {
			if err := go_proto_validators.CallValidatorIfExists(item); err != nil {
				return go_proto_validators.FieldError("Group", err)
			}
		}
	}
	return nil
}
func (this *JoinGroupRequest) Validate() error {
	return nil
}
func (this *LeaveGroupRequest) Validate() error {
	return nil
}