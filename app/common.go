package app

/*** config ***/
type EnvSubsciption struct {
	QueueName string
	Raw       bool
}

type EnvTopic struct {
	Name          string
	Subscriptions []EnvSubsciption
}

type EnvQueue struct {
	Name string
}

type Environment struct {
	Host        string
	Port        string
	SqsPort     string
	SnsPort     string
	Region      string
	LogMessages bool
	LogFile     string
	Topics      []EnvTopic
	Queues      []EnvQueue
}

var CurrentEnvironment Environment

/*** Common ***/
type ResponseMetadata struct {
	RequestId string `xml:"RequestId"`
}

/*** Error Responses ***/
type ErrorResult struct {
	Type      string `xml:"Type,omitempty"`
	Code      string `xml:"Code,omitempty"`
	Message   string `xml:"Message,omitempty"`
	RequestId string `xml:"RequestId,omitempty"`
}

type ErrorResponse struct {
	Result ErrorResult `xml:"Error"`
}
