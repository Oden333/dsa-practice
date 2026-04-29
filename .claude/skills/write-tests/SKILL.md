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

## Обязательный стиль тестов

**Table-driven всегда.** Несколько последовательных `t.Run` без `[]tests` — антипаттерн.

```go
func TestFoo(t *testing.T) {
    t.Parallel()

    tests := []struct {
        name  string
        setup func() *T   // для объектов-указателей
        input SomeType    // для примитивов/срезов — напрямую
        want  ResultType
    }{
        { name: "...", ... },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            t.Parallel()
            got := Foo(tt.input)
            assert.Equal(t, tt.want, got)
        })
    }
}
```

**Правила:**
1. `t.Parallel()` — в TestXxx и внутри каждого `t.Run`
2. Сетап — предпочитать поле с сериализованным входом (`args [][]int`, `vals []int`) и вызывать builder в цикле (`structs.BuildGraph(tt.args)`). `setup func() T` — только когда объект нельзя описать сериализованным входом
3. Сравнение — под каждый тип:
   - примитивы, срезы, map — `assert.Equal` напрямую
   - указатели на структуры (GraphNode, TreeNode) — сериализовать через вспомогательную функцию (например `PrintGraph`)
4. Сообщение в assert — только если неочевидно что сравнивается, одна короткая фраза
5. Имена кейсов — описывают сценарий, не входные данные (`"two sources on linear graph"`, не `"starts=[0,4]"`)
6. Для сложных структур (деревья, графы) — оставлять ASCII-комментарий прямо в тест-кейсе:

```go
{
    //     1
    //    / \
    //   2   3
    //    \
    //     5
    name: "right child only on node 2",
    root: structs.BuildTree([]int{1, 2, 3, structs.Null, 5}),
    want: ...,
},
```
