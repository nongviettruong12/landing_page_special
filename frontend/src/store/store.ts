import { defineStore } from "pinia";
import api from "../api/axios";

export const useUserStore = defineStore("user", {
  state: () => ({
    token: localStorage.getItem("token") as string | null,
  }),
  actions: {
    async login(email: string, password: string) {
      const res = await api.post("/auth/login", { email, password });
      this.token = res.data.token;
      if (this.token) {
        localStorage.setItem("token", this.token);
      }
    },
  },
});
