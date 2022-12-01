function submitData() {
  let name = document.getElementById("name").value;
  let email = document.getElementById("email").value;
  let phone = document.getElementById("phone").value;
  let subject = document.getElementById("subject").value;
  let message = document.getElementById("message").value;
  console.log(name, email, phone, subject, message);

  // Untuk Validasi
  if (name == "") {
    return alert("Nama harus diisi");
  } else if (email == "") {
    return alert("Email harus diisi");
  } else if (phone == "") {
    return alert("Phone harus diisi");
  } else if (subject == "") {
    return alert("Subject harus diisi");
  } else if (message == "") {
    return alert("Message harus diisi");
  }

  //   Untuk Ngelink ke Mail
  let link = document.createElement("a");
  link.href = `mailto:${email}?subject=${subject}&body=Hallo nama saya ${name}, ${message}, Hubungi saya di Whatsapp ${phone}`;
  link.click();

  let student = {
    name,
    email,
    phone,
    subject,
    message,
  };

  console.log(student);
}
