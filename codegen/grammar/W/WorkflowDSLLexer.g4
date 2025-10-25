lexer grammar WorkflowDSLLexer;

// ----------------- Keywords -----------------
WORKFLOW    : 'workflow';
OUT         : 'out';
MODELS      : 'models';
IF          : 'if';
ELSE        : 'else';

//

STRING_TYPE : 'string';
INT_TYPE    : 'int';
FLOAT_TYPE  : 'float';
AUTO        : 'auto';
STRUCT      : 'struct';

// ----------------- Symbols -----------------
LBRACE      : '{';
RBRACE      : '}';
LBRACK      : '[';
RBRACK      : ']';
LPAREN      : '(';
RPAREN      : ')';
COLON       : ':';
ARROW       : '<-';
COMMA       : ',';
DOT         : '.';
SLASH       : '/';

// ----------------- Operators -----------------
PLUS        : '+';
MINUS       : '-';
MULT        : '*';
// DIV         : '/';
EQ          : '==';
NEQ         : '!=';
LT          : '<';
LTE         : '<=';
GT          : '>';
GTE         : '>=';
DOLLAR      : '$';  
AMP         : '&';
PIPE        : '|';
PERCENT     : '%';
AT          : '@';
HASH        : '#';
CARET       : '^';
TILDE       : '~';
AND         : '&&';
OR          : '||';
NOT         : '!';
// arithmeticOP
SUM         :'sum';

// ----------------- Literals -----------------
IN_ID       : 'In' [a-zA-Z0-9_]* ;
IN_ID_ARRAY       : '[]In' [a-zA-Z0-9_]* ;
ID      : [a-zA-Z_][a-zA-Z_0-9]* ;
INT     : [0-9]+ ;
FLOAT   : [0-9]+ '.' [0-9]+ ;

STRING  : '"' (~["\\] | '\\' .)* '"' ;

// ----------------- Misc -----------------
WS      : [ \t\r\n]+ -> skip ;
COMMENT : '//' ~[\r\n]* -> skip ;
