<script lang="ts">
  import FormInputCheckboxGroup from "$lib/components/form/FormInputCheckboxGroup.svelte";
  import FormInputField from "$lib/components/form/FormInputField.svelte";
  import FormInputTextArea from "$lib/components/form/FormInputTextArea.svelte";
  import PageHeader from "$lib/components/header/PageHeader.svelte";
  import Button from "$lib/components/button/Button.svelte";
  import { createLookoutMutation } from "$lib/usecase/lookout/query/createLookout";
  import { goto } from "$app/navigation";

  const createLookout = createLookoutMutation();

  const onSubmit = async (event: any) => {
    const data = new FormData(event.target);
    const name = data.get("name")?.toString();
    const cron = data.get("cron")?.toString();
    const query = data.get("query")?.toString();
    const notifyLocal = data.has("notifyLocal");
    const notifyMail = data.has("notifyMail");

    if (!name || !cron || !query) {
      alert("Name, cron and query must not be null");
      return;
    }

    const result = await createLookout({
      variables: { name, cron, query, notifyMail, notifyLocal },
    });

    if (result.errors) {
      alert("Something went wrong");
      return;
    }

    goto("/lookout", { state: { refetch: true } });
  };
</script>

<PageHeader title="Create Lookout" showBackButton />

<form method="POST" class="w-full max-w-lg" on:submit|preventDefault={onSubmit}>
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
      <Button title="Create" type="submit" leadingIcon="check" />
    </div>
  </div>
</form>
