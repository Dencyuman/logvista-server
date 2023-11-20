package converter

import (
	"encoding/json"
	"log"

	"github.com/Dencyuman/logvista-server/src/models"
	"github.com/Dencyuman/logvista-server/src/schemas"
	"github.com/Dencyuman/logvista-server/src/utils"
)

// schemas.Logをmodels.Logに変換
func ConvertLogSchemaToModel(logSchema *schemas.Log, SystemID string) *models.Log {
	if logSchema == nil {
		return nil
	}

	// AttributesをJSON文字列に変換
	attributesStr, err := json.Marshal(logSchema.Attributes)
	if err != nil {
		log.Printf("Error converting attributes to JSON: %v\n", err)
		return nil
	}

	logModel := &models.Log{
		ID:              logSchema.ID,
		SystemID:        SystemID,
		CPUPercent:      logSchema.CPUPercent,
		ExcType:         logSchema.ExcType,
		ExcValue:        logSchema.ExcValue,
		ExcDetail:       utils.ReplaceDummyLf(logSchema.ExcDetail),
		FileName:        logSchema.FileName,
		FuncName:        logSchema.FuncName,
		Lineno:          logSchema.Lineno,
		Message:         logSchema.Message,
		Module:          logSchema.Module,
		Name:            logSchema.Name,
		LevelName:       logSchema.LevelName,
		Levelno:         logSchema.Levelno,
		Process:         logSchema.Process,
		ProcessName:     logSchema.ProcessName,
		Thread:          logSchema.Thread,
		ThreadName:      logSchema.ThreadName,
		TotalMemory:     logSchema.TotalMemory,
		AvailableMemory: logSchema.AvailableMemory,
		MemoryPercent:   logSchema.MemoryPercent,
		UsedMemory:      logSchema.UsedMemory,
		FreeMemory:      logSchema.FreeMemory,
		CPUUserTime:     logSchema.CPUUserTime,
		CPUSystemTime:   logSchema.CPUSystemTime,
		CPUIdleTime:     logSchema.CPUIdleTime,
		Timestamp:       logSchema.Timestamp,
		Attributes:      string(attributesStr),
	}

	// ExcTracebackを変換
	for _, tbSchema := range logSchema.ExcTraceback {
		tbModel := ConvertTracebackSchemaToModel(&tbSchema, logSchema.ID)
		logModel.ExcTraceback = append(logModel.ExcTraceback, *tbModel)
	}

	return logModel
}

// schemas.Tracebackをmodels.Tracebackに変換
func ConvertTracebackSchemaToModel(tbSchema *schemas.Traceback, logID string) *models.Traceback {
	if tbSchema == nil {
		return nil
	}
	if logID == "" {
		return nil
	}

	return &models.Traceback{
		LogID:      logID,
		TbFilename: tbSchema.TbFilename,
		TbLineno:   tbSchema.TbLineno,
		TbName:     tbSchema.TbName,
		TbLine:     tbSchema.TbLine,
	}
}

// models.Logをschemas.LogResponseに変換
func ConvertLogModelToResponseSchema(logModel *models.Log) *schemas.LogResponse {
	if logModel == nil {
		return nil
	}

	// AttributesをJSON文字列からmap[string]interface{}に変換
	var attributes map[string]interface{}
	err := json.Unmarshal([]byte(logModel.Attributes), &attributes)
	if err != nil {
		log.Printf("Error converting attributes to JSON: %v\n", err)
		return nil
	}

	logSchema := &schemas.LogResponse{
		Log: schemas.Log{
			ID:              logModel.ID,
			CPUPercent:      logModel.CPUPercent,
			ExcType:         logModel.ExcType,
			ExcValue:        logModel.ExcValue,
			ExcDetail:       logModel.ExcDetail,
			FileName:        logModel.FileName,
			FuncName:        logModel.FuncName,
			Lineno:          logModel.Lineno,
			Message:         logModel.Message,
			Module:          logModel.Module,
			Name:            logModel.Name,
			LevelName:       logModel.LevelName,
			Levelno:         logModel.Levelno,
			Process:         logModel.Process,
			ProcessName:     logModel.ProcessName,
			Thread:          logModel.Thread,
			ThreadName:      logModel.ThreadName,
			TotalMemory:     logModel.TotalMemory,
			AvailableMemory: logModel.AvailableMemory,
			MemoryPercent:   logModel.MemoryPercent,
			UsedMemory:      logModel.UsedMemory,
			FreeMemory:      logModel.FreeMemory,
			CPUUserTime:     logModel.CPUUserTime,
			CPUSystemTime:   logModel.CPUSystemTime,
			CPUIdleTime:     logModel.CPUIdleTime,
			Timestamp:       logModel.Timestamp,
			Attributes:      attributes,
		},
		System: schemas.SystemResponse{
			ID:        logModel.SystemID,
			System:    schemas.System{Name: logModel.System.Name, Category: logModel.System.Category},
			CreatedAt: logModel.System.CreatedAt,
			UpdatedAt: logModel.System.UpdatedAt,
		},
		CreatedAt: logModel.CreatedAt,
		UpdatedAt: logModel.UpdatedAt,
	}

	// ExcTracebackを変換
	logSchema.ExcTraceback = make([]schemas.TracebackResponse, 0)
	for _, tbModel := range logModel.ExcTraceback {
		logSchema.ExcTraceback = append(logSchema.ExcTraceback, *ConvertTracebackModelToResponseSchema(&tbModel))
	}

	return logSchema
}

// models.Tracebackをschemas.TracebackResponseに変換
func ConvertTracebackModelToResponseSchema(tbModel *models.Traceback) *schemas.TracebackResponse {
	if tbModel == nil {
		return nil
	}

	return &schemas.TracebackResponse{
		Traceback: schemas.Traceback{
			TbFilename: tbModel.TbFilename,
			TbLineno:   tbModel.TbLineno,
			TbName:     tbModel.TbName,
			TbLine:     tbModel.TbLine,
		},
	}
}

// models.Systemとmodels.Log,[]schemas.SummaryDataをschemas.Summaryに変換
func ConvertSystemModelAndSummaryDataToSchema(
	systemModel *models.System,
	summaryData []schemas.SummaryData,
	latestLog *models.Log,
) *schemas.Summary {
	if systemModel == nil || summaryData == nil {
		return nil
	}

	return &schemas.Summary{
		SystemResponse: schemas.SystemResponse{
			ID: systemModel.ID,
			System: schemas.System{
				Name:     systemModel.Name,
				Category: systemModel.Category,
			},
			CreatedAt: systemModel.CreatedAt,
			UpdatedAt: systemModel.UpdatedAt,
		},
		LatestLog: *ConvertLogModelToResponseSchema(latestLog),
		Data:      summaryData,
	}
}
