package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCyr2lat(t *testing.T) {
	m := map[string]string{
		"Љубазни фењерџија чађавог лица хоће да ми покаже штос.":    "Ljubazni fenjerdžija čađavog lica hoće da mi pokaže štos.",
		"ЉУБАЗНИ ФЕЊЕРЏИЈА ЧАЂАВОГ ЛИЦА ХОЋЕ ДА МИ ПОКАЖЕ ШТОС.":    "LjUBAZNI FENjERDžIJA ČAĐAVOG LICA HOĆE DA MI POKAŽE ŠTOS.",
		"Дебљој згужвах смеђ филц — њен шкрт џепчић.":               "Debljoj zgužvah smeđ filc — njen škrt džepčić.",
		"Ljubazni fenjerdžija čađavog lica hoće da mi pokaže štos.": "Ljubazni fenjerdžija čađavog lica hoće da mi pokaže štos.",
	}

	for cyrVal, latVal := range m {
		if cyr2lat(cyrVal) != latVal {
			assert.Equal(t, cyr2lat(cyrVal), latVal)
		}
	}
}
