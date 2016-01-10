package syn

type (
	Statement interface {
		statement()
	}

	Expression interface {
		expression()
	}
)

type (
	ExpressionStatement struct {
		Expr Expression
	}

	VarDefStatement struct {
		VarName string
		Expr Expression
	}

	EchoStatement struct {
		Expr    Expression
	}

	ClassStatement struct {
		VarName string
		Statements []Statement
	}

	FunctionStatement struct {
		VarName string
		Statements []Statement
	}
)

func (x *ExpressionStatement) statement() {}
func (x *VarDefStatement) statement() {}
func (x *EchoStatement) statement() {}
func (x *ClassStatement) statement() {}
func (x *FunctionStatement) statement() {}

type (
	NumberExpression struct {
		Lit string
	}

	IdentifierExpression struct {
		Lit string
	}

	UnaryMinusExpression struct {
		SubExpr Expression
	}

	ParenExpression struct {
		SubExpr Expression
	}

	BinOpExpression struct {
		LHS      Expression
		Operator int
		RHS      Expression
	}
)

func (x *NumberExpression) expression()     {}
func (x *IdentifierExpression) expression() {}
func (x *UnaryMinusExpression) expression() {}
func (x *ParenExpression) expression()      {}
func (x *BinOpExpression) expression()      {}
