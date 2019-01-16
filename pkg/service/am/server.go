// Copyright 2019 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package am

import (
	"openpitrix.io/logger"

	"openpitrix.io/iam/pkg/config"
	"openpitrix.io/iam/pkg/service/am/db"
)

type Server struct {
	cfg *config.Config
	db  *db.Database
}

func OpenServer(cfg *config.Config) (*Server, error) {
	cfg = cfg.Clone()

	db, err := db.OpenDatabase(cfg)
	if err != nil {
		logger.Criticalf(nil, "%v", err)
	}

	p := &Server{
		cfg: cfg,
		db:  db,
	}
	return p, nil
}

func (p *Server) Close() error {
	var lastErr error

	if p.db != nil {
		if err := p.db.Close(); err != nil {
			lastErr = err
			p.db = nil
		}
	}

	if lastErr != nil {
		return lastErr
	}
	return nil
}