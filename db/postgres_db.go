package db

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

const MaxPingTry = 5

func NewPostgresDB(dataSouceName string) (*PostgreDB, error) {
	db, err := sqlx.Open("postgres", dataSouceName)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	for i := 0; i < MaxPingTry; i++ {
		if err := db.Ping(); err != nil {
			fmt.Printf("pinging data base(%d/%d)\n", i, MaxPingTry)
			time.Sleep(time.Second)
		} else {
			break
		}
	}

	return &PostgreDB{db}, nil
}

type StorageRepository interface {
	CreateUser(u *User) error
	CreateUserSession(s *Session, id uuid.UUID) error
	GetUser(id uuid.UUID) (User, error)
	GetUserByEmail(email string) error
	IsEmailTaken(email string) (bool, error)
	UpdateUser(newName string, id uuid.UUID) error
	DeleteUser(id uuid.UUID) error
	GetUserSessionByToken(token string) (Session, error)
	ExtendUserSession(id uuid.UUID) error
	RevokeUserSession(token string, id uuid.UUID) error
	RevokeUserSessions(id uuid.UUID) error
	CreateGroup(g *Group) error
	GetGroup(userId uuid.UUID, groupId uuid.UUID) (Group, error)
	GetGroups(userId uuid.UUID) ([]Group, error)
	UpdateGroup(g *Group) error
	DeleteGroup(g *Group) error
	AddMemberToGroup(m *GroupMember) error
	GetGroupMembersIDs(groupId uuid.UUID) ([]uuid.UUID, error)
	DeleteGroupMember(groupId uuid.UUID, userId uuid.UUID) error
	SendPendingRequest(pm *PendingMember) error
	RemovePendingRequest(id uuid.UUID) error
	IsGroupExist(id uuid.UUID) (bool, error)
	CreateExpense(e *Expense) error
	GetExpense(userId uuid.UUID, expId uuid.UUID) (Expense, error)
	GetExpenses(userId uuid.UUID) ([]Expense, error)
	UpdateExpense(e *Expense) error
	DeleteExpense(userId uuid.UUID, expId uuid.UUID) error
	AddExpensePayment(e *ExpensePayment) error
	GetExpensePayment(expId uuid.UUID, userId uuid.UUID) (ExpensePayment, error)
	UpdateExpensePayment(ep *ExpensePayment) error
	AddExpenseSplit(es *ExpenseSplit) error
	GetExpenseSplit(expId uuid.UUID, groupId uuid.UUID, userId uuid.UUID) (ExpenseSplit, error)
	UpdateExpenseSplit(es *ExpenseSplit) error
	AddSettleUp(su *SettleUp) error
	GetSettleUp(expId uuid.UUID, userId uuid.UUID) (SettleUp, error)
	GetSettleUps(expId uuid.UUID) ([]SettleUp, error)
	AddNotification(n *Notification) error
	MarkAsReadNotification(id uuid.UUID) error
	DeleteUserNotifications(id uuid.UUID) error
	GetDev(id uuid.UUID) (Dev, error)
	GetDevIDBySession(token string) (uuid.UUID, error)
	UpdatePrivacyPolicy(content string) error
	StatUsersCount() (int, error)
	StatsExpensesCount() (int, error)
}

type PostgreDB struct {
	*sqlx.DB
}

// User & Session

func (db *PostgreDB) CreateUser(u *User) error {
	if _, err := db.Exec(`INSERT INTO groups VALUES ($1, $2, $3, $4, $5, $6)`,
		u.ID,
		u.Name,
		u.Email,
		u.PasswordHash,
		u.CreatedAt,
	); err != nil {
		return fmt.Errorf("failed to save user: %w", err)
	}
	return nil
}

func (db *PostgreDB) CreateUserSession(s *Session, id uuid.UUID) error {
	if _, err := db.Exec(`INSERT INTO sessions VALUES ($1, $2, $3, $4)`,
		s.ID,
		s.Token,
		s.UserID,
		s.ExpireAt,
	); err != nil {
		return fmt.Errorf("failed to create session for user")
	}
	return nil
}

func (db *PostgreDB) GetUser(id uuid.UUID) (User, error) {
	var u User
	if err := db.Get(&u, `SELECT * FROM users WHERE id = $1`, id); err != nil {
		return User{}, fmt.Errorf("failed to get user: %w", err)
	}
	return u, nil
}

func (db *PostgreDB) GetUserByEmail(email string) (User, error) {
	var u User
	if err := db.Get(&u, `SELECT * FROM users WHERE email = $1`, email); err != nil {
		return User{}, fmt.Errorf("failed to get user by email: %w", err)
	}

	return u, nil
}

func (db *PostgreDB) IsEmailTaken(email string) (bool, error) {
	var exists bool
	err := db.Get(&exists, `SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)`, email)
	return exists, err
}

func (db *PostgreDB) UpdateUser(newName string, id uuid.UUID) error {
	if _, err := db.Exec(`UPDATE users SET name = $1  WHERE id = $2`,
		newName,
		id,
	); err != nil {
		return fmt.Errorf("failed to update user info: %w", err)
	}
	return nil
}

func (db *PostgreDB) DeleteUser(id uuid.UUID) error {
	// TOOD: implement TXN for ACID
	if _, err := db.Exec(`DELETE FROM users WHERE id = $1`,
		id,
	); err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}
	return nil
}

func (db *PostgreDB) GetUserSessionByToken(token string) (Session, error) {
	var s Session
	if err := db.Get(&s, `SELECT * FROM sessions WHERE token = $1`, token); err != nil {
		return Session{}, fmt.Errorf("failed to get session by token: %w", err)
	}
	return s, nil
}

func (db *PostgreDB) ExtendUserSession(id uuid.UUID) error {
	if _, err := db.Exec(`UPDATE sessions SET expire_at = NOW() + INTERVAL '7 days' WHERE user_id = $1`, id); err != nil {
		return fmt.Errorf("failed to extend user session: %w", err)
	}
	return nil
}

func (db *PostgreDB) RevokeUserSession(token string, id uuid.UUID) error {
	if _, err := db.Exec(`DELETE FROM sessions WHERE token = $1 AND user_id = $2`, token, id); err != nil {
		return fmt.Errorf("failed to revoke user session: %w", err)
	}
	return nil
}

func (db *PostgreDB) RevokeUserSessions(id uuid.UUID) error {
	if _, err := db.Exec(`DELETE FROM sessions WHERE user_id = $1`, id); err != nil {
		return fmt.Errorf("failed to expire user session: %w", err)
	}
	return nil
}

// Groups

func (db *PostgreDB) CreateGroup(g *Group) error {
	if _, err := db.Exec(`INSERT INTO groups VALUES ($1, $2, $3, $4, $5, $6, $7)`,
		g.ID,
		g.Name,
		g.Type,
		g.CurrencyISO,
		g.CreatedBy,
		g.CreatedAt,
		g.UpdatedAt); err != nil {
		return fmt.Errorf("failed to save group: %w", err)
	}
	return nil
}

func (db *PostgreDB) GetGroup(userId uuid.UUID, groupId uuid.UUID) (Group, error) {
	var g Group
	if err := db.Get(&g, `SELECT * FROM groups WHERE id = $1 AND created_by = $2`, groupId, userId); err != nil {
		return Group{}, fmt.Errorf("failed to get group: %w", err)
	}
	return g, nil
}

func (db *PostgreDB) GetGroups(userId uuid.UUID) ([]Group, error) {
	var gg []Group
	if err := db.Select(&gg, `SELECT * FROM groups WHERE created_by = $1`, userId); err != nil {
		return nil, fmt.Errorf("failed to get group: %w", err)
	}
	return gg, nil
}

func (db *PostgreDB) UpdateGroup(g *Group) error {
	if _, err := db.Exec(`UPDATE groups SET name = $1, type = $2, currency_iso = $3`, g.Name, g.Type, g.CurrencyISO); err != nil {
		return fmt.Errorf("failed to update group: %w", err)
	}
	return nil
}

func (db *PostgreDB) DeleteGroup(g *Group) error {
	// TOOD: implement TXN for ACID
	if _, err := db.Exec(`UPDATE groups SET name = $1, type = $2, currency_iso = $3`, g.Name, g.Type, g.CurrencyISO); err != nil {
		return fmt.Errorf("failed to update group: %w", err)
	}
	return nil
}

func (db *PostgreDB) AddMemberToGroup(m *GroupMember) error {
	if _, err := db.Exec(`INSERT INTO group_members VALUES ($1, $2, $3, $4)`, m.ID, m.GroupID, m.UserID, m.JoinedAt); err != nil {
		return fmt.Errorf("failed to add member to group: %w", err)
	}
	return nil
}

func (db *PostgreDB) GetGroupMembersIDs(groupId uuid.UUID) ([]uuid.UUID, error) {
	var uids []uuid.UUID
	if err := db.Select(&uids, `SELECT user_id FROM group_members WHERE group_id = $1`, groupId); err != nil {
		return nil, fmt.Errorf("failed to get group members: %w", err)
	}

	return uids, nil
}

func (db *PostgreDB) DeleteGroupMember(groupId uuid.UUID, userId uuid.UUID) error {
	if _, err := db.Exec(`DELETE FROM group_members WHERE group_id = $1 AND user_id = $1`, groupId, userId); err != nil {
		return fmt.Errorf("failed to delete group member: %w", err)
	}

	return nil
}

func (db *PostgreDB) SendPendingRequest(pm *PendingMember) error {
	if _, err := db.Exec(`INSERT INTO pending_members VALUES ($1, $2, $3, $4)`, pm.ID, pm.GroupID, pm.SenderID, pm.SentAt); err != nil {
		return fmt.Errorf("failed to add pending request: %w", err)
	}
	return nil
}

func (db *PostgreDB) RemovePendingRequest(id uuid.UUID) error {
	if _, err := db.Exec(`DELETE FROM pending_members WHERE id = $1`, id); err != nil {
		return fmt.Errorf("failed to remove pending request: %w", err)
	}
	return nil
}

func (db *PostgreDB) IsGroupExist(id uuid.UUID) (bool, error) {
	var exists bool
	err := db.Get(&exists, `SELECT EXISTS(SELECT 1 FROM groups WHERE id = $1)`, id)
	return exists, err

}

// Expenses

func (db *PostgreDB) CreateExpense(e *Expense) error {
	if _, err := db.Exec(`INSERT INTO expenses VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`,
		e.ID,
		e.Title,
		e.TotalAmount,
		e.CurrencyISO,
		e.GroupID,
		e.CreatedBy,
		e.CreatedAt,
		e.UpdatedAt,
		e.Note); err != nil {
		return fmt.Errorf("failed to create expense: %w", err)
	}
	return nil
}

func (db *PostgreDB) GetExpense(userId uuid.UUID, expId uuid.UUID) (Expense, error) {
	var e Expense
	if err := db.Get(&e, `SELECT * FROM expenses WHERE user_id = $1 AND id = $2`,
		userId,
		expId,
	); err != nil {
		return Expense{}, fmt.Errorf("failed to get expense: %w", err)
	}
	return e, nil
}

func (db *PostgreDB) GetExpenses(userId uuid.UUID) ([]Expense, error) {
	var ee []Expense
	if err := db.Select(&ee, `SELECT * FROM expenses WHERE user_id = $1`,
		userId,
	); err != nil {
		return nil, fmt.Errorf("failed to get expense: %w", err)
	}
	return ee, nil
}

func (db *PostgreDB) UpdateExpense(e *Expense) error {
	if _, err := db.Exec(`UPDATE expenses SET title = $1, total_amount = $2, currency_iso = $3, updated_at = $4, note = $5 WHERE created_by =  $6 AND id = $7`,
		e.Title,
		e.TotalAmount,
		e.CurrencyISO,
		e.UpdatedAt,
		e.Note,
		e.CreatedBy,
		e.ID,
	); err != nil {
		return fmt.Errorf("failed to update expense: %w", err)
	}
	return nil
}

func (db *PostgreDB) DeleteExpense(userId uuid.UUID, expId uuid.UUID) error {
	if _, err := db.Exec(`UPDATE expenses SET is_deleted = true WHERE created_by = $1 AND id = $2`, userId, expId); err != nil {
		return fmt.Errorf("failed to delete expense: %w", err)
	}

	return nil
}

// Expense payments (BY)

func (db *PostgreDB) AddExpensePayment(e *ExpensePayment) error {
	if _, err := db.Exec(`INSERT INTO expense_payments VALUES ($1, $2, $3, $4)`,
		e.ID,
		e.ExpenseID,
		e.PaidAmount,
		e.PayerID,
	); err != nil {
		return fmt.Errorf("failed to add expense payment: %w", err)
	}
	return nil
}

func (db *PostgreDB) GetExpensePayment(expId uuid.UUID, userId uuid.UUID) (ExpensePayment, error) {
	var e ExpensePayment
	if err := db.Get(&e, `SELECT * FROM expense_payments WHERE expense_id = $1 and payer_id = $2`, expId, userId); err != nil {
		return ExpensePayment{}, fmt.Errorf("failed to get expense payment: %w", err)
	}
	return e, nil
}

func (db *PostgreDB) UpdateExpensePayment(ep *ExpensePayment) error {
	if _, err := db.Exec(`UPDATE expense_payments SET = paid_amount = $1, payer_id = $2 WHERE id = $3 AND expense_id = $4`, ep.PaidAmount, ep.PayerID, ep.ID, ep.ExpenseID); err != nil {
		return fmt.Errorf("failed to update expense payment: %w", err)
	}

	return nil
}

// Expense splits (FOR)

func (db *PostgreDB) AddExpenseSplit(es *ExpenseSplit) error {
	if _, err := db.Exec(`INSERT INTO expense_splits VALUES($1, $2, $3, $4, $5, $6, $7)`, es.ID, es.GroupID, es.UserID, es.ExpenseID, es.SplitMethod, es.MethodValue, es.IsExclude); err != nil {
		return fmt.Errorf("failed to add expense split: %w", err)
	}
	return nil
}

func (db *PostgreDB) GetExpenseSplit(expId uuid.UUID, groupId uuid.UUID, userId uuid.UUID) (ExpenseSplit, error) {
	var es ExpenseSplit
	if err := db.Get(&es, `SELECT * FROM expense_splits WHERE expense_id = $1 AND group_id = $2 AND user_id = $3`, expId, groupId, userId); err != nil {
		return ExpenseSplit{}, fmt.Errorf("failed to get expense split: %w", err)
	}

	return es, nil
}

func (db *PostgreDB) UpdateExpenseSplit(es *ExpenseSplit) error {
	if _, err := db.Exec(`UPDATE expense_splits SET = split_method = $1, method_value = $2, is_exclude = $3 WHERE id = $4`, es.SplitMethod, es.MethodValue, es.IsExclude, es.ID); err != nil {
		return fmt.Errorf("failed to update expense split: %w", err)
	}

	return nil
}

func (db *PostgreDB) AddSettleUp(su *SettleUp) error {
	if _, err := db.Exec(`INSERT INTO settle_ups VALUES($1, $2, $3, $4, $5, $6, $7, $8)`,
		su.ID,
		su.Amount,
		su.ExpenseID,
		su.PayerID,
		su.ReceiverID,
		su.Note,
		su.CreatedBy,
		su.CreatedAt,
	); err != nil {
		return fmt.Errorf("failed to add settle up: %w", err)
	}

	return nil
}

func (db *PostgreDB) GetSettleUp(expId uuid.UUID, userId uuid.UUID) (SettleUp, error) {
	var s SettleUp
	if err := db.Get(&s, `SELECT * FROM settle_ups WHERE expense_id = $1 AND created_by = $2`, expId, userId); err != nil {
		return SettleUp{}, fmt.Errorf("failed to get settle up: %w", err)
	}

	return s, nil
}

func (db *PostgreDB) GetSettleUps(expId uuid.UUID) ([]SettleUp, error) {
	var ss []SettleUp
	if err := db.Select(&ss, `SELECT * FROM settle_ups WHERE expense_id = $1`, expId); err != nil {
		return nil, fmt.Errorf("failed to get settle ups: %w", err)
	}

	return ss, nil
}

// Notification system

func (db *PostgreDB) AddNotification(n *Notification) error {
	if _, err := db.Exec(`INSERT INTO notifications VALUES($1, $2, $3, $4, $5)`,
		n.ID,
		n.Content,
		n.ReceiverID,
		n.CreatedAt,
		n.IsRead,
	); err != nil {
		return fmt.Errorf("failed to add notification: %w", err)
	}
	return nil
}

func (db *PostgreDB) MarkAsReadNotification(id uuid.UUID) error {
	if _, err := db.Exec(`UPDATE notifications SET is_read = true WHERE id = $1`,
		id,
	); err != nil {
		return fmt.Errorf("failed to mark as read notification: %w", err)
	}
	return nil
}

func (db *PostgreDB) DeleteUserNotifications(id uuid.UUID) error {
	if _, err := db.Exec(`DELETE FROM notifications WHERE receiver_id = $1`, id); err != nil {
		return fmt.Errorf("failed to delete user notifications")
	}

	return nil
}

// Dev Layer

func (db *PostgreDB) GetDev(id uuid.UUID) (Dev, error) {
	var d Dev
	if err := db.Get(&d, `SELECT * FROM dev WHERE id = $1`, id); err != nil {
		return d, fmt.Errorf("failed to get dev: %w", err)
	}

	return d, nil
}

func (db *PostgreDB) GetDevIDBySession(token string) (uuid.UUID, error) {
	var id uuid.UUID
	if err := db.Get(&id, `SELECT dev_id FROM dev_session WHERE token = $1`, token); err != nil {
		return id, fmt.Errorf("failed to get dev: %w", err)
	}

	return id, nil
}

func (db *PostgreDB) UpdatePrivacyPolicy(content string) error {
	if _, err := db.Exec(`UPDATE privacy_policy SET content = $1`, content); err != nil {
		return fmt.Errorf("failed to update privacy policy: %w", err)
	}

	return nil
}

func (db *PostgreDB) StatUsersCount() (int, error) {
	var n int

	if err := db.Get(&n, `SELECT COUNT(id) FROM users`); err != nil {
		return n, fmt.Errorf("failed to get users count: %w", err)
	}

	return n, nil
}

func (db *PostgreDB) StatsExpensesCount() (int, error) {
	var n int

	if err := db.Get(&n, `SELECT COUNT(id) FROM expenses`); err != nil {
		return n, fmt.Errorf("failed to get expenses count: %w", err)
	}

	return n, nil
}
