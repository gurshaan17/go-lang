package tasks

import (
    "encoding/csv"
    "fmt"
    "os"
    "strconv"
    "time"

    "golang.org/x/sys/unix"
)

// Task represents a single to-do item.
type Task struct {
    ID          int
    Description string
    CreatedAt   time.Time
    IsComplete  bool
}

// LoadTasks reads tasks from a CSV file.
func LoadTasks(filePath string) ([]Task, error) {
    f, err := loadFile(filePath)
    if err != nil {
        return nil, err
    }
    defer closeFile(f)

    reader := csv.NewReader(f)
    records, err := reader.ReadAll()
    if err != nil {
        return nil, fmt.Errorf("failed to read file: %w", err)
    }

    // Check if there are any records (skip reading header if file is empty)
    if len(records) < 1 {
        return []Task{}, nil // return an empty slice if there are no records
    }

    var tasks []Task
    for _, record := range records[1:] { // Skip the header row
        task, err := parseCSVLine(record)
        if err != nil {
            return nil, err
        }
        tasks = append(tasks, task)
    }

    return tasks, nil
}

// SaveTasks writes tasks to a CSV file.
func SaveTasks(filePath string, tasks []Task) error {
    f, err := loadFile(filePath)
    if err != nil {
        return err
    }
    defer closeFile(f)

    writer := csv.NewWriter(f)
    defer writer.Flush()

    // Write header
    writer.Write([]string{"ID", "Description", "CreatedAt", "IsComplete"})
    for _, task := range tasks {
        writer.Write([]string{
            strconv.Itoa(task.ID),
            task.Description,
            task.CreatedAt.Format(time.RFC3339),
            strconv.FormatBool(task.IsComplete),
        })
    }
    return nil
}

// Helper functions for file handling.
func loadFile(filepath string) (*os.File, error) {
    f, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE, os.ModePerm)
    if err != nil {
        return nil, fmt.Errorf("failed to open file for reading")
    }

    // Exclusive lock obtained on the file descriptor
    if err := unix.Flock(int(f.Fd()), unix.LOCK_EX); err != nil {
        _ = f.Close()
        return nil, err
    }

    return f, nil
}

func closeFile(f *os.File) error {
    unix.Flock(int(f.Fd()), unix.LOCK_UN)
    return f.Close()
}

func parseCSVLine(line []string) (Task, error) {
    id, err := strconv.Atoi(line[0])
    if err != nil {
        return Task{}, err
    }

    createdAt, err := time.Parse(time.RFC3339, line[2])
    if err != nil {
        return Task{}, err
    }

    isComplete, err := strconv.ParseBool(line[3])
    if err != nil {
        return Task{}, err
    }

    return Task{
        ID:          id,
        Description: line[1],
        CreatedAt:   createdAt,
        IsComplete:  isComplete,
    }, nil
}