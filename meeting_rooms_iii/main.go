package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func main() {
	fmt.Println(mostBooked(2, [][]int{{0, 10}, {1, 5}, {2, 7}, {3, 4}}))
	fmt.Println(mostBooked(3, [][]int{{1, 20}, {2, 10}, {3, 5}, {4, 9}, {6, 8}}))
}

func mostBooked(n int, meetings [][]int) int {
	// Sort meetings by start time
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][0] < meetings[j][0]
	})

	// Initialize min-heaps
	freeRooms := &MinHeap{}
	heap.Init(freeRooms)
	for i := 0; i < n; i++ {
		heap.Push(freeRooms, Room{0, i})
	}

	// Initialize counters for meetings held in each room
	meetingCount := make([]int, n)

	for _, m := range meetings {
		start, end := m[0], m[1]

		// Get the next available room
		for freeRooms.Len() > 0 && freeRooms[0].availableAt <= start {
			heap.Push(freeRooms, heap.Pop(freeRooms))
		}

		// If no free room is available, get the one that becomes available the earliest
		room := heap.Pop(freeRooms).(Room)
		if room.availableAt <= start {
			room.availableAt = end
		} else {
			room.availableAt += end - start
		}
		meetingCount[room.index]++
		heap.Push(freeRooms, room)
	}

	// Find the room with the maximum number of meetings
	maxMeetings, result := 0, 0
	for i := 0; i < n; i++ {
		if meetingCount[i] > maxMeetings || (meetingCount[i] == maxMeetings && i < result) {
			maxMeetings = meetingCount[i]
			result = i
		}
	}

	return result
}

type Meeting struct {
	start, end int
}

// Room represents a room with its next available time and its index
type Room struct {
	availableAt, index int
}

// MinHeap for available rooms
type MinHeap []Room

func (h MinHeap) Len() int { return len(h) }
func (h MinHeap) Less(i, j int) bool {
	return h[i].availableAt < h[j].availableAt || (h[i].availableAt == h[j].availableAt && h[i].index < h[j].index)
}
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(Room)) }
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]
	return item
}
