// Copyright (c) 2020 Red Hat, Inc.
// Copyright Contributors to the Open Cluster Management project

package dbsyncers

import (
	"fmt"
	"time"

	"github.com/stolostron/hub-of-hubs-all-in-one/pkg/db"
	"github.com/stolostron/hub-of-hubs-all-in-one/pkg/statussyncer/dbtostatus/dbsyncer"
	ctrl "sigs.k8s.io/controller-runtime"
)

// AddDBSyncers adds all the DBSyncers to the Manager.
func AddDBSyncers(mgr ctrl.Manager, database db.DB, statusSyncInterval time.Duration) error {
	addDBSyncerFunctions := []func(ctrl.Manager, db.DB, time.Duration) error{
		dbsyncer.AddPolicyDBSyncer,
	}

	for _, addDBSyncerFunction := range addDBSyncerFunctions {
		if err := addDBSyncerFunction(mgr, database, statusSyncInterval); err != nil {
			return fmt.Errorf("failed to add DB Syncer: %w", err)
		}
	}

	return nil
}
