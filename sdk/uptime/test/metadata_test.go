package uptime_test

// import (
// 	"testing"

// 	"github.com/stretchr/testify/assert"

// 	"github.com/dlanderson/betterstack-go/models"
// )

// ! I cannot test this feature because I am using a free account and I can't get the desired responses
// var metadataID string

// func TestListMetadata(t *testing.T) {
// 	metadata, err := bs.ListMetadata()
// 	assert.Nil(t, err)
// 	assert.NotNil(t, metadata)
// 	assert.IsType(t, &models.Metadata{}, metadata)
// }

// func TestCreateMetadataRecord(t *testing.T) {
// 	reqBody := models.MetadataRecordReqBody{

// 	}
// 	heartbeat, err := bs.CreateMetadataRecord(reqBody)
// 	assert.Nil(t, err)
// 	assert.NotNil(t, heartbeat)
// 	assert.IsType(t, &models.MetadataRecord{}, heartbeat)
// 	metadataID = heartbeat.ID
// }

// func TestGetMetadataRecord(t *testing.T) {
// 	heartbeat, err := bs.GetMetadataRecord(metadataID)
// 	assert.Nil(t, err)
// 	assert.NotNil(t, heartbeat)
// 	assert.IsType(t, &models.MetadataRecord{}, heartbeat)
// 	assert.Equal(t, metadataID, heartbeat.ID)
// }

// func TestUpdateMetadataRecord(t *testing.T) {
// 	reqBody := models.MetadataRecordReqBody{

// 	}
// 	heartbeat, err := bs.UpdateMetadataRecord(metadataID, reqBody)
// 	assert.Nil(t, err)
// 	assert.NotNil(t, heartbeat)
// 	assert.IsType(t, &models.MetadataRecord{}, heartbeat)
// 	assert.Equal(t, metadataID, heartbeat.ID)
// }

// func TestDeleteMetadataRecord(t *testing.T) {
// 	err := bs.DeleteMetadataRecord(metadataID)
// 	assert.Nil(t, err)
// }
