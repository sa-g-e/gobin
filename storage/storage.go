package storage

import (
    "fmt"
    "os"
    "time"
    "sync"
    "encoding/json"
)

type Paste struct {
    ID      string `json:"id"`
    Content string `json:"content"`
    Expire  time.Time `json:"expire"`
}

var (
    mu sync.Mutex
    storageDir = "data"
)

func SavePaste(content string, duration time.Duration) (string, error) {
    mu.Lock()
    defer mu.Unlock()

    id := generateID()
    paste := Paste{
        ID:      id,
        Content: content,
        Expire:  time.Now().Add(duration),
    }

    data, err := json.Marshal(paste)
    if err != nil {
        return "", err
    }

    if err := os.WriteFile(fmt.Sprintf("%s/%s.json", storageDir, id), data, 0644); err != nil {
        return "", err
    }

    return id, nil
}

func LoadPaste(id string) (*Paste, error) {
    mu.Lock()
    defer mu.Unlock()

    data, err := os.ReadFile(fmt.Sprintf("%s/%s.json", storageDir, id))
    if err != nil {
        return nil, err
    }

    var paste Paste
    if err := json.Unmarshal(data, &paste); err != nil {
        return nil, err
    }

    if time.Now().After(paste.Expire) {
		os.Remove(fmt.Sprintf("%s/%s.json", storageDir, id))
        return nil, fmt.Errorf("paste expired")
    }

    return &paste, nil
}

func generateID() string {
    return fmt.Sprintf("%d", time.Now().UnixNano())
}
