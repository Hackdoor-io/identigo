package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringToUint32Success(t *testing.T) {
	expected := []uint32{
		1320097209,
		728769438,
		3408452011,
		2797316833,
	}

	actual := []uint32{
		stringToUint32("username"),
		stringToUint32("mitch"),
		stringToUint32("hello_world!"),
		stringToUint32("196b0f14eba66e10fba74dbf9e99c22f"),
	}

	for i := 0; i < len(expected); i++ {
		assert.Equal(t, expected[i], actual[i])
	}
}

func TestStringToUint32CorrectFailure(t *testing.T) {
	expected := []uint32{
		1320097209,
		728769438,
		3408452011,
		2797316833,
	}

	actual := []uint32{
		stringToUint32("_"),
		stringToUint32("linux"),
		stringToUint32("goodbye_world?"),
		stringToUint32("7815696ecbf1c96e6894b779456d330e"),
	}

	for i := 0; i < len(expected); i++ {
		assert.NotEqual(t, expected[i], actual[i])
	}
}

func TestGenerateRGBFromStringSuccess(t *testing.T) {

	expected := []RGBColor{
		RGBColor{38, 130, 40},
		RGBColor{130, 82, 6},
		RGBColor{33, 136, 98},
	}

	actual := []RGBColor{
		generateRGBFromString("fae53351b9effc708e764e871bef3119", 2),
		generateRGBFromString("fae53351b9effc708e764e871bef3119", 4),
		generateRGBFromString("4e3ef92ed95e0776ff69797b475ccd58", 1),
	}

	for i := 0; i < len(expected); i++ {
		expectedRed := expected[i].R
		actualRed := actual[i].R

		expectedGreen := expected[i].G
		actualGreen := actual[i].G

		expectedBlue := expected[i].B
		actualBlue := actual[i].B

		assert.Equal(t, expectedRed, actualRed)
		assert.Equal(t, expectedGreen, actualGreen)
		assert.Equal(t, expectedBlue, actualBlue)
	}

}

func TestGenerateRGBFromStringFailure(t *testing.T) {

	expected := []RGBColor{
		RGBColor{32, 130, 40},
		RGBColor{120, 82, 6},
		RGBColor{33, 131, 98},
	}

	actual := []RGBColor{
		generateRGBFromString("fae5d351b9effc708e764e871bef3119", 5),
		generateRGBFromString("fae533a1b9effc708e764e871bef3119", 3),
		generateRGBFromString("4e3ef9aed95e0776ff69797b475ccd58", 2),
	}

	for i := 0; i < len(expected); i++ {
		expectedRed := expected[i].R
		actualRed := actual[i].R

		expectedGreen := expected[i].G
		actualGreen := actual[i].G

		expectedBlue := expected[i].B
		actualBlue := actual[i].B

		assert.NotEqual(t, expectedRed, actualRed)
		assert.NotEqual(t, expectedGreen, actualGreen)
		assert.NotEqual(t, expectedBlue, actualBlue)
	}

}
