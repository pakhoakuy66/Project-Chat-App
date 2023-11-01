export default async (e: Event) => {
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
};
