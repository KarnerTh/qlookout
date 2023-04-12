<script lang="ts">
  export let label: string;
  export let name: string;
  export let placeholder: string = "";
  export let hasError: boolean = false;
  export let required: boolean = false;
</script>

<div class="w-full px-3 mb-6 md:mb-0">
  <label
    class="block uppercase tracking-wide text-gray-700 text-xs font-bold mb-2"
    for={name}
  >
    {label}
  </label>
  <textarea
    class="appearance-none block w-full bg-gray-200 text-gray-700 border rounded py-3 px-4 mb-3 leading-tight focus:outline-none focus:bg-white"
    class:border-red-500={hasError}
    id={name}
    rows="10"
    {name}
    {required}
    {placeholder}
  />
</div>
