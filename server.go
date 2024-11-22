package main

import (
    "fmt"
    "log"
    "net/http"
    "os/exec"
    "path/filepath"
)

func main() {
    rootDir := "./phpdocs" 

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        servePHP(w, r, rootDir)
    })

    port := ":8000"
    fmt.Printf("Starting server at http://localhost%s\n", port)
    log.Fatal(http.ListenAndServe(port, nil))
}

func servePHP(w http.ResponseWriter, r *http.Request, rootDir string) {
    path := r.URL.Path
    if path == "/" {
        path = "/index.php"
    }

    fullPath := filepath.Join(rootDir, path)

    if filepath.Ext(fullPath) == ".php" {
        cmd := exec.Command("php", fullPath)
        output, err := cmd.Output()
        if err != nil {
            http.Error(w, "Error executing PHP script", http.StatusInternalServerError)
            fmt.Println("PHP Error:", err)
            return
        }

        w.Header().Set("Content-Type", "text/html")
        w.Write(output)
    } else {
        http.ServeFile(w, r, fullPath)
    }
}
