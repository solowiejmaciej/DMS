<template>
  <Toast />
  <ConfirmDialog></ConfirmDialog>
  <Dialog
    modal
    :draggable="false"
    v-model:visible="isPhoneNumberDialogVisible"
    header="Confirm phone number"
    :pt="{
      root: 'border-none',
      mask: {
        style: 'backdrop-filter: blur(2px)',
      },
    }"
  >
    <ConfirmPhoneNumber
      :phone-number="enteredPhoneNumber"
      @phone-number-confirmed="handlePhoneNumberConfirmed"
    />
  </Dialog>
  <div class="flex justify-between">
    <Fieldset legend="Personal data">
      <div
        class="flex flex-col gap-2 mx-auto"
        style="min-height: 20rem; max-width: 45rem; min-width: 30rem"
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
              :optionValue="value"
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
      </div>
    </Fieldset>
    <Fieldset legend="Credentials">
      <div class="flex">
        <div class="flex flex-col gap-2 mx-auto" style="min-height: 16rem">
          <div class="field p-fluid m-2">
            <IconField>
              <FloatLabel>
                <InputText
                  id="apiKey"
                  v-model="apiKey"
                  :type="inputType"
                  disabled
                  style="min-width: 40rem"
                />
                <label for="apiKey">API Key</label>
              </FloatLabel>
              <InputIcon>
                <i
                  class="pi pi-copy"
                  @click="copyApiKeyToClipboard"
                  v-tooltip="'Copy API Key to clipboard'"
                />
                <i
                  class="pi pi-eye ml-2"
                  @click="showApiKey"
                  v-tooltip="'Show API Key'"
                  v-if="!isApiKeyVisible"
                />
                <i
                  v-else
                  class="pi pi-eye-slash ml-2"
                  @click="hideApiKey"
                  v-tooltip="'Hide API Key'"
                />
              </InputIcon>
            </IconField>
          </div>
          <div class="field p-fluid m-2">
            <Button
              type="button"
              label="Regenerate API Key"
              @click="confirmRegenerateApiKey()"
            />
          </div>
          <div class="field p-fluid m-2">
            <IconField>
              <FloatLabel>
                <InputNumber
                  v-model="enteredPhoneNumber"
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
            <div v-if="v$.enteredPhoneNumber.$error" class="error-message">
              Phone number is required
            </div>
          </div>
          <div v-if="!isPhoneNumberConfirmed" class="field p-fluid m-2">
            <Button
              type="button"
              label="Confirm Phone Number"
              @click="openPhoneNumberDialog"
            />
          </div>
        </div>
      </div>
    </Fieldset>
  </div>
  <div class="flex justify-end mt-4">
    <Button
      text
      type="button"
      label="Cancel"
      class="w-32 mr-2"
      @click="handleCancel"
    />
    <Button
      icon="pi pi-check"
      type="button"
      label="Save"
      @click="saveData"
      :disabled="v$.value"
      class="w-32"
    />
  </div>
</template>

<script setup>
import {
  ref,
  computed,
  watch,
  defineProps,
  watchEffect,
  defineEmits,
} from "vue";
import { useVuelidate } from "@vuelidate/core";
import { useToast } from "primevue/usetoast";
import { useConfirm } from "primevue/useconfirm";
import ConfirmPhoneNumber from "./ConfirmPhoneNumber.vue";
import apiClient from "../Api/apiClient";
import { useStore } from "vuex";

const store = useStore();

import { required, email, minLength, maxLength } from "@vuelidate/validators";
import countries from "../consts/countries";
import "../styles/flags.css";

let emit = defineEmits(["hide", "userUpdated"]);

const toast = useToast();
const confirm = useConfirm();

const props = defineProps({
  user: {
    type: {
      firstname: String,
      lastname: String,
      email: String,
      password: String,
      phonenumber: String,
      birthday: String,
      country: String,
      isPhoneNumberConfirmed: Boolean,
    },
    required: true,
  },
});

let firstName = ref("");
let lastName = ref("");
let birthDay = ref("");
let userEmail = ref("");
let selectedCountry = ref("");
let apiKey = ref("");
let isApiKeyVisible = ref(false);
let isPhoneNumberDialogVisible = ref(false);
let isPhoneNumberConfirmed = ref(false);
let enteredPhoneNumber = ref("");

let initPhoneNumber = ref(null);

watchEffect(() => {
  if (props.user) {
    firstName.value = props.user.firstname;
    lastName.value = props.user.lastname;
    userEmail.value = props.user.email;
    birthDay.value = props.user.birthDate;
    selectedCountry.value = {
      name: props.user.country,
      code: countries.find((c) => c.name === props.user.country)?.code,
    };
    isPhoneNumberConfirmed.value = props.user.isPhoneNumberConfirmed;
    enteredPhoneNumber.value = props.user.phoneNumber;
    apiKey.value = props.user.dmsApiKey;
    initPhoneNumber.value = props.user.phoneNumber;
  }
});

watch(enteredPhoneNumber, (newVal, oldVal) => {
  if (newVal != oldVal) {
    if (newVal == initPhoneNumber.value) {
      isPhoneNumberConfirmed.value = true;
    } else {
      isPhoneNumberConfirmed.value = false;
    }
  } else {
    isPhoneNumberConfirmed.value = true;
  }
});

const handlePhoneNumberConfirmed = () => {
  isPhoneNumberDialogVisible.value = false;
  isPhoneNumberConfirmed.value = true;
};

const inputType = computed(() => (isApiKeyVisible.value ? "text" : "password"));

const regenerateApiKey = async () => {
  try {
    let data = await apiClient.post("/token/apikey/regenerate");
    apiKey.value = data.data.apiKey;
    toast.add({
      severity: "success",
      summary: "API Key Regenerated",
      life: 3000,
    });
  } catch (error) {
    toast.add({
      severity: "error",
      summary: "Error",
      detail: "Failed to regenerate API key",
      life: 3000,
    });
  }
};

const handleCancel = () => {
  emit("hide");
};

const saveData = async () => {
  let userId = ref(store.state.userId);

  if (!isFormValid()) return;

  let date = new Date(birthDay.value);
  let formattedDate = `${date.getFullYear()}-${String(
    date.getMonth() + 1
  ).padStart(2, "0")}-${String(date.getDate()).padStart(2, "0")}`;
  try {
    let data = {
      firstName: firstName.value,
      lastName: lastName.value,
      email: userEmail.value,
      birthDate: formattedDate,
      country: selectedCountry.value.name,
      phoneNumber: String(enteredPhoneNumber.value),
    };

    let response = await apiClient.put("/user/" + userId.value, data);
    if (response.status === 200) {
      emit("hide");
      emit("userUpdated");
    }
  } catch (error) {
    toast.add({
      severity: "error",
      summary: "Error",
      detail: "Failed to save user data",
      life: 3000,
    });
  }
};

const confirmRegenerateApiKey = () => {
  confirm.require({
    message: "Are you sure you want to proceed? This action cannot be undone.",
    header: "Confirmation",
    icon: "pi pi-exclamation-triangle",
    rejectClass: "p-button-secondary p-button-outlined",
    rejectLabel: "Cancel",
    acceptLabel: "Yes",
    accept: async () => {
      try {
        await regenerateApiKey();
      } catch (error) {
        console.error(error);
        toast.add({
          severity: "error",
          summary: "Error",
          detail: "Failed to regenerate API key",
          life: 3000,
        });
      }
    },
    reject: () => {
      toast.add({
        severity: "error",
        summary: "Rejected",
        detail: "You have rejected",
        life: 3000,
      });
    },
  });
};

const showApiKey = () => {
  isApiKeyVisible.value = true;
};

const hideApiKey = () => {
  isApiKeyVisible.value = false;
};

const copyApiKeyToClipboard = () => {
  navigator.clipboard.writeText(apiKey.value);
  toast.add({
    severity: "success",
    summary: "API Key Copied",
    life: 3000,
  });
};

const openPhoneNumberDialog = () => {
  isPhoneNumberDialogVisible.value = true;
};

const rules = {
  userEmail: { required, email },
  firstName: { required },
  lastName: { required },
  enteredPhoneNumber: { required },
  birthDay: { required },
  selectedCountry: { required },
};

let v$ = useVuelidate(rules, {
  userEmail,
  firstName,
  lastName,
  enteredPhoneNumber,
  birthDay,
  selectedCountry,
});

const isFormValid = () => {
  v$.value.firstName.$touch();
  v$.value.lastName.$touch();
  v$.value.enteredPhoneNumber.$touch();
  v$.value.birthDay.$touch();
  v$.value.selectedCountry.$touch();
  v$.value.userEmail.$touch();
  if (
    !v$.value.firstName.$invalid &&
    !v$.value.lastName.$invalid &&
    !v$.value.enteredPhoneNumber.$invalid &&
    !v$.value.birthDay.$invalid &&
    !v$.value.selectedCountry.$invalid &&
    !v$.value.userEmail.$invalid
  ) {
    return true;
  } else {
    toast.add({
      severity: "error",
      summary: "Error Message",
      detail: "Please fill in the required fields",
      life: 3000,
    });
    return false;
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
