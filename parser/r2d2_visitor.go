// Code generated from R2D2.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // R2D2

import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by R2D2Parser.
type R2D2Visitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by R2D2Parser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by R2D2Parser#declaration.
	VisitDeclaration(ctx *DeclarationContext) interface{}

	// Visit a parse tree produced by R2D2Parser#globalDeclaration.
	VisitGlobalDeclaration(ctx *GlobalDeclarationContext) interface{}

	// Visit a parse tree produced by R2D2Parser#importDeclaration.
	VisitImportDeclaration(ctx *ImportDeclarationContext) interface{}

	// Visit a parse tree produced by R2D2Parser#interfaceDeclaration.
	VisitInterfaceDeclaration(ctx *InterfaceDeclarationContext) interface{}

	// Visit a parse tree produced by R2D2Parser#moduleDeclaration.
	VisitModuleDeclaration(ctx *ModuleDeclarationContext) interface{}

	// Visit a parse tree produced by R2D2Parser#functionDeclaration.
	VisitFunctionDeclaration(ctx *FunctionDeclarationContext) interface{}

	// Visit a parse tree produced by R2D2Parser#functionCallStatement.
	VisitFunctionCallStatement(ctx *FunctionCallStatementContext) interface{}

	// Visit a parse tree produced by R2D2Parser#functionCall.
	VisitFunctionCall(ctx *FunctionCallContext) interface{}

	// Visit a parse tree produced by R2D2Parser#parameterList.
	VisitParameterList(ctx *ParameterListContext) interface{}

	// Visit a parse tree produced by R2D2Parser#parameter.
	VisitParameter(ctx *ParameterContext) interface{}

	// Visit a parse tree produced by R2D2Parser#typeExpression.
	VisitTypeExpression(ctx *TypeExpressionContext) interface{}

	// Visit a parse tree produced by R2D2Parser#arrayDimensions.
	VisitArrayDimensions(ctx *ArrayDimensionsContext) interface{}

	// Visit a parse tree produced by R2D2Parser#baseType.
	VisitBaseType(ctx *BaseTypeContext) interface{}

	// Visit a parse tree produced by R2D2Parser#genericType.
	VisitGenericType(ctx *GenericTypeContext) interface{}

	// Visit a parse tree produced by R2D2Parser#typeDeclaration.
	VisitTypeDeclaration(ctx *TypeDeclarationContext) interface{}

	// Visit a parse tree produced by R2D2Parser#variableDeclaration.
	VisitVariableDeclaration(ctx *VariableDeclarationContext) interface{}

	// Visit a parse tree produced by R2D2Parser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by R2D2Parser#expressionStatement.
	VisitExpressionStatement(ctx *ExpressionStatementContext) interface{}

	// Visit a parse tree produced by R2D2Parser#ifStatement.
	VisitIfStatement(ctx *IfStatementContext) interface{}

	// Visit a parse tree produced by R2D2Parser#forStatement.
	VisitForStatement(ctx *ForStatementContext) interface{}

	// Visit a parse tree produced by R2D2Parser#assignmentDeclaration.
	VisitAssignmentDeclaration(ctx *AssignmentDeclarationContext) interface{}

	// Visit a parse tree produced by R2D2Parser#assignment.
	VisitAssignment(ctx *AssignmentContext) interface{}

	// Visit a parse tree produced by R2D2Parser#assignmentOperator.
	VisitAssignmentOperator(ctx *AssignmentOperatorContext) interface{}

	// Visit a parse tree produced by R2D2Parser#simpleFor.
	VisitSimpleFor(ctx *SimpleForContext) interface{}

	// Visit a parse tree produced by R2D2Parser#whileStatement.
	VisitWhileStatement(ctx *WhileStatementContext) interface{}

	// Visit a parse tree produced by R2D2Parser#loopStatement.
	VisitLoopStatement(ctx *LoopStatementContext) interface{}

	// Visit a parse tree produced by R2D2Parser#cicleControl.
	VisitCicleControl(ctx *CicleControlContext) interface{}

	// Visit a parse tree produced by R2D2Parser#breakStatement.
	VisitBreakStatement(ctx *BreakStatementContext) interface{}

	// Visit a parse tree produced by R2D2Parser#continueStatement.
	VisitContinueStatement(ctx *ContinueStatementContext) interface{}

	// Visit a parse tree produced by R2D2Parser#returnStatement.
	VisitReturnStatement(ctx *ReturnStatementContext) interface{}

	// Visit a parse tree produced by R2D2Parser#arrayAccessExpression.
	VisitArrayAccessExpression(ctx *ArrayAccessExpressionContext) interface{}

	// Visit a parse tree produced by R2D2Parser#additiveExpression.
	VisitAdditiveExpression(ctx *AdditiveExpressionContext) interface{}

	// Visit a parse tree produced by R2D2Parser#identifierExpression.
	VisitIdentifierExpression(ctx *IdentifierExpressionContext) interface{}

	// Visit a parse tree produced by R2D2Parser#functionCallExpression.
	VisitFunctionCallExpression(ctx *FunctionCallExpressionContext) interface{}

	// Visit a parse tree produced by R2D2Parser#parenthesisExpression.
	VisitParenthesisExpression(ctx *ParenthesisExpressionContext) interface{}

	// Visit a parse tree produced by R2D2Parser#comparisonExpression.
	VisitComparisonExpression(ctx *ComparisonExpressionContext) interface{}

	// Visit a parse tree produced by R2D2Parser#multiplicativeExpression.
	VisitMultiplicativeExpression(ctx *MultiplicativeExpressionContext) interface{}

	// Visit a parse tree produced by R2D2Parser#literalExpression.
	VisitLiteralExpression(ctx *LiteralExpressionContext) interface{}

	// Visit a parse tree produced by R2D2Parser#unaryExpression.
	VisitUnaryExpression(ctx *UnaryExpressionContext) interface{}

	// Visit a parse tree produced by R2D2Parser#logicalExpression.
	VisitLogicalExpression(ctx *LogicalExpressionContext) interface{}

	// Visit a parse tree produced by R2D2Parser#memberExpression.
	VisitMemberExpression(ctx *MemberExpressionContext) interface{}

	// Visit a parse tree produced by R2D2Parser#memberPart.
	VisitMemberPart(ctx *MemberPartContext) interface{}

	// Visit a parse tree produced by R2D2Parser#argumentList.
	VisitArgumentList(ctx *ArgumentListContext) interface{}

	// Visit a parse tree produced by R2D2Parser#primaryExpression.
	VisitPrimaryExpression(ctx *PrimaryExpressionContext) interface{}

	// Visit a parse tree produced by R2D2Parser#arrayLiteral.
	VisitArrayLiteral(ctx *ArrayLiteralContext) interface{}

	// Visit a parse tree produced by R2D2Parser#literal.
	VisitLiteral(ctx *LiteralContext) interface{}

	// Visit a parse tree produced by R2D2Parser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by R2D2Parser#switchStatement.
	VisitSwitchStatement(ctx *SwitchStatementContext) interface{}

	// Visit a parse tree produced by R2D2Parser#switchCase.
	VisitSwitchCase(ctx *SwitchCaseContext) interface{}

	// Visit a parse tree produced by R2D2Parser#defaultCase.
	VisitDefaultCase(ctx *DefaultCaseContext) interface{}

	// Visit a parse tree produced by R2D2Parser#jsStatement.
	VisitJsStatement(ctx *JsStatementContext) interface{}
}
