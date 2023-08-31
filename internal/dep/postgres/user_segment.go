package postgres

import (
	"database/sql"
	"fmt"
	"github.com/fanat1k666/backend-trainee-assignment-2023/internal/repository"
	_ "github.com/lib/pq"
)

type Postgres struct {
	db *sql.DB
}

func NewPostgres(host string, port int, user string, password string, dbname string) (*Postgres, error) {
	db, err := sql.Open("postgres",
		fmt.Sprintf(`host=%s port=%d user=%s password=%s dbname=%s sslmode=disable`,
			host, port, user, password, dbname),
	)
	if err != nil {
		return nil, fmt.Errorf("can't open db connection: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("can't ping db: %w", err)
	}

	return &Postgres{
		db: db,
	}, nil
}

func (p *Postgres) CreateSegment(name string) error {
	_, err := p.db.Exec(`INSERT INTO segment (name) VALUES ($1)`, name)
	if err != nil {
		return fmt.Errorf("can't create segment: %w", err)
	}
	return nil
}

func (p *Postgres) DropSegment(name string) error {
	_, err := p.db.Exec(`DELETE FROM segment WHERE name=$1`, name)
	if err != nil {
		return fmt.Errorf("can't drop segment: %w", err)
	}
	return nil
}

func (p *Postgres) DropUserFromSegment(userId int, name string) error {
	_, err := p.db.Exec(`DELETE FROM user_segment
WHERE EXISTS
(SELECT u.user_id, name FROM segment
                                 INNER JOIN user_segment ON segment.id = user_segment.segment_id
                                 INNER JOIN public."user" u on user_segment.user_id = u.id
 WHERE u.user_id = $1 AND segment.name = $2)`, userId, name)
	if err != nil {
		return fmt.Errorf("can't drop user from segment: %w", err)
	}
	return nil
}
func (p *Postgres) UserSegment(userId int, segmentId []string) error {
	_, err := p.db.Exec(`INSERT INTO "user" (user_id) VALUES ($1) ON CONFLICT DO NOTHING`, userId)
	if err != nil {
		return fmt.Errorf("can't add new User: %w", err)
	}

	sqlStr := `
WITH user_name AS (SELECT id FROM "user" WHERE "user".user_id = $1)
INSERT INTO user_segment (user_id, segment_id)
VALUES 
`
	counter := 2
	for range segmentId {
		sqlStr += fmt.Sprintf(`((SELECT user_name.id FROM user_name),
        (SELECT id FROM segment WHERE segment.name = $%d)),`, counter)
		counter++
	}

	execArgs := []interface{}{userId}
	for i := range segmentId {
		execArgs = append(execArgs, interface{}(segmentId[i]))
	}
	_, err = p.db.Exec(sqlStr[:len(sqlStr)-1], execArgs...)
	if err != nil {
		return fmt.Errorf("!!!can't add segment to user: %w", err)
	}
	return nil
}

func (p *Postgres) ShowSegment(userId int) ([]repository.ShowUsersSegment, error) {
	rows, err := p.db.Query(`SELECT u.user_id, name FROM segment
	INNER JOIN user_segment ON segment.id = user_segment.segment_id
	INNER JOIN public."user" u on user_segment.user_id = u.id
	WHERE u.user_id = $1`, userId)
	if err != nil {
		return nil, fmt.Errorf("can't found user: %w", err)
	}
	defer rows.Close()

	var segments []repository.ShowUsersSegment
	for rows.Next() {
		s := repository.ShowUsersSegment{}
		err = rows.Scan(&s.UserId, &s.Name)
		if err != nil {
			fmt.Println(err)
			continue
		}
		segments = append(segments, s)
	}
	return segments, nil
}
