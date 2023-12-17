package datastore

import (
	"library/model"

	"gofr.dev/pkg/gofr"
)

type Book interface {
	// GetByID retrieves a Book record based on its ID.
	GetByID(ctx *gofr.Context, id int) (*model.Book, error)
	// Create inserts a new Book record into the database.
	Create(ctx *gofr.Context, model *model.Book) (*model.Book, error)
	// Update updates an existing Book record with the provided information.
	Update(ctx *gofr.Context, model *model.Book) (*model.Book, error)
	// Delete removes a Book record from the database based on its ID.
	Delete(ctx *gofr.Context, id int) error
}