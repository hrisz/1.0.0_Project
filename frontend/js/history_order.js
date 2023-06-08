fetch("http://127.0.0.1:3000/api/history")
  .then((result) => {
    return result.json();
  })
  .then((data) => {
    let tableData = "";
    data.data.map((values) => {
      // Manipulasi data order dan masukkan ke dalam bentuk tabel
      tableData += `
                <tr>
				<td>${values.id}</td>
                <td>${values.tanggal}</td>
                <td>${values.nama}</td>
                <td>${values.nomor_hp}</td>
                <td>${values.barang}</td>
                <td>${values.desc_masalah}</td>
                <td>${values.lokasi}</td>
                <td>Pesanan selesai</td>
                <td></td>
                </tr>`;
    });
    // Tampilkan data order ke dalam tabel
    document.getElementById("history-table").innerHTML = tableData;
  })
  .catch((error) => {
    console.log("error", error);
    alert("Terjadi kesalahan pada server");
  });
