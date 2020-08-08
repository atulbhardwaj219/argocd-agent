package codefresh

import "fmt"

type Environment struct {
	FinishedAt   string                `json:"finishedAt"`
	HealthStatus string                `json:"healthStatus"`
	SyncStatus   string                `json:"status"`
	HistoryId    int64                 `json:"historyId"`
	SyncRevision string                `json:"revision"`
	Name         string                `json:"name"`
	Activities   []EnvironmentActivity `json:"activities"`
	RepoUrl      string                `json:"repoUrl"`
}

type EnvironmentActivity struct {
	Name         string   `json:"name"`
	TargetImages []string `json:"targetImages"`
	Status       string   `json:"status"`
	LiveImages   []string `json:"liveImages"`
}

type CodefreshError struct {
	Status  int         `json:"status"`
	Code    string      `json:"code"`
	Name    string      `json:"name"`
	Message string      `json:"message"`
	Context interface{} `json:"context"`
}

func (e *CodefreshError) Error() string {
	return fmt.Sprintf("Request failed, %s - %s", e.Code, e.Message)
}

type AgentApplication struct {
	Name      string `json:"name"`
	UID       string `json:"uid"`
	Project   string `json:"project"`
	Namespace string `json:"namespace"`
	Server    string `json:"server"`
}

type AgentProject struct {
	Name string `json:"name"`
	UID  string `json:"uid"`
}

type AgentState struct {
	Kind  string      `json:"type"`
	Items interface{} `json:"items"`
}

type IntegrationPayloadData struct {
	Name     string `json:"name"`
	Url      string `json:"url"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type IntegrationPayload struct {
	Type string                 `json:"type"`
	Data IntegrationPayloadData `json:"data"`
}

type requestOptions struct {
	path   string
	method string
	body   interface{}
	qs     map[string]string
}
