<template>
  <form @submit.prevent="handleClick" class="space-y-4">
    <div class="flex items-center space-x-2">
      <i class="pi pi-ellipsis-v text-lg"></i>
      <InputText
        id="placeInTheSystem"
        v-model="placeInTheSystem"
        class="w-full"
        placeholder="Where in the system bug occured?"
      />
    </div>

    <div v-if="v$.placeInTheSystem.$error" class="text-red-500 text-sm mt-1">
      <div v-if="!v$.placeInTheSystem.required.$model">
        Place in system is required
      </div>
    </div>

    <div class="flex items-center space-x-2">
      <i class="pi pi-check-circle text-lg"></i>
      <InputText
        id="expectedBehavior"
        v-model="expectedBehavior"
        class="w-full"
        placeholder="What should happen correctly?"
      />
    </div>
    <div v-if="v$.expectedBehavior.$error" class="text-red-500 text-sm mt-1">
      <div v-if="!v$.expectedBehavior.required.$model">
        Expected behavior is required
      </div>
    </div>
    <div class="flex items-start space-x-2">
      <i class="pi pi-pencil text-lg mt-1"></i>
      <Textarea
        v-model="bugDescription"
        autoResize
        rows="5"
        cols="30"
        class="text-lg w-full"
        placeholder="Bug Description"
      />
    </div>

    <div v-if="v$.bugDescription.$error" class="text-red-500 text-sm mt-1">
      <div v-if="!v$.bugDescription.required.$model">
        Bug Description is required
      </div>
      <div v-if="!v$.bugDescription.minLength.$model">
        Bug Description must be at least 3 characters long
      </div>
      <div v-if="!v$.bugDescription.maxLength.$model">
        Bug Description must be at most 1000 characters long
      </div>
    </div>

    <div class="flex justify-end">
      <Button type="submit" label="Submit" icon="pi pi-check" />
    </div>
  </form>
</template>

<script setup>
import { ref, watch } from "vue";
import { required, minLength, maxLength } from "@vuelidate/validators";
import { useVuelidate } from "@vuelidate/core";
import apiClient from "../../Api/apiClient";

let emit = defineEmits(["bugReported"]);

let placeInTheSystem = ref("");
let expectedBehavior = ref("");
let bugDescription = ref("");

const rules = {
  placeInTheSystem: { required, minLength: minLength(3) },
  expectedBehavior: { required, minLength: minLength(3) },
  bugDescription: {
    required,
    minLength: minLength(3),
    maxLength: maxLength(1000),
  },
};

const v$ = useVuelidate(rules, {
  placeInTheSystem,
  expectedBehavior,
  bugDescription,
});

const handleClick = async () => {
  v$.value.placeInTheSystem.$touch();
  v$.value.expectedBehavior.$touch();
  v$.value.bugDescription.$touch();
  if (v$.value.$error) {
    return;
  }

  let response = await apiClient.post("/bug", {
    placeInTheSystem: placeInTheSystem.value,
    expectedBehavior: expectedBehavior.value,
    bugDescription: bugDescription.value,
  });

  if (response.status === 200) {
    emit("bugReported");
  }
};

watch([placeInTheSystem, expectedBehavior, bugDescription], () => {
  v$.value.$reset();
  v$.value.$validate();
});
</script>
