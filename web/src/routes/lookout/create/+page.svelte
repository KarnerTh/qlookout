<script lang="ts">
  import FormInputCheckboxGroup from "$lib/components/form/FormInputCheckboxGroup.svelte";
  import FormInputField from "$lib/components/form/FormInputField.svelte";
  import FormInputTextArea from "$lib/components/form/FormInputTextArea.svelte";
  import PageHeader from "$lib/components/header/PageHeader.svelte";
  import { enhance } from "$app/forms";
  import type { ActionData } from "./$types";
  import Button from "$lib/components/button/Button.svelte";

  export let form: ActionData;
</script>

<PageHeader title="Create Lookout" showBackButton />

<form method="POST" class="w-full max-w-lg" use:enhance>
  <div class="flex flex-wrap -mx-3 mb-6">
    <FormInputField
      name="name"
      required
      label="Lookout name"
      placeholder="My important query"
    />
    <FormInputField
      name="cron"
      required
      label="Cron expression"
      placeholder="0 7 * * 1-5"
    />
    <FormInputTextArea
      name="query"
      required
      label="SQL Query"
      placeholder="SELECT COUNT(*) FROM something"
    />

    <FormInputCheckboxGroup
      label="Notifications"
      items={[
        {
          name: "notifyLocal",
          title: "Local notification",
          description: "Only works if run locally",
        },
        {
          name: "notifyMail",
          title: "Mail notification",
        },
      ]}
    />

    <div class="mt-4 mx-3">
      {#if form}
        <p class="text-red-500 text-xs italic">{form.description}</p>
      {/if}

      <Button title="Create" type="submit" leadingIcon="check" />
    </div>
  </div>
</form>
