document
  .getElementById("login-form")
  .addEventListener("submit", async (event) => {
    event.preventDefault();

    const email = document.getElementById("email").value;
    const username = document.getElementById("username").value;
    const password = document.getElementById("password").value;

    try {
      const response = await fetch("/login", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          eMail: email,
          userName: username,
          userPassword: password,
        }),
      });

      if (!response.ok) {
        throw new Error("Giriş Başarısız");
      }

      const result = await response.json();
      if (result.message === "Giriş Başarılı") {
        localStorage.setItem("token", result.token);
        localStorage.setItem(
          "loggedInUser",
          JSON.stringify({
            userID: result.userID,
            userName: result.userName,
            eMail: result.eMail,
          })
        );
        window.location.href = "blog.html";
      } else {
        throw new Error(result.error);
      }
    } catch (error) {
      document.getElementById("error-message").textContent = error.message;
      document.getElementById("error-message").style.display = "block";
    }
  });
