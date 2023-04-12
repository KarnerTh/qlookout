import type { LookoutConfigCreateModel } from "$lib/usecase/lookout/lookoutConfigModel";
import { fail, redirect } from "@sveltejs/kit";
import type { Actions } from "./$types";

export const actions = {
  default: async (event) => {
    const data = await event.request.formData();
    const name = data.get("name")?.toString();
    const cron = data.get("cron")?.toString();
    const query = data.get("query")?.toString();
    const notifyLocal = data.has("notifyLocal");
    const notifyMail = data.has("notifyMail");

    if (!name || !cron || !query) {
      return fail(400, {
        description: "Name, cron and query must not be null",
      });
    }

    const createRequest: LookoutConfigCreateModel = {
      name: name,
      cron: cron,
      query: query,
      notifyLocal: notifyLocal,
      notifyMail: notifyMail,
    };

    // TODO: use api
    console.log("works", createRequest);

    throw redirect(303, "/lookout")
  },
} satisfies Actions;
