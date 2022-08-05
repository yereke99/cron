package main

import(
   "fmt"
   "time"
   "sync"
   cron "github.com/robfig/cron/v3"
)

type Visited struct{
  mux sync.Mutex
  visited map[string]int
}

func (v *Visited) Inc(url string){
  v.mux.Lock()
  defer v.mux.Unlock()
  v.visited[url]++
}

func (v *Visited) Value(url string) int {
  v.mux.Lock()
  defer v.mux.Unlock()

  return v.visited[url]
}


func main(){
  s1 := cron.New()
  s2 := cron.New()
  s3 := cron.New()

  s := Visited{visited: make(map[string]int)}

  s1.AddFunc("@every 0h0m1s", func() {
    for{
      time.Sleep(100*time.Millisecond)
      s.Inc("google.com")
    }
  })
  s2.AddFunc("@every 0h0m2s", func() {
    for{
      time.Sleep(100*time.Millisecond)
      s.Inc("google.kz")
    }
  })
  s2.AddFunc("@every 0h0m3s", func() {
    for{
      time.Sleep(10*time.Millisecond)
      s.Inc("google.ru")
    }
  })

  defer s1.Stop()
  defer s2.Stop()
  defer s3.Stop()

  go s1.Start()
  go s2.Start()  
  go s3.Start()

  time.Sleep(10 * time.Second)
  fmt.Println(s.Value("google.com"))
  fmt.Println(s.Value("google.kz"))
  fmt.Println(s.Value("google.ru"))

}
/*
type Visited struct{
  mux sync.Mutex
  visited map[string]int
}

func (v *Visited) Inc(url string){
   v.mux.Lock()
   v.visited[url]++
   v.mux.Unlock()
}

func (v *Visited) Value(url string) int {
  v.mux.Lock()
  defer v.mux.Unlock()

  return v.visited[url]
}


func main(){
  s1 := cron.New()
  s2 := cron.New()
  s3 := cron.New()

  s := Visited{visited: make(map[string]int)}

  s1.AddFunc("@every 0h0m1s", func() {
    for{
      time.Sleep(100 * time.Millisecond)
      s.Inc("google.com")
    }
  })

  s2.AddFunc("@every 0h0m2s", func() {
    for{
      time.Sleep(100 * time.Millisecond)
      s.Inc("google.kz")
    }
  })

  s3.AddFunc("@every 0h0m2s", func(){
    for{
      time.Sleep(10 * time.Millisecond)
      s.Inc("amazon.com")
    }
  })

  defer s1.Stop()
  defer s2.Stop()
  defer s3.Stop()

  go s1.Start()
  go s2.Start()
  go s3.Start()

  time.Sleep(10 * time.Second)
  fmt.Println(s.Value("google.com"))
  fmt.Println(s.Value("google.kz"))
  fmt.Println(s.Value("amazon.com"))

}
*/
/*
type Visited struct{
  mux sync.Mutex
  visited map[string]int
}

func (v *Visited) Inc(url string){
     v.mux.Lock()
     v.visited[url]++
     v.mux.Unlock()
}

func (v *Visited) Value(url string) int {
    v.mux.Lock()
    defer v.mux.Unlock()
    return v.visited[url]
}


func main(){
  s1 := cron.New()
  s2 := cron.New()
  s3 := cron.New()

  s := Visited{visited: make(map[string]int)}

  s1.AddFunc("@every 0h0m1s", func() {
    for{
      time.Sleep(100 * time.Millisecond)
      s.Inc("google.kz")
    }
  })
  s2.AddFunc("@every 0h0m2s", func() {
    for{
      time.Sleep(100 * time.Millisecond)
      s.Inc("google.com")
    }
  })
  s3.AddFunc("@every 0h0m1s", func() {
    for{
      time.Sleep(100 * time.Millisecond)
      s.Inc("amazon.com")
    }
  })


  defer s1.Stop()
  defer s2.Stop()
  defer s3.Stop()

  go s1.Start()
  go s2.Start()
  go s3.Start()

  time.Sleep(10 * time.Second)

  fmt.Println(s.Value("google.kz"))
  fmt.Println(s.Value("google.com"))
  fmt.Println(s.Value("amazon.com"))


}

*/
/*
func hello_world(){
  for{
    time.Sleep(100 * time.Millisecond)
    fmt.Println("Hello world")
  }

}


func main(){
  scheduler1 := cron.New()
  scheduler2 := cron.New()
  scheduler3 := cron.New()

  scheduler1.AddFunc("@every 0h0m1s", func() {fmt.Println("Every second")})
  scheduler2.AddFunc("@every 0h0m2s", func() {fmt.Println("Every two second")})
  scheduler3.AddFunc("@every 0h0m3s", hello_world)

  defer scheduler1.Stop()
  defer scheduler2.Stop()
  defer scheduler3.Stop()

  go scheduler1.Start()
  go scheduler2.Start()
  go scheduler3.Start()
  time.Sleep(10 * time.Second)
}
*/
/*
func main(){
  ch := make(chan string)

  go func(){ ch <- "Hello I am gorutine!"}()

  msg := <-ch
  fmt.Println(msg)

}
*/
