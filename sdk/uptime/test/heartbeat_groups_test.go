package uptime_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/dlanderson/betterstack-go/models"
)

var heartbeatGroupID string

func TestListHeartbeatGroups(t *testing.T) {
	heartbeatGroups, err := bs.ListHeartbeatGroups()
	assert.Nil(t, err)
	assert.NotNil(t, heartbeatGroups)
	assert.IsType(t, &models.Groups{}, heartbeatGroups)
}

func TestCreateHeartbeatGroup(t *testing.T) {
	reqBody := models.GroupReqBody{
		Name: "heartbeatGroup_test",
	}
	heartbeatGroup, err := bs.CreateHeartbeatGroup(reqBody)
	assert.Nil(t, err)
	assert.NotNil(t, heartbeatGroup)
	assert.IsType(t, &models.Group{}, heartbeatGroup)
	assert.Equal(t, reqBody.Name, heartbeatGroup.Attributes.Name)
	heartbeatGroupID = heartbeatGroup.ID
}

func TestGetHeartbeatGroup(t *testing.T) {
	heartbeatGroup, err := bs.GetHeartbeatGroup(heartbeatGroupID)
	assert.Nil(t, err)
	assert.NotNil(t, heartbeatGroup)
	assert.IsType(t, &models.Group{}, heartbeatGroup)
	assert.Equal(t, heartbeatGroupID, heartbeatGroup.ID)
}

func TestUpdateHeartbeatGroup(t *testing.T) {
	reqBody := models.GroupReqBody{
		Name:      "heartbeatGroup_test_update",
		SortIndex: &[]int{1}[0],
	}
	heartbeatGroup, err := bs.UpdateHeartbeatGroup(heartbeatGroupID, reqBody)
	assert.Nil(t, err)
	assert.NotNil(t, heartbeatGroup)
	assert.IsType(t, &models.Group{}, heartbeatGroup)
	assert.Equal(t, heartbeatGroupID, heartbeatGroup.ID)
	assert.Equal(t, reqBody.Name, heartbeatGroup.Attributes.Name)
	assert.Equal(t, *reqBody.SortIndex, heartbeatGroup.Attributes.SortIndex)
}

func TestDeleteHeartbeatGroup(t *testing.T) {
	err := bs.DeleteHeartbeatGroup(heartbeatGroupID)
	assert.Nil(t, err)
}
