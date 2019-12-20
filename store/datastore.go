package store

import (
	"context"
	"errors"

	"github.com/mendersoftware/workflows/model"
)

var (
	ErrWorkflowNotFound      = errors.New("Workflow not found")
	ErrWorkflowMissingName   = errors.New("Workflow missing name")
	ErrWorkflowAlreadyExists = errors.New("Workflow already exists")
)

// DataStoreMongoInterface for DataStore  services
type DataStore interface {
	InsertWorkflows(ctx context.Context, workflow ...model.Workflow) (int, error)
	GetWorkflowByName(ctx context.Context, workflowName string) (*model.Workflow, error)
	GetWorkflows() []model.Workflow
	InsertJob(ctx context.Context, job *model.Job) (*model.Job, error)
	GetJobs(ctx context.Context) (<-chan interface{}, error)
	AquireJob(ctx context.Context, job *model.Job) (*model.Job, error)
	UpdateJobAddResult(ctx context.Context, job *model.Job, result *model.TaskResult) error
	UpdateJobStatus(ctx context.Context, job *model.Job, status int) error
	GetJobByNameAndID(ctx context.Context, name string, ID string) (*model.Job, error)
}
