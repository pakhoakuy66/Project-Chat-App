import "/src/index.postcss";

const loginForm = document.querySelector("[data-login-form]");

loginForm?.addEventListener("submit", async (e) => {
  e.preventDefault();
  const loginInfo = new FormData(e.currentTarget as HTMLFormElement);
  const loginResponse = await fetch("http://localhost:8080/auth/login", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    credentials: "include",
    body: JSON.stringify(Object.fromEntries(loginInfo.entries())),
  });
  if (loginResponse.ok) {
  }
});
