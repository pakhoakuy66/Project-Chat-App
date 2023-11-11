export type Credentials = {
  jwt: string;
  expiredAt: number;
};

export function getCredentials() {
  const creds = localStorage.getItem("credentials");
  if (creds === null) {
    return null;
  }
  return JSON.parse(creds) as Credentials;
}

export function setCredentials(creds: Credentials) {
  localStorage.setItem("credentials", JSON.stringify(creds));
}

export function clearCredentials() {
  localStorage.removeItem("credentials");
}
