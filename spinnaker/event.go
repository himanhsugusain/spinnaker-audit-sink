// Package spinnaker defines spinnaker artifacts
package spinnaker

// Root represents the top-level structure of the event
type Root struct {
	EventName string  `json:"eventName"`
	Payload   Payload `json:"payload"`
}

// Payload contains the main event data
type Payload struct {
	Content Content `json:"content"`
	Details Details `json:"details"`
}

// Content holds the execution details
type Content struct {
	Execution Execution `json:"execution"`
}

// Execution contains information about the execution
type Execution struct {
	Application             string         `json:"application"`
	Authentication          Authentication `json:"authentication"`
	BuildTime               float64        `json:"buildTime"`
	Canceled                bool           `json:"canceled"`
	EndTime                 float64        `json:"endTime"`
	ID                      string         `json:"id"`
	InitialConfig           map[string]any `json:"initialConfig"`
	KeepWaitingPipelines    bool           `json:"keepWaitingPipelines"`
	LimitConcurrent         bool           `json:"limitConcurrent"`
	MaxConcurrentExecutions int            `json:"maxConcurrentExecutions"`
	Name                    string         `json:"name"`
	Notifications           []any          `json:"notifications"`
	Origin                  string         `json:"origin"`
	PipelineConfigID        string         `json:"pipelineConfigId"`
	SpelEvaluator           string         `json:"spelEvaluator"`
	Stages                  []Stage        `json:"stages"`
	StartTime               float64        `json:"startTime"`
	Status                  string         `json:"status"`
	SystemNotifications     []any          `json:"systemNotifications"`
	Trigger                 Trigger        `json:"trigger"`
	ExecutionID             string         `json:"executionId"`
}

// Authentication holds authentication details
type Authentication struct {
	AllowedAccounts []string `json:"allowedAccounts"`
	User            string   `json:"user"`
}

// Stage represents a stage in the execution
type Stage struct {
	Context              map[string]any `json:"context"`
	EndTime              float64        `json:"endTime"`
	ID                   string         `json:"id"`
	Name                 string         `json:"name"`
	Outputs              map[string]any `json:"outputs"`
	RefID                int            `json:"refId"`
	RequisiteStageRefIDs []any          `json:"requisiteStageRefIds"`
	StartTime            float64        `json:"startTime"`
	Status               string         `json:"status"`
	Tasks                []Task         `json:"tasks"`
	Type                 string         `json:"type"`
}

// Task represents a task within a stage
type Task struct {
	EndTime              float64        `json:"endTime"`
	ID                   string         `json:"id"`
	ImplementingClass    string         `json:"implementingClass"`
	LoopEnd              bool           `json:"loopEnd"`
	LoopStart            bool           `json:"loopStart"`
	Name                 string         `json:"name"`
	StageEnd             bool           `json:"stageEnd"`
	StageStart           bool           `json:"stageStart"`
	StartTime            float64        `json:"startTime"`
	Status               string         `json:"status"`
	TaskExceptionDetails map[string]any `json:"taskExceptionDetails"`
}

// Trigger contains information about the trigger
type Trigger struct {
	Artifacts                 []any          `json:"artifacts"`
	DryRun                    bool           `json:"dryRun"`
	Notifications             []any          `json:"notifications"`
	Other                     OtherTrigger   `json:"other"`
	Parameters                map[string]any `json:"parameters"`
	Rebake                    bool           `json:"rebake"`
	ResolvedExpectedArtifacts []any          `json:"resolvedExpectedArtifacts"`
	Strategy                  bool           `json:"strategy"`
	Type                      string         `json:"type"`
	User                      []string       `json:"user"`
}

// OtherTrigger contains additional trigger details
type OtherTrigger struct {
	Artifacts                 []any          `json:"artifacts"`
	DryRun                    bool           `json:"dryRun"`
	Enabled                   bool           `json:"enabled"`
	EventID                   string         `json:"eventId"`
	ExecutionID               string         `json:"executionId"`
	ExpectedArtifacts         []any          `json:"expectedArtifacts"`
	Notifications             []any          `json:"notifications"`
	Parameters                map[string]any `json:"parameters"`
	Preferred                 bool           `json:"preferred"`
	Rebake                    bool           `json:"rebake"`
	ResolvedExpectedArtifacts []any          `json:"resolvedExpectedArtifacts"`
	Strategy                  bool           `json:"strategy"`
	Type                      string         `json:"type"`
	User                      []string       `json:"user"`
}

// Details contains metadata about the event
type Details struct {
	Application    string         `json:"application"`
	Created        int64          `json:"created"`
	RequestHeaders map[string]any `json:"requestHeaders"`
	Source         string         `json:"source"`
	Type           string         `json:"type"`
	EventID        string         `json:"eventId"`
}
