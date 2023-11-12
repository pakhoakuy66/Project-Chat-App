import { Credentials, setCredentials } from "../../lib/credentials";

export default async (e: Event) => {
  e.preventDefault();
  const loginInfo = new FormData(e.currentTarget as HTMLFormElement);
  const loginResponse = await fetch("http://localhost:8080/auth/login", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(Object.fromEntries(loginInfo.entries())),
  });
  if (loginResponse.ok) {
    const creds = (await loginResponse.json()) as Credentials;
    setCredentials(creds);
    window.location.href = "/profile/";
  } else {
    window.alert("username or password is incorrect");
  }
};
