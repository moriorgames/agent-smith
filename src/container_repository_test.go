package smith

import (
	"errors"
	"testing"
	"time"
	"github.com/stretchr/testify/assert"
)

func TestRepositoryIsAbleFindModel(t *testing.T) {
	mockClient := new(RedisMock)
	mockClient.result = `{
        "id":"fake-uuid",
        "name":"container_name",
        "image":"container_image",
        "created_at":"2018-08-08T22:00:00.0+02:00",
        "ports":"8080:8080"
    }`
	mockClient.err = errors.New("error")

	containerRepository := NewContainerRepository(mockClient)
	container, _ := containerRepository.FindByID("fake-uuid")

	createdAt, _ := time.Parse(time.RFC3339, "2018-08-08T22:00:00.0+02:00")

	assert.Equal(t, "fake-uuid", container.ID)
	assert.Equal(t, "container_name", container.Name)
	assert.Equal(t, "container_image", container.Image)
	assert.Equal(t, createdAt, container.CreatedAt)
	assert.Equal(t, "8080:8080", container.Ports)
}

func TestRepositoryIsAbleToPersistModel(t *testing.T) {
	mockClient := new(RedisMock)
	mockClient.result = "OK"
	mockClient.err = nil

	container := new(Container)
	createdAt, _ := time.Parse(time.RFC3339, "2018-08-08T22:00:00.0+02:00")
	container.ID = "fake-uuid"
	container.Name = "container_name"
	container.Image = "container_image"
	container.CreatedAt = createdAt
	container.Ports = "8080:8080"

	containerRepository := NewContainerRepository(mockClient)
	result := containerRepository.Persist(container)

	assert.Equal(t, nil, result)
}

func TestRepositoryIsNotAbleToPersistModel(t *testing.T) {
	mockClient := new(RedisMock)
	mockClient.result = "none"
	mockClient.err = errors.New("error")

	container := new(Container)

	containerRepository := NewContainerRepository(mockClient)
	result := containerRepository.Persist(container)

	assert.Equal(t, result, errors.New("error"), "Values must be Equal.")
}

func TestRepositoryIsAbleToPersistModelOnProduction(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	redisClient := ConnectToRedis()

	container := new(Container)
	container.ID = "prod-fake-uuid"
	container.Name = "prod_container_name"
	container.Image = "prod_container_image"
	container.CreatedAt = time.Now()
	container.Ports = "443:443"

	containerRepository := NewContainerRepository(redisClient)
	result := containerRepository.Persist(container)

	assert.Equal(t, nil, result)
}

func TestRepositoryIsAbleFindModelOnProduction(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	redisClient := ConnectToRedis()

	containerRepository := NewContainerRepository(redisClient)
	container, _ := containerRepository.FindByID("prod-fake-uuid")

	assert.Equal(t, "prod-fake-uuid", container.ID)
	assert.Equal(t, "prod_container_name", container.Name)
	assert.Equal(t, "prod_container_image", container.Image)
	assert.Equal(t, "443:443", container.Ports)
}
