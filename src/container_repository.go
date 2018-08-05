package smith

import (
	"encoding/json"
)

type ContainerRepository interface {
	FindByID(id string) (*Container, error)
}

func (c *ContainerRepo) FindByID(id string) (*Container, error) {
	data, err := c.redis.Get(id)
	bytes := []byte(data)
	container := new(Container)

	json.Unmarshal(bytes, &container)

	return container, err
}

type ContainerRepo struct {
	redis RedisConnectable
}

func NewContainerRepository(redis RedisConnectable) *ContainerRepo {
	containerRepo := ContainerRepo{}
	containerRepo.redis = redis

	return &containerRepo
}
