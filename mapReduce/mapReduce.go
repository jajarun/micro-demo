package mapReduce

import "sync"

type MapReduce struct {
	MapResults []any
	wg         sync.WaitGroup
	lock       sync.RWMutex
}

type MapFunc func() any

type ReduceFunc func(mapResult []any)

func (mr *MapReduce) AddMap(mapFunc MapFunc) {
	mr.wg.Add(1)
	go func() {
		//mr.MapResults.mapFunc()
		mapResult := mapFunc()
		mr.lock.Lock()
		mr.MapResults = append(mr.MapResults, mapResult)
		mr.lock.Unlock()
		mr.wg.Done()
	}()
}

func (mr *MapReduce) Reduce(reduceFunc ReduceFunc) {
	mr.wg.Wait()
	//for _, mapResult := range mr.MapResults {
	//	reduceFunc(mapResult)
	//}
	reduceFunc(mr.MapResults)
}
