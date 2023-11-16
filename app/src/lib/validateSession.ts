import { Credentials, getCredentials, setCredentials } from "./credentials";

export default async function validateSession() {
  const creds = getCredentials();
  if (creds === null) {
    return null;
  } else if (creds.expiredAt - Date.now() < 300_000) {
    const refreshRequest = await fetch("http://localhost:8080/auth/refresh", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${creds.jwt}`,
      },
    });
    if (refreshRequest.ok) {
      const newCreds = (await refreshRequest.json()) as Credentials;
      setCredentials(newCreds);
      return newCreds;
    } else {
      return null;
    }
  }
  return creds;
}
