package constants

// op codes
// nolint golint
const (
	OP_IDENTITY = iota + 1
    OP_JMP
    OP_GOTO
    OP_DEBUG

	OP_SERIALIZE
	OP_DESERIALIZE

	//TODO: remove undeclared type, this is stupid
	START_OF_OPERATORS //WHAT KIND OF "OPERATORS"
	START_OF_COMPARISON_OPERATORS = iota + START_OF_OPERATORS
	OP_EQUAL
	OP_UNEQUAL
	OP_LT
	OP_GT
	OP_LTEQ
	OP_GTEQ
	END_OF_COMPARISON_OPERATORS
	//TODO: WTF?
	START_OF_ARITHMETIC_OPERATORS = iota + END_OF_COMPARISON_OPERATORS
	OP_BITAND
	OP_BITOR
	OP_BITXOR
	OP_BITCLEAR
	OP_BITSHL
	OP_BITSHR
	OP_MUL
	OP_DIV
	OP_MOD
	OP_ADD
	OP_SUB
	OP_NEG
	END_OF_ARITHMETIC_OPERATORS
	END_OF_OPERATORS = iota + END_OF_ARITHMETIC_OPERATORS
	//TODO: Why this start/end for operators?

	OP_UND_LEN
	OP_UND_PRINTF
	OP_UND_SPRINTF
	OP_UND_READ

	START_PARSE_OPS

	OP_I8_STR
	OP_I8_I16
	OP_I8_I32
	OP_I8_I64
	OP_I8_UI8
	OP_I8_UI16
	OP_I8_UI32
	OP_I8_UI64
	OP_I8_F32
	OP_I8_F64

	OP_I16_STR
	OP_I16_I8
	OP_I16_I32
	OP_I16_I64
	OP_I16_UI8
	OP_I16_UI16
	OP_I16_UI32
	OP_I16_UI64
	OP_I16_F32
	OP_I16_F64

	OP_I32_STR
	OP_I32_I8
	OP_I32_I16
	OP_I32_I64
	OP_I32_UI8
	OP_I32_UI16
	OP_I32_UI32
	OP_I32_UI64
	OP_I32_F32
	OP_I32_F64

	OP_I64_STR
	OP_I64_I8
	OP_I64_I16
	OP_I64_I32
	OP_I64_UI8
	OP_I64_UI16
	OP_I64_UI32
	OP_I64_UI64
	OP_I64_F32
	OP_I64_F64

	OP_UI8_STR
	OP_UI8_I8
	OP_UI8_I16
	OP_UI8_I32
	OP_UI8_I64
	OP_UI8_UI16
	OP_UI8_UI32
	OP_UI8_UI64
	OP_UI8_F32
	OP_UI8_F64

	OP_UI16_STR
	OP_UI16_I8
	OP_UI16_I16
	OP_UI16_I32
	OP_UI16_I64
	OP_UI16_UI8
	OP_UI16_UI32
	OP_UI16_UI64
	OP_UI16_F32
	OP_UI16_F64

	OP_UI32_STR
	OP_UI32_I8
	OP_UI32_I16
	OP_UI32_I32
	OP_UI32_I64
	OP_UI32_UI8
	OP_UI32_UI16
	OP_UI32_UI64
	OP_UI32_F32
	OP_UI32_F64

	OP_UI64_STR
	OP_UI64_I8
	OP_UI64_I16
	OP_UI64_I32
	OP_UI64_I64
	OP_UI64_UI8
	OP_UI64_UI16
	OP_UI64_UI32
	OP_UI64_F32
	OP_UI64_F64

	OP_F32_STR
	OP_F32_I8
	OP_F32_I16
	OP_F32_I32
	OP_F32_I64
	OP_F32_UI8
	OP_F32_UI16
	OP_F32_UI32
	OP_F32_UI64
	OP_F32_F64

	OP_F64_STR
	OP_F64_I8
	OP_F64_I16
	OP_F64_I32
	OP_F64_I64
	OP_F64_UI8
	OP_F64_UI16
	OP_F64_UI32
	OP_F64_UI64
	OP_F64_F32

	OP_STR_I8
	OP_STR_I16
	OP_STR_I32
	OP_STR_I64
	OP_STR_UI8
	OP_STR_UI16
	OP_STR_UI32
	OP_STR_UI64
	OP_STR_F32
	OP_STR_F64

	END_PARSE_OPS

	OP_BOOL_PRINT
	OP_BOOL_EQUAL
	OP_BOOL_UNEQUAL
	OP_BOOL_NOT
	OP_BOOL_OR
	OP_BOOL_AND

	OP_I8_PRINT
	OP_I8_ADD
	OP_I8_SUB
	OP_I8_NEG
	OP_I8_MUL
	OP_I8_DIV
	OP_I8_MOD
	OP_I8_ABS
	OP_I8_GT
	OP_I8_GTEQ
	OP_I8_LT
	OP_I8_LTEQ
	OP_I8_EQ
	OP_I8_UNEQ
	OP_I8_BITAND
	OP_I8_BITOR
	OP_I8_BITXOR
	OP_I8_BITCLEAR
	OP_I8_BITSHL
	OP_I8_BITSHR
	OP_I8_MAX
	OP_I8_MIN
	OP_I8_RAND

	OP_I16_PRINT
	OP_I16_ADD
	OP_I16_SUB
	OP_I16_NEG
	OP_I16_MUL
	OP_I16_DIV
	OP_I16_MOD
	OP_I16_ABS
	OP_I16_GT
	OP_I16_GTEQ
	OP_I16_LT
	OP_I16_LTEQ
	OP_I16_EQ
	OP_I16_UNEQ
	OP_I16_BITAND
	OP_I16_BITOR
	OP_I16_BITXOR
	OP_I16_BITCLEAR
	OP_I16_BITSHL
	OP_I16_BITSHR
	OP_I16_MAX
	OP_I16_MIN
	OP_I16_RAND

	OP_I32_PRINT
	OP_I32_ADD
	OP_I32_SUB
	OP_I32_NEG
	OP_I32_MUL
	OP_I32_DIV
	OP_I32_MOD
	OP_I32_ABS
	OP_I32_GT
	OP_I32_GTEQ
	OP_I32_LT
	OP_I32_LTEQ
	OP_I32_EQ
	OP_I32_UNEQ
	OP_I32_BITAND
	OP_I32_BITOR
	OP_I32_BITXOR
	OP_I32_BITCLEAR
	OP_I32_BITSHL
	OP_I32_BITSHR
	OP_I32_MAX
	OP_I32_MIN
	OP_I32_RAND

	OP_I64_PRINT
	OP_I64_ADD
	OP_I64_SUB
	OP_I64_NEG
	OP_I64_MUL
	OP_I64_DIV
	OP_I64_MOD
	OP_I64_ABS
	OP_I64_GT
	OP_I64_GTEQ
	OP_I64_LT
	OP_I64_LTEQ
	OP_I64_EQ
	OP_I64_UNEQ
	OP_I64_BITAND
	OP_I64_BITOR
	OP_I64_BITXOR
	OP_I64_BITCLEAR
	OP_I64_BITSHL
	OP_I64_BITSHR
	OP_I64_MAX
	OP_I64_MIN
	OP_I64_RAND

	OP_UI8_PRINT
	OP_UI8_ADD
	OP_UI8_SUB
	OP_UI8_MUL
	OP_UI8_DIV
	OP_UI8_MOD
	OP_UI8_GT
	OP_UI8_GTEQ
	OP_UI8_LT
	OP_UI8_LTEQ
	OP_UI8_EQ
	OP_UI8_UNEQ
	OP_UI8_BITAND
	OP_UI8_BITOR
	OP_UI8_BITXOR
	OP_UI8_BITCLEAR
	OP_UI8_BITSHL
	OP_UI8_BITSHR
	OP_UI8_MAX
	OP_UI8_MIN
	OP_UI8_RAND

	OP_UI16_PRINT
	OP_UI16_ADD
	OP_UI16_SUB
	OP_UI16_MUL
	OP_UI16_DIV
	OP_UI16_MOD
	OP_UI16_GT
	OP_UI16_GTEQ
	OP_UI16_LT
	OP_UI16_LTEQ
	OP_UI16_EQ
	OP_UI16_UNEQ
	OP_UI16_BITAND
	OP_UI16_BITOR
	OP_UI16_BITXOR
	OP_UI16_BITCLEAR
	OP_UI16_BITSHL
	OP_UI16_BITSHR
	OP_UI16_MAX
	OP_UI16_MIN
	OP_UI16_RAND

	OP_UI32_PRINT
	OP_UI32_ADD
	OP_UI32_SUB
	OP_UI32_MUL
	OP_UI32_DIV
	OP_UI32_MOD
	OP_UI32_GT
	OP_UI32_GTEQ
	OP_UI32_LT
	OP_UI32_LTEQ
	OP_UI32_EQ
	OP_UI32_UNEQ
	OP_UI32_BITAND
	OP_UI32_BITOR
	OP_UI32_BITXOR
	OP_UI32_BITCLEAR
	OP_UI32_BITSHL
	OP_UI32_BITSHR
	OP_UI32_MAX
	OP_UI32_MIN
	OP_UI32_RAND

	OP_UI64_PRINT
	OP_UI64_ADD
	OP_UI64_SUB
	OP_UI64_MUL
	OP_UI64_DIV
	OP_UI64_MOD
	OP_UI64_GT
	OP_UI64_GTEQ
	OP_UI64_LT
	OP_UI64_LTEQ
	OP_UI64_EQ
	OP_UI64_UNEQ
	OP_UI64_BITAND
	OP_UI64_BITOR
	OP_UI64_BITXOR
	OP_UI64_BITCLEAR
	OP_UI64_BITSHL
	OP_UI64_BITSHR
	OP_UI64_MAX
	OP_UI64_MIN
	OP_UI64_RAND

	OP_F32_IS_NAN
	OP_F32_PRINT
	OP_F32_ADD
	OP_F32_SUB
	OP_F32_NEG
	OP_F32_MUL
	OP_F32_DIV
	OP_F32_MOD
	OP_F32_ABS
	OP_F32_POW
	OP_F32_GT
	OP_F32_GTEQ
	OP_F32_LT
	OP_F32_LTEQ
	OP_F32_EQ
	OP_F32_UNEQ
	OP_F32_ACOS
	OP_F32_COS
	OP_F32_ASIN
	OP_F32_SIN
	OP_F32_SQRT
	OP_F32_LOG
	OP_F32_LOG2
	OP_F32_LOG10
	OP_F32_MAX
	OP_F32_MIN
	OP_F32_RAND

	OP_F64_IS_NAN
	OP_F64_PRINT
	OP_F64_ADD
	OP_F64_SUB
	OP_F64_NEG
	OP_F64_MUL
	OP_F64_DIV
	OP_F64_MOD
	OP_F64_ABS
	OP_F64_POW
	OP_F64_GT
	OP_F64_GTEQ
	OP_F64_LT
	OP_F64_LTEQ
	OP_F64_EQ
	OP_F64_UNEQ
	OP_F64_ACOS
	OP_F64_COS
	OP_F64_ASIN
	OP_F64_SIN
	OP_F64_SQRT
	OP_F64_LOG
	OP_F64_LOG2
	OP_F64_LOG10
	OP_F64_MAX
	OP_F64_MIN
	OP_F64_RAND

	OP_STR_PRINT
	OP_STR_CONCAT
	OP_STR_SUBSTR
	OP_STR_INDEX
	OP_STR_LAST_INDEX
	OP_STR_TRIM_SPACE
	OP_STR_EQ
	OP_STR_UNEQ

	OP_APPEND
	OP_RESIZE
	OP_INSERT
	OP_REMOVE
	OP_COPY

	OP_ASSERT
	OP_TEST
	OP_PANIC
	OP_PANIC_IF
	OP_PANIC_IF_NOT
	OP_STRERROR

	OP_AFF_PRINT
	OP_AFF_QUERY
	OP_AFF_ON
	OP_AFF_OF
	OP_AFF_INFORM
	OP_AFF_REQUEST

	OP_TCP_DIAL
	OP_TCP_LISTEN

	OP_TCP_ACCEPT

	OP_TCP_CLOSE

	END_OF_CORE_OPS
)
