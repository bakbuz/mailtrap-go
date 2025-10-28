package emails_test

import (
	"testing"

	"github.com/bakbuz/mailtrap-go/core/common"
	"github.com/bakbuz/mailtrap-go/emails/model"
	"github.com/stretchr/testify/assert"
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
	attachment := model.NewAttachment(_testContent, _testFileName, "", common.MimeType_TextPlain, "")

	result := attachment.Validate()

	assert.Nil(t, result)
}

func TestValidation_ShouldNotFail_WhenMimeTypeLessThan3(t *testing.T) {
	attachment := model.NewAttachment(_testContent, _testFileName, "", "AB", "")

	result := attachment.Validate()

	assert.NotNil(t, result)
}
