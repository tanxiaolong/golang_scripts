package main

import "fmt"
import "container/heap"

type UserScore struct {
    UserID string
    Score  int
}

type UserHeap []UserScore

func (h UserHeap) Len() int           { return len(h) }
func (h UserHeap) Less(i, j int) bool { return h[i].Score < h[j].Score }
func (h UserHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *UserHeap) Push(x interface{}) {
    *h = append(*h, x.(UserScore))
}

func (h *UserHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}

func GetTop10(users []UserScore) []UserScore {
    h := &UserHeap{}
    heap.Init(h)

    for _, user := range users {
        heap.Push(h, user)
        if h.Len() > 10 {
            heap.Pop(h)
        }
    }

    top10 := make([]UserScore, h.Len())
    for i := h.Len() - 1; i >= 0; i-- {
        top10[i] = heap.Pop(h).(UserScore)
    }

    return top10
}

func main(){
  users := []UserScore{
    {UserID: "user1", Score: 100},
    {UserID: "user2", Score: 200},
    {UserID: "user3", Score: 300},
    {UserID: "user4", Score: 400},
    {UserID: "user5", Score: 500},
    {UserID: "user6", Score: 600},
    {UserID: "user7", Score: 700},
    {UserID: "user8", Score: 800},
    {UserID: "user9", Score: 900},
    {UserID: "user10", Score: 1000},
    {UserID: "user11", Score: 1100},
    {UserID: "user12", Score: 1200},
  }  

  top10 := GetTop10(users)
  fmt.Println(top10)
}


