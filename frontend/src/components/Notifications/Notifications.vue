<template>
  <Accordion :activeIndex="0">
    <AccordionTab
      v-for="notification in notifications"
      :key="notification.header"
      :icon="'pi pi-github'"
    >
      <template #header>
        <span class="flex align-items-center gap-2 w-full">
          <i :class="getIcon(notification.severity)" />
          <span class="font-bold white-space-nowrap">{{
            notification.header
          }}</span>
        </span>
      </template>
      <p class="m-0" v-html="notification.text"></p>
    </AccordionTab>
  </Accordion>
</template>

<script setup>
import { ref } from "vue";
let notifications = ref([
  {
    header: "Welcome to DMS!",
    text: "<strong>Thank you for using our software!</strong>",
    severity: "INFO",
  },
  {
    header: "Warning!",
    text: "<strong>Devices are not loading correclty, sorry </strong>",
    severity: "WARNING",
  },
  // More notifications here
]);

const getIcon = (severity) => {
  switch (severity) {
    case "INFO":
      return "pi pi-info-circle";
    case "WARNING":
      return "pi pi-exclamation-triangle";
    case "DANGER":
      return "pi pi-times-circle";
    default:
      return "";
  }
};
</script>
