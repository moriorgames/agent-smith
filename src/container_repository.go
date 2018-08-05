package smith

import (
	"encoding/json"
)

type ContainerRepository interface {
	FindByID(id string) (*Container, error)
	Persist(*Container) (error)
}

func (c *ContainerRepo) FindByID(id string) (*Container, error) {
	data, err := c.redis.Get(id)
	bytes := []byte(data)
	container := new(Container)
	json.Unmarshal(bytes, &container)

	return container, err
}

func (c *ContainerRepo) Persist(container *Container) (error) {
	bytes, err := json.Marshal(container)
	if err != nil {
		return err
	}

	_, err = c.redis.Set(container.ID, string(bytes), 0)
	if err != nil {
		return err
	}

	return nil
}

type ContainerRepo struct {
	redis RedisConnectable
}

func NewContainerRepository(redis RedisConnectable) *ContainerRepo {
	containerRepo := ContainerRepo{}
	containerRepo.redis = redis

	return &containerRepo
}
