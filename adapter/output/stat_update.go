package adapter

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/justahmed99/delos-farm/core/domain"
	"github.com/justahmed99/delos-farm/program_constant"
)

func UpdateRequestRecordAndAgentForFarm(repo *GormFarmRepository, context *gin.Context, title string) (*domain.RequestRecord, error) {
	var agent_name = context.GetHeader("agent")

	if agent_name == "" {
		return nil, errors.New(program_constant.NO_AGENT_ERROR)
	}

	requestRecordRepo := NewGormRecordRequestRepository(repo.db)
	record, err_find_record := requestRecordRepo.GetRequestRecordByTitle(title)
	agentRepo := NewGormAgentRepository(repo.db)
	if err_find_record != nil {
		record = &domain.RequestRecord{
			RecordTitle:     title,
			Count:           1,
			UniqueUserAgent: 1,
		}
		_, err1 := requestRecordRepo.SaveRecord(record)
		if err1 != nil {
			return nil, err1
		}
	}
	_, err_find_agent := agentRepo.GetAgentsByRequestRecordIdAndName(record.ID, agent_name)
	if err_find_agent != nil {
		record.UniqueUserAgent += 1
		var new_agent = &domain.Agent{
			Name:            agent_name,
			RequestRecordId: record.ID,
		}
		_, err_save_agent := agentRepo.SaveAgent(new_agent)
		if err_save_agent != nil {
			return nil, err_save_agent
		}
	}
	record.Count += 1
	_, err_update_record := requestRecordRepo.UpdateRecord(record)
	if err_update_record != nil {
		return nil, err_update_record
	}

	return record, nil
}

func UpdateRequestRecordAndAgentForPond(repo *GormPondRepository, context *gin.Context, title string) (*domain.RequestRecord, error) {
	var agent_name = context.GetHeader("agent")

	if agent_name == "" {
		return nil, errors.New(program_constant.NO_AGENT_ERROR)
	}

	requestRecordRepo := NewGormRecordRequestRepository(repo.db)
	record, err_find_record := requestRecordRepo.GetRequestRecordByTitle(title)
	agentRepo := NewGormAgentRepository(repo.db)
	if err_find_record != nil {
		record = &domain.RequestRecord{
			RecordTitle:     title,
			Count:           0,
			UniqueUserAgent: 1,
		}
		_, err1 := requestRecordRepo.SaveRecord(record)
		if err1 != nil {
			return nil, err1
		}
	}

	_, err_find_agent := agentRepo.GetAgentsByRequestRecordIdAndName(record.ID, agent_name)
	if err_find_agent != nil {
		record.UniqueUserAgent += 1
		var new_agent = &domain.Agent{
			Name:            agent_name,
			RequestRecordId: record.ID,
		}
		_, err_save_agent := agentRepo.SaveAgent(new_agent)
		if err_save_agent != nil {
			return nil, err_save_agent
		}
	}
	record.Count += 1
	_, err_update_record := requestRecordRepo.UpdateRecord(record)
	if err_update_record != nil {
		return nil, err_update_record
	}

	return record, nil
}
