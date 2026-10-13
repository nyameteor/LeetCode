use std::collections::HashMap;

// =============================================================================
// Lexer
// =============================================================================

#[derive(Debug, PartialEq, Eq, Clone)]
enum TokenKind {
    OpenParen,
    CloseParen,
    Identifier(String),
    Integer(i32),
    Unknown(char),
    Invalid(String),
    Eof,
}

#[derive(Debug, PartialEq, Eq, Clone)]
struct Token {
    kind: TokenKind,
    start: usize,
}

impl Token {
    fn new(kind: TokenKind, start: usize) -> Token {
        Token { kind, start }
    }
}

struct Lexer<'a> {
    input: &'a str,
    pos: usize,
}

impl<'a> Lexer<'a> {
    fn new(input: &'a str) -> Lexer<'a> {
        Lexer { input, pos: 0 }
    }

    fn tokenize(&mut self) -> Token {
        self.skip_whitespace();

        match self.peek() {
            Some(ch) => match ch {
                '(' => {
                    let start = self.pos;
                    self.bump();
                    Token::new(TokenKind::OpenParen, start)
                }
                ')' => {
                    let start = self.pos;
                    self.bump();
                    Token::new(TokenKind::CloseParen, start)
                }
                ch if ch.is_ascii_lowercase() => self.tokenize_identifier(),
                ch if ch.is_ascii_digit() || ch == '-' => self.tokenize_integer(),
                _ => {
                    let start = self.pos;
                    self.bump();
                    Token::new(TokenKind::Unknown(ch), start)
                }
            },
            None => Token::new(TokenKind::Eof, self.pos),
        }
    }

    fn skip_whitespace(&mut self) {
        while self.peek().is_some_and(|x| x.is_whitespace()) {
            self.bump();
        }
    }

    fn tokenize_identifier(&mut self) -> Token {
        let start = self.pos;
        while self
            .peek()
            .is_some_and(|x| x.is_ascii_lowercase() || x.is_ascii_digit())
        {
            self.bump();
        }

        let name = self.input[start..self.pos].to_owned();
        Token::new(TokenKind::Identifier(name), start)
    }

    fn tokenize_integer(&mut self) -> Token {
        let start = self.pos;

        if self.peek().is_some_and(|x| x == '-') {
            self.bump();
        }
        while self.peek().is_some_and(|x| x.is_ascii_digit()) {
            self.bump();
        }

        let numeric = self.input[start..self.pos].to_owned();
        match numeric.parse::<i32>() {
            Ok(integer) => Token::new(TokenKind::Integer(integer), start),
            Err(_) => Token::new(TokenKind::Invalid(numeric), start),
        }
    }

    fn peek(&self) -> Option<char> {
        self.input[self.pos..].chars().next()
    }

    fn bump(&mut self) {
        if let Some(ch) = self.peek() {
            self.pos += ch.len_utf8();
        }
    }
}

fn tokenize(input: &str) -> Vec<Token> {
    let mut tokens = Vec::new();
    let mut lexer = Lexer::new(input);
    loop {
        match lexer.tokenize() {
            token if token.kind == TokenKind::Eof => {
                tokens.push(token);
                break;
            }
            token => tokens.push(token),
        }
    }
    tokens
}

fn test_tokenizer() {
    let test_cases = vec![
        ("", vec![Token::new(TokenKind::Eof, 0)]),
        (
            "1",
            vec![
                Token::new(TokenKind::Integer(1), 0),
                Token::new(TokenKind::Eof, 1),
            ],
        ),
        (
            "(add -12 3)",
            vec![
                Token::new(TokenKind::OpenParen, 0),
                Token::new(TokenKind::Identifier("add".to_owned()), 1),
                Token::new(TokenKind::Integer(-12), 5),
                Token::new(TokenKind::Integer(3), 9),
                Token::new(TokenKind::CloseParen, 10),
                Token::new(TokenKind::Eof, 11),
            ],
        ),
        (
            "(add x1 2)",
            vec![
                Token::new(TokenKind::OpenParen, 0),
                Token::new(TokenKind::Identifier("add".to_owned()), 1),
                Token::new(TokenKind::Identifier("x1".to_owned()), 5),
                Token::new(TokenKind::Integer(2), 8),
                Token::new(TokenKind::CloseParen, 9),
                Token::new(TokenKind::Eof, 10),
            ],
        ),
        (
            "(let x 2 (mult x 3))",
            vec![
                Token::new(TokenKind::OpenParen, 0),
                Token::new(TokenKind::Identifier("let".to_owned()), 1),
                Token::new(TokenKind::Identifier("x".to_owned()), 5),
                Token::new(TokenKind::Integer(2), 7),
                Token::new(TokenKind::OpenParen, 9),
                Token::new(TokenKind::Identifier("mult".to_owned()), 10),
                Token::new(TokenKind::Identifier("x".to_owned()), 15),
                Token::new(TokenKind::Integer(3), 17),
                Token::new(TokenKind::CloseParen, 18),
                Token::new(TokenKind::CloseParen, 19),
                Token::new(TokenKind::Eof, 20),
            ],
        ),
        (
            "(let x 1 y 2 (add x y))",
            vec![
                Token::new(TokenKind::OpenParen, 0),
                Token::new(TokenKind::Identifier("let".to_owned()), 1),
                Token::new(TokenKind::Identifier("x".to_owned()), 5),
                Token::new(TokenKind::Integer(1), 7),
                Token::new(TokenKind::Identifier("y".to_owned()), 9),
                Token::new(TokenKind::Integer(2), 11),
                Token::new(TokenKind::OpenParen, 13),
                Token::new(TokenKind::Identifier("add".to_owned()), 14),
                Token::new(TokenKind::Identifier("x".to_owned()), 18),
                Token::new(TokenKind::Identifier("y".to_owned()), 20),
                Token::new(TokenKind::CloseParen, 21),
                Token::new(TokenKind::CloseParen, 22),
                Token::new(TokenKind::Eof, 23),
            ],
        ),
        (
            "(add @ 1)",
            vec![
                Token::new(TokenKind::OpenParen, 0),
                Token::new(TokenKind::Identifier("add".to_owned()), 1),
                Token::new(TokenKind::Unknown('@'), 5),
                Token::new(TokenKind::Integer(1), 7),
                Token::new(TokenKind::CloseParen, 8),
                Token::new(TokenKind::Eof, 9),
            ],
        ),
        (
            "-",
            vec![
                Token::new(TokenKind::Invalid("-".to_string()), 0),
                Token::new(TokenKind::Eof, 1),
            ],
        ),
        (
            "999999999999999999999",
            vec![
                Token::new(TokenKind::Invalid("999999999999999999999".to_string()), 0),
                Token::new(TokenKind::Eof, 21),
            ],
        ),
    ];

    for (input, expected) in test_cases {
        let actual = tokenize(input);
        assert_eq!(actual, expected, "input: {input:?}");
    }
}

// =============================================================================
// Parser
// Expr = Integer | Symbol | List(Expr*)
// =============================================================================

#[derive(Debug, PartialEq, Eq)]
enum ExprKind {
    Integer(i32),
    Symbol(String),
    List(Vec<Expr>),
}

#[derive(Debug, PartialEq, Eq)]
struct Expr {
    kind: ExprKind,
    start: usize,
}

impl Expr {
    fn new(kind: ExprKind, start: usize) -> Expr {
        Expr { kind, start }
    }
}

struct Parser<'a> {
    tokens: &'a [Token],
    pos: usize,
}

#[derive(Debug, PartialEq, Eq)]
enum ParseError {
    UnexpectedToken(Token),
}

impl<'a> Parser<'a> {
    fn new(tokens: &'a [Token]) -> Parser<'a> {
        Parser { tokens, pos: 0 }
    }

    fn parse(&mut self) -> Result<Expr, ParseError> {
        let token = self.peek();
        match &token.kind {
            TokenKind::Integer(value) => {
                let value = *value;
                let start = token.start;
                self.bump();
                Ok(Expr::new(ExprKind::Integer(value), start))
            }
            TokenKind::Identifier(name) => {
                let name = name.clone();
                let start = token.start;
                self.bump();
                Ok(Expr::new(ExprKind::Symbol(name), start))
            }
            TokenKind::OpenParen => self.parse_list(),
            _ => Err(Self::unexpected(token)),
        }
    }

    fn parse_list(&mut self) -> Result<Expr, ParseError> {
        let start = self.peek().start;
        self.expect(TokenKind::OpenParen)?;
        let mut items = Vec::new();
        while self.peek().kind != TokenKind::CloseParen {
            let expr = self.parse()?;
            items.push(expr);
        }
        self.expect(TokenKind::CloseParen)?;
        Ok(Expr::new(ExprKind::List(items), start))
    }

    fn peek(&self) -> &Token {
        &self.tokens[self.pos]
    }

    fn bump(&mut self) {
        self.pos += 1;
    }

    fn expect(&mut self, expected: TokenKind) -> Result<(), ParseError> {
        let token = self.peek();
        if !(token.kind == expected) {
            Err(ParseError::UnexpectedToken(token.clone()))
        } else {
            self.bump();
            Ok(())
        }
    }

    fn unexpected(token: &Token) -> ParseError {
        ParseError::UnexpectedToken(token.clone())
    }
}

fn parse(tokens: &[Token]) -> Result<Expr, ParseError> {
    let mut parser = Parser::new(tokens);
    let expr = parser.parse()?;
    let token = parser.peek();
    match &token.kind {
        TokenKind::Eof => Ok(expr),
        _ => Err(Parser::unexpected(token)),
    }
}

fn parse_input(input: &str) -> Result<Expr, ParseError> {
    let tokens = tokenize(input);
    parse(&tokens)
}

fn test_parser() {
    let ok_cases = vec![
        ("42", Expr::new(ExprKind::Integer(42), 0)),
        ("-12", Expr::new(ExprKind::Integer(-12), 0)),
        ("x", Expr::new(ExprKind::Symbol("x".to_owned()), 0)),
        ("y1", Expr::new(ExprKind::Symbol("y1".to_owned()), 0)),
        (
            "(add 1 2)",
            Expr::new(
                ExprKind::List(vec![
                    Expr::new(ExprKind::Symbol("add".to_owned()), 1),
                    Expr::new(ExprKind::Integer(1), 5),
                    Expr::new(ExprKind::Integer(2), 7),
                ]),
                0,
            ),
        ),
        (
            "(mult 3 4)",
            Expr::new(
                ExprKind::List(vec![
                    Expr::new(ExprKind::Symbol("mult".to_owned()), 1),
                    Expr::new(ExprKind::Integer(3), 6),
                    Expr::new(ExprKind::Integer(4), 8),
                ]),
                0,
            ),
        ),
        (
            "(add 1 (mult 2 3))",
            Expr::new(
                ExprKind::List(vec![
                    Expr::new(ExprKind::Symbol("add".to_owned()), 1),
                    Expr::new(ExprKind::Integer(1), 5),
                    Expr::new(
                        ExprKind::List(vec![
                            Expr::new(ExprKind::Symbol("mult".to_owned()), 8),
                            Expr::new(ExprKind::Integer(2), 13),
                            Expr::new(ExprKind::Integer(3), 15),
                        ]),
                        7,
                    ),
                ]),
                0,
            ),
        ),
        (
            "(mult (add 1 2) 3)",
            Expr::new(
                ExprKind::List(vec![
                    Expr::new(ExprKind::Symbol("mult".to_owned()), 1),
                    Expr::new(
                        ExprKind::List(vec![
                            Expr::new(ExprKind::Symbol("add".to_owned()), 7),
                            Expr::new(ExprKind::Integer(1), 11),
                            Expr::new(ExprKind::Integer(2), 13),
                        ]),
                        6,
                    ),
                    Expr::new(ExprKind::Integer(3), 16),
                ]),
                0,
            ),
        ),
        (
            "(let x 1 x)",
            Expr::new(
                ExprKind::List(vec![
                    Expr::new(ExprKind::Symbol("let".to_owned()), 1),
                    Expr::new(ExprKind::Symbol("x".to_owned()), 5),
                    Expr::new(ExprKind::Integer(1), 7),
                    Expr::new(ExprKind::Symbol("x".to_owned()), 9),
                ]),
                0,
            ),
        ),
        (
            "(let x 1 y 2 (add x y))",
            Expr::new(
                ExprKind::List(vec![
                    Expr::new(ExprKind::Symbol("let".to_owned()), 1),
                    Expr::new(ExprKind::Symbol("x".to_owned()), 5),
                    Expr::new(ExprKind::Integer(1), 7),
                    Expr::new(ExprKind::Symbol("y".to_owned()), 9),
                    Expr::new(ExprKind::Integer(2), 11),
                    Expr::new(
                        ExprKind::List(vec![
                            Expr::new(ExprKind::Symbol("add".to_owned()), 14),
                            Expr::new(ExprKind::Symbol("x".to_owned()), 18),
                            Expr::new(ExprKind::Symbol("y".to_owned()), 20),
                        ]),
                        13,
                    ),
                ]),
                0,
            ),
        ),
        (
            "(let x 1 y (add x 2) (add x y))",
            Expr::new(
                ExprKind::List(vec![
                    Expr::new(ExprKind::Symbol("let".to_owned()), 1),
                    Expr::new(ExprKind::Symbol("x".to_owned()), 5),
                    Expr::new(ExprKind::Integer(1), 7),
                    Expr::new(ExprKind::Symbol("y".to_owned()), 9),
                    Expr::new(
                        ExprKind::List(vec![
                            Expr::new(ExprKind::Symbol("add".to_owned()), 12),
                            Expr::new(ExprKind::Symbol("x".to_owned()), 16),
                            Expr::new(ExprKind::Integer(2), 18),
                        ]),
                        11,
                    ),
                    Expr::new(
                        ExprKind::List(vec![
                            Expr::new(ExprKind::Symbol("add".to_owned()), 22),
                            Expr::new(ExprKind::Symbol("x".to_owned()), 26),
                            Expr::new(ExprKind::Symbol("y".to_owned()), 28),
                        ]),
                        21,
                    ),
                ]),
                0,
            ),
        ),
    ];

    for (input, expected) in ok_cases {
        let actual = parse_input(input);
        assert_eq!(actual, Ok(expected), "input: {input:?}");
    }

    let error_cases = vec![
        (
            "",
            ParseError::UnexpectedToken(Token::new(TokenKind::Eof, 0)),
        ),
        (
            "   ",
            ParseError::UnexpectedToken(Token::new(TokenKind::Eof, 3)),
        ),
        (
            "add 1 2)",
            ParseError::UnexpectedToken(Token::new(TokenKind::Integer(1), 4)),
        ),
        (
            "(add 1 2",
            ParseError::UnexpectedToken(Token::new(TokenKind::Eof, 8)),
        ),
        (
            "(add 1 2))",
            ParseError::UnexpectedToken(Token::new(TokenKind::CloseParen, 9)),
        ),
        (
            "(add 1 2) (mult 3 4)",
            ParseError::UnexpectedToken(Token::new(TokenKind::OpenParen, 10)),
        ),
        (
            "(add (mult 2 3) 4",
            ParseError::UnexpectedToken(Token::new(TokenKind::Eof, 17)),
        ),
    ];

    for (input, expected) in error_cases {
        assert_eq!(parse_input(input), Err(expected), "input: {input:?}");
    }
}

// =============================================================================
// Evaluator
// =============================================================================

struct Frame {
    symbol_table: HashMap<String, i32>,
}

impl Frame {
    fn new() -> Frame {
        Frame {
            symbol_table: HashMap::new(),
        }
    }

    fn insert(&mut self, name: &str, value: i32) {
        self.symbol_table.insert(name.to_owned(), value);
    }

    fn get(&self, name: &str) -> Option<&i32> {
        self.symbol_table.get(name)
    }
}

struct Env {
    frames: Vec<Frame>,
}

impl Env {
    fn new() -> Env {
        Env { frames: Vec::new() }
    }

    fn push(&mut self) {
        self.frames.push(Frame::new());
    }

    fn pop(&mut self) {
        self.frames.pop();
    }

    fn bind(&mut self, name: &str, value: i32) {
        self.frames
            .last_mut()
            .expect("Cannot bind without an active frame")
            .insert(name, value);
    }

    fn lookup(&self, name: &str) -> Option<i32> {
        for frame in self.frames.iter().rev() {
            match frame.get(name) {
                Some(value) => return Some(*value),
                None => continue,
            }
        }
        None
    }
}

#[derive(Debug, PartialEq, Eq)]
enum EvalError {
    UnboundVariable(String),
    InvalidExpression(usize),
}

#[derive(Debug, PartialEq, Eq)]
enum InterpreterError {
    Parse(ParseError),
    Eval(EvalError),
}

impl From<ParseError> for InterpreterError {
    fn from(value: ParseError) -> Self {
        InterpreterError::Parse(value)
    }
}

impl From<EvalError> for InterpreterError {
    fn from(value: EvalError) -> Self {
        InterpreterError::Eval(value)
    }
}

struct Evaluator {
    expr: Expr,
    env: Env,
}

impl Evaluator {
    fn new(expr: Expr) -> Evaluator {
        Evaluator {
            expr,
            env: Env::new(),
        }
    }

    fn evaluate(&mut self) -> Result<i32, EvalError> {
        Self::evaluate_expr(&mut self.env, &self.expr)
    }

    fn evaluate_expr(env: &mut Env, expr: &Expr) -> Result<i32, EvalError> {
        match &expr.kind {
            ExprKind::Integer(value) => Ok(*value),
            ExprKind::Symbol(name) => match env.lookup(name) {
                Some(value) => Ok(value),
                None => Err(EvalError::UnboundVariable(name.to_string())),
            },
            ExprKind::List(exprs) => match exprs.first().map(|x| &x.kind) {
                Some(ExprKind::Symbol(symbol)) => match symbol.as_str() {
                    "add" | "mult" => {
                        if exprs.len() != 3 {
                            return Err(EvalError::InvalidExpression(expr.start));
                        }
                        let left = Self::evaluate_expr(env, &exprs[1])?;
                        let right = Self::evaluate_expr(env, &exprs[2])?;
                        match symbol.as_str() {
                            "add" => Ok(left + right),
                            "mult" => Ok(left * right),
                            _ => unreachable!(),
                        }
                    }
                    "let" => {
                        if !(exprs.len() >= 4 && exprs.len() % 2 == 0) {
                            return Err(EvalError::InvalidExpression(expr.start));
                        }
                        env.push();
                        let result = (|| {
                            let (body, bindings) = exprs[1..].split_last().unwrap();
                            for pair in bindings.chunks_exact(2) {
                                let [name_expr, value_expr] = pair else {
                                    return Err(EvalError::InvalidExpression(expr.start));
                                };
                                let ExprKind::Symbol(name) = &name_expr.kind else {
                                    return Err(EvalError::InvalidExpression(name_expr.start));
                                };
                                let value = Self::evaluate_expr(env, value_expr)?;
                                env.bind(name.as_str(), value);
                            }
                            Self::evaluate_expr(env, body)
                        })();
                        env.pop();
                        result
                    }
                    _ => Err(EvalError::InvalidExpression(expr.start)),
                },
                _ => Err(EvalError::InvalidExpression(expr.start)),
            },
        }
    }
}

fn evaluate(expr: Expr) -> Result<i32, EvalError> {
    let mut evaluator = Evaluator::new(expr);
    evaluator.evaluate()
}

fn evaluate_input(input: &str) -> Result<i32, InterpreterError> {
    let tokens = tokenize(input);
    let expr = parse(&tokens)?;
    Ok(evaluate(expr)?)
}

fn test_evaluate() {
    let ok_cases = vec![
        ("1", 1),
        ("-42", -42),
        ("0", 0),
        ("(add 1 (mult 2 3))", 7),
        ("(mult (add 2 3) 4)", 20),
        ("(add (mult 2 (add 3 4)) 5)", 19),
        ("(let x 5 x)", 5),
        ("(let x 2 y 3 (add x y))", 5),
        ("(let x 2 y (mult x 4) (add x y))", 10),
        ("(let x 1 (let x 2 x))", 2),
        ("(let x 1 (let x (add x 1) x))", 2),
        ("(let x 1 y (let x 2 (add x 3)) (add x y))", 6),
    ];

    for (input, expected) in ok_cases {
        let actual = evaluate_input(input).unwrap();
        assert_eq!(actual, expected, "input: {input:?}");
    }

    let error_cases = vec![
        (
            "x",
            InterpreterError::Eval(EvalError::UnboundVariable("x".to_owned())),
        ),
        (
            "(unknown 1 2)",
            InterpreterError::Eval(EvalError::InvalidExpression(0)),
        ),
        (
            "(add 1)",
            InterpreterError::Eval(EvalError::InvalidExpression(0)),
        ),
        (
            "(mult 2 3 4)",
            InterpreterError::Eval(EvalError::InvalidExpression(0)),
        ),
        (
            "()",
            InterpreterError::Eval(EvalError::InvalidExpression(0)),
        ),
        (
            "(let x 1 y)",
            InterpreterError::Eval(EvalError::UnboundVariable("y".to_owned())),
        ),
        (
            "(let 1 2 3)",
            InterpreterError::Eval(EvalError::InvalidExpression(5)),
        ),
    ];

    for (input, expected) in error_cases {
        assert_eq!(evaluate_input(input), Err(expected), "input: {input:?}");
    }
}

struct Solution;

impl Solution {
    pub fn evaluate(expression: String) -> i32 {
        evaluate_input(&expression).unwrap()
    }
}

fn test_solution() {
    let test_cases = vec![
        ("(let x 2 (mult x (let x 3 y 4 (add x y))))", 14),
        ("(let x 3 x 2 x)", 2),
        ("(let x 1 y 2 x (add x y) (add x y))", 5),
    ];

    for (input, expected) in test_cases {
        let actual = Solution::evaluate(input.to_owned());
        assert_eq!(actual, expected, "input: {input:?}");
    }
}

fn main() {
    test_tokenizer();
    test_parser();
    test_evaluate();
    test_solution();
}
