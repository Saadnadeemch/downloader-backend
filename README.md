# 🚀 Direct Video Downloader Backend (Go + yt-dlp)

A high-performance backend built with Go and yt-dlp, designed to fetch direct video URLs from platforms that provide publicly streamable media.  
This backend powers the open-source frontend available here:

👉 **Frontend Repo:** https://github.com/Saadnadeemch/downloader-frontend

---

## 🧩 Overview

This project acts as a lightweight alternative to traditional desktop downloaders.  
Instead of requiring users to install yt-dlp locally, the system provides a browser-based downloader powered by this backend.

It delivers:

- Video metadata  
- High-quality direct source URLs  
- Secure proxy downloading  
- Multi-platform support  

The backend exposes clean and simple routes:

- **Get Video Info** — extracts metadata + direct file URL  
- **Proxy Download** — downloads files safely through the backend

Since many platforms generate temporary URLs that expire quickly, the frontend first fetches video info, then triggers the proxy download to ensure the file is delivered correctly.

The system is optimized for speed, minimal memory usage, and deployment behind an NGINX reverse proxy on a VPS.

---

## ⚡ Supported Platforms

This backend works with any platform that provides directly streamable media URLs.

Currently supported:

- Instagram  
- Facebook  
- Reddit  
- LinkedIn  

More platforms can be added easily — all that’s required is direct-streaming support.

---

## 🛠 Tech Stack

**Core Technologies**

- Go (Golang 1.20+)  
- yt-dlp (system-installed)  
- Standard Go HTTP packages  
- Clean `cmd/server` architecture  
- NGINX as reverse proxy  
- VPS hosting (Ubuntu recommended)

**Why Go?**

- High-performance concurrency  
- Faster execution for downloads  
- Low latency handling  
- Stable under heavy traffic  

---

## 🔌 API Endpoints

### **1) `/api/video?url=...`**

Returns metadata + direct video URL extracted via yt-dlp.


{
  "title": "Sample Video",
  "thumbnail": "https://...",
  "platform": "instagram",
  "directUrl": "https://video-cdn..."
}
⚠ Direct URLs usually expire quickly — use the proxy endpoint for actual downloading.

2) /api/proxy-download
Streams the video directly from the backend to the user.

Required Query:

url= — direct media URL obtained from /api/video

Optional:

title= custom filename

platform= origin platform

This ensures stable downloading with no expiration issues.

🔧 Installation & Setup
Install Go
bash
Copy code
sudo apt install golang
Install yt-dlp
bash
Copy code
sudo apt install yt-dlp
▶ Running the Project
bash
Copy code
git clone https://github.com/Saadnadeemch/downloader-backend
cd downloader-backend

go run cmd/server/main.go
The backend starts immediately — no database or additional services required.

# ScreenShots

[![Screenshot-2025-11-19-143713.png](https://i.postimg.cc/RCdhVXtP/Screenshot-2025-11-19-143713.png)](https://postimg.cc/PNLTSb2Z)

[![Screenshot-2025-11-19-143806.png](https://i.postimg.cc/B6M3s7z3/Screenshot-2025-11-19-143806.png)](https://postimg.cc/sQQq4myH)

[![Screenshot-2025-11-19-143849.png](https://i.postimg.cc/0Q2gtdy3/Screenshot-2025-11-19-143849.png)](https://postimg.cc/v1jSTn9L)

[![Screenshot-2025-11-19-143925.png](https://i.postimg.cc/L5Yb7yFF/Screenshot-2025-11-19-143925.png)](https://postimg.cc/64Kc4fQM)

