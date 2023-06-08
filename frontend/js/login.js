var admin = [
  {
    id: "1",
    nama: "Haris",
    email: "admin@teknisi.com",
    password: "haris123",
  },
];

var teknisi = [
  {
    id: "1",
    nama: "Budi Raynar",
    email: "budiraynar@teknisi.com",
    password: "raynarbudi123",
  },
];

function getLogin() {
  var email = document.getElementById("email").value;
  var password = document.getElementById("password").value;

  for (i = 0; i < teknisi.length; i++) {
    if (email == teknisi[i].email && password == teknisi[i].password) {
      console.log("Teknisi logged successfully.");
      alert("Selamat datang di menu Teknisi!");
      window.location.href = "teknisi_menu.html";
      return;
    }
  }

  for (i = 0; i < admin.length; i++) {
    if (email == admin[i].email && password == admin[i].password) {
      console.log("Admin logged successfully.");
      alert("Selamat datang di menu Admin!");
      window.location.href = "admin_menu.html";

      return;
    } else if (!email || !password) {
      console.log("Form belum terisi semua.");
      alert("Harap isi semua bidang!");
    } else {
      console.log("Incorrect email or password");
      alert("Email atau password salah");
    }
  }
}
