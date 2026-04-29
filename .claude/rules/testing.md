# Правила написания тестов

## Структура

Всегда использовать table-driven тесты:

```go
func TestFoo(t *testing.T) {
    t.Parallel()

    // необязательный сетап, общий для всех кейсов
    sharedGraph := buildSomeGraph()

    tests := []struct {
        name  string
        // поля входа и ожидаемого результата
        setup func() *Node   // если нужен объект — через setup func
        want  map[int]int
    }{
        { name: "...", ... },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            t.Parallel()
            got := Foo(tt.setup())
            assert.Equal(t, tt.want, got, "short failure hint")
        })
    }
}
```

## Правила

1. **Table-driven всегда.** Несколько последовательных `t.Run` без `[]tests` — антипаттерн.
2. **`t.Parallel()`** — в TestXxx и внутри каждого `t.Run`.
3. **Сетап** — предпочитать поле с сериализованным входом (`args [][]int`, `vals []int`) и вызывать builder в цикле (`structs.BuildGraph(tt.args)`). `setup func() T` — только когда нет удобного сериализованного формата и объект нужно собирать вручную.
4. **Сравнение** — думать под каждый тип:
   - примитивы, срезы, map — `assert.Equal` напрямую
   - указатели на структуры (GraphNode, TreeNode) — сериализовать через вспомогательную функцию (например `PrintGraph`), чтобы diff был читаемым
5. **Сообщение в assert** — одна короткая фраза, только если неочевидно что сравнивается.
6. **Имена кейсов** — строка на английском, описывает сценарий, а не входные данные (`"two sources on linear graph"`, не `"starts=[0,4]"`).
7. **Визуализация сложных кейсов** — для деревьев и графов оставлять ASCII-комментарий прямо внутри тест-кейса:

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
