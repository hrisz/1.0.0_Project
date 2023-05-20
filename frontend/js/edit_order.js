// Ambil data dari URL
const urlParams = new URLSearchParams(window.location.search);
const idOrder = urlParams.get("order_id");
// Ambil elemen formulir
const form = document.getElementById("order-form");
const namaInput = document.getElementById("nama");
const barangInput = document.getElementById("barang");
const lokasiInput = document.getElementById("lokasi");
const masalahInput = document.getElementById("masalah");
// Isi formulir dengan data dari server berdasarkan ID
fetch(`http://127.0.0.1:3000/api/order/${idOrder}`)
  .then((result) => {
    return result.json();
  })
  .then((data) => {
    const order = data.data;
    namaInput.value = order.nama;
    barangInput.value = order.barang;
    lokasiInput.value = order.lokasi;
    masalahInput.value = order.desc_masalah;
  });
// Fungsi untuk mengirim data ke server
function updateData() {
  // Buat objek data yang akan dikirim ke server
  const data = {
    nama: namaInput.value,
    barang: barangInput.value,
    lokasi: lokasiInput.value,
    desc_masalah: masalahInput.value,
  };
  // Buat konfigurasi untuk request
  const options = {
    method: "PUT",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(data),
  };
  // Kirim request ke server
  fetch(`http://127.0.0.1:3000/api/order/${idOrder}`, options)
    .then((result) => {
      return result.json();
    })
    .then((data) => {
      alert("Order berhasil diupdate");
      window.location.href = "index.html";
    })
    .catch((error) => {
      console.error(error);
      alert("Order gagal diupdate");
    });
}
// Tambahkan event listener pada tombol Edit
const editButton = document.querySelector(".btn-primary");
editButton.addEventListener("click", updateData);
