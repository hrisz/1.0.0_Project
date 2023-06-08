// Update tanggal secara otomatis

var dateTimeInput = document.getElementById("tanggalorder");
var currentDateTime = new Date();

currentDateTime.setDate(currentDateTime.getDay() + 4);
currentDateTime.setHours(currentDateTime.getHours() + 7);

var formattedDateTime = currentDateTime.toISOString().slice(0, 16);
dateTimeInput.value = formattedDateTime;