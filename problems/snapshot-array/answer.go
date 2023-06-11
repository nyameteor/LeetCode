package main

import "fmt"

type snapshot struct {
	id    int
	value int
}

type SnapshotArray struct {
	numsSnapShots [][]snapshot // Snapshots for each nums[i]
	versionId     int          // Current version id
}

func Constructor(length int) SnapshotArray {
	return SnapshotArray{
		numsSnapShots: make([][]snapshot, length),
		versionId:     0,
	}
}

func (this *SnapshotArray) Set(index int, val int) {
	if len(this.numsSnapShots[index]) == 0 {
		this.numsSnapShots[index] = []snapshot{{this.versionId, val}}
	} else {
		// Update value if we at the same version
		if this.numsSnapShots[index][len(this.numsSnapShots[index])-1].id == this.versionId {
			this.numsSnapShots[index][len(this.numsSnapShots[index])-1].value = val
		} else {
			this.numsSnapShots[index] = append(this.numsSnapShots[index], snapshot{this.versionId, val})
		}
	}
}

func (this *SnapshotArray) Snap() int {
	defer func() {
		this.versionId++
	}()
	return this.versionId
}

func (this *SnapshotArray) Get(index int, snap_id int) int {
	// Binary search to find the index of rightmost snapshot
	snapshots := this.numsSnapShots[index]
	l, r := 0, len(snapshots)-1
	for l <= r {
		m := l + (r-l)/2
		if snapshots[m].id == snap_id {
			return snapshots[m].value
		} else if snapshots[m].id < snap_id {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	if l-1 <= len(snapshots) && l-1 >= 0 {
		return snapshots[l-1].value
	}

	// If no snapshot present, return default value
	return 0
}

/**
 * Your SnapshotArray object will be instantiated and called as such:
 * obj := Constructor(length);
 * obj.Set(index,val);
 * param_2 := obj.Snap();
 * param_3 := obj.Get(index,snap_id);
 */

func main() {
	var obj SnapshotArray

	obj = Constructor(3)
	obj.Set(0, 5)
	fmt.Println(obj.Snap())    // 0
	fmt.Println(obj.Get(0, 0)) // 5

	obj = Constructor(1)
	obj.Set(0, 4)
	obj.Set(0, 16)
	obj.Set(0, 13)
	fmt.Println(obj.Snap())    // 0
	fmt.Println(obj.Get(0, 0)) // 13
	fmt.Println(obj.Snap())    // 1

	obj = Constructor(5)
	obj.Set(0, 49352)
	fmt.Println(obj.Snap())    // 0
	fmt.Println(obj.Get(0, 0)) // 49352
	obj.Set(0, 5373)
	fmt.Println(obj.Snap())    // 1
	fmt.Println(obj.Get(0, 0)) // 49352
	obj.Set(0, 23341)
	fmt.Println(obj.Snap())    // 2
	fmt.Println(obj.Get(0, 0)) // 49352
	obj.Set(0, 12771)
	fmt.Println(obj.Snap())    // 3
	fmt.Println(obj.Get(0, 2)) // 23341
	obj.Set(0, 25105)
	fmt.Println(obj.Snap())    // 4
	fmt.Println(obj.Get(0, 1)) // 5373
}
