<script lang="ts">
  import { iconData } from "$lib/components/icons";

  export let title: string;
</script>

<div class="flex items-center justify-center">
  <div class="relative inline-block text-left dropdown">
    <span class="rounded-md shadow-sm">
      <button
        class="inline-flex justify-center w-full px-4 py-2 text-sm font-medium text-gray-700 transition duration-150 ease-in-out bg-white border border-gray-300 rounded-md hover:text-gray-500 focus:outline-none active:bg-gray-50 active:text-gray-800"
        type="button"
        aria-haspopup="true"
        aria-expanded="true"
        aria-controls="headlessui-menu-items-117"
      >
        <span>{title}</span>
        <span class="pl-2">
          {@html iconData.chevron_down}
        </span>
      </button>
    </span>
    <div
      class="opacity-0 invisible dropdown-menu transition-all duration-300 transform origin-top-right -translate-y-2 scale-95"
    >
      <div
        class="absolute left-0 w-56 mt-2 origin-top-right bg-white border border-gray-200 divide-y divide-gray-100 rounded-md shadow-lg outline-none"
        aria-labelledby="headlessui-menu-button-1"
        id="headlessui-menu-items-117"
        role="menu"
      >
        <div class="py-1">
          <slot />
        </div>
      </div>
    </div>
  </div>
</div>

<style>
  .dropdown:focus-within .dropdown-menu {
    opacity: 1;
    transform: translate(0) scale(1);
    visibility: visible;
  }
</style>
