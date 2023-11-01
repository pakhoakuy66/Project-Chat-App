import handleLogin from "./modules/handleLogin";
import "/src/index.postcss";

const loginForm = document.querySelector("[data-login-form]");

loginForm?.addEventListener("submit", handleLogin);
