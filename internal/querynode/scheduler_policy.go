package querynode

/*
#cgo pkg-config: milvus_segcore milvus_common

#include "segcore/collection_c.h"
#include "segcore/segment_c.h"
#include "segcore/segcore_init_c.h"
#include "common/init_c.h"

*/
import "C"

import (
	"container/list"
	"fmt"
)

type scheduleReadTaskPolicy func(sqTasks *list.List) []readTask

func defaultScheduleReadPolicy(sqTasks *list.List) []readTask {
	// for e := sqTasks.Front(); e != nil; e = e.Next() {
	// 	t, ok := e.Value.(*searchTask)
	// 	if !ok {
	// 		fmt.Println("FUCK")
	// 	}
	// 	fmt.Print(t.NQ, " ")
	// }
	// fmt.Println()
	var ret []readTask
	var next *list.Element
	cnt := int64(0)
	for e := sqTasks.Front(); e != nil; e = next {
		next = e.Next()
		t, ok := e.Value.(readTask)
		if !ok {
			fmt.Println("FUCK")
		}
		tt, ok := e.Value.(*searchTask)
		if !ok {
			fmt.Println("FUCKFUCK")
		}
		nq := tt.NQ
		nt := int64(C.Segcoreabc())
		cnt += nq
		if nt+cnt > 100 || (nt+cnt > 50 && nq < 40) {
			continue
		}
		sqTasks.Remove(e)
		rateCol.rtCounter.sub(t, readyQueueType)
		ret = append(ret, t)
	}
	return ret
}
