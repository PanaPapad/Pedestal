package repositories

import (
	"database/sql"

	"github.com/PanaPapad/Pedestal/backend/models"
)

type PodcastRepository struct {
	DB *sql.DB
}

func (r *PodcastRepository) GetAllPodcasts() ([]models.PodcastEpisode, error) {

	query := `
		SELECT id, title, slug, description, audio_url, duration, created_at, updated_at
		FROM podcasts
		ORDER BY created_at DESC
	`
	rows, err := r.DB.Query(query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var podcasts []models.PodcastEpisode

	for rows.Next() {
		var p models.PodcastEpisode
		err := rows.Scan(
			&p.ID,
			&p.Title,
			&p.Slug,
			&p.Description,
			&p.AudioURL,
			&p.Duration,
			&p.CreatedAt,
			&p.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		podcasts = append(podcasts, p)
	}
	return podcasts, nil
}

func (r *PodcastRepository) CreatePodcast(podcast *models.PodcastEpisode) error {
	/*
		Function that inserts into the datbase a new Podcast .
		Takes as input a reference to a models.PodcastEpisode struct
	*/

	query := `
		INSERT INTO podcasts (title, slug, description, audio_url, duration)
		VALUES (?, ?, ?, ?, ?)
	`

	results, err := r.DB.Exec(query,
		podcast.Title,
		podcast.Slug,
		podcast.Description,
		podcast.AudioURL,
		podcast.Duration,
	)

	if err != nil {
		return err
	}

	podcast.ID, _ = results.LastInsertId()
	return nil

}

func (r *PodcastRepository) GetOnePodcast(slug string) (*models.PodcastEpisode, error) {

	query := `
		SELECT id, title, slug, description, duration, audio_url, created_at, updated_at
		FROM podcasts
		WHERE slug = ? 
		LIMIT 1
	`
	var podcast models.PodcastEpisode
	err := r.DB.QueryRow(query, slug).Scan(
		&podcast.ID,
		&podcast.Title,
		&podcast.Slug,
		&podcast.Description,
		&podcast.Duration,
		&podcast.AudioURL,
		&podcast.CreatedAt,
		&podcast.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &podcast, nil

}

func (r *PodcastRepository) DeletePodcast(slug string) error {
	/*
		Function that deletes a podcast from the database by its slug
	*/
	query :=
		`
	DELETE FROM podcasts
	WHERE slug = ?
	`
	_, err := r.DB.Exec(query, slug)

	return err

}
