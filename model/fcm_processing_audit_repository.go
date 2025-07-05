package models

import (
	"database/sql"
	"fmt"
	"time"
)

// FCMProcessingAuditRepository handles database operations for FCM processing audit
type FCMProcessingAuditRepository struct {
	db *sql.DB
}

// NewFCMProcessingAuditRepository creates a new FCM processing audit repository
func NewFCMProcessingAuditRepository(db *sql.DB) *FCMProcessingAuditRepository {
	return &FCMProcessingAuditRepository{db: db}
}

// CreateAuditRecord creates a new audit record for FCM processing errors
func (r *FCMProcessingAuditRepository) CreateAuditRecord(audit *FCMProcessingAudit) error {
	query := `
		INSERT INTO fcm_processing_audit 
		(message_key, message_value, error_type, error_message, processing_stage, retry_count, is_resolved, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	now := time.Now()
	_, err := r.db.Exec(query,
		audit.MessageKey,
		audit.MessageValue,
		audit.ErrorType,
		audit.ErrorMessage,
		audit.ProcessingStage,
		audit.RetryCount,
		audit.IsResolved,
		now,
		now,
	)

	if err != nil {
		return fmt.Errorf("failed to create FCM processing audit record: %w", err)
	}

	return nil
}

// UpdateRetryCount updates the retry count for a specific message
func (r *FCMProcessingAuditRepository) UpdateRetryCount(messageKey string) error {
	query := `
		UPDATE fcm_processing_audit 
		SET retry_count = retry_count + 1, updated_at = ?
		WHERE message_key = ? AND is_resolved = false
		ORDER BY created_at DESC
		LIMIT 1
	`

	_, err := r.db.Exec(query, time.Now(), messageKey)
	if err != nil {
		return fmt.Errorf("failed to update retry count: %w", err)
	}

	return nil
}

// MarkAsResolved marks an audit record as resolved
func (r *FCMProcessingAuditRepository) MarkAsResolved(messageKey string, resolutionNote string) error {
	query := `
		UPDATE fcm_processing_audit 
		SET is_resolved = true, resolution_note = ?, updated_at = ?
		WHERE message_key = ? AND is_resolved = false
		ORDER BY created_at DESC
		LIMIT 1
	`

	_, err := r.db.Exec(query, resolutionNote, time.Now(), messageKey)
	if err != nil {
		return fmt.Errorf("failed to mark audit record as resolved: %w", err)
	}

	return nil
}

// GetUnresolvedAuditRecords retrieves unresolved audit records
func (r *FCMProcessingAuditRepository) GetUnresolvedAuditRecords(limit int) ([]FCMProcessingAudit, error) {
	query := `
		SELECT id, message_key, message_value, error_type, error_message, 
		       processing_stage, retry_count, is_resolved, resolution_note, 
		       created_at, updated_at
		FROM fcm_processing_audit 
		WHERE is_resolved = false
		ORDER BY created_at DESC
		LIMIT ?
	`

	rows, err := r.db.Query(query, limit)
	if err != nil {
		return nil, fmt.Errorf("failed to get unresolved audit records: %w", err)
	}
	defer rows.Close()

	var audits []FCMProcessingAudit
	for rows.Next() {
		var audit FCMProcessingAudit
		err := rows.Scan(
			&audit.ID,
			&audit.MessageKey,
			&audit.MessageValue,
			&audit.ErrorType,
			&audit.ErrorMessage,
			&audit.ProcessingStage,
			&audit.RetryCount,
			&audit.IsResolved,
			&audit.ResolutionNote,
			&audit.CreatedAt,
			&audit.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan audit record: %w", err)
		}
		audits = append(audits, audit)
	}

	return audits, nil
}
