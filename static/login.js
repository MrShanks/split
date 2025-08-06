document
  .getElementById("login-form")
  .addEventListener("submit", async function (e) {
    e.preventDefault(); // Prevent default form submission

    const email = document.getElementById("email").value.trim();
    const password = document.getElementById("password").value;
    const errorEl = document.getElementById("login-error");
    errorEl.textContent = "";

    try {
      const response = await fetch("http://localhost:8080/login", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ email, password }),
      });

      if (!response.ok) {
        const errorData = await response.json();
        errorEl.textContent = errorData.message || "Login failed.";
        return;
      }

      const data = await response.json();
      alert(data.message);

      // Store JWT or token if returned
      // localStorage.setItem("token", data.token);

      // Redirect or show dashboard
      // window.location.href = "/dashboard.html";
    } catch (error) {
      console.error("Login error:", error);
      errorEl.textContent = "Network error. Please try again later.";
    }
  });
