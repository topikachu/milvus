// Code generated from Plan.g4 by ANTLR 4.9. DO NOT EDIT.

package planparserv2

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 41, 508,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33,
	4, 34, 9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4,
	39, 9, 39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44,
	9, 44, 4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 4, 49, 9,
	49, 4, 50, 9, 50, 4, 51, 9, 51, 4, 52, 9, 52, 4, 53, 9, 53, 4, 54, 9, 54,
	4, 55, 9, 55, 4, 56, 9, 56, 4, 57, 9, 57, 4, 58, 9, 58, 4, 59, 9, 59, 4,
	60, 9, 60, 4, 61, 9, 61, 4, 62, 9, 62, 4, 63, 9, 63, 4, 64, 9, 64, 4, 65,
	9, 65, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7,
	3, 7, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11,
	3, 11, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3,
	13, 3, 13, 5, 13, 166, 10, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14,
	3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 5, 14, 180, 10, 14, 3, 15, 3,
	15, 3, 16, 3, 16, 3, 17, 3, 17, 3, 18, 3, 18, 3, 19, 3, 19, 3, 20, 3, 20,
	3, 20, 3, 21, 3, 21, 3, 21, 3, 22, 3, 22, 3, 22, 3, 23, 3, 23, 3, 24, 3,
	24, 3, 25, 3, 25, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 5, 26, 212, 10, 26,
	3, 27, 3, 27, 3, 27, 3, 27, 5, 27, 218, 10, 27, 3, 28, 3, 28, 3, 29, 3,
	29, 3, 29, 3, 29, 5, 29, 226, 10, 29, 3, 30, 3, 30, 3, 30, 3, 31, 3, 31,
	3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 32, 3, 32, 3, 32, 7, 32, 241, 10,
	32, 12, 32, 14, 32, 244, 11, 32, 3, 32, 3, 32, 3, 33, 3, 33, 3, 33, 3,
	33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33,
	3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3,
	33, 3, 33, 3, 33, 5, 33, 275, 10, 33, 3, 34, 3, 34, 3, 34, 3, 34, 5, 34,
	281, 10, 34, 3, 35, 3, 35, 5, 35, 285, 10, 35, 3, 36, 3, 36, 3, 36, 7,
	36, 290, 10, 36, 12, 36, 14, 36, 293, 11, 36, 3, 36, 3, 36, 3, 36, 3, 36,
	3, 36, 5, 36, 300, 10, 36, 3, 37, 5, 37, 303, 10, 37, 3, 37, 3, 37, 5,
	37, 307, 10, 37, 3, 37, 3, 37, 3, 37, 5, 37, 312, 10, 37, 3, 37, 5, 37,
	315, 10, 37, 3, 38, 3, 38, 3, 38, 3, 38, 5, 38, 321, 10, 38, 3, 38, 3,
	38, 6, 38, 325, 10, 38, 13, 38, 14, 38, 326, 3, 39, 3, 39, 3, 39, 5, 39,
	332, 10, 39, 3, 40, 6, 40, 335, 10, 40, 13, 40, 14, 40, 336, 3, 41, 6,
	41, 340, 10, 41, 13, 41, 14, 41, 341, 3, 42, 3, 42, 3, 42, 3, 42, 3, 42,
	3, 42, 3, 42, 5, 42, 351, 10, 42, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3,
	43, 3, 43, 5, 43, 360, 10, 43, 3, 44, 3, 44, 3, 45, 3, 45, 3, 46, 3, 46,
	3, 46, 6, 46, 369, 10, 46, 13, 46, 14, 46, 370, 3, 47, 3, 47, 7, 47, 375,
	10, 47, 12, 47, 14, 47, 378, 11, 47, 3, 47, 5, 47, 381, 10, 47, 3, 48,
	3, 48, 7, 48, 385, 10, 48, 12, 48, 14, 48, 388, 11, 48, 3, 49, 3, 49, 3,
	49, 3, 49, 3, 50, 3, 50, 3, 51, 3, 51, 3, 52, 3, 52, 3, 53, 3, 53, 3, 53,
	3, 53, 3, 53, 3, 54, 3, 54, 3, 54, 3, 54, 3, 54, 3, 54, 3, 54, 3, 54, 3,
	54, 3, 54, 5, 54, 415, 10, 54, 3, 55, 3, 55, 5, 55, 419, 10, 55, 3, 55,
	3, 55, 3, 55, 5, 55, 424, 10, 55, 3, 56, 3, 56, 3, 56, 3, 56, 5, 56, 430,
	10, 56, 3, 56, 3, 56, 3, 57, 5, 57, 435, 10, 57, 3, 57, 3, 57, 3, 57, 3,
	57, 3, 57, 5, 57, 442, 10, 57, 3, 58, 3, 58, 5, 58, 446, 10, 58, 3, 58,
	3, 58, 3, 59, 6, 59, 451, 10, 59, 13, 59, 14, 59, 452, 3, 60, 5, 60, 456,
	10, 60, 3, 60, 3, 60, 3, 60, 3, 60, 3, 60, 5, 60, 463, 10, 60, 3, 61, 6,
	61, 466, 10, 61, 13, 61, 14, 61, 467, 3, 62, 3, 62, 5, 62, 472, 10, 62,
	3, 62, 3, 62, 3, 63, 3, 63, 3, 63, 3, 63, 3, 63, 5, 63, 481, 10, 63, 3,
	63, 5, 63, 484, 10, 63, 3, 63, 3, 63, 3, 63, 3, 63, 3, 63, 5, 63, 491,
	10, 63, 3, 64, 6, 64, 494, 10, 64, 13, 64, 14, 64, 495, 3, 64, 3, 64, 3,
	65, 3, 65, 5, 65, 502, 10, 65, 3, 65, 5, 65, 505, 10, 65, 3, 65, 3, 65,
	2, 2, 66, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11,
	21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17, 33, 18, 35, 19, 37, 20,
	39, 21, 41, 22, 43, 23, 45, 24, 47, 25, 49, 26, 51, 27, 53, 28, 55, 29,
	57, 30, 59, 31, 61, 32, 63, 33, 65, 34, 67, 35, 69, 36, 71, 37, 73, 38,
	75, 39, 77, 2, 79, 2, 81, 2, 83, 2, 85, 2, 87, 2, 89, 2, 91, 2, 93, 2,
	95, 2, 97, 2, 99, 2, 101, 2, 103, 2, 105, 2, 107, 2, 109, 2, 111, 2, 113,
	2, 115, 2, 117, 2, 119, 2, 121, 2, 123, 2, 125, 2, 127, 40, 129, 41, 3,
	2, 18, 5, 2, 78, 78, 87, 87, 119, 119, 6, 2, 12, 12, 15, 15, 36, 36, 94,
	94, 6, 2, 12, 12, 15, 15, 41, 41, 94, 94, 5, 2, 67, 92, 97, 97, 99, 124,
	3, 2, 50, 59, 4, 2, 68, 68, 100, 100, 3, 2, 50, 51, 4, 2, 90, 90, 122,
	122, 3, 2, 51, 59, 3, 2, 50, 57, 5, 2, 50, 59, 67, 72, 99, 104, 4, 2, 71,
	71, 103, 103, 4, 2, 45, 45, 47, 47, 4, 2, 82, 82, 114, 114, 12, 2, 36,
	36, 41, 41, 65, 65, 94, 94, 99, 100, 104, 104, 112, 112, 116, 116, 118,
	118, 120, 120, 4, 2, 11, 11, 34, 34, 2, 540, 2, 3, 3, 2, 2, 2, 2, 5, 3,
	2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13,
	3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2,
	21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2,
	2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2,
	2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 2, 43, 3, 2,
	2, 2, 2, 45, 3, 2, 2, 2, 2, 47, 3, 2, 2, 2, 2, 49, 3, 2, 2, 2, 2, 51, 3,
	2, 2, 2, 2, 53, 3, 2, 2, 2, 2, 55, 3, 2, 2, 2, 2, 57, 3, 2, 2, 2, 2, 59,
	3, 2, 2, 2, 2, 61, 3, 2, 2, 2, 2, 63, 3, 2, 2, 2, 2, 65, 3, 2, 2, 2, 2,
	67, 3, 2, 2, 2, 2, 69, 3, 2, 2, 2, 2, 71, 3, 2, 2, 2, 2, 73, 3, 2, 2, 2,
	2, 75, 3, 2, 2, 2, 2, 127, 3, 2, 2, 2, 2, 129, 3, 2, 2, 2, 3, 131, 3, 2,
	2, 2, 5, 133, 3, 2, 2, 2, 7, 135, 3, 2, 2, 2, 9, 137, 3, 2, 2, 2, 11, 139,
	3, 2, 2, 2, 13, 141, 3, 2, 2, 2, 15, 143, 3, 2, 2, 2, 17, 146, 3, 2, 2,
	2, 19, 148, 3, 2, 2, 2, 21, 151, 3, 2, 2, 2, 23, 154, 3, 2, 2, 2, 25, 165,
	3, 2, 2, 2, 27, 179, 3, 2, 2, 2, 29, 181, 3, 2, 2, 2, 31, 183, 3, 2, 2,
	2, 33, 185, 3, 2, 2, 2, 35, 187, 3, 2, 2, 2, 37, 189, 3, 2, 2, 2, 39, 191,
	3, 2, 2, 2, 41, 194, 3, 2, 2, 2, 43, 197, 3, 2, 2, 2, 45, 200, 3, 2, 2,
	2, 47, 202, 3, 2, 2, 2, 49, 204, 3, 2, 2, 2, 51, 211, 3, 2, 2, 2, 53, 217,
	3, 2, 2, 2, 55, 219, 3, 2, 2, 2, 57, 225, 3, 2, 2, 2, 59, 227, 3, 2, 2,
	2, 61, 230, 3, 2, 2, 2, 63, 237, 3, 2, 2, 2, 65, 274, 3, 2, 2, 2, 67, 280,
	3, 2, 2, 2, 69, 284, 3, 2, 2, 2, 71, 299, 3, 2, 2, 2, 73, 302, 3, 2, 2,
	2, 75, 316, 3, 2, 2, 2, 77, 331, 3, 2, 2, 2, 79, 334, 3, 2, 2, 2, 81, 339,
	3, 2, 2, 2, 83, 350, 3, 2, 2, 2, 85, 359, 3, 2, 2, 2, 87, 361, 3, 2, 2,
	2, 89, 363, 3, 2, 2, 2, 91, 365, 3, 2, 2, 2, 93, 380, 3, 2, 2, 2, 95, 382,
	3, 2, 2, 2, 97, 389, 3, 2, 2, 2, 99, 393, 3, 2, 2, 2, 101, 395, 3, 2, 2,
	2, 103, 397, 3, 2, 2, 2, 105, 399, 3, 2, 2, 2, 107, 414, 3, 2, 2, 2, 109,
	423, 3, 2, 2, 2, 111, 425, 3, 2, 2, 2, 113, 441, 3, 2, 2, 2, 115, 443,
	3, 2, 2, 2, 117, 450, 3, 2, 2, 2, 119, 462, 3, 2, 2, 2, 121, 465, 3, 2,
	2, 2, 123, 469, 3, 2, 2, 2, 125, 490, 3, 2, 2, 2, 127, 493, 3, 2, 2, 2,
	129, 504, 3, 2, 2, 2, 131, 132, 7, 42, 2, 2, 132, 4, 3, 2, 2, 2, 133, 134,
	7, 43, 2, 2, 134, 6, 3, 2, 2, 2, 135, 136, 7, 93, 2, 2, 136, 8, 3, 2, 2,
	2, 137, 138, 7, 46, 2, 2, 138, 10, 3, 2, 2, 2, 139, 140, 7, 95, 2, 2, 140,
	12, 3, 2, 2, 2, 141, 142, 7, 62, 2, 2, 142, 14, 3, 2, 2, 2, 143, 144, 7,
	62, 2, 2, 144, 145, 7, 63, 2, 2, 145, 16, 3, 2, 2, 2, 146, 147, 7, 64,
	2, 2, 147, 18, 3, 2, 2, 2, 148, 149, 7, 64, 2, 2, 149, 150, 7, 63, 2, 2,
	150, 20, 3, 2, 2, 2, 151, 152, 7, 63, 2, 2, 152, 153, 7, 63, 2, 2, 153,
	22, 3, 2, 2, 2, 154, 155, 7, 35, 2, 2, 155, 156, 7, 63, 2, 2, 156, 24,
	3, 2, 2, 2, 157, 158, 7, 110, 2, 2, 158, 159, 7, 107, 2, 2, 159, 160, 7,
	109, 2, 2, 160, 166, 7, 103, 2, 2, 161, 162, 7, 78, 2, 2, 162, 163, 7,
	75, 2, 2, 163, 164, 7, 77, 2, 2, 164, 166, 7, 71, 2, 2, 165, 157, 3, 2,
	2, 2, 165, 161, 3, 2, 2, 2, 166, 26, 3, 2, 2, 2, 167, 168, 7, 103, 2, 2,
	168, 169, 7, 122, 2, 2, 169, 170, 7, 107, 2, 2, 170, 171, 7, 117, 2, 2,
	171, 172, 7, 118, 2, 2, 172, 180, 7, 117, 2, 2, 173, 174, 7, 71, 2, 2,
	174, 175, 7, 90, 2, 2, 175, 176, 7, 75, 2, 2, 176, 177, 7, 85, 2, 2, 177,
	178, 7, 86, 2, 2, 178, 180, 7, 85, 2, 2, 179, 167, 3, 2, 2, 2, 179, 173,
	3, 2, 2, 2, 180, 28, 3, 2, 2, 2, 181, 182, 7, 45, 2, 2, 182, 30, 3, 2,
	2, 2, 183, 184, 7, 47, 2, 2, 184, 32, 3, 2, 2, 2, 185, 186, 7, 44, 2, 2,
	186, 34, 3, 2, 2, 2, 187, 188, 7, 49, 2, 2, 188, 36, 3, 2, 2, 2, 189, 190,
	7, 39, 2, 2, 190, 38, 3, 2, 2, 2, 191, 192, 7, 44, 2, 2, 192, 193, 7, 44,
	2, 2, 193, 40, 3, 2, 2, 2, 194, 195, 7, 62, 2, 2, 195, 196, 7, 62, 2, 2,
	196, 42, 3, 2, 2, 2, 197, 198, 7, 64, 2, 2, 198, 199, 7, 64, 2, 2, 199,
	44, 3, 2, 2, 2, 200, 201, 7, 40, 2, 2, 201, 46, 3, 2, 2, 2, 202, 203, 7,
	126, 2, 2, 203, 48, 3, 2, 2, 2, 204, 205, 7, 96, 2, 2, 205, 50, 3, 2, 2,
	2, 206, 207, 7, 40, 2, 2, 207, 212, 7, 40, 2, 2, 208, 209, 7, 99, 2, 2,
	209, 210, 7, 112, 2, 2, 210, 212, 7, 102, 2, 2, 211, 206, 3, 2, 2, 2, 211,
	208, 3, 2, 2, 2, 212, 52, 3, 2, 2, 2, 213, 214, 7, 126, 2, 2, 214, 218,
	7, 126, 2, 2, 215, 216, 7, 113, 2, 2, 216, 218, 7, 116, 2, 2, 217, 213,
	3, 2, 2, 2, 217, 215, 3, 2, 2, 2, 218, 54, 3, 2, 2, 2, 219, 220, 7, 128,
	2, 2, 220, 56, 3, 2, 2, 2, 221, 226, 7, 35, 2, 2, 222, 223, 7, 112, 2,
	2, 223, 224, 7, 113, 2, 2, 224, 226, 7, 118, 2, 2, 225, 221, 3, 2, 2, 2,
	225, 222, 3, 2, 2, 2, 226, 58, 3, 2, 2, 2, 227, 228, 7, 107, 2, 2, 228,
	229, 7, 112, 2, 2, 229, 60, 3, 2, 2, 2, 230, 231, 7, 112, 2, 2, 231, 232,
	7, 113, 2, 2, 232, 233, 7, 118, 2, 2, 233, 234, 7, 34, 2, 2, 234, 235,
	7, 107, 2, 2, 235, 236, 7, 112, 2, 2, 236, 62, 3, 2, 2, 2, 237, 242, 7,
	93, 2, 2, 238, 241, 5, 127, 64, 2, 239, 241, 5, 129, 65, 2, 240, 238, 3,
	2, 2, 2, 240, 239, 3, 2, 2, 2, 241, 244, 3, 2, 2, 2, 242, 240, 3, 2, 2,
	2, 242, 243, 3, 2, 2, 2, 243, 245, 3, 2, 2, 2, 244, 242, 3, 2, 2, 2, 245,
	246, 7, 95, 2, 2, 246, 64, 3, 2, 2, 2, 247, 248, 7, 118, 2, 2, 248, 249,
	7, 116, 2, 2, 249, 250, 7, 119, 2, 2, 250, 275, 7, 103, 2, 2, 251, 252,
	7, 86, 2, 2, 252, 253, 7, 116, 2, 2, 253, 254, 7, 119, 2, 2, 254, 275,
	7, 103, 2, 2, 255, 256, 7, 86, 2, 2, 256, 257, 7, 84, 2, 2, 257, 258, 7,
	87, 2, 2, 258, 275, 7, 71, 2, 2, 259, 260, 7, 104, 2, 2, 260, 261, 7, 99,
	2, 2, 261, 262, 7, 110, 2, 2, 262, 263, 7, 117, 2, 2, 263, 275, 7, 103,
	2, 2, 264, 265, 7, 72, 2, 2, 265, 266, 7, 99, 2, 2, 266, 267, 7, 110, 2,
	2, 267, 268, 7, 117, 2, 2, 268, 275, 7, 103, 2, 2, 269, 270, 7, 72, 2,
	2, 270, 271, 7, 67, 2, 2, 271, 272, 7, 78, 2, 2, 272, 273, 7, 85, 2, 2,
	273, 275, 7, 71, 2, 2, 274, 247, 3, 2, 2, 2, 274, 251, 3, 2, 2, 2, 274,
	255, 3, 2, 2, 2, 274, 259, 3, 2, 2, 2, 274, 264, 3, 2, 2, 2, 274, 269,
	3, 2, 2, 2, 275, 66, 3, 2, 2, 2, 276, 281, 5, 93, 47, 2, 277, 281, 5, 95,
	48, 2, 278, 281, 5, 97, 49, 2, 279, 281, 5, 91, 46, 2, 280, 276, 3, 2,
	2, 2, 280, 277, 3, 2, 2, 2, 280, 278, 3, 2, 2, 2, 280, 279, 3, 2, 2, 2,
	281, 68, 3, 2, 2, 2, 282, 285, 5, 109, 55, 2, 283, 285, 5, 111, 56, 2,
	284, 282, 3, 2, 2, 2, 284, 283, 3, 2, 2, 2, 285, 70, 3, 2, 2, 2, 286, 291,
	5, 87, 44, 2, 287, 290, 5, 87, 44, 2, 288, 290, 5, 89, 45, 2, 289, 287,
	3, 2, 2, 2, 289, 288, 3, 2, 2, 2, 290, 293, 3, 2, 2, 2, 291, 289, 3, 2,
	2, 2, 291, 292, 3, 2, 2, 2, 292, 300, 3, 2, 2, 2, 293, 291, 3, 2, 2, 2,
	294, 295, 7, 38, 2, 2, 295, 296, 7, 111, 2, 2, 296, 297, 7, 103, 2, 2,
	297, 298, 7, 118, 2, 2, 298, 300, 7, 99, 2, 2, 299, 286, 3, 2, 2, 2, 299,
	294, 3, 2, 2, 2, 300, 72, 3, 2, 2, 2, 301, 303, 5, 77, 39, 2, 302, 301,
	3, 2, 2, 2, 302, 303, 3, 2, 2, 2, 303, 314, 3, 2, 2, 2, 304, 306, 7, 36,
	2, 2, 305, 307, 5, 79, 40, 2, 306, 305, 3, 2, 2, 2, 306, 307, 3, 2, 2,
	2, 307, 308, 3, 2, 2, 2, 308, 315, 7, 36, 2, 2, 309, 311, 7, 41, 2, 2,
	310, 312, 5, 81, 41, 2, 311, 310, 3, 2, 2, 2, 311, 312, 3, 2, 2, 2, 312,
	313, 3, 2, 2, 2, 313, 315, 7, 41, 2, 2, 314, 304, 3, 2, 2, 2, 314, 309,
	3, 2, 2, 2, 315, 74, 3, 2, 2, 2, 316, 324, 5, 71, 36, 2, 317, 320, 7, 93,
	2, 2, 318, 321, 5, 73, 37, 2, 319, 321, 5, 93, 47, 2, 320, 318, 3, 2, 2,
	2, 320, 319, 3, 2, 2, 2, 321, 322, 3, 2, 2, 2, 322, 323, 7, 95, 2, 2, 323,
	325, 3, 2, 2, 2, 324, 317, 3, 2, 2, 2, 325, 326, 3, 2, 2, 2, 326, 324,
	3, 2, 2, 2, 326, 327, 3, 2, 2, 2, 327, 76, 3, 2, 2, 2, 328, 329, 7, 119,
	2, 2, 329, 332, 7, 58, 2, 2, 330, 332, 9, 2, 2, 2, 331, 328, 3, 2, 2, 2,
	331, 330, 3, 2, 2, 2, 332, 78, 3, 2, 2, 2, 333, 335, 5, 83, 42, 2, 334,
	333, 3, 2, 2, 2, 335, 336, 3, 2, 2, 2, 336, 334, 3, 2, 2, 2, 336, 337,
	3, 2, 2, 2, 337, 80, 3, 2, 2, 2, 338, 340, 5, 85, 43, 2, 339, 338, 3, 2,
	2, 2, 340, 341, 3, 2, 2, 2, 341, 339, 3, 2, 2, 2, 341, 342, 3, 2, 2, 2,
	342, 82, 3, 2, 2, 2, 343, 351, 10, 3, 2, 2, 344, 351, 5, 125, 63, 2, 345,
	346, 7, 94, 2, 2, 346, 351, 7, 12, 2, 2, 347, 348, 7, 94, 2, 2, 348, 349,
	7, 15, 2, 2, 349, 351, 7, 12, 2, 2, 350, 343, 3, 2, 2, 2, 350, 344, 3,
	2, 2, 2, 350, 345, 3, 2, 2, 2, 350, 347, 3, 2, 2, 2, 351, 84, 3, 2, 2,
	2, 352, 360, 10, 4, 2, 2, 353, 360, 5, 125, 63, 2, 354, 355, 7, 94, 2,
	2, 355, 360, 7, 12, 2, 2, 356, 357, 7, 94, 2, 2, 357, 358, 7, 15, 2, 2,
	358, 360, 7, 12, 2, 2, 359, 352, 3, 2, 2, 2, 359, 353, 3, 2, 2, 2, 359,
	354, 3, 2, 2, 2, 359, 356, 3, 2, 2, 2, 360, 86, 3, 2, 2, 2, 361, 362, 9,
	5, 2, 2, 362, 88, 3, 2, 2, 2, 363, 364, 9, 6, 2, 2, 364, 90, 3, 2, 2, 2,
	365, 366, 7, 50, 2, 2, 366, 368, 9, 7, 2, 2, 367, 369, 9, 8, 2, 2, 368,
	367, 3, 2, 2, 2, 369, 370, 3, 2, 2, 2, 370, 368, 3, 2, 2, 2, 370, 371,
	3, 2, 2, 2, 371, 92, 3, 2, 2, 2, 372, 376, 5, 99, 50, 2, 373, 375, 5, 89,
	45, 2, 374, 373, 3, 2, 2, 2, 375, 378, 3, 2, 2, 2, 376, 374, 3, 2, 2, 2,
	376, 377, 3, 2, 2, 2, 377, 381, 3, 2, 2, 2, 378, 376, 3, 2, 2, 2, 379,
	381, 7, 50, 2, 2, 380, 372, 3, 2, 2, 2, 380, 379, 3, 2, 2, 2, 381, 94,
	3, 2, 2, 2, 382, 386, 7, 50, 2, 2, 383, 385, 5, 101, 51, 2, 384, 383, 3,
	2, 2, 2, 385, 388, 3, 2, 2, 2, 386, 384, 3, 2, 2, 2, 386, 387, 3, 2, 2,
	2, 387, 96, 3, 2, 2, 2, 388, 386, 3, 2, 2, 2, 389, 390, 7, 50, 2, 2, 390,
	391, 9, 9, 2, 2, 391, 392, 5, 121, 61, 2, 392, 98, 3, 2, 2, 2, 393, 394,
	9, 10, 2, 2, 394, 100, 3, 2, 2, 2, 395, 396, 9, 11, 2, 2, 396, 102, 3,
	2, 2, 2, 397, 398, 9, 12, 2, 2, 398, 104, 3, 2, 2, 2, 399, 400, 5, 103,
	52, 2, 400, 401, 5, 103, 52, 2, 401, 402, 5, 103, 52, 2, 402, 403, 5, 103,
	52, 2, 403, 106, 3, 2, 2, 2, 404, 405, 7, 94, 2, 2, 405, 406, 7, 119, 2,
	2, 406, 407, 3, 2, 2, 2, 407, 415, 5, 105, 53, 2, 408, 409, 7, 94, 2, 2,
	409, 410, 7, 87, 2, 2, 410, 411, 3, 2, 2, 2, 411, 412, 5, 105, 53, 2, 412,
	413, 5, 105, 53, 2, 413, 415, 3, 2, 2, 2, 414, 404, 3, 2, 2, 2, 414, 408,
	3, 2, 2, 2, 415, 108, 3, 2, 2, 2, 416, 418, 5, 113, 57, 2, 417, 419, 5,
	115, 58, 2, 418, 417, 3, 2, 2, 2, 418, 419, 3, 2, 2, 2, 419, 424, 3, 2,
	2, 2, 420, 421, 5, 117, 59, 2, 421, 422, 5, 115, 58, 2, 422, 424, 3, 2,
	2, 2, 423, 416, 3, 2, 2, 2, 423, 420, 3, 2, 2, 2, 424, 110, 3, 2, 2, 2,
	425, 426, 7, 50, 2, 2, 426, 429, 9, 9, 2, 2, 427, 430, 5, 119, 60, 2, 428,
	430, 5, 121, 61, 2, 429, 427, 3, 2, 2, 2, 429, 428, 3, 2, 2, 2, 430, 431,
	3, 2, 2, 2, 431, 432, 5, 123, 62, 2, 432, 112, 3, 2, 2, 2, 433, 435, 5,
	117, 59, 2, 434, 433, 3, 2, 2, 2, 434, 435, 3, 2, 2, 2, 435, 436, 3, 2,
	2, 2, 436, 437, 7, 48, 2, 2, 437, 442, 5, 117, 59, 2, 438, 439, 5, 117,
	59, 2, 439, 440, 7, 48, 2, 2, 440, 442, 3, 2, 2, 2, 441, 434, 3, 2, 2,
	2, 441, 438, 3, 2, 2, 2, 442, 114, 3, 2, 2, 2, 443, 445, 9, 13, 2, 2, 444,
	446, 9, 14, 2, 2, 445, 444, 3, 2, 2, 2, 445, 446, 3, 2, 2, 2, 446, 447,
	3, 2, 2, 2, 447, 448, 5, 117, 59, 2, 448, 116, 3, 2, 2, 2, 449, 451, 5,
	89, 45, 2, 450, 449, 3, 2, 2, 2, 451, 452, 3, 2, 2, 2, 452, 450, 3, 2,
	2, 2, 452, 453, 3, 2, 2, 2, 453, 118, 3, 2, 2, 2, 454, 456, 5, 121, 61,
	2, 455, 454, 3, 2, 2, 2, 455, 456, 3, 2, 2, 2, 456, 457, 3, 2, 2, 2, 457,
	458, 7, 48, 2, 2, 458, 463, 5, 121, 61, 2, 459, 460, 5, 121, 61, 2, 460,
	461, 7, 48, 2, 2, 461, 463, 3, 2, 2, 2, 462, 455, 3, 2, 2, 2, 462, 459,
	3, 2, 2, 2, 463, 120, 3, 2, 2, 2, 464, 466, 5, 103, 52, 2, 465, 464, 3,
	2, 2, 2, 466, 467, 3, 2, 2, 2, 467, 465, 3, 2, 2, 2, 467, 468, 3, 2, 2,
	2, 468, 122, 3, 2, 2, 2, 469, 471, 9, 15, 2, 2, 470, 472, 9, 14, 2, 2,
	471, 470, 3, 2, 2, 2, 471, 472, 3, 2, 2, 2, 472, 473, 3, 2, 2, 2, 473,
	474, 5, 117, 59, 2, 474, 124, 3, 2, 2, 2, 475, 476, 7, 94, 2, 2, 476, 491,
	9, 16, 2, 2, 477, 478, 7, 94, 2, 2, 478, 480, 5, 101, 51, 2, 479, 481,
	5, 101, 51, 2, 480, 479, 3, 2, 2, 2, 480, 481, 3, 2, 2, 2, 481, 483, 3,
	2, 2, 2, 482, 484, 5, 101, 51, 2, 483, 482, 3, 2, 2, 2, 483, 484, 3, 2,
	2, 2, 484, 491, 3, 2, 2, 2, 485, 486, 7, 94, 2, 2, 486, 487, 7, 122, 2,
	2, 487, 488, 3, 2, 2, 2, 488, 491, 5, 121, 61, 2, 489, 491, 5, 107, 54,
	2, 490, 475, 3, 2, 2, 2, 490, 477, 3, 2, 2, 2, 490, 485, 3, 2, 2, 2, 490,
	489, 3, 2, 2, 2, 491, 126, 3, 2, 2, 2, 492, 494, 9, 17, 2, 2, 493, 492,
	3, 2, 2, 2, 494, 495, 3, 2, 2, 2, 495, 493, 3, 2, 2, 2, 495, 496, 3, 2,
	2, 2, 496, 497, 3, 2, 2, 2, 497, 498, 8, 64, 2, 2, 498, 128, 3, 2, 2, 2,
	499, 501, 7, 15, 2, 2, 500, 502, 7, 12, 2, 2, 501, 500, 3, 2, 2, 2, 501,
	502, 3, 2, 2, 2, 502, 505, 3, 2, 2, 2, 503, 505, 7, 12, 2, 2, 504, 499,
	3, 2, 2, 2, 504, 503, 3, 2, 2, 2, 505, 506, 3, 2, 2, 2, 506, 507, 8, 65,
	2, 2, 507, 130, 3, 2, 2, 2, 49, 2, 165, 179, 211, 217, 225, 240, 242, 274,
	280, 284, 289, 291, 299, 302, 306, 311, 314, 320, 326, 331, 336, 341, 350,
	359, 370, 376, 380, 386, 414, 418, 423, 429, 434, 441, 445, 452, 455, 462,
	467, 471, 480, 483, 490, 495, 501, 504, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'('", "')'", "'['", "','", "']'", "'<'", "'<='", "'>'", "'>='", "'=='",
	"'!='", "", "", "'+'", "'-'", "'*'", "'/'", "'%'", "'**'", "'<<'", "'>>'",
	"'&'", "'|'", "'^'", "", "", "'~'", "", "'in'", "'not in'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "LT", "LE", "GT", "GE", "EQ", "NE", "LIKE", "EXISTS",
	"ADD", "SUB", "MUL", "DIV", "MOD", "POW", "SHL", "SHR", "BAND", "BOR",
	"BXOR", "AND", "OR", "BNOT", "NOT", "IN", "NIN", "EmptyTerm", "BooleanConstant",
	"IntegerConstant", "FloatingConstant", "Identifier", "StringLiteral", "JSONIdentifier",
	"Whitespace", "Newline",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "LT", "LE", "GT", "GE", "EQ", "NE",
	"LIKE", "EXISTS", "ADD", "SUB", "MUL", "DIV", "MOD", "POW", "SHL", "SHR",
	"BAND", "BOR", "BXOR", "AND", "OR", "BNOT", "NOT", "IN", "NIN", "EmptyTerm",
	"BooleanConstant", "IntegerConstant", "FloatingConstant", "Identifier",
	"StringLiteral", "JSONIdentifier", "EncodingPrefix", "DoubleSCharSequence",
	"SingleSCharSequence", "DoubleSChar", "SingleSChar", "Nondigit", "Digit",
	"BinaryConstant", "DecimalConstant", "OctalConstant", "HexadecimalConstant",
	"NonzeroDigit", "OctalDigit", "HexadecimalDigit", "HexQuad", "UniversalCharacterName",
	"DecimalFloatingConstant", "HexadecimalFloatingConstant", "FractionalConstant",
	"ExponentPart", "DigitSequence", "HexadecimalFractionalConstant", "HexadecimalDigitSequence",
	"BinaryExponentPart", "EscapeSequence", "Whitespace", "Newline",
}

type PlanLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewPlanLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *PlanLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewPlanLexer(input antlr.CharStream) *PlanLexer {
	l := new(PlanLexer)
	lexerDeserializer := antlr.NewATNDeserializer(nil)
	lexerAtn := lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)
	lexerDecisionToDFA := make([]*antlr.DFA, len(lexerAtn.DecisionToState))
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Plan.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// PlanLexer tokens.
const (
	PlanLexerT__0             = 1
	PlanLexerT__1             = 2
	PlanLexerT__2             = 3
	PlanLexerT__3             = 4
	PlanLexerT__4             = 5
	PlanLexerLT               = 6
	PlanLexerLE               = 7
	PlanLexerGT               = 8
	PlanLexerGE               = 9
	PlanLexerEQ               = 10
	PlanLexerNE               = 11
	PlanLexerLIKE             = 12
	PlanLexerEXISTS           = 13
	PlanLexerADD              = 14
	PlanLexerSUB              = 15
	PlanLexerMUL              = 16
	PlanLexerDIV              = 17
	PlanLexerMOD              = 18
	PlanLexerPOW              = 19
	PlanLexerSHL              = 20
	PlanLexerSHR              = 21
	PlanLexerBAND             = 22
	PlanLexerBOR              = 23
	PlanLexerBXOR             = 24
	PlanLexerAND              = 25
	PlanLexerOR               = 26
	PlanLexerBNOT             = 27
	PlanLexerNOT              = 28
	PlanLexerIN               = 29
	PlanLexerNIN              = 30
	PlanLexerEmptyTerm        = 31
	PlanLexerBooleanConstant  = 32
	PlanLexerIntegerConstant  = 33
	PlanLexerFloatingConstant = 34
	PlanLexerIdentifier       = 35
	PlanLexerStringLiteral    = 36
	PlanLexerJSONIdentifier   = 37
	PlanLexerWhitespace       = 38
	PlanLexerNewline          = 39
)
