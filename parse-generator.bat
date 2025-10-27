@echo off
chcp 65001


REM 生成词法分析器
java -jar D:/App/antlr-4.13.2-complete.jar -Dlanguage=Go -package parser -o ./codegen/parser ./codegen/grammar/P/PromptDSLLexer.g4

REM 生成语法分析器
java -jar D:/App/antlr-4.13.2-complete.jar -Dlanguage=Go -package parser -o ./codegen/parser -lib  ./codegen/parser ./codegen/grammar/P/PromptDSLParser.g4

echo 解析器生成完成。
@REM pausec