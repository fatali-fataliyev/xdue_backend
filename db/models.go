package db

import (
	"time"

	"github.com/google/uuid"
)

// --- User & Session Layer ---

type User struct {
	ID           uuid.UUID `db:"id"`
	Name         string    `db:"name"`
	Email        string    `db:"email"`
	PasswordHash string    `db:"password_hash"`
	CreatedAt    time.Time `db:"created_at"`
}

type Session struct {
	ID       uuid.UUID `db:"id"`
	Token    string    `db:"token"`
	UserID   uuid.UUID `db:"user_id"`
	ExpireAt time.Time `db:"expire_at"`
}

// --- Group Layer ---

type Group struct {
	ID          uuid.UUID `db:"id"`
	Name        string    `db:"name"`
	Type        string    `db:"type"`
	CurrencyISO string    `db:"currency_iso"`
	CreatedBy   uuid.UUID `db:"created_by"`
	UpdatedAt   time.Time `db:"updated_at"`
	CreatedAt   time.Time `db:"created_at"`
}

type GroupMember struct {
	ID       uuid.UUID `db:"id"`
	GroupID  uuid.UUID `db:"group_id"`
	UserID   uuid.UUID `db:"user_id"`
	JoinedAt time.Time `db:"joined_at"`
}

type PendingMember struct {
	ID       uuid.UUID `db:"id"`
	GroupID  uuid.UUID `db:"group_id"`
	SenderID uuid.UUID `db:"sender_id"`
	SentAt   time.Time `db:"sent_at"`
}

// --- Expense Layer ---

type Expense struct {
	ID          uuid.UUID `db:"id"`
	Title       string    `db:"title"`
	TotalAmount int       `db:"total_amount"`
	CurrencyISO string    `db:"currency_iso"`
	GroupID     uuid.UUID `db:"group_id"`
	CreatedBy   uuid.UUID `db:"created_by"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
	Note        string    `db:"note"`
}

type DeletedExpense struct {
	ID          uuid.UUID `db:"id"`
	Title       string    `db:"title"`
	TotalAmount int       `db:"total_amount"`
	CurrencyISO string    `db:"currency_iso"`
	GroupID     uuid.UUID `db:"group_id"`
	DeletedBy   uuid.UUID `db:"deleted_by"`
	CreatedAt   time.Time `db:"created_at"`
	DeletedAt   time.Time `db:"deleted_at"`
}

type ExpensePayment struct {
	ID         uuid.UUID `db:"id"`
	ExpenseID  uuid.UUID `db:"expense_id"`
	PaidAmount int       `db:"paid_amount"`
	PayerID    uuid.UUID `db:"payer_id"`
}

type ExpenseSplit struct {
	ID          uuid.UUID `db:"id"`
	GroupID     uuid.UUID `db:"group_id"`
	UserID      uuid.UUID `db:"user_id"`
	ExpenseID   uuid.UUID `db:"expense_id"`
	SplitMethod string    `db:"split_method"`
	MethodValue int       `db:"method_value"`
	IsExclude   bool      `db:"is_exclude"`
}

// --- Communication & Feedback ---

type Notification struct {
	ID         uuid.UUID `db:"id"`
	Content    string    `db:"content"`
	ReceiverID uuid.UUID `db:"receiver_id"`
	CreatedAt  time.Time `db:"created_at"`
	IsRead     bool      `db:"is_read"`
}

type Feedback struct {
	ID        uuid.UUID `db:"id"`
	Category  string    `db:"category"`
	AppRating int       `db:"app_rating"`
	Feedback  string    `db:"feedback"`
	Name      *string   `db:"name"`
	Email     *string   `db:"email"`
}

type SettleUp struct {
	ID         uuid.UUID `db:"id"`
	Amount     int       `db:"amount"`
	ExpenseID  uuid.UUID `db:"expense_id"`
	PayerID    uuid.UUID `db:"payer_id"`
	ReceiverID uuid.UUID `db:"receiver_id"`
	Note       string    `db:"note"`
	CreatedBy  uuid.UUID `db:"created_by"`
	CreatedAt  time.Time `db:"created_at"`
}

// --- Developer Layer ---

type Dev struct {
	ID           uuid.UUID `db:"id"`
	Email        string    `db:"email"`
	PasswordHash string    `db:"password_hash"`
}

type DevSession struct {
	ID       uuid.UUID `db:"id"`
	Token    string    `db:"token"`
	DevID    uuid.UUID `db:"dev_id"`
	ExpireAt time.Time `db:"expire_at"`
}

type PrivacyPolicy struct {
	Content   string    `db:"content"`
	UpdatedAt time.Time `db:"updated_at"`
}
