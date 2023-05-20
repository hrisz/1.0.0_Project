// Ambil elemen formulir
const form = document.getElementById("order-form");
const namaInput = document.getElementById("nama");
const tanggalInput = document.getElementById("tanggal")
const barangInput = document.getElementById("barang");
const lokasiInput = document.getElementById("lokasi");
const masalahInput = document.getElementById("masalah");
// Tambahkan event listener pada tombol "Input Data"
const inputButton = document.querySelector(".btn-primary");
inputButton.addEventListener("click", () => {
  // Ambil nilai input dari formulir
  const nama = namaInput.value;
  const tanggal = tanggalInput.value;
  const barang = barangInput.value;
  const lokasi = lokasiInput.value;
  const masalah = masalahInput.value;
  // Validasi input
  if (!nama || !barang || !lokasi || !masalah) {
    alert("Silakan lengkapi semua field");
    return;
  }
  // Buat objek data yang akan dikirim ke server
  const data = {
    nama: nama,
    tanggal: tanggal,
    barang: barang,
    lokasi: lokasi,
    desc_masalah: masalah,
  };
  // Kirim permintaan POST ke server untuk menambahkan data
  fetch("http://127.0.0.1:3000/api/order", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(data),
  })
    .then((response) => response.json())
    .then((result) => {
      console.log("Success:", result);
      alert("Order berhasil ditambahkan!");
      window.location.href = "index.html";
    })
    .catch((error) => {
      console.error("Error:", error);
      alert("Terjadi kesalahan saat mengirim pesanan.");
    });
});
