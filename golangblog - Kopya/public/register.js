document
  .getElementById("register-form")
  .addEventListener("submit", async function (e) {
    e.preventDefault();

    const email = document.getElementById("email").value;
    const username = document.getElementById("username").value;
    const password = document.getElementById("password").value;

    const response = await fetch("/register", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        email,
        userName: username,
        userPassword: password,
      }),
    });

    const data = await response.json();

    if (response.ok) {
      window.location.href = "/index.html";
    } else {
      document.getElementById("error-message").innerText =
        data.message || "Kayıt sırasında bir hata oluştu";
      document.getElementById("error-message").style.display = "block";
    }
  });
