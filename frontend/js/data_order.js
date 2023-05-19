fetch("http://127.0.0.1:3000/api/order")
	.then((result) => {
		return result.json();
	})
	.then((data) => {
		let tableData = "";
		data.data.map((values) => {
			// Manipulasi data pegawai dan masukkan ke dalam bentuk tabel
			tableData += `
                <tr>
				<td>${values.order_id}</td>
                <td>${values.nama}</td>
                <td>${values.tanggal}</td>
                <td>${values.barang}</td>
                <td>${values.desc_masalah}</td>
                <td>${values.lokasi}</td>
                <td>${values.pembayaran}</td>
                </tr>`;
		});
		// Tampilkan data pegawai ke dalam tabel
		document.getElementById("order-table").innerHTML = tableData;

		// Tambahkan event listener pada setiap tombol delete
		const deleteButtons = document.querySelectorAll('.btn-danger');
		deleteButtons.forEach((button) => {
			button.addEventListener('click', (event) => {
				const order_id = event.target.getAttribute('order-id');
				// Kirim permintaan DELETE ke server
				fetch(`http://127.0.0.1:3000/api/order/${order_id}`, {
					method: 'DELETE'
				})
					.then((response) => {
						if (response.ok) {
							// Sukses menghapus data, lakukan tindakan sesuai kebutuhan (misalnya, memperbarui tampilan tabel)
							// Contoh:
							event.target.parentNode.parentNode.remove(); // Menghapus baris tabel yang sesuai dengan tombol delete yang diklik
						} else {
							// Gagal menghapus data, berikan umpan balik kepada pengguna atau tangani kesalahan
							throw new Error('Gagal menghapus data');
						}
					})
					.catch(error => {
						console.log('error', error);
						alert('Terjadi kesalahan pada server');
					});
			});
		});

		// Tambahkan event listener pada setiap tombol detail
		const detailButtons = document.querySelectorAll('.btn-success');
		detailButtons.forEach((button) => {
			button.addEventListener('click', (event) => {
				const order_id = event.target.getAttribute('order-id');
				window.location.href = `update.html?order_id=${order_id}`;
			});
		});
	})
	.catch(error => {
		console.log('error', error);
		alert('Terjadi kesalahan pada server');
	});