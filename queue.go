type Deque struct {
    sync.RWMutex
    not_empty_notify        chan int
    container *list.List
}
func (s *Deque) Put(item interface{}) {
    s.Lock()
    s.container.PushFront(item)
    s.Unlock()
    select {
    case s.not_empty_notify <-1:
    default:
    }
}
func (s *Deque) Get(timeout int) (interface{}, error) {
    s.Lock()
    var item interface{} = nil
    var lastContainerItem *list.Element = nil
    endTime := time.Now().Add(time.Duration(timeout) * time.Second)
    for {
        if s.container.Back() != nil {
            break
        }   
        remaining := endTime.Sub(time.Now())
        s.Unlock()
        if remaining < 0 {
            return nil, errors.New("time out in Pop")
        }
        select {
        case <-s.not_empty_notify:
        case <-time.After(remaining):
        return nil, errors.New("time out in Pop")
        }
        s.Lock()
    }
    lastContainerItem = s.container.Back()
    item = s.container.Remove(lastContainerItem)
    s.Unlock()
    return item, nil
}
 
queueFreeServer := Deque{ container: list.New(),  not_empty_notify: make(chan int)}

