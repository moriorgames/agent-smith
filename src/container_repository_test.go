package smith

import (
	"testing"
	"errors"
	"github.com/stretchr/testify/assert"
	"time"
)

func TestRepositoryIsAbleToPersistModel(t *testing.T) {
	mockClient := new(RedisMock)
	mockClient.result = `{
        "id":"fake-uuid",
        "name":"container_name",
        "ip":"some_ip",
        "created_at":"2018-08-08T22:00:00.0+02:00",
        "ports":"8080:8080",
        "status":true
    }`
	mockClient.err = errors.New("error")

	containerRepository := NewContainerRepository(mockClient)
	container, _ := containerRepository.FindByID("fake-uuid")

	createdAt, _ := time.Parse(time.RFC3339, "2018-08-08T22:00:00.0+02:00")

	assert.Equal(t, container.ID, "fake-uuid", "Values must be Equal.")
	assert.Equal(t, container.Name, "container_name", "Values must be Equal.")
	assert.Equal(t, container.Ip, "some_ip", "Values must be Equal.")
	assert.Equal(t, container.CreatedAt, createdAt, "Values must be Equal.")
	assert.Equal(t, container.Ports, "8080:8080", "Values must be Equal.")
	assert.Equal(t, container.Status, true, "Values must be Equal.")
}
