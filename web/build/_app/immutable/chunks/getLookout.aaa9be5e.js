import{k as e,q as u}from"./svelte-apollo.e94973ad.js";const l=e`
  query Lookout($id: Int!) {
    lookout(id: $id) {
      id
      name
      cron
      query
      notifyLocal
      notifyMail
      rules {
        id
        lookoutId
        columnName
        columnType
        rowIndex
        exactValue
        greaterThan
        lessThan
        shouldBeNull
      }
    }
  }
`,t=o=>u(l,{variables:{id:o}});export{t as u};
