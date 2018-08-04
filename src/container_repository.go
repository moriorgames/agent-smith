package smith

type ContainerRepository interface {
	FindByID(id string) (string, error)
}

func (c *ContainerRepo) FindByID(id string) (string, error) {
	return c.redis.Get(id)
}

type ContainerRepo struct {
	redis RedisConnectable
}

func NewContainerRepository(redis RedisConnectable) *ContainerRepo {
	containerRepo := ContainerRepo{}
	containerRepo.redis = redis

	return &containerRepo
}
