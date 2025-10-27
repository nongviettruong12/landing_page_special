<template>
  <div class="login-container">
    <el-card class="login-card">
      <h2 class="title">Đăng nhập</h2>

      <el-form :model="form" :rules="rules" ref="loginForm" label-position="top">
        <el-form-item label="Email" prop="email">
          <el-input v-model="form.email" placeholder="Nhập email" />
        </el-form-item>

        <el-form-item label="Mật khẩu" prop="password">
          <el-input
            v-model="form.password"
            type="password"
            placeholder="Nhập mật khẩu"
          />
        </el-form-item>

        <el-form-item>
          <el-button
            type="primary"
            :loading="loading"
            style="width: 100%"
            @click="handleLogin"
          >
            Đăng nhập
          </el-button>
        </el-form-item>
      </el-form>

      <div class="extra">
        <span>Bạn chưa có tài khoản?</span>
        <router-link to="/register">Đăng ký ngay</router-link>
      </div>
    </el-card>
  </div>
</template>

<script setup>
import { ref } from "vue";
import { ElMessage } from "element-plus";
import { useRouter } from "vue-router";
import { useUserStore } from "../store/user.js";

const router = useRouter();
const userStore = useUserStore();

const form = ref({
  email: "",
  password: "",
});

const rules = {
  email: [{ required: true, message: "Vui lòng nhập email", trigger: "blur" }],
  password: [
    { required: true, message: "Vui lòng nhập mật khẩu", trigger: "blur" },
  ],
};

const loading = ref(false);
const loginForm = ref(null);

const handleLogin = async () => {
  await loginForm.value.validate(async (valid) => {
    if (!valid) return;
    loading.value = true;
    try {
      await userStore.login(form.value.email, form.value.password);
      ElMessage.success("Đăng nhập thành công!");
      router.push("/");
    } catch (error) {
      ElMessage.error("Sai thông tin đăng nhập!");
    } finally {
      loading.value = false;
    }
  });
};
</script>

<style scoped>
.login-container {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
  background: #f5f6fa;
}

.login-card {
  width: 380px;
  padding: 20px;
  border-radius: 12px;
  box-shadow: 0 3px 10px rgba(0, 0, 0, 0.1);
}

.title {
  text-align: center;
  margin-bottom: 20px;
  color: #333;
}

.extra {
  text-align: center;
  margin-top: 10px;
}

.extra a {
  color: #409eff;
  margin-left: 5px;
}
</style>
