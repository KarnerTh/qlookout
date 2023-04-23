<script lang="ts">
  import { afterNavigate, goto } from "$app/navigation";
  import { page } from "$app/stores";
  import Button from "$lib/components/button/Button.svelte";
  import DialogButton from "$lib/components/dialog/DialogButton.svelte";
  import PageHeader from "$lib/components/header/PageHeader.svelte";
  import LoadingSpinner from "$lib/components/loading/LoadingSpinner.svelte";
  import { useLookoutDeleteMutation } from "$lib/usecase/lookout/mutation/deleteLookout";
  import { useLookout } from "$lib/usecase/lookout/query/getLookout";
  import LookoutBaseInfo from "./LookoutBaseInfo.svelte";
  import LookoutQueryInfo from "./LookoutQueryInfo.svelte";
  import LookoutRuleInfo from "./LookoutRuleInfo.svelte";

  const lookoutId = +$page.params.lookoutId;
  const lookout = useLookout(lookoutId);
  const deleteLookout = useLookoutDeleteMutation();

  afterNavigate(() => {
    if (history.state.refetch) {
      lookout.refetch();
    }
  });

  const onDelete = async () => {
    const result = await deleteLookout({ variables: { id: lookoutId } });
    if (result.errors) {
      alert("Something went wrong");
      return;
    }

    goto("/lookout", { state: { refetch: true } });
  };
</script>

<PageHeader title={"Lookout Detail"} backAction={() => goto("/lookout")} />

{#if $lookout.loading}
  <div class="flex justify-center pt-9">
    <LoadingSpinner />
  </div>
{:else if $lookout.data}
  <div class="mb-2 flex gap-2">
    <Button
      title="Edit"
      leadingIcon="pencil"
      on:click={() => goto(`/lookout/${$page.params.lookoutId}/edit`)}
    />
    <DialogButton
      buttonText="Delete"
      buttonIcon="trash"
      buttonType="red"
      dialogTitle="Delete Lookout"
      dialogDescription="Do you realy want to delete this lookout? This action cannot be undone and all related data will be lost"
      actionButtons={[
        {
          title: "Cancel",
          onClick() {},
        },
        {
          title: "Delete",
          buttonType: "red",
          onClick() {
            onDelete();
          },
        },
      ]}
    />
  </div>

  <div class="flex flex-wrap gap-2">
    <div class="max-w-sm">
      <LookoutBaseInfo lookout={$lookout.data.lookout} />
    </div>

    <div class="grow">
      <LookoutQueryInfo query={$lookout.data.lookout.query} />
    </div>
  </div>

  <div class="mt-2">
    <LookoutRuleInfo rules={$lookout.data.lookout.rules} />
  </div>
{/if}
