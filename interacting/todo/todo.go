package todo

import (
  "encoding/json"
  "errors"
  "fmt"
  "os"
  "time"
)

// item struct represents a ToDo item
type item struct {
  Task string
  Done bool
  CreatedAt time.Time
  CompletedAt time.Time
}

// List represents a list of ToDo items
type List []item

// Add creates a new todo item and appends it to the list
func (l *List) Add(task string) {
  t := item {
    Task: task, 
    Done: false, 
    CreatedAt: time.Now(),
    CompletedAt: time.Time{},
  }

  *l = append(*l, t)
}

// Complete method marks a ToDo item as completed by 
// setting Done = true and CompletedAt to the current time
func (l *List) Complete(i int) error {
  ls := *l

  if i <= 0 || i > len(ls) {
    return fmt.Errorf("Item %d does not exist", i)
  }

  ls[i - 1].Done = true
  ls[i - 1].CompletedAt = time.Now()

  return nil
}

// Delete method deletes a ToDo item from the list
func (l *List) Delete(i int) error {
  ls := *l

  if i <= 0 || i > len(ls) {
    return fmt.Errorf("Item %d does not exist", i)
  }

  *l = append(ls[:i-1], ls[i:]...)

  return nil
}

// Save method encodes the List as JSON and saves it
// using the provided file name
func (l *List) Save(filename string) error {
  js, err := json.Marshal(l)

  if err != nil {
    return err
  }

  return os.WriteFile(filename, js, 0644)
}

// Get method opens the provided file name, decodes 
// the JSON data and parses it into a List
func (l *List) Get(filename string) error {
  file, err := os.ReadFile(filename)

  if err != nil {
    if errors.Is(err, os.ErrNotExist) {
      return nil
    }
    
    return err
  }

  if len(file) == 0 {
    return nil
  }

  return json.Unmarshal(file, l)
}

// String prints out a formatted list
// Implements the fmt.Stringer interface
func (l *List) String() string {
  formatted := ""

  for k, t := range *l {
    prefix := " "
    if t.Done {
      prefix = "X "
    }

    formatted += fmt.Sprintf("%s%d: %s\n", prefix, k+1, t.Task)
  }

  return formatted
}

// Verbose prints out a formatted list
// along with date and time information
func (l *List) Verbose() string {
  formatted := ""

  for k, t := range *l {
    prefix := " "
    if t.Done {
      prefix = "X "
    }

    formatted += fmt.Sprintf("%s%d: %s %s\n", prefix, k+1, t.Task, t.CreatedAt.String())
  }

  return formatted
}

// Open prints out a formatted list
// of only open tasks
func (l *List) Open() string {
  formatted := ""

  for k, t := range *l {
    if !t.Done {
      formatted += fmt.Sprintf("%d: %s\n", k+1, t.Task)
    }
  }

  return formatted
}
