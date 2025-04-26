# Folder Konfigurasi Monitoring (monitored/)

Folder ini berisi file konfigurasi untuk stack observability dan monitoring pada proyek Nusaloka. Berikut adalah beberapa file yang termasuk dalam folder ini:

## Isi Folder:
1. **`prometheus.yml`**: Konfigurasi untuk Prometheus, yang digunakan untuk monitoring dan scraping metrics dari berbagai service.
2. **`loki-config.yaml`**: Konfigurasi untuk Grafana Loki, digunakan untuk log aggregation dan penyimpanan log dari aplikasi.
3. **`grafana-datasource.yml`** (Opsional): Konfigurasi sumber data untuk Grafana, jika diperlukan untuk setup otomatis.

## Tujuan Folder Ini:
- **Observability**: Menyediakan monitoring, logging, dan alerting untuk proyek.
- **Version Control**: Menyimpan dan mengelola file konfigurasi monitoring agar bisa dikelola secara terpusat dan terverifikasi.

## Penggunaan:
File-file ini digunakan oleh Docker containers dan tools lainnya (seperti Prometheus dan Loki) untuk mengumpulkan data observabilitas dan menyajikannya di Grafana.

## Penting:
Pastikan untuk **tidak menyertakan kredensial sensitif** dalam file konfigurasi ini. Jika ada informasi yang sensitif, simpan dalam variabel lingkungan atau volume Docker secara terpisah.
