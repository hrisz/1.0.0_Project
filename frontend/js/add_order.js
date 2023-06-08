// Ambil elemen formulir
const form = document.getElementById("order-form");
const tanggalInput = document.getElementById("tanggalorder");
const namaInput = document.getElementById("nama");
const kontakInput = document.getElementById("nomor_hp");
const barangInput = document.getElementById("barang");
const lokasiInput = document.getElementById("lokasi");
const masalahInput = document.getElementById("masalah");

// Tambahkan event listener pada tombol "Order!"
const inputButton = document.querySelector(".btn-primary");
inputButton.addEventListener("click", () => {
  // Ambil nilai input dari formulir
  const nama = namaInput.value;
  const tanggal = tanggalInput.value;
  const nomor_hp = kontakInput.value;
  const barang = barangInput.value;
  const lokasi = lokasiInput.value;
  const masalah = masalahInput.value;
  // Validasi input
  if (!nama || !nomor_hp || !barang || !lokasi || !masalah) {
    alert("Silakan lengkapi semua field");
    return;
  }
  // Buat objek data yang akan dikirim ke server
  const data = {
    nama: nama,
    tanggal: tanggal,
    nomor_hp: nomor_hp,
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
      window.location.href = "thanks_order.html";
    })
    .catch((error) => {
      console.error("Error:", error);
      alert("Terjadi kesalahan saat mengirim pesanan.");
    });
});
