<script lang="ts">
  import { goto } from "$app/navigation";
  import { page } from "$app/stores";
  import Button from "$lib/components/button/Button.svelte";
  import FormInputCheckboxGroup from "$lib/components/form/FormInputCheckboxGroup.svelte";
  import FormInputField from "$lib/components/form/FormInputField.svelte";
  import FormInputTextArea from "$lib/components/form/FormInputTextArea.svelte";
  import PageHeader from "$lib/components/header/PageHeader.svelte";
  import LoadingSpinner from "$lib/components/loading/LoadingSpinner.svelte";
  import { useLookoutUpdateMutation } from "$lib/usecase/lookout/mutation/updateLookout";
  import { useLookout } from "$lib/usecase/lookout/query/getLookout";

  const lookout = useLookout(+$page.params.lookoutId);
  const updateLookout = useLookoutUpdateMutation();

  const onSubmit = async (event: SubmitEvent) => {
    const data = new FormData(event.target as HTMLFormElement);
    const name = data.get("name")?.toString();
    const cron = data.get("cron")?.toString();
    const query = data.get("query")?.toString();
    const notifyLocal = data.has("notifyLocal");
    const notifyMail = data.has("notifyMail");

    const result = await updateLookout({
      variables: {
        id: +$page.params.lookoutId,
        name,
        cron,
        query,
        notifyMail,
        notifyLocal,
      },
    });

    if (result.errors) {
      alert("Something went wrong");
      return;
    }

    goto(`/lookout/${$page.params.lookoutId}`, { state: { refetch: true } });
  };
</script>

<PageHeader
  title={"Update Lookout"}
  backAction={() => goto(`/lookout/${$page.params.lookoutId}`)}
/>

{#if $lookout.loading}
  <div class="flex justify-center pt-9">
    <LoadingSpinner />
  </div>
{:else if $lookout.data}
  <form
    method="POST"
    class="w-full max-w-lg"
    on:submit|preventDefault={onSubmit}
  >
    <div class="flex flex-wrap -mx-3 mb-6">
      <FormInputField
        name="name"
        value={$lookout.data.lookout.name}
        required
        label="Lookout name"
        placeholder="My important query"
      />
      <FormInputField
        name="cron"
        required
        value={$lookout.data.lookout.cron}
        label="Cron expression"
        placeholder="0 7 * * 1-5"
      />
      <FormInputTextArea
        name="query"
        required
        value={$lookout.data.lookout.query}
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
            initialValue: $lookout.data.lookout.notifyLocal,
          },
          {
            name: "notifyMail",
            title: "Mail notification",
            initialValue: $lookout.data.lookout.notifyMail,
          },
        ]}
      />

      <div class="mt-4 mx-3 w-full flex flex-row-reverse">
        <Button title="Update" type="submit" leadingIcon="check" />
      </div>
    </div>
  </form>
{/if}
