import { defineConfig } from "vite";
import { resolve } from "path";

export default defineConfig({
  build: {
    rollupOptions: {
      input: {
        main: resolve(__dirname, "index.html"),
        login: resolve(__dirname, "login/index.html"),
        register: resolve(__dirname, "register/index.html"),
        profile: resolve(__dirname, "profile/index.html"),
        friends: resolve(__dirname, "friends/index.html"),
        friendrequests: resolve(__dirname, "friendrequests/index.html"),
      },
    },
  },
});
