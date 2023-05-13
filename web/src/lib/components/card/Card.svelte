<script lang="ts">
  export let cursor: boolean = false;
</script>

<!-- svelte-ignore a11y-click-events-have-key-events -->
<div
  class="block rounded-lg bg-white p-6 shadow dark:bg-neutral-700"
  class:cursor-pointer={cursor}
  on:click
>
  <slot />
</div>
