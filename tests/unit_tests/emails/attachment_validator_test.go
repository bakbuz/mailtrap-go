package emails_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/maydere/mailtrap-go/core/common"
	"gitlab.com/maydere/mailtrap-go/emails/model"
)

const (
	_testContent  string = "Test content"
	_testFileName string = "test.txt"
)

func TestValidation_ShouldNotFail_WhenContentIsNotNullOrEmpty(t *testing.T) {
	attachment := &model.Attachment{Content: _testContent, FileName: _testFileName}

	result := attachment.Validate()

	assert.Nil(t, result)
}

func TestValidation_ShouldNotFail_WhenFileNameIsNotNullOrEmpty(t *testing.T) {
	attachment := model.NewAttachment(_testContent, _testFileName, common.DispositionType_Attachment, "", "")

	result := attachment.Validate()

	assert.Nil(t, result)
}

// func TestValidation_ShouldNotFail_WhenMimeTypeIsNull(t *testing.T) {
// 	attachment := model.NewAttachment(_testContent, _testFileName, nil, "", "")
// 	result := attachment.Validate()
// 	assert.Nil(t, result)
// }

func TestValidation_ShouldFail_WhenMimeTypeIsEmpty(t *testing.T) {
	attachment := model.NewAttachment(_testContent, _testFileName, "", "", "")

	result := attachment.Validate()

	assert.Nil(t, result)
}

func TestValidation_ShouldNotFail_WhenMimeTypeIsNotNullOrEmpty(t *testing.T) {
	attachment := model.NewAttachment(_testContent, _testFileName, "", common.MediaType_TextPlain, "")

	result := attachment.Validate()

	assert.Nil(t, result)
}

func TestValidation_ShouldNotFail_WhenMimeTypeLessThan3(t *testing.T) {
	attachment := model.NewAttachment(_testContent, _testFileName, "", "AB", "")

	result := attachment.Validate()

	assert.NotNil(t, result)
}
