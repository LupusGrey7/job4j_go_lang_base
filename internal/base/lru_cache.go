package base

type Node struct {
	Key   string
	Value string
	Prev  *Node
	Next  *Node
}

/*
LruCache -
НОВЫЙ элемент → HEAD (most recently used)
Полный кеш → удаляем TAIL (least recently used)
*/
type LruCache struct {
	size  int
	count int
	Head  *Node // голова
	Tail  *Node // хвост
}

// NewLruCache - this is func
func NewLruCache(size int) *LruCache {
	return &LruCache{
		size: size,
	}
}

// Put - this is method not function.
// Он привязан к объекту l *LruCache, может быть вызван только через него (lru.Put(key, value))
func (l *LruCache) Put(key string, value string) {
	// ШАГ 1: Проверяем дубликат
	for current := l.Head; current != nil; current = current.Next {
		if current.Key == key {
			current.Value = value // Обновляем значение
			l.moveToHead(current) // Перемещаем в HEAD (LRU)
			return
		}
	}

	//При переполнении надо удалять TAIL, а не печатать ошибку.
	// ШАГ 2: Кеш полный? → Удаляем TAIL
	if l.count >= l.size {
		l.removeTail()
	}

	// ШАГ 3: Добавляем НОВЫЙ в HEAD
	newNode := &Node{
		Key:   key,
		Value: value,
		Next:  l.Head, // ← ВАЖНО! Новый становится HEAD
	}

	if l.Head != nil {
		l.Head.Prev = newNode
	}
	l.Head = newNode

	// Если был пустой → HEAD = TAIL
	if l.Tail == nil {
		l.Tail = l.Head
	}

	l.count++
}

// Перемещает ноду в HEAD (LRU логика)
func (l *LruCache) moveToHead(node *Node) {
	if node == l.Head {
		return // Уже в HEAD
	}

	// Отсоединяем от текущего места
	if node.Prev != nil {
		node.Prev.Next = node.Next
	}
	if node.Next != nil {
		node.Next.Prev = node.Prev
	}
	if l.Tail == node {
		l.Tail = node.Prev
	}

	// Добавляем в HEAD
	node.Next = l.Head
	node.Prev = nil
	l.Head.Prev = node
	l.Head = node
}

// Удаляет TAIL (LRU eviction)
func (l *LruCache) removeTail() {
	if l.Tail == nil {
		return
	}

	if l.Tail.Prev != nil {
		l.Tail.Prev.Next = nil
		l.Tail = l.Tail.Prev
	} else {
		l.Head = nil
		l.Tail = nil
	}
	l.count--
}

func (l *LruCache) Get(key string) string {
	// Начинаем с Head и идем до Tail
	for current := l.Head; current != nil; current = current.Next {
		if current.Key == key {
			return current.Value // Возвращаем указатель на Value
		}
	}

	return "" // Ключ не найден
}

func (l *LruCache) GetSize() int {
	return l.size
}
