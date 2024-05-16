package main

import (
    "bufio"
    "fmt"
    "log"
    "net/http"
    "os"
    "strings"
    "time"
)

func main() {
    // Membaca kode autentikasi dari pengguna
    fmt.Print("Masukkan kode autentikasi Warpcast Anda: ")
    reader := bufio.NewReader(os.Stdin)
    authCode, err := reader.ReadString('\n')
    if err != nil {
        log.Fatal("Gagal membaca kode autentikasi:", err)
    }
    authCode = strings.TrimSpace(authCode)

    // Meminta daftar profil dari pengguna
    fmt.Println("Masukkan daftar profil Warpcast yang akan di-like (pisahkan dengan koma): ")
    profilesInput, err := reader.ReadString('\n')
    if err != nil {
        log.Fatal("Gagal membaca daftar profil:", err)
    }
    profiles := strings.Split(strings.TrimSpace(profilesInput), ",")

    // Melakukan auto-like pada setiap profil dalam daftar dengan jeda waktu
    for _, profile := range profiles {
        profile = strings.TrimSpace(profile)
        if profile != "" {
            err := autoLike(profile, authCode)
            if err != nil {
                fmt.Printf("Gagal melakukan like pada profil %s: %v\n", profile, err)
            } else {
                fmt.Printf("Berhasil melakukan like pada profil %s\n", profile)
            }
            // Menambahkan jeda waktu 2 detik antara setiap permintaan like
            time.Sleep(2 * time.Second)
        }
    }
}

// Fungsi untuk melakukan like pada profil Warpcast
func autoLike(profile string, authCode string) error {
    url := fmt.Sprintf("https://api.warpcast.com/profiles/%s/like", profile)

    // Membuat permintaan HTTP dengan metode POST
    req, err := http.NewRequest("POST", url, nil)
    if err != nil {
        return fmt.Errorf("gagal membuat permintaan: %v", err)
    }

    // Menambahkan header autentikasi
    req.Header.Set("Authorization", "Bearer "+authCode)

    // Mengirim permintaan
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return fmt.Errorf("gagal mengirim permintaan: %v", err)
    }
    defer resp.Body.Close()

    // Memeriksa kode status dari respons
    if resp.StatusCode != http.StatusOK {
        return fmt.Errorf("permintaan gagal dengan status: %s", resp.Status)
    }

    return nil
}
