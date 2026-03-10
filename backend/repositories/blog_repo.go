package repositories

import (
	"database/sql"
	"fmt"

	"github.com/PanaPapad/Pedestal/backend/models"
)

// define a handler object: this is a struct that provides access to all sql.DB methods, we name it DB
type BlogRepository struct {
	DB *sql.DB
}

// first paranthesis is a method receiver. This method now has access to the above struct
func (r *BlogRepository) CreateBlog(post *models.BlogPost) error {
	/*
		Function that inserts into the datbase a new Blog post.
		Takes as input a reference to a models.BlogPost struct
	*/
	query := `
		INSERT INTO blog_posts (title, slug, content, status)
		VALUES (?, ?, ?, ?)
	`

	fmt.Println(query)

	result, err := r.DB.Exec(query,
		post.Title,
		post.Slug,
		post.Content,
		post.Status,
	)

	if err != nil {
		return err
	}

	post.ID, _ = result.LastInsertId()
	return nil
}

func (r *BlogRepository) GetAllBlogs() ([]models.BlogPost, error) {
	/*

	 */
	rows, err := r.DB.Query(`
		SELECT id, title, slug, content, status, created_at, updated_at
		FROM blog_posts
		WHERE status = 'published'
		ORDER BY created_at DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.BlogPost

	for rows.Next() {
		var p models.BlogPost
		err := rows.Scan(
			&p.ID,
			&p.Title,
			&p.Slug,
			&p.Content,
			&p.Status,
			&p.CreatedAt,
			&p.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, p)
	}

	return posts, nil
}

func (r *BlogRepository) GetOneBlog(slug string) (*models.BlogPost, error) {
	/*
		Function that queries the database to find one blog based on the slug
	*/
	query := `
		SELECT id, title, slug, content, status, created_at, updated_at
		FROM blog_posts
		WHERE slug = ? AND status = 'published'
		LIMIT 1
	`

	var blog_post models.BlogPost

	err := r.DB.QueryRow(query, slug).Scan(
		&blog_post.ID,
		&blog_post.Title,
		&blog_post.Slug,
		&blog_post.Content,
		&blog_post.Status,
		&blog_post.CreatedAt,
		&blog_post.UpdatedAt,
	)

	// this means that we did not find the blog
	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &blog_post, nil
}

func (r *BlogRepository) DeleteBlog(slug string) error {
	/*
		Function that deletes a blog post from the database by its slug
	*/
	query :=
		`
	DELETE FROM blog_posts
	WHERE slug = ?
	`
	_, err := r.DB.Exec(query, slug)

	return err

}
