package db

import (
	"time"

	"github.com/google/uuid"
)

// User Layer

type Expense struct {
	ID          uuid.UUID `db:"id"`
	Title       string    `db:"title"`
	TotalAmount int       `db:"total_amount"`
	CurrencyISO string    `db:"currency_iso"`
	GroupID     uuid.UUID `db:"group_id"`
	CreatedBy   uuid.UUID `db:"created_by"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

type DeletedExpense struct {
	ID          uuid.UUID `db:"id"`
	Title       string    `db:"title"`
	TotalAmount int       `db:"total_amount"`
	CurrencyISO string    `db:"currency_iso"`
	GroupID     uuid.UUID `db:"group_id"`
	CreatedBy   uuid.UUID `db:"created_by"`
	DeletedBy   uuid.UUID `db:"deleted_by"`
	CreatedAt   time.Time `db:"created_at"`
	DeletedAt   time.Time `db:"deleted_at"`
}

type ExpensePayments struct {
	ID         uuid.UUID `db:"id"`
	ExpenseID  uuid.UUID `db:"expense_id"`
	PaidAmount int       `db:"paid_amount"`
	PayerID    uuid.UUID `db:"payer_id"`
}

type ExpenseSplits struct {
	ID          uuid.UUID `db:"id"`
	GroupID     uuid.UUID `db:"group_id"`
	MemberID    uuid.UUID `db:"member_id"`
	ExpenseID   uuid.UUID `db:"expense_id"`
	SplitMethod string    `db:"split_method"`
	MethodValue int       `db:"method_value"`
	IsExclude   bool      `db:"is_exclude"`
}

type Groups struct {
	ID          uuid.UUID `db:"id"`
	Name        string    `db:"name"`
	Type        string    `db:"type"`
	CurrencyISO string    `db:"currency_iso"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

type Members struct {
	ID        uuid.UUID `db:"id"`
	Name      string    `db:"name"`
	BackupKey string    `db:"backup_key"`
	GroupID   uuid.UUID `db:"group_id"`
	Role      string    `db:"role"`
}

type PendingMembers struct {
	ID      uuid.UUID `db:"id"`
	GroupID uuid.UUID `db:"group_id"`
	Name    string    `db:"name"`
	SentAt  time.Time `db:"sent_at"`
}

type Sessions struct {
	ID       uuid.UUID `db:"id"`
	Token    string    `db:"token"`
	MemberID uuid.UUID `db:"member_id"`
	ExpireAt time.Time `db:"expire_at"`
}

// Developer Layer

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
