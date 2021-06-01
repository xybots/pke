/*
 * Pipeline API
 *
 * Pipeline is a feature rich application platform, built for containers on top of Kubernetes to automate the DevOps experience, continuous application development and the lifecycle of deployments. 
 *
 * API version: latest
 * Contact: info@banzaicloud.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pipeline
import (
	"time"
)
// Process struct for Process
type Process struct {
	Id string `json:"id,omitempty"`
	ParentId string `json:"parentId,omitempty"`
	OrgId int32 `json:"orgId,omitempty"`
	Type string `json:"type,omitempty"`
	Log string `json:"log,omitempty"`
	ResourceId string `json:"resourceId,omitempty"`
	ResourceType string `json:"resourceType,omitempty"`
	Status ProcessStatus `json:"status,omitempty"`
	StartedAt time.Time `json:"startedAt,omitempty"`
	FinishedAt *time.Time `json:"finishedAt,omitempty"`
	Events []ProcessEvent `json:"events,omitempty"`
}
