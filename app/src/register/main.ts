import createRegisterHandler from "./modules/handleRegister";
import "/src/index.postcss";

const registerForm = document.querySelector("[data-register-form]");
const passwordField: HTMLInputElement | null =
  document.querySelector("#password");
const confirmPasswordField: HTMLInputElement | null =
  document.querySelector("#confirm-password");

if (passwordField !== null && confirmPasswordField !== null)
  registerForm?.addEventListener(
    "submit",
    createRegisterHandler(passwordField, confirmPasswordField),
  );
