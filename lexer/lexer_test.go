package lexer

import (
	"testing"

	"github.com/sriramr98/modern_js/token"
)

type TestCases struct {
	expectedType    token.TokenType
	expectedLiteral string
}

func testWithCases(t *testing.T, input string, tests []TestCases) {
	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests [%d] - token type wrong. Expected %q, got %q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests [%d] - token type wrong. Expected %q, got %q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}

func TestNextTokenWithOnlySpecialChars(t *testing.T) {
	input := `=+();<>`

	tests := []TestCases{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.LT, "<"},
		{token.GT, ">"},
		{token.EOF, ""},
	}

	testWithCases(t, input, tests)
}

func TestNextTokenForValidCode(t *testing.T) {
	input := `let five = 5;
    let ten = 10;

    function add(x, y) {
        return x + y;
    }

    let result = add(five, ten);
    `

	tests := []TestCases{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.FUNCTION, "function"},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	testWithCases(t, input, tests)
}

func TestNextTokenWithConditionals(t *testing.T) {
	input := `
        if (result == 10) {
            return true;
        }

        if (result >= 20) {
            return true;
        }

        if (result <= 50) {
            return false;
        }

        if (result != 10) {
            return true;
        }
    }
    `

	tests := []TestCases{
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.IDENT, "result"},
		{token.EQ, "=="},
		{token.INT, "10"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.IDENT, "result"},
		{token.GT_OR_EQ, ">="},
		{token.INT, "20"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.IDENT, "result"},
		{token.LT_OR_EQ, "<="},
		{token.INT, "50"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.FALSE, "false"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.IDENT, "result"},
		{token.NOT_EQ, "!="},
		{token.INT, "10"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
	}

	testWithCases(t, input, tests)
}
