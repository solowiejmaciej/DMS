<template>
  <Toast />
  <div class="flex flex-col justify-center items-center">
    <div class="bg-white rounded-full h-auto w-auto flex items-center">
      <img src="/src/assets/dms-logo.png" alt="DMS Logo" class="w-64 h-64" />
    </div>
    <div class="text-3xl font-bold mb-4">Login</div>
    <div class="card w-1/2">
      <div
        class="flex flex-col gap-2 mx-auto"
        style="min-height: 16rem; max-width: 20rem"
      >
        <div class="field p-fluid m-2">
          <IconField>
            <FloatLabel>
              <InputText id="userEmail" v-model="userEmail" />
              <label for="userEmail">Email</label>
            </FloatLabel>
            <InputIcon>
              <i class="pi pi-envelope" />
            </InputIcon>
          </IconField>
          <div v-if="v$.userEmail.$error" class="error-message">
            Email is required
          </div>
        </div>
        <div class="field p-fluid m-2">
          <IconField>
            <FloatLabel>
              <Password
                id="password"
                v-model="userPassword"
                toggleMask
                :feedback="false"
              />
              <label for="password">Password</label>
            </FloatLabel>
            <div v-if="v$.userPassword.$error" class="error-message">
              Password is required and must be between 6 and 20 characters
            </div>
          </IconField>
        </div>
      </div>
      <div class="flex pt-4 justify-between">
        <p class="text-gray-500 block mb-5">
          Don't have an account?
          <router-link to="/register">Register</router-link>
        </p>
        <Button
          label="Next"
          icon="pi pi-arrow-right"
          iconPos="right"
          @click="nextPage"
        />
      </div>
    </div>
  </div>
</template>
<script setup>
import { ref } from "vue";
import { useToast } from "primevue/usetoast";
import { useRouter } from "vue-router";
import { useVuelidate } from "@vuelidate/core";
import { required, email, minLength, maxLength } from "@vuelidate/validators";
import { useStore } from "vuex";

const store = useStore();
const toast = useToast();
const router = useRouter();

let userEmail = ref("");
let userPassword = ref("");

const redirectToHome = () => {
  //set jwt token to local storage
  router.push("/home");
};

const rules = {
  userEmail: { required, email },
  userPassword: { required, minLength: minLength(6), maxLength: maxLength(20) },
};

let v$ = useVuelidate(rules, {
  userEmail,
  userPassword,
});

const handleLogin = () => {
  store
    .dispatch("login", {
      email: userEmail.value,
      password: userPassword.value,
    })
    .then(() => {
      redirectToHome();
    })
    .catch((err) => {
      toast.add({
        severity: "error",
        summary: "Login Error",
        detail: err.message,
        life: 3000,
      });
    });
};

const nextPage = () => {
  v$.value.userEmail.$touch();
  v$.value.userPassword.$touch();
  if (!v$.value.userEmail.$invalid && !v$.value.userPassword.$invalid) {
    handleLogin();
  } else {
    toast.add({
      severity: "error",
      summary: "Error Message",
      detail: "Please fill in the required fields",
      life: 3000,
    });
  }
};
</script>

<style scoped>
.error-message {
  color: #cd4141; /* red color */
  font-size: 0.875rem; /* font size */
  margin-top: 0.25rem; /* space above the error message */
  display: block; /* use flexbox */
}
</style>
