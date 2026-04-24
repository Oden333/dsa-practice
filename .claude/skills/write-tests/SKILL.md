---
name: write-tests
description: Generate comprehensive unit tests for a Go function or problem file. Delegates to the test-writer subagent.
argument-hint: <path-to-file>
agent: test-writer
allowed-tools: Read Write Bash(go test *)
user-invocable: true
---

Делегируй задачу агенту `test-writer`. Файл для покрытия тестами: $ARGUMENTS

Контекст для агента:
- Модуль: `dsa`
- Тестирование: testify assert
- Вспомогательные типы: `dsa/structs` (TreeNode, ListNode, Graph и др.)
- Паттерн именования тест-файла: тот же путь с суффиксом `_test.go`
