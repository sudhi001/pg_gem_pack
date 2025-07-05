package models

import (
	"database/sql"
	"time"
)

// NotificationRepository handles database operations for notifications
type NotificationRepository struct {
	db *sql.DB
}

// NewNotificationRepository creates a new notification repository
func NewNotificationRepository(db *sql.DB) *NotificationRepository {
	return &NotificationRepository{db: db}
}

// CreateNotification creates a new notification record
func (r *NotificationRepository) CreateNotification(notification *Notification) error {
	query := `
		INSERT INTO notifications (
			title, body, type, image_url, space_uid, user_id, created_by_uid,
			total_devices, success_count, failed_count, status, is_read,
			delivered_at, read_at, created_at, updated_at
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	now := time.Now()
	_, err := r.db.Exec(query,
		notification.Title,
		notification.Body,
		notification.Type,
		notification.ImageURL,
		notification.SpaceUID,
		notification.UserID,
		notification.CreatedByUID,
		notification.TotalDevices,
		notification.SuccessCount,
		notification.FailedCount,
		notification.Status,
		notification.IsRead,
		notification.DeliveredAt,
		notification.ReadAt,
		now,
		now,
	)

	return err
}

// UpdateNotificationStatus updates the notification status
func (r *NotificationRepository) UpdateNotificationStatus(id uint, status string) error {
	query := `UPDATE notifications SET status = ?, updated_at = ? WHERE id = ?`
	_, err := r.db.Exec(query, status, time.Now(), id)
	return err
}

// UpdateNotificationCounts updates success and failed counts
func (r *NotificationRepository) UpdateNotificationCounts(id uint, successCount, failedCount int) error {
	query := `UPDATE notifications SET success_count = ?, failed_count = ?, updated_at = ? WHERE id = ?`
	_, err := r.db.Exec(query, successCount, failedCount, time.Now(), id)
	return err
}

// MarkAsDelivered marks a notification as delivered
func (r *NotificationRepository) MarkAsDelivered(id uint) error {
	now := time.Now()
	query := `UPDATE notifications SET status = 'delivered', delivered_at = ?, updated_at = ? WHERE id = ?`
	_, err := r.db.Exec(query, now, now, id)
	return err
}

// MarkAsRead marks a notification as read
func (r *NotificationRepository) MarkAsRead(id uint) error {
	now := time.Now()
	query := `UPDATE notifications SET status = 'read', is_read = true, read_at = ?, updated_at = ? WHERE id = ?`
	_, err := r.db.Exec(query, now, now, id)
	return err
}

// GetNotificationByID retrieves a notification by ID
func (r *NotificationRepository) GetNotificationByID(id uint) (*Notification, error) {
	query := `SELECT * FROM notifications WHERE id = ?`
	row := r.db.QueryRow(query, id)

	notification := &Notification{}
	err := row.Scan(
		&notification.ID,
		&notification.Title,
		&notification.Body,
		&notification.Type,
		&notification.ImageURL,
		&notification.SpaceUID,
		&notification.UserID,
		&notification.CreatedByUID,
		&notification.TotalDevices,
		&notification.SuccessCount,
		&notification.FailedCount,
		&notification.Status,
		&notification.IsRead,
		&notification.DeliveredAt,
		&notification.ReadAt,
		&notification.CreatedAt,
		&notification.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return notification, nil
}

// GetNotificationsByUserID retrieves all notifications for a user
func (r *NotificationRepository) GetNotificationsByUserID(userID string, limit, offset int) ([]*Notification, error) {
	query := `SELECT * FROM notifications WHERE user_id = ? ORDER BY created_at DESC LIMIT ? OFFSET ?`
	rows, err := r.db.Query(query, userID, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notifications []*Notification
	for rows.Next() {
		notification := &Notification{}
		err := rows.Scan(
			&notification.ID,
			&notification.Title,
			&notification.Body,
			&notification.Type,
			&notification.ImageURL,
			&notification.SpaceUID,
			&notification.UserID,
			&notification.CreatedByUID,
			&notification.TotalDevices,
			&notification.SuccessCount,
			&notification.FailedCount,
			&notification.Status,
			&notification.IsRead,
			&notification.DeliveredAt,
			&notification.ReadAt,
			&notification.CreatedAt,
			&notification.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		notifications = append(notifications, notification)
	}

	return notifications, nil
}

// GetUnreadNotificationsCount returns the count of unread notifications for a user
func (r *NotificationRepository) GetUnreadNotificationsCount(userID string) (int, error) {
	query := `SELECT COUNT(*) FROM notifications WHERE user_id = ? AND is_read = false`
	var count int
	err := r.db.QueryRow(query, userID).Scan(&count)
	return count, err
}
