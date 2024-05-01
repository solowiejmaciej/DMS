<template>
  <Dialog
    modal
    :draggable="false"
    @hide="closeDialog"
    dismissableMask
    v-model:visible="isDialogVisible"
    :header="dialogHeader"
    :pt="{
      root: 'border-none',
      mask: {
        style: 'backdrop-filter: blur(2px)',
      },
    }"
  >
    <ManageSoftware v-if="showManageSoftwareDialog"></ManageSoftware>
    <SoftwareUpload
      v-if="showSoftwareUploadDialog"
      @upload-finished="closeDialog"
    ></SoftwareUpload>
  </Dialog>
  <div class="flex flex-col gap-2 mx-auto">
    <div class="flex justify-between">
      <div class="field p-fluid m-2">
        <Button
          @click="openUploadSoftwareDialog"
          type="button"
          label="Upload new software to DMS"
          icon="pi pi-upload"
        >
        </Button>
      </div>
      <div class="field p-fluid m-2">
        <Button
          @click="openManageSoftwareDialog"
          type="button"
          label="Manage software updates"
          icon="pi pi-cog"
        >
        </Button>
      </div>
      <div class="field p-fluid m-2">
        <Button
          :loading="isUpdateLoading"
          @click="sentUpdateCommnad"
          type="button"
          label="Update device"
          icon="pi pi-sync"
        >
        </Button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from "vue";
import ManageSoftware from "./ManageSoftware.vue";
import SoftwareUpload from "./SoftwareUpload.vue";
import { useToast } from "primevue/usetoast";

const toast = useToast();

const showSoftwareUploadDialog = ref(false);
const showManageSoftwareDialog = ref(false);
const isDialogVisible = ref(false);

const dialogHeader = ref("");
const isUpdateLoading = ref(false);

const closeDialog = () => {
  isDialogVisible.value = false;
  showSoftwareUploadDialog.value = false;
  showManageSoftwareDialog.value = false;
};

const openUploadSoftwareDialog = () => {
  isDialogVisible.value = true;
  showSoftwareUploadDialog.value = true;
  dialogHeader.value = "Upload new software";
};

const openManageSoftwareDialog = () => {
  isDialogVisible.value = true;
  showManageSoftwareDialog.value = true;
  dialogHeader.value = "Manage software updates";
};

const sentUpdateCommnad = () => {
  isUpdateLoading.value = true;

  setTimeout(() => {
    isUpdateLoading.value = false;
    toast.add({
      severity: "success",
      summary: "Success",
      detail: "Update command sent",
      life: 3000,
    });
  }, 2000);
};
</script>
