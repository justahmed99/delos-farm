package port

import "github.com/justahmed99/delos-farm/core/domain"

type AgentRepository interface {
	SaveAgent(agent *domain.Agent) (*domain.Agent, error)
	GetAgentsByRequestRecordIdAndName(recordId int64, name string) (*domain.Agent, error)
	CountUniqueAgent(recordId int64) (int64, error)
}
