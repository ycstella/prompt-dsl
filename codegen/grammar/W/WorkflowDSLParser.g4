parser grammar WorkflowDSLParser;

options { tokenVocab=WorkflowDSLLexer; }

// ----------------- Parser Rules -----------------
workflow
    : WORKFLOW LBRACE workflowBody RBRACE
    ;

workflowBody
    : inputs out models
    ;

inputs
    : inputBlock+   
    ;

inputBlock
    :  (IN_ID|IN_ID_ARRAY) LBRACE fieldDef* RBRACE
    ;


fieldDef
    : ID COLON type?
    ;

type
    : INT_TYPE
    | FLOAT_TYPE
    | STRING_TYPE
    | STRUCT LBRACE fieldDef* RBRACE   // 嵌套对象
    | LBRACK RBRACK type       // 数组
    ;

outtype
    : INT_TYPE
    | FLOAT_TYPE
    | STRING_TYPE
    | STRUCT LBRACE outField* RBRACE   // 嵌套对象
    | LBRACK RBRACK outtype       // 数组
    |
    ;

out
    : OUT LBRACE outField* RBRACE
    ;


expr
    : functionCall
    | conditionalExpr
    | arithmeticExpr
    | atom
    | INT
    | parenExpr 
    ;


atom
    : functionCall (DOT ID)*
    | (ID | IN_ID) (DOT ID)*
    ;

functionCall
    : ID LPAREN (expr (COMMA expr)*)? RPAREN
    ;


conditionalExpr
    : LBRACK IF condition RBRACK LPAREN expr RPAREN LBRACK ELSE RBRACK LPAREN expr RPAREN
    ;

// conditionalExpr
//     : LBRACK IF condition RBRACK exprWithoutConditional 
//       (LBRACK ELSE RBRACK exprWithoutConditional)?
//     ;

// exprWithoutConditional
//     : functionCall
//     | arithmeticExpr
//     | atom
//     | INT
//     ;

condition
    : LPAREN expr (comparator expr)? RPAREN 
    | LPAREN condition (comparator condition)? RPAREN
    ;
 
comparator
    : EQ | NEQ | LT | LTE | GT | GTE | AND | OR 
    ;

arithmeticExpr
    : LBRACK DOLLAR SUM arithExprContent RBRACK   # DollarSumExpr
    ;


parenExpr
    : LPAREN expr RPAREN 
    ;

arithExprContent
    : parenExpr MULT parenExpr  
    | parenExpr PLUS parenExpr 
    ;

models
    : MODELS LBRACE modelEntry* RBRACE
    ;

modelEntry
    : ID COLON (INT | AUTO)
    ;

outField
    : ID COLON outtype (ARROW expr)?
    ;
