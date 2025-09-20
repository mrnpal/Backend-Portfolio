package repository

import (
	"database/sql"
	"portfolio-website/internal/models"
)

type ContactRepository struct {
	db *sql.DB
}

func NewContactRepository(db *sql.DB) *ContactRepository {
	return &ContactRepository{db: db}
}

func (r *ContactRepository) GetAllContacts() ([]models.Contact, error) {
	rows, err := r.db.Query("SELECT id, name, email, message, created_at FROM contacts ORDER BY created_at DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var contacts []models.Contact
	for rows.Next() {
		var c models.Contact
		if err := rows.Scan(&c.ID, &c.Name, &c.Email, &c.Message, &c.CreatedAt); err != nil {
			return nil, err
		}
		contacts = append(contacts, c)
	}

	return contacts, nil
}

func (r *ContactRepository) GetContactByID(id int) (*models.Contact, error) {
	var c models.Contact
	err := r.db.QueryRow("SELECT id, name, email, message, created_at FROM contacts WHERE id = ?", id).
		Scan(&c.ID, &c.Name, &c.Email, &c.Message, &c.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &c, nil
}

func (r *ContactRepository) CreateContact(contact *models.Contact) (int64, error) {
	result, err := r.db.Exec(
		"INSERT INTO contacts (name, email, message) VALUES (?, ?, ?)",
		contact.Name, contact.Email, contact.Message,
	)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *ContactRepository) DeleteContact(id int) error {
	_, err := r.db.Exec("DELETE FROM contacts WHERE id = ?", id)
	return err
}
