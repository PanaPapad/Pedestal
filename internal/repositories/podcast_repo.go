package repositories

import (
	"database/sql"

	"github.com/PanaPapad/Pedestal/internal/models"
)

type PodcastRepository struct {
	DB *sql.DB
}

func (r *PodcastRepository) GetAllPodcasts() ([]models.PodcastEpisode, error) {

	query := `
		SELECT id, title, description audio_url, duration, publishedAt, createdAt, updatedAt
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
			&p.Description,
			&p.AudioURL,
			&p.Duration,
			&p.PublishedAt,
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
