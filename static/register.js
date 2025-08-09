document
  .getElementById("register-form")
  .addEventListener("submit", async function (e) {
    e.preventDefault(); // Prevent default form submission

    const firstname = document.getElementById("firstname").value.trim();
    const lastname = document.getElementById("lastname").value.trim();
    const email = document.getElementById("email").value.trim();
    const password = document.getElementById("password").value;
    const errorEl = document.getElementById("register-error");
    errorEl.textContent = "";

    try {
      const response = await fetch("http://localhost:8080/signup", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ firstname, lastname, email, password }),
      });

      if (!response.ok) {
        const errorData = await response.json();
        errorEl.textContent = errorData.message || "Registration failed.";
        return;
      }

      const data = await response.json();
      alert(data.message);

      // Store JWT or token if returned
      // localStorage.setItem("token", data.token);

      // Redirect or show dashboard
      // window.location.href = "/dashboard.html";
    } catch (error) {
      console.error("Registration error:", error);
      errorEl.textContent = "Network error. Please try again later.";
    }
  });
