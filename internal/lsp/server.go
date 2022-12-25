package lsp

import (
	"context"
	"fmt"
	"time"

	"github.com/thegrumpylion/jot/internal/adoc"
	"go.lsp.dev/jsonrpc2"
	"go.lsp.dev/protocol"
)

// server is a language server.
type server struct {
	ver        string
	workspaces map[string]*Workspace
}

// NewServer creates a new server.
func NewServer(ver string) *server {
	return &server{
		ver:        ver,
		workspaces: make(map[string]*Workspace),
	}
}

// Serve is the entry point for the language server.
func (s *server) Serve(ctx context.Context, to time.Duration) error {
	sh := protocol.ServerHandler(s, nil)
	hs := jsonrpc2.HandlerServer(sh)
	return jsonrpc2.ListenAndServe(ctx, "tcp", "0.0.0.0:56857", hs, to)
}

func (s *server) Initialize(ctx context.Context, params *protocol.InitializeParams) (result *protocol.InitializeResult, err error) {
	fmt.Println("initialize", params.RootPath)

	s.workspaces[params.RootPath] = NewWorkspace(params.RootPath)

	res := &protocol.InitializeResult{
		Capabilities: protocol.ServerCapabilities{
			TextDocumentSync: protocol.TextDocumentSyncKindFull,
			CompletionProvider: &protocol.CompletionOptions{
				ResolveProvider: true,
			},
			CodeLensProvider: &protocol.CodeLensOptions{},
		},
		ServerInfo: &protocol.ServerInfo{
			Name:    "jotls",
			Version: s.ver,
		},
	}
	return res, nil
}

func (s *server) Initialized(ctx context.Context, params *protocol.InitializedParams) (err error) {
	fmt.Println("initialized", *params)
	return nil
}

func (s *server) Shutdown(ctx context.Context) (err error) {
	fmt.Println("shutdown, bye!")
	return nil
}

func (s *server) Exit(ctx context.Context) (err error) {
	fmt.Println("exit")
	return nil
}

func (s *server) WorkDoneProgressCancel(ctx context.Context, params *protocol.WorkDoneProgressCancelParams) (err error) {
	panic("not implemented") // TODO: Implement
}

func (s *server) LogTrace(ctx context.Context, params *protocol.LogTraceParams) (err error) {
	panic("not implemented") // TODO: Implement
}

func (s *server) SetTrace(ctx context.Context, params *protocol.SetTraceParams) (err error) {
	panic("not implemented") // TODO: Implement
}

func (s *server) CodeAction(ctx context.Context, params *protocol.CodeActionParams) (result []protocol.CodeAction, err error) {
	panic("not implemented") // TODO: Implement
}

func (s *server) CodeLens(ctx context.Context, params *protocol.CodeLensParams) (result []protocol.CodeLens, err error) {
	panic("not implemented") // TODO: Implement
}

func (s *server) CodeLensResolve(ctx context.Context, params *protocol.CodeLens) (result *protocol.CodeLens, err error) {
	panic("not implemented") // TODO: Implement
}

func (s *server) ColorPresentation(ctx context.Context, params *protocol.ColorPresentationParams) (result []protocol.ColorPresentation, err error) {
	panic("not implemented") // TODO: Implement
}

func (s *server) Completion(ctx context.Context, params *protocol.CompletionParams) (result *protocol.CompletionList, err error) {
	fmt.Println("completion", *params)
	fn := params.TextDocument.URI.Filename()

	fmt.Println("fn", fn)
	doc, err := adoc.ParseFile(fn)
	if err != nil {
		fmt.Println("error parsing file", err)
	}
	fmt.Println("doc", doc)

	return &protocol.CompletionList{
		IsIncomplete: false,
		Items: []protocol.CompletionItem{
			{
				Label: "foo",
				Kind:  protocol.CompletionItemKindText,
				Data:  1,
			},
			{
				Label: "bar",
				Kind:  protocol.CompletionItemKindText,
				Data:  2,
			},
			{
				Label: "bar",
				Kind:  protocol.CompletionItemKindKeyword,
				Data:  3,
			},
		},
	}, nil
}

func (s *server) CompletionResolve(ctx context.Context, params *protocol.CompletionItem) (result *protocol.CompletionItem, err error) {
	panic("not implemented") // TODO: Implement
}

func (s *server) Declaration(ctx context.Context, params *protocol.DeclarationParams) (result []protocol.Location, err error) {
	panic("not implemented") // TODO: Implement
}

func (s *server) Definition(ctx context.Context, params *protocol.DefinitionParams) (result []protocol.Location, err error) {
	panic("not implemented") // TODO: Implement
}

func (s *server) DidChange(ctx context.Context, params *protocol.DidChangeTextDocumentParams) (err error) {
	fmt.Println("didChange")

	// with full sync we expect only one change
	if len(params.ContentChanges) != 1 {
		return fmt.Errorf("expected one change, got %d", len(params.ContentChanges))
	}

	// get document text
	text := params.ContentChanges[0].Text

	fmt.Println("text", text)

	return nil
}

func (s *server) DidChangeConfiguration(ctx context.Context, params *protocol.DidChangeConfigurationParams) (err error) {
	panic("not implemented") // TODO: Implement
}

func (s *server) DidChangeWatchedFiles(ctx context.Context, params *protocol.DidChangeWatchedFilesParams) (err error) {
	panic("not implemented") // TODO: Implement
}

func (s *server) DidChangeWorkspaceFolders(ctx context.Context, params *protocol.DidChangeWorkspaceFoldersParams) (err error) {
	panic("not implemented") // TODO: Implement
}

func (s *server) DidClose(ctx context.Context, params *protocol.DidCloseTextDocumentParams) (err error) {
	panic("not implemented") // TODO: Implement
}

func (s *server) DidOpen(ctx context.Context, params *protocol.DidOpenTextDocumentParams) (err error) {
	fmt.Println("didOpen", *params)
	return nil
}

func (s *server) DidSave(ctx context.Context, params *protocol.DidSaveTextDocumentParams) (err error) {
	fmt.Println("didSave", *params)
	return nil
}

func (s *server) DocumentColor(ctx context.Context, params *protocol.DocumentColorParams) (result []protocol.ColorInformation, err error) {
	panic("not implemented") // TODO: Implement
}

func (s *server) DocumentHighlight(ctx context.Context, params *protocol.DocumentHighlightParams) (result []protocol.DocumentHighlight, err error) {
	panic("not implemented") // TODO: Implement
}

func (s *server) DocumentLink(ctx context.Context, params *protocol.DocumentLinkParams) (result []protocol.DocumentLink, err error) {
	panic("not implemented") // TODO: Implement
}

func (s *server) DocumentLinkResolve(ctx context.Context, params *protocol.DocumentLink) (result *protocol.DocumentLink, err error) {
	panic("not implemented") // TODO: Implement
}

func (s *server) DocumentSymbol(ctx context.Context, params *protocol.DocumentSymbolParams) (result []interface{}, err error) {
	panic("not implemented") // TODO: Implement
}

func (s *server) ExecuteCommand(ctx context.Context, params *protocol.ExecuteCommandParams) (result interface{}, err error) {
	panic("not implemented") // TODO: Implement
}

func (s *server) FoldingRanges(ctx context.Context, params *protocol.FoldingRangeParams) (result []protocol.FoldingRange, err error) {
	panic("not implemented") // TODO: Implement
}

func (s *server) Formatting(ctx context.Context, params *protocol.DocumentFormattingParams) (result []protocol.TextEdit, err error) {
	panic("not implemented") // TODO: Implement
}

func (s *server) Hover(ctx context.Context, params *protocol.HoverParams) (result *protocol.Hover, err error) {
	panic("not implemented") // TODO: Implement
}

func (s *server) Implementation(ctx context.Context, params *protocol.ImplementationParams) (result []protocol.Location, err error) {
	panic("not implemented") // TODO: Implement
}

func (s *server) OnTypeFormatting(ctx context.Context, params *protocol.DocumentOnTypeFormattingParams) (result []protocol.TextEdit, err error) {
	panic("not implemented") // TODO: Implement
}

func (s *server) PrepareRename(ctx context.Context, params *protocol.PrepareRenameParams) (result *protocol.Range, err error) {
	panic("not implemented") // TODO: Implement
}

func (s *server) RangeFormatting(ctx context.Context, params *protocol.DocumentRangeFormattingParams) (result []protocol.TextEdit, err error) {
	panic("not implemented") // TODO: Implement
}

func (s *server) References(ctx context.Context, params *protocol.ReferenceParams) (result []protocol.Location, err error) {
	panic("not implemented") // TODO: Implement
}

func (s *server) Rename(ctx context.Context, params *protocol.RenameParams) (result *protocol.WorkspaceEdit, err error) {
	panic("not implemented") // TODO: Implement
}

func (s *server) SignatureHelp(ctx context.Context, params *protocol.SignatureHelpParams) (result *protocol.SignatureHelp, err error) {
	panic("not implemented") // TODO: Implement
}

func (s *server) Symbols(ctx context.Context, params *protocol.WorkspaceSymbolParams) (result []protocol.SymbolInformation, err error) {
	panic("not implemented") // TODO: Implement
}

func (s *server) TypeDefinition(ctx context.Context, params *protocol.TypeDefinitionParams) (result []protocol.Location, err error) {
	panic("not implemented") // TODO: Implement
}

func (s *server) WillSave(ctx context.Context, params *protocol.WillSaveTextDocumentParams) (err error) {
	panic("not implemented") // TODO: Implement
}

func (s *server) WillSaveWaitUntil(ctx context.Context, params *protocol.WillSaveTextDocumentParams) (result []protocol.TextEdit, err error) {
	panic("not implemented") // TODO: Implement
}

func (s *server) ShowDocument(ctx context.Context, params *protocol.ShowDocumentParams) (result *protocol.ShowDocumentResult, err error) {
	panic("not implemented") // TODO: Implement
}

func (s *server) WillCreateFiles(ctx context.Context, params *protocol.CreateFilesParams) (result *protocol.WorkspaceEdit, err error) {
	panic("not implemented") // TODO: Implement
}

func (s *server) DidCreateFiles(ctx context.Context, params *protocol.CreateFilesParams) (err error) {
	panic("not implemented") // TODO: Implement
}

func (s *server) WillRenameFiles(ctx context.Context, params *protocol.RenameFilesParams) (result *protocol.WorkspaceEdit, err error) {
	panic("not implemented") // TODO: Implement
}

func (s *server) DidRenameFiles(ctx context.Context, params *protocol.RenameFilesParams) (err error) {
	panic("not implemented") // TODO: Implement
}

func (s *server) WillDeleteFiles(ctx context.Context, params *protocol.DeleteFilesParams) (result *protocol.WorkspaceEdit, err error) {
	panic("not implemented") // TODO: Implement
}

func (s *server) DidDeleteFiles(ctx context.Context, params *protocol.DeleteFilesParams) (err error) {
	panic("not implemented") // TODO: Implement
}

func (s *server) CodeLensRefresh(ctx context.Context) (err error) {
	panic("not implemented") // TODO: Implement
}

func (s *server) PrepareCallHierarchy(ctx context.Context, params *protocol.CallHierarchyPrepareParams) (result []protocol.CallHierarchyItem, err error) {
	panic("not implemented") // TODO: Implement
}

func (s *server) IncomingCalls(ctx context.Context, params *protocol.CallHierarchyIncomingCallsParams) (result []protocol.CallHierarchyIncomingCall, err error) {
	panic("not implemented") // TODO: Implement
}

func (s *server) OutgoingCalls(ctx context.Context, params *protocol.CallHierarchyOutgoingCallsParams) (result []protocol.CallHierarchyOutgoingCall, err error) {
	panic("not implemented") // TODO: Implement
}

func (s *server) SemanticTokensFull(ctx context.Context, params *protocol.SemanticTokensParams) (result *protocol.SemanticTokens, err error) {
	panic("not implemented") // TODO: Implement
}

func (s *server) SemanticTokensFullDelta(ctx context.Context, params *protocol.SemanticTokensDeltaParams) (result interface{}, err error) {
	panic("not implemented") // TODO: Implement
}

func (s *server) SemanticTokensRange(ctx context.Context, params *protocol.SemanticTokensRangeParams) (result *protocol.SemanticTokens, err error) {
	panic("not implemented") // TODO: Implement
}

func (s *server) SemanticTokensRefresh(ctx context.Context) (err error) {
	panic("not implemented") // TODO: Implement
}

func (s *server) LinkedEditingRange(ctx context.Context, params *protocol.LinkedEditingRangeParams) (result *protocol.LinkedEditingRanges, err error) {
	panic("not implemented") // TODO: Implement
}

func (s *server) Moniker(ctx context.Context, params *protocol.MonikerParams) (result []protocol.Moniker, err error) {
	panic("not implemented") // TODO: Implement
}

func (s *server) Request(ctx context.Context, method string, params interface{}) (result interface{}, err error) {
	panic("not implemented") // TODO: Implement
}
