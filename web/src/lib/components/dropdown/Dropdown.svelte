<script lang="ts">
  import Button from "../button/Button.svelte";

  export let title: string;
</script>

<div class="flex items-center justify-center">
  <div class="relative inline-block text-left dropdown">
    <Button {title} trailingIcon="chevron_down" />
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
