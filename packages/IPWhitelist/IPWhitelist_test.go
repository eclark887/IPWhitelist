package IPWhitelist

import (
	"testing"
	"github.com/stretchr/testify/assert"

)

func TestGetRecordFromIPForIP(t *testing.T) {
	record, err := GetRecordFromIP("81.2.69.142")
	assert.Nil(t, err)
	t.Log(record.Country.Names)

	assert.Equal(t, record.Country.Names["en"], "United Kingdom")
	assert.Equal(t, record.Country.Names["de"], "Vereinigtes Königreich")
	assert.Equal(t, record.Country.Names["es"], "Reino Unido")
	assert.Equal(t, record.Country.Names["fr"], "Royaume-Uni")
	assert.Equal(t, record.Country.Names["ja"], "イギリス")
	assert.Equal(t, record.Country.Names["pt-BR"], "Reino Unido")
	assert.Equal(t, record.Country.Names["ru"], "Великобритания")
	assert.Equal(t, record.Country.Names["zh-CN"], "英国")
}

func TestGetCountryNameFromRecord(t *testing.T) {
	record, err := GetRecordFromIP("81.2.69.142")
	assert.Nil(t, err)

	name, err := GetCountryNameFromRecord(record, "en")
	assert.Nil(t, err)
	assert.Equal(t, name, "United Kingdom")

	name, err = GetCountryNameFromRecord(record, "de")
	assert.Nil(t, err)
	assert.Equal(t, name, "Vereinigtes Königreich")

	name, err = GetCountryNameFromRecord(record, "es")
	assert.Nil(t, err)
	assert.Equal(t, name, "Reino Unido")

	name, err = GetCountryNameFromRecord(record, "fr")
	assert.Nil(t, err)
	assert.Equal(t, name, "Royaume-Uni")

	name, err = GetCountryNameFromRecord(record, "ja")
	assert.Nil(t, err)
	assert.Equal(t, name, "イギリス")

	name, err = GetCountryNameFromRecord(record, "pt-BR")
	assert.Equal(t, name, "")
	assert.Equal(t, err, LocaleErr)

	name, err = GetCountryNameFromRecord(record, "err")
	assert.Equal(t, name, "")
	assert.Equal(t, err, LocaleErr)
}

func TestGetISOFromRecord(t *testing.T) {
	record, err := GetRecordFromIP("81.2.69.142")
	assert.Nil(t, err)
	iso := GetISOFromRecord(record)

	assert.Equal(t, iso, "GB")
}

func TestVerifyAllowedLocale(t *testing.T) {
	allowed := VerifyAllowedLocale("what")
	assert.False(t, allowed)

	allowed = VerifyAllowedLocale("en")
	assert.True(t, allowed)
}