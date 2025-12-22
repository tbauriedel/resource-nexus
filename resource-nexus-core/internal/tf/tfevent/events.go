package tfevent

import "time"

type EventType string

// BaseEvent represents the base event structure. Each EventType builds on that base.
type BaseEvent struct {
	Level     string    `json:"@level"`
	Message   string    `json:"@message"`
	Module    string    `json:"@module"`
	Timestamp time.Time `json:"@timestamp"`
	Type      EventType `json:"type"`
}

// EventVersion represents the event type 'version'.
type EventVersion struct {
	BaseEvent
	UI        string `json:"ui"`
	Terraform string `json:"terraform"`
}

// LogEvent represents the event type 'log'.
type LogEvent struct {
	BaseEvent
}

// InitOutputEvent represents the event type 'init_output'.
type InitOutputEvent struct {
	BaseEvent
	MessageCode string `json:"message_code"` //nolint:tagliatelle
}

// PlannedChangeEvent represents the event type 'planned_change'.
type PlannedChangeEvent struct {
	BaseEvent
	Change Change
}

// ChangeSummaryEvent represents the event type 'change_summary'.
type ChangeSummaryEvent struct {
	BaseEvent
	Changes Changes
}

// ApplyStartEvent represents the event type 'apply_start'.
type ApplyStartEvent struct {
	BaseEvent
	Hook Hook
}

// ApplyProgressEvent represents the event type 'apply_progress'.
type ApplyProgressEvent struct {
	BaseEvent
	Hook Hook
}

// ApplyCompleteEvent represents the event type 'apply_complete'.
type ApplyCompleteEvent struct {
	BaseEvent
	Hook Hook
}

// OutputsEvent represents the event type 'outputs'.
type OutputsEvent struct {
	BaseEvent
	Outputs Outputs
}

type Change struct {
	Resource Resource `json:"resource"`
	Action   string   `json:"action"`
}

type Resource struct {
	Addr            string `json:"addr"`
	Module          string `json:"module"`
	Resource        string `json:"resource"`
	ImpliedProvider string `json:"implied_provider"` //nolint:tagliatelle
	ResourceType    string `json:"resource_type"`    //nolint:tagliatelle
	ResourceName    string `json:"resource_name"`    //nolint:tagliatelle
	ResourceKey     string `json:"resource_key"`     //nolint:tagliatelle
}

type Changes struct {
	Add               int    `json:"add"`
	Change            int    `json:"change"`
	Import            int    `json:"import"`
	Remove            int    `json:"remove"`
	ActionInvocations int    `json:"action_invocations"` //nolint:tagliatelle
	Operation         string `json:"operation"`
}

type Hook struct {
	Resource    Resource `json:"resource"`
	Action      string   `json:"action"`
	ElapsedTime int      `json:"elapsed_time,omitempty"` //nolint:tagliatelle
	IDKey       string   `json:"id_key,omitempty"`       //nolint:tagliatelle
	IDValue     string   `json:"id_value,omitempty"`     //nolint:tagliatelle
}

type Outputs struct {
	Outputs map[string]any
}
