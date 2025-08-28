// Code generated from ./codegen/grammar/PromptDSLParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // PromptDSLParser
import "github.com/antlr4-go/antlr/v4"

// BasePromptDSLParserListener is a complete listener for a parse tree produced by PromptDSLParser.
type BasePromptDSLParserListener struct{}

var _ PromptDSLParserListener = &BasePromptDSLParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasePromptDSLParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasePromptDSLParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasePromptDSLParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasePromptDSLParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterPromptFile is called when production promptFile is entered.
func (s *BasePromptDSLParserListener) EnterPromptFile(ctx *PromptFileContext) {}

// ExitPromptFile is called when production promptFile is exited.
func (s *BasePromptDSLParserListener) ExitPromptFile(ctx *PromptFileContext) {}

// EnterRunexeDef is called when production runexeDef is entered.
func (s *BasePromptDSLParserListener) EnterRunexeDef(ctx *RunexeDefContext) {}

// ExitRunexeDef is called when production runexeDef is exited.
func (s *BasePromptDSLParserListener) ExitRunexeDef(ctx *RunexeDefContext) {}

// EnterPromptDef is called when production promptDef is entered.
func (s *BasePromptDSLParserListener) EnterPromptDef(ctx *PromptDefContext) {}

// ExitPromptDef is called when production promptDef is exited.
func (s *BasePromptDSLParserListener) ExitPromptDef(ctx *PromptDefContext) {}

// EnterPromptBlock is called when production promptBlock is entered.
func (s *BasePromptDSLParserListener) EnterPromptBlock(ctx *PromptBlockContext) {}

// ExitPromptBlock is called when production promptBlock is exited.
func (s *BasePromptDSLParserListener) ExitPromptBlock(ctx *PromptBlockContext) {}

// EnterGoimportSection is called when production goimportSection is entered.
func (s *BasePromptDSLParserListener) EnterGoimportSection(ctx *GoimportSectionContext) {}

// ExitGoimportSection is called when production goimportSection is exited.
func (s *BasePromptDSLParserListener) ExitGoimportSection(ctx *GoimportSectionContext) {}

// EnterGoimportEntry is called when production goimportEntry is entered.
func (s *BasePromptDSLParserListener) EnterGoimportEntry(ctx *GoimportEntryContext) {}

// ExitGoimportEntry is called when production goimportEntry is exited.
func (s *BasePromptDSLParserListener) ExitGoimportEntry(ctx *GoimportEntryContext) {}

// EnterInputSection is called when production inputSection is entered.
func (s *BasePromptDSLParserListener) EnterInputSection(ctx *InputSectionContext) {}

// ExitInputSection is called when production inputSection is exited.
func (s *BasePromptDSLParserListener) ExitInputSection(ctx *InputSectionContext) {}

// EnterOutputSection is called when production outputSection is entered.
func (s *BasePromptDSLParserListener) EnterOutputSection(ctx *OutputSectionContext) {}

// ExitOutputSection is called when production outputSection is exited.
func (s *BasePromptDSLParserListener) ExitOutputSection(ctx *OutputSectionContext) {}

// EnterOutputStruct is called when production outputStruct is entered.
func (s *BasePromptDSLParserListener) EnterOutputStruct(ctx *OutputStructContext) {}

// ExitOutputStruct is called when production outputStruct is exited.
func (s *BasePromptDSLParserListener) ExitOutputStruct(ctx *OutputStructContext) {}

// EnterOutputMarkdown is called when production outputMarkdown is entered.
func (s *BasePromptDSLParserListener) EnterOutputMarkdown(ctx *OutputMarkdownContext) {}

// ExitOutputMarkdown is called when production outputMarkdown is exited.
func (s *BasePromptDSLParserListener) ExitOutputMarkdown(ctx *OutputMarkdownContext) {}

// EnterBeforeSection is called when production beforeSection is entered.
func (s *BasePromptDSLParserListener) EnterBeforeSection(ctx *BeforeSectionContext) {}

// ExitBeforeSection is called when production beforeSection is exited.
func (s *BasePromptDSLParserListener) ExitBeforeSection(ctx *BeforeSectionContext) {}

// EnterSystemSection is called when production systemSection is entered.
func (s *BasePromptDSLParserListener) EnterSystemSection(ctx *SystemSectionContext) {}

// ExitSystemSection is called when production systemSection is exited.
func (s *BasePromptDSLParserListener) ExitSystemSection(ctx *SystemSectionContext) {}

// EnterSysContent is called when production sysContent is entered.
func (s *BasePromptDSLParserListener) EnterSysContent(ctx *SysContentContext) {}

// ExitSysContent is called when production sysContent is exited.
func (s *BasePromptDSLParserListener) ExitSysContent(ctx *SysContentContext) {}

// EnterUserSection is called when production userSection is entered.
func (s *BasePromptDSLParserListener) EnterUserSection(ctx *UserSectionContext) {}

// ExitUserSection is called when production userSection is exited.
func (s *BasePromptDSLParserListener) ExitUserSection(ctx *UserSectionContext) {}

// EnterUserContent is called when production userContent is entered.
func (s *BasePromptDSLParserListener) EnterUserContent(ctx *UserContentContext) {}

// ExitUserContent is called when production userContent is exited.
func (s *BasePromptDSLParserListener) ExitUserContent(ctx *UserContentContext) {}

// EnterModuleDef is called when production moduleDef is entered.
func (s *BasePromptDSLParserListener) EnterModuleDef(ctx *ModuleDefContext) {}

// ExitModuleDef is called when production moduleDef is exited.
func (s *BasePromptDSLParserListener) ExitModuleDef(ctx *ModuleDefContext) {}

// EnterModuleContent is called when production moduleContent is entered.
func (s *BasePromptDSLParserListener) EnterModuleContent(ctx *ModuleContentContext) {}

// ExitModuleContent is called when production moduleContent is exited.
func (s *BasePromptDSLParserListener) ExitModuleContent(ctx *ModuleContentContext) {}

// EnterThencontent is called when production thencontent is entered.
func (s *BasePromptDSLParserListener) EnterThencontent(ctx *ThencontentContext) {}

// ExitThencontent is called when production thencontent is exited.
func (s *BasePromptDSLParserListener) ExitThencontent(ctx *ThencontentContext) {}

// EnterElsecontent is called when production elsecontent is entered.
func (s *BasePromptDSLParserListener) EnterElsecontent(ctx *ElsecontentContext) {}

// ExitElsecontent is called when production elsecontent is exited.
func (s *BasePromptDSLParserListener) ExitElsecontent(ctx *ElsecontentContext) {}

// EnterForcontent is called when production forcontent is entered.
func (s *BasePromptDSLParserListener) EnterForcontent(ctx *ForcontentContext) {}

// ExitForcontent is called when production forcontent is exited.
func (s *BasePromptDSLParserListener) ExitForcontent(ctx *ForcontentContext) {}

// EnterIfStatement is called when production ifStatement is entered.
func (s *BasePromptDSLParserListener) EnterIfStatement(ctx *IfStatementContext) {}

// ExitIfStatement is called when production ifStatement is exited.
func (s *BasePromptDSLParserListener) ExitIfStatement(ctx *IfStatementContext) {}

// EnterCondition is called when production condition is entered.
func (s *BasePromptDSLParserListener) EnterCondition(ctx *ConditionContext) {}

// ExitCondition is called when production condition is exited.
func (s *BasePromptDSLParserListener) ExitCondition(ctx *ConditionContext) {}

// EnterForTraditional is called when production ForTraditional is entered.
func (s *BasePromptDSLParserListener) EnterForTraditional(ctx *ForTraditionalContext) {}

// ExitForTraditional is called when production ForTraditional is exited.
func (s *BasePromptDSLParserListener) ExitForTraditional(ctx *ForTraditionalContext) {}

// EnterForRangeWithIndex is called when production ForRangeWithIndex is entered.
func (s *BasePromptDSLParserListener) EnterForRangeWithIndex(ctx *ForRangeWithIndexContext) {}

// ExitForRangeWithIndex is called when production ForRangeWithIndex is exited.
func (s *BasePromptDSLParserListener) ExitForRangeWithIndex(ctx *ForRangeWithIndexContext) {}

// EnterForRangeNoIndex is called when production ForRangeNoIndex is entered.
func (s *BasePromptDSLParserListener) EnterForRangeNoIndex(ctx *ForRangeNoIndexContext) {}

// ExitForRangeNoIndex is called when production ForRangeNoIndex is exited.
func (s *BasePromptDSLParserListener) ExitForRangeNoIndex(ctx *ForRangeNoIndexContext) {}

// EnterAssignExpr is called when production assignExpr is entered.
func (s *BasePromptDSLParserListener) EnterAssignExpr(ctx *AssignExprContext) {}

// ExitAssignExpr is called when production assignExpr is exited.
func (s *BasePromptDSLParserListener) ExitAssignExpr(ctx *AssignExprContext) {}

// EnterUpdateExpr is called when production updateExpr is entered.
func (s *BasePromptDSLParserListener) EnterUpdateExpr(ctx *UpdateExprContext) {}

// ExitUpdateExpr is called when production updateExpr is exited.
func (s *BasePromptDSLParserListener) ExitUpdateExpr(ctx *UpdateExprContext) {}

// EnterSwitchStatement is called when production switchStatement is entered.
func (s *BasePromptDSLParserListener) EnterSwitchStatement(ctx *SwitchStatementContext) {}

// ExitSwitchStatement is called when production switchStatement is exited.
func (s *BasePromptDSLParserListener) ExitSwitchStatement(ctx *SwitchStatementContext) {}

// EnterSwitchCase is called when production switchCase is entered.
func (s *BasePromptDSLParserListener) EnterSwitchCase(ctx *SwitchCaseContext) {}

// ExitSwitchCase is called when production switchCase is exited.
func (s *BasePromptDSLParserListener) ExitSwitchCase(ctx *SwitchCaseContext) {}

// EnterSwitchDefault is called when production switchDefault is entered.
func (s *BasePromptDSLParserListener) EnterSwitchDefault(ctx *SwitchDefaultContext) {}

// ExitSwitchDefault is called when production switchDefault is exited.
func (s *BasePromptDSLParserListener) ExitSwitchDefault(ctx *SwitchDefaultContext) {}

// EnterTypeCase is called when production typeCase is entered.
func (s *BasePromptDSLParserListener) EnterTypeCase(ctx *TypeCaseContext) {}

// ExitTypeCase is called when production typeCase is exited.
func (s *BasePromptDSLParserListener) ExitTypeCase(ctx *TypeCaseContext) {}

// EnterTypeDefault is called when production typeDefault is entered.
func (s *BasePromptDSLParserListener) EnterTypeDefault(ctx *TypeDefaultContext) {}

// ExitTypeDefault is called when production typeDefault is exited.
func (s *BasePromptDSLParserListener) ExitTypeDefault(ctx *TypeDefaultContext) {}

// EnterTypeName is called when production typeName is entered.
func (s *BasePromptDSLParserListener) EnterTypeName(ctx *TypeNameContext) {}

// ExitTypeName is called when production typeName is exited.
func (s *BasePromptDSLParserListener) ExitTypeName(ctx *TypeNameContext) {}

// EnterDslCallExpr is called when production dslCallExpr is entered.
func (s *BasePromptDSLParserListener) EnterDslCallExpr(ctx *DslCallExprContext) {}

// ExitDslCallExpr is called when production dslCallExpr is exited.
func (s *BasePromptDSLParserListener) ExitDslCallExpr(ctx *DslCallExprContext) {}

// EnterExpr is called when production expr is entered.
func (s *BasePromptDSLParserListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BasePromptDSLParserListener) ExitExpr(ctx *ExprContext) {}

// EnterFieldDef is called when production fieldDef is entered.
func (s *BasePromptDSLParserListener) EnterFieldDef(ctx *FieldDefContext) {}

// ExitFieldDef is called when production fieldDef is exited.
func (s *BasePromptDSLParserListener) ExitFieldDef(ctx *FieldDefContext) {}

// EnterTextLine is called when production textLine is entered.
func (s *BasePromptDSLParserListener) EnterTextLine(ctx *TextLineContext) {}

// ExitTextLine is called when production textLine is exited.
func (s *BasePromptDSLParserListener) ExitTextLine(ctx *TextLineContext) {}

// EnterParamPath is called when production paramPath is entered.
func (s *BasePromptDSLParserListener) EnterParamPath(ctx *ParamPathContext) {}

// ExitParamPath is called when production paramPath is exited.
func (s *BasePromptDSLParserListener) ExitParamPath(ctx *ParamPathContext) {}

// EnterStructDef is called when production structDef is entered.
func (s *BasePromptDSLParserListener) EnterStructDef(ctx *StructDefContext) {}

// ExitStructDef is called when production structDef is exited.
func (s *BasePromptDSLParserListener) ExitStructDef(ctx *StructDefContext) {}

// EnterAnnotation is called when production annotation is entered.
func (s *BasePromptDSLParserListener) EnterAnnotation(ctx *AnnotationContext) {}

// ExitAnnotation is called when production annotation is exited.
func (s *BasePromptDSLParserListener) ExitAnnotation(ctx *AnnotationContext) {}

// EnterAnnotationArgs is called when production annotationArgs is entered.
func (s *BasePromptDSLParserListener) EnterAnnotationArgs(ctx *AnnotationArgsContext) {}

// ExitAnnotationArgs is called when production annotationArgs is exited.
func (s *BasePromptDSLParserListener) ExitAnnotationArgs(ctx *AnnotationArgsContext) {}

// EnterAnnotationValue is called when production annotationValue is entered.
func (s *BasePromptDSLParserListener) EnterAnnotationValue(ctx *AnnotationValueContext) {}

// ExitAnnotationValue is called when production annotationValue is exited.
func (s *BasePromptDSLParserListener) ExitAnnotationValue(ctx *AnnotationValueContext) {}

// EnterArrayLiteral is called when production arrayLiteral is entered.
func (s *BasePromptDSLParserListener) EnterArrayLiteral(ctx *ArrayLiteralContext) {}

// ExitArrayLiteral is called when production arrayLiteral is exited.
func (s *BasePromptDSLParserListener) ExitArrayLiteral(ctx *ArrayLiteralContext) {}

// EnterDefaultAnnotation is called when production defaultAnnotation is entered.
func (s *BasePromptDSLParserListener) EnterDefaultAnnotation(ctx *DefaultAnnotationContext) {}

// ExitDefaultAnnotation is called when production defaultAnnotation is exited.
func (s *BasePromptDSLParserListener) ExitDefaultAnnotation(ctx *DefaultAnnotationContext) {}

// EnterFixSection is called when production fixSection is entered.
func (s *BasePromptDSLParserListener) EnterFixSection(ctx *FixSectionContext) {}

// ExitFixSection is called when production fixSection is exited.
func (s *BasePromptDSLParserListener) ExitFixSection(ctx *FixSectionContext) {}

// EnterAfterSection is called when production afterSection is entered.
func (s *BasePromptDSLParserListener) EnterAfterSection(ctx *AfterSectionContext) {}

// ExitAfterSection is called when production afterSection is exited.
func (s *BasePromptDSLParserListener) ExitAfterSection(ctx *AfterSectionContext) {}

// EnterCodeBlockContent is called when production codeBlockContent is entered.
func (s *BasePromptDSLParserListener) EnterCodeBlockContent(ctx *CodeBlockContentContext) {}

// ExitCodeBlockContent is called when production codeBlockContent is exited.
func (s *BasePromptDSLParserListener) ExitCodeBlockContent(ctx *CodeBlockContentContext) {}

// EnterType is called when production type is entered.
func (s *BasePromptDSLParserListener) EnterType(ctx *TypeContext) {}

// ExitType is called when production type is exited.
func (s *BasePromptDSLParserListener) ExitType(ctx *TypeContext) {}

// EnterValue is called when production value is entered.
func (s *BasePromptDSLParserListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BasePromptDSLParserListener) ExitValue(ctx *ValueContext) {}

// EnterFormatType is called when production formatType is entered.
func (s *BasePromptDSLParserListener) EnterFormatType(ctx *FormatTypeContext) {}

// ExitFormatType is called when production formatType is exited.
func (s *BasePromptDSLParserListener) ExitFormatType(ctx *FormatTypeContext) {}
