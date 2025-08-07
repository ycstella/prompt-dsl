parser grammar PromptDSLParser;

options { tokenVocab=PromptDSLLexer; }

// 解析规则
promptFile  : promptDef+ EOF ;

promptDef   : PROMPT ID LBRACE promptBlock+ RBRACE ;

promptBlock 
    : inputSection
    | outputSection
    | systemSection
    | userSection
    | afterSection
    | fixSection
    | moduleDef
    | goimportSection
    ;

goimportSection
    : GOIMPORT LBRACE goimportEntry* RBRACE
    ;

goimportEntry
    : alias=ID? path=STRING
    ;
    
inputSection
    : INPUT LBRACE fieldDef+ RBRACE
    | INPUT fieldDef    
    ;

outputSection
    : defaultAnnotation* OUTPUT (outputStruct | outputMarkdown)
    ;

outputStruct : LBRACE fieldDef+ RBRACE ;
outputMarkdown : COLON MARKDOWN ;

beforeSection
    : BEFORE LBRACE beforeContent* RBRACE
    ;

beforeContent
    : varDef
    | expr
    | ifStatement
    | textLine
    ;

varDef
    : ID EQUAL expr
    ;

systemSection
    : SYSTEM LBRACE ID+ RBRACE
    | SYSTEM LBRACE sysContent+ RBRACE
    ;

sysContent
    : ifStatement
    | paramPath
    | forStatement 
    | ARRAY_OUTPUTSPEC
    | OUTPUTSPEC
    | expr
    | textLine
    ;



userSection
    : USER LBRACE ID+ RBRACE
    | USER LBRACE userContent+ RBRACE
    ;

userContent
    : ifStatement
    | paramPath
    | forStatement 
    | ARRAY_OUTPUTSPEC
    | OUTPUTSPEC
    | expr
    | textLine
    ;

moduleDef : ID LBRACE moduleContent* RBRACE ;
moduleContent
    : ifStatement
    | paramPath
    | forStatement 
    | switchStatement
    | ARRAY_OUTPUTSPEC
    | OUTPUTSPEC
    | expr
    | textLine
    ;
thencontent : userContent ;
elsecontent : userContent ;
forcontent  : userContent ;

ifStatement 
    : IF LPAREN condition RPAREN LBRACE thencontent* RBRACE (ELSE LBRACE elsecontent* RBRACE)?
    ;

condition
    : lhs=expr op=(EQEQ | NOTEQ | LT | LTE | GT | GTE | DECL_ASSIGN) rhs=expr
    | single=expr
    ;

// 完整 forStatement 结构
forStatement
    : FOR init=assignExpr SEMI condition SEMI update=updateExpr LBRACE forcontent* RBRACE              # ForTraditional
    | FOR key=( ID | UNDERSCORE ) COMMA val=ID DECL_ASSIGN RANGE iterable=expr LBRACE forcontent* RBRACE                # ForRangeWithIndex
    | FOR val=ID DECL_ASSIGN RANGE iterable=expr LBRACE forcontent* RBRACE                             # ForRangeNoIndex
    ;

// 支持赋值表达式
assignExpr
    : lhs=paramPath EQUAL rhs=expr
    ;

// 支持更新表达式：i++ / i-- / i += 1 等
updateExpr
    : paramPath (INCREMENT | DECREMENT | PLUSEQ expr | MINUSEQ expr | MULTEQ expr | DIVEQ expr | MODEQ expr)
    ;


//switch case
switchStatement
    : SWITCH condition LBRACE switchCase* switchDefault? RBRACE
    ;

switchCase
    : CASE condition COLON userContent*
    ;

switchDefault
    : DEFAULT COLON userContent*
    ;
typeCase
    : CASE typeName COLON userContent*
    ;

typeDefault
    : DEFAULT COLON userContent*
    ;

typeName
    : STAR? ID
    ;

dslCallExpr
    : paramPath LPAREN (expr (COMMA expr)*)? RPAREN
    ;

expr
    : expr op=(PLUS | MINUS | STAR | SLASH | MOD) expr   // 二元运算
    | paramPath
    | DASH_STRING
    | STRING
    | NUMBER
    | BOOL
    | LPAREN expr RPAREN
    ;

fieldDef
    : ID COLON type ( EQUAL value )? annotation* SEMI?
    ;

textLine
    : DASH_STRING
    | STRING
    | LINE_COMMENT
    | paramPath
    // | RAW_TEXT_LINE
    ;

paramPath
    : (ID | INPUT | OUTPUT | AFTER | BEFORE) (DOT ID | LBRACK expr RBRACK)*
    ;


structDef
    : ID STRUCT LBRACE fieldDef+ RBRACE 
    ;

annotation
    : AT ID (LPAREN annotationArgs? RPAREN)?
    ;
annotationArgs
    : annotationValue (COMMA annotationValue)* 
    ;

annotationValue
    : STRING
    | arrayLiteral
    ;

arrayLiteral
    : LBRACE (STRING (COMMA STRING)*)? RBRACE 
    ;

defaultAnnotation
    : AT ID (LPAREN annotationArgs? RPAREN)?
    ;

// Fix section - uses CODE_BLOCK mode
fixSection
    : FIX  codeBlockContent RBRACE
    ;

// After section - uses CODE_BLOCK mode
afterSection
    : AFTER  codeBlockContent RBRACE
    ;

// Handle nested braces and strings in code blocks
codeBlockContent
    : (CODE_TEXT | CODE_STRING | LBRACE codeBlockContent RBRACE)*
    ;

type
    : STRUCT LBRACE fieldDef* RBRACE
    | FLOAT_TYPE
    | INT_TYPE
    | LBRACK RBRACK type  // 这里你之前用的是 '[]' type，建议改写对应你Lexer中的表示数组的token，如果没有，需要补充
    | STRING_TYPE   
    | ID 
    ;

value
    : DASH_STRING
    | STRING
    | NUMBER
    | BOOL
    ;

formatType
    : MD
    | JSON
    ;
