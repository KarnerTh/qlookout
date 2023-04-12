<script lang="ts">
  import { iconData, type Icon } from "../icons";

  export let icon: Icon;
</script>

<span class="rounded-md">
  <button
    class="flex justify-center items-center w-full text-sm font-medium text-gray-700 hover:text-gray-500 focus:outline-none active:text-gray-800"
    type="button"
    on:click
  >
    {@html iconData[icon]}
  </button>
</span>
