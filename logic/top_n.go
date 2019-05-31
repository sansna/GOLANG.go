/*
 * Copyright (C) Wentao Wang
 * Date: 19 Fri May.31 11:01
 */

package logic

type linkstruct struct {
	r   *linklist
	cur *linklist
}
type linklist struct {
	Thing      interface{}
	Prev, Next *linklist
}

type topn interface {
	N() int
	Len() int
	Less(i, j interface{}) bool
	Get(i int) interface{}
}

func FindN(data topn) (ret []interface{}) {
	// init
	lr := linkstruct{
		r: &linklist{},
	}
	lr.cur = lr.r
	for i := 0; i < data.N()-1; i++ {
		lr.cur.Next = &linklist{}
		lr.cur.Next.Prev = lr.cur
		lr.cur = lr.cur.Next
	}
	lr.cur.Next = lr.r
	lr.r.Prev = lr.cur
	lr.cur = lr.r
	// init done

	for i := 0; i < data.Len(); i++ {
		item := data.Get(i)
		if i < data.N() {
			lr.cur.Thing = data.Get(i)
			lr.cur = lr.cur.Next
			if i > 0 {
				y := lr.cur.Prev
				x := y.Prev
				var l *linklist
				shouldReposition := false
				for l = lr.r; l.Prev != x; l = l.Next {
					if data.Less(item, l.Thing) {
						shouldReposition = true
						break
					}
				}
				if shouldReposition {
					y.Prev.Next = y.Next
					y.Next.Prev = x
					y.Prev = l.Prev
					y.Next = l
					y.Prev.Next = y
					l.Prev = y
					if l == lr.r {
						// re-head
						lr.r = y
					}
				}
			}
		} else {
			shouldReplace := false
			var l *linklist
			for l = lr.r; l != lr.r.Prev; l = l.Next {
				if data.Less(item, l.Thing) {
					shouldReplace = true
					break
				}
			}
			if shouldReplace {
				y := lr.r.Prev
				y.Thing = item

				if l == lr.r {
					lr.r = y
				} else {
					y.Prev.Next = y.Next
					y.Next.Prev = y.Prev

					y.Prev = l.Prev
					y.Next = l
					l.Prev = y
					y.Prev.Next = y
				}
			} else if !shouldReplace && data.Less(item, lr.r.Prev.Thing) {
				lr.r.Prev.Thing = item
			}
		}
	}

	l := lr.r
	for i := 0; i < data.N(); i++ {
		if l.Thing != nil {
			ret = append(ret, l.Thing)
		}
		l = l.Next
	}
	return ret
}
