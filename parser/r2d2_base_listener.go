// Code generated from R2D2.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // R2D2

import "github.com/antlr4-go/antlr/v4"

// BaseR2D2Listener is a complete listener for a parse tree produced by R2D2Parser.
type BaseR2D2Listener struct{}

var _ R2D2Listener = &BaseR2D2Listener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseR2D2Listener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseR2D2Listener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseR2D2Listener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseR2D2Listener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseR2D2Listener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseR2D2Listener) ExitProgram(ctx *ProgramContext) {}

// EnterDeclaration is called when production declaration is entered.
func (s *BaseR2D2Listener) EnterDeclaration(ctx *DeclarationContext) {}

// ExitDeclaration is called when production declaration is exited.
func (s *BaseR2D2Listener) ExitDeclaration(ctx *DeclarationContext) {}

// EnterGlobalDeclaration is called when production globalDeclaration is entered.
func (s *BaseR2D2Listener) EnterGlobalDeclaration(ctx *GlobalDeclarationContext) {}

// ExitGlobalDeclaration is called when production globalDeclaration is exited.
func (s *BaseR2D2Listener) ExitGlobalDeclaration(ctx *GlobalDeclarationContext) {}

// EnterImportDeclaration is called when production importDeclaration is entered.
func (s *BaseR2D2Listener) EnterImportDeclaration(ctx *ImportDeclarationContext) {}

// ExitImportDeclaration is called when production importDeclaration is exited.
func (s *BaseR2D2Listener) ExitImportDeclaration(ctx *ImportDeclarationContext) {}

// EnterInterfaceDeclaration is called when production interfaceDeclaration is entered.
func (s *BaseR2D2Listener) EnterInterfaceDeclaration(ctx *InterfaceDeclarationContext) {}

// ExitInterfaceDeclaration is called when production interfaceDeclaration is exited.
func (s *BaseR2D2Listener) ExitInterfaceDeclaration(ctx *InterfaceDeclarationContext) {}

// EnterModuleDeclaration is called when production moduleDeclaration is entered.
func (s *BaseR2D2Listener) EnterModuleDeclaration(ctx *ModuleDeclarationContext) {}

// ExitModuleDeclaration is called when production moduleDeclaration is exited.
func (s *BaseR2D2Listener) ExitModuleDeclaration(ctx *ModuleDeclarationContext) {}

// EnterFunctionDeclaration is called when production functionDeclaration is entered.
func (s *BaseR2D2Listener) EnterFunctionDeclaration(ctx *FunctionDeclarationContext) {}

// ExitFunctionDeclaration is called when production functionDeclaration is exited.
func (s *BaseR2D2Listener) ExitFunctionDeclaration(ctx *FunctionDeclarationContext) {}

// EnterFunctionCallStatement is called when production functionCallStatement is entered.
func (s *BaseR2D2Listener) EnterFunctionCallStatement(ctx *FunctionCallStatementContext) {}

// ExitFunctionCallStatement is called when production functionCallStatement is exited.
func (s *BaseR2D2Listener) ExitFunctionCallStatement(ctx *FunctionCallStatementContext) {}

// EnterFunctionCall is called when production functionCall is entered.
func (s *BaseR2D2Listener) EnterFunctionCall(ctx *FunctionCallContext) {}

// ExitFunctionCall is called when production functionCall is exited.
func (s *BaseR2D2Listener) ExitFunctionCall(ctx *FunctionCallContext) {}

// EnterParameterList is called when production parameterList is entered.
func (s *BaseR2D2Listener) EnterParameterList(ctx *ParameterListContext) {}

// ExitParameterList is called when production parameterList is exited.
func (s *BaseR2D2Listener) ExitParameterList(ctx *ParameterListContext) {}

// EnterParameter is called when production parameter is entered.
func (s *BaseR2D2Listener) EnterParameter(ctx *ParameterContext) {}

// ExitParameter is called when production parameter is exited.
func (s *BaseR2D2Listener) ExitParameter(ctx *ParameterContext) {}

// EnterTypeExpression is called when production typeExpression is entered.
func (s *BaseR2D2Listener) EnterTypeExpression(ctx *TypeExpressionContext) {}

// ExitTypeExpression is called when production typeExpression is exited.
func (s *BaseR2D2Listener) ExitTypeExpression(ctx *TypeExpressionContext) {}

// EnterArrayDimensions is called when production arrayDimensions is entered.
func (s *BaseR2D2Listener) EnterArrayDimensions(ctx *ArrayDimensionsContext) {}

// ExitArrayDimensions is called when production arrayDimensions is exited.
func (s *BaseR2D2Listener) ExitArrayDimensions(ctx *ArrayDimensionsContext) {}

// EnterBaseType is called when production baseType is entered.
func (s *BaseR2D2Listener) EnterBaseType(ctx *BaseTypeContext) {}

// ExitBaseType is called when production baseType is exited.
func (s *BaseR2D2Listener) ExitBaseType(ctx *BaseTypeContext) {}

// EnterGenericType is called when production genericType is entered.
func (s *BaseR2D2Listener) EnterGenericType(ctx *GenericTypeContext) {}

// ExitGenericType is called when production genericType is exited.
func (s *BaseR2D2Listener) ExitGenericType(ctx *GenericTypeContext) {}

// EnterTypeDeclaration is called when production typeDeclaration is entered.
func (s *BaseR2D2Listener) EnterTypeDeclaration(ctx *TypeDeclarationContext) {}

// ExitTypeDeclaration is called when production typeDeclaration is exited.
func (s *BaseR2D2Listener) ExitTypeDeclaration(ctx *TypeDeclarationContext) {}

// EnterVariableDeclaration is called when production variableDeclaration is entered.
func (s *BaseR2D2Listener) EnterVariableDeclaration(ctx *VariableDeclarationContext) {}

// ExitVariableDeclaration is called when production variableDeclaration is exited.
func (s *BaseR2D2Listener) ExitVariableDeclaration(ctx *VariableDeclarationContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseR2D2Listener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseR2D2Listener) ExitStatement(ctx *StatementContext) {}

// EnterExpressionStatement is called when production expressionStatement is entered.
func (s *BaseR2D2Listener) EnterExpressionStatement(ctx *ExpressionStatementContext) {}

// ExitExpressionStatement is called when production expressionStatement is exited.
func (s *BaseR2D2Listener) ExitExpressionStatement(ctx *ExpressionStatementContext) {}

// EnterIfStatement is called when production ifStatement is entered.
func (s *BaseR2D2Listener) EnterIfStatement(ctx *IfStatementContext) {}

// ExitIfStatement is called when production ifStatement is exited.
func (s *BaseR2D2Listener) ExitIfStatement(ctx *IfStatementContext) {}

// EnterForStatement is called when production forStatement is entered.
func (s *BaseR2D2Listener) EnterForStatement(ctx *ForStatementContext) {}

// ExitForStatement is called when production forStatement is exited.
func (s *BaseR2D2Listener) ExitForStatement(ctx *ForStatementContext) {}

// EnterAssignmentDeclaration is called when production assignmentDeclaration is entered.
func (s *BaseR2D2Listener) EnterAssignmentDeclaration(ctx *AssignmentDeclarationContext) {}

// ExitAssignmentDeclaration is called when production assignmentDeclaration is exited.
func (s *BaseR2D2Listener) ExitAssignmentDeclaration(ctx *AssignmentDeclarationContext) {}

// EnterAssignment is called when production assignment is entered.
func (s *BaseR2D2Listener) EnterAssignment(ctx *AssignmentContext) {}

// ExitAssignment is called when production assignment is exited.
func (s *BaseR2D2Listener) ExitAssignment(ctx *AssignmentContext) {}

// EnterAssignmentOperator is called when production assignmentOperator is entered.
func (s *BaseR2D2Listener) EnterAssignmentOperator(ctx *AssignmentOperatorContext) {}

// ExitAssignmentOperator is called when production assignmentOperator is exited.
func (s *BaseR2D2Listener) ExitAssignmentOperator(ctx *AssignmentOperatorContext) {}

// EnterSimpleFor is called when production simpleFor is entered.
func (s *BaseR2D2Listener) EnterSimpleFor(ctx *SimpleForContext) {}

// ExitSimpleFor is called when production simpleFor is exited.
func (s *BaseR2D2Listener) ExitSimpleFor(ctx *SimpleForContext) {}

// EnterWhileStatement is called when production whileStatement is entered.
func (s *BaseR2D2Listener) EnterWhileStatement(ctx *WhileStatementContext) {}

// ExitWhileStatement is called when production whileStatement is exited.
func (s *BaseR2D2Listener) ExitWhileStatement(ctx *WhileStatementContext) {}

// EnterLoopStatement is called when production loopStatement is entered.
func (s *BaseR2D2Listener) EnterLoopStatement(ctx *LoopStatementContext) {}

// ExitLoopStatement is called when production loopStatement is exited.
func (s *BaseR2D2Listener) ExitLoopStatement(ctx *LoopStatementContext) {}

// EnterCicleControl is called when production cicleControl is entered.
func (s *BaseR2D2Listener) EnterCicleControl(ctx *CicleControlContext) {}

// ExitCicleControl is called when production cicleControl is exited.
func (s *BaseR2D2Listener) ExitCicleControl(ctx *CicleControlContext) {}

// EnterBreakStatement is called when production breakStatement is entered.
func (s *BaseR2D2Listener) EnterBreakStatement(ctx *BreakStatementContext) {}

// ExitBreakStatement is called when production breakStatement is exited.
func (s *BaseR2D2Listener) ExitBreakStatement(ctx *BreakStatementContext) {}

// EnterContinueStatement is called when production continueStatement is entered.
func (s *BaseR2D2Listener) EnterContinueStatement(ctx *ContinueStatementContext) {}

// ExitContinueStatement is called when production continueStatement is exited.
func (s *BaseR2D2Listener) ExitContinueStatement(ctx *ContinueStatementContext) {}

// EnterReturnStatement is called when production returnStatement is entered.
func (s *BaseR2D2Listener) EnterReturnStatement(ctx *ReturnStatementContext) {}

// ExitReturnStatement is called when production returnStatement is exited.
func (s *BaseR2D2Listener) ExitReturnStatement(ctx *ReturnStatementContext) {}

// EnterArrayAccessExpression is called when production arrayAccessExpression is entered.
func (s *BaseR2D2Listener) EnterArrayAccessExpression(ctx *ArrayAccessExpressionContext) {}

// ExitArrayAccessExpression is called when production arrayAccessExpression is exited.
func (s *BaseR2D2Listener) ExitArrayAccessExpression(ctx *ArrayAccessExpressionContext) {}

// EnterAdditiveExpression is called when production additiveExpression is entered.
func (s *BaseR2D2Listener) EnterAdditiveExpression(ctx *AdditiveExpressionContext) {}

// ExitAdditiveExpression is called when production additiveExpression is exited.
func (s *BaseR2D2Listener) ExitAdditiveExpression(ctx *AdditiveExpressionContext) {}

// EnterIdentifierExpression is called when production identifierExpression is entered.
func (s *BaseR2D2Listener) EnterIdentifierExpression(ctx *IdentifierExpressionContext) {}

// ExitIdentifierExpression is called when production identifierExpression is exited.
func (s *BaseR2D2Listener) ExitIdentifierExpression(ctx *IdentifierExpressionContext) {}

// EnterFunctionCallExpression is called when production functionCallExpression is entered.
func (s *BaseR2D2Listener) EnterFunctionCallExpression(ctx *FunctionCallExpressionContext) {}

// ExitFunctionCallExpression is called when production functionCallExpression is exited.
func (s *BaseR2D2Listener) ExitFunctionCallExpression(ctx *FunctionCallExpressionContext) {}

// EnterParenthesisExpression is called when production parenthesisExpression is entered.
func (s *BaseR2D2Listener) EnterParenthesisExpression(ctx *ParenthesisExpressionContext) {}

// ExitParenthesisExpression is called when production parenthesisExpression is exited.
func (s *BaseR2D2Listener) ExitParenthesisExpression(ctx *ParenthesisExpressionContext) {}

// EnterComparisonExpression is called when production comparisonExpression is entered.
func (s *BaseR2D2Listener) EnterComparisonExpression(ctx *ComparisonExpressionContext) {}

// ExitComparisonExpression is called when production comparisonExpression is exited.
func (s *BaseR2D2Listener) ExitComparisonExpression(ctx *ComparisonExpressionContext) {}

// EnterMultiplicativeExpression is called when production multiplicativeExpression is entered.
func (s *BaseR2D2Listener) EnterMultiplicativeExpression(ctx *MultiplicativeExpressionContext) {}

// ExitMultiplicativeExpression is called when production multiplicativeExpression is exited.
func (s *BaseR2D2Listener) ExitMultiplicativeExpression(ctx *MultiplicativeExpressionContext) {}

// EnterLiteralExpression is called when production literalExpression is entered.
func (s *BaseR2D2Listener) EnterLiteralExpression(ctx *LiteralExpressionContext) {}

// ExitLiteralExpression is called when production literalExpression is exited.
func (s *BaseR2D2Listener) ExitLiteralExpression(ctx *LiteralExpressionContext) {}

// EnterUnaryExpression is called when production unaryExpression is entered.
func (s *BaseR2D2Listener) EnterUnaryExpression(ctx *UnaryExpressionContext) {}

// ExitUnaryExpression is called when production unaryExpression is exited.
func (s *BaseR2D2Listener) ExitUnaryExpression(ctx *UnaryExpressionContext) {}

// EnterLogicalExpression is called when production logicalExpression is entered.
func (s *BaseR2D2Listener) EnterLogicalExpression(ctx *LogicalExpressionContext) {}

// ExitLogicalExpression is called when production logicalExpression is exited.
func (s *BaseR2D2Listener) ExitLogicalExpression(ctx *LogicalExpressionContext) {}

// EnterMemberExpression is called when production memberExpression is entered.
func (s *BaseR2D2Listener) EnterMemberExpression(ctx *MemberExpressionContext) {}

// ExitMemberExpression is called when production memberExpression is exited.
func (s *BaseR2D2Listener) ExitMemberExpression(ctx *MemberExpressionContext) {}

// EnterMemberPart is called when production memberPart is entered.
func (s *BaseR2D2Listener) EnterMemberPart(ctx *MemberPartContext) {}

// ExitMemberPart is called when production memberPart is exited.
func (s *BaseR2D2Listener) ExitMemberPart(ctx *MemberPartContext) {}

// EnterArgumentList is called when production argumentList is entered.
func (s *BaseR2D2Listener) EnterArgumentList(ctx *ArgumentListContext) {}

// ExitArgumentList is called when production argumentList is exited.
func (s *BaseR2D2Listener) ExitArgumentList(ctx *ArgumentListContext) {}

// EnterPrimaryExpression is called when production primaryExpression is entered.
func (s *BaseR2D2Listener) EnterPrimaryExpression(ctx *PrimaryExpressionContext) {}

// ExitPrimaryExpression is called when production primaryExpression is exited.
func (s *BaseR2D2Listener) ExitPrimaryExpression(ctx *PrimaryExpressionContext) {}

// EnterArrayLiteral is called when production arrayLiteral is entered.
func (s *BaseR2D2Listener) EnterArrayLiteral(ctx *ArrayLiteralContext) {}

// ExitArrayLiteral is called when production arrayLiteral is exited.
func (s *BaseR2D2Listener) ExitArrayLiteral(ctx *ArrayLiteralContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BaseR2D2Listener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BaseR2D2Listener) ExitLiteral(ctx *LiteralContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseR2D2Listener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseR2D2Listener) ExitBlock(ctx *BlockContext) {}

// EnterSwitchStatement is called when production switchStatement is entered.
func (s *BaseR2D2Listener) EnterSwitchStatement(ctx *SwitchStatementContext) {}

// ExitSwitchStatement is called when production switchStatement is exited.
func (s *BaseR2D2Listener) ExitSwitchStatement(ctx *SwitchStatementContext) {}

// EnterSwitchCase is called when production switchCase is entered.
func (s *BaseR2D2Listener) EnterSwitchCase(ctx *SwitchCaseContext) {}

// ExitSwitchCase is called when production switchCase is exited.
func (s *BaseR2D2Listener) ExitSwitchCase(ctx *SwitchCaseContext) {}

// EnterDefaultCase is called when production defaultCase is entered.
func (s *BaseR2D2Listener) EnterDefaultCase(ctx *DefaultCaseContext) {}

// ExitDefaultCase is called when production defaultCase is exited.
func (s *BaseR2D2Listener) ExitDefaultCase(ctx *DefaultCaseContext) {}

// EnterJsStatement is called when production jsStatement is entered.
func (s *BaseR2D2Listener) EnterJsStatement(ctx *JsStatementContext) {}

// ExitJsStatement is called when production jsStatement is exited.
func (s *BaseR2D2Listener) ExitJsStatement(ctx *JsStatementContext) {}
