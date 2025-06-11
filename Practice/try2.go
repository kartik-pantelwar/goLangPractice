package main

import (
    "fmt"
    "net/http"
    "sync"
    "time"
)

// Job represents a single web task (checking a URL)
type Job struct {
    ID  int
    URL string
}

// Result represents the outcome of a web check
type Result struct {
    JobID      int
    URL        string
    StatusCode int
    Err        error
}

// worker is a goroutine that takes Jobs from the jobs channel, processes them, and sends Results to the results channel
func worker(id int, jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) {
    //jobs is a receiver channel (that receive jobs)
    //result is a sender channel (that send results)
    defer wg.Done()
    client := &http.Client{
        Timeout: 5 * time.Second,
    }

    for job := range jobs {
        // Try to GET the URL
        res, err := client.Get(job.URL)
        status := 0
        if err == nil {
            status = res.StatusCode
            res.Body.Close()
        }
        results <- Result{
            JobID:      job.ID,
            URL:        job.URL,
            StatusCode: status,
            Err:        err,
        }
    }
}

func main() {
    // List of URLs to check
    urlList := []string{
        "https://www.google.com",
        "https://www.github.com",
        "https://www.golang.org",
        "https://notarealwebsite12345.com",
        "https://www.stackoverflow.com",
        "https://www.reddit.com",
        "https://outlier.org", // Add more URLs as desired
    }

    numWorkers := 3 // Size of worker pool
    jobs := make(chan Job, len(urlList))
    results := make(chan Result, len(urlList))

    var wg sync.WaitGroup

    // Start workers
    for i := 0; i < numWorkers; i++ {
        wg.Add(1)
        go worker(i+1, jobs, results, &wg)
    }

    // Send jobs
    for idx, url := range urlList {
        jobs <- Job{ID: idx + 1, URL: url}
    }
    close(jobs) // No more jobs

    // Wait for workers to finish, then close results channel
    go func() {
        wg.Wait()
        close(results)
    }()

    // Collect and print results
    fmt.Println("Results:")
    for res := range results {
        if res.Err != nil {
            fmt.Printf("[Job %d] %s --> ERROR: %v\n", res.JobID, res.URL, res.Err)
        } else {
            fmt.Printf("[Job %d] %s --> Status Code: %d\n", res.JobID, res.URL, res.StatusCode)
        }
    }

    fmt.Println("All done!")
}