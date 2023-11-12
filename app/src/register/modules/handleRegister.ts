const passwordField: HTMLInputElement | null =
  document.querySelector("#password");
const confirmPasswordField: HTMLInputElement | null =
  document.querySelector("#confirm-password");

export default async (e: Event) => {
  e.preventDefault();
  if (passwordField?.value !== confirmPasswordField?.value) {
    window.alert("Confirm password is incorrect");
    return;
  }
  const registerInfo = new FormData(e.currentTarget as HTMLFormElement);
  const birthday = registerInfo.get("birthday");
  if (birthday !== null)
    registerInfo.set("birthday", birthday.toString() + "T00:00:00Z");
  const registerResponse = await fetch("http://localhost:8080/auth/register", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(Object.fromEntries(registerInfo.entries())),
  });
  if (registerResponse.ok) {
    window.location.href = "/login/";
  } else {
    window.alert("something went wrong");
  }
};
