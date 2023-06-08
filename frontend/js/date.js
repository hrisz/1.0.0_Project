function displayDate() {
  var waktuSekarang = new Date();
  var hari = waktuSekarang.getDay() + 4;
  var bulan = waktuSekarang.getMonth() + 1;
  var tahun = waktuSekarang.getFullYear();
  var dateString = hari + "/" + bulan + "/" + tahun;
  document.getElementById("tanggal").innerHTML = dateString;
}

function displayTime() {
  var currentTime = new Date();
  var hours = currentTime.getHours();
  var minutes = currentTime.getMinutes();
  var timeString = hours + ":" + minutes;
  document.getElementById("waktu").innerHTML = timeString;
}

displayTime();
displayDate();