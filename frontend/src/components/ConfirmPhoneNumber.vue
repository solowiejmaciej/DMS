<template>
  <div class="flex flex-col items-center">
    <div class="font-bold text-xl mb-2">{{ props.headerText }}</div>
    <p class="text-gray-500 block mb-5">
      Please enter the code sent to your phone.
    </p>
    <InputOtp v-model="confirmPhoneNumberCode" :length="6"> </InputOtp>
    <div class="flex justify-between w-full mt-5">
      <div class="flex items-center space-x-2">
        <Button
          :disabled="shouldDisableResend"
          label="Resend Code"
          @click="resend2faCode"
          link
          class="p-0"
        ></Button>
        <div v-if="shouldDisableResend">{{ state.countdown }}s</div>
      </div>
    </div>
  </div>
</template>
<script setup>
import { ref, reactive, watch, defineProps, defineEmits } from "vue";
import { useToast } from "primevue/usetoast";
import { onMounted } from "vue";
import apiClient from "../Api/apiClient";

const props = defineProps({
  headerText: String,
});
const emit = defineEmits(["phone-number-confirmed"]);

const toast = useToast();

onMounted(async () => {
  await send2faCode();
});

let shouldDisableResend = ref(false);
let confirmPhoneNumberCode = ref("");

let state = reactive({
  countdown: 60,
  intervalId: null,
});

watch(shouldDisableResend, (newValue) => {
  if (newValue) {
    state.countdown = 60;
    state.intervalId = setInterval(() => {
      if (state.countdown > 0) {
        state.countdown--;
      } else {
        clearInterval(state.intervalId);
        shouldDisableResend.value = false;
      }
    }, 1000);
  } else if (state.intervalId) {
    clearInterval(state.intervalId);
  }
});

watch(confirmPhoneNumberCode, (newValue) => {
  if (newValue) {
    if (newValue.length === 6) {
      submit2faCode();
    }
  }
});

const submit2faCode = async () => {
  try {
    let res = await apiClient.post("/user/confirm-phone", {
      code: confirmPhoneNumberCode.value,
    });

    if (res.status === 200) {
      toast.add({
        severity: "success",
        summary: "Success",
        detail: "Phone number confirmed",
        life: 3000,
      });
      emit("phone-number-confirmed");
    }
  } catch (error) {
    toast.add({
      severity: "error",
      summary: "Error Message",
      detail: "Invalid code",
      life: 3000,
    });
  }
};

const send2faCode = async () => {
  let res = await apiClient.post("/user/send-confirm-phone");
  if (res.status === 200) {
    toast.add({
      severity: "success",
      summary: "Success",
      detail: "Code sent successfully",
      life: 3000,
    });
  } else {
    toast.add({
      severity: "error",
      summary: "Error Message",
      detail: "Failed to send code",
      life: 3000,
    });
  }
};

const resend2faCode = async () => {
  shouldDisableResend.value = true;
  await send2faCode();
};
</script>
