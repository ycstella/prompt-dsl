@echo off
chcp 65001


REM 生成词法分析器
java -jar D:/App/antlr-4.13.2-complete.jar -Dlanguage=Go -package Wparser -o ./codegen/Wparser ./codegen/grammar/W/WorkflowDSLLexer.g4

REM 生成语法分析器
java -jar D:/App/antlr-4.13.2-complete.jar -Dlanguage=Go -package Wparser -o ./codegen/Wparser -lib  ./codegen/parser ./codegen/grammar/W/WorkflowDSLParser.g4

echo 解析器生成完成。
@REM pausec