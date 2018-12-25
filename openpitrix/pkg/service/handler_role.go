// Copyright 2018 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package service

import (
	"context"

	"openpitrix.io/iam/openpitrix/pkg/pb"
)

func (p *Server) CreateRole(ctx context.Context, req *pb.CreateRoleRequest) (*pb.CreateRoleResponse, error) {
	return p.db.CreateRole(ctx, req)
}

func (p *Server) DeleteRoles(ctx context.Context, req *pb.DeleteRolesRequest) (*pb.DeleteRolesResponse, error) {
	return p.db.DeleteRoles(ctx, req)
}
func (p *Server) ModifyRole(ctx context.Context, req *pb.ModifyRoleRequest) (*pb.ModifyRoleResponse, error) {
	return p.db.ModifyRole(ctx, req)
}
func (p *Server) GetRole(context.Context, *pb.GetRoleRequest) (*pb.GetRoleResponse, error) {
	panic("TODO")
}
func (p *Server) DescribeRoles(context.Context, *pb.DescribeRolesRequest) (*pb.DescribeRolesResponse, error) {
	panic("TODO")
}
