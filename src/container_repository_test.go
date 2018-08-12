package smith

import (
	"testing"
	"errors"
	"github.com/stretchr/testify/assert"
	"time"
)

func TestRepositoryIsAbleFindModel(t *testing.T) {
	mockClient := new(RedisMock)
	mockClient.result = `{
        "id":"fake-uuid",
        "name":"container_name",
        "image":"container_image",
        "ip":"some_ip",
        "created_at":"2018-08-08T22:00:00.0+02:00",
        "ports":"8080:8080",
        "status":true
    }`
	mockClient.err = errors.New("error")

	containerRepository := NewContainerRepository(mockClient)
	container, _ := containerRepository.FindByID("fake-uuid")

	createdAt, _ := time.Parse(time.RFC3339, "2018-08-08T22:00:00.0+02:00")

	assert.Equal(t, container.ID, "fake-uuid")
	assert.Equal(t, container.Name, "container_name")
	assert.Equal(t, container.Image, "container_image")
	assert.Equal(t, container.Ip, "some_ip")
	assert.Equal(t, container.CreatedAt, createdAt)
	assert.Equal(t, container.Ports, "8080:8080")
	assert.Equal(t, container.Status, true)
}

func TestRepositoryIsAbleToPersistModel(t *testing.T) {
	mockClient := new(RedisMock)
	mockClient.result = "none"
	mockClient.err = nil

	container := new(Container)
	createdAt, _ := time.Parse(time.RFC3339, "2018-08-08T22:00:00.0+02:00")
	container.ID = "fake-uuid"
	container.Name = "container_name"
	container.Ip = "some_ip"
	container.CreatedAt = createdAt
	container.Ports = "8080:8080"
	container.Status = true

	containerRepository := NewContainerRepository(mockClient)
	result := containerRepository.Persist(container)

	assert.Equal(t, result, nil, "Values must be Equal.")
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
