<template>
  <Toast />
  <div class="flex flex-col justify-center items-center min-h-screen mt-0">
    <div
      class="bg-white rounded-full h-auto w-auto flex items-center justify-center mb-4"
    >
      <img src="/src/assets/dms-logo.png" alt="DMS Logo" class="w-64 h-64" />
    </div>
    <div class="text-3xl font-bold mb-4">Register</div>
    <div class="card w-1/2">
      <div
        class="rounded border-2 w-12 h-12 inline-flex items-center justify-center"
      ></div>
      <Stepper linear v-model:activeStep="active">
        <StepperPanel>
          <template #header="{ clickCallback }">
            <button
              class="bg-transparent border-none inline-flex flex-col gap-2"
              @click="clickCallback"
            >
              <span
                :class="'rounded border-2 w-12 h-12 inline-flex items-center justify-center'"
              >
                <i class="pi pi-at" style="font-size: 1.8rem"></i>
              </span>
            </button>
          </template>
          <template #content="">
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
                  Email is required and must be valid and unique
                </div>
              </div>
              <div class="field p-fluid m-2">
                <IconField>
                  <FloatLabel>
                    <Password id="password" v-model="userPassword" toggleMask />
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
                Already have an account?
                <router-link to="/login">Login</router-link>
              </p>
              <Button
                label="Next"
                icon="pi pi-arrow-right"
                iconPos="right"
                @click="nextPage"
              />
            </div>
          </template>
        </StepperPanel>
        <StepperPanel>
          <template #header="{ index, clickCallback }">
            <button
              class="bg-transparent border-none inline-flex flex-col gap-2"
              @click="clickCallback"
            >
              <span
                :class="'rounded border-2 w-12 h-12 inline-flex items-center justify-center'"
              >
                <i class="pi pi-user" style="font-size: 1.8rem"></i>
              </span>
            </button>
          </template>
          <template #content="{ prevCallback, nextCallback }">
            <div
              class="flex flex-col gap-2 mx-auto"
              style="min-height: 16rem; max-width: 20rem"
            >
              <div class="field p-fluid m-2">
                <IconField>
                  <FloatLabel>
                    <InputText id="firstName" v-model="firstName" />
                    <label for="firstName">First name</label>
                  </FloatLabel>
                  <InputIcon>
                    <InputIcon>
                      <i class="pi pi-address-book" />
                    </InputIcon>
                  </InputIcon>
                </IconField>
                <div v-if="v$.firstName.$error" class="error-message">
                  First name is required
                </div>
              </div>
              <div class="field p-fluid m-2">
                <IconField>
                  <FloatLabel>
                    <InputText id="lastName" v-model="lastName" />
                    <label for="lastName">Last name</label>
                  </FloatLabel>
                  <InputIcon>
                    <InputIcon>
                      <i class="pi pi-address-book" />
                    </InputIcon>
                  </InputIcon>
                </IconField>
                <div v-if="v$.lastName.$error" class="error-message">
                  Last name is required
                </div>
              </div>
              <div class="field p-fluid m-2">
                <IconField>
                  <FloatLabel>
                    <InputText
                      v-model="phoneNumber"
                      inputId="phoneNumber"
                      :useGrouping="false"
                    />
                    <label for="phoneNumber">Phone number</label>
                  </FloatLabel>
                  <InputIcon>
                    <InputIcon>
                      <i class="pi pi-phone" />
                    </InputIcon>
                  </InputIcon>
                </IconField>
                <div v-if="v$.phoneNumber.$error" class="error-message">
                  {{ v$.phoneNumber.$message }}
                </div>
              </div>
              <div class="field p-fluid m-2">
                <FloatLabel>
                  <Calendar
                    id="birthDay"
                    v-model="birthDay"
                    showIcon
                    iconDisplay="input"
                    :showOnFocus="true"
                    dateFormat="yy-mm-dd"
                  />
                  <label for="birthDay">Birthday</label>
                </FloatLabel>
                <div v-if="v$.birthDay.$error" class="error-message">
                  Birthday is required
                </div>
              </div>
              <div class="field p-fluid m-2">
                <FloatLabel>
                  <Dropdown
                    id="country"
                    showClear
                    v-model="selectedCountry"
                    :options="countries"
                    filter
                    optionLabel="name"
                    class="w-full md:w-14rem"
                  >
                    <template #value="slotProps">
                      <img
                        v-if="slotProps.value"
                        :alt="slotProps.value.label"
                        src="https://primefaces.org/cdn/primevue/images/flag/flag_placeholder.png"
                        :class="`me-2 flag flag-${slotProps.value.code?.toLowerCase()}`"
                        style="width: 18px; display: inline-block"
                      />
                      <div v-if="slotProps.value" style="display: inline-block">
                        {{ slotProps.value.name }}
                      </div>
                    </template>
                    <template #option="slotProps">
                      <img
                        :alt="slotProps.option.label"
                        src="https://primefaces.org/cdn/primevue/images/flag/flag_placeholder.png"
                        :class="`me-2 flag flag-${slotProps.option.code?.toLowerCase()}`"
                        style="width: 18px"
                      />
                      <div>{{ slotProps.option.name }}</div>
                    </template>
                  </Dropdown>
                  <label for="country">Country of origin</label>
                </FloatLabel>
                <div v-if="v$.selectedCountry.$error" class="error-message">
                  Country is required
                </div>
              </div>
              <div class="field p-fluid">
                <div class="flex items-center space-x-2">
                  <p class="text-lg font-semibold">Agree to</p>
                  <router-link to="/terms" class="text-cyan-500 hover:underline"
                    >terms and conditions</router-link
                  >
                  <Checkbox
                    v-model="termsAgreed"
                    :binary="true"
                    class="text-xl"
                  />
                </div>
                <div v-if="v$.termsAgreed.$error" class="error-message">
                  You must agree to the terms and conditions
                </div>
              </div>
            </div>
            <div class="flex pt-4 justify-between">
              <Button
                label="Back"
                severity="secondary"
                icon="pi pi-arrow-left"
                @click="prevCallback"
              />
              <Button
                label="Next"
                icon="pi pi-arrow-right"
                iconPos="right"
                @click="nextPage"
              />
            </div>
          </template>
        </StepperPanel>
        <StepperPanel>
          <template #header="{ index, clickCallback }">
            <button
              class="bg-transparent border-none inline-flex flex-col gap-2"
              @click="clickCallback"
            >
              <span
                :class="'rounded border-2 w-12 h-12 inline-flex items-center justify-center'"
              >
                <i class="pi pi-lock" style="font-size: 1.8rem"></i>
              </span>
            </button>
          </template>
          <template #content="{ prevCallback, nextCallback }">
            <ConfirmPhoneNumber
              v-if="shouldShowConfirm && active === 2"
              headerText="Authenticate Your Account"
              @phone-number-confirmed="handlePhoneNumberConfirmed"
            />
            <div class="flex pt-4 justify-between">
              <Button
                label="Back"
                severity="secondary"
                icon="pi pi-arrow-left"
                @click="prevCallback"
              />
              <Button
                label="Next"
                icon="pi pi-arrow-right"
                iconPos="right"
                @click="nextPage"
                :disabled="shouldDisableNext"
              />
            </div>
          </template>
        </StepperPanel>
        <StepperPanel>
          <template #header="{ index, clickCallback }">
            <button
              class="bg-transparent border-none inline-flex flex-col gap-2"
              @click="clickCallback"
            >
              <span
                :class="'rounded border-2 w-12 h-12 inline-flex items-center justify-center'"
              >
                <i class="pi pi-check" style="font-size: 1.8rem"></i>
              </span>
            </button>
          </template>
          <template #content="{ prevCallback }">
            <div class="flex flex-col items-center">
              <div class="font-bold text-xl mb-2">Registration Complete</div>
              <p class="text-white-500 block mb-5">
                You have successfully registered.
              </p>
            </div>
            <div class="flex pt-4 justify-center">
              <Button
                label="Finish"
                icon="pi pi-check"
                iconPos="right"
                @click="redirectToHome"
              >
              </Button>
            </div>
          </template>
        </StepperPanel>
      </Stepper>
    </div>
  </div>
</template>
<script setup>
import "../../styles/flags.css";
import countries from "../../consts/countries";
import ConfirmPhoneNumber from "../../components/ConfirmPhoneNumber.vue";

import { ref } from "vue";
import { useToast } from "primevue/usetoast";
import { useRouter } from "vue-router";
import { useVuelidate } from "@vuelidate/core";
import { required, email, minLength, maxLength } from "@vuelidate/validators";
import apiClient from "../../Api/apiClient";
import { useStore } from "vuex";

const store = useStore();

const toast = useToast();
const router = useRouter();

let active = ref(0);
let firstName = ref("");
let lastName = ref("");
let phoneNumber = ref("");
let birthDay = ref("");
let termsAgreed = ref();
let userEmail = ref("");
let userPassword = ref("");
let selectedCountry = ref("");
let shouldShowConfirm = ref(false);

let shouldDisableNext = ref(true);

const redirectToHome = () => {
  router.push("/home");
};

const isChecked = (value) => value === true;

const handlePhoneNumberConfirmed = () => {
  shouldDisableNext.value = false;
};

const rules = {
  userEmail: { required, email },
  userPassword: { required, minLength: minLength(6), maxLength: maxLength(20) },
  firstName: { required },
  lastName: { required },
  phoneNumber: { required },
  birthDay: { required },
  selectedCountry: { required },
  termsAgreed: { isChecked },
};

let v$ = useVuelidate(rules, {
  userEmail,
  userPassword,
  firstName,
  lastName,
  phoneNumber,
  birthDay,
  selectedCountry,
  termsAgreed,
});

const nextPage = () => {
  if (active.value === 0) {
    validateEmailAndPassword();
  } else if (active.value === 1) {
    validatePersonalInfo();
    handleRegisterUser();
  } else {
    active.value++;
  }
};

const validateEmailAndPassword = () => {
  v$.value.userEmail.$touch();
  v$.value.userPassword.$touch();
  if (!v$.value.userEmail.$invalid && !v$.value.userPassword.$invalid) {
    active.value++;
  } else {
    toast.add({
      severity: "error",
      summary: "Error Message",
      detail: "Please fill in the required fields",
      life: 3000,
    });
  }
};

const validatePersonalInfo = () => {
  v$.value.firstName.$touch();
  v$.value.lastName.$touch();
  v$.value.phoneNumber.$touch();
  v$.value.birthDay.$touch();
  v$.value.selectedCountry.$touch();
  v$.value.termsAgreed.$touch();
  if (
    !v$.value.firstName.$invalid &&
    !v$.value.lastName.$invalid &&
    !v$.value.phoneNumber.$invalid &&
    !v$.value.birthDay.$invalid &&
    !v$.value.selectedCountry.$invalid &&
    !v$.value.termsAgreed.$invalid
  ) {
    active.value++;
  } else {
    toast.add({
      severity: "error",
      summary: "Error Message",
      detail: "Please fill in the required fields",
      life: 3000,
    });
  }
};

const handleRegisterUser = async () => {
  let date = new Date(birthDay.value);
  let formattedDate = `${date.getFullYear()}-${String(
    date.getMonth() + 1
  ).padStart(2, "0")}-${String(date.getDate()).padStart(2, "0")}`;

  try {
    let res = await apiClient.post("/user", {
      email: userEmail.value,
      password: userPassword.value,
      firstname: firstName.value,
      lastname: lastName.value,
      phoneNumber: phoneNumber.value,
      birthDate: formattedDate,
      country: selectedCountry.value.name,
    });

    if (res.status === 201) {
      await store.dispatch("login", {
        email: userEmail.value,
        password: userPassword.value,
      });
      shouldShowConfirm.value = true;
      toast.add({
        severity: "success",
        summary: "Success Message",
        detail: "User registered successfully",
        life: 3000,
      });
    }
  } catch {
    toast.add({
      severity: "error",
      summary: "Error Message",
      detail: "Failed to register user",
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
