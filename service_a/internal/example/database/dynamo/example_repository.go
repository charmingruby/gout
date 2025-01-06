package dynamo

import (
	"errors"

	"github.com/charmingruby/gout-service-a/internal/example/core/model"
)

type ExampleRepository struct {
}

func NewExampleRepository() *ExampleRepository {
	return &ExampleRepository{}
}

func (r *ExampleRepository) Store(model *model.Example) error {
	return errors.New("unimplemented method")
}

func (r *ExampleRepository) FindByID(id string) (*model.Example, error) {
	return nil, errors.New("unimplemented method")
}

func (r *ExampleRepository) Save(model *model.Example) error {
	return errors.New("unimplemented method")
}

func (r *ExampleRepository) Delete(model *model.Example) error {
	return errors.New("unimplemented method")
}
