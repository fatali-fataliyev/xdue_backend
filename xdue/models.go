package xdue

import (
	"time"

	"github.com/google/uuid"
)

// MODELS

type User struct {
	ID           uuid.UUID
	Name         string
	Email        string
	PasswordHash string
	CreatedAt    time.Time
}

type Session struct {
	ID       uuid.UUID
	Token    string
	UserID   uuid.UUID
	ExpireAt time.Time
}

// --- Group Layer ---

type Group struct {
	ID          uuid.UUID
	Name        string
	Type        string
	CurrencyISO string
	CreatedBy   uuid.UUID
	UpdatedAt   time.Time
	CreatedAt   time.Time
}

type GroupMember struct {
	ID       uuid.UUID
	GroupID  uuid.UUID
	UserID   uuid.UUID
	JoinedAt time.Time
}

type PendingMember struct {
	ID       uuid.UUID
	GroupID  uuid.UUID
	SenderID uuid.UUID
	SentAt   time.Time
}

// --- Expense Layer ---

type Expense struct {
	ID          uuid.UUID
	Title       string
	TotalAmount int
	CurrencyISO string
	GroupID     uuid.UUID
	CreatedBy   uuid.UUID
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Note        string
}

type DeletedExpense struct {
	ID          uuid.UUID
	Title       string
	TotalAmount int
	CurrencyISO string
	GroupID     uuid.UUID
	DeletedBy   uuid.UUID
	CreatedAt   time.Time
	DeletedAt   time.Time
}

type ExpensePayment struct {
	ID         uuid.UUID
	ExpenseID  uuid.UUID
	PaidAmount int
	PayerID    uuid.UUID
}

type ExpenseSplit struct {
	ID          uuid.UUID
	GroupID     uuid.UUID
	UserID      uuid.UUID
	ExpenseID   uuid.UUID
	SplitMethod string
	MethodValue int
	IsExclude   bool
}

// --- Communication & Feedback ---

type Notification struct {
	ID         uuid.UUID
	Content    string
	ReceiverID uuid.UUID
	CreatedAt  time.Time
	IsRead     bool
}

type Feedback struct {
	ID        uuid.UUID
	Category  string
	AppRating int
	Feedback  string
	Name      *string
	Email     *string
}

type SettleUp struct {
	ID         uuid.UUID
	Amount     int
	ExpenseID  uuid.UUID
	PayerID    uuid.UUID
	ReceiverID uuid.UUID
	Note       string
	CreatedBy  uuid.UUID
	CreatedAt  time.Time
}

// --- Developer Layer ---

type Dev struct {
	ID           uuid.UUID
	Email        string
	PasswordHash string
}

type DevSession struct {
	ID       uuid.UUID
	Token    string
	DevID    uuid.UUID
	ExpireAt time.Time
}

type PrivacyPolicy struct {
	Content   string
	UpdatedAt time.Time
}
