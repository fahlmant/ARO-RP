package redhatopenshift

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"crypto/tls"
	"net/http"

	gofrsuuid "github.com/gofrs/uuid"

	"github.com/Azure/go-autorest/autorest"

	mgmtredhatopenshift20250725 "github.com/Azure/ARO-RP/pkg/client/services/redhatopenshift/mgmt/2025-07-25/redhatopenshift"
	"github.com/Azure/ARO-RP/pkg/env"
	"github.com/Azure/ARO-RP/pkg/util/azureclient"
)

// OperationsClient is a minimal interface for azure OperationsClient
type OperationsClient interface {
	OperationsClientAddons
}

type operationsClient struct {
	mgmtredhatopenshift20250725.OperationsClient
}

var _ OperationsClient = &operationsClient{}

// NewOperationsClient creates a new OperationsClient
func NewOperationsClient(environment *azureclient.AROEnvironment, subscriptionID string, authorizer autorest.Authorizer) OperationsClient {
	var client mgmtredhatopenshift20250725.OperationsClient
	if env.IsLocalDevelopmentMode() {
		client = mgmtredhatopenshift20250725.NewOperationsClientWithBaseURI("https://localhost:8443", gofrsuuid.FromStringOrNil(subscriptionID))
		client.Sender = &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true, // #nosec G402
				},
			},
		}
	} else {
		client = mgmtredhatopenshift20250725.NewOperationsClientWithBaseURI(environment.ResourceManagerEndpoint, gofrsuuid.FromStringOrNil(subscriptionID))
		client.Authorizer = authorizer
	}

	return &operationsClient{
		OperationsClient: client,
	}
}
