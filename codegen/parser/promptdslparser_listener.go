// Code generated from ./codegen/grammar/PromptDSLParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // PromptDSLParser
import "github.com/antlr4-go/antlr/v4"

// PromptDSLParserListener is a complete listener for a parse tree produced by PromptDSLParser.
type PromptDSLParserListener interface {
	antlr.ParseTreeListener

	// EnterPromptFile is called when entering the promptFile production.
	EnterPromptFile(c *PromptFileContext)

	// EnterRunexeDef is called when entering the runexeDef production.
	EnterRunexeDef(c *RunexeDefContext)

	// EnterPromptDef is called when entering the promptDef production.
	EnterPromptDef(c *PromptDefContext)

	// EnterPromptBlock is called when entering the promptBlock production.
	EnterPromptBlock(c *PromptBlockContext)

	// EnterGoimportSection is called when entering the goimportSection production.
	EnterGoimportSection(c *GoimportSectionContext)

	// EnterGoimportEntry is called when entering the goimportEntry production.
	EnterGoimportEntry(c *GoimportEntryContext)

	// EnterInputSection is called when entering the inputSection production.
	EnterInputSection(c *InputSectionContext)

	// EnterIterater is called when entering the iterater production.
	EnterIterater(c *IteraterContext)

	// EnterLoopRangeSection is called when entering the loopRangeSection production.
	EnterLoopRangeSection(c *LoopRangeSectionContext)

	// EnterOutputSection is called when entering the outputSection production.
	EnterOutputSection(c *OutputSectionContext)

	// EnterOutputStruct is called when entering the outputStruct production.
	EnterOutputStruct(c *OutputStructContext)

	// EnterOutputMarkdown is called when entering the outputMarkdown production.
	EnterOutputMarkdown(c *OutputMarkdownContext)

	// EnterBeforeSection is called when entering the beforeSection production.
	EnterBeforeSection(c *BeforeSectionContext)

	// EnterSystemSection is called when entering the systemSection production.
	EnterSystemSection(c *SystemSectionContext)

	// EnterSysContent is called when entering the sysContent production.
	EnterSysContent(c *SysContentContext)

	// EnterUserSection is called when entering the userSection production.
	EnterUserSection(c *UserSectionContext)

	// EnterUserContent is called when entering the userContent production.
	EnterUserContent(c *UserContentContext)

	// EnterModuleDef is called when entering the moduleDef production.
	EnterModuleDef(c *ModuleDefContext)

	// EnterModuleContent is called when entering the moduleContent production.
	EnterModuleContent(c *ModuleContentContext)

	// EnterThencontent is called when entering the thencontent production.
	EnterThencontent(c *ThencontentContext)

	// EnterElsecontent is called when entering the elsecontent production.
	EnterElsecontent(c *ElsecontentContext)

	// EnterForcontent is called when entering the forcontent production.
	EnterForcontent(c *ForcontentContext)

	// EnterIfStatement is called when entering the ifStatement production.
	EnterIfStatement(c *IfStatementContext)

	// EnterCondition is called when entering the condition production.
	EnterCondition(c *ConditionContext)

	// EnterForTraditional is called when entering the ForTraditional production.
	EnterForTraditional(c *ForTraditionalContext)

	// EnterForRangeWithIndex is called when entering the ForRangeWithIndex production.
	EnterForRangeWithIndex(c *ForRangeWithIndexContext)

	// EnterForRangeNoIndex is called when entering the ForRangeNoIndex production.
	EnterForRangeNoIndex(c *ForRangeNoIndexContext)

	// EnterAssignExpr is called when entering the assignExpr production.
	EnterAssignExpr(c *AssignExprContext)

	// EnterUpdateExpr is called when entering the updateExpr production.
	EnterUpdateExpr(c *UpdateExprContext)

	// EnterSwitchStatement is called when entering the switchStatement production.
	EnterSwitchStatement(c *SwitchStatementContext)

	// EnterSwitchCase is called when entering the switchCase production.
	EnterSwitchCase(c *SwitchCaseContext)

	// EnterSwitchDefault is called when entering the switchDefault production.
	EnterSwitchDefault(c *SwitchDefaultContext)

	// EnterTypeCase is called when entering the typeCase production.
	EnterTypeCase(c *TypeCaseContext)

	// EnterTypeDefault is called when entering the typeDefault production.
	EnterTypeDefault(c *TypeDefaultContext)

	// EnterTypeName is called when entering the typeName production.
	EnterTypeName(c *TypeNameContext)

	// EnterDslCallExpr is called when entering the dslCallExpr production.
	EnterDslCallExpr(c *DslCallExprContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// EnterFieldDef is called when entering the fieldDef production.
	EnterFieldDef(c *FieldDefContext)

	// EnterTextLine is called when entering the textLine production.
	EnterTextLine(c *TextLineContext)

	// EnterParamPath is called when entering the paramPath production.
	EnterParamPath(c *ParamPathContext)

	// EnterStructDef is called when entering the structDef production.
	EnterStructDef(c *StructDefContext)

	// EnterAnnotation is called when entering the annotation production.
	EnterAnnotation(c *AnnotationContext)

	// EnterAnnotationArgs is called when entering the annotationArgs production.
	EnterAnnotationArgs(c *AnnotationArgsContext)

	// EnterAnnotationValue is called when entering the annotationValue production.
	EnterAnnotationValue(c *AnnotationValueContext)

	// EnterArrayLiteral is called when entering the arrayLiteral production.
	EnterArrayLiteral(c *ArrayLiteralContext)

	// EnterDefaultAnnotation is called when entering the defaultAnnotation production.
	EnterDefaultAnnotation(c *DefaultAnnotationContext)

	// EnterFixSection is called when entering the fixSection production.
	EnterFixSection(c *FixSectionContext)

	// EnterAfterSection is called when entering the afterSection production.
	EnterAfterSection(c *AfterSectionContext)

	// EnterCodeBlockContent is called when entering the codeBlockContent production.
	EnterCodeBlockContent(c *CodeBlockContentContext)

	// EnterType is called when entering the type production.
	EnterType(c *TypeContext)

	// EnterValue is called when entering the value production.
	EnterValue(c *ValueContext)

	// EnterFormatType is called when entering the formatType production.
	EnterFormatType(c *FormatTypeContext)

	// ExitPromptFile is called when exiting the promptFile production.
	ExitPromptFile(c *PromptFileContext)

	// ExitRunexeDef is called when exiting the runexeDef production.
	ExitRunexeDef(c *RunexeDefContext)

	// ExitPromptDef is called when exiting the promptDef production.
	ExitPromptDef(c *PromptDefContext)

	// ExitPromptBlock is called when exiting the promptBlock production.
	ExitPromptBlock(c *PromptBlockContext)

	// ExitGoimportSection is called when exiting the goimportSection production.
	ExitGoimportSection(c *GoimportSectionContext)

	// ExitGoimportEntry is called when exiting the goimportEntry production.
	ExitGoimportEntry(c *GoimportEntryContext)

	// ExitInputSection is called when exiting the inputSection production.
	ExitInputSection(c *InputSectionContext)

	// ExitIterater is called when exiting the iterater production.
	ExitIterater(c *IteraterContext)

	// ExitLoopRangeSection is called when exiting the loopRangeSection production.
	ExitLoopRangeSection(c *LoopRangeSectionContext)

	// ExitOutputSection is called when exiting the outputSection production.
	ExitOutputSection(c *OutputSectionContext)

	// ExitOutputStruct is called when exiting the outputStruct production.
	ExitOutputStruct(c *OutputStructContext)

	// ExitOutputMarkdown is called when exiting the outputMarkdown production.
	ExitOutputMarkdown(c *OutputMarkdownContext)

	// ExitBeforeSection is called when exiting the beforeSection production.
	ExitBeforeSection(c *BeforeSectionContext)

	// ExitSystemSection is called when exiting the systemSection production.
	ExitSystemSection(c *SystemSectionContext)

	// ExitSysContent is called when exiting the sysContent production.
	ExitSysContent(c *SysContentContext)

	// ExitUserSection is called when exiting the userSection production.
	ExitUserSection(c *UserSectionContext)

	// ExitUserContent is called when exiting the userContent production.
	ExitUserContent(c *UserContentContext)

	// ExitModuleDef is called when exiting the moduleDef production.
	ExitModuleDef(c *ModuleDefContext)

	// ExitModuleContent is called when exiting the moduleContent production.
	ExitModuleContent(c *ModuleContentContext)

	// ExitThencontent is called when exiting the thencontent production.
	ExitThencontent(c *ThencontentContext)

	// ExitElsecontent is called when exiting the elsecontent production.
	ExitElsecontent(c *ElsecontentContext)

	// ExitForcontent is called when exiting the forcontent production.
	ExitForcontent(c *ForcontentContext)

	// ExitIfStatement is called when exiting the ifStatement production.
	ExitIfStatement(c *IfStatementContext)

	// ExitCondition is called when exiting the condition production.
	ExitCondition(c *ConditionContext)

	// ExitForTraditional is called when exiting the ForTraditional production.
	ExitForTraditional(c *ForTraditionalContext)

	// ExitForRangeWithIndex is called when exiting the ForRangeWithIndex production.
	ExitForRangeWithIndex(c *ForRangeWithIndexContext)

	// ExitForRangeNoIndex is called when exiting the ForRangeNoIndex production.
	ExitForRangeNoIndex(c *ForRangeNoIndexContext)

	// ExitAssignExpr is called when exiting the assignExpr production.
	ExitAssignExpr(c *AssignExprContext)

	// ExitUpdateExpr is called when exiting the updateExpr production.
	ExitUpdateExpr(c *UpdateExprContext)

	// ExitSwitchStatement is called when exiting the switchStatement production.
	ExitSwitchStatement(c *SwitchStatementContext)

	// ExitSwitchCase is called when exiting the switchCase production.
	ExitSwitchCase(c *SwitchCaseContext)

	// ExitSwitchDefault is called when exiting the switchDefault production.
	ExitSwitchDefault(c *SwitchDefaultContext)

	// ExitTypeCase is called when exiting the typeCase production.
	ExitTypeCase(c *TypeCaseContext)

	// ExitTypeDefault is called when exiting the typeDefault production.
	ExitTypeDefault(c *TypeDefaultContext)

	// ExitTypeName is called when exiting the typeName production.
	ExitTypeName(c *TypeNameContext)

	// ExitDslCallExpr is called when exiting the dslCallExpr production.
	ExitDslCallExpr(c *DslCallExprContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

	// ExitFieldDef is called when exiting the fieldDef production.
	ExitFieldDef(c *FieldDefContext)

	// ExitTextLine is called when exiting the textLine production.
	ExitTextLine(c *TextLineContext)

	// ExitParamPath is called when exiting the paramPath production.
	ExitParamPath(c *ParamPathContext)

	// ExitStructDef is called when exiting the structDef production.
	ExitStructDef(c *StructDefContext)

	// ExitAnnotation is called when exiting the annotation production.
	ExitAnnotation(c *AnnotationContext)

	// ExitAnnotationArgs is called when exiting the annotationArgs production.
	ExitAnnotationArgs(c *AnnotationArgsContext)

	// ExitAnnotationValue is called when exiting the annotationValue production.
	ExitAnnotationValue(c *AnnotationValueContext)

	// ExitArrayLiteral is called when exiting the arrayLiteral production.
	ExitArrayLiteral(c *ArrayLiteralContext)

	// ExitDefaultAnnotation is called when exiting the defaultAnnotation production.
	ExitDefaultAnnotation(c *DefaultAnnotationContext)

	// ExitFixSection is called when exiting the fixSection production.
	ExitFixSection(c *FixSectionContext)

	// ExitAfterSection is called when exiting the afterSection production.
	ExitAfterSection(c *AfterSectionContext)

	// ExitCodeBlockContent is called when exiting the codeBlockContent production.
	ExitCodeBlockContent(c *CodeBlockContentContext)

	// ExitType is called when exiting the type production.
	ExitType(c *TypeContext)

	// ExitValue is called when exiting the value production.
	ExitValue(c *ValueContext)

	// ExitFormatType is called when exiting the formatType production.
	ExitFormatType(c *FormatTypeContext)
}
