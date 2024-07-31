package repositories

import (
	"database/sql"
	"repo/internal/domain/entities"
	"repo/internal/domain/repositories"
)

type MySQLReportRepository struct {
	db *sql.DB
}

func NewMySQLReportRepository(db *sql.DB) repositories.ReportRepository {
	return &MySQLReportRepository{db: db}
}

func (r *MySQLReportRepository) GetAll() ([]*entities.Report, error) {
	rows, err := r.db.Query("SELECT id, user_id, title, word_count, writing_style, language FROM reports")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reports []*entities.Report
	for rows.Next() {
		var report entities.Report
		err := rows.Scan(&report.ID, &report.UserID, &report.Title, &report.WordCount, &report.WritingStyle, &report.Language)
		if err != nil {
			return nil, err
		}
		reports = append(reports, &report)
	}
	return reports, nil
}

func (r *MySQLReportRepository) GetByID(id int) (*entities.Report, error) {
	var report entities.Report
	err := r.db.QueryRow("SELECT id, user_id, title, word_count, writing_style, language FROM reports WHERE id = ?", id).
		Scan(&report.ID, &report.UserID, &report.Title, &report.WordCount, &report.WritingStyle, &report.Language)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // レポートが見つからない場合はnilを返す
		}
		return nil, err
	}
	return &report, nil
}

func (r *MySQLReportRepository) Create(report *entities.Report) error {
	result, err := r.db.Exec("INSERT INTO reports (user_id, title, word_count, writing_style, language) VALUES (?, ?, ?, ?, ?)",
		report.UserID, report.Title, report.WordCount, report.WritingStyle, report.Language)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	report.ID = int(id)
	return nil
}

func (r *MySQLReportRepository) Update(report *entities.Report) error {
	_, err := r.db.Exec("UPDATE reports SET user_id = ?, title = ?, word_count = ?, writing_style = ?, language = ? WHERE id = ?",
		report.UserID, report.Title, report.WordCount, report.WritingStyle, report.Language, report.ID)
	return err
}

func (r *MySQLReportRepository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM reports WHERE id = ?", id)
	return err
}
