package db

type transaction struct {
	storage map[string]any
	next    *transaction
}

type InMemoryDB struct {
	mainStorage map[string]any
	top         *transaction
}

func NewInMemoryDB() *InMemoryDB {
	return &InMemoryDB{mainStorage: make(map[string]any)}
}

func (db *InMemoryDB) StartTransaction() {
	storage := make(map[string]any)
	if db.top != nil {
		for k, v := range db.top.storage {
			storage[k] = v
		}
	}

	curTr := &transaction{storage: storage}
	curTr.next = db.top
	db.top = curTr
}

func (db *InMemoryDB) RollBackTransaction() {
	if db.top != nil {
		for key := range db.top.storage {
			delete(db.top.storage, key)
		}

		if db.top.next != nil {
			db.top = db.top.next
		}
	}
}

func (db *InMemoryDB) CommitTransaction() {
	curTr := db.top

	if curTr != nil && len(curTr.storage) != 0 {
		for key, value := range curTr.storage {
			db.mainStorage[key] = value
			if curTr.next != nil {
				curTr.next.storage[key] = value
			}
		}
	} else {
		db.mainStorage = make(map[string]any)
	}
}

func (db *InMemoryDB) Get(key string) any {
	curTr := db.top

	if curTr == nil || len(curTr.storage) == 0 {
		if val, ok := db.mainStorage[key]; ok {
			return val
		}
	} else {
		if val, ok := curTr.storage[key]; ok {
			return val
		}
	}

	return nil
}

func (db *InMemoryDB) Set(key string, value any) {
	if key == "" || value == "" {
		return
	}

	curTr := db.top
	if curTr == nil {
		db.mainStorage[key] = value
	} else {
		curTr.storage[key] = value
	}
}

func (db *InMemoryDB) Delete(key string) {
	if db.top == nil {
		delete(db.mainStorage, key)
	} else {
		delete(db.top.storage, key)
	}
}
