package swagger

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"encoding/json"
	"os"

	"github.com/Azure/ARO-RP/pkg/util/stringutils"
)

func (g *generator) generateExamples(outputDir string, s *Swagger) error {
	err := os.RemoveAll(outputDir + "/examples")
	if err != nil {
		return err
	}

	err = os.MkdirAll(outputDir+"/examples", 0777)
	if err != nil {
		return err
	}

	for _, pi := range s.Paths {
		for _, op := range []*Operation{pi.Get, pi.Put, pi.Post, pi.Delete, pi.Options, pi.Head, pi.Patch} {
			if op == nil {
				continue
			}

			example := struct {
				Parameters NameParameters `json:"parameters,omitempty"`
				Responses  Responses      `json:"responses,omitempty"`
			}{
				Responses: Responses{},
			}
			for _, param := range op.Parameters {
				switch param := param.(type) {
				case Reference:
					switch param.Ref {
					case "../../../../../../common-types/resource-management/" + g.commonTypesVersion + "/types.json#/parameters/ApiVersionParameter":
						example.Parameters = append(example.Parameters, NameParameter{
							Name:      "api-version",
							Parameter: stringutils.LastTokenByte(outputDir, '/'),
						})
					case "../../../../../../common-types/resource-management/" + g.commonTypesVersion + "/types.json#/parameters/SubscriptionIdParameter":
						example.Parameters = append(example.Parameters, NameParameter{
							Name:      "subscriptionId",
							Parameter: "00000000-0000-0000-0000-000000000000",
						})
					case "../../../../../../common-types/resource-management/" + g.commonTypesVersion + "/types.json#/parameters/ResourceGroupNameParameter":
						example.Parameters = append(example.Parameters, NameParameter{
							Name:      "resourceGroupName",
							Parameter: "resourceGroup",
						})
					case "../../../../../../common-types/resource-management/" + g.commonTypesVersion + "/types.json#/parameters/LocationParameter":
						example.Parameters = append(example.Parameters, NameParameter{
							Name:      "location",
							Parameter: "location",
						})
					}
				case Parameter:
					switch param.Name {
					case "resourceName":
						example.Parameters = append(example.Parameters, NameParameter{
							Name:      param.Name,
							Parameter: "resourceName",
						})
					case "childResourceName":
						example.Parameters = append(example.Parameters, NameParameter{
							Name:      "childResourceName",
							Parameter: "childResourceName",
						})
					case "openShiftVersion":
						example.Parameters = append(example.Parameters, NameParameter{
							Name:      "openShiftVersion",
							Parameter: "4.14.40",
						})
					case "openShiftMinorVersion":
						example.Parameters = append(example.Parameters, NameParameter{
							Name:      "openShiftMinorVersion",
							Parameter: "4.14",
						})
					case "parameters":
						switch param.Schema.Ref {
						case "#/definitions/OpenShiftCluster":
							example.Parameters = append(example.Parameters, NameParameter{
								Name:      param.Name,
								Parameter: g.exampleOpenShiftClusterPutParameter(),
							})
						case "#/definitions/OpenShiftClusterUpdate":
							example.Parameters = append(example.Parameters, NameParameter{
								Name:      param.Name,
								Parameter: g.exampleOpenShiftClusterPatchParameter(),
							})
						}
					}
				}
			}

			for statusCode, resp := range op.Responses {
				if statusCode == "default" {
					continue
				}

				response := resp.(Response)

				var body interface{}
				if response.Schema != nil {
					switch response.Schema.Ref {
					case "#/definitions/OpenShiftCluster":
						if g.workerProfilesStatus {
							switch op {
							case pi.Get:
								body = g.exampleOpenShiftClusterGetResponse()
							case pi.Put, pi.Patch:
								body = g.exampleOpenShiftClusterPutOrPatchResponse()
							}
						} else {
							body = g.exampleOpenShiftClusterResponse()
						}
					case "#/definitions/OpenShiftClusterCredentials":
						body = g.exampleOpenShiftClusterCredentialsResponse()
					case "#/definitions/OpenShiftClusterAdminKubeconfig":
						body = g.exampleOpenShiftClusterAdminKubeconfigResponse()
					case "#/definitions/OpenShiftClusterList":
						body = g.exampleOpenShiftClusterListResponse()
					case "#/definitions/OperationList":
						body = g.exampleOperationListResponse()
					case "#/definitions/OpenShiftVersion":
						body = g.exampleOpenShiftVersionResponse()
					case "#/definitions/OpenShiftVersionList":
						body = g.exampleOpenShiftVersionListResponse()
					case "#/definitions/PlatformWorkloadIdentityRoleSet":
						body = g.examplePlatformWorkloadIdentityRoleSetResponse()
					case "#/definitions/PlatformWorkloadIdentityRoleSetList":
						body = g.examplePlatformWorkloadIdentityRoleSetListResponse()
					}
				}

				if statusCode == "202" {
					// If the response code is 202 Accepted, then it's a long-running operation and must have
					// a "location" header.
					example.Responses[statusCode] = struct {
						Body    interface{} `json:"body,omitempty"`
						Headers interface{} `json:"headers,omitempty"`
					}{
						Body: body,
						Headers: map[string]string{
							"location": "https://management.azure.com/subscriptions/subid/providers/Microsoft.Cache/...pathToOperationResult...",
						},
					}
				} else {
					example.Responses[statusCode] = struct {
						Body interface{} `json:"body,omitempty"`
					}{
						Body: body,
					}
				}
			}

			b, err := json.MarshalIndent(example, "", "  ")
			if err != nil {
				return err
			}

			b = append(b, '\n')

			err = os.WriteFile(outputDir+"/examples/"+op.OperationID+".json", b, 0666)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
