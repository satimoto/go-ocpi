package evid

import (
	"regexp"
	"strings"
)

type EvIdValidator interface {
	ValueWithSeparator(separator string) string
	Value() string
	IsValid() bool
}

type EvIdValidate struct {
	evId     string
	checksum string
}

func NewEvid(evId string) EvIdValidator {
	replaceSpacingRegex := regexp.MustCompile(`[-\*]`)
	validate := &EvIdValidate{
		evId:     strings.ToUpper(replaceSpacingRegex.ReplaceAllString(evId, "")),
		checksum: "",
	}
	validate.calculateChecksum()

	return validate
}

var matrix = map[string][]int{
	"0": {0, 0, 0, 0, 0},
	"1": {0, 0, 0, 1, 16},
	"2": {0, 0, 0, 2, 32},
	"3": {0, 0, 1, 0, 4},
	"4": {0, 0, 1, 1, 20},
	"5": {0, 0, 1, 2, 36},
	"6": {0, 0, 2, 0, 8},
	"7": {0, 0, 2, 1, 24},
	"8": {0, 0, 2, 2, 40},
	"9": {0, 1, 0, 0, 2},
	"A": {0, 1, 0, 1, 18},
	"B": {0, 1, 0, 2, 34},
	"C": {0, 1, 1, 0, 6},
	"D": {0, 1, 1, 1, 22},
	"E": {0, 1, 1, 2, 38},
	"F": {0, 1, 2, 0, 10},
	"G": {0, 1, 2, 1, 26},
	"H": {0, 1, 2, 2, 42},
	"I": {1, 0, 0, 0, 1},
	"J": {1, 0, 0, 1, 17},
	"K": {1, 0, 0, 2, 33},
	"L": {1, 0, 1, 0, 5},
	"M": {1, 0, 1, 1, 21},
	"N": {1, 0, 1, 2, 37},
	"O": {1, 0, 2, 0, 9},
	"P": {1, 0, 2, 1, 25},
	"Q": {1, 0, 2, 2, 41},
	"R": {1, 1, 0, 0, 3},
	"S": {1, 1, 0, 1, 19},
	"T": {1, 1, 0, 2, 35},
	"U": {1, 1, 1, 0, 7},
	"V": {1, 1, 1, 1, 23},
	"W": {1, 1, 1, 2, 39},
	"X": {1, 1, 2, 0, 11},
	"Y": {1, 1, 2, 1, 27},
	"Z": {1, 1, 2, 2, 43},
}

var p1 = [][]int{
	{0, 1, 1, 1},
	{1, 1, 1, 0},
	{1, 0, 0, 1},
	{0, 1, 1, 1},
	{1, 1, 1, 0},
	{1, 0, 0, 1},
	{0, 1, 1, 1},
	{1, 1, 1, 0},
	{1, 0, 0, 1},
	{0, 1, 1, 1},
	{1, 1, 1, 0},
	{1, 0, 0, 1},
	{0, 1, 1, 1},
	{1, 1, 1, 0},
	{1, 0, 0, 1},
}

var p2 = [][]int{
	{0, 1, 1, 2},
	{1, 2, 2, 2},
	{2, 2, 2, 0},
	{2, 0, 0, 2},
	{0, 2, 2, 1},
	{2, 1, 1, 1},
	{1, 1, 1, 0},
	{1, 0, 0, 1},
	{0, 1, 1, 2},
	{1, 2, 2, 2},
	{2, 2, 2, 0},
	{2, 0, 0, 2},
	{0, 2, 2, 1},
	{2, 1, 1, 1},
	{1, 1, 1, 0},
}

func (v *EvIdValidate) IsValid() bool {
	if !v.isValid() {
		return false
	} else if len(v.evId) < 15 {
		return true
	}

	return string(v.evId[14]) == v.checksum
}

func (v *EvIdValidate) isValid() bool {
	validationRegex := regexp.MustCompile(`^([A-Za-z]{2})([A-Za-z0-9]{3})([A-Za-z0-9]{9})([A-Za-z0-9]?)$`)
	return validationRegex.MatchString(v.evId)
}

func (v *EvIdValidate) calculateChecksum() {
	if v.isValid() {
		var matrixValues []int

		for i := 0; i < 14; i++ {
			for k := 0; k < 4; k++ {
				matrixValues = append(matrixValues, matrix[string(v.evId[i])][k])
			}
		}

		var c1, c2, c3, c4 = 0, 0, 0, 0
		for i := 0; i < 14; i++ {
			c1 += matrixValues[i*4]*p1[i][0] + matrixValues[i*4+1]*p1[i][2]
			c2 += matrixValues[i*4]*p1[i][1] + matrixValues[i*4+1]*p1[i][3]
			c3 += matrixValues[i*4+2]*p2[i][0] + matrixValues[i*4+3]*p2[i][2]
			c4 += matrixValues[i*4+2]*p2[i][1] + matrixValues[i*4+3]*p2[i][3]
		}

		c1 = c1 % 2
		c2 = c2 % 2
		c3 = c3 % 3
		c4 = c4 % 3

		var q1, q2, r1, r2 = c1, c2, 0, 0

		switch c4 {
		case 0:
			r1 = 0
		case 1:
			r1 = 2
		case 2:
			r1 = 1
		}
		switch c3 + r1 {
		case 0:
			r2 = 0
		case 1:
			r2 = 2
		case 2:
			r2 = 1
		case 3:
			r2 = 0
		case 4:
			r2 = 2
		}

		var digit = q1 + q2*2 + r1*4 + r2*16

		for key, val := range matrix {
			if val[4] == digit {
				v.checksum = key
				break
			}
		}
	}
}

func (v *EvIdValidate) Value() string {
	if !v.isValid() {
		return ""
	} else if len(v.evId) < 15 {
		return v.evId + v.checksum
	}

	return v.evId
}

func (v *EvIdValidate) ValueWithSeparator(separator string) string {
	var evId = v.Value()

	if len(evId) == 0 {
		return evId
	}

	return evId[0:2] + separator + evId[2:5] + separator + evId[5:14] + separator + evId[14:15]
}
