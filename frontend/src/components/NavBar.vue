<template>
  <Sidebar
    v-model:visible="showNotifications"
    position="left"
    :baseZIndex="10000"
  >
    <Notifications />
  </Sidebar>
  <Menubar :model="items">
    <template #start>
      <div
        class="bg-white rounded-full h-auto w-auto flex items-center justify-center"
      >
        <img src="/src/assets/dms-logo.png" alt="DMS Logo" class="w-20 h-20" />
      </div>
    </template>
    <template #item="{ item, props }">
      <a
        :class="{ 'active-menu': item.isActive }"
        v-ripple
        class="flex align-items-center"
        v-bind="props.action"
        @click="item.onClick"
      >
        <span :class="item.icon" />
        <span class="ml-2" :class="{ 'font-bold': item.isActive }">{{
          item.label
        }}</span>
      </a>
    </template>
    <template #end>
      <div
        class="flex flex-col items-center justify-center cursor-pointer"
        @click="openUserMenu"
        aria-haspopup="true"
        aria-controls="overlay_menu"
      >
        <Avatar shape="circle" size="large" :label="label" />
        <p>
          <strong>{{ user.firstname + " " + user.lastname }}</strong>
        </p>
        <Menu ref="menu" id="overlay_menu" :model="userMenuItems" :popup="true">
          <template #item="{ item, props }">
            <a v-ripple class="flex align-items-center" v-bind="props.action">
              <span :class="item.icon" />
              <span class="ml-2">{{ item.label }}</span>
              <Badge v-if="item.badge" class="ml-auto" :value="item.badge" />
              <span
                v-if="item.shortcut"
                class="ml-auto border-1 surface-border border-round surface-100 text-xs p-1"
                >{{ item.shortcut }}</span
              >
            </a>
          </template>
        </Menu>

        <Dialog
          modal
          dismissableMask
          :draggable="false"
          v-model:visible="showDialog"
          :header="showDialogHeader"
          @hide="handleDialogHide"
          :pt="{
            root: 'border-none',
            mask: {
              style: 'backdrop-filter: blur(2px)',
            },
          }"
        >
          <UserData
            v-if="showUserData"
            :user="user"
            @userUpdated="handleUserUpdated"
            @hide="handleDialogHide"
          />
          <ReportBug v-if="showReportBug" @bugReported="handleBugReported" />
        </Dialog>
      </div>
    </template>
  </Menubar>
</template>

<script setup>
import { ref, onMounted } from "vue";
import apiClient from "../Api/apiClient";
import { useRouter } from "vue-router";
import UserData from "./UserData.vue";
import Notifications from "./Notifications/Notifications.vue";
import ReportBug from "./ReportBug/ReportBug.vue";
import { useStore } from "vuex";
import { useToast } from "primevue/usetoast";

const toast = useToast();
const store = useStore();
const router = useRouter();

let showDialog = ref(false);
let showUserData = ref(false);
let showNotifications = ref(false);
let showReportBug = ref(false);
let user = ref("");

let label = ref("");

let showDialogHeader = ref("");

const getUser = async () => {
  try {
    const response = await apiClient.post("/me");
    user.value = response.data.user;
    label.value = user.value.firstname[0] + user.value.lastname[0];
  } catch (error) {
    console.error(error);
  }
};

onMounted(async () => {
  await getUser();
});

const getNotificationsCount = () => {
  return "5";
};

const menu = ref();
const userMenuItems = ref([
  {
    label: "Edit profile",
    icon: "pi pi-user",
    command: () => {
      handleOpenEditUser();
    },
  },
  {
    label: "Notifications",
    badge: getNotificationsCount(),
    icon: "pi pi-bell",
    command: () => {
      handleOpenNotifications();
    },
  },
  {
    label: "Become supporter",
    icon: "pi pi-heart",
    command: () => {
      handleBecomeSuporter();
    },
  },
  {
    label: "Report a bug",
    icon: "pi pi-hammer",
    command: () => {
      handleReportBug();
    },
  },
  {
    label: "Logout",
    icon: "pi pi-sign-out",
    command: () => {
      handleLogout();
    },
  },
]);

const emit = defineEmits(["open-devices", "open-dashboard"]);

const items = ref([
  {
    label: "Dashboard",
    icon: "pi pi-chart-bar",
    isActive: true,
    onClick: () => {
      openDashboard();
    },
  },
  {
    label: "Devices",
    icon: "pi pi-microchip",
    isActive: false,
    onClick: () => {
      openDevices();
    },
  },
]);

const openDevices = () => {
  items.value.forEach((item) => {
    if (item.label === "Devices") {
      item.isActive = true;
    } else {
      item.isActive = false;
    }
  });
  emit("open-devices");
};

const openDashboard = () => {
  items.value.forEach((item) => {
    if (item.label === "Dashboard") {
      item.isActive = true;
    } else {
      item.isActive = false;
    }
  });
  emit("open-dashboard");
};

const openUserMenu = (event) => {
  menu.value.toggle(event);
};

const handleLogout = async () => {
  await store.dispatch("logout");
};

const handleOpenNotifications = () => {
  showNotifications.value = true;
};

const handleOpenEditUser = () => {
  showDialog.value = true;
  showUserData.value = true;
  showDialogHeader.value = "Edit User";
  showNotifications.value = false;
  showReportBug.value = false;
};

const handleBecomeSuporter = () => {
  toast.add({
    severity: "info",
    summary: "Info",
    detail: "This feature is not available yet",
    life: 3000,
  });
};

const handleReportBug = () => {
  showDialog.value = true;
  showReportBug.value = true;
  showDialogHeader.value = "Report a bug";
  showUserData.value = false;
  showNotifications.value = false;
};

const handleBugReported = () => {
  handleDialogHide();
  toast.add({
    severity: "success",
    summary: "Success",
    detail: "Bug reported successfully",
    life: 3000,
  });
};

const handleDialogHide = () => {
  showDialog.value = false;
  showUserData.value = false;
};

const handleUserUpdated = async () => {
  toast.add({
    severity: "success",
    summary: "Success",
    detail: "User data saved successfully",
    life: 3000,
  });
  await getUser();
};
</script>

<style scoped>
.active-menu {
  color: var(--primary-color);
}
</style>
