package frontend

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/Azure/ARO-RP/pkg/api"
	"github.com/Azure/ARO-RP/pkg/frontend/middleware"
	"github.com/Azure/ARO-RP/pkg/mimo"
)

func (f *frontend) putAdminMaintManifestMdsdCertificateRenew(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logIface := ctx.Value(middleware.ContextKeyLog)
	log, ok := logIface.(*logrus.Entry)
	if !ok {
		api.WriteError(w, http.StatusInternalServerError, api.CloudErrorCodeInternalServerError, "", "Log is nil. This is a fallback message.")
	}

	resourceID := resourceIdFromURLParams(r)
	b, err := f._putAdminMaintManifestCreate(ctx, r, resourceID, mimo.MDSD_CERT_ROTATION_ID)

	if cloudErr, ok := err.(*api.CloudError); ok {
		api.WriteCloudError(w, cloudErr)
		return
	}

	err = statusCodeError(http.StatusCreated)
	adminReply(log, w, nil, b, err)
}
