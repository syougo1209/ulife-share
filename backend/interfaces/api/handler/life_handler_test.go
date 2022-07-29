package handler

import (
	"net/http"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	setUp()
	// 各テストケースを実行します。今回だと`TestGetTodos`と`TestPostTodo`です。
	code := m.Run()
	os.Exit(code)
}

func setUp() {
	// テスト用のモックをDIし`Router`のポインタを取得しています。
	// `MockTodoController`は`TodoController`インタフェースの関数を全て実装しているのでDI可能です。
	life
	// テストを実行するマルチプレクサを生成
	mux = http.NewServeMux()
	// マルチプレクサにURLとテスト対象のハンドラ関数を関連付けます。
	mux.HandleFunc("/todos/", target.HandleTodosRequest)
}
